
az login
az acr login --name avwebcontainer

docker build --platform linux/amd64 --tag avwebcontainer.azurecr.io/my-website:v1 .
docker push avwebcontainer.azurecr.io/my-website:v1
