package itemsdao

import (
    "github.com/LiamHowe/ebuy/itemsapi/item"
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "Password123d"
  dbname   = "ebuy"
)

type itemsDao struct {
    psqlInfo string
}

const getUsersQuery =
    "SELECT i.id, i.name, i.number_available, c.currency_symbol, " +
    "   i.price_integer, i.price_decimal, s.id AS seller_id, s.username AS seller_username " +
    "FROM items i " +
    "   JOIN currency c ON c.id = i.currency_id" +
    "   JOIN sellers s ON s.id = i.seller_id"

func (itemsDao itemsDao) GetItems() []item.Item {
    db, err := sql.Open("postgres", itemsDao.psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    rows, err := db.Query(getUsersQuery)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    items := make([]item.Item, 0)
    for rows.Next() {
        var id, numberAvailable, priceInteger, priceDecimal, sellerId int
        var name, currencySymbol, sellerUsername string
        err = rows.Scan(&id, &name, &numberAvailable, &currencySymbol,
            &priceInteger, &priceDecimal, &sellerId, &sellerUsername)
        if err != nil {
            panic(err)
        }
        price := item.NewPriceWithSymbol(currencySymbol, priceInteger, priceDecimal)
        seller := item.Seller{sellerId, sellerUsername}
        items = append(items, item.Item{id, name, price, numberAvailable, seller})
    }

    err = rows.Err()
    if err != nil {
        panic(err)
    }

    return items
}

func (itemsDao itemsDao) AddItem(item item.Item) {
    db, err := sql.Open("postgres", itemsDao.psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    sqlStatement := `
        INSERT INTO items (name, number_available, price_integer, price_decimal, currency_id, seller_id)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`
    id := 0
    err = db.QueryRow(sqlStatement, item.Name, item.NumberAvailable,
        item.Price.Integer, item.Price.Decimal, item.Price.CurrencyId, item.Seller.ID,
    ).Scan(&id)
    if err != nil {
      panic(err)
    }
}

func NewItemsDao() itemsDao {
    psqlInfo := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    return itemsDao{psqlInfo}
}
