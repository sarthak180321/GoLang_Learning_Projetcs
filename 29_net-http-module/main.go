package main
import(
	"fmt"
	"net/http"
	
)
func handlRoute(writer http.ResponseWriter,request *http.Request){
	// body:=request.Body
	// fmt.Println(body)
	if request.Method != http.MethodGet {
		http.Error(writer,"this Method is not allowed yet",http.StatusNotExtended)
	}
	_,_=writer.Write([]byte("this is reply from Go net-http-module"))
}

func QueryHandler(writer http.ResponseWriter,request *http.Request){
	Query:=request.URL.Query().Get("name")
	fmt.Println(Query)
	if Query==""{
		Query="Guest"
	}
	_,_=writer.Write([]byte(Query))
	
}

func main(){
	go Test()
	fmt.Println("Go net-http-module")
	http.HandleFunc("/hello",handlRoute)
	http.HandleFunc("/handleQuery",QueryHandler)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil {
		fmt.Errorf(err.Error())
	}


}