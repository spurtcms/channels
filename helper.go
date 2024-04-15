package channels

import (
	"strconv"

	"gorm.io/gorm"
)

// return all parentid to child ids aray
func Categories(categoryid int, DB *gorm.DB) string {

	var id string

	categoreis, _ := EntryModel.GetChildCategories(categoryid, DB)

	for _, val := range categoreis {

		id += strconv.Itoa(val.Id) + ","
	}

	return id
}

//using name 
func CategoriesByUsingName(categoryname string, DB *gorm.DB) string {

	var id string

	categoreid, _ := EntryModel.GetCategoryIdByName(categoryname, DB)

	categoreis, _ := EntryModel.GetChildCategories(categoreid.Id, DB)

	for _, val := range categoreis {

		id += strconv.Itoa(val.Id) + ","
	}

	return id
}
