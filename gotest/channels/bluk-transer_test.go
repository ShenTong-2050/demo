package channels

import "testing"

const (
	max 	= 	50000000	// 数据统计上限
	block	= 	500			// 数据块大小
	bufSize	= 	100			// 缓冲区大小
)

func BenchmarkPer(t *testing.B) {
	done := make(chan struct{})
	c := make(chan int,bufSize)
	go func() {
		count := 0
		for x := range c {
			count += x
		}
		close(done)
	}()
	for i:=0; i<max;i++ {
		c <- i
	}
	close(c)
	<-done
}

func BenchmarkBlock(t *testing.B) {			// 块模式,每次将 500 个数字打包成块传输
	done := make(chan struct{})
	c := make(chan [block]int,bufSize)
	go func() {
		count := 0
		for a := range c {
			for x := range a {
				count += x
			}
		}
		close(done)
	}()
	for i:=0; i<max; i+=block {
		var b [block]int		// 使用数组对数据打包
		for n := 0; n<block; n++ {
			b[n] = i+n
			if i+n == max-1 {
				break
			}
		}
		c <- b
	}
	close(c)
	<-done
}


