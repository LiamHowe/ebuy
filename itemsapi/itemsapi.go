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

func main() {
    router := mux.NewRouter()
    itemsservice.AddItem(item.NewItem(1, "toaster", item.NewPrice("GBP", 15, 0), 5))
    itemsservice.AddItem(item.NewItem(2, "microwave", item.NewPrice("GBP", 25, 99), 0))
    itemsservice.AddItem(item.NewItem(3, "sofa", item.NewPrice("USD", 425, 0), 0))
    router.HandleFunc("/items", GetItemsEndpoint).Methods("GET")
    fmt.Println(itemsservice.GetItems())
    log.Fatal(http.ListenAndServe(":1234", router))
}
