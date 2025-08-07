# dotnet Metrics
### metric.dotnet.assembly.count

The number of .NET assemblies that are currently loaded.


### metric.dotnet.exceptions

The number of exceptions that have been thrown in managed code.


### metric.dotnet.gc.collections

The number of garbage collections that have occurred since the process has started.


### metric.dotnet.gc.heap.total_allocated

The *approximate* number of bytes allocated on the managed GC heap since the process has started. The returned value does not include any native allocations.



### metric.dotnet.gc.last_collection.heap.fragmentation.size

The heap fragmentation, as observed during the latest garbage collection.



### metric.dotnet.gc.last_collection.heap.size

The managed GC heap size (including fragmentation), as observed during the latest garbage collection.



### metric.dotnet.gc.last_collection.memory.committed_size

The amount of committed virtual memory in use by the .NET GC, as observed during the latest garbage collection.



### metric.dotnet.gc.pause.time

The total amount of time paused in GC since the process has started.


### metric.dotnet.jit.compilation.time

The amount of time the JIT compiler has spent compiling methods since the process has started.



### metric.dotnet.jit.compiled_il.size

Count of bytes of intermediate language that have been compiled since the process has started.


### metric.dotnet.jit.compiled_methods

The number of times the JIT compiler (re)compiled methods since the process has started.



### metric.dotnet.monitor.lock_contentions

The number of times there was contention when trying to acquire a monitor lock since the process has started.



### metric.dotnet.process.cpu.count

The number of processors available to the process.


### metric.dotnet.process.cpu.time

CPU time used by the process.


### metric.dotnet.process.memory.working_set

The number of bytes of physical memory mapped to the process context.


### metric.dotnet.thread_pool.queue.length

The number of work items that are currently queued to be processed by the thread pool.



### metric.dotnet.thread_pool.thread.count

The number of thread pool threads that currently exist.


### metric.dotnet.thread_pool.work_item.count

The number of work items that the thread pool has completed since the process has started.



### metric.dotnet.timer.count

The number of timer instances that are currently active.

