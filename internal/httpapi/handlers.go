package httpapi

import "net/http"

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// _ indicates to junk the return, Write returns with (int, error) which are not needed here
	// []byte is required as Write takes a byte slice
	_, _ = w.Write([]byte("ok\n"))
}
