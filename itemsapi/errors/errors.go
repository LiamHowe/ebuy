package errors

type ErrorResponse struct {
    ErrorCode int           `json:"errorCode"`
    ErrorMessage string     `json:"errorMessage"`
}

func GetInvalidRequestParametersError(param string) ErrorResponse {
    return ErrorResponse{10001, "Invalid Request Parameter Provided: " + param}
}
