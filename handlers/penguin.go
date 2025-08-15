package netc

import "os"
//a function to read the penguin file
func peng()[]byte{
hellomsg ,err:= os.ReadFile("handlers/peng.txt")
if err != nil{
	return []byte("erorr read peng")
}
return hellomsg
}