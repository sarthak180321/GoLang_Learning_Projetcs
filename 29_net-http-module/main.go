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

func main(){
	fmt.Println("Go net-http-module")
	http.HandleFunc("/hello",handlRoute)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil {
		fmt.Errorf(err.Error())
	}
}