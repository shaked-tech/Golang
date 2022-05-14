### This folder contains the docker file and configuration needed to create a custom-default-nginx-backend for nginx to pass errored replies in order to serve custom html pages, instead of the default nginx.

#### simple build:
docker build -f Dockerfile .

#### to build and push use (cd /home/develeap/Desktop/demystify/bitbucket/starfair_ui/nginx-backend/):
docker build -t eu.gcr.io/starfair/starfair_ui/nginx-backend:stable-json \
             -t eu.gcr.io/starfair/starfair_ui/nginx-backend:latest -f Dockerfile . && \
docker push eu.gcr.io/starfair/starfair_ui/nginx-backend:stable-json && \
docker push eu.gcr.io/starfair/starfair_ui/nginx-backend:latest


<!-- Debugging: -->
$ k port-forward svc/nginx-ingress-nginx-defaultbackend 8080:80

curl -H "X-Code:{***}" http://localhost:8080/

e.g.: $ curl -H "X-Code:503" http://localhost:8080/