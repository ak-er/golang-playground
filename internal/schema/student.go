package schema

type Student struct {
	Id     int64  `json:"id"`
	Name   string `json:"name" validate:"required"`
	Course string `json:"course" validate:"required"`
	City   string `json:"city" validate:"required"`
	Age    int    `json:"age" validate:"required"`
}
