# exam_shopee
exam

# how to build image
docker build -t getip exam_shopee/.

# how to create the container and start it.
docker container create --name getip -p 80:8080 getip && docker start getip

