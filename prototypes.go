package main

type prototypes interface {
	catch()
	finally()
	then()
}

type Methods struct {
	OnRejected  error
	OnFulfilled interface{}
}

func (a Methods) catch() error {
	return a.OnRejected
}

func (a Methods) then() string {
	if a.OnFulfilled != nil {
		return "Received Data"
	} else {
		return "Data is Loading"
	}
}
func (a Methods) finally() interface{} {
	return a.OnFulfilled
}

type MyFloat string

func (b MyFloat) catch() string {
	var a = "skjdnck"
	return a
}
