# lumiere
Lumiere (light in French and a cute candelabra in Beauty and The Beast) illuminates users about their favorite movies and actors. Naming  service is very important and in a containerized environment the earlier the better. 
Lumiere comes with a dockerized Neo4j service that loads the database on the first run.
Lumiere uses the example movie neo4j database and adds to it:
* user and annotation node
* user added annotation relation
* user follow user relation
* user likes user movie person relation
* (So far only the user node was added)



  


### Buid Neo4j
```
docker-compose build neo4j
```
### Run Neo4j
```
docker-compose up neo4j
```
### Run Api
(Dockerfile to come)
```
go build .
go run api
```
