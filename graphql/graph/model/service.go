package model

import "github.com/jackc/pgtype"

type Service struct {
	ID            pgtype.UUID
	ServiceTypeId pgtype.UUID
	Name          pgtype.Varchar
	Description   pgtype.Varchar
	Length        pgtype.Int8
	Price         pgtype.Int8
}

func (c *Service) Technicians() []*Technician {
	return make([]*Technician, 0)
}
