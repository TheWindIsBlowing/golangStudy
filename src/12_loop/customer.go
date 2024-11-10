package loop

type Customer struct {
	ID      string
	Balence float64
}

type Store struct {
	m map[string]*Customer
}

func (s *Store) storeCustomer(customers []Customer) {
	for _, customer := range customers {
		s.m[customer.ID] = &customer
	}
}
