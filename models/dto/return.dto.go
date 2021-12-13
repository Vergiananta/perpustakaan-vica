package dto

import uuid "github.com/satori/go.uuid"

type ReturnRequest struct {
	BukuID 			uuid.UUID `json:"buku_id"`
	PeminjamanID	uuid.UUID `json:"peminjaman_id"`
}

type ReturnResponse struct {

}
