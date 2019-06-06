package main

import (
    "sync"
	"poollisting/pool"
	"poollisting/db"
    "log"
    "time"
    "math/rand"
)

const (
    maxGoroutines   = 5 // 要使用的goroutine的数量
    pooledResources = 2 // 池中的资源的数量
)

func performQuery(query int, p *pool.Pool) {
    // 1. 获取连接
    conn, err := p.Acquire()
    if err != nil {
        log.Println("acquire conn error, ", err)
        return
    }

    // 使用结束后，释放链接
    defer p.Release(conn)

    // 该 log 模拟对连接的使用
    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    log.Printf("QID[%d] CID[%d]\n", query, conn.(*db.DBConnection).ID)
}

func main() {
    var wg sync.WaitGroup
    wg.Add(maxGoroutines)
    // 1. 创建一个 Pool
    p, err := pool.New(db.CreateConn, pooledResources)
    if err != nil {
        log.Println("create Pool error")
    }

    // 2. 开启 goroutine 执行任务
    for query := 0; query < maxGoroutines; query++ {
        // 每个goroutine需要自己复制一份要、查询值的副本，
        // 不然所有的查询会共享同一个查询变量，即所有的 goroutine 最后的 query 值都是3
        go func(q int) {
            performQuery(q, p)
            wg.Done()
        }(query)
        //time.Sleep(1000*time.Millisecond) // 用于测试从 resources channel 中获取资源
    }

    // 3. 关闭连接池
    wg.Wait()
    p.Close()
    log.Println("pool closed - main")
}
