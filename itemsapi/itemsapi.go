package main

import (
    "fmt"
    "github.com/LiamHowe/ebuy/itemsapi/itemsservice"
    "github.com/LiamHowe/ebuy/itemsapi/item"
    "github.com/LiamHowe/ebuy/itemsapi/errors"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
)

func getItemsEndpoint(w http.ResponseWriter, r *http.Request) {
    var sellerId, currencyId int
    var err error
    if sellerIdStr := r.FormValue("sellerId"); sellerIdStr != "" {
        sellerId, err = strconv.Atoi(sellerIdStr)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(errors.GetInvalidRequestParametersError("sellerId"))
            return
        }
    }
    if currencyIdStr := r.FormValue("currencyId"); currencyIdStr != "" {
        currencyId, err = strconv.Atoi(currencyIdStr)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(errors.GetInvalidRequestParametersError("currencyId"))
            return
        }
    }

    json.NewEncoder(w).Encode(itemsservice.GetItems(sellerId, currencyId))
}

func getItemEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(errors.GetInvalidRequestParametersError("id"))
        return
    }

    itemResponse := itemsservice.GetItem(id)
    if itemResponse == item.EmptyItemResponse() {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(itemsservice.GetItem(id))
}

func addItemEndpoint(w http.ResponseWriter, req *http.Request) {
    var itemRequest item.ItemRequest
    _ = json.NewDecoder(req.Body).Decode(&itemRequest)

    itemsservice.AddItem(itemRequest)
    w.WriteHeader(http.StatusCreated)
}

func deleteItemEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(errors.GetInvalidRequestParametersError("id"))
        return
    }

    success := itemsservice.DeleteItem(id)
    if !success {
        w.WriteHeader(http.StatusNotFound)
        return
    }
}

func main() {
    router := mux.NewRouter()
    router.Use(commonMiddleware)
    router.HandleFunc("/items", getItemsEndpoint).Methods("GET")
    router.HandleFunc("/items", getItemsEndpoint).Methods("GET").
        Queries("sellerId", "{sellerId}").
        Queries("currencyId", "{currencyId}")
    router.HandleFunc("/item/{id}", getItemEndpoint).Methods("GET")
    router.HandleFunc("/item/{id}", deleteItemEndpoint).Methods("DELETE")
    router.HandleFunc("/item", addItemEndpoint).Methods("POST")
    fmt.Println("eBuy Items API started")
    log.Fatal(http.ListenAndServe(":1234", router))
}

func commonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}
