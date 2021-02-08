package util

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

func SendMail(to []string, subject, message string) error {
	smtp_port := os.Getenv("SMTP_PORT")
	smtp_ports, err := strconv.Atoi(smtp_port)
	if err != nil {
		return err
	}
	body := "From: " + os.Getenv("SMPTP_SENDER_NAME") + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", os.Getenv("SMTP_CONFIG_AUTH_EMAIL"), os.Getenv("SMTP_CONFIG_AUTH_PASSWORD"), os.Getenv("SMTP_HOST"))
	smtpAddr := fmt.Sprintf("%s:%d", os.Getenv("SMTP_HOST"), smtp_ports)

	err = smtp.SendMail(smtpAddr, auth, os.Getenv("SMTP_CONFIG_AUTH_EMAIL"), append(to), []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func SendMailTelkom(tos string, subject, messages string) error {
	from := mail.Address{"", "admin.nprm@telkom.co.id"}
	to := mail.Address{"", tos}
	subj := subject
	body := messages
	// Setup headers
	headers := make(map[string]string)
	fmt.Printf(from.String())
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	fmt.Printf(message)
	auth := smtp.PlainAuth("", os.Getenv("SMTP_CONFIG_AUTH_EMAIL"), os.Getenv("SMTP_CONFIG_AUTH_PASSWORD"), os.Getenv("SMTP_HOST"))
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         os.Getenv("SMTP_HOST"),
	}

	conn, err := tls.Dial("tcp", os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, os.Getenv("SMTP_HOST"))
	if err != nil {
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

	return err

}

func SendMailVerify(email string, firstname string, code string, verifyExp string) error {
	subject := "NPRM Email Verification"
	message := "Halo " + firstname + ",\n \n" +
		"Berikut link aktifasi akun NPRM anda : " + "\n \n" +
		"http://" + os.Getenv("IP") + ":8080/verify/activation/" + code + "\n \n" +
		"Mohon segera melakukan konfirmasi akun anda sebelum " + verifyExp + "\n \n" +
		"Terimakasih"
	err := SendMailTelkom(email, subject, message)
	return err
}

func SendMailForget(email string, firstname string, code string, verifyExp string) error {
	// to := []string{email}
	subject := "NPRM Email Request Forget Password"
	message := "Halo " + firstname + ",\n \n" +
		"Berikut link untuk reset password akun NPRM anda : " + "\n \n" +
		"http://" + os.Getenv("IP") + ":8080/verify/forget/" + code + "\n \n" +
		"Mohon segera melakukan reset password akun anda sebelum " + verifyExp + "\n \n" +
		"Terimakasih"
	err := SendMailTelkom(email, subject, message)
	return err
}
