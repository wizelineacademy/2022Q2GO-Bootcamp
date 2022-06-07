# 2022 Q2 Go Bootcamp

## Documentation
How to use.

Go to folder src/cmd/api
run command "go run ."

Endpoints:

Read csv file
http://localhost:8000/getcsvdata

Read external Api (cat fact)
http://localhost:8000/getexternalapidata

Read csv concurrently
http://localhost:8000/getcsvdataconcurrently?type=even&items=10&items_per_workers=2

## Introduction

Thank you for participating in the GO Bootcamp course!
Here, you'll find instructions for completing your certification.

## The Challenge

The purpose of the challenge is for you to demonstrate your GO skills. This is your chance to show off everything you've learned during the course!!

You will build and deliver a whole GO project on your own. We don't want to limit you by providing some fill-in-the-blanks exercises, but instead request you to build it from scratch.
We hope you find this exercise challenging and engaging.

The goal is to build a REST API which must include:

- An endpoint for reading from an external API
  - Write the information in a CSV file
- An endpoint for reading the CSV
  - Display the information as a JSON
- An endpoint for reading the CSV concurrently with some criteria (details below)
- Unit testing for the principal logic
- Follow conventions, best practices
- Clean architecture
- Go routines usage

## Requirements

These are the main requirements we will evaluate:

- Use all that you've learned in the course:
  - Best practices
  - Go basics
  - HTTP handlers
  - Error handling
  - Structs and interfaces
  - Clean architecture
  - Unit testing
  - CSV file fetching
  - Concurrency

## Getting Started

To get started, follow these steps:

1. Fork this project
1. Commit periodically
1. Apply changes according to the reviewer's comments
1. Have fun!

## Deliverables

We provide the delivery dates so you can plan accordingly; please take this challenge seriously and try to make progress constantly.

For the final deliverable, we will provide some feedback, but there is no extra review date. If you are struggling with something, contact the mentors and peers to get help on time. Feel free to use the slack channel available.

## First Deliverable (due Friday May 13th, 23:59PM)

Based on the self-study material and mentorship covered until this deliverable, we suggest you perform the following:

- Create an API
- Add an endpoint to read from a CSV file
- The CSV should have any information, for example:

```txt
1,bulbasaur
2,ivysaur
3,venusaur
```

- The items in the CSV must have an ID element (int value)
- The endpoint should get information from the CSV by some field ***(example: ID)***
- The result should be displayed as a response
- Clean architecture proposal
- Use best practices
- Handle the Errors ***(CSV not valid, error connection, etc)***

> Note: what’s listed in this deliverable is just for guidance and to help you distribute your workload; you can deliver more or fewer items if necessary. However, if you deliver fewer items at this point, you have to cover the remaining tasks in the next deliverable.

## Second Deliverable (due Friday May 27th, 23:59PM)

Based on the self-study material and mentorship covered until this deliverable, we suggest you perform the following:

- Create a client to consume an external API
- Add an endpoint to consume the external API client
- The information obtained should be stored in the CSV file
- Add unit testing
- Update the endpoint made in the first deliverable to display the result as a JSON
- Refator if needed

> Note: what’s listed in this deliverable is just for guidance and to help you distribute your workload; you can deliver more or fewer items if necessary. However, if you deliver fewer items at this point, you have to cover the remaining tasks in the next deliverable.

## 3rd Deliverable (due Monday June 6th, 9:00AM)

- Add a new endpoint
- The endpoint must read items from the CSV concurrently using a worker pool
- The endpoint must support the following query params:

```text
type: Only support "odd" or "even"
items: Is an Int and is the amount of valid items you need to display as a response
items_per_workers: Is an Int and is the amount of valid items the worker should append to the response
```

- Reject the values according to the query param ***type*** (you could use an ID column)
- Instruct the workers to shut down according to the query param ***items_per_workers*** collected
- The result should be displayed as a response
- The response should be displayed when:

  - The workers reached the limit
  - EOF
  - Valid items completed

> Important: In this deliverable all the requirements must be included. You will have 2 more days to make final changes and improve your project based on the feedback provided by your mentor, so you can submit your final project on Wednesday June 8th

## Final Deliverable (due Wednesday June 8th, 2:00PM)
> Important: this is the final deliverable, so all the requirements must be included.

## Submitting the deliverables

For submitting your work, you should follow these steps:

1. Create a pull request with your code, targeting the master branch of your fork.
2. Fill this [form](https://forms.gle/urV6szfnCVMqp4UL9) including the PR’s url
3. Stay tune for feedback
4. Do the changes according to the reviewer's comments

