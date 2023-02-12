# Go Projects

- [Go Projects](#go-projects)
  - [Some Point to Keep In Mind](#some-point-to-keep-in-mind)
  - [Steps to Commit/Checkin the Code](#steps-to-commitcheckin-the-code)
  - [Commit using command line](#commit-using-command-line)
  - [Steps](#steps)
  - [Information About Go](#information-about-go)
  - [Why Is Go A Top Programming Language](#why-is-go-a-top-programming-language)
  - [Prerequisites To Learn Golang](#prerequisites-to-learn-golang)
  - [Why Golang Is Faster Than Others](#why-golang-is-faster-than-others)
  - [1. Golang vs. Java](#1-golang-vs-java)
  - [2. Golang vs. Node.js](#2-golang-vs-nodejs)
  - [3. Golang vs. C++](#3-golang-vs-c)
  - [When Should You Use Golang?](#when-should-you-use-golang)

## Some Point to Keep In Mind

- Always check the branch you are working on
- Also always make sure that your branch is up to date with your parent branch i.e to take latest pull of both child and parent branch and then run command `git merge parent-branch`

## Steps to Commit/Checkin the Code

- First add the files to staging that you want to commit
- Make sure you don't add files that are not intended to commit
- Now once files are in staging you can now commit the same using an appropriate commit message
- After commit make sure you push those changes

## Commit using command line

- Add the untarcked file to tracking (if any) usinf command `git add filename`
- Now your file is in staging
- Now you can commit your file using command `git commit -m "Message"`
- Now you could have also done the above 2 steps in single command using `git commit -am "Message"` but avoid this to avoid unnecessary checkins
- Now you have to push those commit as well using command `git push`

## Steps

1. Remember The Branch Name
2. Switch To Current Working Branch
3. Take pull using "git pull" command also make sure you take latest pull of the parent branch before merging it to child branch
4. update the branch from main/parent using "git merge parent-branch-name" command
5. Continue The Modifications Or Updates You Do After It's Done Then Commit All The Changes With Proper Commit Message Take Pull And Then Push The Changes And At Last Create PR

## Information About Go

1. Go Is Not OOP Nor POP Language
2. Go (also called GoLang or Go Language) is an open source programming language used for general purpose
3. It is a relatively Fast Programing Language

## Why Is Go A Top Programming Language

1. Development:
Because of the language design, Golang is easy to understand with simple syntax. Though sometimes it is called as language with verbose, but on the other hand it also assures the clarity due to its roots from C. When it comes to development, IDE (Integrated Development Environment) is equally important and plays a critical role in overall productivity. And with the evolving world there is good support for Golang in IDEs for rapid development, debugging tools and integration with deployment environments further making it programmer-friendly.

2. Binaries:
Go generates binaries for your applications with all the dependencies built in. This removes the need for you to install runtimes that are necessary for running your application. This eases the task of deploying applications and providing necessary updates across thousands of installations. With its support for multiple OS and processor architectures, this is a big win for the language.

3. Language Design:
The designers of the language made a conscious decision to keep the language simple and easy to understand. The entire specification is in a small number of pages and some interesting design decisions were made vis-à-vis Object-Oriented support in the language, that keeps the features limited. Towards this, the language is opinionated and recommends an idiomatic way of achieving things. It prefers Composition over Inheritance, and its Type System is elegant and allows for behavior to be added without tight coupling the components too much. In Go Language, “Do More with Less” is the mantra. The language is more pragmatic since it came into existence to solve practical problems at Google.

4. Powerful Standard Library:
Go comes with a powerful standard library, distributed as packages. This library caters to most components, libraries that developers have come to expect from 3rd party packages when it comes to other languages. A look at the Packages available in the standard library is good enough indication of the power available in them.

5. Package Management:
Go combines modern day developer workflow of working with Open Source projects and includes that in the way it manages external packages. Support is provided directly in the tooling to get external packages and publish your own packages in a set of easy commands.

6. Static Typing:
Go is a statically typed language, and the compiler works hard to ensure that the code is not just able to compile correctly but other type conversions and compatibility are taken care of. This can avoid the problems that one faces in dynamically typed languages, where you discover the issues only when the code is executed.

7. Concurrency Support:
One area where the language shines is its first-class support for Concurrency in the language itself. If you have programming concurrency in other languages, you understand that it is quite complex to do so. Go Concurrency primitives via go routines and channels makes concurrent programming easy. Its ability to take advantage of multi-core processor architectures and efficient memory is one of the reasons while Go code is today running some of the most heavily used applications that can scale.

8. Testing Support:
Go Language brings Unit Testing right into the language itself. It provides a simple mechanism to write your unit tests in parallel with your code. The tooling also provides support to understand code coverage by your tests, benchmarking tests and writing example code that is used in generating your code documentation.

## Prerequisites To Learn Golang

1. Some programming experience. The code here is pretty simple, but it helps to know something about functions.
2. A tool to edit your code. Any text editor you have will work fine. Most text editors have good support for Go.    The most popular are VSCode (free), GoLand (paid), and Vim (free).
3. A command terminal. Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows.

## Why Golang Is Faster Than Others

The main reason for Golang’s popularity is its speed. Many programmers have compared Golang to its competitors and found Golang to have the fastest compile time. Here’s a comparison of Golang with popular programming languages.

## 1. Golang vs. Java

Golang is much faster than Java in terms of performance and speed. Since Java is compiled on a virtual machine, its code must be changed to bytecode before passing through the compiler. Even though this step makes it a platform-independent language, it significantly slows down the compilation process. Golang doesn’t rely on a virtual machine for code compilation and is directly compiled from the binary file. That’s why it is much faster than Java when it comes to application development. Golang’s automatic garbage collection also contributes to its speed and makes it much faster than Java.

## 2. Golang vs. Node.js

Golang processing is faster and more lightweight than Node.js. Golang can also handle subroutines concurrently (i.e., it can execute threads in parallel). This is different from Node.js, which is single-threaded. This feature also improves the scalability of the final application. Golang’s goroutines are faster and more efficient compared to the single-thread architecture of Node.js. They also allow engineers to handle issues such as network timeout and database failures, which makes Golang a perfect candidate for high-performance system design.

## 3. Golang vs. C++

C++ modules often take a lot of time to parse and compile headers. Sometimes they also need to parse and compile symbol tables, which further slows down the application compilation. In contrast, Golang only uses packages that are necessary to run the program. Golang has a feature that reminds the developer to remove unused packages from the final build. It throws a compilation error whenever it detects an unused variable or import, trading short-term convenience for performance and efficiency.

## When Should You Use Golang?

Golang is a very versatile and promising programming language. Many companies use it to optimize their product performance and improve their efficiency. Furthermore, due to its goroutines that handle concurrency issues, it is the perfect language for preventing scalability bottlenecks.
Golang allows developers to create performant and lightweight applications that are optimized for performance. Golang applications are easily maintainable and vastly outperform Node.js/Java-based applications in terms of performance. For microservices and enterprise cloud projects with a focus on performance, Golang is a great option.
