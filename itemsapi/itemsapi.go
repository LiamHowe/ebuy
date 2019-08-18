package main

import (
    "fmt"
    "github.com/LiamHowe/ebuy/itemsapi/itemsservice"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func GetItemsEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(itemsservice.GetItems())
}

func main() {
    router := mux.NewRouter()
    router.Use(commonMiddleware)
    router.HandleFunc("/items", GetItemsEndpoint).Methods("GET")
    fmt.Println("eBuy Items API started")
    log.Fatal(http.ListenAndServe(":1234", router))
}

func commonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}
