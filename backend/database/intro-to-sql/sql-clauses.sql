-- SQL Clauses

-- WHERE
	-- AND, OR, LIKE 
-- LIMIT, ORDER BY

--SELECT * 
--FROM Employees
--WHERE EmployeeName = 'Jane' AND EmployeeSalary = 20000000;


--SELECT *
--FROM Employees
--WHERE EmployeeName = 'Jane' OR EmployeeSalary > 20000000;

-- SELECT * 
-- FROM Employees
-- WHERE EmployeeName LIKE 'J%'; // cari semua yang memiliki awalan J, setelah itu bebeas belakangnya
-- WHERE EmployeeName LIKE 'ja%'; // cari semua yang memiliki awalan Ja, setelah itu bebas
-- WHERE EmployeeName LIKE '%e'; // cari semua bebas yang penting belakangnya ada 'e'
-- WHERE EmployeeName LIKE '%ne';
-- WHERE EmployeeName LIKE '%ne%';

-- SELECT *
-- FROM Employees
-- LIMIT 100; -- LIMIT 100

--SELECT *`
--FROM Employees
--ORDER BY EmployeeName ASC; -- ASC | DESC

-- GROUP BY : ada aggregate function

-- SELECT EmployeeName, SUM(EmployeeSalary)
-- FROM Employees
-- GROUP BY EmployeeName;

-- HAVING
SELECT *, COUNT(*) FROM Employees
GROUP BY EmployeeName
HAVING COUNT(EmployeeName) > 1;

