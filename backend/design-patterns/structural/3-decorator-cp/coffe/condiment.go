package coffe

type Mocha struct {
	Coffe Coffe
}

func (m Mocha) GetCost() float64 {
	return m.Coffe.GetCost() + 1.00

}

func (m Mocha) GetDescription() string {
	//return ""
	// ditambah ", mocha"

	return m.Coffe.GetDescription() + ", Mocha"
}

type Whipcream struct {
	Coffe Coffe
}

func (w Whipcream) GetCost() float64 {
	return w.Coffe.GetCost() + 0.10
}

func (w Whipcream) GetDescription() string {
	//return ""
	return w.Coffe.GetDescription() + ", Whipcream"
}
