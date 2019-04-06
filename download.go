
package main

import (
	"net/http"
	"os"
	"io"
	"fmt"
	"strings"
)

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}


func download(filename string) error  {

	token:= strings.Split(filename, "/")
	name := token[len(token)-1]

	fmt.Printf("Downloading the file %s\n", filename)


	 // Get the data
    resp, err := http.Get(filename)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create the file
    if fileExists(name)  {
    	os.Remove(name)
    }

    out, err := os.Create(name)
    if err != nil {
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err

 
}