//go:build dev
// +build dev

//
package main

import (
	"log"
	"net/http"
	"os"
)

func public() http.Handler {
	log.Println("building static files for development")
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}
