package errors

import "net/http"

func HttpError(w http.ResponseWriter, status int) {
	switch status {
	case 500:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
