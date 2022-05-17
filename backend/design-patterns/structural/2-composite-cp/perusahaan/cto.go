package perusahaan

type CTO struct {
	Subordinate []Employee
}

func (c CTO) GetSalary() int {
	return 30
}

// CTO 1 => 30
// VP 2 => 40
// VP1 - 2 junior => 20
// VP2 - 1 junior => 10

// total 100

func (c CTO) TotalDivisonSalary() int {
	var total int

	for _, employee := range c.Subordinate {
		totalSalaryDivision := employee.TotalDivisonSalary()

		total += totalSalaryDivision
	}

	totalSalary := c.GetSalary() + total

	return totalSalary
}
