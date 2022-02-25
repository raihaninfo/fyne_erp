package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

const (
	bannerHt = 94.0
	xIndent  = 40.0
	taxRate  = 0.09
)

type LineItem struct {
	UnitName       string
	PricePerUnit   int
	UnitsPurchased int
}

func ShowInvoice() {
	lineItems := []LineItem{
		{
			UnitName:       "2x6 Lumber - 8'",
			PricePerUnit:   375, // in poisa
			UnitsPurchased: 220,
		}, {
			UnitName:       "Drywall Sheet",
			PricePerUnit:   822, // in poisa
			UnitsPurchased: 50,
		}, {
			UnitName:       "Paint",
			PricePerUnit:   1455, // in poisa
			UnitsPurchased: 3,
		}, {
			UnitName:       "This is a line",
			PricePerUnit:   3211, // in poisa
			UnitsPurchased: 3,
		}, {
			UnitName:       "Paint",
			PricePerUnit:   5, // in poisa
			UnitsPurchased: 3300,
		}, {
			UnitName:       "Paint",
			PricePerUnit:   332, // in poisa
			UnitsPurchased: 44,
		},
	}
	subtotal := 0
	for _, li := range lineItems {
		subtotal += li.PricePerUnit * li.UnitsPurchased
	}
	tax := int(float64(subtotal) * taxRate)
	total := subtotal + tax
	totalBDT := toBDT(total)

	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	// Top Maroon Banner
	pdf.SetFillColor(19, 143, 214)
	pdf.Polygon([]gofpdf.PointType{
		{0, 0},
		{w, 0},
		{w, bannerHt},
		{0, bannerHt * 0.9},
	}, "F")
	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - (bannerHt * 0.2)},
		{w, h - (bannerHt * 0.1)},
		{w, h},
	}, "F")

	// Banner - INVOICE
	pdf.SetFont("arial", "B", 40)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt := pdf.GetFontSize()
	pdf.Text(xIndent, bannerHt-(bannerHt/3.0)+lineHt/4.1, "miniERP")

	// Banner - Logo
	pdf.ImageOptions("img/master.png", 248.0, 0+(bannerHt-(bannerHt/1.5))/2.0, 0, bannerHt/1.5, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")

	// Banner - Phone, email, domain
	pdf.SetFont("arial", "", 12)
	pdf.SetTextColor(0, 0, 0)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w-xIndent-2.0*124.0, (bannerHt-(lineHt*1.5*3.0))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "+880 1819 523 690\nmtmart@yahoo.com\nwww.romdevelop.com", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// Banner - Address
	pdf.SetFont("arial", "", 12)
	pdf.SetTextColor(0, 0, 0)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w-xIndent-124.0, (bannerHt-(lineHt*1.5*3.0))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "2no gate Nasirabad,\n86/26 Sheikh Farid Market,Ground Floor.\n Chittagong.", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// Summary - Billed To, Invoice #, Date of Issue
	_, sy := summaryBlock(pdf, xIndent, bannerHt+lineHt*2.0, "Billed To", "Md Abu Raihan", "Shimulia,khoksa,", "Kushtia, 7021")
	summaryBlock(pdf, xIndent*2.0+lineHt*12.5, bannerHt+lineHt*2.0, "Invoice Number", "INVSL#0001")
	summaryBlock(pdf, xIndent*2.0+lineHt*12.5, bannerHt+lineHt*6.25, "Date of Issue", "05/02/2022")

	// Summary - Invoice Total
	x, y := w-xIndent-124.0, bannerHt+lineHt*2.25
	pdf.MoveTo(x, y)
	pdf.SetFont("times", "", 14)
	_, lineHt = pdf.GetFontSize()
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(100.0, lineHt, "Invoice Total", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x, y = x+2.0, y+lineHt*1.5
	pdf.MoveTo(x, y)
	pdf.SetFont("times", "", 48)
	_, lineHt = pdf.GetFontSize()
	alpha := 58
	pdf.SetTextColor(72+alpha, 42+alpha, 55+alpha)
	pdf.CellFormat(124.0, lineHt, totalBDT, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x, y = x-2.0, y+lineHt*1.25

	if sy > y {
		y = sy
	}
	x, y = xIndent-20.0, y+30.0
	pdf.Rect(x, y, w-(xIndent*2.0)+40.0, 3.0, "F")

	// Line Items - headers
	pdf.SetFont("times", "", 14)
	_, lineHt = pdf.GetFontSize()
	pdf.SetTextColor(52, 74, 110)
	x, y = xIndent-2.0, y+lineHt
	pdf.MoveTo(x, y)
	pdf.CellFormat(w/2.65+1.5, lineHt, "Description", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignLeft, false, 0, "")
	x = x + w/2.65 + 1.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(100.0, lineHt, "Price Per Unit", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = x + 100.0
	pdf.MoveTo(x, y)
	pdf.CellFormat(80.0, lineHt, "Quantity", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = w - xIndent - 2.0 - 119.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(119.5, lineHt, "Amount", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")

	// Line Items - real data
	y = y + lineHt
	for _, li := range lineItems {
		x, y = lineItem(pdf, x, y, li)
	}

	// Subtotal etc
	x, y = w/1.75, y+lineHt*2.25
	x, y = trailerLine(pdf, x, y, "Subtotal", subtotal)
	x, y = trailerLine(pdf, x, y, "Tax", tax)
	pdf.SetDrawColor(255, 255, 180)
	pdf.Line(x+20.0, y, x+220.0, y)
	y = y + lineHt*0.5
	x, y = trailerLine(pdf, x, y, "Total", total)

	// Grid
	// drawGrid(pdf)
	err := pdf.OutputFileAndClose("miniERP.pdf")
	if err != nil {
		panic(err)
	}
}

func trailerLine(pdf *gofpdf.Fpdf, x, y float64, label string, amount int) (float64, float64) {
	origX := x
	w, _ := pdf.GetPageSize()
	pdf.SetFont("times", "", 14)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(180, 180, 180)
	pdf.MoveTo(x, y)
	pdf.CellFormat(80.0, lineHt, label, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = w - xIndent - 2.0 - 119.5
	pdf.MoveTo(x, y)
	pdf.SetTextColor(50, 50, 50)
	pdf.CellFormat(119.5, lineHt, toBDT(amount), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	y = y + lineHt*1.5
	return origX, y
}

func toBDT(poisa int) string {
	poisaStr := fmt.Sprintf("%d", poisa%100)
	if len(poisaStr) < 2 {
		poisaStr = "0" + poisaStr
	}
	return fmt.Sprintf("TK %d.%s", poisa/100, poisaStr)
}

func lineItem(pdf *gofpdf.Fpdf, x, y float64, lineItem LineItem) (float64, float64) {
	origX := x
	w, _ := pdf.GetPageSize()
	pdf.SetFont("times", "", 14)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(50, 50, 50)
	pdf.MoveTo(x, y)
	x, y = xIndent-2.0, y+lineHt*.75
	pdf.MoveTo(x, y)
	pdf.MultiCell(w/2.65+1.5, lineHt, lineItem.UnitName, gofpdf.BorderNone, gofpdf.AlignLeft, false)
	tmp := pdf.SplitLines([]byte(lineItem.UnitName), w/2.65+1.5)
	maxY := y + float64(len(tmp)-1)*lineHt
	x = x + w/2.65 + 1.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(100.0, lineHt, toBDT(lineItem.PricePerUnit), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = x + 100.0
	pdf.MoveTo(x, y)
	pdf.CellFormat(80.0, lineHt, fmt.Sprintf("%d", lineItem.UnitsPurchased), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = w - xIndent - 2.0 - 119.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(119.5, lineHt, toBDT(lineItem.PricePerUnit*lineItem.UnitsPurchased), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	if maxY > y {
		y = maxY
	}
	y = y + lineHt*1.75
	pdf.SetDrawColor(180, 180, 180)
	pdf.Line(xIndent-10.0, y, w-xIndent+10.0, y)
	return origX, y
}

func summaryBlock(pdf *gofpdf.Fpdf, x, y float64, title string, data ...string) (float64, float64) {
	pdf.SetFont("times", "", 14)
	pdf.SetTextColor(180, 180, 180)
	_, lineHt := pdf.GetFontSize()
	y = y + lineHt
	pdf.Text(x, y, title)
	y = y + lineHt*.25
	pdf.SetTextColor(50, 50, 50)
	for _, str := range data {
		y = y + lineHt*1.25
		pdf.Text(x, y, str)
	}
	return x, y
}

func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) {
		pdf.SetTextColor(200, 200, 200)
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (w / 20.0) {
		if y < bannerHt*.9 {
			pdf.SetTextColor(200, 200, 200)
		} else {
			pdf.SetTextColor(80, 80, 80)
		}
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}
