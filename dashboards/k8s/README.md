# k8s Metrics
### metric.k8s.container.cpu.limit

Maximum CPU resource limit set for the container.


### metric.k8s.container.cpu.request

CPU resource requested for the container.


### metric.k8s.container.ephemeral_storage.limit

Maximum ephemeral storage resource limit set for the container.


### metric.k8s.container.ephemeral_storage.request

Ephemeral storage resource requested for the container.


### metric.k8s.container.memory.limit

Maximum memory resource limit set for the container.


### metric.k8s.container.memory.request

Memory resource requested for the container.


### metric.k8s.container.ready

Indicates whether the container is currently marked as ready to accept traffic, based on its readiness probe (1 = ready, 0 = not ready).



### metric.k8s.container.restart.count

Describes how many times the container has restarted (since the last counter reset).


### metric.k8s.container.status.reason

Describes the number of K8s containers that are currently in a state for a given reason.


### metric.k8s.container.status.state

Describes the number of K8s containers that are currently in a given state.


### metric.k8s.container.storage.limit

Maximum storage resource limit set for the container.


### metric.k8s.container.storage.request

Storage resource requested for the container.


### metric.k8s.cronjob.active_jobs

The number of actively running jobs for a cronjob.


### metric.k8s.daemonset.current_scheduled_nodes

Number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod.


### metric.k8s.daemonset.desired_scheduled_nodes

Number of nodes that should be running the daemon pod (including nodes currently running the daemon pod).


### metric.k8s.daemonset.misscheduled_nodes

Number of nodes that are running the daemon pod, but are not supposed to run the daemon pod.


### metric.k8s.daemonset.ready_nodes

Number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.


### metric.k8s.deployment.available_pods

Total number of available replica pods (ready for at least minReadySeconds) targeted by this deployment.


### metric.k8s.deployment.desired_pods

Number of desired replica pods in this deployment.


### metric.k8s.hpa.current_pods

Current number of replica pods managed by this horizontal pod autoscaler, as last seen by the autoscaler.


### metric.k8s.hpa.desired_pods

Desired number of replica pods managed by this horizontal pod autoscaler, as last calculated by the autoscaler.


### metric.k8s.hpa.max_pods

The upper limit for the number of replica pods to which the autoscaler can scale up.


### metric.k8s.hpa.metric.target.cpu.average_utilization

Target average utilization, in percentage, for CPU resource in HPA config.


### metric.k8s.hpa.metric.target.cpu.average_value

Target average value for CPU resource in HPA config.


### metric.k8s.hpa.metric.target.cpu.value

Target value for CPU resource in HPA config.


### metric.k8s.hpa.min_pods

The lower limit for the number of replica pods to which the autoscaler can scale down.


### metric.k8s.job.active_pods

The number of pending and actively running pods for a job.


### metric.k8s.job.desired_successful_pods

The desired number of successfully finished pods the job should be run with.


### metric.k8s.job.failed_pods

The number of pods which reached phase Failed for a job.


### metric.k8s.job.max_parallel_pods

The max desired number of pods the job should run at any given time.


### metric.k8s.job.successful_pods

The number of pods which reached phase Succeeded for a job.


### metric.k8s.namespace.phase

Describes number of K8s namespaces that are currently in a given phase.


### metric.k8s.node.allocatable.cpu

Amount of cpu allocatable on the node.


### metric.k8s.node.allocatable.ephemeral_storage

Amount of ephemeral-storage allocatable on the node.


### metric.k8s.node.allocatable.memory

Amount of memory allocatable on the node.


### metric.k8s.node.allocatable.pods

Amount of pods allocatable on the node.


### metric.k8s.node.condition.status

Describes the condition of a particular Node.


### metric.k8s.node.cpu.time

Total CPU time consumed.


### metric.k8s.node.cpu.usage

Node's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs.


### metric.k8s.node.filesystem.available

Node filesystem available bytes.


### metric.k8s.node.filesystem.capacity

Node filesystem capacity.


### metric.k8s.node.filesystem.usage

Node filesystem usage.


### metric.k8s.node.memory.usage

Memory usage of the Node.


### metric.k8s.node.network.errors

Node network errors.


### metric.k8s.node.network.io

Network bytes for the Node.


### metric.k8s.node.uptime

The time the Node has been running.


### metric.k8s.pod.cpu.time

Total CPU time consumed.


### metric.k8s.pod.cpu.usage

Pod's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs.


### metric.k8s.pod.filesystem.available

Pod filesystem available bytes.


### metric.k8s.pod.filesystem.capacity

Pod filesystem capacity.


### metric.k8s.pod.filesystem.usage

Pod filesystem usage.


### metric.k8s.pod.memory.usage

Memory usage of the Pod.


### metric.k8s.pod.network.errors

Pod network errors.


### metric.k8s.pod.network.io

Network bytes for the Pod.


### metric.k8s.pod.uptime

The time the Pod has been running.


### metric.k8s.pod.volume.available

Pod volume storage space available.


### metric.k8s.pod.volume.capacity

Pod volume total capacity.


### metric.k8s.pod.volume.inode.count

The total inodes in the filesystem of the Pod's volume.


### metric.k8s.pod.volume.inode.free

The free inodes in the filesystem of the Pod's volume.


### metric.k8s.pod.volume.inode.used

The inodes used by the filesystem of the Pod's volume.


### metric.k8s.pod.volume.usage

Pod volume usage.


### metric.k8s.replicaset.available_pods

Total number of available replica pods (ready for at least minReadySeconds) targeted by this replicaset.


### metric.k8s.replicaset.desired_pods

Number of desired replica pods in this replicaset.


### metric.k8s.replication_controller.available_pods

Deprecated, use `k8s.replicationcontroller.available_pods` instead.


### metric.k8s.replication_controller.desired_pods

Deprecated, use `k8s.replicationcontroller.desired_pods` instead.


### metric.k8s.replicationcontroller.available_pods

Total number of available replica pods (ready for at least minReadySeconds) targeted by this replication controller.


### metric.k8s.replicationcontroller.desired_pods

Number of desired replica pods in this replication controller.


### metric.k8s.resourcequota.cpu.limit.hard

The CPU limits in a specific namespace.
The value represents the configured quota limit of the resource in the namespace.



### metric.k8s.resourcequota.cpu.limit.used

The CPU limits in a specific namespace.
The value represents the current observed total usage of the resource in the namespace.



### metric.k8s.resourcequota.cpu.request.hard

The CPU requests in a specific namespace.
The value represents the configured quota limit of the resource in the namespace.



### metric.k8s.resourcequota.cpu.request.used

The CPU requests in a specific namespace.
The value represents the current observed total usage of the resource in the namespace.



### metric.k8s.resourcequota.ephemeral_storage.limit.hard

The sum of local ephemeral storage limits in the namespace.
The value represents the configured quota limit of the resource in the namespace.



### metric.k8s.resourcequota.ephemeral_storage.limit.used

The sum of local ephemeral storage limits in the namespace.
The value represents the current observed total usage of the resource in the namespace.



### metric.k8s.resourcequota.ephemeral_storage.request.hard

The sum of local ephemeral storage requests in the namespace.
The value represents the configured quota limit of the resource in the namespace.



### metric.k8s.resourcequota.ephemeral_storage.request.used

The sum of local ephemeral storage requests in the namespace.
The value represents the current observed total usage of the resource in the namespace.



### metric.k8s.resourcequota.hugepage_count.request.hard

The huge page requests in a specific namespace.
The value represents the configured quota limit of the resource in the namespace.



### metric.k8s.resourcequota.hugepage_count.request.used

The huge page requests in a specific namespace.
The value represents the current observed total usage of the resource in the namespace.



### metric.k8s.resourcequota.memory.limit.hard

The memory limits in a specific namespace.
The value represents the configured quota limit of the resource in the namespace.



### metric.k8s.resourcequota.memory.limit.used

The memory limits in a specific namespace.
The value represents the current observed total usage of the resource in the namespace.



### metric.k8s.resourcequota.memory.request.hard

The memory requests in a specific namespace.
The value represents the configured quota limit of the resource in the namespace.



### metric.k8s.resourcequota.memory.request.used

The memory requests in a specific namespace.
The value represents the current observed total usage of the resource in the namespace.



### metric.k8s.resourcequota.object_count.hard

The object count limits in a specific namespace.
The value represents the configured quota limit of the resource in the namespace.



### metric.k8s.resourcequota.object_count.used

The object count limits in a specific namespace.
The value represents the current observed total usage of the resource in the namespace.



### metric.k8s.resourcequota.persistentvolumeclaim_count.hard

The total number of PersistentVolumeClaims that can exist in the namespace.
The value represents the configured quota limit of the resource in the namespace.



### metric.k8s.resourcequota.persistentvolumeclaim_count.used

The total number of PersistentVolumeClaims that can exist in the namespace.
The value represents the current observed total usage of the resource in the namespace.



### metric.k8s.resourcequota.storage.request.hard

The storage requests in a specific namespace.
The value represents the configured quota limit of the resource in the namespace.



### metric.k8s.resourcequota.storage.request.used

The storage requests in a specific namespace.
The value represents the current observed total usage of the resource in the namespace.



### metric.k8s.statefulset.current_pods

The number of replica pods created by the statefulset controller from the statefulset version indicated by currentRevision.


### metric.k8s.statefulset.desired_pods

Number of desired replica pods in this statefulset.


### metric.k8s.statefulset.ready_pods

The number of replica pods created for this statefulset with a Ready Condition.


### metric.k8s.statefulset.updated_pods

Number of replica pods created by the statefulset controller from the statefulset version indicated by updateRevision.

