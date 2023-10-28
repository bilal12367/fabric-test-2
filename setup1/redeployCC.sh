
cd vm1

./deployChaincode.sh start

cd ../vm2

./installAndApprove.sh

cd ../vm3

./installAndApprove.sh

cd ../vm1

./deployChaincode.sh init