package main

import "fmt"

type AuthMeth interface {
	Auth(int) error
}

type SocMedia struct {
	Name  string
	Email string
	Pswrd string
}

type PinCode struct {
	Name string
	Code string
}

type Password struct {
	Name string
	Pass string
}

func (s *SocMedia) Auth(n int) error {
	if len(s.Pswrd) < n {
		return fmt.Errorf("Password must be more than 6 characters long")
	}
	return nil
}

func (ps *Password) Auth(n int) error {
	if len(ps.Pass) < n {
		return fmt.Errorf("Password must be more than 6 characters long")
	}
	return nil
}

func (p *PinCode) Auth(n int) error {
	if len(p.Code) < n {
		return fmt.Errorf("Pincode must be more than 6 digits long")
	}
	return nil
}

func Enter(a AuthMeth) {
	switch a.(type) {
	case *PinCode:
		fmt.Println("Enter pin code")
	case *Password:
		fmt.Println("Enter password")
	case *SocMedia:
		fmt.Println("Enter Email and password")
	default:
		fmt.Println("Unknown entry method")
	}

	er := a.Auth(6)
	if er != nil {
		panic(er)
	}

	fmt.Println("You have entered")
}

func main() {
	Pas := &Password{"Aliakbar", "ALIAKBAR"}
	Pin := &PinCode{"Aizat", "123"}
	Soc := &SocMedia{"Erlan", "Erlan4ik@gmail.com", "987456"}
	Enter(Pas)
	Enter(Pin)
	Enter(Soc)
}
