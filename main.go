package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jordan-wright/email"
	"net/smtp"
	"os"
	"time"
	_ "time"
)

var yesterday = time.Now().AddDate(0, 0, -1).Format("2006-01-02")

func main() {

	err := getCsv()
	if err != nil {
		fmt.Printf("Get csv returned an error %s", err)
	}

	e := email.NewEmail()
	password := "My app Password"
	from := "mercymaina@infinitytechafrica.com"
	e.From = "Infinity Tech <mercymaina@infinitytechafrica.com>"
	e.To = []string{"mercymaina567@gmail.com"}

	e.Subject = "MO Reports"
	e.HTML = []byte("<h1>Find attached MO report</h1>")
	file := yesterday + ".csv"
	e.AttachFile(file)
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", from, password, "smtp.gmail.com"))

}
func getCsv() error {
	db, err := sql.Open("mysql", "ucm:4rfvBHU*@tcp(44.239.52.145:3306)/ucm")
	if err != nil {
		// Handle error
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("Cant connect %s", err)
	}
	fmt.Printf("Ping successful ")

	query := fmt.Sprintf("SELECT org_id, network, mnc, mcc, cc, msisdn, flow, src_address, created_on FROM ucm.tbl_campaign_messages WHERE src_address='40023' AND DATE(created_on)='%s' AND flow='MO' LIMIT 100000", yesterday)
	rows, err := db.Query(query)
	//rows, err := db.Query("SELECT org_id,network,mnc,mcc,cc, msisdn,flow,src_address,created_on FROM ucm.tbl_campaign_messages where src_address=\"40023\" and date(created_on)=%s and flow =\"MO\" limit 100000", time.Now().AddDate(0, 0, -1).Format("2023-06-10"))
	if err != nil {
		// Handle error
		fmt.Printf("Error selecting %s", err)
	}
	fmt.Printf("select successful")
	defer rows.Close()
	//create csv file using encoding/csv
	filename := yesterday + ".csv"
	file, err := os.Create(filename)
	if err != nil {
		// Handle error
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header to the CSV file
	header := []string{"org_id", "network", "MNC", "MCC", "cc", "msisdn", "flow", "src_address", "created_on"}
	writer.Write(header)
	// Iterate over the result set
	for rows.Next() {
		// Process each row
		var org_id, network, mnc, mcc, cc, msisdn, flow, src_address, created_on string
		err := rows.Scan(&org_id, &network, &mnc, &mcc, &cc, &msisdn, &flow, &src_address, &created_on)
		if err != nil {
			// Handle error
		}
		row := []string{org_id, network, mnc, mcc, cc, msisdn, flow, src_address, created_on}
		writer.Write(row)
	}
	//
	if err := rows.Err(); err != nil {
		// Handle error
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		// Handle error
	}

	return nil
}
