#!/usr/bin/env sh
cd backend
docker build -t fuse-go-image .
cd ..
echo "Service running at http://localhost:3000"
docker run -it --rm -d --name tmp -p 3000:3000 fuse-go-image:latest

cd frontend
docker build -t fuse-svelte-image .
cd ..
echo "Frontend running at http://localhost:8080"
docker run -it --rm -d --name tmp1 -p 80:80 fuse-svelte-image:latest


