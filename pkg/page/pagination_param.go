package page

import (
	"net/http"
)

func ParamPagination(param string, def int, r *http.Request) int {
	var page int

	p, err := CheckParam(param, r)
	if err != nil {
		page = def
	} else {
		page = p
	}

	return page
}

func ParamOffset(size int, r *http.Request) (int, int) {
	var page int

	p, err := CheckParam("page", r)
	if err != nil {
		page = 1
	} else {
		page = p
	}

	offsite := page * size

	return page, offsite
}
