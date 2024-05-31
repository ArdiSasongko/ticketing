package helper

import (
	"fmt"
	"net/url"
	"strconv"
)

func ExtractFilterSort(params url.Values) (map[string]string, string, int, int) {
	filters := make(map[string]string)
	var sort string

	for key, values := range params {
		if len(values) > 0 {
			fmt.Println(key)
			if len(key) > 8 && key[0:6] == "filter" {
				name := key[7 : len(key)-1]
				filters[name] = values[0]
			} else if len(key) == 4 && key[0:4] == "sort" {
				fmt.Println(values[0])
				sort = values[0]
			}
		}
	}
	limit, _ := strconv.Atoi(params.Get("limit"))
	page, _ := strconv.Atoi(params.Get("page"))

	return filters, sort, limit, page
}
