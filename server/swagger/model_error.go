package swagger

type ModelError struct {
	Message string  `json:"message"`
	Code    float64 `json:"code"`
}
