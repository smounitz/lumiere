FROM neo4j:4.4.3-community

COPY --chown=neo4j dumps/db.dump db.dump
COPY --chown=neo4j load-entrypoint.sh load-entrypoint.sh
RUN chmod +x load-entrypoint.sh

ENTRYPOINT ["./load-entrypoint.sh"]