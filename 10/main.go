package main

import (
	"fmt"

	filemanager "example.com/calculator/fileManager"
	"example.com/calculator/price"
)

func main() {
	taxes := []float64{0, 0.07, 0.1, 0.15}

	doneChan := make([]chan bool, len(taxes))
	errChan := make([]chan error, len(taxes))

	for index, taxRate := range taxes {
		doneChan[index] = make(chan bool)
		errChan[index] = make(chan error)
		fm := filemanager.New("Price.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdmanager := cmdmanager.New()
		priceJob := price.NewPriceIncludeTaxJob(fm, taxRate)
		go priceJob.Process(doneChan[index], errChan[index])

		// if err != nil {
		// 	fmt.Println("Could not process the job")
		// 	fmt.Println(err)
		// 	break
		// }
	}

	for index := range taxes {
		select {
		case err := <-errChan[index]:
			if err != nil {
				fmt.Println("Could not process the job")
				fmt.Println(err)
			}
		case <-doneChan[index]:
			fmt.Println("Done!")
		}
	}
}
