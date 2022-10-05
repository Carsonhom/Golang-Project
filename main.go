package main

import ( // Libraries
    "fmt"
    "log"
    "net/http"
)

func main() {

    // http.HandleFunc("/", HelloHandler) // HelloHandler


    //---------------Handler Functions---------------
    http.HandleFunc("/status", func(w http.ResponseWriter, _ *http.Request) { // localhost:8080/status Handler, responds with http status code
        w.WriteHeader(http.StatusOK)
    })

    th := &CounterHandler{counter: 0} // Page visit counter
    http.Handle("/count", th) // localhost:8080/count Handler

    http.HandleFunc("/ua", func(w http.ResponseWriter, r *http.Request) { // HTTP Get, identifies client (localhost:8080/ua)

        ua := r.Header.Get("User-Agent")

        fmt.Fprintf(w, "User agent: %s\n", ua) // returns HTTP header name:value pair
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // default handler, 

        if r.URL.Path != "/" { // client specified invalid endpoint, return 404 error
            w.WriteHeader(404)
            w.Write([]byte("404 - not found\n"))
            return
        }

        fmt.Fprintln(w, "Hello World") // Client does not specify endpoint, return "Hello World" 
    })

    //---------------TCP connection---------------
    log.Println("Listening...")
    log.Fatal(http.ListenAndServe(":8080", nil)) // Listens on TCP network address and handles requests on incoming connections
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) { // hello handler, responds with "Hello World"

    fmt.Fprintf(w, "Hello World\n")
}

type CounterHandler struct { // Page visit count struct
    counter int
}

func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // Counter handler, responds with the number of times client has requested /count
    fmt.Println(ct.counter)
    ct.counter++
    fmt.Fprintln(w, "Counter:", ct.counter)
}
