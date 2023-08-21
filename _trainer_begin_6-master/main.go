package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var tpl *template.Template

type NewPerson struct {
	Image      string
	FullName   string
	UserName   string
	Email      string
	Phone      string
	References string
	Date       string
	Dob        string
	Gender     string
	Desc       string
	Password   string
	CountRefer int
	Role       string
	ListRefer  []Refer
}

type Refer struct {
	NameRefer string
	Date      string
}

type Response struct {
	Error   string
	Success string
}

var db *sql.DB

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	var err error
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	db, err = sql.Open("mysql", "root:1234567890@tcp(localhost:3306)/test1")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/", home)
	http.HandleFunc("/detail", detail)
	http.HandleFunc("/infor", information)
	log.Fatal(http.ListenAndServe("localhost:9091", nil))
}

func register(w http.ResponseWriter, r *http.Request) {
	_, error := r.Cookie("acc")
	if error == nil {
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
		return
	}
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "register.html", nil)
		return
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02")
	fullName := r.FormValue("fullName")
	userName := r.FormValue("userName")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	references := r.FormValue("references")
	password := r.FormValue("password")
	rePassword := r.FormValue("rePassword")
	checkImg := r.FormValue("checkImg")
	var re Response
	r.ParseMultipartForm(10)
	if !confirmPassword(password, rePassword) {
		re.Error = "Password and RePassword not available!!!"
		tpl.ExecuteTemplate(w, "register.html", re)
		return
	}
	if checkUserName(userName) {
		re.Error = "UserName already exist!!!"
		tpl.ExecuteTemplate(w, "register.html", re)
		return
	}
	if checkPhone(phone) {
		re.Error = "Phone already exist!!!"
		tpl.ExecuteTemplate(w, "register.html", re)
		return
	}
	if checkEmail(email) {
		re.Error = "Email already exist!!!"
		tpl.ExecuteTemplate(w, "register.html", re)
		return
	}
	if references != "" {
		if !checkUserName(references) && !checkEmail(references) {
			re.Error = "References not exist!!!"
			tpl.ExecuteTemplate(w, "register.html", re)
			return
		} else {
			refer := getUserNameByEmail(references)
			if refer != "" {
				references = refer
			}
			insert := "INSERT INTO `test1`.`refer` (`UserName`, `NameRefer`,`Date`) VALUES (?,?,?);"
			row, _ := db.Query(insert, userName, references, formattedTime)
			defer row.Close()
		}
	}

	img := "../static/img/img.png"
	if checkImg != "" {
		file, fileHeader, _ := r.FormFile("myFile")
		defer file.Close()
		contentType := fileHeader.Header["Content-Type"][0]
		fmt.Println("Content Type:", contentType)
		var osFile *os.File //newFile
		defer osFile.Close()
		if contentType == "image/jpeg" {
			osFile, _ = ioutil.TempFile("static/img/", "*.jpg")
		} else {
			re.Error = "The file is not an image!!!"
			tpl.ExecuteTemplate(w, "register.html", re)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		osFile.Write(fileBytes)
		s := osFile.Name()
		fmt.Println(osFile.Name())
		img = "../" + s
	}

	insert := "INSERT INTO `test1`.`information` (`UserName`, `FullName`, `Email`, `Phone`, `Picture`,`Date`, `References`, `DoB`, `Gender`, `Description` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "
	row, _ := db.Query(insert, userName, fullName, email, phone, img, formattedTime, references, "", "", "")
	insert = "INSERT INTO `test1`.`account` (`Email`, `UserName`, `Password`, `role`) VALUES (?,?,?,?);"
	row, _ = db.Query(insert, email, userName, password, 3)
	defer row.Close()
	re.Success = "Register success!"
	tpl.ExecuteTemplate(w, "register.html", re)
}

var count int

func login(w http.ResponseWriter, r *http.Request) {
	cookie, error := r.Cookie("acc")
	if error == nil {
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
		return
	}
	_, error = r.Cookie("ban")
	if error == nil {
		tpl.ExecuteTemplate(w, "login.html", "Ban 2h oke")
		return
	}
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "login.html", nil)
		return
	}
	userName := r.FormValue("userName")
	password := r.FormValue("password")
	var err string
	if !checkUserName(userName) {
		if getUserNameByEmail(userName) == "" {
			err = "UserName or Email not exist!"
			tpl.ExecuteTemplate(w, "login.html", err)
			return
		}
	} else if !checkPassword(userName, password) {
		if count >= 2 {
			cookie := &http.Cookie{
				Name:    "ban",
				Value:   userName,
				Expires: time.Now().Add(2 * time.Hour),
			}
			http.SetCookie(w, cookie)
			count = 0
		}
		count++
		err = "Password not correct!"
		tpl.ExecuteTemplate(w, "login.html", err)
		return
	}

	if error != nil {
		fmt.Println("cookie was not found")
		cookie = &http.Cookie{
			Name:    "acc",
			Value:   userName + "/" + findRole(userName),
			Expires: time.Now().Add(24 * time.Hour),
		}
		http.SetCookie(w, cookie)
	}
	http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
}

var dy = 10.0
var tb = 0.0

func home(w http.ResponseWriter, r *http.Request) {
	dis := r.FormValue("display")
	ta := r.FormValue("tabs")
	var display, tab int
	if dis != "" {
		display, _ = strconv.Atoi(dis)
		dy = float64(display)
	}
	if ta != "" {
		tab, _ = strconv.Atoi(ta)
		tb = float64(tab) - 1
	}
	fmt.Println(dy, tb)
	rows, _ := db.Query("SELECT Picture, FullName, `References`, Date, UserName FROM test1.information LIMIT ? OFFSET ?", int(dy), int(tb*dy))
	defer rows.Close()
	var list []NewPerson
	for rows.Next() {
		var p NewPerson
		rows.Scan(&p.Image, &p.FullName, &p.References, &p.Date, &p.UserName)
		list = append(list, p)
	}
	rows, _ = db.Query("SELECT count(*) FROM test1.information")
	var s string
	for rows.Next() {
		rows.Scan(&s)
	}
	n, _ := strconv.Atoi(s)
	nu := strconv.FormatFloat(float64(n)/dy, 'f', 2, 64)
	cookie := &http.Cookie{
		Name:    "li",
		Value:   nu,
		Expires: time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, cookie)

	tpl.ExecuteTemplate(w, "home.html", list)
}

var p1 NewPerson

type Te struct {
	NewPerson
	Response
}

func information(w http.ResponseWriter, r *http.Request) {
	cookie, error := r.Cookie("acc")
	if error != nil {
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
		return
	}

	if r.Method == "GET" {
		s := strings.Split(cookie.Value, "/")
		fmt.Println(s[0])
		rows, _ := db.Query("SELECT `FullName`, `UserName`, `Email`, `Phone`, `Gender`, `DoB`, `Description`, `Picture`, `Date` FROM test1.information where UserName = ?", s[0])
		defer rows.Close()
		var p NewPerson
		for rows.Next() {
			err := rows.Scan(&p.FullName, &p.UserName, &p.Email, &p.Phone, &p.Gender, &p.Dob, &p.Desc, &p.Image, &p.Date)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}
		rows, _ = db.Query("SELECT Password FROM test1.account where UserName = ?", s[0])
		for rows.Next() {
			rows.Scan(&p.Password)
		}
		rows, _ = db.Query("SELECT UserName, Date FROM test1.refer where NameRefer = ?", s[0])
		var re Refer
		for rows.Next() {
			rows.Scan(&re.NameRefer, &re.Date)
			p.ListRefer = append(p.ListRefer, re)
			p.CountRefer++
		}
		p1 = p
		tpl.ExecuteTemplate(w, "information.html", p)
		return
	}
	fullName := r.FormValue("fullName")
	userName := r.FormValue("userName")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	gender := r.FormValue("gender")
	dob := r.FormValue("dob")
	description := r.FormValue("des")
	password := r.FormValue("password")
	rePassword := r.FormValue("rePassword")
	var re Response
	if !confirmPassword(password, rePassword) {
		re.Error = "Password and RePassword not available!!!"
		tpl.ExecuteTemplate(w, "information.html", p1)
		return
	}
	if checkUserName(userName) {
		if p1.UserName != userName {
			re.Error = "UserName already exist!!!"
			tpl.ExecuteTemplate(w, "information.html", p1)
			return
		}
	}
	if checkPhone(phone) {
		if p1.Phone != phone {
			re.Error = "Phone already exist!!!"
			tpl.ExecuteTemplate(w, "information.html", p1)
			return
		}
	}
	if checkEmail(email) {
		if p1.Email != email {
			re.Error = "Email already exist!!!"
			tpl.ExecuteTemplate(w, "information.html", p1)
			return
		}
	}
	checkImg := r.FormValue("checkImg")
	img := p1.Image
	if checkImg != p1.Image {
		file, fileHeader, _ := r.FormFile("myFile")
		defer file.Close()
		contentType := fileHeader.Header["Content-Type"][0]
		fmt.Println("Content Type:", contentType)
		var osFile *os.File //newFile
		defer osFile.Close()
		if contentType == "image/jpeg" {
			osFile, _ = ioutil.TempFile("static/img/", "*.jpg")
		} else {
			re.Error = "The file is not an image!!!"
			tpl.ExecuteTemplate(w, "information.html", p1)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		osFile.Write(fileBytes)
		s1 := osFile.Name()
		fmt.Println(osFile.Name())
		img = "../" + s1
	}
	insert := "UPDATE `test1`.`information` SET `Picture` = ?,  `FullName` = ?, `UserName` = ?, `Email` = ?,`Phone` = ?,`Gender` = ?,`DoB` = ?,`Description` = ? WHERE `UserName` = ?; "
	row, err1 := db.Query(insert, img, fullName, userName, email, phone, gender, dob, description, userName)
	if err1 != nil {
		fmt.Println("errr1")
		tpl.ExecuteTemplate(w, "information.html", p1)
		return
	}
	insert = "UPDATE `test1`.`account` SET `Email` = ?,`UserName` = ?,`Password`=? WHERE (`Email` = ?);"
	row, err1 = db.Query(insert, email, userName, password, email)
	if err1 != nil {
		fmt.Println("errr2")
		tpl.ExecuteTemplate(w, "information.html", p1)
		return
	}
	changeRefer(userName)
	defer row.Close()
	re.Success = "Update success!"

	s := strings.Split(cookie.Value, "/")
	fmt.Println(s[0])
	rows, _ := db.Query("SELECT `FullName`, `UserName`, `Email`, `Phone`, `Gender`, `DoB`, `Description`, `Picture`, `Date` FROM test1.information where UserName = ?", s[0])
	defer rows.Close()
	var p NewPerson
	for rows.Next() {
		err := rows.Scan(&p.FullName, &p.UserName, &p.Email, &p.Phone, &p.Gender, &p.Dob, &p.Desc, &p.Image, &p.Date)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	rows, _ = db.Query("SELECT Password FROM test1.account where UserName = ?", s[0])
	for rows.Next() {
		rows.Scan(&p.Password)
	}
	tpl.ExecuteTemplate(w, "information.html", p)
}

func detail(w http.ResponseWriter, r *http.Request) {
	cookie, error := r.Cookie("acc")
	userName := r.FormValue("u")
	fmt.Println(userName)
	if error != nil {
		rows, err := db.Query("SELECT `FullName`, `Picture` FROM test1.information where UserName = ?", userName)
		if err != nil {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		}
		defer rows.Close()
		var p NewPerson
		for rows.Next() {
			err := rows.Scan(&p.FullName, &p.Image)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}
		tpl.ExecuteTemplate(w, "detail.html", p)
		return
	}
	s := strings.Split(cookie.Value, "/")
	if r.Method == "GET" || !checkRole(userName, s[0]) {
		rows, _ := db.Query("SELECT `FullName`, `UserName`, `Email`, `Phone`, `Gender`, `DoB`, `Description`, `Picture`, `Date` FROM test1.information where UserName = ?", userName)
		defer rows.Close()
		var p NewPerson
		for rows.Next() {
			err := rows.Scan(&p.FullName, &p.UserName, &p.Email, &p.Phone, &p.Gender, &p.Dob, &p.Desc, &p.Image, &p.Date)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}
		if s[1] == "1" {
			p.Role = findRole(userName)
		}
		p1 = p
		tpl.ExecuteTemplate(w, "detail.html", p)
		return
	}
	userName = r.FormValue("userName")
	fullName := r.FormValue("fullName")
	phone := r.FormValue("phone")
	gender := r.FormValue("gender")
	dob := r.FormValue("dob")
	des := r.FormValue("des")
	role := r.FormValue("role")
	var re Response
	if checkPhone(phone) {
		if p1.Phone != phone {
			re.Error = "Phone already exist!!!"
			tpl.ExecuteTemplate(w, "detail.html", p1)
			return
		}
	}
	checkImg := r.FormValue("checkImg")
	img := p1.Image
	if checkImg != p1.Image {
		file, fileHeader, _ := r.FormFile("myFile")
		defer file.Close()
		contentType := fileHeader.Header["Content-Type"][0]
		fmt.Println("Content Type:", contentType)
		var osFile *os.File //newFile
		defer osFile.Close()
		if contentType == "image/jpeg" {
			osFile, _ = ioutil.TempFile("static/img/", "*.jpg")
		} else {
			re.Error = "The file is not an image!!!"
			tpl.ExecuteTemplate(w, "detail.html", p1)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		osFile.Write(fileBytes)
		s1 := osFile.Name()
		fmt.Println(osFile.Name())
		img = "../" + s1
	}
	insert := "UPDATE `test1`.`information` SET `Picture` = ?,  `FullName` = ?, `Phone` = ?,`Gender` = ?,`DoB` = ?,`Description` = ? WHERE `UserName` = ?; "
	row, err1 := db.Query(insert, img, fullName, phone, gender, dob, des, userName)
	insert = "UPDATE `test1`.`account` SET `role` = ? WHERE `UserName` = ? "
	row, err1 = db.Query(insert, role, userName)
	defer row.Close()
	if err1 != nil {
		fmt.Println("errr1")
		tpl.ExecuteTemplate(w, "detail.html", p1)
		return
	}

	rows, _ := db.Query("SELECT `FullName`, `UserName`, `Email`, `Phone`, `Gender`, `DoB`, `Description`, `Picture`, `Date` FROM test1.information where UserName = ?", userName)
	defer rows.Close()
	var p NewPerson
	for rows.Next() {
		err := rows.Scan(&p.FullName, &p.UserName, &p.Email, &p.Phone, &p.Gender, &p.Dob, &p.Desc, &p.Image, &p.Date)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	if s[1] == "1" {
		p.Role = findRole(userName)
	}
	fmt.Println("done")
	tpl.ExecuteTemplate(w, "detail.html", p)
}

func checkRole(userName, userLogin string) bool {
	rows, _ := db.Query("SELECT role FROM test1.account where UserName = ?", userName)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&userName)
	}
	rows, _ = db.Query("SELECT role FROM test1.account where UserName = ?", userLogin)
	for rows.Next() {
		rows.Scan(&userLogin)
	}
	un, _ := strconv.Atoi(userName)
	ul, _ := strconv.Atoi(userLogin)
	if ul < un {
		return true
	}
	return false
}

func changeRefer(references string) {
	insert := "SELECT UserName FROM test1.refer where NameRefer = ?"
	rows, _ := db.Query(insert, references)
	defer rows.Close()
	var list []string
	var s string
	for rows.Next() {
		rows.Scan(&s)
		list = append(list, s)
	}
	for _, i := range list {
		insert = "UPDATE `test1`.`refer` SET `NameRefer` = ? WHERE (`UserName` = ?);"
		rows, _ = db.Query(insert, references, i)
	}
}

func confirmPassword(password, rePassword string) bool {
	if password == rePassword {
		return true
	}
	return false
}

func checkUserName(userName string) bool {
	rows, _ := db.Query("SELECT UserName FROM test1.information")
	defer rows.Close()
	var listString []string
	var s string
	for rows.Next() {
		rows.Scan(&s)
		listString = append(listString, s)
	}
	for i := 0; i < len(listString); i++ {
		if listString[i] == userName {
			return true
		}
	}
	return false
}

func checkEmail(email string) bool {
	rows, _ := db.Query("SELECT Email FROM test1.information")
	defer rows.Close()
	var listString []string
	var s string
	for rows.Next() {
		rows.Scan(&s)
		listString = append(listString, s)
	}
	for i := 0; i < len(listString); i++ {
		if listString[i] == email {
			return true
		}
	}
	return false
}

func getUserNameByEmail(email string) string {
	rows, _ := db.Query("SELECT Email FROM test1.information")
	defer rows.Close()
	var listString []string
	var s string
	for rows.Next() {
		rows.Scan(&s)
		listString = append(listString, s)
	}
	for i := 0; i < len(listString); i++ {
		if listString[i] == email {
			row, _ := db.Query("SELECT UserName FROM test1.information where Email = ?", listString[i])
			defer row.Close()
			var s1 string
			for row.Next() {
				row.Scan(&s1)
				return s1
			}
		}
	}
	return ""
}

func checkPhone(phone string) bool {
	rows, _ := db.Query("SELECT Email FROM test1.information")
	defer rows.Close()
	var listString []string
	var s string
	for rows.Next() {
		rows.Scan(&s)
		listString = append(listString, s)
	}
	for i := 0; i < len(listString); i++ {
		if listString[i] == phone {
			return true
		}
	}
	return false
}

func checkPassword(userName, password string) bool {
	rows, _ := db.Query("SELECT Password FROM test1.account where UserName = ?;", userName)
	defer rows.Close()
	var s string
	for rows.Next() {
		rows.Scan(&s)
	}
	if password == s {
		return true
	}
	return false
}

func findRole(userName string) string {
	rows, _ := db.Query("SELECT role FROM test1.account where UserName = ?;", userName)
	defer rows.Close()
	var s string
	for rows.Next() {
		rows.Scan(&s)
	}
	return s
}
