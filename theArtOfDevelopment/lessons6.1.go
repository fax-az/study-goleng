package main

var _ User = &user{}

type user struct {
	FIO, Adress, Phone string
}

func (u *user) ChangeFIO(string) int {
	u.FIO = newFIO
	return 0
}

func (u *user) ChangeFddress(string) {
	u.Adress = newAdress
}

type User interface {
	CangeFIO(newFIO string) int
	ChangeFddress(newAdress string)
}

func NewUser(fio, adress, phone string) User {
	u := user{
		FIO:    fio,
		Adress: adress,
		Phone:  phone,
	}
	return &u
}

func main() {

}
