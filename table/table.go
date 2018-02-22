package table

import (
	"encoding/csv"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strings"
)

func TrimRecord(record []string) (trimmed []string) {
	t := []string{record[0], record[3], record[7]}
	return t
}

func PrintRecords(records [][]string) {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"User", "Password Enabled", "MFA Status"})
	t.SetBorder(false)
	for _, record := range records[1:] {
		trimmed := TrimRecord(record)
		t.Append(trimmed)
	}
	t.Render()
}

func WaitForCredentialReport(s *session.Session) {
	iam_client := iam.New(s)
	generate_credential_report, err := iam_client.GenerateCredentialReport(
		&iam.GenerateCredentialReportInput{},
	)

	if err != nil {
		log.Fatal(err)
	}

	if *generate_credential_report.State != "COMPLETE" {
		WaitForCredentialReport(s)
	}
	return
}

func PrintCredentialReport(s *session.Session) {
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

	PrintRecords(records)
}
