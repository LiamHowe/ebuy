package main

import (
    "fmt"
    "github.com/LiamHowe/ebuy/itemsapi/itemsdao"
    "github.com/LiamHowe/ebuy/itemsapi/item"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func GetItemsEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(itemsdao.GetItems())
}

func main() {
    router := mux.NewRouter()
    itemsdao.AddItem(item.NewItem(1, "toaster", item.NewPrice("GBP", 15, 0), 5))
    itemsdao.AddItem(item.NewItem(2, "microwave", item.NewPrice("GBP", 25, 99), 0))
    router.HandleFunc("/items", GetItemsEndpoint).Methods("GET")
    fmt.Println(itemsdao.GetItems())
    log.Fatal(http.ListenAndServe(":1234", router))
}
