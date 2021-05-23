package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/tidwall/buntdb"
)

type Employee struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Salary  int32  `json:"salary"`
	Active  bool   `json:"is_active"`
}

func (e *Employee) Prepare() {
	e.Name = html.EscapeString(strings.TrimSpace(e.Name))
	e.Contact = html.EscapeString(strings.TrimSpace(e.Contact))
	id, _ := uuid.NewUUID()
	e.ID = id.String()
}

func (e *Employee) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if e.Name == "" {
			return errors.New("Required Name")
		}
		if e.Salary == 0 {
			return errors.New("Required Employee Salary")
		}
		if e.Contact == "" {
			return errors.New("Required Contact")
		}

		return nil

	default:
		if e.Name == "" {
			return errors.New("Required Name")
		}
		if e.Salary == 0 {
			return errors.New("Employee Salary must be greater than ZERO")
		}
		if e.Contact == "" {
			return errors.New("Required Contact")
		}
		return nil
	}
}

func (e *Employee) AddEmployee(db *buntdb.DB) (*Employee, error) {

	val, _ := json.Marshal(e)

	fmt.Println(string(val))

	err := db.Update(func(tx *buntdb.Tx) error {
		tx.Set(e.ID, string(val), nil)
		return nil
	})
	if err != nil {
		return &Employee{}, err
	}
	return e, nil
}

func (u *Employee) FindAllEmployees(db *buntdb.DB) (*[]Employee, error) {
	var err error
	users := []Employee{}

	err = db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("name", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			emp := Employee{}
			err = json.Unmarshal([]byte(value), &emp)

			fmt.Println(value)

			users = append(users, emp)
			return true
		})
		return nil
	})

	if err != nil {
		return &[]Employee{}, err
	}
	return &users, err
}

func FindByID(db *buntdb.DB, uid string) (*Employee, error) {
	e := Employee{}
	err := db.View(func(tx *buntdb.Tx) error {
		value, err := tx.Get(uid)

		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(value), &e)
		return nil
	})
	if err != nil {
		return &Employee{}, err
	}
	return &e, err
}

func (e *Employee) UpdateAUser(db *buntdb.DB, uid string) (*Employee, error) {
	e.ID = uid
	val, _ := json.Marshal(e)

	err := db.Update(func(tx *buntdb.Tx) error {
		_, res, err := tx.Set(uid, string(val), nil)

		if err != nil {
			return err
		}
		if res {
			fmt.Println("Successfully Replaced Value")
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return e, nil
}

func DeleteEmp(db *buntdb.DB, uid string) error {

	err := db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(uid)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
