package keys

import (
	"log"

	"github.com/DaveAppleton/ether_go/ethKeys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func KeyTx(key *ethKeys.AccountKey) *bind.TransactOpts {
	return bind.NewKeyedTransactor(key.GetKey())
}

func Role(job string) *bind.TransactOpts {
	return KeyTx(RoleKey(job))
}

func RoleKey(job string) *ethKeys.AccountKey {
	transactorAcc := ethKeys.NewKey("adminKeys/" + job)
	if err := transactorAcc.RestoreOrCreate(); err != nil {
		log.Println(err)
	}
	return transactorAcc
}

func RoleAddress(job string) common.Address {
	return RoleKey(job).PublicKey()
}
