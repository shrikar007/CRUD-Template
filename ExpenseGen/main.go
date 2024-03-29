package main

import (
	"flag"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/gobuffalo/packr/v2"
	"github.com/shrikar007/ExpenseGen/expenses"
	"html/template"
	"log"
	"os"

	"strings"
	//"gopkg.in/reform.v1"
)
type Json struct{

	Struct string
	SliceStruct string
	Package string

}
// go:generate main -package="ExpenseGen" -Opname="Create" -Opname1="Update" -Opname2="Delete" -Opname3="GetbyId" -Opname4="GetAll"
func main() {
	var struc string
	//names := structs.Names(&Json{})
	//fmt.Println(names)
	//var a Json
	//val := reflect.Indirect(reflect.ValueOf(a))
	//fmt.Println(val.Type().Field(0).Type)

	flag.StringVar(&struc, "Structurename", "", "name of the structure")
	flag.Parse()
	box := packr.New("temp", "./templates")
	t, err := box.FindString("request-create.gotpl")
	t1, err := box.FindString("responsegen.gotpl")
	t2, err := box.FindString("crud-op.gotpl")
	if err != nil {
		log.Fatal(err)
	}
	tpl, err := template.New("request").Parse(t)
	if err != nil {
		fmt.Println(err)
	}
	tpl1, err1 := template.New("request").Parse(t1)
	if err1 != nil {
		fmt.Println(err)
	}
	tpl2, err2 := template.New("request").Parse(t2)
	if err2 != nil {
		fmt.Println(err)
	}
	pluralize := pluralize.NewClient()
	StructSlice := pluralize.Plural(struc)
	pkgname := strings.ToLower(StructSlice)
	data := Json{
		Struct:      struc,
		SliceStruct: StructSlice,
		Package:     pkgname,
	}

	file, err := os.Create("./" + pkgname + "/request.go")
	file1, err := os.Create("./" + pkgname + "/response.go")
	file2, err := os.Create("./" + pkgname + "/crud.go")

	err1 = tpl.Execute(file, data)
	err = tpl1.Execute(file1, data)
	err = tpl2.Execute(file2, data)
	if err1 != nil {
			panic(err1)

	}
expenses.Init()

}