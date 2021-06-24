# Luizalabs Challenge

This project aims to develop a Rest API for querying products, buyers and managing buyers' favorite products.

# Getting Started

These instructions will provide you a copy of the project that can be run on your local machine for development and testing purposes. Consult deployment item for notes on how to deploy the project on a live system.

# Prerequisites

This package was created with go1.16 and all you need is the standard go library.

# Installing

This is what you need to install the application from the source code:

```bash
    git clone https://github.com/paraizofelipe/luizalabs-challenge
```

To build the docker version you can use the `Makefile`:

```bash
    make dk-build 
```

# Running the tests

Until I finish this README there is not so much Unit tests written.

But I will try to coverage unless 80% of unit tests for this code as soon as possible.

You can run tests like this:

```bash
    make test
```

To run this code locally for test purposes use:

```bash
    make dk-start
```

# Deployment

This codebase is cloud-native by design so you can use lots of environments to make this run anywhere you want.

But to make this even easier to you the codebase also provides a Dockerfile.

Deploy with docker:

```bash
    make dk-deploy
```

# API

## Resources

### product

- **POST** /api/product

Add a product

```bash
    curl -i -X POST "http://localhost:300/api/product/" \
    -H "Content-Type: application/json" \
    -d '{"title": "Smart TV xpto", "brand": "Paraizo", "price": 100.50, "review_score": 4.5, "image": "http://brand.com/image.png" }' 
```

- **GET** /api/product/**\<id\>**

Fetch a product by ID

```bash
    curl -i -X GET "http://localhost:300/api/product/925fa490-7fa5-4fee-8035-3b691af02cb8" 
```

- **GET** /api/product/?page=**\<page number\>**

List a page of products

```bash
    curl -i -X GET "http://localhost:300/api/product/?page=10" 
```

- **PATCH** /api/product/**\<id\>**

Updates product's information

```bash
    curl -i -X PATCH "http://localhost:300/api/product/925fa490-7fa5-4fee-8035-3b691af02cb8" \
    -H "Content-Type: application/json" \
    -d '{"title": "Smart TV xpto", "brand": "Hell", "price": 100.50, "review_score": 4.5, "image": "http://brand.com/image.png" }' 
```

- **DELETE** /api/product/**\<id\>**

Remove a product of database

```bash
    curl -i -X DELETE "http://localhost:300/api/product/925fa490-7fa5-4fee-8035-3b691af02cb8" 
```

### buyer

- **POST** /api/buyer

Create a buyer

```bash
    curl -i -X POST "http://localhost:300/api/buyer/" \
    -H "Content-Type: application/json" \
    -d '{"title": "Smart TV xpto", "brand": "Paraizo", "price": 100.50, "review_score": 4.5, "image": "http://brand.com/image.png" }' 
```
Add favorite product to buyer

```bash
    curl -i -X POST "http://localhost:300/api/buyer/925fa490-7fa5-4fee-8035-3b691af02cb8/product/f4cacf02-4e71-42ad-9fd8-ace6d27a4c87"
```

- **GET** /api/buyer/**\<id\>**

Fetch a buyer by id

```bash
    curl -i -X GET "http://localhost:300/api/buyer/925fa490-7fa5-4fee-8035-3b691af02cb8" 
```

- **PATCH** /api/buyer/**\<id\>**

Updates buyer's information

```bash
    curl -i -X PATCH "http://localhost:300/api/buyer/925fa490-7fa5-4fee-8035-3b691af02cb8" \
    -H "Content-Type: application/json" \
    -d '{"title": "Smart TV xpto", "brand": "Hell", "price": 100.50, "review_score": 4.5, "image": "http://brand.com/image.png" }' 
```

- **DELETE** /api/buyer/**\<id\>**

Remove buyer of database

```bash
    curl -i -X DELETE "http://localhost:300/api/buyer/925fa490-7fa5-4fee-8035-3b691af02cb8" 
```
