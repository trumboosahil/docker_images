# I'll create a markdown file with the provided README content

readme_content = """

# Go API with Docker

This project contains a simple Go API with two endpoints:

- `GET /hello`: Returns a greeting message.
- `POST /echo`: Accepts a JSON payload and echoes it back.

The steps below explain how to build the Docker image for this API, store it locally, transfer it to another machine, and run it.

## Table of Contents

1. [Create the Docker Image](#1-create-the-docker-image)
2. [Store the Image Locally](#2-store-the-image-locally)
3. [Load the Image on Another Machine](#4-load-the-image-on-another-machine)
4. [Run the Docker Image](#5-run-the-docker-image)
5. [Access the API on localhost](#6-access-the-api-on-localhost)

---

## 1. Create the Docker Image

To build the Docker image for the Go API, use the following command from the project root directory where the Dockerfile is located:

```bash
docker build -t my-go-api .
```

docker images


## 2. Store the Image Locally

After building the image, you can save it to a file for transferring to another machine or storing locally:

docker save -o my-go-api.tar my-go-api


## 3. Load the Image on Another Machine

Once the image file (`my-go-api.tar`) is transferred to another machine, you can load it into Docker using:

docker load -i my-go-api.tar


## 4. Run the Docker Image

After loading the Docker image, you can run it using:

docker run -d -p 8080:8080 my-go-api

curl http://localhost:8080/hello
