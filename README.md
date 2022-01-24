# lumiere
Lumiere (light in French and a cute candelabra in Beauty and The Beast) illuminates users about their favorite movies and actors. Naming  service is very important and in a containerized environment the earlier the better. 
Lumiere comes with a dockerized Neo4j service that loads the database on the first run.
Lumiere uses the example movie neo4j database and adds to it:
* user and annotation node
* user added annotation relation
* user follow user relation
* user likes user movie person relation
* (So far only the user node was added)


Lumiere will expose a graphql endpoint (not implemented yet) to query movies and their annotations. We want to show movies relevant to the user, so querying will only be allowed for logged in users. The movies and annotations will be ordered based on the relations between the loged in user and other users in the database.

Graphql is great, but for things like health checks, authentication, and social network sign in, a simple Rest endpoint is better.  So we added a few Rest endpoint. But inspired by Graphql mutations we chose to name the endpoint with verbs (register, signing, etc) and not http verb and noun.

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
