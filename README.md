# GoLang
Sample snippets for Golang

# Starting Go Project
1. Install and configure Go [Download Here](https://golang.org/dl/)
2. Using the Go Command
3. Setting up an editor (Visula Studio which is widely used for Go)
4. Organizing source code

# Lets Begin With Installing Go Tools (Windows)
* Navigate [Here](https://golang.org/dl/)
* Download starts automatically and click on next each time and accept service aggrement
* When asked for Installation location , you can choose default [Recommended] or choose your own 
* After following above steps , to verify go installation type the below command in cmd :
  - go version 
* If you are able to see a valid Go version something as below :
  - **go version go1.12.6 windows/amd64**  we are ready to move ahead :smile:
  
# Seting up an Editor (Visual Studio Code)   
* You can choose any editor of your choice but I used Visual Studio (referred as VS further) which can be [Download Here](https://code.visualstudio.com/)
* Click through the prompts and accept license agreement , which launches VS automatically
* By default VS do not support Go in order to make it support for Go we need to install Go plugin , click on package manager (Square solid   box on left panel) and enter Go in search box and click install 
* Now hit Ctrl+Shift+P in VS code editor which opens up a command window search for :
  - **GO: Install/Update Tools** 
  - which lists down several plugins select them all and click OK , which makes your VS code fully equipped to work with Go
* If plugin installation is failed then follow below sequence of steps :
  1. Check whether GOROOT and GOPATH path variables are set [FOR MORE INFO](https://golang.org/doc/code.html#GOPATH)
  2. If it fails with **git command not found** ensure GIT is installed and PATH variable is set to bin(ex: C:\Program Files\Git\bin) 
  
# Writing and executing Go code snippet in VS
* In VS File->New File , create new file named main.go , save it and copy the code from [Here] (https://gobyexample.com/hello-world)
* Now open terminal in VS code using View-> Integrated Terminal **or**  Ctrl+`
* Execute command : **go run <filename.go> **
* Hurray !!! We learnt how to accomplish setup required to run Go program , long way to go :running:
  
