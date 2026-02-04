package main

//Go knows where our application starts, at which files it should take a look first, by looking for a package named main.

//main is a special package name in Go. bcz it tells Go that this package will be the main entry point of the application we are building.

//one module consist of multiple packages.
//module is just go project.

/*
While working with go you split your code across packages you must have at least one package per application, per Go program, but you can have multiple packages.
And a single package can also be split across multiple files.

Note:- You can have multiple files that make up one package, and you can have multiple packages in one Go project.
*/

import "fmt"

//importing fmt package to use its functions

//which code to execute when the application starts
func main() {
	fmt.Println("Hello, World!")
}

//go mod init - this will turn this project into a module
// go mod init <github-url> - where you plan to host your module later if you plan to expose it as a third party library to the world.
// go mod init <dummy-path>
//eg. go mod init example.com/hello-world

//go build - in your go project folder, and this will then tell Go to in the end create an executable file that could also be executed on systems that do not have Go installed.

// ./hello-world - this is the name of the executable file that will be created after running go build command. You can give any name to your executable file.
