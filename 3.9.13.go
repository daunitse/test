package main

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)

		for {
			select {
			case num := <-firstChan:
				resultChan <- num * num
			case num := <-secondChan:
				resultChan <- num * 3
			case <-stopChan:
				return
			}
		}
	}()

	return resultChan

}
