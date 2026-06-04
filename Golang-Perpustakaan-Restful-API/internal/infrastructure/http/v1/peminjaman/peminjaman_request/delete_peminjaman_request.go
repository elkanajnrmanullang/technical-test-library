package peminjaman_request

type DeletePeminjamanRequest struct {
	IDPeminjaman string `json:"id_peminjaman" validate:"required"`
}
