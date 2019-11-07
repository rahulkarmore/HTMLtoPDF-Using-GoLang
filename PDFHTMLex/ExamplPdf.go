package main

import (
	"fmt"

	"log"
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return
	}
	htmlStr := `<!DOCTYPE html>
	<html>
	<head>
	<style>
	#customers {
	  font-family: "Trebuchet MS", Arial, Helvetica, sans-serif;
	  border-collapse: collapse;
	  width: 100%;
	}
	
	#customers td, #customers th {
	  border: 1px solid #ddd;
	  padding: 8px;
	}
	
	#customers tr:nth-child(even){background-color: #f2f2f2;}
	
	#customers tr:hover {background-color: #ddd;}
	
	#customers th {
	  padding-top: 12px;
	  padding-bottom: 12px;
	  text-align: left;
	  background-color: #4CAF50;
	  color: white;
	}
	#red{
		color:red;
		size=2;
	}
	</style>
	</head>
	<body>
	
	
		<h1>Company Name :  MSPGCL</h1>
		<h3>Unit Name : CENTRAL PURCHASE AGENCY (MUMBAI)</h3>
		<h3  id="red">Kindly go through the details, In case of any problems/ issues related to the reported data email to cgmstores@mahagenco.in</h3>
		<br>
		<!-- <img src="logo_mkcl_w.png" alt="Smiley face"> -->
	
	<!-- table 1 -->
		<table id="customers">
	  <tr>
		<th>ENQUIRY </th>
		<th>DETAILS</th>
	  
	  </tr>
	  <tr>
		<td>Enquiry Cod</td>
		<td> 63941 </td>
		
	  </tr>
	  <tr>
		<td>Mode</td>
		<td>Quantity Contract </td>
		
	  </tr>
	  <tr>
		<td>Is Over All</td>
		<td>Yes</td>
		
	  </tr>
	  <tr>
		<td>Enquiry Type</td>
		<td> LIMITED TENDER </td>
	   
	  </tr>
	  <tr>
		<td>Enquiry Category Type</td>
		<td>SUPPLY</td>
		
	  </tr>
	  <tr>
		<td> Section</td>
		<td>CPAGR1 </td>
		
	  </tr>
	  <tr>
		<td>Type Of Bid</td>
		<td>Single Bid </td>
	  
	  </tr>
	  <tr>
		<td> Material Description </td>
		<td> testing overall for legend in price bid. </td>
		
	  </tr>
	  <tr>
		<td>Estimated Cost [In Rupees]</td>
		<td>NA</td>
		
	  </tr>
	  <tr>
		<td>Submission Date</td>
		<td> Friday, November 16, 2018 5:39:00 PM </td>
		
	  </tr>
	
	  <!--  -->
	  <tr>
			<td> Enquiry Validity</td>
			<td>100 Days</td>
			
		  </tr>
		  <tr>
			<td>Delivery Period </td>
			<td>80 Days</td>
			
		  </tr>
		  <tr>
			<td> Contact Email Id</td>
			<td>  cgmstores@mahagenco.in </td>
			
		  </tr>
		  <!--  -->
		  <tr>
				<td>  Commencement Period </td>
				<td> 40 days from 'First order'</td>
				
			  </tr>
			  <tr>
				<td> Basis Of Prices </td>
				<td> Firm Price</td>
				
			  </tr>
			  <tr>
				<td> Sample Required</td>
				<td>   Not Required </td>
				
			  </tr>
			</tr>
			<tr>
			  <td>  Deviation </td>
			  <td>   NO </td>
			  
			</tr>
	
			<tr>
					<td> BID OPENING DETAILS ( Single Bid)  </td>
					<td>  Will be declared later </td>
					
				  </tr>
	
				  
	</table>
	<br>
	<!-- table 2 -->
	<table id="customers" >
			<tr>
					<th>ENQUIRY DOCUMENTS</th>
					
				  
				  </tr>
				  <tr>
					<td>Total Number of documents : 1</td>
				   
					
				  </tr>
	</table>
	</body>
	</html>
	`

	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	//Your Pdf Name
	err = pdfg.WriteFile("Your_pdfname.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

//  this code for importing html file in go and genrate pdf
// but problem is we cant not applay css effect to pdf
// so above code is working fine
// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
// )

// const (
// 	OrientationLandscape = "Landscape" // Landscape mode
// 	OrientationPortrait  = "Portrait"  // Portrait mode
// )

// func main() {
// 	path := "D:/workspace/src/github.com/SebastiaanKlippert/go-wkhtmltopdf/wkhtmltopdf.go"
// 	wkhtmltopdf.SetPath(path)
// 	pdfg, err := wkhtmltopdf.NewPDFGenerator()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Set global options
// 	pdfg.Dpi.Set(300)
// 	pdfg.Orientation.Set(OrientationLandscape)
// 	pdfg.Grayscale.Set(true)

// 	// Create a new input page from an URL
// 	page := wkhtmltopdf.NewPage("input.html")

// 	// Set options for this page
// 	page.FooterRight.Set("[page]")
// 	page.FooterFontSize.Set(10)
// 	page.Zoom.Set(0.95)

// 	// Add to document
// 	pdfg.AddPage(page)

// 	// Create PDF document in internal buffer
// 	err = pdfg.Create()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Write buffer contents to file on disk
// 	err = pdfg.WriteFile("simplesample.pdf")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Done")

// }
// package main

// import (
// 	"fmt"
// 	"log"

// 	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
// 	"github.com/jung-kurt/gofpdf"
// )

// func main() {

// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.AddPage()
// 	pdf.SetFont("Arial", "", 15)

// 	// Create new PDF generator
// 	pdfg, err := wkhtml.NewPDFGenerator()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Set global options
// 	pdfg.Dpi.Set(300)
// 	// pdfg.Orientation.Set(OrientationLandscape)
// 	pdfg.Grayscale.Set(true)

// 	// Create a new input page from an URL
// 	page := wkhtml.NewPage("input.html")

// 	// Set options for this page
// 	page.FooterRight.Set("[page]")
// 	page.FooterFontSize.Set(10)
// 	page.Zoom.Set(0.95)

// 	// Add to document
// 	pdfg.AddPage(page)

// 	// Create PDF document in internal buffer
// 	err = pdfg.Create()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Write buffer contents to file on disk
// 	err = pdfg.WriteFile("./PdfGenrated.pdf")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("pdf genrated ")

// }
