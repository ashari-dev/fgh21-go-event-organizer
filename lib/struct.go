package lib

type Respont struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results,omitempty"`
}
