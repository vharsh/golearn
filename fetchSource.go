package main    //package should compile as an executable program and not as a shared library

import (
  "fmt"
  "os"          //system calls
  "net/http"    //retrieving the source-code of the webpage requested
  "io/ioutil"   //reads and prints the fetched source-code
)

func main() {
	args := os.Args[1:]
    if len(args) < 1 {
	fmt.Fprintln(os.Stderr, "Please Enter the URL")
	os.Exit(1)
    }
    fmt.Println("Downloading: ", args)
    retrieve(args[0])
}

func retrieve(url string) {// Prints the source code
    resp, err := http.Get(url)
    if err != nil{
      fmt.Println("read error is:", err)
      return
    }
    body, err := ioutil.ReadAll(resp.Body);
    if err != nil{
      fmt.Println("read error is:", err)
      return
    } else{
      fmt.Println(string(body))
    }
}

