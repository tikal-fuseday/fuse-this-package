#!/usr/bin/env sh
cd backend
docker build -t fuse_go_image .
cd ..
echo "Service running at http://localhost:3000"
docker run -it --rm --name tmp -p 3000:3000 fuse_go_image:latest


