package main

import "fmt"

func main() {
	/*
		Convert the given second to 00:00:00 hour minute second format

		Example Input/Output
		int 30 -> "00:00:30" string
		70 -> "00:01:10" // 60
		67812 -> 18:50:12
		678120 -> 188:22:00

	*/
	res := ConvertSecondToTimeString(30)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := ConvertSecondToTimeStringCorrect(arr)
	// fmt.Println(resCorrect)
}

func ConvertSecondToTimeString(second int) string {
	hours := second / 3600
	minutes := second / 60

	timeString := fmt.Sprintf("%d:%d:%d", hours, minutes, second)
	return timeString
}

func ConvertSecondToTimeStringCorrect(second int) string {
	//return "" // TODO: replace this
	hours := second / 3600
	minutes := second / 60 % 60
	second = second % 60

	timeString := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, second)
	return timeString
}
