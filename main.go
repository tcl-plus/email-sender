package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"gopkg.in/gomail.v2"
)

var (
	from           string
	to             string
	cc             string
	subject        string
	body           string
	server         string
	port           int
	user           string
	password       string
	attachmentPath string
)

func init() {
	pflag.StringVar(&from, "from", "", "Sender's email address with optional name (e.g., 'Name <user@example.com>' or 'user@example.com')")
	pflag.StringVar(&to, "to", "", "Recipient's email addresses (comma-separated)")

	pflag.StringVar(&cc, "cc", "", "CC email addresses (comma-separated)")
	pflag.StringVar(&subject, "subject", "", "Email subject")
	pflag.StringVar(&body, "body", "", "Email body")
	pflag.StringVar(&server, "server", "", "SMTP server address")
	pflag.IntVar(&port, "port", 0, "SMTP server port")
	pflag.StringVar(&user, "user", "", "SMTP server username")
	pflag.StringVar(&password, "password", "", "SMTP server password")
	pflag.StringVar(&attachmentPath, "attachment", "", "Path to the attachment file (comma-separated")
	pflag.Parse()
}

func checkRequiredParameters() {
	var missingParams []string
	if from == "" {
		missingParams = append(missingParams, "from")
	}
	if to == "" {
		missingParams = append(missingParams, "to")
	}
	if subject == "" {
		missingParams = append(missingParams, "subject")
	}
	if body == "" {
		missingParams = append(missingParams, "body")
	}
	if server == "" {
		server = os.Getenv("MAIL_SERVER")
		if server == "" {
			missingParams = append(missingParams, "server")
		}
	}
	if port == 0 {
		portStr := os.Getenv("MAIL_SERVER_PORT")
		if portStr != "" {
			var err error
			port, err = strconv.Atoi(portStr)
			if err != nil {
				fmt.Println("Invalid port: ", port)
				os.Exit(1)
			}
		}
		if port == 0 {
			missingParams = append(missingParams, "port")
		}
	}

	if user == "" {
		user = os.Getenv("MAIL_SERVER_USER")
		if user == "" {
			missingParams = append(missingParams, "user")
		}
	}
	if password == "" {
		password = os.Getenv("MAIL_SERVER_PASSWORD")
		if password == "" {
			missingParams = append(missingParams, "password")
		}
	}

	if len(missingParams) > 0 {
		fmt.Printf("Missing required parameter(s): %s\n", strings.Join(missingParams, ", "))
		pflag.PrintDefaults()
		os.Exit(1)
	}
}

func splitByComma(value string) []string {
	if value == "" {
		return nil
	}
	return strings.Split(value, ",")
}

func createAndSendEmail() {
	// Split email addresses
	toAddresses := splitByComma(to)
	ccAddresses := splitByComma(cc)
	attachmentPaths := splitByComma(attachmentPath)

	// Create and send email
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", toAddresses...)
	if ccAddresses != nil {
		m.SetHeader("Cc", ccAddresses...)
	}
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	for _, attachment := range attachmentPaths {
		m.Attach(attachment)
	}

	d := gomail.NewDialer(server, port, user, password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		os.Exit(1)
	}

	fmt.Println("Email sent successfully.")
}

func main() {
	checkRequiredParameters()
	createAndSendEmail()
}
