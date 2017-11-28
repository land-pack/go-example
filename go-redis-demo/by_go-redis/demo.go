package main

import (
    "github.com/go-redis/redis"
)

func main() {
    ring := redis.NewRing(&redis.RingOptions{
        Addrs: map[string]string{
            "shard1": ":6379",
        },
    })
    ring.ForEachShard(func(client *redis.Client) error {
        println(client)
        return nil
    })

    /*
    for {
        ring.Ping()
    }
    */
}
