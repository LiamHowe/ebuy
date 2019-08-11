package item

import (
    "strconv"
)

type ItemResponse struct {
    ID int                  `json:"id,omitempty"`
    Name string             `json:"name,omitempty"`
    Price string            `json:"price,omitempty"`
    NumberAvailable int     `json:"numberAvailable"`
}

type Item struct {
    ID int
    Name string
    Price Price
    NumberAvailable int
}

type Price struct {
    CurrencyCode string
    Integer int
    Decimal int
}

func (price Price) String() string {
    currencySymbol := getCurrency(price.CurrencyCode).CurrencySymbol
    priceString := currencySymbol + strconv.Itoa(price.Integer)
    if price.Decimal > 0 {
        priceString += "." + strconv.Itoa(price.Decimal)
    }
    return priceString
}

type Currency struct {
    CurrencyCode string
    CurrencySymbol string
}

var currencies = [...]Currency{
    Currency{"GBP", "£"},
    Currency{"EUR", "€"},
    Currency{"USD", "$"},
}

func getCurrency(currencyCode string) *Currency {
    for _, currency := range currencies {
        if currency.CurrencyCode == currencyCode {
            return &currency
        }
    }
    return nil
}

func NewItemResponse(id int, name string, price string, numberAvailable int) ItemResponse {
    return ItemResponse{id, name, price, numberAvailable}
}

func NewItem(id int, name string, price Price, numberAvailable int) Item {
    return Item{id, name, price, numberAvailable}
}

func NewPrice(currency string, integer int, decimal int) Price {
    return Price{currency, integer, decimal}
}
