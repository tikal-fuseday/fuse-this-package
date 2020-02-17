#!/usr/bin/env sh
pushd backend
docker build -t fuse_go_image .
popd
echo "Service running at http://localhost:3000"
docker run -it --rm --name tmp -p 3000:3000 fuse_go_image:latest


