#! /bin/sh
docker build -t test-web-app .
docker run -it \
--network="host" \
-p 127.0.0.1:8000:8000 -t test-web-app
