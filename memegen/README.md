# Memegen

A simple web service that generates a meme image given text and an image URL. 

### URL parameters:

* `image`: URL of the image
* `top`:  text to add at the top of the image
* `bottom`:  text to add at the bottom of the image

## Running the server locally

* Build with `docker build . -t memegen`
* Start with `docker run -p 8080:8080 memegen`
* Open in your browser at `http://localhost:8080/`


## local
```bash
# Set up the environment
$ export PROJECT_ID=test-yokoo
# build
$ docker build -t asia-northeast1-docker.pkg.dev/$PROJECT_ID/gitops-demo/memegen:blue . --no-cache
# tag
$ docker tag asia-northeast1-docker.pkg.dev/$PROJECT_ID/gitops-demo/memegen:blue memegen:blue
# push
$ docker push asia-northeast1-docker.pkg.dev/$PROJECT_ID/gitops-demo/memegen:blue
```