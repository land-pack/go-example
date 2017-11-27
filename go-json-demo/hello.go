package main

import "fmt"
import "github.com/Jeffail/gabs"



func main(){

    jsonParsed, err := gabs.ParseJSON([]byte(`{
        "debug":1,
    	"outter":{
    		"inner":{
    			"value1":10,
    			"value2":22
    		},
    		"alsoInner":{
    			"value1":20
    		}
    	}
    }`))




    var value float64
    var ok bool

    value, ok = jsonParsed.Path("outter.inner.value1").Data().(float64)
    // value == 10.0, ok == true
    if ok == true{
        fmt.Printf("\n")
        println("value is ==>", value)
    }else{
        println("error -->", err)
        }

    if jsonParsed.Exists("debug-again"){
        println("the data is complete")
    }
}
