package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewHandler(store Store) *Handler {
	return &Handler{
		store: store,
	}
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Handler struct {
	store Store
}

// GetCountries return list of supported countries
func (h *Handler) GetCountries(c *gin.Context) {
	countries, err := h.store.GetCountries(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Status: "Fail", Data: "Internal Server"})
		return
	}

	c.JSON(http.StatusOK, Response{Status: "success", Data: countries})
}

// GetStates return list supported states in a country
func (h *Handler) GetStates(c *gin.Context) {
	country := c.Query("country")
	if len(country) == 0 {
		c.JSON(http.StatusBadRequest, Response{Status: "Fail", Data: "Bad Request"})
		return
	}

	states, err := h.store.GetStates(c, country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Status: "Fail", Data: "Internal Server"})
		return
	}

	c.JSON(http.StatusOK, Response{Status: "success", Data: states})
}

// GetCities return list supported cities in a state
func (h *Handler) GetCities(c *gin.Context) {
	fmt.Println(c.Request.URL.String())
	country := c.Query("country")
	if len(country) == 0 {
		c.JSON(http.StatusBadRequest, Response{Status: "Fail", Data: "Bad Request"})
		return
	}

	state := c.Query("state")
	if len(state) == 0 {
		c.JSON(http.StatusBadRequest, Response{Status: "Fail", Data: "Bad Request"})
		return
	}

	cities, err := h.store.GetCities(c, country, state)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Status: "Fail", Data: "Internal Server"})
		return
	}

	c.JSON(http.StatusOK, Response{Status: "success", Data: cities})
}

func (h *Handler) GetNearestCity(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Status: "Success", Data: M{
		"city":    "Port Harcourt",
		"state":   "Rivers",
		"country": "Nigeria",
		"location": M{
			"type":        "Point",
			"coordinates": []float64{7.048623, 4.854166}},
		"forecasts": []M{
			{
				"ts":     "2019-08-15T12:00:00.000Z",
				"aqius":  137,
				"aqicn":  69,
				"tp":     23,
				"tp_min": 23,
				"pr":     996,
				"hu":     100,
				"ws":     2,
				"wd":     225,
				"ic":     "10d",
			},
		}}})
}
