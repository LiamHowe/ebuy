package main

import (
    "fmt"
    "github.com/LiamHowe/ebuy/itemsapi/itemsservice"
    "github.com/LiamHowe/ebuy/itemsapi/item"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func GetItemsEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(itemsservice.GetItems())
}

func AddItemEndpoint(w http.ResponseWriter, req *http.Request) {
    var itemRequest item.ItemRequest
    _ = json.NewDecoder(req.Body).Decode(&itemRequest)

    // w.WriteHeader(http.StatusBadRequest)
    // json.NewEncoder(w).Encode(errors.GetInvalidRequestParametersError("id"))

    itemsservice.AddItem(itemRequest)
    w.WriteHeader(http.StatusCreated)
}

func main() {
    router := mux.NewRouter()
    router.Use(commonMiddleware)
    router.HandleFunc("/items", GetItemsEndpoint).Methods("GET")
    router.HandleFunc("/item", AddItemEndpoint).Methods("POST")
    fmt.Println("eBuy Items API started")
    log.Fatal(http.ListenAndServe(":1234", router))
}

func commonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}
