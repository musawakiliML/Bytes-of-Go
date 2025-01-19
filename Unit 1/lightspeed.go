package main

import "fmt"

func main() {

	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km

	fmt.Println(lightSpeed/distance, "Seconds")

	distance = 401000000

	fmt.Println(lightSpeed/distance, "Seconds")
}
