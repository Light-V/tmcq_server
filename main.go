package main

import (
	"fmt"
	"math/rand"
)

type Data struct {
	value       int
	isProcessed chan bool
}

func main() {

	//c := NewController()
	//c.Run()
	inChan := make(chan *Data, 2)
	outChanNum := 3
	outChans := make([]chan *Data, 0, outChanNum)
	for i := 0; i < outChanNum; i++ {
		outChans = append(outChans, make(chan *Data, 1))
	}

	go func() {
		cnt := 0
		for {
			inData := Data{cnt, make(chan bool)}
			inChan <- &inData
			fmt.Printf("Data %d in\n", inData.value)
			<-inData.isProcessed
			cnt += 1
			if cnt >= 10 {
				close(inChan)
				return
			}
		}
	}()

	for {
		selectedChan := rand.Intn(outChanNum)
		inData, isOn := <-inChan
		if !isOn {
			return
		}
		outChans[selectedChan] <- inData
		outData := <-outChans[selectedChan]
		fmt.Printf("OutChannel %d outputs: %d\n", selectedChan, outData.value)
		close(outData.isProcessed)
	}

}
