#!/bin/bash

# taken from: https://community.neo4j.com/t/copy-dump-file-to-docker-container-and-load-it-on-startup/36816/2

# THANK YOU! Special shout-out to @marcellodesales on GitHub
# https://github.com/marcellodesales/neo4j-with-cypher-seed-docker/blob/master/wrapper.sh for such a great example script

# Log the info with the same format as NEO4J outputs
log_info() {
  # https://www.howtogeek.com/410442/how-to-display-the-date-and-time-in-the-linux-terminal-and-use-it-in-bash-scripts/
  printf '%s %s\n' "$(date -u +"%Y-%m-%d %H:%M:%S:%3N%z") INFO  MGT: $1"
  return
}

# turn on bash's job control
# https://stackoverflow.com/questions/11821378/what-does-bashno-job-control-in-this-shell-mean/46829294#46829294
set -m

log_info "Import database dump"
neo4j-admin load --from=./db.dump --database=neo4j --force
log_info "DONE"

/docker-entrypoint.sh neo4j