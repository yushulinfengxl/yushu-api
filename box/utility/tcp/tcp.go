package tcp

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
	"yushu/util"
)

func LocalHostScan() {
	// 65535
	var wg sync.WaitGroup
	for i := 21; i < 65535; i++ {
		wg.Add(1)
		go func(j int) {
			t := fmt.Sprintf("handler.ip-ns.eu.org:%d", j)
			conn, err := net.Dial("tcp", t)
			wg.Done()
			if err == nil {
				fmt.Printf("%v \t已被占用\n", j)
				err := conn.Close()
				if err != nil {
					fmt.Printf("%v \t关闭失败\n", t)
				}
			}
		}(i)
	}
	wg.Wait()
}

func PortScan(host string, startPort int, endPort int) {
	// 65535
	var wg sync.WaitGroup
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(port int) {
			addr := fmt.Sprintf("%v:%d", host, port)
			conn, err := net.DialTimeout("tcp", addr, 30*time.Second)
			wg.Done()
			if err != nil {
				return
			}
			fmt.Printf("端口:%v\t已被占用\n", port)
			err = conn.Close()
			if err != nil {
				fmt.Printf("%v\t关闭失败\n", addr)
			}
		}(port)
	}
	wg.Wait()
	fmt.Println("Scan end...")
}

func Ping(addr string) {
	// 耗时检测time
	durationTime := util.MeasureTime(func() {
		conn, err := net.DialTimeout("tcp", addr, 30*time.Second)
		if err != nil {
			return
		}
		fmt.Printf("端口:%v\t已被占用\n", addr)
		err = conn.Close()
		if err != nil {
			fmt.Printf("%v\t关闭失败\n", addr)
		}
	})
	log.Println(durationTime, "kkk")
}
