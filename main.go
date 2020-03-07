package main

import ( //format
	// "html/template" //allows us to do templating

	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath" //to use filepath.Ext(*fileFlag) to trim file extension
	"regexp"
	"strings"
	//"reflect" //package has TypeOf() which returns the Type of an object
	// "text/template"
	// "oset/http"
)

type Directory struct {
	Path  string
	Name  string
	Files []File
}

type File struct {
	Path     string
	Name     string //file name
	Codes    []Code
	Contents []string
}

type Code struct {
	LineNumber  int    //total number of lines
	LineContent string //line content
	// Content    []string //all contents
}

type Project struct {
	Name            string
	Directories     []Directory
	HasConstantFile bool
	ConstantFile    File
}

// note, that variables are pointers
var fileFlag = flag.String("file", "", "Name of file")
var dirFlag = flag.String("dir", "", "Name of directory")
var kCONSTANTFILEPATH string

func main() {
	var projectPath = getDirectoryName()
	fmt.Println("Directory is=", projectPath)
	var project = Project{Name: projectPath}
	project = setupConstantFile(projectPath, project)
	project = searchForStrings(projectPath, project)
}

//Loop through each files and look for each strings in each lines
func searchForStrings(path string, project Project) (currentProject Project) {
	files, err := ioutil.ReadDir(path) //ReadDir returns a slice of FileInfo structs
	if isError(err) {
		return
	}
	for _, file := range files { //loop through each files and directories
		var fileName = file.Name()
		// fmt.Println("\nVisting...", fileName, "isDir=", file.IsDir())
		if file.IsDir() { //skip if file is directory
			if fileName == "Pods" || fileName == ".git" { //ignore Pods and .git directories
				continue
			}
			var prevPath = path
			path = path + "/" + fileName              //update directory path by adding /fileName
			project = searchForStrings(path, project) //recursively call this function again
			path = prevPath
		} else { //if file...
			var fileExtension = filepath.Ext(strings.TrimSpace(fileName)) //gets the file extension from file name
			if fileExtension == ".swift" {                                //if we find a Swift file... look for strings
				path = path + "/" + fileName
				var fileContents = readFile(path)
				lines := stringLineToArray(fileContents) //turn lines of strings to array of strings
				for _, line := range lines {             //loop through each lines
					var startIndex = strings.Index(line, "\"") //gets first index of "
					var endIndex = -1
					if startIndex == -1 { //if line has no "
						continue
					}
					quotedWord := line[startIndex:]          //+1 to not include the first "
					for i := 1; i < len(quotedWord)-1; i++ { //loop through until we reach the end of the line. i:=1 so we ignore the first "
						if string(quotedWord[i]) == "\"" { //if char is second "
							endIndex = i
							break //DONT BREAK IN THE FUTURE and implement a way to check strings after this line instead of moving to next line
						}
					}
					if endIndex == -1 { //check if we did get an endIndex
						continue
					}
					var doubleQuotedWord = quotedWord[:endIndex+1]
					var variableName = capitalizedWord(doubleQuotedWord)
					print("\n\nChanged: ", path, "'s ", doubleQuotedWord, " with ", variableName, "\n")
					var newFileLines = strings.Replace(fileContents, doubleQuotedWord, variableName, 1) //from fileContents, replace the doubleQuotedWord with our variableName, -1 means globally
					print("\nNew File LINES = ", newFileLines)
					replaceFile(path, newFileLines)                                        //write fileContents to our file
					project = updateConstantsFile(doubleQuotedWord, variableName, project) //lastly, write it to our Constant file
				}
				// fmt.Println("\n\nPath BEFORE trimming", path)
				path = trimPathAfterLastSlash(path)
				// fmt.Println("\n\nPath AFTER trimming", path)
			} //if not .swift file
		}
	}
	return
}

func updateConstantsFile(quotedWord, variableName string, project Project) Project {
	// var fileContents = readFile(project.ConstantFile.Path)
	var constantVariable = "\npublic let " + variableName + ": String = " + quotedWord
	var updatedFileContents = constantVariable
	replaceFile(kCONSTANTFILEPATH, updatedFileContents)
	return project
}

//turns word to kWORD
func capitalizedWord(word string) string {
	var processedWord = removeAllSymbols(word)
	return "k" + strings.ToUpper(processedWord)
}

//recursively reads a directory and get .swift files
func setupConstantFile(path string, project Project) Project {
	//1. Make sure we have a Constant file
	var isFound, filePath = searchFileLocation(path, "Constant", false) //search for any files containing Constant
	var constantFile = File{}
	if isFound { //if a Constant file originally exist...
		// fileContents := readFile(filePath)
		// contentDic := stringLineToArray(fileContents) //turn lines of strings to array of strings
		// for index, content := range contentDic {
		// 	code := Code{LineNumber: index, LineContent: content}
		// 	// fmt.Println("\nLine:", index, "=", content)
		// 	constantFile.Codes = append(constantFile.Codes, code) //append all codes line by line
		// }
		constantFile.Name = trimPathBeforeLastSlash(filePath, false) //get file name from path
		constantFile.Path = filePath
		// fmt.Println("\n========================= Swift file: ", " contents =========================\n", fileContents)
	} else { //create a Constants.swift file to the same directory AppDelegate.swift is at
		constantFile = createNewConstantFile(path)
	}
	// fmt.Println("\nConstant file's path", constantFile.Path)
	project.HasConstantFile = true
	project.ConstantFile = constantFile
	kCONSTANTFILEPATH = constantFile.Path
	return project
}

//Creates a Constant.swift file on the same directory as the AppDelegate.swift file
func createNewConstantFile(path string) (constant File) {
	var fileNameToSearch = "AppDelegate.swift"
	constant.Name = "Constants.swift"
	var isFound, filePath = searchFileLocation(path, fileNameToSearch, true) //get AppDelegate's path
	if isFound {                                                             //if AppDelegate is found, create our Constants.swift in this directory
		var trimmedPath = trimPathAfterLastSlash(filePath)
		// print(filePath, " trimmed is=", trimmedPath)
		constant.Path = trimmedPath + "/" + constant.Name                                               //remove AppDelegate.swift from the path which will be used to write our Constant file into
		writeToFile(constant.Path, "//Thank you for using Samuel Folledo's Go Utility\n\nimport UIKit") //NOTE: writing to xcode project doesn't automatically add the Constant.swift file to the project
	} else {
		fmt.Println("Error: Failed to find ", fileNameToSearch)
	}
	return
}

func handleSwiftFile(filePath string) {
	fileContents := readFile(filePath)
	print("\n========================= Swift file: ", " contents =========================\n", fileContents)
}

//replaces everything inside a file
//Note: Read first if you dont want to remove everything
func replaceFile(filePath, lines string) {
	err := ioutil.WriteFile(filePath, []byte(lines), 0)
	if isError(err) {
		return
	}
}

//append a string to a file at the end
//Usage - add constant variable to Constants.swift file
func writeToFile(fileName, line string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(line); err != nil {
		panic(err)
	}
	// bytesToWrite := []byte(lines)                         //data written
	// err := ioutil.WriteFile(fileName, bytesToWrite, 0644) //filename, byte array (binary representation), and 0644 which represents permission number. (0-777) //will create a new text file if that text file does not exist yet
	// if isError(err) {
	// 	fmt.Println("Error Writing to file:", fileName, "====", err)
	// 	return
	// }
}

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

//////////////////////////////////////////////////// MARK: HELPER METHODS ////////////////////////////////////////////////////

//Removes all strings before the last "/"
func trimPathBeforeLastSlash(path string, removeExtension bool) (fileName string) {
	if index := strings.LastIndex(path, "/"); index != -1 {
		// fmt.Println(path, " Trimmed =", path[:index])
		fileName = path[index+1:]
	}
	if removeExtension {
		if index := strings.LastIndex(fileName, "."); index != -1 {
			// fmt.Println(path, " Trimmed =", path[:index])
			fileName = path[:index] //remove all strings after the last .
		}
	}
	return fileName
}

//Removes all strings after the last "/"
func trimPathAfterLastSlash(path string) string {
	if index := strings.LastIndex(path, "/"); index != -1 {
		// fmt.Println(path, " Trimmed =", path[:index])
		return path[:index]
	}
	fmt.Println("Failed to trim strings after last '/'")
	return path
}

//Search a path until it finds a path that contains a fileName we are searching for. isExactName will determine if fileName must exactly match or must contain only
func searchFileLocation(path, fileNameToSearch string, isExactName bool) (isFound bool, filePath string) {
	files, err := ioutil.ReadDir(path) //ReadDir returns a slice of FileInfo structs
	if isError(err) {
		return
	}
	for _, file := range files { //loop through each files and directories
		var fileName = file.Name()
		if file.IsDir() { //skip if file is directory
			if fileName == "Pods" || fileName == ".git" { //ignore Pods and .git directories
				continue
			}
			var prevPath = path
			path = path + "/" + fileName                                                //update directory path by adding /fileName
			isFound, filePath = searchFileLocation(path, fileNameToSearch, isExactName) //recursively call this function again
			if isFound {                                                                //if we found it then keep returning
				return
			}
			path = prevPath //if not found, go to next directory, but update our path
		}
		var fileExtension = filepath.Ext(strings.TrimSpace(fileName)) //gets the file extension from file name
		if fileExtension == ".swift" {                                //if file is a .swift file
			filePath = path + "/" + fileName //path of file
			if isExactName {                 //if we want the exact fileName...
				if fileName == fileNameToSearch {
					fmt.Println("Searched and EXACTLY found ", fileNameToSearch, " at ", filePath)
					isFound = true
					return
				}
			} else { //if we want fileName to only contain
				if strings.Contains(filePath, fileNameToSearch) { //if fileName contains name of file we are looking for... it means we found our file's path
					fmt.Println("Searched and found ", fileNameToSearch, " CONTAINS at ", filePath)
					isFound = true
					return
				}
			}
			continue
		}
		continue
	}
	return
}

//get the directory flag name
func getDirectoryName() string {
	flag.Parse()    //parse flags
	return *dirFlag //after flag.Parse(), *fileFlag is now user's --file= input
}

func readFile(fileName string) (content string) { //method that will read a file and return lines or error
	fileContents, err := ioutil.ReadFile(fileName)
	if isError(err) {
		print("Error reading ", fileName, "====", err)
		return
	}
	// fmt.Print("READING ", fileName, " = \n", string(fileContents))
	// for index, fileContent := range fileContents {
	// 	// fmt.Println(index, " === ", string(fileContent))
	// 	if string(fileContent) == "\n" {
	// 		fmt.Println("Found newLine at", index, "\n")
	// 	}
	// 	fmt.Println("Char= ", string(fileContent))
	// }
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

func stringLineToArray(str string) (results []string) {
	// - strings.Fields function to split a string into substrings removing any space characters, including newlines.
	// - strings.Split function to split a string into its comma separated values
	results = strings.Split(str, "\n") //split strings by line
	// print("REULTS ARE", results)
	// output := strings.Join(lines, "\n") //puts array to string
	return
}

//removes all symbols in a word
func removeAllSymbols(word string) string {
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if isError(err) {
		return ""
	}
	processedString := reg.ReplaceAllString(word, "")
	return processedString
}
