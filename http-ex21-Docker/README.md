Ref:  
 https://medium.com/@shijuvar/deploying-go-web-apps-with-docker-1b7561b36f53  
 https://blog.hasura.io/the-ultimate-guide-to-writing-dockerfiles-for-go-web-apps-336efad7012c/


```
docker build -t go-web-docker .
docker run --rm -it -p 8080:8080 go-web-docker
```
