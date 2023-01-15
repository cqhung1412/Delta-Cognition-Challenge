# Delta-Cognition-Challenge

## Introduction

This repo is for the Delta Cognition coding challenge. The purpose of this challenge is to develop an end-to-end web application that includes the user interface (UI), APIs and database and meets the following requirements:

- The user shall login to the system using email and password.
- The system must display the list of images of dogs to the user.
- The user shall be able to upload an image of a dog.
- The user shall enter the information about a dog in the uploaded image.
- Based on the uploaded image and the entered information, the system must recommend the similar images of dogs to the user. The similarity can be defined via either the overlapping information or the appearance of the dogs.
- ***[TODO]*** The user shall select a dog image and save it to the favourite list.

## Architecture

The architecture of the application is described as below:

- UI (Reactjs + TS, Vite, AntD) authenticates users and displays the list of images of dogs to the user and allows the user to upload an image of a dog.
- Server (Golang, Gin, SQLC) provides APIs to the UI and requests from other SDK.
- Database (PostgreSQL) stores the information of the users and the dogs.
- Storage (Amazon S3) stores the images and labels of the dogs.
- Dog Detection API (AWS Lambda + Python, AWS Rekognition)is triggered by upload to S3 bucket, detects the dog in the uploaded image and store the result in another S3 bucket.

## ERD and Queries

The database schema is described in this [link](https://dbdiagram.io/d/63c1a9f5296d97641d798f8a), or in this [file](backend/db/migration/000001_init_schema.up.sql).

The queries are written in this [folder](backend/db/query)

## Results

The results is described in this [file](results/RESULT.md).

## Missing Features

- The user shall select a dog image and save it to the favourite list.

## Deployment Status

- The serverless code is deployed to AWS Lambda. This does not have any CI/CD pipeline due to IAM policy and role requirement.
- The others are not deployed to any cloud platform.
- The backend can be run locally by `docker-compose up -d` in the `backend` folder, but it requires the .env file to be configured.
- The frontend can be run locally by `pnpm run dev` in the `frontend` folder.

## Future Improvements

- Create a Github Actions pipeline to deploy the application to AWS Fargate or EKS.
- Refactor code (weakly types in Reactjs, poor error handling, etc.)
