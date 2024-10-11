package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	db := &Sqlite{}
	err := db.New()
	if err != nil {
		log.Fatal("Can't open database: ", err)
	}

	err = db.CreateTable()
	if err != nil {
		log.Fatal("Can't create database: ", err)
	}

	defer db.db.Close()

	cmdFlag := flag.NewFlagSet("cmdFlag", flag.ExitOnError)
	username := cmdFlag.String("u", "", "username for the email server")
	password := cmdFlag.String("p", "", "password for the email server")
	smtpHost := cmdFlag.String("host", "", "smtp host for the email server")
	smtpPort := cmdFlag.String("port", "", "smtp port for the email server")
	thingsEmail := cmdFlag.String("things", "", "things email for the email server")
	title := cmdFlag.String("t", "", "task title for things")
	desc := cmdFlag.String("d", "", "description for the task")

	cmdFlag.Parse(os.Args[1:])

	credential := &Credential{}
	emailServer := &EmailServer{}
	things3 := &Things3{}
	Email := &Email{}

	switch {
	case *username != "" && *password != "":
		credential.Username = *username
		credential.Password = *password

		err := db.InsertCredentials(credential)
		if err != nil {
			log.Fatal(err)
		}

	case *smtpHost != "" && *smtpPort != "":
		emailServer.SMTPHost = *smtpHost
		emailServer.SMTPPort = *smtpPort

		err := db.InsertEmailServer(emailServer)
		if err != nil {
			log.Fatal(err)
		}

	case *thingsEmail != "":
		things3.Email = *thingsEmail

		err := db.InsertThings3(things3)
		if err != nil {
			log.Fatal(err)
		}

	case *title != "" || *desc != "":
		cred, err := db.GetLastCredential()
		if err != nil {
			log.Fatal("No credential found, please add credential first by using -u and -p")
		}
		credential.Username = cred.Username
		credential.Password = cred.Password

		es, err := db.GetLastEmailServer()
		if err != nil {
			log.Fatal("No email server found, please add mail server first by using -host and -port")
		}
		emailServer.SMTPHost = es.SMTPHost
		emailServer.SMTPPort = es.SMTPPort

		things3, err := db.GetLastThings3()
		if err != nil {
			log.Fatal("No things email found, please add things email first by using -things")
		}

		Email.To = []string{things3.Email}
		Email.From = credential.Username
		Email.Subject = *title
		Email.Body = *desc

		err = Email.SendEmail(*credential, *emailServer)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Task has been added to Things 3")

	default:
		log.Fatal("Please provide the correct flag")
	}

	// validation

	_, err = db.GetLastEmailServer()
	if err != nil {
		log.Fatal("No email server found, please add mail server first by using -host and -port")
	}

	_, err = db.GetLastCredential()
	if err != nil {
		log.Fatal("No credential found, please add credential first by using -u and -p")
	}

	_, err = db.GetLastThings3()
	if err != nil {
		log.Fatal("No things email found, please add things email first by using -things")
	}

	if *title == "" {
		log.Fatal("Please provide the title")
	}
}
