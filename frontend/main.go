package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tophatsteve/munge/frontend/reverse"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	router := httprouter.New()
	router.GET("/", rootHandler)
	router.GET("/:text", textHandler)
	log.Fatal(http.ListenAndServe(":8333", router))
}

func reverseText(text string) (string, error) {
	conn, err := grpc.Dial(
		"reverse-service:9191",
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

func textHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	text := p.ByName("text")
	reveresedText, err := reverseText(text)

	if err != nil {
		respondErr(w, r, err, http.StatusBadRequest)
	} else {
		respond(w, r, reveresedText, http.StatusOK)
	}
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
