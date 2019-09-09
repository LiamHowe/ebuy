package itemsservice

import (
    "github.com/LiamHowe/ebuy/itemsapi/item"
    "github.com/LiamHowe/ebuy/itemsapi/itemsdao"
)

func GetItems(sellerId, currencyId int) []item.ItemResponse {
    itemsDao := itemsdao.NewItemsDao()
    items := itemsDao.GetItems(sellerId, currencyId)
    itemResponses := make([]item.ItemResponse, len(items))
    for index, item := range items {
        itemResponses[index] = convertItemToResponse(item)
    }

    return itemResponses
}

func GetItem(id int) item.ItemResponse {
    itemsDao := itemsdao.NewItemsDao()
    requestItem := itemsDao.GetItem(id)
    var itemResponse item.ItemResponse
    if (requestItem != nil) {
        itemResponse = convertItemToResponse(*requestItem)
    }
    return itemResponse
}


func DeleteItem(id int) bool {
    itemsDao := itemsdao.NewItemsDao()
    return itemsDao.DeleteItem(id)
}

func AddItem(itemRequest item.ItemRequest) {
    itemsDao := itemsdao.NewItemsDao()
    itemsDao.AddItem(convertItemRequest(itemRequest))
}

func convertItemRequest(itemRequest item.ItemRequest) item.Item {
    return item.Item{0, itemRequest.Name,
        item.Price{itemRequest.CurrencyId, "", itemRequest.PriceInteger, itemRequest.PriceDecimal},
        itemRequest.NumberAvailable, item.Seller{itemRequest.SellerId, ""}}
}

func convertItemToResponse(itemToConvert item.Item) item.ItemResponse {
    return item.NewItemResponse(
        itemToConvert.ID,
        itemToConvert.Name,
        itemToConvert.Price.String(),
        itemToConvert.NumberAvailable,
        convertSellerTpResonse(itemToConvert.Seller),
    )
}

func convertSellerTpResonse(sellerToConvert item.Seller) item.SellerResponse {
    return item.SellerResponse{sellerToConvert.ID, sellerToConvert.Username}
}
