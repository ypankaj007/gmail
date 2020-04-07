package smtp

import "testing"

func TestSend(t *testing.T) {
	from := "xxxxxxxxx@xxxxxxx.xxx"
	pass := "xxxxxxxx"
	to := []string{"xxxxx@xxxx.xx"}
	message := []byte("To: xxxxx@xxxx.xx\r\n" +
		"Subject: Testing\r\n" +
		"\r\n" +
		"Here’s the space for our great sales pitch\r\n")
	cl := Client{Host: "smtp.gmail.com", Port: "587"}
	err := cl.Send(from, pass, to, message)
	if err != nil {
		t.Errorf("Send email")
	}
}
