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

type ItemRequest struct {
    ID int                  `json:"id"`
    Name string             `json:"name"`
    CurrencyId int          `json:"currencyId"`
    PriceInteger int        `json:"priceInteger"`
    PriceDecimal int        `json:"priceDecimal"`
    NumberAvailable int     `json:"numberAvailable"`
    SellerId int            `json:"sellerId"`
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
    CurrencyId int
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

func EmptyItemResponse() ItemResponse {
    return ItemResponse{}
}

func NewItemResponse(id int, name string, price string, numberAvailable int, seller SellerResponse) ItemResponse {
    return ItemResponse{id, name, price, numberAvailable, seller}
}

func NewPriceWithSymbol(currencySymbol string, integer int, decimal int) Price {
    return Price{0, currencySymbol, integer, decimal}
}

func NewPriceWithId(currencyId int, integer int, decimal int) Price {
    return Price{currencyId, "", integer, decimal}
}
