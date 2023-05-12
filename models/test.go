package models

type CreateTest struct {
	Name          string `validate:"required"`
	Desc          string `validate:"required"`
	Img           string `validate:"required"`
	Minute        int32  `validate:"required"`
	AgeCls        string `validate:"required"`
	BeforeDesc    string `validate:"required"`
	AfterDesc     string `validate:"required"`
	ExampleReport string
	IsActive      bool `validate:"required"`
}
