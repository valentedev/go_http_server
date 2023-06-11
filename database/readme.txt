Create image
docker build -t basic-img .

Run the docker container
docker run -d --name basic-ctr -p 5432:5432 basic-img

Open new terminal to access the container
docker exec -it basic-ctr bash

Once inside the container, enter the USER and DATABASE provided on dockerfile
psql -U basicadmin -d basic