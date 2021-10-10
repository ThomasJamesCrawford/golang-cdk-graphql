package model

type Service struct {
	ID          string
	Name        string
	Description string
	Length      int
	Price       int
}

func (c *Service) Technicians() []*Technician {
	return make([]*Technician, 0)
}
