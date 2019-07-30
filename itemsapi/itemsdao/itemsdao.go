package itemsdao

import (
    "github.com/LiamHowe/ebuy/itemsapi/item"
)

var items []item.Item

func GetItems() []item.Item {
    return items
}

func AddItem(item item.Item) {
    items = append(items, item)
}
