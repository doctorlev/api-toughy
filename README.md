# api-toughy
some microservices API, DB, docker

The Project contains golang SW with the following components

1. **  HTTPAPI (runs on container) -
the program to main.go only imports local HTTPMANAGE  package,
which provides http request analysis and response creation based on:
- analyzing the uri (..../<string>) and the method (POST, GET, etc)
- choosing accordingly the operation about the User with the DB (redis)
the HTTPMANAGE imports local DATAMANGE(to work with Redis DB) and
HTTPHELPER (to run TokenCheck)
To run on container - use run.sh (and env.sh)

2. **  STORAGEAPI (runs on container)  -
the program main.go only imports local STORAGEMANAGE package,
which provides the functionality of :
- uploading and downloading the files
to/from DB
- getting the combined information from DB.
It imports local HTTPHELPER (to run TokenCheck).
To run on container - use run.sh (and env.sh)

3. Project packages:
- DATAMANAGE (to work with Redis DB)
- HTTPHELPER (to run TokenCheck)

4. Redis DB (runs on standard redis container)  -
- up.sh- start the container
- down.sh - stop,kill and remove the container

5. TESTS
contains:
- scripts to run tests, including #5 - consolidated
- examples of CURL command
- ufile.txt file for uploading

