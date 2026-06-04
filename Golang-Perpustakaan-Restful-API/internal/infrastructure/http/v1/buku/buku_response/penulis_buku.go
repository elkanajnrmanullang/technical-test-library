package buku_response

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"

type ResponseSingleDataPenulis struct {
	Error  bool                 `json:"error"`
	Status string               `json:"status"`
	Data   *domain.Penulis_Buku `json:"data"`
}

type ResponseManyDataPenulis struct {
	Error  bool                  `json:"error"`
	Status string                `json:"status"`
	Data   []domain.Penulis_Buku `json:"data"`
}

func GetPenulisBukuByIdResponse(data *domain.Penulis_Buku, msg string, err bool) ResponseSingleDataPenulis {
	var dt ResponseSingleDataPenulis
	dt.Error = err
	dt.Status = msg
	dt.Data = data
	return dt
}

func GetAllPenulisBukuResponse(data []domain.Penulis_Buku, msg string, err bool) ResponseManyDataPenulis {
	var dt ResponseManyDataPenulis
	dt.Error = err
	dt.Status = msg
	dt.Data = data
	return dt
}

