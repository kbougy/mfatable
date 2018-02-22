package main

import (
	"encoding/csv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strings"
)

func print_records(records [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"User", "Password Enabled", "MFA Status"})
	table.SetBorder(false)
	for _, record := range records[1:] {
		trimmed_record := []string{record[0], record[3], record[7]}
		table.Append(trimmed_record)
	}
	table.Render()
}

func wait_for_credential_report(s *session.Session) {
	iam_client := iam.New(s)
	generate_credential_report, err := iam_client.GenerateCredentialReport(
		&iam.GenerateCredentialReportInput{},
	)

	if err != nil {
		log.Fatal(err)
	}

	if *generate_credential_report.State != "COMPLETE" {
		wait_for_credential_report(s)
	}
	return
}

func print_credential_report(s *session.Session) {
	iam_client := iam.New(s)
	get_credential_report, err := iam_client.GetCredentialReport(
		&iam.GetCredentialReportInput{},
	)

	if err != nil {
		log.Fatal(err)
	}

	// convert bytes to string
	string_content := string(get_credential_report.Content[:])
	r := csv.NewReader(strings.NewReader(string_content))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	print_records(records)
}

func main() {
	s := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	wait_for_credential_report(s)
	print_credential_report(s)
}
