package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"strings"
	"regexp"
	"strconv"
	//"golang.org/x/mobile/event/key"
)

func main() {

	construir()
}

func construir() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Calculadora")
	window.Connect("destroy", gtk.MainQuit)
	vbox := gtk.NewVBox(false, 1)

	vpaned := gtk.NewVPaned()
	vbox.Add(vpaned)

	frame1 := gtk.NewFrame("Calculadora")
	framebox1 := gtk.NewVBox(false, 1)
	frame1.Add(framebox1)

	frame2 := gtk.NewFrame("")
	framebox2 := gtk.NewVBox(false, 1)
	frame2.Add(framebox2)

	vpaned.Pack1(frame1, false, false)
	vpaned.Pack2(frame2, false, false)

	entry := gtk.NewEntry()
	entry.SetText("")
	framebox1.Add(entry)

	swin := gtk.NewScrolledWindow(nil, nil)
	tableNumber := gtk.NewTable(3, 4, true)
	for y := uint(0); y < 4; y++ {
		for x := uint(0); x < 3; x++ {
			var format = "%01d"
			k := uint(0)
			 
			if y == 0 {
				k = (x + 1)
			} else if y == 1 {
				k = x + 4
			} else if y == 2 {
				k = x + 7
			} else if y == 3 {
				if x == 0 {
					k = 00
					format = "%02d"
				} else if x == 2 {
					k = 000
					format = "%03d"
				}

			}
		
			btn := gtk.NewButtonWithLabel(fmt.Sprintf(format, k))		
			
			btn.Clicked(func() {
				fmt.Println("expressao =", entry.GetText())
				entry.SetText(entry.GetText() + "" + btn.GetLabel())
			})
			tableNumber.Attach(btn, x, x+1, y, y+1, gtk.FILL, gtk.FILL, 15, 15)
		}
	}

	buttons := gtk.NewHBox(false, 1)
	soma := gtk.NewButtonWithLabel("+")
	//soma.set
	soma.Clicked(func() {
 		entry.SetText(entry.GetText() + " " + soma.GetLabel() + " ")
	})
	sub := gtk.NewButtonWithLabel("-")
	//soma.set
	sub.Clicked(func() {
 		entry.SetText(entry.GetText() + " " + sub.GetLabel() + " ")
	})
	multi := gtk.NewButtonWithLabel("*")
	//soma.set
	multi.Clicked(func() {
 		entry.SetText(entry.GetText() + " " + multi.GetLabel() + " ")
	})
	divs := gtk.NewButtonWithLabel("/")
	//soma.set
	divs.Clicked(func() {
 		entry.SetText(entry.GetText() + " " + divs.GetLabel() + " ")
	})
	result := gtk.NewButtonWithLabel("=")
	//soma.set
	result.Clicked(func() {
		//function cacola
		retorno := calcolo( entry.GetText())
		entry.SetText(retorno)
	})
	backspace := gtk.NewButtonWithLabel("<-")
	//soma.set
	backspace.Clicked(func() {
 		s := "" + entry.GetText()
		lastBin := len(s)
		if lastBin > 0 {
			s = s[0 : lastBin-1]
			fmt.Println(s)
			entry.SetText(s)
		}
	})
	buttons.Add(soma)
	buttons.Add(sub)
	buttons.Add(multi)
	buttons.Add(divs)
	buttons.Add(result)
	buttons.Add(backspace)

	//adiciona a tabela un btn
	//event := make(chan interface{})

	swin.AddWithViewPort(tableNumber)

	framebox2.PackStart(buttons, false, false, 0)
	framebox2.Add(swin)
	vbox.PackStart(framebox1, false, false, 0)
	vbox.PackStart(framebox2, false, false, 0)
	window.Add(vbox)
	window.SetDefaultSize(300, 650)
	window.ShowAll()

	gtk.Main()

}

func calcolo(a string)  string {

	fmt.Println("%s",a)
	fmt.Println(strings.Split(a,"+"))
	var s []string = strings.Split(a," ")
	//var size int = len(s)
	var rr string = ""
	for i,n := range s{
		fmt.Print("FOR: "+n+",")	
		fmt.Println(i)
		if isOperator(n){
			switch n{
				case "+": 
					a,_  := strconv.ParseFloat((s[i-1]), 64)
					b,_  := strconv.ParseFloat((s[i+1]), 64) 
					
					c  := soma(a,b)
					//remover os tres valore uzads e substituir pelo resultado c
 					if len(s) > 2{ 
 						s[i+1] = strconv.FormatFloat(c, 'f', 2, 64) 
					}else{
 						s[i+1] = strconv.FormatFloat(c, 'f', 2, 64)

					} 
					rr = strconv.FormatFloat(c, 'f', 2, 64)
				break;
				case "-":
						a,_  := strconv.ParseFloat((s[i-1]), 64)
					b,_  := strconv.ParseFloat((s[i+1]), 64) 
					
					c  := subtracao(a,b)
					//remover os tres valore uzads e substituir pelo resultado c
 					if len(s) > 2{ 
 						s[i+1] = strconv.FormatFloat(c, 'f', 2, 64) 
					}else{
 						s[i+1] = strconv.FormatFloat(c, 'f', 2, 64)

					} 
					rr = strconv.FormatFloat(c, 'f', 2, 64)
				break;
				case "*":	
					a,_  := strconv.ParseFloat((s[i-1]), 64)
					b,_  := strconv.ParseFloat((s[i+1]), 64) 
					
					c  := multiplicacao(a,b)
					//remover os tres valore uzads e substituir pelo resultado c
 					if len(s) > 2{ 
 						s[i+1] = strconv.FormatFloat(c, 'f', 2, 64) 
					}else{
 						s[i+1] = strconv.FormatFloat(c, 'f', 2, 64)

					} 
					rr = strconv.FormatFloat(c, 'f', 2, 64)

				break;
				case "/":
					a,_  := strconv.ParseFloat((s[i-1]), 64)
					b,_  := strconv.ParseFloat((s[i+1]), 64) 
					
					c  := divisao(a,b)
					//remover os tres valore uzads e substituir pelo resultado c
 					if len(s) > 2{ 
 						s[i+1] = strconv.FormatFloat(c, 'f', 2, 64) 
					}else{
 						s[i+1] = strconv.FormatFloat(c, 'f', 2, 64)

					} 
					rr = strconv.FormatFloat(c, 'f', 2, 64)

				break;
			}
		}	
	}
	//ans := soma(1, 2)
	//fmt.Println("%s",a)
	return rr
}
func soma(a float64, b float64) float64 {
	return (a + b)
}

func subtracao(a float64, b float64) float64 {
	return (a - b)
}
func multiplicacao(a float64, b float64) float64 {
	return a * b
}

func divisao(a float64, b float64) float64 {
	return a / b
}

func isOperator(a string) bool {
	var validID = regexp.MustCompile(`[+-/*]`)
	return validID.MatchString(a)
	
}