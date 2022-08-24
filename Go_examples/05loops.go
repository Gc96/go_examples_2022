// array -> colelctuon of values of datatypes

package main

import (
	"fmt"
	"reflect"
)

func main() {

	// declare array
	var scores = [8]int{4, 6, 8, 2, 0, 3, 5, 8}
	fmt.Println(scores)

	var friends = [3]string{"aoob", "klia", "cwg"}
	fmt.Println(friends)

	var specificPosition = [9]int{0: 7, 4: 22, 8: 18}
	fmt.Println(specificPosition)

	fmt.Printf("Length of array specificPosition is %v \n", len(specificPosition))
	fmt.Printf("Datatype specificPosition is %T", specificPosition)
	fmt.Println("Length of array specificPosition is ", reflect.TypeOf(specificPosition).Kind())

}
