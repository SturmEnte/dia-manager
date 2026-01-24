package inventory

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	inventoryService "dia-manager-backend/service/inventory"
	"dia-manager-backend/utils"
)

// Item Structure
func GetItemStructures(c *gin.Context) {

	var userId string = utils.GetUserIdByContext(c)

	structures, err := inventoryService.GetItemStructures(userId)

	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, inventoryService.ErrUserInput) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating item structure"})
		return
	}

	c.JSON(http.StatusOK, structures)

}

func CreateItemStructure(c *gin.Context) {

	var req CreateItemStructureRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	var userId string = utils.GetUserIdByContext(c)

	structureId, err := inventoryService.CreateItemStructure(userId, req.Name, req.Attributes, req.GeneralInformation)

	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, inventoryService.ErrUserInput) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating item structure"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": structureId})
}

// Items
func GetItems(c *gin.Context) {

	var userId string = utils.GetUserIdByContext(c)

	items, err := inventoryService.GetItems(userId)

	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, inventoryService.ErrUserInput) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting items"})
		return
	}

	c.JSON(http.StatusOK, items)
}

func CreateItems(c *gin.Context) { 

	var req CreateItemsRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	var userId string = utils.GetUserIdByContext(c)

	itemIds, err := inventoryService.CreateItems(userId, req.StructureId, req.Items)

	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, inventoryService.ErrUserInput) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating item structure"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ids": itemIds})

}