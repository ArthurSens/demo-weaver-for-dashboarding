# db Metrics
### metric.db.client.connection.count

The number of connections that are currently in state described by the `state` attribute.


### metric.db.client.connection.create_time

The time it took to create a new connection.


### metric.db.client.connection.idle.max

The maximum number of idle open connections allowed.


### metric.db.client.connection.idle.min

The minimum number of idle open connections allowed.


### metric.db.client.connection.max

The maximum number of open connections allowed.


### metric.db.client.connection.pending_requests

The number of current pending requests for an open connection.


### metric.db.client.connection.timeouts

The number of connection timeouts that have occurred trying to obtain a connection from the pool.


### metric.db.client.connection.use_time

The time between borrowing a connection and returning it to the pool.


### metric.db.client.connection.wait_time

The time it took to obtain an open connection from the pool.


### metric.db.client.connections.create_time

Deprecated, use `db.client.connection.create_time` instead. Note: the unit also changed from `ms` to `s`.


### metric.db.client.connections.idle.max

Deprecated, use `db.client.connection.idle.max` instead.


### metric.db.client.connections.idle.min

Deprecated, use `db.client.connection.idle.min` instead.


### metric.db.client.connections.max

Deprecated, use `db.client.connection.max` instead.


### metric.db.client.connections.pending_requests

Deprecated, use `db.client.connection.pending_requests` instead.


### metric.db.client.connections.timeouts

Deprecated, use `db.client.connection.timeouts` instead.


### metric.db.client.connections.usage

Deprecated, use `db.client.connection.count` instead.


### metric.db.client.connections.use_time

Deprecated, use `db.client.connection.use_time` instead. Note: the unit also changed from `ms` to `s`.


### metric.db.client.connections.wait_time

Deprecated, use `db.client.connection.wait_time` instead. Note: the unit also changed from `ms` to `s`.


### metric.db.client.cosmosdb.active_instance.count

Deprecated, use `azure.cosmosdb.client.active_instance.count` instead.


### metric.db.client.cosmosdb.operation.request_charge

Deprecated, use `azure.cosmosdb.client.operation.request_charge` instead.


### metric.db.client.operation.duration

Duration of database client operations.


### metric.db.client.response.returned_rows

The actual number of records returned by the database operation.

