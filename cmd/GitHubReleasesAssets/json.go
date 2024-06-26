package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/valyala/fastjson"
)

var client *http.Client
var parser *fastjson.Parser

func init() {
	client = &http.Client{
		Timeout: 15 * time.Second,
	}

	parser = &fastjson.Parser{}
}

func GetJSON(Url string) *fastjson.Value {
	var Error error

	//
	var Request *http.Request
	if Request, Error = http.NewRequest(http.MethodGet, Url, nil); Error == nil {
		Request.Header.Set("Content-Type", "application/json")
	} else {
		fmt.Println("[http.NewRequest]", Error)
		os.Exit(1)
	}

	var Response *http.Response
	if Response, Error = client.Do(Request); Error == nil {
		defer Response.Body.Close()
	} else {
		fmt.Println("[client.Do]", Error)
		os.Exit(1)
	}

	var Body []byte
	if Body, Error = io.ReadAll(Response.Body); Error != nil {
		fmt.Println("[io.ReadAll]", Error)
		os.Exit(1)
	}

	switch Response.StatusCode {
	case http.StatusOK:
		break
	case http.StatusNotFound:
		fmt.Print("Repository not found")
		os.Exit(1)
	case http.StatusTeapot:
		fmt.Print("Repository is a teapot")
		os.Exit(1)
	default:
		fmt.Printf("HTTP status %d %s\n", Response.StatusCode, http.StatusText(Response.StatusCode))
		os.Exit(1) // oof
	}

	var Result *fastjson.Value
	if Result, Error = parser.ParseBytes(Body); Error != nil {
		fmt.Println("[Parser.ParseBytes]", Error)
		os.Exit(1)
	}

	return Result
}
