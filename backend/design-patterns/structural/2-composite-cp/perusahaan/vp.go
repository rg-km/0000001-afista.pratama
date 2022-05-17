package perusahaan

type VP struct {
	Subordinate []Employee // membawahi banyak junior
}

func (vp VP) GetSalary() int {
	return 20
}

// junior 100
// vp 1
// 1000 (100 * 10) + 20 = 1020
func (vp VP) TotalDivisonSalary() int {
	var total int

	for _, employee := range vp.Subordinate {
		salaryEmployee := employee.GetSalary()
		total += salaryEmployee
	}

	totalSalary := vp.GetSalary() + total
	return totalSalary
}
