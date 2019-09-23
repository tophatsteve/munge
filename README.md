# munge

A small set of microservices for trying out various cloud infrastructure tools.

The project is made up of 3 services:

- *frontend* - a webserver that takes a string of text as input to its root url. It calls the other 2 services passing the input text and returns the result.
- *capitalise* - takes a text string as input and returns it capitalised.
- *reverse* - takes a text string as input and returns it reversed.

These services are built into the following Docker images and pushed to Docker Hub:

- tophatsteve/frontend:latest
- tophatsteve/capitalise:latest
- tophatsteve/reverse:latest
