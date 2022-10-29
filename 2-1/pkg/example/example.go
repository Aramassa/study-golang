package example

type Example struct {
	Name string
	Age  int
}

func (e Example) MyName() string {
	return e.Name
}
