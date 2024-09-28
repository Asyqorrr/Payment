package controller

import (
	"database/sql"
	"net/http"
	db "payment/db/sqlc"
	"payment/models"
	"payment/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type entityController struct {
	serviceManager *services.ServiceManager
}

// constructor
func NewEntityController(servicesManager services.ServiceManager) *entityController {
	return &entityController{
		serviceManager: &servicesManager,
	}
}

type entityCreateReq struct {
	EntityName string `json:"entity_name" binding:"required"`
}

type entityUpdateReq struct {
	EntityName string `json:"entity_name" binding:"required"`
}

// create entity
func (handler *entityController) CreateEntity(c *gin.Context) {
	var payload *entityCreateReq

	if err := c.ShouldBindJSON(&payload); err != nil{
		// c.JSON(http.StatusUnprocessableEntity, gin.H{"status" : "fail", "message" : err})
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}
	args := payload.EntityName

	entity, err := handler.serviceManager.EntityService.CreateEntity(c, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusCreated, entity)
}

func(handler *entityController) UpdateEntity(c *gin.Context){
	
	var payload *entityUpdateReq
	entyId, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}
	args := &db.UpdateEntityParams{
		EntityID:   int32(entyId),
		EntityName: payload.EntityName,
	}


	entity, err := handler.serviceManager.EntityService.UpdateEntity(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusCreated, entity)
}

// get list entity
func (handler *entityController) GetListEntity(c *gin.Context){
	entity, err := handler.serviceManager.EntityService.FindAllEntity(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
	}

	c.JSON(http.StatusOK, entity)
}

func (handler *entityController) DeleteEntity(c *gin.Context){
	entyId, _ := strconv.Atoi(c.Param("id"))

	err := handler.serviceManager.EntityService.DeleteEntity(c, int32(entyId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message" : "data has been deleted"})
}

func (handler *entityController) GetEntityById (c *gin.Context){
	entyId, _ := strconv.Atoi(c.Param("id"))
	entity, err := handler.serviceManager.EntityService.FindEntityById(c, int32(entyId))
	
	if err != nil {
		if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, models.ErrDataNotFound)
		return
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusOK, entity)

}

