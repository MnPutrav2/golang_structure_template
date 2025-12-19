package page

import (
	"fmt"
	"net/http"
	"strconv"
)

func CheckParam(param string, r *http.Request) (int, error) {
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
