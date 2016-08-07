package main

import (
    "log"
    "net/http"
)

func main() {
    res, err := http.Get("http://google.com")
    if err != nil {
        log.Fatal(err)
    }
    defer closeBody(res)

    log.Println("...")
}

func closeBody(res *http.Response) {
    log.Println("Closing body on end of function with defer")
    res.Body.Close()
}