# Strings Utility

<p>
  <a>
    <a href="https://goreportcard.com/badge/github.com/SamuelFolledo/StringsUtility" />
    <img alt="commits" src="https://goreportcard.com/badge/github.com/SamuelFolledo/StringsUtility" target="_blank" />
    <a href="https://github.com/SamuelFolledo/StringsUtility/commits/master">
    <img alt="commits" src="https://img.shields.io/github/commit-activity/w/SamuelFolledo/StringsUtility?color=green" target="_blank" />
  </a> 
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://github.com/imthaghost/gitmoji-changelog">
    <img src="https://img.shields.io/badge/changelog-gitmoji-brightgreen.svg" alt="gitmoji-changelog">
  </a>
</p>

A CLI app written in [Go](https://golang.org/) that takes an Xcode project, and replace all strings in all ```.swift``` files to a constant variable scoped globally and writing them into a ```Constants``` file.

## Why Use?
- Avoid unintended typos
- Have strings autocompleted
- Easily manage all your strings in one file

## [Tips](Tips.md):
- Currently does not support multi line strings
- To avoid common errors, strings which contains the following will substrings not be put to the constant file or translated. Edit files accordingly 
```"/", "\\", "{", "}", "http", "https", ".com", "#", "%", "img_", "vid_", "gif_", ".jpg", ".png", ".mp4", ".mp3", ".mov", "gif", "identifier"```
    - Image named like ```UIImage(named: "heart")``` will have translate "heart" unintentionally, so consider editing the image name so it can work like this ```UIImage(named: "IMG_heart")```

## Upcoming Features
- searches and translates strings from .xib and .storyboard files
- Translate strings using [Google Cloud Translator](https://cloud.google.com/translate/docs)

## How to Use?

### Download StringsUtility
- clone the repo
  ```
  $ git clone https://github.com/SamuelFolledo/StringsUtility
  ```
- go to the project
  ```
  $ cd StringsUtility
  ```

### [Install Golang](https://sourabhbajaj.com/mac-setup/Go/README.html) with Homebrew:
  ```
  $ brew update
  $ brew install golang
  ```

### Setup Google Cloud Translator: [Basic Setup Instruction](https://cloud.google.com/translate/docs/basic/setup-basic)
- [ ] Create or select a project
- [ ] Enable the Cloud Translation API for this project
- [ ] Download a private key as JSON
- [ ] While inside StringsUtility, run in terminal
  ```
  $ go get -u cloud.google.com/go/translate
  ```
- [ ] __Important:__ Run this command __once__ each time the project starts to set the environment variable. Replace the ```[PATH]``` to the path of the ```.json``` file downloaded from setup 2 step 3. [Instructions](https://cloud.google.com/docs/authentication/production) for more info or for Windows setup
    ```
    export GOOGLE_APPLICATION_CREDENTIALS="[PATH]"
    ```
    For example:
    ```
    export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/[FILE_NAME].json"
    ```
    
### Run StringsUtility
- run the program locally replacing PATH_TO_YOUR_PROJECT with your project directory
  ```
  $ go build && go run main.go -dir=PATH_TO_YOUR_PROJECT
  ``` 

## Demo
<img src="https://github.com/SamuelFolledo/StringsUtility/blob/master/static/stringsUtilityDemo.gif" width="896" height="521">

### Links
- [MakeSchool's](makeschool.com)'s [BEW2.5: Patterns & Practices in Strongly Typed Languages](https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Languages/#/) - the only university that teaches Go
- [MakeUtility Requirements](https://github.com/Make-School-Courses/BEW-2.5-Strongly-Typed-Languages/blob/master/Project/MakeUtility.md)
- [How to localized your iOS app](https://github.com/Make-School-Courses/BEW-2.5-Strongly-Typed-Languages/blob/master/Project/MakeUtility.md)
- [Make School Spring Intensive 1.3](https://github.com/Make-School-Courses/INT-1.3-AND-INT-2.3-Spring-Intensive)
- [Make School Spring Intensive 1.3 Tracker](https://docs.google.com/spreadsheets/u/2/d/1VwXNWcWpcLQuZCEwvPO1_W0JiarsNiL0nBzUZZEyAGQ/edit#gid=0)

Lincense under [MIT License](LICENSE)