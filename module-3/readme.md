mkdir httpserver && cd httpserver

vim Dockerfile
```
FROM ubuntu
ADD bin/amd64/httpserver /httpserver
EXPOSE 80
ENTRYPOINT /httpserver
```

docker build . -t 121640093/httpserver:v1

docker login -u 121640093 -p 'xxx'

docer push 121640093/httpserver:v1

docker rmi 121640093/httpserver:v1

docker run -itd --name httpserver -p 8080:80 121640093/httpserver:v1

export pid=$(docker inspect -f "{{.State.Pid}}" httpserver)

nsenter -t $pid -n ip a


