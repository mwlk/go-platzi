package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id         int
		dni        string
		mockFunc   func()
		ftExpected FullTimeEmployee
	}{
		{id: 1, dni: "1", mockFunc: func() {
			GetEmployeeById = func(id int) (Employee, error) {
				return Employee{
					Id:       1,
					Position: "CEO",
				}, nil
			}
			GetPersonByDni = func(dni string) (Person, error) {
				return Person{
					Name: "Mirko Wlk",
					Age:  27,
					Dni:  "35480",
				}, nil
			}
		}, ftExpected: FullTimeEmployee{
			Person: Person{
				Age:  27,
				Dni:  "35480",
				Name: "Mirko Wlk",
			}, Employee: Employee{
				Id:       1,
				Position: "CEO",
			},
		}},
	}

	originalGetEmployeeById := GetEmployeeById
	originalPersonByDni := GetPersonByDni

	for _, item := range table {
		item.mockFunc()

		ft, err := GetFullTimeEmployeeById(item.id, item.dni)

		if err != nil {
			t.Errorf(" error when getting employee")
		}

		if ft.Age != item.ftExpected.Age {
			t.Errorf("age not expected")
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDni = originalPersonByDni
}
