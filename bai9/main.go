package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"time"
)

var tpl *template.Template

func main() {
	tpl, _ = template.ParseGlob("templates/*.html")
	http.HandleFunc("/1", ex1)
	http.HandleFunc("/2", ex2)
	http.HandleFunc("/4", ex4)
	http.HandleFunc("/5", ex5)
	http.HandleFunc("/6", ex6)
	http.HandleFunc("/7", ex7)
	http.ListenAndServe(":8080", nil)
}

func ex1(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index1.html", nil)
	fmt.Fprint(w, "<div style='text-align: left;margin: auto;font-size:20px;max-width: 800px;width: 100%;margin: auto;'>")
	fmt.Fprint(w, "<table style='border: 2px solid black;'>")
	fmt.Fprint(w, "<tr>")
	for i := 1; i <= 5; i++ {
		fmt.Fprint(w, "<td style='border:1px solid black;'>")
		for j := 1; j <= 10; j++ {
			fmt.Fprint(w, i, "x", j, "=", j*i, "<br>")
		}
		fmt.Fprint(w, "</td>")
	}
	fmt.Fprint(w, "</tr>")
	fmt.Fprint(w, "<tr>")

	for i := 6; i <= 10; i++ {
		fmt.Fprint(w, "<td style='border:1px solid black;'>")
		for j := 1; j <= 10; j++ {
			fmt.Fprint(w, i, "x", j, "=", j*i, "<br>")
		}
		fmt.Fprint(w, "</td>")
	}
	fmt.Fprint(w, "</tr>")

	fmt.Fprint(w, "</table>")
	fmt.Fprint(w, "</div>")
}

func ex2(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index2.html", nil)

	myArray := [...]string{"Anh", "Phap", "Duc", "VN"}
	myInf := [len(myArray)]int{}
	fmt.Fprintln(w, "<form action=\"2\" method=\"get\">")
	for i := 0; i < len(myArray); i++ {
		fmt.Fprintln(w, myArray[i]+":"+"<input type=\"text\" name=", myArray[i], " placeholder=", myArray[i], " value=\"0\"><br>")
	}
	fmt.Fprint(w, "<input type=\"submit\" value=\"Submit\">")
	fmt.Fprint(w, "</form>")

	for i := 0; i < len(myArray); i++ {
		myInf[i], _ = strconv.Atoi(r.FormValue(myArray[i]))
	}

	fmt.Fprint(w, "<table style=\"margin: auto;\">")
	for i := 0; i < len(myArray); i++ {
		fmt.Fprint(w, "<tr>")
		fmt.Fprint(w, "<td>", myArray[i], "</td>")
		fmt.Fprint(w, "<td style=\"display: flex;gap: 20px;\">")
		fmt.Fprint(w, "<div style=\"display: flex;gap: 10px;\">")
		fmt.Fprint(w, "<p id=", myArray[i], " style=\"background-color: red;padding: 10px ", myInf[i], "px 10px 0;\"></p>")
		fmt.Fprint(w, "<p id=", i, ">", myInf[i], " % </p>")
		fmt.Fprint(w, "</div></td></tr>")
	}
	fmt.Fprint(w, "</table>")
	//fmt.Fprint(w)
	//fmt.Fprint(w)
}

func ex4(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index4.html", nil)

	number, _ := strconv.Atoi(r.FormValue("number"))
	space := "&nbsp;&nbsp;&nbsp;"
	for i := 0; i < number; i++ {
		if i > 9 {
			for j := i; j > 0; j-- {
				if j > 9 {
					x := strconv.Itoa(j)
					fmt.Fprint(w, x[len(x)-1:len(x)], space)
				} else {
					fmt.Fprint(w, j, space)
				}
			}
			fmt.Fprint(w, "<br>")
		} else {
			for j := i; j > 0; j-- {
				fmt.Fprint(w, j, space)
			}
			fmt.Fprint(w, "<br>")
		}
	}
}

func ex5(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index5.html", nil)

	number, _ := strconv.Atoi(r.FormValue("number"))
	space := "&nbsp;&nbsp;"
	for i := 1; i <= number; i++ {
		for j := 0; j <= number-i; j++ {
			fmt.Fprint(w, space)
		}
		if i > 9 {
			for j := i; j > i/2; j-- {
				if j > 9 {
					x := strconv.Itoa(j)
					fmt.Fprint(w, x[len(x)-1:len(x)], space)
				} else {
					fmt.Fprint(w, j, space)
				}
			}
			for j := math.Round(float64(i)/2) + 1; j <= float64(i); j++ {
				if j > 9 {
					x := strconv.Itoa(int(j))
					fmt.Fprint(w, x[len(x)-1:len(x)], space)
				} else {
					fmt.Fprint(w, j, space)
				}
			}
			fmt.Fprint(w, "<br>")
		} else {
			for j := i; j > i/2; j-- {
				fmt.Fprint(w, j, space)
			}
			for j := math.Round(float64(i)/2) + 1; j <= float64(i); j++ {
				fmt.Fprint(w, j, space)
			}
			fmt.Fprint(w, "<br>")
		}
	}

}

func ex6(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index6.html", nil)

	a, _ := strconv.Atoi(r.FormValue("a"))
	b, _ := strconv.Atoi(r.FormValue("b"))
	c := r.FormValue("operator")
	rs := 0
	rs1 := 0.0
	switch c {
	case "+":
		rs = a + b
		break
	case "-":
		rs = a - b
		break
	case "/":
		rs1 = float64(a) / float64(b)
		break
	case "x":
		rs = a * b
		break
	}
	if rs != 0 {
		fmt.Fprint(w, "<p class=\"result\">", rs, "</p>")
	} else if rs1 != 0 {
		fmt.Fprint(w, "<p class=\"result\">", rs1, "</p>")
	}
}

func ex7(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index7.html", nil)

	dayNames := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	date := time.Now()
	//nowDay := 0
	//for i, day := range dayNames {
	//	if day == dayNames[date.Weekday()] {
	//		nowDay = i
	//		break
	//	}
	//}
	day := 0
	fmt.Fprint(w, "<div class=\"calendar\">")
	fmt.Fprint(w, "<p class=\"year\">", date.Month(), "&nbsp;", date.Year(), "</p>")
	fmt.Fprint(w, "<div style=\"display: inline-flex; margin-top: 5px; gap: 20px\"><p>Su</p><p>M</p><p>Tu</p><p>W</p><p>Th</p><p>F</p><p>Sa</p></div>")
	for i := 1; i <= 31; i++ {
		date1 := time.Date(date.Year(), date.Month(), i, 12, 30, 0, 0, time.UTC)
		for j, day1 := range dayNames {
			if day1 == dayNames[date1.Weekday()] {
				day = j
				break
			}
		}
		if i == 1 {
			fmt.Fprint(w, "<div style=\"display: inline-flex; margin-top: 5px; gap: 20px\">")

			for j := 0; j < day; j++ {
				fmt.Fprint(w, "<p style=\"border: 0; padding-right: 3px\"></p>")
			}
		}
		if i != 1 && date1.Day() == 1 {
			fmt.Fprint(w, "</div>")
			break
		} else if date.Day() == date1.Day() {
			fmt.Fprint(w, "<p class=\"today\">", date1.Day(), "</p>")
		} else {
			fmt.Fprint(w, "<p>", date1.Day(), "</p>")
		}
		if day >= 6 {
			fmt.Fprint(w, "</div>")
			fmt.Fprint(w, "<div style=\"display: inline-flex; margin-top: 5px; gap: 20px\">")
		}
	}
	fmt.Fprint(w, "</div>")
}
