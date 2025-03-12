package data

type Rate struct {
	Currency string
	Price    float64
}

func (r *Rate) NewRate(curr string, price float64) *Rate {
	return &Rate{Currency: curr, Price: price}
}
