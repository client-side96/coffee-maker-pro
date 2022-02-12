# coffee-maker-pro

## Development setup

1. Create the log directory for the sensor logs
```shell
sudo mkdir -p /var/log/coffee-maker-pro
sudo chmod 777 /var/log/coffee-maker-pro
```

2. Setup the local `mongodb instances`
```shell
docker-compose up -d
docker ps # copy the container id of the instance running on port 27011
docker exec -it <container_id> /bin/bash

# inside the container
mongo
rs.initiate({_id : "rs0",members: [{"_id": 0,"host": "localmongo1:27017","priority": 4},{"_id": 1,"host": "localmongo2:27017","priority": 2},{"_id": 2,"host": "localmongo3:27017","priority": 1}]})
```
