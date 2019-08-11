package itemsservice

import (
    "github.com/LiamHowe/ebuy/itemsapi/item"
    "github.com/LiamHowe/ebuy/itemsapi/itemsdao"
)

func GetItems() []item.ItemResponse {
    items := itemsdao.GetItems()
    itemResponses := make([]item.ItemResponse, len(items))
    for index, item := range items {
        itemResponses[index] = convert(item)
    }

    return itemResponses
}

func AddItem(item item.Item) {
    itemsdao.AddItem(item)
}

func convert(itemToConvert item.Item) item.ItemResponse {
    return item.NewItemResponse(
        itemToConvert.ID,
        itemToConvert.Name,
        itemToConvert.Price.String(),
        itemToConvert.NumberAvailable,
    )
}
