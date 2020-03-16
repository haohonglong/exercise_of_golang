package helper
import (
	"runtime"
    "fmt"
    "time"
)

func Log() string {
    _, fileName, fileLine, ok := runtime.Caller(1)
    timelocal, err  := time.LoadLocation("Asia/Chongqing")
    if err != nil {
        panic(err)
    }
    time.Local = timelocal
    curNow := time.Now().Local()
    var s string
    if ok {
        s = fmt.Sprintf("%s %s:%d", curNow.Format("2006-01-02 15:04:05"), fileName, fileLine)
    } else {
        s = ""
    }
    return s
}