package delivery

import (
	"api-alta-dashboard/features/class"
	"api-alta-dashboard/middlewares"
	"api-alta-dashboard/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ClassDelivery struct {
	classService class.ServiceInterface
}

func New(service class.ServiceInterface, e *echo.Echo) {
	handler := &ClassDelivery{
		classService: service,
	}

	e.GET("/classes", handler.GetAllClass, middlewares.JWTMiddleware())
	e.GET("/classes/:id", handler.GetByIdClass, middlewares.JWTMiddleware())
	e.POST("/classes", handler.CreateClass, middlewares.JWTMiddleware())
	e.PUT("/classes/:id", handler.UpdateClass, middlewares.JWTMiddleware())
	e.DELETE("/classes/:id", handler.DeleteClass, middlewares.JWTMiddleware())

	//middlewares.IsAdmin = untuk membatasi akses endpoint hanya admin
	//middlewares.UserOnlySameId = untuk membatasi akses user mengelola data diri sendiri saja

}

func (delivery *ClassDelivery) GetAllClass(c echo.Context) error {
	query := c.QueryParam("name")
	helper.LogDebug("isi query = ", query)
	results, err := delivery.classService.GetAllClass(query)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all classes", dataResponse))
}

func (delivery *ClassDelivery) GetByIdClass(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.classService.GetByIdClass(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read class.", dataResponse))
}

func (delivery *ClassDelivery) CreateClass(c echo.Context) error {
	classInput := CreateClassRequest{}
	errBind := c.Bind(&classInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(classInput)
	err := delivery.classService.CreateClass(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed create data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *ClassDelivery) UpdateClass(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	classInput := UpdateClassRequest{}
	errBind := c.Bind(&classInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(classInput)
	err := delivery.classService.UpdateClass(dataCore, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data."))
}

func (delivery *ClassDelivery) DeleteClass(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	err := delivery.classService.DeleteClass(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
