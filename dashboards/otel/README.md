# otel Metrics
### metric.otel.sdk.exporter.log.exported

The number of log records for which the export has finished, either successful or failed.


### metric.otel.sdk.exporter.log.inflight

The number of log records which were passed to the exporter, but that have not been exported yet (neither successful, nor failed).


### metric.otel.sdk.exporter.metric_data_point.exported

The number of metric data points for which the export has finished, either successful or failed.


### metric.otel.sdk.exporter.metric_data_point.inflight

The number of metric data points which were passed to the exporter, but that have not been exported yet (neither successful, nor failed).


### metric.otel.sdk.exporter.operation.duration

The duration of exporting a batch of telemetry records.


### metric.otel.sdk.exporter.span.exported

The number of spans for which the export has finished, either successful or failed.


### metric.otel.sdk.exporter.span.exported.count

Deprecated, use `otel.sdk.exporter.span.exported` instead.


### metric.otel.sdk.exporter.span.inflight

The number of spans which were passed to the exporter, but that have not been exported yet (neither successful, nor failed).


### metric.otel.sdk.exporter.span.inflight.count

Deprecated, use `otel.sdk.exporter.span.inflight` instead.


### metric.otel.sdk.log.created

The number of logs submitted to enabled SDK Loggers.


### metric.otel.sdk.metric_reader.collection.duration

The duration of the collect operation of the metric reader.


### metric.otel.sdk.processor.log.processed

The number of log records for which the processing has finished, either successful or failed.


### metric.otel.sdk.processor.log.queue.capacity

The maximum number of log records the queue of a given instance of an SDK Log Record processor can hold.


### metric.otel.sdk.processor.log.queue.size

The number of log records in the queue of a given instance of an SDK log processor.


### metric.otel.sdk.processor.span.processed

The number of spans for which the processing has finished, either successful or failed.


### metric.otel.sdk.processor.span.processed.count

Deprecated, use `otel.sdk.processor.span.processed` instead.


### metric.otel.sdk.processor.span.queue.capacity

The maximum number of spans the queue of a given instance of an SDK span processor can hold.


### metric.otel.sdk.processor.span.queue.size

The number of spans in the queue of a given instance of an SDK span processor.


### metric.otel.sdk.span.ended

Use `otel.sdk.span.started` minus `otel.sdk.span.live` to derive this value.


### metric.otel.sdk.span.ended.count

Use `otel.sdk.span.started` minus `otel.sdk.span.live` to derive this value.


### metric.otel.sdk.span.live

The number of created spans with `recording=true` for which the end operation has not been called yet.


### metric.otel.sdk.span.live.count

Deprecated, use `otel.sdk.span.live` instead.


### metric.otel.sdk.span.started

The number of created spans.

