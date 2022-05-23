-- DDL : create
CREATE TABLE Employees (
	EmployeeID INT,
	EmployeeName VARCHAR(50),
	EmployeeSalary INT
);

-- DDL : alter
ALTER TABLE Employees ADD EmployeeJob VARCHAR;

-- sqlite itu tidak support MODIFY ke column
ALTER TABLE Employees RENAME TO table_employees;

-- drop column
ALTER TABLE table_employees DROP COLUMN EmployeeJob;

-- DDL : drop
DROP TABLE table_employees;

-- DDL : delete data
-- mysql, postgresql, sqlite tidak support
DELETE FROM users;