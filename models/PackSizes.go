package models

type NumberOfPack struct {
	PackageSize     int32 `json:"packageSize"`
	NumberOfPackage int32 `json:"numberOfPackage"`
}

type CalculatedPackages struct {
	ItemsOrdered         int32          `json:"itemsOrdered"`
	CorrectNumberOfPacks []NumberOfPack `json:"correctNumberOfPacks"`
}

type CalculatePackagesDto struct {
	OrderedItems    int32   `json:"orderedItems"`
	PackageSizeList []int32 `json:"packageSizeList"`
}
