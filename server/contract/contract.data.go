package contract

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

var contractMap = struct {
	sync.RWMutex
	m map[int]Contract
}{m: make(map[int]Contract)}

func init() {
	fmt.Println("loading contracts...")
	contMap, err := loadContractMap()
	contractMap.m = contMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%d contracts loaded...\n", len(contractMap.m))
}

func loadContractMap() (map[int]Contract, error) {
	fileName := "contracts.json"

	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exists", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	contractsList := make([]Contract, 0)
	err = json.Unmarshal([]byte(file), &contractsList)
	if err != nil {
		log.Fatal(err)
	}
	contMap := make(map[int]Contract)
	for i := 0; i < len(contractsList); i++ {
		contMap[contractsList[i].ID] = contractsList[i]
	}
	return contMap, nil
}

func getContractbyId(contractID int) (*Contract, error) {
	contractMap.RLock()
	defer contractMap.RUnlock()
	if contract, ok := contractMap.m[contractID]; ok {
		return &contract, nil
	}
	return &Contract{}, fmt.Errorf("contract with such id '%v' does not exists", contractID)
}

func removeContract(contractID int) {
	contractMap.Lock()
	defer contractMap.Unlock()
	delete(contractMap.m, contractID)
}

func getContractList() []Contract {
	contractMap.RLock()
	contracts := make([]Contract, 0, len(contractMap.m))
	for _, value := range contractMap.m {
		contracts = append(contracts, value)
	}
	contractMap.RUnlock()
	return contracts
}

// helper
// get next ID

func getContractIds() []int {
	contractMap.RLock()
	contractIds := []int{}
	for key := range contractMap.m {
		contractIds = append(contractIds, key)
	}
	contractMap.RUnlock()
	sort.Ints(contractIds)
	return contractIds
}

func getNextContractID() int {
	contractIDs := getContractIds()
	return contractIDs[len(contractIDs)-1] + 1
}

func addOrUpdateContract(contract Contract) (int, error) {
	// if the product id is set, update, otherwise add
	addOrUpdateID := -1
	if contract.ID > 0 {
		_, err := getContractbyId(contract.ID)
		if err != nil {
			return 0, fmt.Errorf("product id [%d] doesn't exist", contract.ID)
		}
		addOrUpdateID = contract.ID
	} else {
		addOrUpdateID = getNextContractID()
		contract.ID = addOrUpdateID
	}
	contractMap.Lock()
	contractMap.m[addOrUpdateID] = contract
	contractMap.Unlock()
	return addOrUpdateID, nil
}
