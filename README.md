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

## Upcoming Features
- Localizable.strings generator
- Translate strings using [Google Cloud Translator](https://cloud.google.com/translate/docs)

## How to Use?

### [Install Golang](https://sourabhbajaj.com/mac-setup/Go/README.html) with Homebrew:
  ```
  $ brew update
  $ brew install golang
  ```

### Run StringsUtility
- clone the repo
  ```
  git clone https://github.com/SamuelFolledo/StringsUtility
  ```
- cd into the repo
  ```
  cd StringsUtility
  ```
- run the program locally replacing PATH_TO_YOUR_PROJECT with your project directory
  ```
  go build && go run main.go -dir=PATH_To_YOUR_PROJECT
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