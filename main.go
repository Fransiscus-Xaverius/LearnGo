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

	//if else and switch cases
	//if else
	var myVar int = 10
	if myVar > 5 {
		fmt.Println("myVar is greater than 5")
	} else if myVar < 5 {
		fmt.Println("myVar is less than 5")
	} else {
		fmt.Println("myVar is equal to 5")
	}

	//switch
	//conditional switch statement
	switch myVar {
		case 10:
			fmt.Println("myVar is equal to 10")
		case 5:
			fmt.Println("myVar is equal to 5")
		default:
			fmt.Println("myVar is not equal to 10 or 5")
	}

	//non-conditional switch statement
	switch {
		case myVar > 5:
			fmt.Println("myVar is greater than 5")
		case myVar < 5:
			fmt.Println("myVar is less than 5")
		default:
			fmt.Println("myVar is equal to 5")
	}

	//array
	var myArray [3]int32 //this is an array with a length of 3 and the type of int32. because the value is not initialized, the array value of each element is 0 because the default value of an int32 is 0.
	myArray[0] = 123
	fmt.Println(myArray[0])
	fmt.Println(myArray[1:3])
	//array pointers
	fmt.Println(&myArray[0]) //this will print the memory address of the first element of the array

	//alternatives to array declaration
	myArray2 := [3]int32{1, 2, 3} //myArray2 := [...]int32{1,2,3}
	fmt.Println(myArray2)

	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	//appending an array. this makes 7 to be appended to the end of the array as the last element.
	intSlice = append(intSlice, 7) 
	fmt.Println(intSlice)
	//cap function to get the capacity of an array
	fmt.Println(cap(intSlice))

	//Spread operator to append multiple elements to an array
	var intSlice2 []int32 = []int32{8,9,10}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	//map. A key-value pair data structure
	var myTestMap map[string]uint8 = make(map[string]uint8) //this creates an empty map
	fmt.Println(myTestMap)
	myTestMap["key1"] = 1
	fmt.Println(myTestMap)

	//alternative way to make a map with predifined keys and values
	var myMap map[string]int32 = map[string]int32{"key1": 1, "key2": 2}
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	//if you tried to access a key that doesn't exist you'll get the default value of the type of value of the map.
	fmt.Println(myMap["key3"]) //this returns 0 because 0 is the default value of the type int32.

	//check if key is in map
	var num, ok = myMap["key3"] //to check if key value pair exists, make an 'ok' variable and use it to check if the key-value pair exists in the map
	if ok{ //the ok variable will be true if the key is in the map
		fmt.Println(num)
	}else{
		fmt.Println("key3" + " is not in the map")
	}

	//delete key-value pair from map
	fmt.Println(myMap)
	delete(myMap, "key1")
	fmt.Println(myMap)

	//for loops
	for i := 0; i < 5; i++{
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	//for range loops
	var intSlice3 []int32 = []int32{11,12,13}
	for i, value := range intSlice3{
		fmt.Println(i, value)
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
