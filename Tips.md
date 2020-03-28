# StringsUtility Project Recommendations
Project tips in order to make StringsUtility as effective and error-free as possible.

## Table Of Contents:
1. [Tips](#tips)
2. [Requirements](#requirements)
    - [Feature 1](#feature1)
    - [Feature 2](#feature2)
    - [Feature 3](#feature3)
3. [Common Errors](#commonErrors)
4. [Languages Supported](#languagesSupported)

<a name="tips"></a>
## Tips:
- StringsUtility currently does not support multi-line strings ```""" """``` and strings from ```.storyboard``` and ```.xib``` files
- To avoid common errors, strings which contains the following will substrings not be put to the constant file or translated. Edit files accordingly 
```"/", "\\", "{", "}", "http", "https", ".com", "#", "%", "img_", "vid_", "gif_", ".jpg", ".jpeg", ".png", ".mp4", ".mp3", ".wav", ".mov", "gif", "identifier", "json_", "dic_"```
    - Image named like ```UIImage(named: "heart")``` will have translate "heart" unintentionally, so consider editing the image name so it can work like this ```UIImage(named: "IMG_heart")```

<a name="requirements"></a>
## Requirements
<a name="feature1"></a>
### FEATURE 1: Moving Strings to Constant file
- [ ] Have at least __1 constants file__ (e.g. ```Constants.swift```) in your project for the strings to get stored into.

<a name="feature2"></a>
### FEATURE 2: Strings Localization
- [ ] Create a ```Localizable.strings``` file. In your project, New File -> String File -> Name it ```Localizable``` __exactly__
    <img src="https://github.com/SamuelFolledo/StringsUtility/blob/master/static/pics/mediumImages/createStrings.png" width="750" height="273">

- [ ] To support more languages, go to Project -> Info -> Localizations -> ```+``` like the demo below
<img src="https://github.com/SamuelFolledo/StringsUtility/blob/master/static/gifs/multipleLocalizable.gif" width="750" height="493">

<a name="feature3"></a>
### FEATURE 3: String Translation
- [ ] Have [Google Cloud Translation API](https://console.cloud.google.com/apis/library/translate.googleapis.com?q=translation&project=go-makesite&folder&organizationId) setup

<a name="commonErrors"></a>
## Common Error Fixes
- ```dialing: google: could not find default credentials.```   
    - __Make sure you have done the following:__
    - [ ] ```go get -u cloud.google.com/go/translate```
    - [ ] ```export GOOGLE_APPLICATION_CREDENTIALS=[PATH]```


<a name="languagesSupported"></a>
## Languages Currently Supported
| Language                                                   	| Xcode```.lproj``` Key 	| Google Key  	|
|------------------------------------------------------------	|-----------------------	|-------------	|
| English                                                    	| en.lproj              	| en          	|
| English (Australia)                                        	| en-AU.lproj           	| en          	|
| English (India)                                            	| en-IN.lproj           	| en          	|
| English (United Kingdom)                                   	| en-GB.lproj           	| en-GB       	|
| Filipino                                                   	| fil.lproj             	| tl          	|
| Filipino (Philippines)                                     	| fil-PH.lproj          	| tl          	|
| Spanish                                                    	| es.lproj              	| es          	|
| Spanish (Latin-America)                                    	| es-419.lproj          	| es          	|
| French                                                     	| fr.lproj              	| fr          	|
| French (Canada)                                            	| fr-CA.lproj           	| fr          	|
| Chinese, Simplified [Mandarin @ Mainland China, Singapore] 	| zh-Hans.lproj         	| zh-CN or zh 	|
| Chinese, Traditional [Mandarin @ Taiwan]                   	| zh-Hant.lproj         	| zh-TW       	|
| Chinese (Hong Kong) [Cantonese @ Hong Kong]                	| zh-HK.lproj           	| zh-CN or zh 	|
| Japanese                                                   	| ja.lproj              	| ja          	|
| Germany                                                    	| de.lproj              	| de          	|
| Russian                                                    	| ru.lproj              	| ru          	|
| Portugese (Portugal)                                       	| pt-PT.lproj           	| pt-PT       	|
| Portugese (Brazil)                                         	| pt-BR.lproj           	| pt-BR       	|
| Italian                                                    	| it.lproj              	| it          	|
| Korean                                                     	| ko.lproj              	| ko          	|
| Arabic                                                     	| ar.lproj              	| ar          	|
| Turkish                                                    	| tr.lproj              	| tr          	|
| Thailand                                                   	| th.lproj              	| th          	|
| Dutch                                                      	| nl.lproj              	| nl          	|
| Swedish                                                    	| sv.lproj              	| sv          	|
| Danish                                                     	| da.lproj              	| da          	|
| Vietnamese                                                 	| vi.lproj              	| vi          	|
| Norgwegian                                                 	| nb.lproj              	| nb          	|
| Polish                                                     	| pl.lproj              	| pl          	|
| Finnish                                                    	| fi.lproj              	| fi          	|
| Indonesian                                                 	| id.lproj              	| id          	|
| Hebrew                                                     	| he.lproj              	| he or iw    	|
| Greek                                                      	| el.lproj              	| el          	|
| Romanian                                                   	| ro.lproj              	| ro          	|
| Hungarian                                                  	| hu.lproj              	| hu          	|
| Czech                                                      	| cs.lproj              	| cs          	|
| Catalan                                                    	| ca.lproj              	| ca          	|
| Slovak                                                     	| sk.lproj              	| sk          	|
| Ukranian                                                   	| uk.lproj              	| uk          	|
| Croatian                                                   	| hr.lproj              	| hr          	|
| Malay                                                      	| ms.lproj              	| ms          	|
| Hindi                                                      	| hi.lproj              	| hi          	|