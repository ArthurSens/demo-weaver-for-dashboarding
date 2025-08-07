# vcs Metrics
### metric.vcs.change.count

The number of changes (pull requests/merge requests/changelists) in a repository, categorized by their state (e.g. open or merged).


### metric.vcs.change.duration

The time duration a change (pull request/merge request/changelist) has been in a given state.


### metric.vcs.change.time_to_approval

The amount of time since its creation it took a change (pull request/merge request/changelist) to get the first approval.


### metric.vcs.change.time_to_merge

The amount of time since its creation it took a change (pull request/merge request/changelist) to get merged into the target(base) ref.


### metric.vcs.contributor.count

The number of unique contributors to a repository.


### metric.vcs.ref.count

The number of refs of type branch or tag in a repository.


### metric.vcs.ref.lines_delta

The number of lines added/removed in a ref (branch) relative to the ref from the `vcs.ref.base.name` attribute.


### metric.vcs.ref.revisions_delta

The number of revisions (commits) a ref (branch) is ahead/behind the branch from the `vcs.ref.base.name` attribute.


### metric.vcs.ref.time

Time a ref (branch) created from the default branch (trunk) has existed. The `ref.type` attribute will always be `branch`.


### metric.vcs.repository.count

The number of repositories in an organization.

