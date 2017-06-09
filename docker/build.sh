#!/bin/bash

if [ "$1" == "clean" ]; then docker rmi snoby/spark-pivot:build; fi

docker build -t snoby/spark-pivot:BUILD . -f Dockerfile.build

#
#
#Next step is to extract the files that you built out of the container.
docker create --name built snoby/spark-pivot:BUILD
docker cp built:/output/webhook ./webhook
docker cp built:/output/spark-pivot ./spark-pivot
docker rm -f built

# Now add these new files not a container
docker build --no-cache -t snoby/spark-pivot:latest .
rm -rf spark-pivot
rm -rf webhook
echo "Build Complete"
