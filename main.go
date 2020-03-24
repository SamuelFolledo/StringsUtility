package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath" //to use filepath.Ext(*fileFlag) to trim file extension
	"regexp"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"

	"github.com/SamuelFolledo/StringsUtility/github.com/copy"
	"github.com/gookit/color" //for adding colors to CLI outputs
)

type Language struct {
	Name      string
	LProj     string
	GoogleKey string
	Path      string
	Exist     bool
}

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

type ConstantVariable struct {
	Name     string
	Value    string
	Variable string
}

type Project struct {
	Name            string
	Path            string
	Directories     []Directory
	Languages       []Language
	HasConstantFile bool
	ConstantFile    File
}

// note, that variables are pointers
var fileFlag = flag.String("file", "", "Name of file")
var dirFlag = flag.String("dir", "", "Name of directory")
var kCONSTANTFILEPATH string
var kCONSTANTFILENAME string
var kCONSTANTDASHES string = "--------------------------------------------------------------------------------------------------"
var supportedLanguages = []Language{
	Language{Name: "Filipino", LProj: "fil.lproj", GoogleKey: "tl"},
	Language{Name: "Filipino (Philippines)", LProj: "fil-PH.lproj", GoogleKey: "tl"},
	Language{Name: "English", LProj: "en.lproj", GoogleKey: "en"},
	Language{Name: "English (Australia)", LProj: "en-AU.lproj", GoogleKey: "en"},
	Language{Name: "English (India)", LProj: "en-IN.lproj", GoogleKey: "en"},
	Language{Name: "English (United Kingdom)", LProj: "en-GB.lproj", GoogleKey: "en-GB"}, //
	Language{Name: "Spanish", LProj: "es.lproj", GoogleKey: "es"},
	Language{Name: "Spanish (Latin-America)", LProj: "es-419.lproj", GoogleKey: "es"},
	Language{Name: "French", LProj: "fr.lproj", GoogleKey: "fr"},
	Language{Name: "French (Canada)", LProj: "fr-CA.lproj", GoogleKey: "fr"},
	Language{Name: "Chinese, Simplified", LProj: "zh-Hans.lproj", GoogleKey: "zh-CN"},
	Language{Name: "Chinese, Traditional", LProj: "zh-Hant.lproj", GoogleKey: "zh-CN"},
	Language{Name: "Chinese (Hong Kong)", LProj: "zh-HK.lproj", GoogleKey: "zh-CN"},
	Language{Name: "Japanese", LProj: "ja.lproj", GoogleKey: "ja"},
	Language{Name: "Germany", LProj: "de.lproj", GoogleKey: "de"},
	Language{Name: "Russian", LProj: "ru.lproj", GoogleKey: "ru"},
	Language{Name: "Portugese (Portugal)", LProj: "pt-PT.lproj", GoogleKey: "pt-PT"},
	Language{Name: "Portugese (Brazil)", LProj: "pt-BR.lproj", GoogleKey: "pt-BR"},
	Language{Name: "Italian", LProj: "it.lproj", GoogleKey: "it"},
	Language{Name: "Korean", LProj: "ko.lproj", GoogleKey: "ko"},
	Language{Name: "Arabic", LProj: "ar.lproj", GoogleKey: "ar"},
	Language{Name: "Turkish", LProj: "tr.lproj", GoogleKey: "tr"},
	Language{Name: "Thailand", LProj: "th.lproj", GoogleKey: "th"},
	Language{Name: "Dutch", LProj: "nl.lproj", GoogleKey: "nl"},
	Language{Name: "Swedish", LProj: "sv.lproj", GoogleKey: "sv"},
	Language{Name: "Danish", LProj: "da.lproj", GoogleKey: "da"},
	Language{Name: "Vietnamese", LProj: "vi.lproj", GoogleKey: "vi"},
	Language{Name: "Norgwegian", LProj: "nb.lproj", GoogleKey: "no"},
	Language{Name: "Polish", LProj: "pl.lproj", GoogleKey: "pl"},
	Language{Name: "Finnish", LProj: "fi.lproj", GoogleKey: "fi"},
	Language{Name: "Indonesian", LProj: "id.lproj", GoogleKey: "id"},
	Language{Name: "Hebrew", LProj: "he.lproj", GoogleKey: "iw"},
	Language{Name: "Greek", LProj: "el.lproj", GoogleKey: "el"},
	Language{Name: "Romanian", LProj: "ro.lproj", GoogleKey: "ro"},
	Language{Name: "Hungarian", LProj: "hu.lproj", GoogleKey: "hu"},
	Language{Name: "Czech", LProj: "cs.lproj", GoogleKey: "cs"},
	Language{Name: "Catalan", LProj: "ca.lproj", GoogleKey: "ca"},
	Language{Name: "Slovak", LProj: "sk.lproj", GoogleKey: "sk"},
	Language{Name: "Ukranian", LProj: "uk.lproj", GoogleKey: "uk"},
	Language{Name: "Croatian", LProj: "hr.lproj", GoogleKey: "hr"},
	Language{Name: "Malay", LProj: "ms.lproj", GoogleKey: "ms"},
	Language{Name: "Hindi", LProj: "hi.lproj", GoogleKey: ""},
}

func main() {
	var projectPath = getDirectoryName() //get project's path directory flag
	//1) Welcome
	fmt.Print("\n" + kCONSTANTDASHES)
	color.Bold.Print("\n\n\nThank you for using Strings Utility. Our priority is to not cause any error to your project. If you see any errors, please send me an email at samuelfolledo@gmail.com or create an issue at github.com/SamuelFolledo/StringsUtility\n\n\n")
	fmt.Print(kCONSTANTDASHES, "\n")
	//2) Prompt fresh commit
	promptCommitAnyChanges()
	//3) Clone project
	fmt.Print("\n\nFinished cloning "+trimPathBeforeLastSlash(projectPath, false)+". StringsUtility is ready to make changes\n\n"+kCONSTANTDASHES, "\n")
	copy.CopyDir(projectPath, projectPath+"_previous") //clones project in the same place where the project exist"
	//4 Initialize project
	var project = Project{Name: trimPathBeforeLastSlash(projectPath, true), Path: projectPath}
	project = setupConstantFile(projectPath, project)
	//5) FEATURE 1: Prompt if user wants to put all strings to the constant file
	project = promptMoveStringsToConstant(project, projectPath, kCONSTANTFILEPATH)
	//6) Get languages supported
	project, _, _ = getProjectLanguages(project, 0, project.Path, "Localizable.strings")
	//7) FEATURE 2: Prompt if user wants to move strings in constant file to all Localizable.strings file
	project = promptMoveStringsToLocalizable(project)
	//8) FEATURE 3: Prompt if user wants to translate strings
	project = promptTranslateStrings(project)
	//9) Prompt to undo by copying contents from the cloned project
	promptToUndo(projectPath+"_previous", projectPath)
	//10) Delete the cloned project
	deleteAllFiles(projectPath + "_previous")
}

//Loop through each files and look for each strings in each lines
func moveStringsToConstant(path string, project Project) Project {
	// fmt.Println("00 Project path I need =", project.ConstantFile.Path)
	files, err := ioutil.ReadDir(path) //ReadDir returns a slice of FileInfo structs
	if isError(err) {
		return project
	}
	for _, file := range files { //loop through each files and directories
		var fileName = file.Name()
		if file.IsDir() { //if directory...
			// fmt.Println("WHILE AT PATH: ", fileName, "\t1-Project path I need =", project.ConstantFile.Path)
			if fileName == "Pods" || fileName == ".git" { //ignore Pods and .git directories
				continue
			}
			path = path + "/" + fileName                   //update directory path by adding /fileName
			project = moveStringsToConstant(path, project) //recursively call this function again
			path = trimPathAfterLastSlash(path)            //reset path by removing the / + fileName
		} else { //if file...
			// fmt.Println("WHILE AT PATH: ", fileName, "\t2-Project path I need =", project.ConstantFile.Path)
			var fileExtension = filepath.Ext(strings.TrimSpace(fileName))   //gets the file extension from file name
			if fileExtension == ".swift" && fileName != kCONSTANTFILENAME { //if we find a Swift file that's not the constants file... look for strings
				path = path + "/" + fileName
				project = handleSwiftFile(path, project)
				path = trimPathAfterLastSlash(path) //reset path by removing the / + fileName
			} //not .swift file
		}
	}
	return project
}

//looks for strings in a .swift file and updates the .swift file and Constants file accordingly
func handleSwiftFile(path string, project Project) Project {
	// fmt.Println("WHILE AT PATH: ", path, "\t3-Project path I need =", project.ConstantFile.Path)
	var fileContents = readFile(path)          //get the contents of
	lines := contentToLinesArray(fileContents) //turns fileContents to array of strings
	for _, line := range lines {               //loop through each lines
		var strArray = getStringsFromLine(line)
		var constantArray = []ConstantVariable{}
		for _, str := range strArray {
			var constantVariable = stringToConstantVariable(str)
			constantArray = append(constantArray, constantVariable)
		}
		if len(constantArray) != 0 { //if a constant exist
			for _, constant := range constantArray {
				fileContents = strings.Replace(fileContents, constant.Value, constant.Name, 1) //from fileContents, replace the doubleQuotedWord with our variableName, -1 means globally, but changed it to one at a time
				// print("\nCONTENTS ==== ", fileContents, "\n")
				updateConstantsFile(constant) //lastly, write it to our Constant file
			}
		}
	}
	replaceFile(path, fileContents) //update our .swift file with fileContents
	return project
}

//takes a line with strings and returns an array of strings
func getStringsFromLine(line string) (strArray []string) {
	if strings.Contains(line, "\"\"\"") { //if line contains """ then it's a multi line strings which is currently not supported
		return
	}
	var foundFirstQuote bool                     //initialize as false
	if i := strings.Index(line, "\""); i != -1 { //if line has "
		var startIndex = -1
		var endIndex = -1
		for i := 0; i < len(line); i++ { //loop through until we reach the end of the line. i:=1 so we ignore the first "
			switch string(line[i]) {
			case "\"": //if character is "
				if foundFirstQuote { //if second "
					foundFirstQuote = false
					endIndex = i
					var lineString = line[startIndex : endIndex+1] //line's string is in line's index from startIndex to endIndex+1
					if isValidString(lineString) {                 //append if it's a valid string
						strArray = append(strArray, lineString)
					}
				} else { //if first "... look for the second one
					startIndex = i
					foundFirstQuote = true
				}
			case "\\": //if character is \
				if foundFirstQuote { //if string has concatenation is currently not supported
					return
				}
			case "\n": //if next line... return
				return
			default: //any other characters will be ignored
				break
			}
		}
	}
	return
}

//checks if string is a valid string to be put in constant or translated
func isValidString(str string) bool {
	if len(strings.TrimSpace(str)) <= 2 { //if there is nothing in string other than "", then it is invalid string
		return false
	}
	var invalidSubstrings = []string{"/", "\\", "{", "}", "http", "https", ".com", "#", "%", "img_", "vid_", "gif_", ".jpg", ".png", ".mp4", ".mp3", ".mov", "gif", "identifier"} //these strings are not allowed in a string to be put in constant or translated
	for _, subStr := range invalidSubstrings {
		if strings.Contains(strings.ToLower(str), subStr) { //if lowerCased(str) contains invalid substring, then str is invalid
			return false
		}
	}
	return true
}

//writes constant variable to our Constants file it doesn't exist yet
func updateConstantsFile(constant ConstantVariable) {
	if constantFileContents := readFile(kCONSTANTFILEPATH); !strings.Contains(constantFileContents, constant.Name) { //if constant variable doesn't exist in our Constants file, write it
		writeToFile(kCONSTANTFILEPATH, "\n"+constant.Variable) //append the constant variable
	}
}

//search for Constant file, if it doesn't exist, create a new one
func setupConstantFile(path string, project Project) Project {
	//1. Make sure we have a Constant file
	var isFound, filePath = searchForSwiftFile(path, "Constant", false) //search for any files containing Constant
	var constantFile = File{}
	if isFound { //if a Constant file originally exist...
		constantFile.Name = trimPathBeforeLastSlash(filePath, false) //get file name from path
		constantFile.Path = filePath
	} else { //create a Constants.swift file to the same directory AppDelegate.swift is at
		constantFile = createNewConstantFile(path)
	}
	project.ConstantFile.Name = constantFile.Name
	project.ConstantFile.Path = constantFile.Path
	kCONSTANTFILEPATH = constantFile.Path //keep the reference to the path's file and name
	kCONSTANTFILENAME = constantFile.Name
	return project
}

//Creates a Constant.swift file on the same directory as the AppDelegate.swift file
func createNewConstantFile(path string) (constant File) {
	var fileNameToSearch = "AppDelegate.swift"
	constant.Name = "Constants.swift"
	var isFound, filePath = searchForSwiftFile(path, fileNameToSearch, true) //get AppDelegate's path
	if isFound {                                                             //if AppDelegate is found, create our Constants.swift in this directory
		var trimmedPath = trimPathAfterLastSlash(filePath)
		// print(filePath, " trimmed is=", trimmedPath)
		constant.Path = trimmedPath + "/" + constant.Name                                               //remove AppDelegate.swift from the path which will be used to write our Constant file into
		writeToFile(constant.Path, "//Thank you for using Samuel Folledo's Go Utility\n\nimport UIKit") //NOTE: writing to xcode project doesn't automatically add the Constant.swift file to the project
	} else {
		fmt.Print("Error: Failed to find ", fileNameToSearch)
	}
	return
}

//Search a path until it finds a path that contains a fileName we are searching for. isExactName will determine if fileName must exactly match or must contain only
func searchForSwiftFile(path, fileNameToSearch string, isExactName bool) (isFound bool, filePath string) {
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
			isFound, filePath = searchForSwiftFile(path, fileNameToSearch, isExactName) //recursively call this function again
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
					// fmt.Println("Searched and EXACTLY found ", fileNameToSearch, " at ", filePath)
					isFound = true
					return
				}
			} else { //if we want fileName to only contain
				if strings.Contains(filePath, fileNameToSearch) { //if fileName contains name of file we are looking for... it means we found our file's path
					// fmt.Println("Searched and found ", fileNameToSearch, " CONTAINS at ", filePath)
					isFound = true
					return
				}
			}
		}
	}
	return
}

//count how many Localizable.strings file there are and store their paths
func getProjectLanguages(project Project, counter int, path, fileNameToSearch string) (returnedProject Project, returnedCounter int, filePath string) {
	returnedProject = project
	returnedCounter = counter          //set counter
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
			path = path + "/" + fileName                                                                                               //update directory path by adding /fileName
			returnedProject, returnedCounter, filePath = getProjectLanguages(returnedProject, returnedCounter, path, fileNameToSearch) //recursively call this function again
			path = prevPath                                                                                                            //if not found, go to next directory, but update our path
		}
		filePath = path + "/" + fileName //path of file
		if fileName == fileNameToSearch {
			returnedCounter += 1
			var language = createLanguageFromPath(filePath)
			returnedProject.Languages = append(project.Languages, language)
		}
	}
	return
}

func createLanguageFromPath(path string) Language {
	var language = Language{}
	var languagePath = trimPathAfterLastSlash(path)
	var languageName = trimPathBeforeLastSlash(languagePath, false)
	for _, lan := range supportedLanguages {
		if languageName == lan.LProj {
			language = lan
		}
	}
	language.Path = languagePath
	return language
}

//Turn all Constant's strings to NSLocalizedString("", comment: "") strings
func localizeConstantStrings(project Project) Project {
	var fileContents = readFile(project.ConstantFile.Path)
	var linesArray = contentToLinesArray(fileContents)
	for _, line := range linesArray {
		if !strings.Contains(line, "NSLocalizedString(\"") { //if line does not contain NSLocalizedString
			if strArray := getStringsFromLine(line); len(strArray) > 0 { //ensures that the line has a string that is OK to be translated
				if len(strArray) > 1 { //little error handling that will more than likely not get executed
					unexpectedError("Line " + line + " unexpectedly have multiple strings.")
				}
				var str = strArray[0]
				var localizedStr = "NSLocalizedString(" + str + ", comment: \"\")"
				fileContents = strings.Replace(fileContents, str, localizedStr, 1) //from fileContents, replace the doubleQuotedWord with our variableName, -1 means globally, but changed it to one at a time
				updateLocalizableStrings(project, str)
			}
		}
	}
	replaceFile(project.ConstantFile.Path, fileContents)
	return project
}

//write strings to all project's Localizable.strings file
func updateLocalizableStrings(project Project, str string) Project {
	for _, lang := range project.Languages { //do it to all project's Localizable.strings file
		var path = lang.Path + "/Localizable.strings"
		var fileContents = readFile(path)
		if !strings.Contains(fileContents, str) { //if str does not exist in Localizable.strings...
			var stringToWrite = str + " = \"\";" //equivalent to: "word" = "";
			fmt.Println("Writing", stringToWrite)
			writeToFile(path, "\n"+stringToWrite) //write at the end
		}
	}
	return project
}

func searchForFilePath(counter int, path, fileNameToSearch string, isExactName bool) (returnedCounter int, isFound bool, filePath string) {
	returnedCounter = counter
	files, err := ioutil.ReadDir(path) //ReadDir returns a slice of FileInfo structs
	if isError(err) {
		return
	}
	for _, file := range files { //loop through each files and directories
		// fmt.Println("File name =", file.Name(), " c =", counter, " rc = ", returnedCounter)
		var fileName = file.Name()
		if file.IsDir() { //skip if file is directory
			if fileName == "Pods" || fileName == ".git" { //ignore Pods and .git directories
				continue
			}
			var prevPath = path
			path = path + "/" + fileName                                                                                 //update directory path by adding /fileName
			returnedCounter, isFound, filePath = searchForFilePath(returnedCounter, path, fileNameToSearch, isExactName) //recursively call this function again
			// if isFound {                                                                                                 //if we found it then keep returning
			// 	return
			// }
			path = prevPath //if not found, go to next directory, but update our path
		}
		// var fileExtension = filepath.Ext(strings.TrimSpace(fileName)) //gets the file extension from file name
		filePath = path + "/" + fileName //path of file
		if isExactName {                 //if we want the exact fileName...
			if fileName == fileNameToSearch {
				// fmt.Println("Searched and EXACTLY found ", fileNameToSearch, " at ", filePath)
				isFound = true
				returnedCounter += 1
				// return
			}
		} else { //if we want fileName to only contain
			if strings.Contains(filePath, fileNameToSearch) { //if fileName contains name of file we are looking for... it means we found our file's path
				// fmt.Println("Searched and found ", fileNameToSearch, " CONTAINS at ", filePath)
				isFound = true
				returnedCounter += 1
				// return
			}
		}
	}
	print("WE FOUND languages found", " c =", counter, " rc = ", returnedCounter, "\n\n")
	return
}

func translateProject(project Project) Project {
	var text = "I love you"
	var translatedText, err = translateText("es", text) //translate to Spanish
	if isError(err) {
		return project
	}
	print("\"", text, "\" in SPANISH is \"", translatedText, "\"\n")
	return project
}

//////////////////////////////////////////////////// MARK: Google Translate METHODS ////////////////////////////////////////////////////
//function that takes a text to translate and language to translate to and returns an error or the translatedText
func translateText(targetLanguage, text string) (string, error) {
	ctx := context.Background()
	lang, err := language.Parse(targetLanguage)
	if isError(err) {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx) ///Users/macbookpro15/Downloads/StringsUtility-Tester-785c7f11aedf.json
	if isError(err) {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if isError(err) {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}

//////////////////////////////////////////////////// MARK: PROMPTS METHODS ////////////////////////////////////////////////////

func promptCommitAnyChanges() {
	var commitConfirmation = askBooleanQuestion("RECOMMENDATION: Before using StringsUtility it is recommended to commit any changes? Say yes to continue")
	if !commitConfirmation { //if user said no, then exit program
		fmt.Println("\n" + kCONSTANTDASHES + "\n\nSince you are not ready yet, please finish commiting any changes and run StringsUtility again.")
		os.Exit(100) //exit status 100 means did not finish commmitting
	}
}

func promptMoveStringsToConstant(project Project, projectPath, constantPath string) Project {
	var shouldMoveStrings = askBooleanQuestion("FEATURE 1: Would you like StringsUtility to move all strings in .swift files to a constant file?")
	// var projectName = trimPathBeforeLastSlash(projectPath, true)
	if shouldMoveStrings {
		fmt.Print("\nMoving strings to ", project.ConstantFile.Path, "... ")
		project = moveStringsToConstant(projectPath, project) //MAKE SURE TO UNCOMMENT LATER
		// color.Style{color.Green, color.OpBold}.Print("Finished moving all strings. Reopen project and make sure there is no error.\n")
		color.Style{color.Green}.Print("Finished moving all strings.\n")
	} else {
		fmt.Println("\n\nWill not move strings.")
	}
	fmt.Println("\n" + kCONSTANTDASHES)
	return project
}

func promptMoveStringsToLocalizable(project Project) Project {
	var shouldTranslate = askBooleanQuestion("FEATURE 2: String Localization. Have you created a Localizable.strings?")
	if shouldTranslate {
		print("\nLocalizing strings...")
		project = localizeConstantStrings(project)
		color.Green.Print(" Finished moving and localizing strings.\n")
	} else {
		fmt.Print("\nWill not localize strings...")
	}
	fmt.Print("\n\n" + kCONSTANTDASHES + "\n")
	return project
}

func promptTranslateStrings(project Project) Project {
	var shouldTranslate = askBooleanQuestion("FEATURE 3: String Translation. Have you setup Google Cloud Translator?")
	if shouldTranslate {
		print("\nTranslating strings...")
		project = translateProject(project)
		finishedTranslatingMessage(project)
	} else {
		fmt.Print("\nWill not translate strings...")
	}
	fmt.Println("\n\n" + kCONSTANTDASHES + "\n")
	return project
}

func finishedTranslatingMessage(project Project) {
	color.Green.Print(" Finished translating to:")
	for i, lang := range project.Languages {
		color.Style{color.Green, color.OpBold}.Print(" " + lang.Name)
		if i < len(project.Languages)-1 { //if not in the end, append a comma
			color.Green.Print(",")
		} else {
			color.Green.Print(".")
		}
	}
}

func promptToUndo(srcPath, destPath string) {
	color.Style{color.Green, color.OpBold}.Print("Finished updating project. Reopen project and make sure there is no error.\n")
	var shouldUndo = askBooleanQuestion("QUESTION: Do you want to undo?")
	if shouldUndo {
		fmt.Print("\nUndoing...")
		undoUtilityChanges(srcPath, destPath)
		color.Style{color.Green}.Print(" Finished undoing\n")
	} else {
		color.Bold.Println("\nWe're glad that StringsUtility was a success for you")
	}
	fmt.Print("\n" + kCONSTANTDASHES + "\n")
	fmt.Print("\nFor feedbacks and issues:\n• create an issue at https://github.com/SamuelFolledo/StringsUtility/issues/new\n• or email: samuelfolledo@gmail.com")
	color.Bold.Print("\n\nThank you for using StringsUtility by Samuel P. Folledo.\n")
}

//////////////////////////////////////////////////// MARK: HELPER METHODS ////////////////////////////////////////////////////

//undo changes from project by writing file contents from previous version of the project stored when the program runs
func undoUtilityChanges(prevProjPath, projPath string) {
	files, err := ioutil.ReadDir(prevProjPath) //ReadDir returns a slice of FileInfo structs
	if isError(err) {
		return
	}
	for _, file := range files { //loop through each files and directories
		var fileName = file.Name()
		if file.IsDir() { //if directory...
			if fileName == "Pods" || fileName == ".git" { //ignore Pods and .git directories
				continue
			}
			prevProjPath = prevProjPath + "/" + fileName        //update directory path by adding /fileName
			undoUtilityChanges(prevProjPath, projPath)          //recursively call this function again
			prevProjPath = trimPathAfterLastSlash(prevProjPath) //reset path by removing the / + fileName
		} else { //if file...
			var fileExtension = filepath.Ext(strings.TrimSpace(fileName)) //gets the file extension from file name
			if fileExtension == ".swift" {                                //if we find a Swift file that's not the constants file... look for strings
				prevProjPath = prevProjPath + "/" + fileName
				var isFound, filePath = searchForSwiftFile(projPath, fileName, true) //search project for file with the same name as .swift file from previour version
				if isFound {                                                         //if found... read both file's content
					var prevContents = readFile(prevProjPath)
					var currentContents = readFile(filePath)
					if prevContents != currentContents { //if contents are not the same, replace project's file contents with the previous project's contents
						// fmt.Println("\nCopying contents of " + prevProjPath + " to " + filePath)
						replaceFile(filePath, prevContents)
					}
				} else {
					fmt.Print("Error: Failed to find ", fileName, " during undo. Please remove all changes using version control")
				}
				prevProjPath = trimPathAfterLastSlash(prevProjPath) //reset path by removing the / + fileName
			}
		}
	}
}

func askBooleanQuestion(question string) bool {
	boolAnswer := askQuestionToUser(question + "\nType yes or no: ")
	boolAnswer = strings.ToLower(boolAnswer) //lower case the response
	for {                                    //infinite loop to handle user inputs that are not expected
		if boolAnswer == "yes" || boolAnswer == "no" || boolAnswer == "y" || boolAnswer == "n" { //break if user input expected inputs
			break
		}
		fmt.Print("\n")
		// color.Error.Print("Please respond with yes or no only")
		color.Style{color.FgRed, color.OpBold}.Print("ERROR: Please respond with ")
		color.Style{color.FgRed, color.OpBold, color.OpUnderscore}.Print("yes")
		color.Style{color.FgRed, color.OpBold}.Print(" or ")
		color.Style{color.FgRed, color.OpBold, color.OpUnderscore}.Print("no")
		color.Style{color.FgRed, color.OpBold}.Print(" only")
		boolAnswer = askQuestionToUser(question + ": ")
		boolAnswer = strings.ToLower(boolAnswer) //lower case the response
	}
	if boolAnswer == "yes" || boolAnswer == "y" {
		return true
	}
	return false
}

//given a question, ask and wait for user's CLI input
func askQuestionToUser(question string) string {
	print("\n")
	// color.Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")
	color.Style{color.Cyan, color.OpBold}.Print(question)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

//create a ConstantVariable from a string
func stringToConstantVariable(str string) ConstantVariable {
	var name = capitalizedWord(str)
	return ConstantVariable{
		Name:     name,
		Value:    str,
		Variable: "public let " + name + ": String = " + str,
	}
}

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

//deletes a all files
func deleteAllFiles(path string) {
	var err = os.RemoveAll(path)
	// var err = os.Remove(path) //this will scream: directory not empty, so we must used RemoveAll()
	if isError(err) {
		return
	}
}

//prints in yellow error message and asks for email
func unexpectedError(msg string) {
	color.FgLightYellow.Println(msg + " Please email me at samuelfolledo@gmail.com if this happens")
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

func splitVariableAndString(str string) (variable, quotedWord string) {
	var strArray = strings.Split(str, "=")
	variable = strArray[0]
	quotedWord = strArray[1]
	return
}

//turns strings to array of lines
func contentToLinesArray(str string) (results []string) {
	// - strings.Fields function to split a string into substrings removing any space characters, including newlines.
	// - strings.Split function to split a string into its comma separated values
	results = strings.Split(str, "\n") //split strings by line
	// output := strings.Join(lines, "\n") //puts array to string
	return
}

//removes all non alphabets or numeric in a word
func removeAllSymbols(word string) string {
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if isError(err) {
		fmt.Print("ERROR: removing all symbols for word", word, " err = ", err)
		return word
	}
	processedString := reg.ReplaceAllString(word, "")
	return processedString
}
