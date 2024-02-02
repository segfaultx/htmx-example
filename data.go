package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"sort"
	"strconv"
)

func sortDataBy(data []TableData, field string, direction string) []TableData {
	switch field {
	case "company":
		if direction == "asc" {
			return sortData(data, func(i, j int) bool {
				return data[i].Company < data[j].Company
			})
		}
		return sortData(data, func(i, j int) bool {
			return data[i].Company > data[j].Company
		})
	case "contact":
		if direction == "asc" {
			return sortData(data, func(i, j int) bool {
				return data[i].Contact < data[j].Contact
			})
		}
		return sortData(data, func(i, j int) bool {
			return data[i].Contact > data[j].Contact
		})
	case "country":
		if direction == "asc" {
			return sortData(data, func(i, j int) bool {
				return data[i].Country < data[j].Country
			})
		}
		return sortData(data, func(i, j int) bool {
			return data[i].Country > data[j].Country
		})
	default:
		return data
	}
}
func sortData(data []TableData, compare func(i, j int) bool) []TableData {
	sort.Slice(data[:], compare)
	return data
}

func getById(id int) (*TableData, error) {
	for _, tableData := range Tabledata {
		if tableData.Id == id {
			return &tableData, nil
		}
	}
	return &TableData{}, errors.New("unknown id")
}

func updateById(data TableData) {
	Tabledata[0] = data
}
func HandleData(c echo.Context) error {
	if !isLoggedIn {
		return c.Redirect(http.StatusFound, "/")
	}
	c.Response().Header().Set("HX-Replace-URL", "/data")

	var data = map[string]interface{}{"Data": Tabledata}
	sortBy := c.QueryParam("sort")
	direction := c.QueryParam("direction")
	start, err := strconv.Atoi(c.QueryParam("from"))

	if sortBy != "" {
		switch direction {
		case "asc":
			direction = "desc"
		case "desc":
			direction = "asc"
		default:
			direction = "asc"
		}
		data["SortDirection"] = direction
		var dataSlice []TableData
		if err != nil && ((start + 100) < len(Tabledata)) {
			dataSlice = Tabledata[start : start+100]
		} else {
			dataSlice = Tabledata[0:100]
		}

		data["Data"] = sortDataBy(dataSlice, sortBy, direction)
	} else {
		data["SortDirection"] = nil
		if err == nil && ((start + 100) < len(Tabledata)) {
			data["Data"] = Tabledata[start : start+100]
		} else {
			data["Data"] = Tabledata[0:100]
		}
	}
	data["SortBy"] = sortBy
	if err == nil {
		if (start + 100) > len(Tabledata) {
			return c.JSON(http.StatusOK, "done")
		}
		data["Start"] = start + 100
	} else {
		data["Start"] = 100
		return c.Render(http.StatusOK, "data_table", data)
	}

	return c.Render(http.StatusOK, "data_table_content", data)
}

func HandleDataById(c echo.Context) error {
	entityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	data, err := getById(entityId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.Render(http.StatusOK, "table_row", data)
}

func HandleDataEdit(c echo.Context) error {
	entityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	data, err := getById(entityId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.Render(http.StatusOK, "table_row_edit", data)
}

func HandleDataUpdateById(c echo.Context) error {
	entityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	data, err := getById(entityId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	var requestBody TableData
	err = c.Bind(&requestBody)

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	data.Company = requestBody.Company
	data.Contact = requestBody.Contact
	data.Country = requestBody.Country
	updateById(*data)

	return c.Render(http.StatusOK, "table_row", data)
}
