// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"github.com/f01c5700/avalanchego/api/metrics"
	"github.com/f01c5700/avalanchego/ids"
	"github.com/f01c5700/avalanchego/snow"
	"github.com/f01c5700/avalanchego/snow/validators/validatorstest"
	"github.com/f01c5700/avalanchego/utils/crypto/bls"
	"github.com/f01c5700/avalanchego/utils/logging"
)

func TestSnowContext() *snow.Context {
	sk, err := bls.NewSecretKey()
	if err != nil {
		panic(err)
	}
	pk := bls.PublicFromSecretKey(sk)
	return &snow.Context{
		NetworkID:      0,
		SubnetID:       ids.Empty,
		ChainID:        ids.Empty,
		NodeID:         ids.EmptyNodeID,
		PublicKey:      pk,
		Log:            logging.NoLog{},
		BCLookup:       ids.NewAliaser(),
		Metrics:        metrics.NewPrefixGatherer(),
		ChainDataDir:   "",
		ValidatorState: &validatorstest.State{},
	}
}
