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

    userId, err := utils.GetUserIdByToken(c.GetHeader("Authorization"))

    if err != nil {
        log.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating catheter"})
        return
    }

    catheters, err := catheterService.GetCatheters(userId)

    if err != nil {
        log.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving catheters"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"catheters": catheters})
}

func GetCatheterByID(c *gin.Context) {

    catheterId := c.Param("id")

    userId, err := utils.GetUserIdByToken(c.GetHeader("Authorization"))

    if err != nil {
        log.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting catheter"})
        return
    }

    catheter, err := catheterService.GetCatheter(userId, catheterId)

    if err != nil {
        log.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting catheter"})
        return
    }

    c.JSON(http.StatusFound, catheter)
}

func UpdateCatheter(c *gin.Context) {
    var req UpdateCatheterRequest

    userId, err := utils.GetUserIdByToken(c.GetHeader("Authorization"))

    if err != nil {
        log.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating catheter"})
        return
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    catheterId := c.Param("id")

    err = catheterService.UpdateCatheter(userId, catheterId, req.Start, req.End)

    if err != nil && err.Error() != "nothing to update" {
        log.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating catheter"})
        return
    }

    c.Status(http.StatusNoContent)
}

func DeleteCatheter(c *gin.Context) {

    catheterId := c.Param("id")

    userId, err := utils.GetUserIdByToken(c.GetHeader("Authorization"))

    if err != nil {
        log.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating catheter"})
        return
    }

    err = catheterService.DeleteCatheter(userId, catheterId)

    if err != nil {
        log.Println(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting catheter"})
        return
    }

    c.Status(http.StatusOK)
}