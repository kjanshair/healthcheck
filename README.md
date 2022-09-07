A simple healthcheck that returns HTTP 200 and prints access logs to stdout when request is made to the endpoint.

Pass **ENDPOINT** & **PORT** to the application or container to run healthcheck i.e. `docker container run --rm -d -p 80:80 --name healthcheck -e ENDPOINT=/health -e PORT=80 kjanshair/healthcheck`. 
