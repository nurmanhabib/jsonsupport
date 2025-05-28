package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v3"
	"github.com/xuri/excelize/v2"

	"jsonsupport/app/entity"
)

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{},
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "input",
				Value:       "",
				Destination: nil,
				UsageText:   "",
				Config:      cli.StringConfig{},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			input := cmd.StringArg("input")
			outputDir := "outputs"

			err := os.MkdirAll(outputDir, os.ModePerm)
			if err != nil {
				return fmt.Errorf("create output directory error: %w", err)
			}

			// Step 1: Load JSON file
			jsonFile, err := os.ReadFile(input)
			if err != nil {
				return cli.Exit(fmt.Sprintf("File input \"%s\" tidak dapat di buka", input), 2)
			}

			var tickets []entity.Ticket
			err = json.Unmarshal(jsonFile, &tickets)
			if err != nil {
				panic(fmt.Errorf("error unmarshalling JSON: %w", err))
			}

			// Parse CustomFieldsRaw to struct
			for i := range tickets {
				clean := strings.ReplaceAll(tickets[i].CustomFieldsRaw, "\\n", "")
				if err := json.Unmarshal([]byte(clean), &tickets[i].CustomFields); err != nil {
					fmt.Printf("Error parsing custom fields for ticket %s: %v\n", tickets[i].Title, err)
				}
			}

			// Step 3: Write to excel
			f := excelize.NewFile()
			sheet := "Sheet1"

			// Header
			headers := []string{"Client Name", "Customer Name", "Status", "", "Ticket Category", "Module", "Detail Module", "Created At", "Title", "", "Solved At", "Date"}
			for i, h := range headers {
				cell, _ := excelize.CoordinatesToCellName(i+1, 1)
				f.SetCellValue(sheet, cell, h)
			}

			// Data rows
			for i, row := range tickets {
				f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), row.CustomFields["Client Name"])
				f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), row.CustomerName)
				f.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), mappingStatus(row.Status))
				f.SetCellValue(sheet, fmt.Sprintf("D%d", i+2), "")
				f.SetCellValue(sheet, fmt.Sprintf("E%d", i+2), row.CustomFields["Ticket Category"])
				f.SetCellValue(sheet, fmt.Sprintf("F%d", i+2), row.CustomFields["Module"])
				f.SetCellValue(sheet, fmt.Sprintf("G%d", i+2), row.CustomFields["Detail Module"])
				f.SetCellValue(sheet, fmt.Sprintf("H%d", i+2), convertTimeFormat("January 2, 2006, 3:04 PM", row.CreatedAt, "15:04"))
				f.SetCellValue(sheet, fmt.Sprintf("I%d", i+2), row.Title)
				f.SetCellValue(sheet, fmt.Sprintf("J%d", i+2), "")
				f.SetCellValue(sheet, fmt.Sprintf("K%d", i+2), convertTimeFormat("January 2, 2006, 3:04 PM", row.ResolvedAt, "15:04"))
				f.SetCellValue(sheet, fmt.Sprintf("L%d", i+2), convertTimeFormat("January 2, 2006, 3:04 PM", row.ResolvedAt, "2006-01-02"))
			}

			f.SetColWidth(sheet, "A", "I", 30)

			outputFile := fmt.Sprintf("%s/output_%s.xlsx", outputDir, time.Now().Format("2006_01_02__15_04_05"))

			// Simpan file Excel
			if err := f.SaveAs(outputFile); err != nil {
				panic(fmt.Errorf("error saving Excel file: %w", err))
			}

			fmt.Printf("Data berhasil ditulis ke %s\n", outputFile)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func convertTimeFormat(inputLayout string, timeStr string, outputLayout string) string {
	timed, err := time.Parse(inputLayout, timeStr)
	if err != nil {
		return "-"
	}
	return timed.Format(outputLayout)
}

func mappingStatus(status string) string {
	switch status {
	case "Resolved":
		return "Solved"
	case "On-Hold":
		return "L3"
	case "Open":
		return "L2"
	case "Pending":
		return "L1"
	}
	return status
}
