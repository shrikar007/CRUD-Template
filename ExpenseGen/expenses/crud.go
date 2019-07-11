package expenses

import(

"database/sql"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/mysql"
"net/http"
"context"
"fmt"
"errors"
"log"
"strconv"
"github.com/go-chi/chi"
"github.com/go-chi/chi/middleware"
"github.com/go-chi/render"

_ "github.com/go-sql-driver/mysql"
"time"
"os"
)
var obj Expense
var expenses Expenses
var db1 *reform.DB
var err error

var req Createreq

func Init(){
     d,err:=sql.Open("mysql", "root:root@tcp(localhost:3306)/expense?charset=utf8&parseTime=True")
                 if err != nil {

                     fmt.Println(err)
                 }
                 lg := log.New(os.Stderr, "SQL: ", log.Flags())
                 db1= reform.NewDB(d, mysql.Dialect, reform.NewPrintfLogger(lg.Printf))

       r := chi.NewRouter()
           r.Use(middleware.RequestID)
           r.Use(middleware.RealIP)
           r.Use(middleware.Logger)
           r.Use(middleware.Recoverer)
           r.Use(render.SetContentType(render.ContentTypeJSON))
           r.Route("/expenses", func(r chi.Router) {
               r.Post("/", Create)
               r.Get("/", GetAll)
               r.Route("/{id}", func(r chi.Router) {
                   r.Use(CrudContext)
                   r.Get("/",GetId)
                   r.Put("/", Update)
                   r.Delete("/", Delete)
               })
           })
           log.Fatal(http.ListenAndServe(":8088", r))
   }

func CrudContext(next http.Handler) http.Handler {
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {


ID := chi.URLParam(r, "id")
id,_:=strconv.Atoi(ID)
        temp,err :=db1.FindByPrimaryKeyFrom(ExpenseTable,id)

        if err != nil {
                fmt.Println(err)
                return
            }else{
            ctx := context.WithValue(r.Context(), "key", temp)
            next.ServeHTTP(w, r.WithContext(ctx))
        }
	})

}
func Create(writer http.ResponseWriter , request *http.Request){
err = render.Bind(request, &req)
	temp:=*req.Expense

	t1:=&Expense{
    		Description:temp.Description,
    			Type :temp.Type,
    		  	Amount  :temp.Amount,
    		  	CreatedOn :  time.Now(),
    		  	UpdatedOn  : time.Now(),

    	}

    	db1.Save(t1)
	render.Render(writer, request,List1(req.Expense))
}
func Update(writer http.ResponseWriter , request *http.Request){
s:=request.Context().Value("key").(*Expense)
var upreq Updatereq
err:= render.Bind(request,&upreq)
   if err != nil {
       log.Println(err)
       return
   }
    var temp Expense
       temp=*upreq.Expense

            s.Description=temp.Description
                   s.Type=temp.Type
                   s.Amount=temp.Amount
                   s.UpdatedOn=time.Now()

               err1 := db1.Update(s)

       if err1 != nil{

                  err=errors.New("Expense not found")
                  fmt.Println(err)
                  return
              }else{
                  err=render.Render(writer, request, List1(&temp))
                  fmt.Println(err)
                   }
 }
func Delete(writer http.ResponseWriter , request *http.Request){
s:=request.Context().Value("key").(*Expense)
	err=db1.Delete(s)
	if err != nil {
		panic(err)
	}else{
		_=render.Render(writer, request, List1(s))
	}
}
func GetAll(writer http.ResponseWriter , request *http.Request){
     flag:=1
     var tem Expenses
	tables, err := db1.SelectRows(ExpenseTable, "WHERE id IS NOT NULL")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tables.Close()

	for flag!=0{
		err = db1.NextRow(&obj, tables)
		tem=append(tem,obj)
		//fmt.Println(obj)
		if err!=nil{
			flag=0
			break
		}
	}
	_=render.Render(writer, request, ListAll(&tem))


}
func GetId(writer http.ResponseWriter , request *http.Request){

    s:=request.Context().Value("key").(*Expense)
          _=render.Render(writer, request, List1(s))

}