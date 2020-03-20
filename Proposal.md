# Juniors Spring Intensive Deliverable Proposal

Dates 3/16-3/25

**My Name:** Samuel P. Folledo

**Project Name:** Strings Utility

**Is your project New or Old?** Old

**Is your project Solo or Team?** Solo


## Description
#### Problem
Strings are easily mistyped, hard to track at times, and are a pain overall. Especially for keys for dictionaries, identifiers, JSON, and localization for multiple different languages.

#### Solution
A CLI app written in [Go](https://golang.org/) that takes an Xcode project, and replace all strings in all ```.swift``` files to a constant variable scoped globally and writing them into a ```Constants``` file. 

    //ViewController.swift

    dictionary["userId"] = user.id //will turn to "userId" to kUSERID

Constant file will look like the following:

    //Constants.swift

    public let kUSERID: String = "userId"

**Write a paragraph summary of the current status of your project, what you hope to achieve during the intensive, how and why**

A project I started less 2 weeks ago, on March 5, 2020 which takes a project directory, recursively looks for all strings from all .swift files and put it in a Constants.swift file.


## Challenges I Anticipate
**List out the challenges you anticipate for completing this project**
- Localizable files as I have never done those before
- Pretty new at Go
- Clear instructions on how others can setup the app and work with their own Google Cloud Translator API


## Skateboard
**ONE SINGLE aspect of product.**
- Be able to grab all strings from a .swift files and write it in a Constants.swift file and covering most common if not all edge cases

## Bike
**ONE additional features that get you closer to your idealized product. Examples: CRUD 2nd resource, add comments, API use, authentication, library use** 
- __Undo__ A way to go back to previous codes if for some reason there are bugs
- Interactive terminal

## Car
**ONE additional feature** 
- Write strings to Localizable files as well 
- Translate to different languages

## Airplane
- Get strings from .xib and .storyboard files


## Personal Achievement Goals:
**Each teammate must achieve 2 of 3 of their self-set personal achievement goals. If you're not on a team, delete the other teammate sections as needed.**

### Teammate 1
1. Goal 1: Complete up to Car
2. Goal 2: Write clear instructions on how to use
3. Goal 3: User tests

## Wireframes
**Insert wireframe pictures here**


## Evaluation
**You must meet the following criteria in order to pass the intensive:**
- Students must get proposal approved before starting the project to pass
- SOLO 
    - must score an average above a 2.5 on the [rubric]
- Pitch your product

[rubric]:https://docs.google.com/document/d/1IOQDmohLBEBT-hyr-2vgw1mbZUNsq3fHxVfH0oRmVt0/edit


## Approval Checklist
- [X] If I have a team project, I wrote this proposal to represent my work and only my work
- [X] I have completed all the necessary parts of this proposal
- [X] I linked my proposal in the Spring Intensive Tracker

### Links
- [Make School Spring Intensive 1.3](https://github.com/Make-School-Courses/INT-1.3-AND-INT-2.3-Spring-Intensive)
- [Make School Spring Intensive 1.3 Tracker](https://docs.google.com/spreadsheets/u/2/d/1VwXNWcWpcLQuZCEwvPO1_W0JiarsNiL0nBzUZZEyAGQ/edit#gid=0)
- [Structure of a Proposal](https://github.com/Make-School-Courses/INT-1.3-AND-INT-2.3-Spring-Intensive/blob/master/Proposals/junior-proposal.md)

### Sign off

**Student Name:**                
> Samuel Folledo

**Make School Advisor Name**
> Adriana Gonzalez