package model

type Company struct {
	ID string
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
