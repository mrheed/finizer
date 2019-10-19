# Finizer
File Organizer Tool based on CLI
Supported OS (Linux, Darwin, Free-BSD)
## Getting Started
Helper contents, contains explanation of each argument
	--begin,                directory investigate starting point (string)
	--copy,                 you may use this argument if you wont to move the file from the previous path (boolean)
	--help,                 this argument for detailed description each argument (boolean)
	--ignoreDefaultFormat,  skip the format filtering with default format (boolean)
	--maxStairs,            limit the depth of investigation (int)
	--move,                 you may use this argument if the file has to be moved (boolean)
	--pathToFormat,         the directory where the custom format spotted ---JSON FILE REQUIRED-- (string)
### Prerequisites
You have to install go(golang) into your machine, 
You could download and follow the installation instructions here : https://golang.org/dl/
### Installing
Download this repo using go get
```
go get github.com/syahidnurrohim/finizer
```
And then go to the package directory to build the tool
```
go build
```
Execute the file
```
./finizer
```
## Authors
**Syahid Nurrohim**
