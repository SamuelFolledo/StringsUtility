package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath" //to use filepath.Ext(*fileFlag) to trim file extension
	"regexp"
	"strings"
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
var kCONSTANTFILENAME string

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
		if file.IsDir() { //if directory...
			if fileName == "Pods" || fileName == ".git" { //ignore Pods and .git directories
				continue
			}
			path = path + "/" + fileName                     //update directory path by adding /fileName
			currentProject = searchForStrings(path, project) //recursively call this function again
			path = trimPathAfterLastSlash(path)              //reset path by removing the / + fileName
		} else { //if file...
			var fileExtension = filepath.Ext(strings.TrimSpace(fileName))   //gets the file extension from file name
			if fileExtension == ".swift" && fileName != kCONSTANTFILENAME { //if we find a Swift file that's not the constants file... look for strings
				path = path + "/" + fileName
				currentProject = handleSwiftFile(path, project)
				path = trimPathAfterLastSlash(path) //reset path by removing the / + fileName
			} //not .swift file
		}
	}
	return
}

//looks for strings in a .swift file and updates the .swift file and Constants file accordingly
func handleSwiftFile(path string, project Project) (currentProject Project) {
	var fileContents = readFile(path)        //get the contents of
	lines := stringLineToArray(fileContents) //turns fileContents to array of strings
	for lineIndex, line := range lines {     //loop through each lines
		if startIndex := strings.Index(line, "\""); startIndex != -1 { //gets first index of line that has ", else go to next line
			var endIndex = -1
			quotedWord := line[startIndex:]        //remove all strings before first "
			for i := 1; i < len(quotedWord); i++ { //loop through until we reach the end of the line. i:=1 so we ignore the first "
				if string(quotedWord[i]) == "\"" { //if char is second "
					endIndex = i
					break //DONT BREAK IN THE FUTURE and implement a way to check strings after this line instead of moving to next line
				}
			}
			if endIndex != -1 { //if we found the next "... update our .swift file and Constants file
				var doubleQuotedWord = quotedWord[:endIndex+1]
				var variableName = capitalizedWord(doubleQuotedWord)
				print("\n\nChanged: ", path, ", line: ", lineIndex, " ", doubleQuotedWord, " to ", variableName, "\n")
				fileContents = strings.Replace(fileContents, doubleQuotedWord, variableName, 1) //from fileContents, replace the doubleQuotedWord with our variableName, -1 means globally, but changed it to one at a time
				currentProject = updateConstantsFile(doubleQuotedWord, variableName, project)   //lastly, write it to our Constant file
			}
		}
	}
	replaceFile(path, fileContents) //update our .swift file with fileContents
	return
}

//takes a line with strings and returns an array of strings
func getStringsFromLine(line string) (stringsArray []string) {

	return
}

//writes constant variable to our Constants file it doesn't exist yet
func updateConstantsFile(quotedWord, variableName string, project Project) Project {
	var constantVariable = "\npublic let " + variableName + ": String = " + quotedWord
	if constantFileContents := readFile(kCONSTANTFILEPATH); strings.Contains(constantFileContents, variableName) { //if constant variable doesn't exist in our Constants file, write it
		writeToFile(kCONSTANTFILEPATH, constantVariable)
	}
	return project
}

//search for Constant file, if it doesn't exist, create a new one
func setupConstantFile(path string, project Project) Project {
	//1. Make sure we have a Constant file
	var isFound, filePath = searchFileLocation(path, "Constant", false) //search for any files containing Constant
	var constantFile = File{}
	if isFound { //if a Constant file originally exist...
		constantFile.Name = trimPathBeforeLastSlash(filePath, false) //get file name from path
		constantFile.Path = filePath
	} else { //create a Constants.swift file to the same directory AppDelegate.swift is at
		constantFile = createNewConstantFile(path)
	}
	kCONSTANTFILEPATH = constantFile.Path //keep the reference to the path's file and name
	kCONSTANTFILENAME = constantFile.Name
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
		}
	}
	return
}

//////////////////////////////////////////////////// MARK: HELPER METHODS ////////////////////////////////////////////////////

//replaces everything inside a file
//Note: Read first if you dont want to remove everything before writing
func replaceFile(filePath, lines string) {
	bytesToWrite := []byte(lines)                         //data written
	err := ioutil.WriteFile(filePath, bytesToWrite, 0644) //filename, byte array (binary representation), and 0644 which represents permission number. (0-777) //will create a new text file if that text file does not exist yet
	if isError(err) {
		fmt.Println("Error Writing to file:", filePath, "====", err)
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
}

//turns word to kWORD
func capitalizedWord(word string) string {
	var processedWord = removeAllSymbols(word)
	return "k" + strings.ToUpper(processedWord)
}

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

//Updates path by removing all strings after the last "/".
func trimPathAfterLastSlash(path string) string {
	if index := strings.LastIndex(path, "/"); index != -1 {
		// fmt.Println(path, " Trimmed =", path[:index])
		return path[:index] //remove including the last /
	}
	fmt.Println("Failed to trim strings after last '/'")
	return path
}

//get the directory flag name
func getDirectoryName() string {
	flag.Parse()    //parse flags
	return *dirFlag //after flag.Parse(), *fileFlag is now user's --file= input
}

//reads file given a path
func readFile(fileName string) (content string) { //method that will read a file and return lines or error
	fileContents, err := ioutil.ReadFile(fileName)
	if isError(err) {
		print("Error reading ", fileName, "====", err)
		return
	}
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

//turns strings to array of lines
func stringLineToArray(str string) (results []string) {
	// - strings.Fields function to split a string into substrings removing any space characters, including newlines.
	// - strings.Split function to split a string into its comma separated values
	results = strings.Split(str, "\n") //split strings by line
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
