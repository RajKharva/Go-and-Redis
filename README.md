# Go-and-Redis
Communicate with Redis (containerized using Docker) from Go


Below are the steps that I took for resolving the task:
1. Installed Docker

2. Ran below command to to start a new container with name “redisTest” 
-d - detaches the container and runs it in background
-p - publishes the container’s port to the host
docker run -d -p 6379:6379 --name redisTest redis


3. Check whether the redisTest container is running
docker ps
ps - Process Status

￼

4. Open interactive shell to interact with Redis running on redisTest container
docker exec -it redisTest sh

Verify whether Redis is running and verifying value of key being set from RediGo.go file
I have used RediGo package for communicating with Redis from Go
	The program connects to Redis, 
	creates a counter key in Redis with a value of 1, 
	reads and prints it, increments it, 
	reads and prints it
￼

5. Exit the container and clean up
