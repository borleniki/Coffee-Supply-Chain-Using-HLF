package contracts

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ProcessContract struct {
	contractapi.Contract
}

type Process struct {
	ProcessId     string `json:"processId"`
	CoffeeQuality string `json:"quality"`
	Color         string `json:"color"`
	Name          string `json:"name"`
	TypeofBeans   string `json:"typeofBeans"`
	GrindSize     string `json:"grindSize"`
	BrewTime      string `json:"brewTime"`
}

const collectionName string = "ProcessCollection"

func (p *ProcessContract) ProcessOrderExists(ctx contractapi.TransactionContextInterface, processId string) (bool, error) {

	data, err := ctx.GetStub().GetPrivateDataHash(collectionName, processId)

	if err != nil {
		return false, fmt.Errorf("could not fetch the private data hash. %s", err)
	}

	return data != nil, nil
}

func (p *ProcessContract) CreateProcessOrder(ctx contractapi.TransactionContextInterface, processId string) (string, error) {

	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not fetch client identity. %s", err)
	}

	if clientOrgID == "ProcessorsMSP" {
		exists, err := p.ProcessOrderExists(ctx, processId)
		if err != nil {
			return "", fmt.Errorf("could not read from world state. %s", err)
		} else if exists {
			return "", fmt.Errorf("the order %s already exists", processId)
		}

		var process Process

		transientData, err := ctx.GetStub().GetTransient()
		if err != nil {
			return "", fmt.Errorf("could not fetch transient data. %s", err)
		}

		if len(transientData) == 0 {
			return "", fmt.Errorf("please provide the private data of name, quality, typeofbeans, grindsize, brewtime")
		}

		name, exists := transientData["name"]
		if !exists {
			return "", fmt.Errorf("the name was not specified in transient data. Please try again")
		}
		process.Name = string(name)

		quality, exists := transientData["quality"]
		if !exists {
			return "", fmt.Errorf("the coffequality was not specified in transient data. Please try again")
		}
		process.CoffeeQuality = string(quality)

		color, exists := transientData["color"]
		if !exists {
			return "", fmt.Errorf("the color was not specified in transient data. Please try again")
		}
		process.Color = string(color)

		typeofBeans, exists := transientData["typeofBeans"]
		if !exists {
			return "", fmt.Errorf("the typeofBeans was not specified in transient data. Please try again")
		}
		process.TypeofBeans = string(typeofBeans)

		grindSize, exists := transientData["grindSize"]
		if !exists {
			return "", fmt.Errorf("the grindSize was not specified in transient data. Please try again")
		}
		process.GrindSize = string(grindSize)

		brewTime, exists := transientData["brewTime"]
		if !exists {
			return "", fmt.Errorf("the brewTime was not specified in transient data. Please try again")
		}
		process.BrewTime = string(brewTime)

		process.ProcessId = processId

		bytes, _ := json.Marshal(process)
		err = ctx.GetStub().PutPrivateData(collectionName, processId, bytes)
		if err != nil {
			return "", fmt.Errorf("could not able to write the data")
		}
		return fmt.Sprintf("order with id %v added successfully", processId), nil
	} else {
		return fmt.Sprintf("order cannot be created by organisation with MSPID %v ", clientOrgID), nil
	}
}

func (p *ProcessContract) ReadProcessOrder(ctx contractapi.TransactionContextInterface, processId string) (*Process, error) {
	exists, err := p.ProcessOrderExists(ctx, processId)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", processId)
	}

	bytes, err := ctx.GetStub().GetPrivateData(collectionName, processId)
	if err != nil {
		return nil, fmt.Errorf("could not get the private data. %s", err)
	}
	var process Process

	err = json.Unmarshal(bytes, &process)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal private data collection data to type Order")
	}

	return &process, nil

}


func (p *ProcessContract) DeleteProcessOrder(ctx contractapi.TransactionContextInterface, processId string) error {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("could not read the client identity. %s", err)
	}

	if clientOrgID == "ProcessorsMSP" {

		exists, err := p.ProcessOrderExists(ctx, processId)

		if err != nil {
			return fmt.Errorf("could not read from world state. %s", err)
		} else if !exists {
			return fmt.Errorf("the asset %s does not exist", processId)
		}

		return ctx.GetStub().DelPrivateData(collectionName, processId)
	} else {
		return fmt.Errorf("organisation with %v cannot delete the order", clientOrgID)
	}
}


func (p *ProcessContract) GetAllOrders(ctx contractapi.TransactionContextInterface) ([]*Process, error) {
	queryString := `{"selector":{"assetType":"Process"}}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(collectionName, queryString)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the query result. %s", err)
	}
	defer resultsIterator.Close()
	return OrderResultIteratorFunction(resultsIterator)
}


func (p *ProcessContract) GetOrdersByRange(ctx contractapi.TransactionContextInterface, startKey string, endKey string) ([]*Process, error) {
	resultsIterator, err := ctx.GetStub().GetPrivateDataByRange(collectionName, startKey, endKey)

	if err != nil {
		return nil, fmt.Errorf("could not fetch the private data by range. %s", err)
	}
	defer resultsIterator.Close()

	return OrderResultIteratorFunction(resultsIterator)

}


func OrderResultIteratorFunction(resultsIterator shim.StateQueryIteratorInterface) ([]*Process, error) {
	var processors []*Process
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not fetch the details of result iterator. %s", err)
		}
		var process Process
		err = json.Unmarshal(queryResult.Value, &process)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal the data. %s", err)
		}
		processors = append(processors, &process)
	}

	return processors, nil
}
