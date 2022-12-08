package delivery

import (
	"api-alta-dashboard/features/log"
	"api-alta-dashboard/middlewares"
	"api-alta-dashboard/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LogDelivery struct {
	logService log.ServiceInterface
}

func New(service log.ServiceInterface, e *echo.Echo) {
	handler := &LogDelivery{
		logService: service,
	}

	e.GET("/logs", handler.GetAllLog, middlewares.JWTMiddleware())
	e.POST("/logs", handler.CreateLog, middlewares.JWTMiddleware())
	e.PUT("/logs/:id", handler.UpdateLog, middlewares.JWTMiddleware())
	e.DELETE("/logs/:id", handler.DeleteLog, middlewares.JWTMiddleware())

	//middlewares.IsAdmin = untuk membatasi akses endpoint hanya admin
	//middlewares.UserOnlySameId = untuk membatasi akses user mengelola data diri sendiri saja

}

func (delivery *LogDelivery) GetAllLog(c echo.Context) error {
	query := c.QueryParam("name")
	helper.LogDebug("isi query = ", query)
	results, err := delivery.logService.GetAllLog(query)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all logs", dataResponse))
}

func (delivery *LogDelivery) CreateLog(c echo.Context) error {
	logInput := CreateLogRequest{}
	userID := middlewares.ExtractTokenUserId(c)

	errBind := c.Bind(&logInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(logInput)
	dataCore.UserID = uint(userID)
	err := delivery.logService.CreateLog(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed create data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *LogDelivery) UpdateLog(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	logInput := UpdateLogRequest{}
	errBind := c.Bind(&logInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(logInput)
	err := delivery.logService.UpdateLog(dataCore, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data."))
}

func (delivery *LogDelivery) DeleteLog(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	err := delivery.logService.DeleteLog(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
