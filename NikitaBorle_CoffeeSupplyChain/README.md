### Start your network using 
```
./startNetwork.sh  
```

### To invoke 
```
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.auto.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n Coffeesupply --peerAddresses localhost:7051 --tlsRootCertFiles $FARMERS_PEER_TLSROOTCERT --peerAddresses localhost:9051 --tlsRootCertFiles $PROCESSORS_PEER_TLSROOTCERT --peerAddresses localhost:11051 --tlsRootCertFiles $DISTRIBUTORS_PEER_TLSROOTCERT -c '{"function":"RegisterFarmer","Args":["Nikita", "01", "06/03/2002", "Female", "nikku@gmail.com", "1234567891", "Hyderabad", "Arabica", "Brown Black", "InFarm"]}'
```

### To Query
```
peer chaincode query -C $CHANNEL_NAME -n Coffeesupply -c '{"Args":["GetAllFarmers"]}'
```

### To get history

```
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.auto.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n Coffeesupply --peerAddresses localhost:7051 --tlsRootCertFiles $FARMERS_PEER_TLSROOTCERT --peerAddresses localhost:9051 --tlsRootCertFiles $PROCESSORS_PEER_TLSROOTCERT --peerAddresses localhost:11051 --tlsRootCertFiles $DISTRIBUTORS_PEER_TLSROOTCERT -c '{"function":"RegisterFarmer","Args":["Amjad", "02", "22/05/2002", "Male", "amjad@gmail.com", "1234567891", "Hyderabad", "Arabica", "Brown", "Processed"]}'
```
```
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.auto.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n Coffeesupply --peerAddresses localhost:7051 --tlsRootCertFiles $FARMERS_PEER_TLSROOTCERT --peerAddresses localhost:9051 --tlsRootCertFiles $PROCESSORS_PEER_TLSROOTCERT --peerAddresses localhost:11051 --tlsRootCertFiles $DISTRIBUTORS_PEER_TLSROOTCERT -c '{"function":"RegisterFarmer","Args":["Ankita", "03", "30/03/2002", "Female", "ankita@gmail.com", "1234567891", "Hyderabad", "Arabica", "Black", "InFarm"]}'
```
```
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.auto.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n Coffeesupply --peerAddresses localhost:7051 --tlsRootCertFiles $FARMERS_PEER_TLSROOTCERT --peerAddresses localhost:9051 --tlsRootCertFiles $PROCESSORS_PEER_TLSROOTCERT --peerAddresses localhost:11051 --tlsRootCertFiles $DISTRIBUTORS_PEER_TLSROOTCERT -c '{"function":"RegisterFarmer","Args":["Kavi", "04", "11/08/2002", "Male", "kavi@gmail.com", "1234567891", "Hyderabad", "Arabica", "Brown", "Completes"]}'
```

```
peer chaincode query -C $CHANNEL_NAME -n Coffeesupply -c '{"Args":["GetAllFarmers"]}'
```

### To run PDC 

```
export NAME=$(echo -n "Nikhil" | base64 | tr -d \\n)

export QUALITY=$(echo -n "2" | base64 | tr -d \\n)

export COLOR=$(echo -n "Brown Black" | base64 | tr -d \\n)

export TYPEOFBEANS=$(echo -n "Arabica" | base64 | tr -d \\n)

export GRINDSIZE=$(echo -n "Small" | base64 | tr -d \\n)

export BREWTIME=$(echo -n "15 min" | base64 | tr -d \\n)
```

```
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.auto.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n Coffeesupply --peerAddresses localhost:7051 --tlsRootCertFiles $FARMERS_PEER_TLSROOTCERT --peerAddresses localhost:9051 --tlsRootCertFiles $PROCESSORS_PEER_TLSROOTCERT --peerAddresses localhost:11051 --tlsRootCertFiles $DISTRIBUTORS_PEER_TLSROOTCERT -c '{"Args":["ProcessContract:CreateProcessOrder","01"]}' --transient "{\"name\":\"$NAME\",\"quality\":\"$QUALITY\",\"color\":\"$COLOR\",\"typeofBeans\":\"$TYPEOFBEANS\",\"grindSize\":\"$GRINDSIZE\", \"brewTime\":\"$BREWTIME\"}"
```

```
peer chaincode query -C $CHANNEL_NAME -n Coffeesupply -c '{"Args":["ProcessContract:ReadProcessOrder","01"]}'
```


### To run using API 

To post the farmer
```
http://localhost:3000/api/farmer
```
```
{
    "userName": "anki",
    "name": "Ankita",
    "address": "Hyderabad",
    "dob": "22/12/2012",
    "gender": "female",
    "contactNo": "9090904545",
    "emailId": "anku@gmail.com",
    "typeofBenas": "Arabica",
    "color": "Brown",
    "status": "In farm"
}
```

To get the farmer details
```
http://localhost:3000/api/farmer/anki
```

To post the Privatedate
```
http://localhost:3000/api/order
```
```
{
  "processId": "1",
  "quality": "High",
  "color": "Brown",
  "name": "Nikhil",
  "typeofBeans": "Arabica",
  "grindSize": "Medium",
  "brewTime": "5 minutes"
}

```

To get Private data
```
http://localhost:3000/api/order/1
```
