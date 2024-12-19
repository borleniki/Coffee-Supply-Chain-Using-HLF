package contracts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type FarmerContract struct {
	contractapi.Contract
}

type PaginatedQueryResult struct {
	Records             []*Farmer `json:"records"`
	FetchedRecordsCount int32     `json:"fetchedRecordsCount"`
	Bookmark            string    `json:"bookmark"`
}

type HistoryQueryResult struct {
	Record    *Farmer `json:"record"`
	TxId      string  `json:"txId"`
	Timestamp string  `json:"timestamp"`
	IsDelete  bool    `json:"isDelete"`
}

type Farmer struct {
	Address     string `json:"address"`
	Color       string `json:"color"`
	ContactNo   string `json:"contactNo"`
	DateOfBirth string `json:"dob"`
	EmailId     string `json:"emailId"`
	Gender      string `json:"gender"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	TypeofBeans string `json:"typeofBeans"`
	UserName    string `json:"userName"`
}

// Chech if farmer is already registered
func (f *FarmerContract) FarmerExists(ctx contractapi.TransactionContextInterface, userName string) (bool, error) {
	data, err := ctx.GetStub().GetState(userName)

	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)

	}
	return data != nil, nil
}

func (f *FarmerContract) RegisterFarmer(ctx contractapi.TransactionContextInterface, userName string, name string, dob string, gender string, emailId string, contactNo string, address string, typeofBeans string, color string, status string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID == "FarmersMSP" {
		exists, err := f.FarmerExists(ctx, userName)
		if err != nil {
			return "", fmt.Errorf("%s", err)
		} else if exists {
			return "", fmt.Errorf("the farmer with username %s already present", userName)
		}

		farmer := Farmer{
			Address:     address,
			Color:       color,
			ContactNo:   contactNo,
			DateOfBirth: dob,
			EmailId:     emailId,
			Gender:      gender,
			Name:        name,
			Status:      status,
			TypeofBeans: typeofBeans,
			UserName:    userName,
		}

		bytes, _ := json.Marshal(farmer)

		err = ctx.GetStub().PutState(userName, bytes)

		if err != nil {
			return "", err
		} else {
			return fmt.Sprintf("successfully added farmer with username %v", userName), nil
		}

	} else {
		return "", fmt.Errorf("user under this MSPID %v can't perform this action", clientOrgID)
	}
}

func (f *FarmerContract) ReadFarmerDetails(ctx contractapi.TransactionContextInterface, userName string) (*Farmer, error) {

	bytes, err := ctx.GetStub().GetState(userName)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if bytes == nil {
		return nil, fmt.Errorf("the farmer %s does not exist", userName)
	}

	var farmer Farmer

	err = json.Unmarshal(bytes, &farmer)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type farmer")
	}

	return &farmer, nil
}

func (f *FarmerContract) DeleteFarmer(ctx contractapi.TransactionContextInterface, userName string) (string, error) {

	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not fetch client identity. %s", err)
	}

	if clientOrgID == "FarmersMSP" {

		exists, err := f.FarmerExists(ctx, userName)
		if err != nil {
			return "", fmt.Errorf("%s", err)
		} else if !exists {
			return "", fmt.Errorf("the farmer, %s does not exist", userName)
		}

		err = ctx.GetStub().DelState(userName)
		if err != nil {
			return "", err
		} else {
			return fmt.Sprintf("farmer with id %v is deleted from the world state.", userName), nil
		}

	} else {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}
}

func (f *FarmerContract) GetFarmerByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*Farmer, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the  data by range. %s", err)
	}
	defer resultsIterator.Close()

	return farmerResultIteratorFunction(resultsIterator)
}

func (c *FarmerContract) GetAllFarmers(ctx contractapi.TransactionContextInterface) ([]*Farmer, error) {

	queryString := `{"selector":{}}`

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the query result. %s", err)
	}
	defer resultsIterator.Close()
	return farmerResultIteratorFunction(resultsIterator)
}

func farmerResultIteratorFunction(resultsIterator shim.StateQueryIteratorInterface) ([]*Farmer, error) {
	var farmers []*Farmer
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not fetch the details of the result iterator. %s", err)
		}
		var farmer Farmer
		err = json.Unmarshal(queryResult.Value, &farmer)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal the data. %s", err)
		}
		farmers = append(farmers, &farmer)
	}

	return farmers, nil
}

func (f *FarmerContract) GetFarmerHistory(ctx contractapi.TransactionContextInterface, userName string) ([]*HistoryQueryResult, error) {

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(userName)
	if err != nil {
		return nil, fmt.Errorf("could not get the data. %s", err)
	}
	defer resultsIterator.Close()

	var records []*HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not get the value of resultsIterator. %s", err)
		}

		var farmer Farmer
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &farmer)
			if err != nil {
				return nil, err
			}
		} else {
			farmer = Farmer{
				UserName: userName,
			}
		}

		timestamp := response.Timestamp.AsTime()

		formattedTime := timestamp.Format(time.RFC1123)

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: formattedTime,
			Record:    &farmer,
			IsDelete:  response.IsDelete,
		}
		records = append(records, &record)
	}
	return records, nil
}

func (c *FarmerContract) GetFarmersWithPagination(ctx contractapi.TransactionContextInterface, pageSize int32, bookmark string) (*PaginatedQueryResult, error) {

	queryString := `{"selector":{}}`

	resultsIterator, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, fmt.Errorf("could not get the farmer records. %s", err)
	}
	defer resultsIterator.Close()

	farmers, err := farmerResultIteratorFunction(resultsIterator)
	if err != nil {
		return nil, fmt.Errorf("could not return the farmer records %s", err)
	}

	return &PaginatedQueryResult{
		Records:             farmers,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

func (f *FarmerContract) GetMatchingOrders(ctx contractapi.TransactionContextInterface, userName string) ([]*Process, error) {

	car, err := f.ReadFarmerDetails(ctx, userName)
	if err != nil {
		return nil, fmt.Errorf("error reading farmer %v", err)
	}
	queryString := fmt.Sprintf(`{"selector":{"typeofBeans":"%s","color":"%s"}}`, car.TypeofBeans, car.Color)
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(collectionName, queryString)

	if err != nil {
		return nil, fmt.Errorf("could not get the data. %s", err)
	}
	defer resultsIterator.Close()

	return OrderResultIteratorFunction(resultsIterator)

}

func (f *FarmerContract) MatchOrder(ctx contractapi.TransactionContextInterface, userName string, processId string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not fetch client identity. %s", err)
	}

	if clientOrgID == "FarmersMSP" {
		bytes, err := ctx.GetStub().GetPrivateData(collectionName, processId)
		if err != nil {
			return "", fmt.Errorf("could not get the private data: %s", err)
		}

		var process Process

		err = json.Unmarshal(bytes, &process)

		if err != nil {
			return "", fmt.Errorf("could not unmarshal the data. %s", err)
		}

		farmer, err := f.ReadFarmerDetails(ctx, userName)
		if err != nil {
			return "", fmt.Errorf("could not read the data. %s", err)
		}

		if farmer.TypeofBeans == process.TypeofBeans && farmer.Color == process.Color {
			farmer.Name = process.Name
			farmer.Status = "assigned to a processor"

			bytes, _ := json.Marshal(farmer)

			ctx.GetStub().DelPrivateData(collectionName, processId)

			err = ctx.GetStub().PutState(userName, bytes)

			if err != nil {
				return "", fmt.Errorf("could not add the data %s", err)
			} else {
				return fmt.Sprintf("Deleted order %v and Assigned %v to %v", processId, farmer.UserName, process.Name), nil
			}
		} else {
			return "", fmt.Errorf("order is not matching")
		}
	} else {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}
}



