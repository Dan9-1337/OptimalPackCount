package tests

import (
	"OptimalPackCount/models"
	"OptimalPackCount/service"
	"testing"
)

func TestCalculatePackages(t *testing.T) {
	packSizes := []int32{5000, 2000, 1000, 500, 250}

	testCases := []struct {
		orderSize int32
		expected  models.CalculatedPackages
	}{
		{1, models.CalculatedPackages{ItemsOrdered: 1, CorrectNumberOfPacks: []models.NumberOfPack{{PackageSize: 250, NumberOfPackage: 1}}}},
		{250, models.CalculatedPackages{ItemsOrdered: 250, CorrectNumberOfPacks: []models.NumberOfPack{{PackageSize: 250, NumberOfPackage: 1}}}},
		{251, models.CalculatedPackages{ItemsOrdered: 251, CorrectNumberOfPacks: []models.NumberOfPack{{PackageSize: 500, NumberOfPackage: 1}}}},
		{501, models.CalculatedPackages{ItemsOrdered: 501, CorrectNumberOfPacks: []models.NumberOfPack{{PackageSize: 500, NumberOfPackage: 1}, {PackageSize: 250, NumberOfPackage: 1}}}},
		{12001, models.CalculatedPackages{ItemsOrdered: 12001, CorrectNumberOfPacks: []models.NumberOfPack{{PackageSize: 5000, NumberOfPackage: 2}, {PackageSize: 2000, NumberOfPackage: 1}, {PackageSize: 250, NumberOfPackage: 1}}}},
	}

	for _, tc := range testCases {
		dto := models.CalculatePackagesDto{
			OrderedItems:    tc.orderSize,
			PackageSizeList: packSizes,
		}

		result := service.CalculatePackages(dto)
		if !equal(result.CorrectNumberOfPacks, tc.expected.CorrectNumberOfPacks) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func equal(a, b []models.NumberOfPack) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
