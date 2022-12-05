docker stop efactura
docker rm efactura
docker rmi efactura-image
docker build -t efactura-image -f Dockerfile .
docker run --name efactura -p 9990:8080 efactura-image
docker start efactura