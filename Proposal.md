## Proposal
Strings are easily mistyped, hard to track at times, and are a pain overall. Especially for keys for dictionaries, identifiers, JSON, and localization for multiple different languages.

## Solution
A CLI app written in [Go](https://golang.org/) that takes an Xcode project, and replace all strings in all ```.swift``` files to a constant variable scoped globally and writing them into a ```Constants``` file. 
    ```
    dictionary["userId"] = user.id //will turn to "userId" to kUSERID
    ```

    ```
    //Constants.swift

    public let kUSERID: String = "userId"
    ```

