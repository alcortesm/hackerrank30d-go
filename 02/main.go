package main

import (
	"fmt"
	"log"
)

func main() {

	// This exercise is completely wrong, you should NEVER use floats to
	// handle money. This is quite serious, especially in introductory
	// material, you should not even suggest coming close to floats when
	// dealing with money.  It is not just that float arithmetic
	// subtleties is hard to a newbie, it is that rounding money is
	// complicated even for experts.
	//
	// This kind of exercises should use "room dimensions" or
	// "temperature" or any other phisical unit instead, but not money.
	//
	// That being said, lets do it using floats.

	price, err := readPrice()
	if err != nil {
		log.Fatal(err)
	}

	tipPercent, err := readPercent()
	if err != nil {
		log.Fatal(err)
	}

	taxPercent, err := readPercent()
	if err != nil {
		log.Fatal(err)
	}

	tip := price * float64(tipPercent) / 100.
	tax := price * float64(taxPercent) / 100.
	totalCost := price + tip + tax

	fmt.Printf("The total meal cost is %.0f dollars.\n", totalCost)
}

func readPrice() (float64, error) {
	var price float64
	if _, err := fmt.Scanf("%f\n", &price); err != nil {
		return 0, fmt.Errorf("cannot read price: %s", err)
	}

	if price < 0 {
		return 0, fmt.Errorf("negative prices are not supported")
	}

	return price, nil
}

func readPercent() (cents int, err error) {
	var percent int
	if _, err := fmt.Scanf("%d\n", &percent); err != nil {
		return 0, fmt.Errorf("cannot read percent: %s", err)
	}

	if percent < 0 {
		return 0, fmt.Errorf("negative percents are not allowed")
	}

	return percent, nil
}
