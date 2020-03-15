package main
  
  import (
    "errors"
    "github.com/CodyGuo/win"
    "net/http"
    "time"
  )
  
  func main() {
    url := "http://www.baidu.com"
  
    // 判断链接是否可用，如果不可用则结束执行
    err := Get(url)
    if err != nil {
      return
    }
    // 使用cmd方式调用谷歌浏览器并通过app模式打开指定网址
    execRun("cmd /c start chrome --app=" + url)
  }
  
  // 判断链接是否可用
  func Get(url string) (err error) {
    client := http.Client{Timeout: 5 * time.Second}
    _, error := client.Get(url)
    return error
  }
  
  var (
    winExecError = map[uint32]string{
      0:  "The system is out of memory or resources.",
      2:  "The .exe file is invalid.",
      3:  "The specified file was not found.",
      11: "The specified path was not found.",
    }
  )
  
  // 执行指定cmd命令
  func execRun(cmd string) error {
    lpCmdLine := win.StringToBytePtr(cmd)
    ret := win.WinExec(lpCmdLine, win.SW_HIDE)
    if ret <= 31 {
      return errors.New(winExecError[ret])
    }
    return nil
  }
