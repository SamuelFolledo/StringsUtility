package main

import ( //format
	// "html/template" //allows us to do templating
	"context"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath" //to use filepath.Ext(*fileFlag) to trim file extension
	"strings"

	//"reflect" //package has TypeOf() which returns the Type of an object
	// "text/template"
	// "oset/http"
	"time"

	"cloud.google.com/go/storage"
	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

type FileLines struct {
	Title   string //capital means public, small means private
	Message string
	Done    bool
}

type Article struct {
	Author   string
	NewsList []FileLines
}

// note, that variables are pointers
var fileFlag = flag.String("file", "", "Name of file")
var dirFlag = flag.String("dir", "", "Name of directory")

func main() {
	// saveFileFlag()
	// directoryFlag()
	swiftDirectoryFlag()
}

//looks for swift files
func swiftDirectoryFlag() {
	flag.Parse()                              //parse flags
	fmt.Println("Directory flag =", *dirFlag) //after flag.Parse(), *fileFlag is now user's --file= input
	files, err := ioutil.ReadDir(*dirFlag)    //ReadDir returns a slice of FileInfo structs
	if isError(err) {
		return
	}
	for _, file := range files { //loop through each files
		// print("File: ", file.Name(), "\n")
		if file.IsDir() { //skip if file is directory
			continue
		}
		if filepath.Ext(strings.TrimSpace(file.Name())) == ".swift" { //gets the file extension from file name
			print("File: ", file.Name(), "\n")
			// textFileToHtml(file.Name())
			// -walk directories recursively
			// -filter extensions by .swift

		}
	}
}

//method that runs my implementation of Google Cloud Translation API
func implementTranslationAPI() {
	var text = "Login Error"
	var translatedText, err = translateText("es", text) //translate to Spanish
	if isError(err) {
		return
	}
	print("\"", text, "\" in SPANISH is \"", translatedText, "\"\n")

	var tagalogText, err2 = translateText("tl", text) //translate to Tagalog
	if isError(err2) {
		return
	}
	print("\"", text, "\" in TAGALOG is \"", tagalogText, "\"")
}

//function that takes a text to translate and language to translate to and returns an error or the translatedText
func translateText(targetLanguage, text string) (string, error) {
	ctx := context.Background()
	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}

//Sample code implemented in: https://cloud.google.com/storage/docs/reference/libraries
func translationOverView() { //run this code once in the beginning of translation to create a new bucket
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "go-makesite" //projectID from .json file

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new bucket.
	bucketName := "samuel-new-bucket" //bucket name must not have capital letter

	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)

	// Creates the new bucket.
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	fmt.Printf("Bucket %v created.\n", bucketName)
}

func listSupportedLanguages(w io.Writer, targetLanguage string) error {
	// targetLanguage := "th"
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("translate.NewClient: %v", err)
	}
	defer client.Close()

	langs, err := client.SupportedLanguages(ctx, lang)
	if err != nil {
		return fmt.Errorf("SupportedLanguages: %v", err)
	}

	for _, lang := range langs {
		fmt.Fprintf(w, "%q: %s\n", lang.Tag, lang.Name)
	}

	return nil
}

//method that takes a directory as a flag and find files inside that directory
func directoryFlag() {
	flag.Parse()                              //parse flags
	fmt.Println("Directory flag =", *dirFlag) //after flag.Parse(), *fileFlag is now user's --file= input
	files, err := ioutil.ReadDir(*dirFlag)    //ReadDir returns a slice of FileInfo structs
	if isError(err) {
		return
	}
	for _, file := range files { //loop through each files
		// print("File: ", file.Name(), "\n")
		if file.IsDir() { //skip if file is directory
			continue
		}
		if filepath.Ext(strings.TrimSpace(file.Name())) == ".txt" { //gets the file extension from file name
			print("File: ", file.Name(), "\n")
			textFileToHtml(file.Name())
		}
	}
}

//function that reads a text file from a directory and writes an html version of it using a GO template
func textFileToHtml(fileName string) {
	//1) Get the text of the file passed, and HTML file name
	var fileContents = readFile(("./texts/" + fileName))                       //get the contents of the file
	var trimmedFileName = strings.TrimSuffix(fileName, filepath.Ext(fileName)) //trims the fileName's extension
	var htmlFileName = "./html/" + trimmedFileName + ".html"                   //create the directory and name of the html file
	//2) Get the struct data we will store
	var news = []FileLines{
		FileLines{Title: fileName, Message: fileContents, Done: false},
	}
	var articles = Article{Author: "Samuel", NewsList: news} //contain news to articles variable
	//3) Create the HTML file, parse and execute the template with our data
	var htmlFile = createFile(htmlFileName)
	var t = template.Must(template.New("template.tmpl").ParseFiles(paths...))
	var err = t.Execute(htmlFile, articles)
	if isError(err) {
		return
	}
}

// function used to input filename to generate a new HTML file. Example: `latest-post.txt` flag will generate a `latest-post.html`
func saveFileFlag() {
	flag.Parse()                                                          //parse flags
	fmt.Println("File flag =", *fileFlag)                                 //after flag.Parse(), *fileFlag is now user's --file= input
	var fileName = strings.TrimSuffix(*fileFlag, filepath.Ext(*fileFlag)) //trims the file's extension
	var htmlFileName = "html/" + fileName + ".html"                       //takes a fileName with no extension and add a .html at the end
	//create what we will be storing to html file
	var line = populateLine()
	var news = []FileLines{
		FileLines{Title: "Title 1", Message: line, Done: true},
		FileLines{Title: "Title 2", Message: "MESSAGEE 2", Done: false},
		FileLines{Title: "Title 3", Message: "MESSAGEEE 3", Done: false},
	}
	var articles = Article{Author: "Kobe", NewsList: news}        //contain news to articles variable
	readTmplAndWriteHtml(articles, "template.tmpl", htmlFileName) //create and save an html file from whatever user named the .txt file
}

func readTmplAndWriteHtml(parsedData Article, tmplName, htmlName string) {
	var t = template.Must(template.New(tmplName).ParseFiles(paths...)) //1) parse files //template loader //1h25m is how it is actually read
	var htmlFile = createFile(htmlName)                                //2) Create html file we will be saving to
	var err = t.Execute(htmlFile, parsedData)                          //3) execute //1h26m Stdout prints it in the terminal
	if isError(err) {
		return
	}
}

func writeToFile(fileName, lines string) {
	bytesToWrite := []byte(lines)                         //data written
	err := ioutil.WriteFile(fileName, bytesToWrite, 0644) //filename, byte array (binary representation), and 0644 which represents permission number. (0-777) //will create a new text file if that text file does not exist yet
	if isError(err) {
		return
	}
}

//function that takes a fileName and extension and returns the file created
func createFile(fileName string) (returnedFile *os.File) {
	// check if file exists
	var _, err = os.Stat(fileName)
	if os.IsNotExist(err) == false { //if file exist, then delete file first
		print(fileName, " exist\n")
		deleteFile(fileName)
	}
	//create file
	var file, errr = os.Create(fileName)
	if isError(errr) {
		return
	}
	returnedFile = file
	fmt.Println("File Created Successfully", fileName)
	return
}

func deleteFile(fileName string) {
	var err = os.Remove(fileName)
	if isError(err) {
		return
	}
	fmt.Println("File Deleted")
}

func printLines(news FileLines) {
	t, err := template.New("news").Parse("You have a task titled\n\"{{ .Title}}\"\n\"{{ .Message}}\"") //1) Parse files
	if isError(err) {
		return
	}
	err = t.Execute(os.Stdout, news) //3) Execute and save the parsedFiles to file
	if isError(err) {
		return
	}
}

func populateLine() (line string) {
	directory := "/Users/macbookpro15/Desktop/MakeSite"
	fileName := "texts/sample.txt" //file we will be searching for
	file := findFile(fileName, directory)
	line = ""
	if file != nil {
		line = readFile(fileName)
	}
	return
}

func findFile(fileName, directory string) (fileResult os.FileInfo) { //func that finds a filename from a directory and returns the file found. //[]os.FileInfo is a slice of interfaces
	files, err := ioutil.ReadDir(directory) //ReadDir returns a slice of FileInfo structs
	if isError(err) {
		return
	}
	for _, file := range files { //loop through each files
		// print("File: ", file.Name(), " is ")
		if file.IsDir() { //skip if file is directory
			continue
		}
		// fmt.Print(file.IsDir(), " = ", file.Name(), "\n")
		if file.Name() == fileName {
			// println("\n\nFound", fileName)
			fileResult = file
			return
		}
	}
	return
}

func readDirectory(directory string) []os.FileInfo { //method that takes a directory and returns a list of files and directories
	files, err := ioutil.ReadDir(directory) //ReadDir returns a slice of FileInfo structs
	if isError(err) {
		return nil
	}
	return files
}

func writeFile(fileName, data string) {
	bytesToWrite := []byte("hello\ngo\n")                       //data written
	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644) //filename, byte array (binary representation), and 0644 which represents permission number. (0-777) //will create a new text file if that text file does not exist yet
	if isError(err) {
		return
	}
	print("Successful at writing file")
}

func readFile(fileName string) (content string) { //method that will read a file and return lines or error
	fileContents, err := ioutil.ReadFile(fileName)
	if isError(err) {
		return
	}
	// fmt.Print("READ FILE = \n", string(fileContents))
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
