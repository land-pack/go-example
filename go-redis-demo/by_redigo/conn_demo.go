package main

import "github.com/garyburd/redigo/redis"

func main(){
    n, err := conn.Do("APPEND", "key", "value")
    println(n, err)
}
