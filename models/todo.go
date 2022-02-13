package models

type Todo struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (_ *Todo) TableName() string {
	return "todo"
}

