# Internship Progress - September 28, 2026

## 1. What I Worked On Today

Today, I worked on **database storage and started developing the backend using Go (Golang)**.

I focused on understanding how application data can be stored in a PostgreSQL database and how a Go backend can communicate with the database.

I also started creating the basic structure of the Go backend and prepared it for implementing API functionality.

## 2. What I Learned

Today I learned:

- How data is stored in PostgreSQL
- How to create and manage database tables
- How to insert data into a database
- How to retrieve stored data
- Basic PostgreSQL CRUD operations
- How a Go backend communicates with a database
- Basic Go backend project structure
- How database connections work in Go
- The role of APIs in connecting the frontend, backend, and database
- How backend applications process requests and store data

## 3. What I Implemented

I implemented the initial database storage functionality using PostgreSQL.

I also started creating the backend application using Go.

The work included:

- Creating the required database
- Creating database tables
- Defining appropriate columns and data types
- Storing application data in PostgreSQL
- Testing database operations
- Setting up the initial Go backend structure
- Preparing the backend for API development
- Working on the database connection from the Go backend

## 4. How I Implemented It

First, I created the required PostgreSQL database and tables with appropriate columns and data types.

I then tested storing data in the database using SQL queries and verified that the data was being saved correctly.

After completing the initial database setup, I started creating the Go backend.

The basic architecture I worked on is:

```text
React Frontend
      ↓
Go Backend
      ↓
REST API
      ↓
PostgreSQL Database
      ↓
Stored Data
```

The Go backend will receive requests from the frontend, process the required operations, communicate with PostgreSQL, and return the appropriate response to the frontend.

## 5. Any Problems or Errors I Faced

- Initially, understanding the connection between the Go backend and PostgreSQL was challenging.
- I had some difficulty understanding how database connections are handled in Go.
- Understanding how the backend should communicate with the database required additional practice.
- I also had to make sure that the database tables and data types were defined correctly.
- Understanding the overall flow between the frontend, backend, and database was initially confusing.

## 6. How I Solved Them

I solved these problems by working on the database setup step by step and testing the SQL queries individually.

I verified that the database and tables were created correctly before connecting them to the Go backend.

I also studied the basic Go database connection process and organized the backend into separate components so that the database and API logic can be managed properly.

Testing each part individually helped me identify and understand errors before moving to the next step.

## 7. What I Plan to Work On Next

Next, I plan to continue developing the **Go backend** and connect it more completely with PostgreSQL.

I plan to work on:

- Establishing a proper PostgreSQL connection in Go
- Creating REST API endpoints
- GET, POST, PUT, and DELETE operations
- Implementing CRUD functionality
- Sending and receiving JSON data
- Connecting the frontend with the Go backend
- Handling API requests and responses
- Implementing error handling
- Organizing the Go backend project structure
- Testing the APIs with sample data

My goal is to complete the basic flow:

```text
React Frontend
      ↓
Go REST API
      ↓
PostgreSQL
      ↓
Database Storage
```

and gradually convert it into a working full-stack application.
