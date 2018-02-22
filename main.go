package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/kbougy/mfatable/table"
)

func main() {
	s := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	table.WaitForCredentialReport(s)
	table.PrintCredentialReport(s)
}
