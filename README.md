# REST API for TreatField app

- Docker compose for TIG and Golang simple app:

https://github.com/tochytskyi/treatfield-api/blob/main/docker-compose.yml

- Grafana dashboard with influxdb data source:

![image](https://user-images.githubusercontent.com/7937891/139577919-b4fcc6fe-8af1-4e9b-8871-66f88e6fc5e0.png)

- AB query: Golang http request + Mysql query
```
ab -t 60 -c 100 -n 2000 http://127.0.0.1:5000/api/v1/users/first
```

- Mysql CPU spikes after load tests

![image](https://user-images.githubusercontent.com/7937891/139578218-ca1c823a-9add-4173-bd2c-c051b9eace4a.png)

- API CPU spikes after load tests

![image](https://user-images.githubusercontent.com/7937891/139578233-d6100360-4338-49d0-9eb6-fa8dd0489d80.png)

- Memory

![image](https://user-images.githubusercontent.com/7937891/139578268-5a5eb0f1-e78e-4419-b572-b9e471731f76.png)

- Network

![image](https://user-images.githubusercontent.com/7937891/139578283-b5ebb158-462a-4717-a9ea-e04a2c0a99fd.png)







