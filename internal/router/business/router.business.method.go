package business

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"project-template/model/dto"
	"strconv"
)

func (r *Router) AddBusiness(c echo.Context) error {
	httpRequest := c.Request()
	ctx := httpRequest.Context()
	bus := dto.InsertBusinessDTO{}
	if err := c.Bind(&bus); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Status: http.StatusBadRequest,
			Msg:    "cannot bind request",
			Data:   nil,
		})
	}
	if err := c.Validate(&bus); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Status: http.StatusBadRequest,
			Msg:    "validation error, field(s) are empty",
			Data:   nil,
		})
	}
	resp, err := r.controller.InsertNewBusiness(ctx, bus)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, dto.ResponseWrapper{
			Status: http.StatusInternalServerError,
			Msg:    "cannot insert business",
			Data:   nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseWrapper{
		Status: http.StatusOK,
		Msg:    "ok",
		Data:   resp,
	})
}

func (r *Router) EditBusiness(c echo.Context) error {
	bus := dto.UpdateBusinessDTO{}
	if err := c.Bind(&bus); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Status: http.StatusBadRequest,
			Msg:    "cannot bind request",
			Data:   nil,
		})
	}
	if err := c.Validate(&bus); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Status: http.StatusBadRequest,
			Msg:    "validation error, field(s) are empty",
			Data:   nil,
		})
	}
	httpRequest := c.Request()
	ctx := httpRequest.Context()
	resp, err := r.controller.UpdateBusiness(ctx, bus)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, dto.ResponseWrapper{
			Status: http.StatusInternalServerError,
			Msg:    "cannot update business",
			Data:   nil,
		})
	}
	return c.JSON(http.StatusOK, dto.ResponseWrapper{
		Status: http.StatusOK,
		Msg:    "ok",
		Data:   resp,
	})
}

func (r *Router) DeleteBusiness(c echo.Context) error {
	req := dto.DeleteBusinessDTO{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Status: http.StatusBadRequest,
			Msg:    "cannot bind request",
			Data:   nil,
		})
	}
	err = c.Validate(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseWrapper{
			Status: http.StatusBadRequest,
			Msg:    "id missing",
			Data:   nil,
		})
	}
	httpRequest := c.Request()
	ctx := httpRequest.Context()
	err = r.controller.DeleteBusiness(ctx, req)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, dto.ResponseWrapper{
			Status: http.StatusInternalServerError,
			Msg:    "cannot delete business",
			Data:   nil,
		})
	}
	return c.JSON(http.StatusOK, dto.ResponseWrapper{
		Status: http.StatusOK,
		Msg:    "deleted",
		Data:   nil,
	})
}

func (r *Router) GetBusinesses(c echo.Context) error {
	strPage := c.QueryParam("page")
	if strPage == "" {
		strPage = "1"
	}
	page, err := strconv.ParseUint(strPage, 10, 64)
	if err != nil {
		return err
	}
	strLimit := c.QueryParam("offset")
	if strLimit == "" {
		strLimit = "10"
	}
	limit, err := strconv.ParseUint(strLimit, 10, 64)
	if err != nil {
		return err
	}
	strLatitude := c.QueryParam("latitude")
	if strLatitude == "" {
		strLatitude = "0"
	}
	latitude, err := strconv.ParseFloat(strLatitude, 64)
	if err != nil {
		return err
	}
	strLongitude := c.QueryParam("longitude")
	if strLongitude == "" {
		strLongitude = "0"
	}
	longitude, err := strconv.ParseFloat(strLongitude, 64)
	if err != nil {
		return err
	}
	strOpenNow := c.QueryParam("open_now")
	openNow := true
	if strOpenNow == "false" {
		openNow = false
	}

	strPrice := c.QueryParam("price")
	price, err := strconv.ParseInt(strPrice, 10, 64)
	if err != nil {
		price = 0
	}
	req := dto.SearchBusinessDTO{
		Page:      page,
		Limit:     limit,
		Term:      c.QueryParam("term"),
		Location:  c.QueryParam("location"),
		Latitude:  latitude,
		Longitude: longitude,
		Locale:    c.QueryParam("locale"),
		OpenNow:   openNow,
		OpenAT:    c.QueryParam("open_at"),
		Price:     price,
	}

	ctx := c.Request().Context()
	//if req.Location == "" && req.Longitude == 0 && req.Latitude == 0 {
	//	return c.JSON(http.StatusBadRequest, dto.ResponseWrapper{
	//		Status: http.StatusBadRequest,
	//		Msg:    "location or longitude and latitude is not provided",
	//	})
	//}
	resp, err := r.controller.SearchBusiness(ctx, req)
	if err != nil {
		log.Error(err)
	}

	return c.JSON(http.StatusOK, dto.ResponseWrapper{
		Status: http.StatusOK,
		Msg:    "ok",
		Data:   resp,
	})
}
