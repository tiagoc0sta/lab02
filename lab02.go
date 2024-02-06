package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)
func main() {
    menu := []string{ //slice of strings
        "1. Analysis & design",
        "2. Coding & evidence",
        "3. Deployment & maintenance",
        "4. Exit",
    }
    activities := [][]string{ //2D slice of strings
        {
            "Define the Problem",
            "Plan the Solution",
            "Choose a Programming Language",
        },
        {
            "Set Up Development Environment",
            "Write Code",
            "Test Your Code",
        },
        {
            "Debugging",
            "Run and Test Again",
            "Documentation",
        },
    }
    tips := [][]string{ //2D slice of strings
        {
            "Clearly understand the problem you want to solve with your program. Identify the requirements and constraints.",
            "Devise a plan or algorithm to solve the problem. Break down the solution into smaller, manageable steps.",
            "Select a programming language that is suitable for your problem and aligns with your expertise. ",
        },
        {
            "Install the necessary tools and software to create and run your program. ",
            "Start coding based on your algorithm. Write clear, modular, and well-documented code. ",
            "Regularly test your code to ensure it behaves as expected. Use test cases to verify that your program produces the correct output for different inputs.",
        },
        {
            "If you encounter errors or unexpected behavior, use debugging tools to identify and fix issues in your code.",
            "Execute your program with different inputs to ensure it works correctly in various scenarios.",
            "Document your code by adding comments, explaining complex logic, and providing information on how to use the program.",
        },
    }
    var option int = 0
		tip  := "Professional Tip: "
		activity  := "Your recommended activity is: "

		//Start the program 
    fmt.Println("Welcome to the coding planner's game!")
    fmt.Println("To get started, please tell us how far along you are in your coding planning journey:")

		//Menu presentation
    for option != 4  {
        for _, menu := range menu {
            fmt.Printf("%s\n", menu)
        }
				//receive user input
        fmt.Print("\nSelect your Option : ")
        fmt.Scan(&option)
        switch option {
        case 1:
            fmt.Println(activity, activities[rand.Intn(2)][rand.Intn(2)])
            fmt.Println(tip, tips[rand.Intn(2)][rand.Intn(2)])
            time.Sleep(5 * time.Second) // pause 5 sec 
        case 2:
            fmt.Println(activity, activities[rand.Intn(2)][rand.Intn(2)])
            fmt.Println(tip, tips[rand.Intn(2)][rand.Intn(2)])
            time.Sleep(5 * time.Second)
        case 3:
            fmt.Println(activity, activities[rand.Intn(2)][rand.Intn(2)])
            fmt.Println(tip, tips[rand.Intn(2)][rand.Intn(2)])
            time.Sleep(5 * time.Second)
        case 4:
					fmt.Println("Exiting program")
					os.Exit(0) // exiting program - use of "os" library
        default:
            fmt.Printf("Please enter a valid option (1 to 4)")
            time.Sleep(3 * time.Second)
        }
    }
}