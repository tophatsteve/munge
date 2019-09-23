package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/tophatsteve/munge/capitalise"
	"github.com/tophatsteve/munge/reverse"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var reverseHost string
var reversePort string
var capitaliseHost string
var capitalisePort string
var serverPort string

func init() {
	reverseHost = os.Getenv("REVERSE_SERVICE_HOST")
	reversePort = os.Getenv("REVERSE_SERVICE_PORT")
	capitaliseHost = os.Getenv("CAPITALISE_SERVICE_HOST")
	capitalisePort = os.Getenv("CAPITALISE_SERVICE_PORT")
	serverPort = os.Getenv("PORT")
}

func main() {

	if serverPort == "" {
		serverPort = "80"
	}

	println("Running on port:", serverPort)
	println("Reverse running on host:", reverseHost)
	println("Reverse running on port:", reversePort)
	println("Capitalise running on host:", capitaliseHost)
	println("Capitalise running on port:", capitalisePort)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	router := httprouter.New()
	router.GET("/", rootHandler)
	router.GET("/:text", textHandler)

	httpServer := &http.Server{Addr: ":" + serverPort, Handler: router}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-stop
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	httpServer.Shutdown(ctx)
}

func reverseText(text string) (string, error) {

	conn, err := grpc.Dial(
		reverseHost+":"+reversePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := reverse.NewReverseClient(conn)
	ctx := context.Background()
	reqReverse := reverse.ReverseRequest{
		Text: text,
	}
	respReverse, err := client.Reverse(ctx, &reqReverse)

	if err != nil {
		return "", err
	}

	return respReverse.Text, nil
}

func capitaliseText(text string) (string, error) {

	conn, err := grpc.Dial(
		capitaliseHost+":"+capitalisePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := capitalise.NewCapitaliseClient(conn)
	ctx := context.Background()
	reqCapitalise := capitalise.CapitaliseRequest{
		Text: text,
	}
	respCapitalise, err := client.Capitalise(ctx, &reqCapitalise)

	if err != nil {
		return "", err
	}

	return respCapitalise.Text, nil
}

func textHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	text := p.ByName("text")
	reversedText, err := reverseText(text)

	if err != nil {
		respondErr(w, r, err, http.StatusBadRequest)
	}

	capitalisedText, err := capitaliseText(reversedText)

	if err != nil {
		respondErr(w, r, err, http.StatusBadRequest)
	}

	respond(w, r, capitalisedText, http.StatusOK)
}

func rootHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	respondErr(w, r, errors.New("please pass in some text to munge"), http.StatusBadRequest)
}

func respond(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		respondErr(w, r, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("respond:", err)
	}
}

func respondErr(w http.ResponseWriter, r *http.Request, err error, code int) {
	errObj := struct {
		Error string `json:"error"`
	}{Error: err.Error()}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	err = json.NewEncoder(w).Encode(errObj)
	if err != nil {
		fmt.Println("respondErr:", err)
	}
}
