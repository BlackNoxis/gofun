package main

import (
	"strings"
	"io"
	"fmt"
	"net/http"
	"os"
)

func downloadFromUrl(url string) {
        tokens := strings.Split(url, "/")
        fileName := tokens[len(tokens)-1]
        fmt.Println("Downloading", url, "to", fileName)

        output, err := os.Create(fileName)
        if err != nil {
                fmt.Println("Error while creating\n", fileName, "-", err)
                return
        }
        defer output.Close()

        response, err := http.Get(url)
        if err != nil {
                fmt.Println("Error while downloading\n", url, "-", err)
                return
        }
        defer response.Body.Close()

        n, err := io.Copy(output, response.Body)
        if err != nil {
                fmt.Println("Error while downloading\n", url, "-", err)
                return
        }

        fmt.Println(n, "Bytes downloaded.\n")
}
