package main

import "fmt"
import "log"
import "bytes"


func main() {

    var (
        buf bytes.Buffer
        logger = log.New(&buf, "INFO: ", log.Lshortfile)
    )
        logger.SetPrefix("WARNING :")
        infof := func(info string){
            logger.Output(2, info)
       }
       infof("Hello, World!")
       fmt.Print(&buf)

}
