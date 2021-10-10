package model

import "github.com/jackc/pgtype"

type Company struct {
	ID pgtype.UUID `json:"id"`
}

func (c *Company) Users() []*User {
	return make([]*User, 0)
}

func (c *Company) Technicians() []*Technician {
	return make([]*Technician, 0)
}

func (c *Company) ServiceTypes() []*ServiceType {
	return make([]*ServiceType, 0)
}
