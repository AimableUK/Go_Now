package main

func mainH() {

	// ============== Example 1: ================
	// welcome := "Welcome to Our Pizza App"
	// fmt.Println(welcome)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter the rating for our Pizza:")

	// // comma Ok || err ok
	// input, _ := reader.ReadString('\n')

	// fmt.Println("Thanks for rating,", input)

	// numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Added 1 to your rating: ", numRating+1)
	// }

	// ============== Exercise 1 ==============
	// fmt.Println("Welcome to My GOing App")

	// for {
	// 	fmt.Println("Enter the rating, from 1 to 5:")

	// 	reader := bufio.NewReader(os.Stdin)
	// 	inputB, _ := reader.ReadString('\n')
	// 	ratingB, err := strconv.ParseFloat(strings.TrimSpace(inputB), 64)

	// 	if err != nil {
	// 		fmt.Println("You entered invalid rating")
	// 		continue
	// 	}

	// 	fmt.Printf("Thank you! You entered: %.1f\n", ratingB)
	// 	break
	// }

	// ================= Exercise 2 ================
	// fmt.Println("Welcome to My GOing App")
	// for {
	// 	fmt.Println("Enter the rating, in range 1 to 5:")

	// 	reader := bufio.NewReader(os.Stdin)
	// 	inputC, _ := reader.ReadString('\n')
	// 	ratingC, err := strconv.ParseFloat(strings.TrimSpace(inputC), 64)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	} else {
	// 		if ratingC <= 1 || ratingC >= 5 {
	// 			fmt.Println("Please enter a rating between 1 and 5")
	// 			continue
	// 		} else {
	// 			fmt.Println("Good")
	// 			break
	// 		}
	// 	}
	// }

}
