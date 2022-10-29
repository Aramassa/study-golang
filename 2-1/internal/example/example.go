package example

type Example struct {
	name string
	age  int
}

func (e Example) MyName() string {
	return e.name
}
