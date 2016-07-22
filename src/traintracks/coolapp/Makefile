# Import our abstractions 
include ../../../makefile.defs

name=coolapp

# The root of the docker context
docker_context=..

# Path to all dependencies seperated by space
dependencies=${docker_context}/lib1 ${docker_context}/lib2 ${docker_context}/coolapp

# Where the context goes in the container 
container_src=/go/src/traintracks

# Where artifacts end up after build
container_artifacts=/go/bin

# Complete build command
build_command=go get ./...
