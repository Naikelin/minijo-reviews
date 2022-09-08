package controllers

import (
	"backend/configs"
	"backend/models"
	"backend/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var kbCollection *mongo.Collection = configs.GetCollection(configs.DB, "kbs")
var validate = validator.New()

func CreateKb(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var kb models.Kb
	defer cancel()

	//validate the request body
	if err := c.Bind(&kb); err != nil {
		return c.JSON(http.StatusBadRequest, responses.KbResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&kb); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.KbResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	newKb := models.Kb{
		Id:    primitive.NewObjectID(),
		Desc:  kb.Desc,
		Stars: kb.Stars,
		Title: kb.Title,
	}

	result, err := kbCollection.InsertOne(ctx, newKb)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.KbResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.KbResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func GetAKb(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	kbId := c.Param("kbId")
	var kb models.Kb
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(kbId)

	err := kbCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&kb)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.KbResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.KbResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": kb}})
}

func EditAKb(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	kbId := c.Param("kbId")
	var kb models.Kb
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(kbId)

	//validate the request body
	if err := c.Bind(&kb); err != nil {
		return c.JSON(http.StatusBadRequest, responses.KbResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&kb); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.KbResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"title": kb.Title, "desc": kb.Desc, "Stars": kb.Stars}
	result, err := kbCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.KbResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	var updatedKb models.Kb
	if result.MatchedCount == 1 {
		err := kbCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedKb)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.KbResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}
	}

	return c.JSON(http.StatusOK, responses.KbResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": updatedKb}})
}

func DeleteAKb(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	kbId := c.Param("kbId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(kbId)

	result, err := kbCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.KbResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.JSON(http.StatusNotFound, responses.KbResponse{Status: http.StatusNotFound, Message: "error", Data: &echo.Map{"data": "Kb with specified ID not found!"}})
	}

	return c.JSON(http.StatusOK, responses.KbResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": "Kb successfully deleted!"}})
}

func GetAllKbs(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var kbs []models.Kb
	defer cancel()

	results, err := kbCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.KbResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleKb models.Kb
		if err = results.Decode(&singleKb); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.KbResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		kbs = append(kbs, singleKb)
	}

	return c.JSON(http.StatusOK, responses.KbResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": kbs}})
}
