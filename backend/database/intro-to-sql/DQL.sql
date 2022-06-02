-- DATA QUERY LANGUAGE

-- SELECT EmployeeName, EmployeeSalary FROM Employees;

-- SELECT * FROM Employees;

-- AS, untuk mengubah nama column di view query
-- WHERE digunakan search / queru secara spesific
-- WHERE kita bisa menggunakan alias
-- alias nama table, untuk mempersingkat nama table Employeee jadi e 
-- contoh user u, school : s

SELECT 
	EmployeeSalary AS emp_salary,
	EmployeeID AS emp_id,
	EmployeeName AS emp_name
FROM 
	Employees;


