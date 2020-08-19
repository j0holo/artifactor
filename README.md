# Artifactor

A small file server for uploading and downloading artifacts.

## Goal

The goal is to provide a simple http server for uploading and download artifacts in a trusted environment.

## How to use it

Compile the server.go file with `go build server.go`

Run it with `./server` which will host an artifact server on http://0.0.0.0:8080.

To download files go to http://0.0.0.0:8080/artifact/<path_to_file>/<filename> so ./artifact/test.jpg will be http://0.0.0.0:8080/artifact/test.jpg.

To upload files send a POST request to with Content-Type: multipart/form-data; and as name `file`.

An example with curl:

`curl -F 'file=@minimal_whale.png' http://localhost:8080/upload`