package lib

import "fmt"

var (
	CopyHelper = `
		--copy, you may use this argument if you wont to move the file from the previous path 
	`
	MoveHelper = `
		--move, you may use this argument if the file has to be moved
	`
	BeginHelper = `
		--begin, directory investigate starting point 
	`
	MaxStairsHelper = `
		--maxStairs, limit the depth of investigation
	`
	DefaultMimeHelper = `
		--ignoreDefaultFormat, skip the format filtering with default format
	`
	CustomFormatPathHelper = `
		--pathToFormat, the directory where the custom format spotted
	`
	AllHelper = `
		--help, this argument for detailed description each argument
	`
)

const allHelpers = `Helper contents, contains explanation of each argument

	--begin,                directory investigate starting point (string)

	--copy,                 you may use this argument if you wont to move the file from the previous path (boolean)

	--help,                 this argument for detailed description each argument (boolean)

	--ignoreDefaultFormat,  skip the format filtering with default format (boolean)

	--maxStairs,            limit the depth of investigation (int)

	--move,                 you may use this argument if the file has to be moved (boolean)
	
	--pathToFormat,         the directory where the custom format spotted ---JSON FILE REQUIRED-- (string)

`

func CallHelpers() {
	fmt.Println(allHelpers)
}
