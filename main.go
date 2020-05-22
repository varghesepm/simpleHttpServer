package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type contentWriter struct {
	response http.ResponseWriter
	status   int
}

func loggGen(in http.Handler) (out http.Handler) {
	out = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content := &contentWriter{response: w}
		in.ServeHTTP(content, r)
		log.Printf("%v\t%v\t%v", content.status, r.Method, r.URL)
	})
	return
}

func (c *contentWriter) Write(b []byte) (int, error) {
	return c.response.Write(b)
}

func (c *contentWriter) Header() http.Header {
	return c.response.Header()
}

func (c *contentWriter) WriteHeader(code int) {
	c.status = code
	c.response.WriteHeader(code)
}

func main() {
	fmt.Println("simpleHttpServer is running on : http://localhost:8080")
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	http.ListenAndServe(":8080", loggGen(http.FileServer(http.Dir(cwd))))
}
