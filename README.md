# snyk-without-build-test
Testing running snyk commands without building a go project

This simple example demonstrates that snyk will successfully `test` and `monitor` code when running on go 1.15 without running a build. 

Snyk will handle the fact that there are dependencies missing because we have not generated the files using swagger.
