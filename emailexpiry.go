package emailexpiry

import (
	"fmt"
	"time"
)

func ExpireTime(postdate string)string{
loc,_:=time.LoadLocation("UTC")
dt:=time.Now().In(loc)
ct:=dt.Format("2006-01-02 3:04:05")
x,_:=time.Parse("2006-01-02 3:04:05",postdate)
y,_:=time.Parse("2006-01-02 3:04:05",ct)
diff:=y.Sub(x)
timediff:=int(diff.Minutes())
if timediff<=60{
	fmt.Println(timediff)
	return "verified"
}else{
	fmt.Println(timediff)
	return "expired"
}
}
