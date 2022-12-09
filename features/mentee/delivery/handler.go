package delivery

import (
	"api-alta-dashboard/features/mentee"
	"api-alta-dashboard/middlewares"
	"api-alta-dashboard/utils/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	menteeService mentee.ServiceInterface
}

func New(service mentee.ServiceInterface, e *echo.Echo) {
	handler := &MenteeDelivery{
		menteeService: service,
	}

	e.GET("/mentees", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/mentees/:id", handler.GetById, middlewares.JWTMiddleware())
	e.POST("/mentees", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/mentees/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/mentees/:id", handler.Delete, middlewares.JWTMiddleware())

	//middlewares.IsAdmin = untuk membatasi akses endpoint hanya admin
	//middlewares.UserOnlySameId = untuk membatasi akses user mengelola data diri sendiri saja

}

func (delivery *MenteeDelivery) GetAll(c echo.Context) error {
	queryName := c.QueryParam("name")
	queryIdClass := c.QueryParam("class_id")
	queryStatus := c.QueryParam("status")
	queryEducationType := c.QueryParam("education_type")

	helper.LogDebug("\n\n\nULALA")

	// debug cek query param masuk
	helper.LogDebug("\n isi queryName = ", queryName)
	helper.LogDebug("\n isi queryIDClass= ", queryIdClass)
	helper.LogDebug("\n isi queryStatus = ", queryStatus)
	helper.LogDebug("\n isi queryEdType = ", queryEducationType)

	results, err := delivery.menteeService.GetAll(queryName, queryStatus, queryIdClass, queryEducationType)
	if err != nil {

		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := toResponseMenteeListArray(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read all data.", dataResponse))
}

func (delivery *MenteeDelivery) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	results, err := delivery.menteeService.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	fmt.Println("ini mentee by id : \n\n", results)

	dataResponse := toResponseMentee(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read data.", dataResponse))
}

func (delivery *MenteeDelivery) Create(c echo.Context) error {
	userInput := InsertRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	err := delivery.menteeService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *MenteeDelivery) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}

	userInput := UpdateRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	err := delivery.menteeService.Update(dataCore, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data."))
}

func (delivery *MenteeDelivery) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error. Id must integer."))
	}
	err := delivery.menteeService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
