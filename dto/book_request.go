package dto

type BookRequestBody struct {
	Name string `json:"name"`
}

type BookRequestParams struct {
	Id uint `json:"id"`
}
