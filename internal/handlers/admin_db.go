package handlers

import "context"

type StudentRow struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type StaffRow struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

func getAllStudents() ([]StudentRow, error) {
	rows, err := DB.Query(context.Background(),
		`SELECT first_name, last_name, email FROM students`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []StudentRow
	for rows.Next() {
		var s StudentRow
		rows.Scan(&s.FirstName, &s.LastName, &s.Email)
		students = append(students, s)
	}

	return students, nil
}

func getAllStaff() ([]StaffRow, error) {
	rows, err := DB.Query(context.Background(),
		`SELECT first_name, last_name, email, role FROM staff`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var staff []StaffRow
	for rows.Next() {
		var s StaffRow
		rows.Scan(&s.FirstName, &s.LastName, &s.Email, &s.Role)
		staff = append(staff, s)
	}

	return staff, nil
}
