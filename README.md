### TOKENER - UNIFIED TOKEN GENERATOR

This microservice aims to generate tokens for the clients, so they can use the token to access any microservice on the company network. It is a simple service that generates a token and stores it in a unified database. This microservice aims to solve the problem of generating tokens in disparate environments, where each microservice has its own token generator which leads to hardening application ochestration. 

## Tech Stack

- Golang
- MongoDB
- Redis
- GORM
- AWS EC2

**Golang** is a programming language created by Google. It is a statically-typed language with syntax loosely derived from that of C, adding garbage collection, type safety, some dynamic-typing capabilities, additional built-in types such as variable-length arrays and key-value maps, and a large standard library.I chose Golang for that project because it is a fast, reliable, typed and compiled language and so it is perfect for microservices. 

**MongoDB** is a cross-platform document-oriented database server. Classified as a NoSQL database server, MongoDB uses JSON-like documents with optional schemas. I chose MongoDB because it is a document-oriented database, very fast, scalable and so it is perfect for microservices. MongoDB is used for permanent storage of the tokens.

**Redis** is an in-memory data structure store, used as a distributed, in-memory key–value database, cache and message broker, with optional durability. I chose Redis because it is a very fast in-memory database, perfect for caching and so it is perfect for microservices. Redis is used for caching the tokens and other service data. Persistence is done to the MongoDB database by a background worker.

**GORM** is a Golang ORM library developed on top of database/sql. I chose GORM because it is a very powerful ORM library, easy to use and so it is perfect for microservices. GORM is used for Mongo database operations. ORM makes database operations better integrated with the code.

**AWS EC2** Amazon Elastic Compute Cloud (Amazon EC2) is a widely-used web service offered by Amazon Web Services (AWS) that provides scalable and resizable computing capacity in the cloud. It allows users to create and manage virtual machines, known as instances, on-demand. EC2 instances can be customized with various combinations of CPU, memory, storage, and networking resources to meet specific application requirements. I chose AWS EC2 because it is a very powerful cloud computing service, easy to use and so it is perfect for microservices. AWS EC2 is used for hosting the microservice.

## How it works

Please take a look at the OpenAPI specification file `openapi.yaml` for more details. A fully working instance of the service is available at [http://tokener.orionsoft.site/](http://tokener.orionsoft.site/).

For now the exposes the following routes:

- `GET /` - Returns API presentation

- `GET /reference` - Returns API reference

- `POST /token` - Generates a new token and stores it in the database. The token is valid for 24 hours. The token is returned in response body.


