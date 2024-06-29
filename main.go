package main

// import "fmt"
// import "unicode/utf8"
// import "errors"

//you can also import like this
import(
	"fmt"
	"unicode/utf8"
	"errors"
)

func main(){
	//int variable
	var intNum int = 10
	fmt.Println("Hello, World!")
	fmt.Println(intNum)

	//float variable
	var floatNum float64 = 10.5
	fmt.Println(floatNum)

	//cast math operation
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = float32(intNum32) + floatNum32
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) //uncasted division. will produce an Int.
	fmt.Println(float64(intNum1) / float64(intNum2)) //casted to float64 division.
	//string variable
	var myString string = "Hello \nWorld" //using \n to create a new line
	fmt.Println(myString)
	myString = `Hello 
World`
	fmt.Println(myString) //multi line string, but with no \n but with `
	
	fmt.Println(len("test")) //count the number of bytes is used in a string, this will not return the number of char but rather how many bytes are used. It is usable to count for length, but it will not be accurate if the character is outside the normal vanilla ASCII Character 
	
	fmt.Println(utf8.RuneCountInString("test")) //count the number of characters in a string. Requires import "unicode/utf8" to be used
	//Rune Variable
	var myRune rune = 'a' //this is a rune variable, it defines a single a single unicode character, which is basically a char from many different languages and symbol sets

	fmt.Println(myRune)

	//boolean variable
	var myBoolean bool = true
	fmt.Println(myBoolean)

	//shorthands for variables
	myShorthand := 10 //shorthand variable, this is the same as var myShorthand int = 10 but shorter as Go will infer the type by itself
	fmt.Println(myShorthand)

	//multiple variable declaration
	var var1, var2 int = 1, 2
	fmt.Println(var1, var2)
	//Multiple variable shorthand declaration
	var3, var4 := 3, 4
	fmt.Println(var3, var4)
	//Other ways to declare a variable
	var (
		var5 int = 5
		var6 int = 6
	)
	fmt.Println(var5, var6)

	//You should initialize a variable's type whenever you can, so that it wont be confusing and the code would be much more readable.
	//And you should initialize a variable's type whenever it wasn't obvious. For ex. its a return value from another function that someone doesn't know the return type of.

	//constants
	const myConst int = 10 //constant variable, this is for values that won't change. When declaring a constant you have to initialize its value from the beginning, because again, it can't be changed on the fly.
	fmt.Println(myConst)
	
	//functions
	//you can call functions in Go the same way like in other programming languages.
	myFunc()
	//you can also pass arguments to a function
	myFuncWithParam("This is an argument")
	//functions with return values
	fmt.Println(intDivision(3, 2)) //this will print the result of the division of 10 and 2 as a float64
	//functions with multiple return values
	fmt.Println(intDivisionWithMultipleReturnValues(3, 2))
	//You can also use Printf to print variables in a string. usually prefered if you're trying out/debugging a function with multiple return values.
	var resultOfDivision, remainderOfDivision = intDivisionWithMultipleReturnValues(3, 2)
	fmt.Printf("The result of the division is: %v with remainder %v", resultOfDivision, remainderOfDivision) //%v is a placeholder for the value of the variable, it will print the value of the variable in the order you put them in the function.
	fmt.Println("")
	//functions with error return values
	var resultOfDivisionWithError, err = intDivisionWithErrorReturnVal(3, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultOfDivisionWithError)
	}
}

func myFunc(){ //this is a function in Go. It doesn't have a return value so it will not return anything.
	fmt.Println("Hello, This is a function!")
}

func myFuncWithParam(printValue string){ //this is a function in Go with a parameter. For parameters you should declare their names and types and call them accordingly to their name and their paramaters as well.
	fmt.Println(printValue) //Parameters in functions are type-enforced so you can't give a function that asks for a string and give it an int.
}

func intDivision(numerator int, denominator int) float64{ //this is a function that returns a float64. As you can see from the return value type declared.
	return float64(numerator) / float64(denominator) //this line returns the value of the division of the numerator and the denominator to the caller function.
}

func intDivisionWithMultipleReturnValues(numerator int, denominator int) (float64, int){ //this is a function that returns multiple values
	var result float64 = float64(numerator) / float64(denominator)
	var result2 int = numerator % denominator
	return result, result2
}

func intDivisionWithErrorReturnVal(numerator int, denominator int) (float64, error){
	var err error = nil //this is an error variable that is initialized to nil, which means there is no error
	if denominator == 0 {
		err = errors.New("denominator cannot be 0") //this is an error message that will be returned if the denominator is 0
		return 0, err //this will return 0 and the error message
	}
	return float64(numerator) / float64(denominator), err //this will return the result of the division and the error
}
