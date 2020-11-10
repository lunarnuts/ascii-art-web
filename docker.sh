#!/bin/bash
docker build --tag ascii-art-web .
docker run --publish 8080:8080 --detach --name ascii ascii-art-web 
