package main

import (
	"testing"
	"time"

	"github.com/Aramassa/study-golang/2-1/pkg/example"
)

func TestMyName(t *testing.T) {

	ex := example.Example{Name: "Yosuke", Age: 41}

	if got := ex.MyName(); got != "Yosuke" {
		t.Fatalf("not MyName, got %v", got)
	}
}

func add(a, b int) int {

	time.Sleep(time.Duration(a+b) * time.Second / 3)

	return a + b
}

func TestExample2(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "fail", args: args{a: 1, b: 2}, want: 30},
		{name: "normal", args: args{a: 1, b: 2}, want: 3},
	}
	for _, tt := range tests {

		tt := tt // 重要!!

		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()

			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Fatalf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
