package page

import (
	"fmt"
	"net/http"
	"strconv"
)

func ParamPagination(param string, def int, r *http.Request) int {
	var page int

	p, err := checkParam(param, r)
	if err != nil {
		page = def
	} else {
		page = p
	}

	return page
}

func ParamOffset(size int, r *http.Request) (int, int) {
	var page int

	p, err := checkParam("page", r)
	if err != nil {
		page = 0
	} else {
		page = p
	}

	offsite := page * size

	return page, offsite
}

func checkParam(param string, r *http.Request) (int, error) {
	para := r.URL.Query()
	id := para.Get(param)

	if id == "" {
		return 0, fmt.Errorf("%s", "Empty parameter "+param)
	}

	st, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return st, nil
}
