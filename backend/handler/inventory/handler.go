package inventory

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	inventoryService "dia-manager-backend/service/inventory"
	"dia-manager-backend/utils"
)

// Item Structure
func CreateItemStructure(c *gin.Context) {

	var req CreateItemStructureRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	var userId string = utils.GetUserIdByContext(c)

	structureId, err := inventoryService.CreateItemStructure(userId, req.Name, req.Attributes)

	if err != nil {
        log.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating item structure"})
        return
    }

	c.JSON(http.StatusCreated, gin.H{"id": structureId})
}

// Items
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
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating item structure"})
        return
    }

	c.JSON(http.StatusCreated, gin.H{"ids": itemIds})

}