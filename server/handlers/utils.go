package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hansbala/myncer/core"
)

func WriteJSONOk(resp http.ResponseWriter, body any) error {
	// NOTE: Writing a body also writes the StatusOK to the header.
	if err := json.NewEncoder(resp).Encode(body); err != nil {
		return core.WrappedError(err, "failed to write to response body")
	}
	resp.Header().Add("Content-Type", "application/json")
	return nil
}
