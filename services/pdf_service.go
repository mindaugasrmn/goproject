package main

import (
	"github.com/jung-kurt/gofpdf"
	"log"
)

type Invoice struct {
	InvoiceTitle   string `json:"InvoiceTitle"`
	InvoiceNoName   string `json:"InvoiceNoName"`
	InvoiceDateName   string `json:"InvoiceDateName"`
	InvoicePayTermName string `json:"InvoicePayTermName"`
	InvoicePayTermDateName string `json:"InvoicePayTermDateName"`
	ClientNoName   string `json:"ClientNoName"`
    CompanyName   string `json:"CompanyName`
    CompanyAddress string `json:"CompanyAddress"`
}

func main() {
	y := 10.0
    var inv Invoice
    inv.CompanyName="Swedgarden AB"
    inv.CompanyAddress="Granitvagen 33"
    inv.InvoiceTitle="FAKTURA"
    inv.InvoiceNoName="Fakturanummer"
    inv.ClientNoName="Kundnummer"
    inv.InvoiceDateName="Betalningsvillkor"
    inv.InvoicePayTermName="Betalningsvillkor"
    inv.InvoicePayTermDateName="Förfallodatum"
    pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("arial", "", 13)
	pdf.Cell(10, 10.6, inv.CompanyName)
	
	pdf.SetFont("arial", "", 22)
	pdf.SetTextColor(48, 164, 245)
	pdf.Text(164,17.6, inv.InvoiceTitle)
	pdf.SetXY(0, y)
	y += 17.2
	pdf.SetLineWidth(.1)
	pdf.SetDrawColor(221, 221, 221)
	pdf.ClipRoundedRect(10, y, 120, 32.4, 3, true)
	pdf.ClipEnd()
	pdf.SetXY(5, 18.7)
	pdf.SetFont("times", "", 8)
	pdf.SetTextColor(128, 128, 128)
	pdf.Ln(-1)
	pdf.CellFormat(2, 5.6, "", "0", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5.6, inv.InvoiceNoName, "0", 0, "L", false, 0, "")
    pdf.CellFormat(50, 5.6, "aaaaaa", "0", 0, "R", false, 0, "")
    pdf.Ln(-1)
    pdf.CellFormat(2, 5.6, "", "0", 0, "R", false, 0, "")
    pdf.CellFormat(50, 5.6, inv.ClientNoName, "0", 0, "L", false, 0, "")
    pdf.CellFormat(50, 5.6, "aaaaaa", "0", 0, "R", false, 0, "")
    pdf.Ln(-1)
    pdf.CellFormat(2, 5.6, "", "0", 0, "R", false, 0, "")
    pdf.CellFormat(50, 5.6, inv.InvoiceDateName, "0", 0, "L", false, 0, "")
    pdf.CellFormat(50, 5.6, "aaaaaa", "0", 0, "R", false, 0, "")
    pdf.Ln(-1)
    pdf.CellFormat(2, 5.6, "", "0", 0, "R", false, 0, "")
    pdf.CellFormat(50, 5.6, inv.InvoicePayTermName, "0", 0, "L", false, 0, "")
    pdf.CellFormat(50, 5.6, "aaaaaa", "0", 0, "R", false, 0, "")
    pdf.Ln(-1)
    pdf.CellFormat(2, 5.6, "", "0", 0, "R", false, 0, "")
    pdf.CellFormat(50, 5.6, inv.InvoicePayTermDateName, "0", 0, "L", false, 0, "")
    pdf.CellFormat(50, 5.6, "aaaaaa", "0", 0, "R", false, 0, "")
    pdf.Ln(-1)





	y += 55
	pdf.ClipRect(10, y, 105, 20, true)
	pdf.SetFillColor(255, 255, 255)
	pdf.Rect(10, y, 105, 20, "F")
	pdf.ClipCircle(40, y+10, 15, false)
	pdf.RadialGradient(25, y, 30, 30, 220, 250, 220, 40, 60, 40, 0.3,
	    0.85, 0.3, 0.85, 0.5)
	pdf.ClipEnd()
	pdf.ClipEllipse(80, y+10, 20, 15, false)
	pdf.RadialGradient(60, y, 40, 30, 250, 220, 220, 60, 40, 40, 0.3,
	    0.85, 0.3, 0.85, 0.5)
	pdf.ClipEnd()
	pdf.ClipEnd()

	y += 28
	pdf.ClipEllipse(26, y+10, 16, 10, true)
	pdf.ClipEnd()

	pdf.ClipCircle(60, y+10, 10, true)
	pdf.RadialGradient(50, y, 20, 20, 220, 220, 250, 40, 40, 60, 0.3,
	    0.7, 0.3, 0.7, 0.5)
	pdf.ClipEnd()

	pdf.ClipPolygon([]gofpdf.PointType{{80, y + 20}, {90, y},
	    {100, y + 20}}, true)
	pdf.LinearGradient(80, y, 20, 20, 250, 220, 250, 60, 40, 60, 0.5,
	    1, 0.5, 0.5)
	pdf.ClipEnd()

	
	pdf.SetFont("Times", "", 12)
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Print(err.Error())
		return
	}
}