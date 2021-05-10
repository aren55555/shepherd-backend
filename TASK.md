# Tech challenge
A core flow within Shepherd:
* A broker creates a new application for a client. Depending on the coverage requested, a different set of fields will appear
* The broker fills in some fields
* The broker invites a client to the application
* The client fills in some fields and submits for review
* The broker can approve the application, leave comments and/or flagging fields for revisions on the form
* Once the broker is happy with the submission, a broker moves the application to “complete” and can generate documents for the application (PDF version, Excel sheets for tables in the form, and others)
* The broker can send the documents to a carrier through the tool

This tech challenge is a light version of that flow. It’s purposely created to be loosely defined, affording you space to imagine an ideal solution and capture your thinking around the end-user. Feel free to email me on points where you need further clarification. Below is the product requirements – your task is to convert the requirements into backend code. 

## Product requirements
* User flow
  * A simple web application for a “broker” to use
  * A broker selects from 2 form options
  * Depending on the form selected, they are redirected to a page with the fields for the selected form
    * Each form would render a different set of fields
    * The fields in the form are up to you. This is where you can be creative with how the fields are defined, the validation on fields (required, min length, min number, etc), different field types (input, select, date, etc), options for select/multi fields, and/or nested sections. 
    * Example fields on a form (these are just examples, no need to use them)
      * Company Name (Named Insured)
      * Other Named Insured(s)
      * Primary Email Address
      * Company Website
      * Years in Business
      * Years in business under current name
      * List all business names which applicant has used in the past
      * Describe your principle operations
      * Coverage Effective Date
      * Coverage Expiration Date
      * Federal Tax ID Number (FEIN)
      * Contractors License #
      * State(s) in which you perform operations
      * Will the contractor perform work within the five boroughs or Western New York (Buffalo/Rochester)  at any time during the policy period? If yes, please describe the applicable work.
      * Number of Full-Time Employees
      * Number of Part-Time Employees
      * Developer Name
      * Lender Name
      * Project Start Date
      * Expected Completion Date
      * Company Mailing Address
      * Project Name
      * Project Address
      * Project Description
      * Architect of Record
      * Execution Date of Design contract (if executed)
      * Coverage being requested
  * The broker fills out the form then submits the form. The data is validated on the frontend, then sent to an API
    * API validates the request has an “auth” cookie with the value “shepherd”
    * The data is validated on the backend in accordance with the fields on the form
    * The data is stored in a DB
  * Upon successful submission, the broker is redirected to a page where they can download a PDF version of the form and an Excel version of the form
    * Create either a PDF API route or Excel API route
* No need to deploy this project to the cloud. A working localhost demo is sufficient

## What we are looking for
* Don’t do any of the frontend
* Models for form, fields, and their association (mind you, a field could belong to multiple forms)
* Models for storing the values of a form
* API for all the available forms
* API for the fields for a given form
* API for submitting the data for a given form and validating it
* Storing the values in DB
* Generating a PDF or Excel from the values of a form

## Timeboxing the project
Our goal with this challenge is to simulate the type of work you would do in the role so we can understand your thinking process, approach to problem-solving, and how you might collaborate with us as part of our team. We recommend that you share your repo early and commit often so we can evaluate your thinking, and we encourage you to use frameworks and third-party tools so you don’t invest a ton of time in areas that don’t give us much extra insight. In short, we’re looking for 80/20 problem-solving. Lastly, we want to be mindful of your time, so we ask that you not exceed 4 hours of work on this project. We will accept a submission with what you’ve completed to that point


## Bonus points
These are genuinely bonus points and not required. We’d rather have you prioritize completing the challenge over checking off these bonus points.

* Project is complete
* Handling nested sections (a section has many children of which could be sections or fields)
* Using Typescript
* Using Nodejs
* Using Expressjs
* [2x bonus] share repo early, commit often. Helps us evaluate your thinking

## Submission
* Please send the GitHub repo to Mo
* Create a screen recording showing off the product (we recommend using Loom.com)
