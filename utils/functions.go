package utils

import (
	"databaselineservice/sdk/cervello"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func FindInStringList(list []string, elem string) (int, error) {
	for i, object := range list {
		if object == elem {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

func ParseDate(date string) (time.Time, error) {
	//layout := "2018-01-20"
	return time.Parse("2006-01-02", date)
}

func GetDateFromTime(rTime time.Time) (time.Time, error) {
	s := StringDate(rTime)
	return ParseDate(s)
}

func StringDate(rTime time.Time) string {
	year, month, day := rTime.Date()
	var sMonth string
	var sDay string
	if month > 9 {
		sMonth = fmt.Sprintf("%d", month)
	} else {
		sMonth = fmt.Sprintf("0%d", month)
	}
	if day > 9 {
		sDay = fmt.Sprintf("%d", day)
	} else {
		sDay = fmt.Sprintf("0%d", day)
	}
	return fmt.Sprintf("%d-%s-%s", year, sMonth, sDay)
}

func ConvertToCervelloQuery(queries url.Values) (cervello.QueryParams, error) {
	var err error
	result := cervello.QueryParams{
		Filters:       make([]cervello.Filter, 0),
		PaginationObj: cervello.Pagination{},
		Custom:        nil,
	}
	result.Custom = map[string]interface{}{}
	if len(queries) > 20 {
		return result, errors.New("cant accept more than 20 query params")
	}

	result.PaginationObj.PageSize = 10
	result.PaginationObj.PageNumber = 1
	if queries["pageSize"] != nil {
		var v int64
		v, err = strconv.ParseInt(queries["pageSize"][0], 10, 64)
		if err != nil {
			return result, errors.New(fmt.Sprintf("bad pagination object: %s", err.Error()))
		}
		result.PaginationObj.PageSize = int(v)
	}
	if queries["pageNumber"] != nil {
		var v int64
		v, err = strconv.ParseInt(queries["pageNumber"][0], 10, 64)
		if err != nil {
			return result, errors.New(fmt.Sprintf("bad pagination object: %s", err.Error()))
		}
		result.PaginationObj.PageNumber = int(v)
	}

	if queries["sort"] != nil {
		v := strings.Split(queries["sort"][0], " ")
		result.Custom["sort"] = fmt.Sprintf("order by \"%s\" %s", v[0], v[1])
	} else {

	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("filters[%d][key]", i)
		if queries[key] == nil {
			break
		}

		operator := fmt.Sprintf("filters[%d][operator]", i)
		if queries[operator] == nil {
			return result, errors.New(fmt.Sprintf("missing operator for filter no: %d", i))
		}

		value := fmt.Sprintf("filters[%d][value]", i)
		if queries[value] != nil {
			filter := cervello.Filter{
				Key:   queries[key][0],
				Op:    queries[operator][0],
				Value: queries[value][0],
			}
			result.Filters = append(result.Filters, filter)

		} else if value = fmt.Sprintf("filters[%d][value][0]", i); queries[value] != nil {

			result.Custom[key] = queries[key][0]
			result.Custom[operator] = queries[operator][0]
			for j := 0; j < 5; j++ {
				value = fmt.Sprintf("filters[%d][value][%d]", i, j)
				if queries[value] == nil {
					break
				}
				result.Custom[value] = queries[value][0]
			}
		} else {
			return result, errors.New(fmt.Sprintf("missing value for filter no: %d", i))
		}

	}

	return result, nil
}
