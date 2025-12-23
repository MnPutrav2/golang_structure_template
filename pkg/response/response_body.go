package response

import (
	"clean-arsitektur/internal/model"
	logging "clean-arsitektur/pkg/logging"
	"encoding/json"
	"net/http"
)

func ResponseBody(response any, log string, ty string, w http.ResponseWriter, r *http.Request) {
	var status string

	switch ty {
	case "INFO":
		status = "Success"
	case "WARN":
		status = "Failed"
	case "ERROR":
		status = "Error"
	}

	res, _ := json.Marshal(model.ResponseBody{Status: status, Response: response})
	logging.Log(log, ty, r)
	w.WriteHeader(200)
	w.Write(res)
}
