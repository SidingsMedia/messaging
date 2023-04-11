package responses

type GeneralError struct {
  Code int `json:"code"`
  Message string `json:"message"`
}
