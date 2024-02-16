package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Vehicle struct {
	ID    string  `json:"id"`
	Brand string  `json:"brand"`
	Name  string  `json:"Name"`
	Type  string  `json:"type"`
	Image string  `json:"image"`
	Price float64 `json:"price"`
}

var vehicles []Vehicle

func main() {
	if err := retrieveData(); err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/vehicles", getVehicles)
	router.GET("/vehicles/:id", getVehicleByID)
	router.POST("/vehicle", addVehicle)
	router.PUT("/vehicles/:id", updateVehicleByID)
	router.DELETE("/vehicle/:id", removeVehicleByID)
	router.Run("localhost:5000")
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} 

	router.Use(cors.New(config))
}

func retrieveData() error {
	data, err := ioutil.ReadFile("vehicles.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &vehicles); err != nil {
		return err
	}
	return nil
}

func getVehicles(c *gin.Context) {
	response := struct {
		Status   string    `json:"status"`
		Message  string    `json:"message"`
		Vehicles []Vehicle `json:"vehicles"`
	}{
		Status:   "success",
		Message:  "List of Vehicles",
		Vehicles: vehicles,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func getVehicleByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range vehicles {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "this vehicle could not be found"})
}

func addVehicle(c *gin.Context) {
	var newVehicle Vehicle

	if err := c.BindJSON(&newVehicle); err != nil {
		return
	}

	vehicles = append(vehicles, newVehicle)
	c.IndentedJSON(http.StatusCreated, newVehicle)
}

func updateVehicleByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range vehicles {
		if a.ID == id {
			if err := c.BindJSON(&vehicles[i]); err != nil {
				return
			}
			c.IndentedJSON(http.StatusOK, vehicles[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "this vehicle could not be updated"})
}

func removeVehicleByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range vehicles {
		if a.ID == id {
			vehicles = append(vehicles[:i], vehicles[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "vehicle successfully removed"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "this vehicle could not be removed"})
}
