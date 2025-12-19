package page

import (
	"fmt"
)

func PaginationLink(page, size, count int, keyword string) (string, string) {
	var previousLink string
	var nextLink string

	if page == 0 {
		previousLink = ""
	} else {
		previousLink = fmt.Sprintf("page=%d&size=%d", page-1, size)
	}

	if count <= (page+1)*size {
		nextLink = ""
	} else {
		nextLink = fmt.Sprintf("page=%d&size=%d", page+1, size)
	}

	if keyword != "" {
		if previousLink != "" {
			previousLink += fmt.Sprintf("&keyword=%s", keyword)
		}

		if nextLink != "" {
			nextLink += fmt.Sprintf("&keyword=%s", keyword)
		}
	}

	return previousLink, nextLink
}
