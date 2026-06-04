package denda_request

type DeleteDendaRequest struct {
	IDDenda string `json:"id_denda" validate:"required"`
}
