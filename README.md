# cassandra-go

# Install & run cassnafra using docker
`docker pull cassandra`
`docker run -d --name cassandra -p 9042:9042 cassandra`
# docker-compose
- 7199 - JMX (was 8080 pre Cassandra 0.8.xx)
- 7000 - Internode comminication (not use if TLS enabled)
- 7001 - TLS Internode communcation (use if TLS enabled)
- 9160 - Thrift client API
- 9042 - CQL native transport port
# Importance commands
```
docker exec -it cassandra cqlsh

CREATE KEYSPACE todo WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};

use todo;

create table todo (id text primary key, title text, content text);

select * from todo
```
# Connect with datagrip (ver. 1.4)
```
DBMS: Cassandra (ver. 5.0.3)
Case sensitivity: plain=lower, delimited=exact
Driver: Cassandra JDBC Driver (ver. 1.4, JDBC4.2)

Ping: 35 ms
```

# Cassandra JDBC Driver 4.12.0
![alt](resources/jdbc4.12.0.png)

[cassandra db connection errors](https://youtrack.jetbrains.com/issue/PY-73766/Apache-Cassandra-connection-error-Docker-DB-instance)

