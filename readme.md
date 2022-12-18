# RESTful API Design Pattern
### Introduction  
Created a RESTful API that can perform CRUD operation on a user data	from a persistence database.
### Features  
* Users can be retrieved, created, updated and deleted.   
* Server will save the creation time of the data and when the data is last updated.  
* Error handlers and log messages is implemented.
### Installation Guide  
* Clone this repository [here](https://github.com/chiragjh7/user-rest-api.git).  
* The main branch is the most stable branch at any given time, ensure you're working from it.  
* You can either work with locally installed MongoDB or run you Atlas instance. Do configure to your choice in the application entry file.  
* Create an .env file in your project root folder and add your variables.
### Usage  
* Run below command to start the application.
	`go run main.go `  
* Connect to the API using Postman or Thunderclient extension in VSCode on port 3800.
### API Endpoints  
| HTTP Verbs | Endpoints | Action |  
| --- | --- | --- |  
| GET | /user | To retrieve all existing users |  
| POST | /user | To create a new user.|  
| PUT | /user/:userId | To edit a existing user |  
| GET | /user/:userId | To retrieve a particular user |  
| DELETE | /user/:userId | To delete a single user |
### Technologies Used  
* [Golang](https://go.dev/) Golang is a procedural and statically typed programming language having the syntax similar to [**C**](https://www.geeksforgeeks.org/c-programming-language/) programming language. Sometimes it is termed as **Go Programming Language**. It provides a rich standard library, garbage collection, and dynamic-typing capability.  Golang is one of the most trending programming languages among developers.
* [Go Fiber](https://gofiber.io/) Fiber is a Go web framework built on top of Fasthttp, the  **fastest** HTTP engine for Go. It's designed to  **ease** things up for fast development with  **zero memory allocation** and  **performance** in mind.
* [MongoDB](https://www.mongodb.com/) This is a free open source NoSQL document database with scalability and flexibility. Data are stored in flexible JSON-like documents.
