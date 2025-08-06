package netc

import "os"


func peng()[]byte{
hellomsg ,err:= os.ReadFile("handlers/peng.txt")
if err != nil{
	return []byte("erorr read peng")
}
return hellomsg
}