# Sarapan Bersama

## Description: 

Donation application so people can participate with government program "Makan Bergizi Gratis"

## Background:

> The government’s new program, “Makan Bergizi Gratis” has garnered a wide range of reactions from the public, including praise, criticism, and diverse opinions. One common criticism revolves around the menu options offered in the program. This feedback sparked an idea: what if we, as members of the community, could actively participate in shaping the menu and contributing to the success of the “Makan Bergizi Gratis” program.

## Highlights:

* Microservices Architecture
* Serverless Deployment with Google Cloud Run
* Payment Gateway (Xendit)
* Email notifications

### Tech stacks:

* Go
* Echo
* gRPC
* Docker
* PostgreSQL
* MongoDB
* JWT-Authorization
* 3rd Party APIs (Xendit, SMTP)
* REST
* Swagger

## Application Flow

![Final Flow](./misc/flow.png)

## ERD

![ERD](./misc/ERD.png)

## Deployment

This app is containerized and deployed to Google Cloud Platform as a microservices. This means for each service (user-service, foundation-service, donor-service, restaurant-service and api-gateway) is a separate instance. 

RESTAURANT_SERVICE=https://restaurant-service-75625270837.asia-southeast2.run.app/
FOUNDATION_SERVICE=https://foundation-service-75625270837.asia-southeast2.run.app/
DONOR_SERVICE=https://donor-service-75625270837.asia-southeast2.run.app/
USER_SERVICE=https://user-service-75625270837.asia-southeast2.run.app/

API_GATEWAY=https://api-gateway-75625270837.asia-southeast2.run.app/
