package main
import(
"html/template"
"log"
"net/http"
"time"
)
type PageVariable struct {
Date string
Time string
}
func main(){
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
func HomePage(w http.ResponseWriter,r *http.Request){
now:=time.Now()
HomePageVars :=PageVariable{
Date: now.Format("10-01-19"),
Time: now.Format("15:04:05"),
}
t, err:=template.ParseFiles("homepage.html");
if err!=nil{
log.Print("template parsing error: ",err)
}
err =t.Execute(w,HomePageVars)
if err!=nil{
log.Print("template executing error:",err)
}
}