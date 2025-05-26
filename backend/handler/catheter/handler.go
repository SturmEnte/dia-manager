package catheter

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	catheterService "dia-manager-backend/service/catheter"
	"dia-manager-backend/utils"
)

func CreateCatheter(c *gin.Context) {

	var req CreateCatheterRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userId, err := utils.GetUserIdByToken(c.GetHeader("Authorization"))

    if err != nil {
        log.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating catheter"})
        return
    }

    catheterId, err := catheterService.CreateCatheter(userId, req.Start, req.End)

    if err != nil {
        log.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating catheter"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"id": catheterId})
}

func GetCatheters(c *gin.Context) {
    c.JSON(http.StatusNotImplemented, gin.H{})
}

func GetCatheterByID(c *gin.Context) {
    c.JSON(http.StatusNotImplemented, gin.H{})
}

func UpdateCatheter(c *gin.Context) {
    var req UpdateCatheterRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    catheterId := c.Param("id")

    err := catheterService.UpdateCatheter(catheterId, req.Start, req.End)

    if err != nil && err.Error() != "nothing to update" {
        log.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating catheter"})
        return
    }

    c.Status(http.StatusNoContent)
}

func DeleteCatheter(c *gin.Context) {
    c.JSON(http.StatusNotImplemented, gin.H{})
}