package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
)

const (
	bannerHt = 94.0
	xIndent  = 40.0
)

var (
	discount     float64
	subtotal     float64
	customerName string
)

type LineItem struct {
	UnitName       string
	PricePerUnit   float64
	UnitsPurchased int
}

func ShowInvoice() {
	customerName = input.Selected

	address := GetClientAddress(customerName)
	clientAddress := fmt.Sprintf("%v", address[0]["address"])

	discount = formDiscount
	fmt.Println(formDiscount)
	var lineItems []LineItem
	for i := 1; i < len(invoiceData); i++ {
		price, _ := strconv.ParseFloat(invoiceData[i][2], 32)
		qty, _ := strconv.Atoi(invoiceData[i][1])
		tmp := LineItem{
			UnitName:       invoiceData[i][0],
			PricePerUnit:   price,
			UnitsPurchased: qty,
		}
		lineItems = append(lineItems, tmp)
		subtotal += price * float64(qty)
	}

	total := subtotal - discount
	totalBDT := toBDT(total)

	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	// fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	// Top Maroon Banner
	pdf.SetFillColor(19, 143, 214)
	pdf.Polygon([]gofpdf.PointType{
		{X: 0, Y: 0},
		{X: w, Y: 0},
		{X: w, Y: bannerHt},
		{X: 0, Y: bannerHt * 0.9},
	}, "F")
	pdf.Polygon([]gofpdf.PointType{
		{X: 0, Y: h},
		{X: 0, Y: h - (bannerHt * 0.2)},
		{X: 0, Y: h - (bannerHt * 0.1)},
		{X: w, Y: h},
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
	_, sy := summaryBlock(pdf, xIndent, bannerHt+lineHt*2.0, "Billed To", customerName, clientAddress)
	summaryBlock(pdf, xIndent*2.0+lineHt*12.5, bannerHt+lineHt*2.0, "Invoice Number", "INVSL#0001")

	time := time.Now().Format("01/02/2006")

	summaryBlock(pdf, xIndent*2.0+lineHt*12.5, bannerHt+lineHt*6.25, "Date of Issue", time)

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
	_, y = x-2.0, y+lineHt*1.25

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
	x, y = trailerLine(pdf, x, y, "Discount", discount)
	pdf.SetDrawColor(92, 90, 84)
	pdf.Line(x+20.0, y, x+220.0, y)
	y = y + lineHt*0.5
	_, _ = trailerLine(pdf, x, y, "Total", total)
	// x, y = trailerLine(pdf, x, y, "Pay", 525)
	y = y + lineHt*1.2
	_, _ = trailerLine(pdf, x, y, "Pay", 5)
	y = y + lineHt*1.5
	pdf.SetDrawColor(92, 90, 84)
	pdf.Line(x+20.0, y, x+220.0, y)
	_, _ = trailerLine(pdf, x, y, "Due", 10)

	pathPrefix, _ := CreateFolderUseingDate()
	// rand := rand.Intn(99999)
	fileName := "pdf1"
	folderPath := fmt.Sprintf("pdf/%v/%v.pdf", pathPrefix, fileName)

	err := pdf.OutputFileAndClose(folderPath)
	if err != nil {
		fmt.Println(err)
	}
}

func trailerLine(pdf *gofpdf.Fpdf, x, y float64, label string, amount float64) (float64, float64) {
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

func toBDT(poisha float64) string {
	roundPaisa := math.Round(poisha)
	return fmt.Sprintf("TK %.0f", roundPaisa)
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
	pdf.CellFormat(100.0, lineHt, fmt.Sprintf("%.2f", lineItem.PricePerUnit), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = x + 100.0
	pdf.MoveTo(x, y)
	pdf.CellFormat(80.0, lineHt, fmt.Sprintf("%d", lineItem.UnitsPurchased), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = w - xIndent - 2.0 - 119.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(119.5, lineHt, toBDT(lineItem.PricePerUnit*float64(lineItem.UnitsPurchased)), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
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
