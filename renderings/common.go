package renderings

type CommonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}


type ResultResponse struct{
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}