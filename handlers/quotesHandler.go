package handlers

import (
	"context"
	"net/http"
	"quotes-api/initializers"
	"quotes-api/models"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func HelloIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetQuotesList(c echo.Context) error {
	var quotes []models.Quotes

	cursor, err := initializers.QuotesCollection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var quote models.Quotes

		if err := cursor.Decode(&quote); err != nil {
			return err
		}

		quotes = append(quotes, quote)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"quotes": quotes})
}

func CreateQuotes(c echo.Context) error {
	var quotes models.QuotesList

	if err := c.Bind(&quotes); err != nil {
		return err
	}

	quotes.CreatedAt = time.Now()
	quotes.UpdatedAt = time.Now()

	_, err := initializers.QuotesCollection.InsertOne(context.Background(), quotes)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, quotes)
}
