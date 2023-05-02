### Prerequisites
Install Docker Desktop: [link](https://docs.docker.com/compose/install/)
This will also include Docker Compose and all that is required to run the application.
### Usage
The whole application is Dockerized and also there are some pre-made container build scripts that can be run and will start the application.

 1. clone the repository
 2. cd into the cloned folder structure
 3. enter `source dc_env` in the terminal
 4. to run the application in Docker enter `dc_upd` in the terminal
 5. after the container build phase, the backend application takes around 30 sec to start running (this is known issue, the db service is slow to get running and the golang app needs to wait for it)
 6. after everything is up and running (check in Docker desktop or use commands) the application should be reachable at http://localhost:3000
 7. to stop the application from running enter `dc_down` 
