#/bin/bash
./network.sh down
./network.sh up createChannel -c blockchain2023 -ca 
./network.sh deployCC -c blockchain2023 -ccn test -ccl go -ccv 1.0 -ccs 1 -ccp ../chaincode -cci init
node enrollAdmin.js
node registerUser.js