package Interface

import "fmt"

func Main() {
	AnyLearn()
}

type smsSending interface {
	send(to string, message string)
}

type TwilloSMS struct{}

func (s *TwilloSMS) send(to string, message string) {
	fmt.Printf("Sending OTP To %s, OTP : %s from Twillo\n", to, message)
}

func (s *TwilloSMS) checkBalance() {
	fmt.Printf("Checking Balance : 3lkhs\n")
}

type SnsSMS struct{}

func (s *SnsSMS) send(to string, message string) {
	fmt.Printf("Sending OTP To %s, OTP : %s from SNS\n", to, message)
}

func sending(s smsSending, otp string) {
	s.send("Gurudas", otp)
}

func check(s smsSending) {
	t, ok := s.(*TwilloSMS)
	if ok {
		t.checkBalance()
	} else {
		fmt.Println("Error")
	}
}
