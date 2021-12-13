package dto

import (
	uuid "github.com/satori/go.uuid"
)

type PinjamRequest struct {
	LamaPinjam 	  	int		`json:"lama_pinjam"`
	Keterangan		string 		`json:"keterangan"`
	MemberID 		uuid.UUID 	`json:"member_id"`
	Quantity 		int32 		`json:"quantity"`
	BukuID 			uuid.UUID	`json:"buku_id"`
}

type PinjamResponse struct {
	ID 				uuid.UUID 	`json:"id"`
	TanggalPinjam 	string 		`json:"tanggal_pinjam"`
	TanggalKembali 	string 		`json:"tanggal_kembali"`
	BukuID 			uuid.UUID	`json:"buku_id"`
	Quantity 		int32 		`json:"quantity"`
}

