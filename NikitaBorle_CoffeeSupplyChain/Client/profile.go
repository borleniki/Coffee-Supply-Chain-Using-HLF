package main


type Config struct {
	CertPath     string `json:"certPath"`
	KeyDirectory string `json:"keyPath"`
	TLSCertPath  string `json:"tlsCertPath"`
	PeerEndpoint string `json:"peerEndpoint"`
	GatewayPeer  string `json:"gatewayPeer"`
	MSPID        string `json:"mspID"`
}


var profile = map[string]Config{

	"farmers": {
		CertPath:     "../Fabric_Network/organizations/peerOrganizations/farmers.auto.com/users/User1@farmers.auto.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Fabric_Network/organizations/peerOrganizations/farmers.auto.com/users/User1@farmers.auto.com/msp/keystore/",
		TLSCertPath:  "../Fabric_Network/organizations/peerOrganizations/farmers.auto.com/peers/peer0.farmers.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:7051",
		GatewayPeer:  "peer0.farmers.auto.com",
		MSPID:        "FarmersMSP",
	},

	"processors": {
		CertPath:     "../Fabric_Network/organizations/peerOrganizations/processors.auto.com/users/User1@processors.auto.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Fabric_Network/organizations/peerOrganizations/processors.auto.com/users/User1@processors.auto.com/msp/keystore/",
		TLSCertPath:  "../Fabric_Network/organizations/peerOrganizations/processors.auto.com/peers/peer0.processors.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:9051",
		GatewayPeer:  "peer0.processors.auto.com",
		MSPID:        "ProcessorsMSP",
	},

	"distributors": {
		CertPath:     "../Fabric_Network/organizations/peerOrganizations/distributors.auto.com/users/User1@distributors.auto.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Fabric_Network/organizations/peerOrganizations/distributors.auto.com/users/User1@distributors.auto.com/msp/keystore/",
		TLSCertPath:  "../Fabric_Network/organizations/peerOrganizations/distributors.auto.com/peers/peer0.distributors.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:11051",
		GatewayPeer:  "peer0.distributors.auto.com",
		MSPID:        "DistributorsMSP",
	},

}