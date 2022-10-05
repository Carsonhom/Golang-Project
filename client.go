package main

import (
    "fmt"
    "os"
    "github.com/urfave/cli"
    "io/ioutil"
    "log"
    "net/http"
)

var app = cli.NewApp()

func info() {
    app.Name = "Simple client CLI"
    app.Usage = "A simple CLI for making http requests"
    app.Author = "Carson Hom"
    app.Version = "1.0.0"
}

func commands() {
    app.Commands = []cli.Command{
        {
            Name: "count",
            Aliases: []string{"c"},
            Usage: "Retrieves the number of times count has been called",
            Action: func(c *cli.Context) {
                endpoint := "http://localhost:8080/count"
                resp, err := http.Get(endpoint)
                if err != nil {
                    log.Fatal(err)
                }
                defer resp.Body.Close()
                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                    log.Fatal(err)
                }
                fmt.Println(string(body))
            },
        },
        {
            Name: "user_agent",
            Aliases: []string{"ua"},
            Usage: "Retrieves information about the client", 
            Action: func(c *cli.Context) {
                endpoint := "http://localhost:8080/ua"
                resp, err := http.Get(endpoint)
                if err != nil {
                    log.Fatal(err)
                }
                defer resp.Body.Close()
                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                    log.Fatal(err)
                }
                fmt.Println(string(body))
            },
        },
        {
            Name: "head",
            Aliases: []string{"hd"},
            Usage: "Issues a head request and prints all data from the response", 
            Action: func(c *cli.Context) {
                endpoint := "http://localhost:8080/"
                resp, err := http.Head(endpoint)
                if err != nil {
                    log.Fatal(err)
                }
                for k, v := range resp.Header {
                    fmt.Printf("%s %s\n", k, v)
                }
            },
        },
        {
            Name: "default",
            Aliases: []string{" "},
            Usage: "Default endpoint", 
            Action: func(c *cli.Context) {
                endpoint := "http://localhost:8080/"
                resp, err := http.Get(endpoint)
                if err != nil {
                    log.Fatal(err)
                }
                defer resp.Body.Close()
                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                    log.Fatal(err)
                }
                fmt.Println(string(body))
            },
        },
    }
}

func main() {
    info()
    commands()

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}