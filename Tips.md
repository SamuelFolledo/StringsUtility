# StringsUtility Project Recommendations
Project tips in order to make StringsUtility as effective and error-free as possible.

## Tips:
- To avoid unnecessary common errors, strings which contains the following will substrings not be put to the constant file or translated. Edit files accordingly 
```"/", "\\", "{", "}", "http", "https", ".com", "#", "%", "img_", "IMG_", "vid_", "VID_", "gif_", "GIF_"```
    - Files named like ```UIImage(named: "heart")``` will be translated unintentionally, so consider adding ```"IMG_heart"```

## Feature 1: Moving Strings to Constant file
- [ ] Have at least __1 constants file__ (e.g. ```Constants.swift```) in your project for the strings to get stored into.

## Feature 2: Strings Localization

#### PART 1: Move strings to Localizable
- [ ] Create a ```Localizable.strings``` file. In your project, New File -> String File -> Name it ```Localizable``` __exactly__
    <img src="https://github.com/SamuelFolledo/StringsUtility/blob/master/static/pics/localizableFile.png" width="369" height="265">

- [ ] To support more languages, go to Project -> Info -> Localizations -> ```+``` like the demo below
<img src="https://github.com/SamuelFolledo/StringsUtility/blob/master/static/gifs/multipleLocalizable.gif" width="478" height="238">

#### PART 2: String Translation
- [ ] Have [Google Cloud Translation API](https://console.cloud.google.com/apis/library/translate.googleapis.com?q=translation&project=go-makesite&folder&organizationId) setup