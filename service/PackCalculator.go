// Description: This file contains the implementation of the package service.
// The CalculatePackages function calculates the number of packages required to fulfill an order.

package service

import (
	"OptimalPackCount/models"
)

// CalculatePackages determines the optimal number of packages needed to fulfill an order.
func CalculatePackages(dto models.CalculatePackagesDto) models.CalculatedPackages {
	var calculatedPackages models.CalculatedPackages
	calculatedPackages.ItemsOrdered = dto.OrderedItems
	sortedPackageList := sortIntegerList(dto.PackageSizeList) // Sort the package sizes in ascending order
	var lastPartsPackage int32 = 0
	var remaining int32 = 0

	// Iterate through the sorted package sizes from largest to smallest
	for i := len(sortedPackageList) - 1; i >= 0; i-- {
		packageSize := sortedPackageList[i]

		if packageSize == 0 {
			continue
		}

		remainingOrderedItems := dto.OrderedItems

		if remaining != 0 {
			remainingOrderedItems = remaining // Use remaining items if any
		}

		numOfPackage := remainingOrderedItems / packageSize

		// If we can pack items with the current package size
		if numOfPackage > 0 && (lastPartsPackage == 0 || remainingOrderedItems == packageSize) {
			calculatedPackages.CorrectNumberOfPacks = append(calculatedPackages.CorrectNumberOfPacks,
				models.NumberOfPack{
					PackageSize:     packageSize,
					NumberOfPackage: numOfPackage,
				})

			remaining = remainingOrderedItems - (numOfPackage * packageSize)
			lastPartsPackage = packageSize
		} else if numOfPackage == 0 {
			lastPartsPackage = packageSize
		}
	}

	// Handle remaining items or add a package if necessary to complete the order
	if len(calculatedPackages.CorrectNumberOfPacks) == 0 || remaining != 0 {
		existsInList := false

		// Check if the last package size is already in the list
		for i := 0; i < len(calculatedPackages.CorrectNumberOfPacks); i++ {
			packNo := calculatedPackages.CorrectNumberOfPacks[i].PackageSize
			if packNo == lastPartsPackage {
				existsInList = true
				calculatedPackages.CorrectNumberOfPacks[i].NumberOfPackage = calculatedPackages.CorrectNumberOfPacks[i].NumberOfPackage + 1
				break
			}
		}
		if !existsInList {
			calculatedPackages.CorrectNumberOfPacks = append(calculatedPackages.CorrectNumberOfPacks,
				models.NumberOfPack{
					PackageSize:     lastPartsPackage,
					NumberOfPackage: 1,
				})
		}
	}

	return calculatedPackages
}

// sortIntegerList sorts a list of integers in ascending order.
// It uses the bubble sort algorithm.
func sortIntegerList(list []int32) []int32 {
	if len(list) == 0 {
		return list
	}

	// Bubble sort algorithm itself
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list // return the sorted list
}
