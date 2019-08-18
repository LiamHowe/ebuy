package itemsservice

import (
    "github.com/LiamHowe/ebuy/itemsapi/item"
    "github.com/LiamHowe/ebuy/itemsapi/itemsdao"
)

func GetItems() []item.ItemResponse {
    itemsDao := itemsdao.NewItemsDao()
    items := itemsDao.GetItems()
    itemResponses := make([]item.ItemResponse, len(items))
    for index, item := range items {
        itemResponses[index] = convertItem(item)
    }

    return itemResponses
}

func convertItem(itemToConvert item.Item) item.ItemResponse {
    return item.NewItemResponse(
        itemToConvert.ID,
        itemToConvert.Name,
        itemToConvert.Price.String(),
        itemToConvert.NumberAvailable,
        convertSeller(itemToConvert.Seller),
    )
}

func convertSeller(sellerToConvert item.Seller) item.SellerResponse {
    return item.SellerResponse{sellerToConvert.ID, sellerToConvert.Username}
}
