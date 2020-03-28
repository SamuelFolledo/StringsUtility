# StringsUtility Project Recommendations
Project tips in order to make StringsUtility as effective and error-free as possible.

## Tips:
- Currently does not support multi-line strings
- To avoid common errors, strings which contains the following will substrings not be put to the constant file or translated. Edit files accordingly 
```"/", "\", "{", "}", "http", "https", ".com", "#", "%", "img_", "vid_", "gif_", ".jpg", ".png", ".mp4", ".mp3", ".mov", "gif", "identifier"```
    - Image named like ```UIImage(named: "heart")``` will have translate "heart" unintentionally, so consider editing the image name so it can work like this ```UIImage(named: "IMG_heart")```

## Requirements
### FEATURE 1: Moving Strings to Constant file
- [ ] Have at least __1 constants file__ (e.g. ```Constants.swift```) in your project for the strings to get stored into.

### FEATURE 2: Strings Localization
- [ ] Create a ```Localizable.strings``` file. In your project, New File -> String File -> Name it ```Localizable``` __exactly__
    <img src="https://github.com/SamuelFolledo/StringsUtility/blob/master/static/pics/localizableFile.png" width="369" height="265">

- [ ] To support more languages, go to Project -> Info -> Localizations -> ```+``` like the demo below
<img src="https://github.com/SamuelFolledo/StringsUtility/blob/master/static/gifs/multipleLocalizable.gif" width="478" height="238">

### FEATURE 3: String Translation
- [ ] Have [Google Cloud Translation API](https://console.cloud.google.com/apis/library/translate.googleapis.com?q=translation&project=go-makesite&folder&organizationId) setup

## Common Errors
- ```dialing: google: could not find default credentials.```   
    - __Make sure you have done the following:__
    - [ ] ```go get -u cloud.google.com/go/translate```
    - [ ] ```export GOOGLE_APPLICATION_CREDENTIALS="[PATH]"```


## Languages Currently Supported

| Language                                                   	| Xcode Key    	| Google Key  	|
|------------------------------------------------------------	| --------------| -------------	|
| English                                                    	| en           	| en          	|
| English (Australia)                                        	| en-AU        	| en          	|
| English (India)                                            	| en-IN        	| en          	|
|  English (United Kingdom)                                  	| en-GB        	| en-GB       	|
| Filipino                                                   	| fil          	| tl          	|
| Filipino (Philippines)                                     	| fil-PH       	| tl          	|
| Spanish                                                    	| es           	| es          	|
| Spanish (Latin-America)                                    	| es-419.lproj 	| es          	|
| French                                                     	| fr           	| fr          	|
| French (Canada)                                            	| fr-CA        	| fr          	|
| Chinese, Simplified [Mandarin @ Mainland China, Singapore] 	| zh-Hans      	| zh-CN or zh 	|
| Chinese, Traditional [Mandarin @ Taiwan]                   	| zh-Hant      	| zh-TW       	|
| Chinese (Hong Kong) [Cantonese @ Hong Kong]                	| zh-HK        	| zh-CN or zh 	|
| Japanese                                                   	| ja           	| ja          	|
| Germany                                                    	| de           	| de          	|
| Russian                                                    	| ru           	| ru          	|
| Portugese (Portugal)                                       	| pt-PT        	| pt-PT       	|
| Portugese (Brazil)                                         	| pt-BR        	| pt-BR       	|
| Italian                                                    	| it           	| it          	|
| Korean                                                     	| ko           	| ko          	|
| Arabic                                                     	| ar           	| ar          	|
| Turkish                                                    	| tr           	| tr          	|
| Thailand                                                   	| th           	| th          	|
| Dutch                                                      	| nl           	| nl          	|
| Swedish                                                    	| sv           	| sv          	|
| Danish                                                     	| da           	| da          	|
| Vietnamese                                                 	| vi           	| vi          	|
| Norgwegian                                                 	| nb           	| nb          	|
| Polish                                                     	| pl           	| pl          	|
| Finnish                                                    	| fi           	| fi          	|
| Indonesian                                                 	| id           	| id          	|
| Hebrew                                                     	| he           	| he or iw    	|
| Greek                                                      	| el           	| el          	|
| Romanian                                                   	| ro           	| ro          	|
| Hungarian                                                  	| hu           	| hu          	|
| Czech                                                      	| cs           	| cs          	|
| Catalan                                                    	| ca           	| ca          	|
| Slovak                                                     	| sk           	| sk          	|
| Ukranian                                                   	| uk           	| uk          	|
| Croatian                                                   	| hr           	| hr          	|
| Malay                                                      	| ms           	| ms          	|
| Hindi                                                      	| hi           	| hi          	|