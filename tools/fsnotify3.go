package main
 
import (
    "github.com/fsnotify/fsnotify"
    "log"
    "fmt"
    "os/exec"
    "regexp"
    "strconv"
    "bytes"
    "errors"
    "os"
    "path/filepath"
)
 
const (
    confFilePath = "./conf"
)
 
//获取进程ID
func getPid(processName string) (int, error) {
    //通过wmic process get name,processid | findstr server.exe获取进程ID
    buf := bytes.Buffer{}
    cmd := exec.Command("wmic", "process", "get", "name,processid")
    cmd.Stdout = &buf
    cmd.Run()
    cmd2 := exec.Command("findstr", processName)
    cmd2.Stdin = &buf
    data, _ := cmd2.CombinedOutput()
    if len(data) == 0 {
        return -1, errors.New("not find");
    }
    info := string(data)
    //这里通过正则把进程id提取出来
    reg := regexp.MustCompile(`[0-9]+`)
    pid := reg.FindString(info)
    return strconv.Atoi(pid)
}
 
//启动进程
func startProcess(exePath string, args []string) error {
    attr := &os.ProcAttr{
        //files指定新进程继承的活动文件对象
        //前三个分别为，标准输入、标准输出、标准错误输出
        Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
        //新进程的环境变量
        Env: os.Environ(),
    }
 
    p, err := os.StartProcess(exePath, args, attr);
    if err != nil {
        return err;
    }
    fmt.Println(exePath, "进程启动");
    p.Wait();
    return nil;
}
 
func main() {
    //创建一个监控对象
    watch, err := fsnotify.NewWatcher();
    if err != nil {
        log.Fatal(err);
    }
    defer watch.Close();
    //添加要监控的文件
    err = watch.Add(confFilePath);
    if err != nil {
        log.Fatal(err);
    }
    //我们另启一个goroutine来处理监控对象的事件
    go func() {
        for {
            select {
            case ev := <-watch.Events:
                {
                    //我们只需关心文件的修改
                    if ev.Op&fsnotify.Write == fsnotify.Write {
                        fmt.Println(ev.Name, "文件写入");
                        //查找进程
                        pid, err := getPid("server.exe");
                        //获取运行文件的绝对路径
                        exePath, _ := filepath.Abs("./server.exe")
                        if err != nil {
                            //启动进程
                            go startProcess(exePath, []string{});
                        } else {
                            //找到进程，并退出
                            process, err := os.FindProcess(pid);
                            if err == nil {
                                //让进程退出
                                process.Kill();
                                fmt.Println(exePath, "进程退出");
                            }
                            //启动进程
                            go startProcess(exePath, []string{});
                        }
                    }
                }
            case err := <-watch.Errors:
                {
                    fmt.Println("error : ", err);
                    return;
                }
            }
        }
    }();
 
    //循环
    select {};
}
