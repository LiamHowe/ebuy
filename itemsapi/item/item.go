package item

import (
    "strconv"
)

type ItemResponse struct {
    ID int                  `json:"id,omitempty"`
    Name string             `json:"name,omitempty"`
    Price string            `json:"price,omitempty"`
    NumberAvailable int     `json:"numberAvailable"`
    Seller SellerResponse   `json:"seller,omitempty"`
}

type SellerResponse struct {
    ID int                  `json:"id,omitempty"`
    Username string         `json:"username,omitempty"`
}

type Item struct {
    ID int
    Name string
    Price Price
    NumberAvailable int
    Seller Seller
}

type Seller struct {
    ID int
    Username string
}

type Price struct {
    CurrencySymbol string
    Integer int
    Decimal int
}

func (price Price) isValid() bool {
    return price.CurrencySymbol != "" && price.Integer >= 0 &&
        price.Decimal >= 0 && price.Decimal <= 99
}

func (price Price) isNotValid() bool {
    return !price.isValid()
}

func (price Price) String() string {
    if price.isNotValid() {
        return ""
    }
    priceString := price.CurrencySymbol + strconv.Itoa(price.Integer)
    if price.Decimal > 0 {
        priceString += "." + strconv.Itoa(price.Decimal)
    }
    return priceString
}

func NewItemResponse(id int, name string, price string, numberAvailable int, seller SellerResponse) ItemResponse {
    return ItemResponse{id, name, price, numberAvailable, seller}
}

func NewPrice(currency string, integer int, decimal int) Price {
    return Price{currency, integer, decimal}
}
