package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
	m.Employees = append(m.Employees, e)
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
	for i, emp := range m.Employees {
		if emp.ID == id {
			m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
			break
		}
	}
}

// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
	// TODO: Implement this method
	sum := 0.0

	if len(m.Employees) == 0 {
		return 0
	}

	for _, emp := range m.Employees {
		sum += emp.Salary
	}
	return sum / float64(len(m.Employees))
}

// FindEmployeeByID finds and returns an employee by their ID.
func (m *Manager) FindEmployeeByID(id int) *Employee {
	// TODO: Implement this method
	for _, emp := range m.Employees {
		if emp.ID == id {
			return &emp
		}
	}
	return nil
}

func main() {
	manager := Manager{}
	manager.AddEmployee(Employee{ID: 1, Name: "Aliceee", Age: 30, Salary: 70000})
	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
	manager.RemoveEmployee(1)
	averageSalary := manager.GetAverageSalary()
	employee := manager.FindEmployeeByID(2)

	fmt.Printf("Average Salary: %f\n", averageSalary)
	if employee != nil {
		fmt.Printf("Employee found: %+v\n", *employee)
	}
}
