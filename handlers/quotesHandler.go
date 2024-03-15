package handlers

import (
	"net/http"
	"strconv"
	"time"

	"quotes-api/initializers"
	"quotes-api/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Quotes API")
}

func GetQuotesList(c echo.Context) error {
	var quotes []models.QuotesList
	filter := bson.M{}

	// filter by
	if quoteText := c.QueryParam("quote"); quoteText != "" {
		filter["quotes"] = bson.M{"$regex": primitive.Regex{Pattern: quoteText, Options: "i"}}
	}

	if author := c.QueryParam("author"); author != "" {
		filter["author"] = bson.M{"$regex": primitive.Regex{Pattern: author, Options: "i"}}
	}

	if category := c.QueryParam("category"); category != "" {
		filter["category"] = bson.M{"$regex": primitive.Regex{Pattern: category, Options: "i"}}
	}

	// pagination
	page, err := strconv.Atoi(c.QueryParam("skip"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		pageSize = 10
	}

	// calculate the skip and limit values based on the pagination parameters
	skip := int64((page - 1) * pageSize)
	limit := int64(pageSize)

	cursor, err := initializers.QuotesCollection.Find(c.Request().Context(), filter, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch quotes"})
	}

	defer cursor.Close(c.Request().Context())

	if err = cursor.All(c.Request().Context(), &quotes); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to decode quotes"})
	}

	if len(quotes) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "No quotes found"})
	}

	return c.JSON(http.StatusOK, echo.Map{"quotes": quotes})
}

func CreateQuotes(c echo.Context) error {
	var quote models.Quotes

	if err := c.Bind(&quote); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	quote.CreatedAt = time.Now()
	quote.UpdatedAt = time.Now()
	quote.UsedAsQOTD = false

	_, err := initializers.QuotesCollection.InsertOne(c.Request().Context(), quote)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create quote"})
	}

	return c.JSON(http.StatusCreated, quote)
}

func QuoteOfTheDayHandler(c echo.Context) error {
	var quote models.QuotesList
	filter := bson.M{"used_as_qotd": false}

	// select a random quote
	opts := options.Find().SetSort(bson.D{
		bson.E{Key: "_id", Value: 1},
	})

	cursor, err := initializers.QuotesCollection.Find(c.Request().Context(), filter, opts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch quotes"})
	}

	defer cursor.Close(c.Request().Context())

	if cursor.Next(c.Request().Context()) {
		err = cursor.Decode(&quote)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to decode quote"})
		}
	} else {
		// if no unused quotes are available, reset
		_, err = initializers.QuotesCollection.UpdateMany(c.Request().Context(), bson.M{}, bson.M{"$set": bson.M{"used_as_qotd": false}})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to reset quotes"})
		}

		// finding a quote again after resetting
		cursor, err = initializers.QuotesCollection.Find(c.Request().Context(), filter, opts)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch quotes"})
		}
		defer cursor.Close(c.Request().Context())

		if cursor.Next(c.Request().Context()) {
			err = cursor.Decode(&quote)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to decode quote"})
			}
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No quotes available"})
		}
	}

	// mark the selected quote as used for "Quote of the Day"
	_, err = initializers.QuotesCollection.UpdateOne(c.Request().Context(), bson.M{"_id": quote.ID}, bson.M{"$set": bson.M{"used_as_qotd": true}})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update quote"})
	}

	return c.JSON(http.StatusOK, quote)
}

func ResetQuoteOfTheDay(c echo.Context) error {
	_, err := initializers.QuotesCollection.UpdateMany(
		c.Request().Context(),
		bson.M{},
		bson.M{"$set": bson.M{"used_as_qotd": false}},
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update quotes"})
	}

	return c.JSON(http.StatusAccepted, echo.Map{"message": "Successfully updated all the quotes"})
}
