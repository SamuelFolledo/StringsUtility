# StringsUtility Project Recommendations
Project tips in order to make StringsUtility as effective and error-free as possible.

## Feature 1: Moving Strings to Constant file
- [ ] Have at least __1 constants file__ (e.g. ```Constants.swift```) in your project for the strings to get stored into.

## Feature 2: Strings Localization

#### PART 1: Move strings to Localizable
<<<<<<< HEAD
- [ ] Create/Have a ```Localizable.strings``` file.
    1. In your project, New File
    2. Look for Strings file
    <img src="https://github.com/SamuelFolledo/StringsUtility/blob/master/static/pics/localizableFile.gif" width="369" height="265">
    
    3. Name it ```Localizable``` __exactly__
=======
- [ ] Create a ```Localizable.strings``` file. In your project, New File -> String File -> Name it ```Localizable``` __exactly__
    <img src="https://github.com/SamuelFolledo/StringsUtility/blob/master/static/pics/localizableFile.png" width="369" height="265">
>>>>>>> localizable

- [ ] To support more languages, go to Project -> Info -> Localizations -> ```+``` like the demo below
<img src="https://github.com/SamuelFolledo/StringsUtility/blob/master/static/gifs/multipleLocalizable.gif" width="640" height="320">

#### PART 2: String Translation
- [ ] Have [Google Cloud Translation API](https://console.cloud.google.com/apis/library/translate.googleapis.com?q=translation&project=go-makesite&folder&organizationId) setup