package Interface

import "fmt"

func Main() {
	smsFromTwillo := TwilloSMS{}
	sending(&smsFromTwillo, "234235")

	smsFromSNS := SnsSMS{}
	sending(&smsFromSNS, "367658")
}

type smsSending interface {
	send(to string, message string)
}

type TwilloSMS struct{}

func (s *TwilloSMS) send(to string, message string) {
	fmt.Printf("Sending OTP To %s, OTP : %s from Twillo\n", to, message)
}

type SnsSMS struct{}

func (s *SnsSMS) send(to string, message string) {
	fmt.Printf("Sending OTP To %s, OTP : %s from SNS\n", to, message)
}

func sending(s smsSending, otp string) {
	s.send("Gurudas", otp)
}
