package main

import ( //format
	// "html/template" //allows us to do templating

	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath" //to use filepath.Ext(*fileFlag) to trim file extension
	"strings"
	//"reflect" //package has TypeOf() which returns the Type of an object
	// "text/template"
	// "oset/http"
)

type Directories struct {
	Name  string
	Files []Files
}

type Files struct {
	Name  string //file name
	Codes []Code
}

type Code struct {
	LineNumber string   //total number of lines
	Line       string   //line content
	Content    []string //all contents
}

type Project struct {
	Name            string
	Directories     []Directories
	HasConstantFile bool
}

// note, that variables are pointers
var fileFlag = flag.String("file", "", "Name of file")
var dirFlag = flag.String("dir", "", "Name of directory")

func main() {
	// saveFileFlag()
	// directoryFlag()
	var projectDir = getDirectoryName()
	fmt.Println("Directory is=", projectDir)
	var project = Project{Name: projectDir}
	project = readProjectDirectory(projectDir, project)
	fmt.Println("Project is ", project)
}

func getDirectoryName() string {
	flag.Parse()    //parse flags
	return *dirFlag //after flag.Parse(), *fileFlag is now user's --file= input
}

//recursively reads a directory and get .swift files
func readProjectDirectory(directory string, project Project) Project {
	files, err := ioutil.ReadDir(directory) //ReadDir returns a slice of FileInfo structs
	if isError(err) {
		return project
	}
	for _, file := range files { //loop through each files
		var fileName = file.Name()
		if file.IsDir() { //skip if file is directory
			if fileName == "Pods" || fileName == ".git" { //ignore Pods and .git directories
				continue
			}
			// fmt.Println("Going inside directory=", file.Name())
			var prevDirectory = directory
			directory = directory + "/" + fileName //update directory path
			fmt.Println("\n\nNew Directory: ", directory)
			project = readProjectDirectory(directory, project) //recursivle call this function again
			directory = prevDirectory
			// fmt.Println("\n\nGoing Back to Directory: ", directory)
		}
		var fileExtension = filepath.Ext(strings.TrimSpace(fileName)) //gets the file extension from file name
		if fileExtension == ".swift" {                                //READ if file is a .swift file
			if !project.HasConstantFile { //if project does not have a Constants.swift file yet, check if fileName is a constant file
				if strings.Contains(fileName, "Constants") { //if fileName contains Constants, it means we already have a Constants.swift file
					project.HasConstantFile = true
					fmt.Println("We have a constant file named", fileName, " already")
					continue
				}
			}
			//Start reading files
			var filePath = directory + "/" + fileName
			contents := readFile(filePath)
			print("\n========================= Swift file: ", fileName, " contents =========================\n", contents)
			// handleSwiftFile(file)

		} else { //if fileName is not a .swift file then skip the file
			continue
		}
	}
	return project
}

func handleSwiftFile(file os.FileInfo, directory string) {
	readFile(file.Name())
}

//function that reads a text file from a directory and writes an html version of it using a GO template
// func textFileToHtml(fileName string) {
// 	//1) Get the text of the file passed, and HTML file name
// 	var fileContents = readFile(("./texts/" + fileName))                       //get the contents of the file
// 	var trimmedFileName = strings.TrimSuffix(fileName, filepath.Ext(fileName)) //trims the fileName's extension
// 	var htmlFileName = "./html/" + trimmedFileName + ".html"                   //create the directory and name of the html file
// 	//2) Get the struct data we will store
// 	var news = []FileLines{
// 		FileLines{Title: fileName, Message: fileContents, Done: false},
// 	}
// 	var articles = Article{Author: "Samuel", NewsList: news} //contain news to articles variable
// 	//3) Create the HTML file, parse and execute the template with our data
// 	var htmlFile = createFile(htmlFileName)
// 	var t = template.Must(template.New("template.tmpl").ParseFiles(paths...))
// 	var err = t.Execute(htmlFile, articles)
// 	if isError(err) {
// 		return
// 	}
// }

// // function used to input filename to generate a new HTML file. Example: `latest-post.txt` flag will generate a `latest-post.html`
// func saveFileFlag() {
// 	flag.Parse()                                                          //parse flags
// 	fmt.Println("File flag =", *fileFlag)                                 //after flag.Parse(), *fileFlag is now user's --file= input
// 	var fileName = strings.TrimSuffix(*fileFlag, filepath.Ext(*fileFlag)) //trims the file's extension
// 	var htmlFileName = "html/" + fileName + ".html"                       //takes a fileName with no extension and add a .html at the end
// 	//create what we will be storing to html file
// 	var line = populateLine()
// 	var news = []FileLines{
// 		FileLines{Title: "Title 1", Message: line, Done: true},
// 		FileLines{Title: "Title 2", Message: "MESSAGEE 2", Done: false},
// 		FileLines{Title: "Title 3", Message: "MESSAGEEE 3", Done: false},
// 	}
// 	var articles = Article{Author: "Kobe", NewsList: news}        //contain news to articles variable
// 	readTmplAndWriteHtml(articles, "template.tmpl", htmlFileName) //create and save an html file from whatever user named the .txt file
// }

// func readTmplAndWriteHtml(parsedData Article, tmplName, htmlName string) {
// 	var t = template.Must(template.New(tmplName).ParseFiles(paths...)) //1) parse files //template loader //1h25m is how it is actually read
// 	var htmlFile = createFile(htmlName)                                //2) Create html file we will be saving to
// 	var err = t.Execute(htmlFile, parsedData)                          //3) execute //1h26m Stdout prints it in the terminal
// 	if isError(err) {
// 		return
// 	}
// }

// func writeToFile(fileName, lines string) {
// 	bytesToWrite := []byte(lines)                         //data written
// 	err := ioutil.WriteFile(fileName, bytesToWrite, 0644) //filename, byte array (binary representation), and 0644 which represents permission number. (0-777) //will create a new text file if that text file does not exist yet
// 	if isError(err) {
// 		return
// 	}
// }

// //function that takes a fileName and extension and returns the file created
// func createFile(fileName string) (returnedFile *os.File) {
// 	// check if file exists
// 	var _, err = os.Stat(fileName)
// 	if os.IsNotExist(err) == false { //if file exist, then delete file first
// 		print(fileName, " exist\n")
// 		deleteFile(fileName)
// 	}
// 	//create file
// 	var file, errr = os.Create(fileName)
// 	if isError(errr) {
// 		return
// 	}
// 	returnedFile = file
// 	fmt.Println("File Created Successfully", fileName)
// 	return
// }

// func deleteFile(fileName string) {
// 	var err = os.Remove(fileName)
// 	if isError(err) {
// 		return
// 	}
// 	fmt.Println("File Deleted")
// }

// func populateLine() (line string) {
// 	directory := "/Users/macbookpro15/Desktop/MakeSite"
// 	fileName := "texts/sample.txt" //file we will be searching for
// 	file := findFile(fileName, directory)
// 	line = ""
// 	if file != nil {
// 		line = readFile(fileName)
// 	}
// 	return
// }

// func findFile(fileName, directory string) (fileResult os.FileInfo) { //func that finds a filename from a directory and returns the file found. //[]os.FileInfo is a slice of interfaces
// 	files, err := ioutil.ReadDir(directory) //ReadDir returns a slice of FileInfo structs
// 	if isError(err) {
// 		return
// 	}
// 	for _, file := range files { //loop through each files
// 		// print("File: ", file.Name(), " is ")
// 		if file.IsDir() { //skip if file is directory
// 			continue
// 		}
// 		// fmt.Print(file.IsDir(), " = ", file.Name(), "\n")
// 		if file.Name() == fileName {
// 			// println("\n\nFound", fileName)
// 			fileResult = file
// 			return
// 		}
// 	}
// 	return
// }

// func readDirectory(directory string) []os.FileInfo { //method that takes a directory and returns a list of files and directories
// 	files, err := ioutil.ReadDir(directory) //ReadDir returns a slice of FileInfo structs
// 	if isError(err) {
// 		return nil
// 	}
// 	return files
// }

// func writeFile(fileName, data string) {
// 	bytesToWrite := []byte("hello\ngo\n")                       //data written
// 	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644) //filename, byte array (binary representation), and 0644 which represents permission number. (0-777) //will create a new text file if that text file does not exist yet
// 	if isError(err) {
// 		return
// 	}
// 	print("Successful at writing file")
// }

func readFile(fileName string) (content string) { //method that will read a file and return lines or error
	fileContents, err := ioutil.ReadFile(fileName)
	if isError(err) {
		return
	}
	// fmt.Print("READING ", fileName, " = \n", string(fileContents))
	content = string(fileContents)
	return
}

func isError(err error) bool { //error helper
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return (err != nil)
}
