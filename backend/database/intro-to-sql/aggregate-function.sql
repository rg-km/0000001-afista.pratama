-- DATA QUERY LANGUAGE
-- 5, count, max, sum, min, avg

-- COUNT

--SELECT COUNT(*) FROM Employees;


-- SELECT COUNT(EmployeeSalary) FROM Employees;
-- SELECT EmployeeName, COUNT(EmployeeSalary) FROM Employees; 

-- SUM (tipe data numberic)
-- sum untuk tipe data selain numberic, akan dianggap 0
-- SELECT EmployeeName, SUM(EmployeeSalary) FROM Employees;

-- MIN, MAX, AVG
SELECT 
	MIN(EmployeeName), 
	MAX(EmployeeName), 
	AVG(EmployeeName) 
FROM Employees;

-- MIN ketika dengan tipe data selain numberic, itu works
-- MAX ketika dengan tipe data selain numberic, itu works