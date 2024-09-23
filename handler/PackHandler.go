// Description: This file contains the handler for the pack calculation endpoint.

package handler

import (
	"OptimalPackCount/models"
	"OptimalPackCount/service"
	"encoding/json"
	"net/http"
	"strconv"
)

// Default pack sizes
var defaultPackSizes = []int32{5000, 2000, 1000, 500, 250}

// CalculatePacksHandler handles HTTP requests for calculating the optimal pack count
func CalculatePacksHandler(w http.ResponseWriter, r *http.Request) {
	// Get the order size from the query parameters.
	orderSizeStr := r.URL.Query().Get("orderSize")
	orderSize, err := strconv.Atoi(orderSizeStr)
	if err != nil {
		// Return an error if the order size is invalid.
		http.Error(w, "Invalid order size", http.StatusBadRequest)
		return
	}
	// Prepare the DTO with the order size and default package sizes
	dto := models.CalculatePackagesDto{
		OrderedItems:    int32(orderSize),
		PackageSizeList: defaultPackSizes,
	}
	// Calculate the optimal package breakdown
	calculatedPackages := service.CalculatePackages(dto)

	// Marshal the result to JSON and write it to the response
	response, _ := json.Marshal(calculatedPackages)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
