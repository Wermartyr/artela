package tx

import (
	"github.com/artela-network/artela/ethereum/utils"
	"math"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/artela-network/artela/app"
)

var (
	feeAmt     = math.Pow10(16)
	DefaultFee = sdk.NewCoin(utils.BaseDenom, sdk.NewIntFromUint64(uint64(feeAmt))) // 0.01 Artela
)

// CosmosTxArgs contains the params to create a cosmos txs
type CosmosTxArgs struct {
	// TxCfg is the client txs config
	TxCfg client.TxConfig
	// Priv is the private key that will be used to sign the txs
	Priv cryptotypes.PrivKey
	// ChainID is the chain's id on cosmos format, e.g. 'artela_11820-1'
	ChainID string
	// Gas to be used on the txs
	Gas uint64
	// GasPrice to use on txs
	GasPrice *sdkmath.Int
	// Fees is the fee to be used on the txs (amount and denom)
	Fees sdk.Coins
	// FeeGranter is the account address of the fee granter
	FeeGranter sdk.AccAddress
	// Msgs slice of messages to include on the txs
	Msgs []sdk.Msg
}

// PrepareCosmosTx creates a cosmos txs and signs it with the provided messages and private key.
// It returns the signed txs and an error
func PrepareCosmosTx(
	ctx sdk.Context,
	appArtela *app.Artela,
	args CosmosTxArgs,
) (authsigning.Tx, error) {
	txBuilder := args.TxCfg.NewTxBuilder()

	txBuilder.SetGasLimit(args.Gas)

	var fees sdk.Coins
	if args.GasPrice != nil {
		fees = sdk.Coins{{Denom: utils.BaseDenom, Amount: args.GasPrice.MulRaw(int64(args.Gas))}}
	} else {
		fees = sdk.Coins{DefaultFee}
	}

	txBuilder.SetFeeAmount(fees)
	if err := txBuilder.SetMsgs(args.Msgs...); err != nil {
		return nil, err
	}

	txBuilder.SetFeeGranter(args.FeeGranter)

	return signCosmosTx(
		ctx,
		appArtela,
		args,
		txBuilder,
	)
}

// signCosmosTx signs the cosmos txs on the txBuilder provided using
// the provided private key
func signCosmosTx(
	ctx sdk.Context,
	appArtela *app.Artela,
	args CosmosTxArgs,
	txBuilder client.TxBuilder,
) (authsigning.Tx, error) {
	addr := sdk.AccAddress(args.Priv.PubKey().Address().Bytes())
	seq, err := appArtela.AccountKeeper.GetSequence(ctx, addr)
	if err != nil {
		return nil, err
	}

	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	sigV2 := signing.SignatureV2{
		PubKey: args.Priv.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  args.TxCfg.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: seq,
	}

	sigsV2 := []signing.SignatureV2{sigV2}

	if err := txBuilder.SetSignatures(sigsV2...); err != nil {
		return nil, err
	}

	// Second round: all signer infos are set, so each signer can sign.
	accNumber := appArtela.AccountKeeper.GetAccount(ctx, addr).GetAccountNumber()
	signerData := authsigning.SignerData{
		ChainID:       args.ChainID,
		AccountNumber: accNumber,
		Sequence:      seq,
	}
	sigV2, err = tx.SignWithPrivKey(
		args.TxCfg.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder, args.Priv, args.TxCfg,
		seq,
	)
	if err != nil {
		return nil, err
	}

	sigsV2 = []signing.SignatureV2{sigV2}
	if err = txBuilder.SetSignatures(sigsV2...); err != nil {
		return nil, err
	}
	return txBuilder.GetTx(), nil
}

var _ sdk.Tx = &InvalidTx{}

// InvalidTx defines a type, which satisfies the sdk.Tx interface, but
// holds no valid txs information.
//
// NOTE: This is used for testing purposes, to serve the edge case of invalid data being passed to functions.
type InvalidTx struct{}

func (InvalidTx) GetMsgs() []sdk.Msg { return []sdk.Msg{nil} }

func (InvalidTx) ValidateBasic() error { return nil }
