package main

import (
	"context"
	"fmt"

	"time"
)

func main() {

  secondTest()
}


// ---------------------------------------------------------------------------------------

// agar request 3 soniadan koproq vaqt oladigan bo'lsa, biz kutmay, panic qilishimiz kerak
// berilgan vaqt 5-10 daqiqa
func secondTest() {
ctx, cancel := context.WithTimeout(context.Background(), 2 *time.Second)
defer cancel()

resulst := make(chan int)

go func() {
	resulst <- simulateRequest()
}()
select {
case res := <-resulst:
	fmt.Println("Natija: ", res)
case <-ctx.Done():
	fmt.Println("Timout tugadi , malumot kemadi")
}
}

func simulateRequest() int {
 time.Sleep(time.Second * 1)
 return 1
}