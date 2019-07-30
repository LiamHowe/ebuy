package item

type Item struct {
    ID int                  `json:"id,omitempty"`
    Name string             `json:"name,omitempty"`
    Price Price             `json:"price,omitempty"`
    NumberAvailable int     `json:"numberAvailable,omitempty"`
}

type Price struct {
    Currency string         `json:"currency,omitempty"`
    Integer int             `json:"integer,omitempty"`
    Decimal int             `json:"decimal,omitempty"`
}

func NewItem(id int, name string, price Price, numberAvailable int) Item {
    return Item{id, name, price, numberAvailable}
}

func NewPrice(currency string, integer int, decimal int) Price {
    return Price{currency, integer, decimal}
}
