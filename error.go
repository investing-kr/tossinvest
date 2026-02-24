package tossinvest

import "fmt"

type ErrorResponse struct {
	ErrorBody struct {
		StatusCode int    `json:"statusCode"`
		Code       string `json:"code"`
		Message    string `json:"message"`
	} `json:"error"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("tossinvest: %+v", e.ErrorBody)
}
