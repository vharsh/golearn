package main      //package should compile as an executable program and not as a shared library

import (
  "fmt"
  "flag"        //command line arguments with the predefined methods 
  "os"          //system calls
  "net/http"    //retrieving the source-code of the webpage requested
  "io/ioutil"   //reads and prints the fetched source-code

)

func main() {
    flag.Parse()
    args := flag.Args()
    fmt.Println(args)
    if len(args) < 1 {
    fmt.Println("Please Enter the URL")
    os.Exit(1)
  }
    retrieve(args)
}

func retrieve(url string){// Prints the source code
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

