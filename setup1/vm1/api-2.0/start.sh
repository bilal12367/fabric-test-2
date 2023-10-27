cd config

./generate-ccp.sh

cd ..

docker rm -f fabric_nodeapp_1

docker image rm testnode1

docker-compose up -d

docker logs --follow fabric_nodeapp_1
