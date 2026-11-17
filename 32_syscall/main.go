package main 
import(
	"fmt"
	"syscall"
)
type FileData struct{
	name string
	fd int
	content []byte
	contentsize int
}
func main(){
	// pid := syscall.Getpid()
	// uid := syscall.Getuid()
	// fmt.Println(pid)
	// fmt.Println(uid)
	// fmt.Println("Going OS level with syscall")
	// fd,err:=syscall.Open("test.txt",syscall.O_RDONLY,0)
	// if err != nil {
	// 	fmt.Println("Error opening file:",err)
	// 	panic(err)
	// }
	// defer syscall.Close(fd)
	// buf:=make([]byte,100)
	// n,_:=syscall.Read(fd,buf)
	// syscall.Write(1,buf[:n])


	// fd2,_:=syscall.Open("output.txt",syscall.O_CREAT|syscall.O_WRONLY|syscall.O_TRUNC,0644)
	// defer syscall.Close(fd2)
	// syscall.Write(fd2,buf[:n])
	 
	// _,err2:=syscall.Open("nonexistent.txt",syscall.O_RDONLY,0)
	// if err2 != nil {
	// 	fmt.Println("Error opening file:",err2)
	// 	panic(err2)
	// }
	// for i:=0; ; i++ {
	// 	_,err:=syscall.Open("test.txt",syscall.O_RDONLY,0)
	// 	if err != nil {
	// 		fmt.Println("Error opening file:",err)
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("execs se pehle")
	// err:=syscall.Exec("/bin/ls",[]string{"ls","-l"},syscall.Environ())
	// if err != nil {
	// 	fmt.Println("Error executing command:",err)
	// 	panic(err)
	// }
	mypid:=syscall.Getpid()
	pid,err:=syscall.ForkExec("/bin/ls",[]string{"ls","-l"},&syscall.ProcAttr{
		Files: []uintptr{0,1,2},
		Env: syscall.Environ(),
	})
	if err != nil {
		fmt.Println("Error executing command:",err)
		panic(err)
	}
	fmt.Println(pid,mypid)

}