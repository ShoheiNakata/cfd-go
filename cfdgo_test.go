package cfdgo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// first test
func TestInitialize(t *testing.T) {
	ret := CfdInitialize()
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestInitialize test done.\n")
}

func TestCfdCreateHandle(t *testing.T) {
	ret := CfdCreateHandle(nil)
	assert.Equal(t, (int)(KCfdIllegalArgumentError), ret)

	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdCreateHandle test done.\n")
}

func TestCfdGetLastError(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	lastErr := CfdGetLastErrorCode(handle)
	assert.Equal(t, (int)(KCfdSuccess), lastErr)

	errStr, strret := CfdGoGetLastErrorMessage(handle)
	assert.Equal(t, (int)(KCfdSuccess), strret)
	assert.Equal(t, "", errStr)

	_, _, _, strret = CfdGoCreateAddress(handle, 200, "", "", 200)
	lastErr = CfdGetLastErrorCode(handle)
	assert.Equal(t, (int)(KCfdIllegalArgumentError), lastErr)
	assert.Equal(t, strret, lastErr)
	errStr, _ = CfdGoGetLastErrorMessage(handle)
	assert.Equal(t, "Illegal network type.", errStr)

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdGetLastError test done.\n")
}

func TestCfdGetSupportedFunction(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	flag, cfdRet := CfdGoGetSupportedFunction()
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, uint64(1), (flag & 0x01))

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdGetSupportedFunction test done.\n")
}

func TestCfdGoCreateAddress(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	hashType := (int)(KCfdP2pkh)
	networkType := (int)(KCfdNetworkLiquidv1)
	pubkey := "0279BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798"
	address, lockingScript, segwitLockingScript, cfdRet := CfdGoCreateAddress(handle, hashType, pubkey, "", networkType)
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, "Q7wegLt2qMGhm28vch6VTzvpzs8KXvs4X7", address)
	assert.Equal(t, "76a914751e76e8199196d454941c45d1b3a323f1433bd688ac", lockingScript)
	assert.Equal(t, "", segwitLockingScript)
	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	hashType = (int)(KCfdP2sh)
	redeemScript := "210279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798ac"
	address, lockingScript, segwitLockingScript, cfdRet = CfdGoCreateAddress(
		handle, hashType, "", redeemScript, networkType)
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, "GkSEheszYzEBMgX9G9ueaAyLVg8gfZwiDY", address)
	assert.Equal(t, "a91423b0ad3477f2178bc0b3eed26e4e6316f4e83aa187", lockingScript)
	assert.Equal(t, "", segwitLockingScript)
	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	hashType = (int)(KCfdP2shP2wpkh)
	pubkey = "0205ffcdde75f262d66ada3dd877c7471f8f8ee9ee24d917c3e18d01cee458bafe"
	address, lockingScript, segwitLockingScript, cfdRet = CfdGoCreateAddress(
		handle, hashType, pubkey, "", networkType)
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, "GsaK3GXnFAjdfZDBPPo9PD6UNyAJ53nS9Z", address)
	assert.Equal(t, "a9147200818f884ee12b964442b059c11d0712b6abe787", lockingScript)
	assert.Equal(t, "0014ef692e4bf0cd5ed05235a4fc582ec4a4ff9695b4", segwitLockingScript)
	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	hashType = (int)(KCfdP2wpkh)
	networkType = (int)(KCfdNetworkElementsRegtest)
	pubkey = "02bedf98a38247c1718fdff7e07561b4dc15f10323ebb0accab581778e72c2e995"
	address, lockingScript, segwitLockingScript, cfdRet = CfdGoCreateAddress(
		handle, hashType, pubkey, "", networkType)
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, "ert1qs58jzsgjsteydejyhy32p2v2vm8llh9uns6d93", address)
	assert.Equal(t, "0014850f21411282f246e644b922a0a98a66cfffdcbc", lockingScript)
	assert.Equal(t, "", segwitLockingScript)
	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdGoCreateAddress test done.\n")
}

func TestCfdGoCreateMultisigScript(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	networkType := (int)(KCfdNetworkLiquidv1)
	hashType := (int)(KCfdP2shP2wsh)
	pubkeys := []string{"0205ffcdde75f262d66ada3dd877c7471f8f8ee9ee24d917c3e18d01cee458bafe", "02be61f4350b4ae7544f99649a917f48ba16cf48c983ac1599774958d88ad17ec5"}
	address, redeemScript, witnessScript, cfdRet := CfdGoCreateMultisigScript(handle, networkType, hashType, pubkeys, uint32(2))
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, "H4PB6YPgiTmQLiMU7b772LMFY9vA4gSUC1", address)
	assert.Equal(t, "0020f39f6272ba6b57918eb047c5dc44fb475356b0f24c12fca39b19284e80008a42", redeemScript)
	assert.Equal(t, "52210205ffcdde75f262d66ada3dd877c7471f8f8ee9ee24d917c3e18d01cee458bafe2102be61f4350b4ae7544f99649a917f48ba16cf48c983ac1599774958d88ad17ec552ae", witnessScript)
	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdGoCreateMultisigScript test done.\n")
}

func TestCfdGoGetAddressesFromMultisig(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	networkType := (int)(KCfdNetworkLiquidv1)
	hashType := (int)(KCfdP2shP2wpkh)
	redeemScript := "52210205ffcdde75f262d66ada3dd877c7471f8f8ee9ee24d917c3e18d01cee458bafe2102be61f4350b4ae7544f99649a917f48ba16cf48c983ac1599774958d88ad17ec552ae"
	addressList, pubkeyList, cfdRet := CfdGoGetAddressesFromMultisig(handle, redeemScript, networkType, hashType)
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, 2, len(addressList))
	assert.Equal(t, 2, len(pubkeyList))
	if len(addressList) == 2 {
		assert.Equal(t, "GsaK3GXnFAjdfZDBPPo9PD6UNyAJ53nS9Z", addressList[0])
		assert.Equal(t, "GzGfkxAuJGSE7TL8KgMYmBRftjHPEFTSzS", addressList[1])
	}
	if len(pubkeyList) == 2 {
		assert.Equal(t, "0205ffcdde75f262d66ada3dd877c7471f8f8ee9ee24d917c3e18d01cee458bafe", pubkeyList[0])
		assert.Equal(t, "02be61f4350b4ae7544f99649a917f48ba16cf48c983ac1599774958d88ad17ec5", pubkeyList[1])
	}
	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdGoGetAddressesFromMultisig test done.\n")
}

func TestCfdGoParseDescriptor(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	// PKH
	networkType := (int)(KCfdNetworkLiquidv1)
	descriptorDataList, multisigList, cfdRet := CfdGoParseDescriptor(handle,
		"pkh(02c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5)",
		networkType,
		"")
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, 1, len(descriptorDataList))
	assert.Equal(t, 0, len(multisigList))
	if len(descriptorDataList) == 1 {
		assert.Equal(t, uint32(0), descriptorDataList[0].depth)
		assert.Equal(t, (int)(KCfdDescriptorScriptPkh), descriptorDataList[0].scriptType)
		assert.Equal(t, "76a91406afd46bcdfd22ef94ac122aa11f241244a37ecc88ac", descriptorDataList[0].lockingScript)
		assert.Equal(t, "PwsjpD1YkjcfZ95WGVZuvGfypkKmpogoA3", descriptorDataList[0].address)
		assert.Equal(t, (int)(KCfdP2pkh), descriptorDataList[0].hashType)
		assert.Equal(t, "", descriptorDataList[0].redeemScript)
		assert.Equal(t, (int)(KCfdDescriptorKeyPublic), descriptorDataList[0].keyType)
		assert.Equal(t, "02c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5", descriptorDataList[0].pubkey)
		assert.Equal(t, "", descriptorDataList[0].extPubkey)
		assert.Equal(t, "", descriptorDataList[0].extPrivkey)
		assert.Equal(t, false, descriptorDataList[0].isMultisig)
	}
	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	// p2sh-p2wsh(pkh)
	networkType = (int)(KCfdNetworkLiquidv1)
	descriptorDataList, multisigList, cfdRet = CfdGoParseDescriptor(handle,
		"sh(wsh(pkh(02e493dbf1c10d80f3581e4904930b1404cc6c13900ee0758474fa94abe8c4cd13)))",
		networkType, "")
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, 3, len(descriptorDataList))
	assert.Equal(t, 0, len(multisigList))
	if len(descriptorDataList) == 3 {
		// 0
		assert.Equal(t, uint32(0), descriptorDataList[0].depth)
		assert.Equal(t, (int)(KCfdDescriptorScriptSh), descriptorDataList[0].scriptType)
		assert.Equal(t, "a91455e8d5e8ee4f3604aba23c71c2684fa0a56a3a1287", descriptorDataList[0].lockingScript)
		assert.Equal(t, "Gq1mmExLuSEwfzzk6YtUxJ769grv6T5Tak", descriptorDataList[0].address)
		assert.Equal(t, (int)(KCfdP2shP2wsh), descriptorDataList[0].hashType)
		assert.Equal(t, "0020fc5acc302aab97f821f9a61e1cc572e7968a603551e95d4ba12b51df6581482f", descriptorDataList[0].redeemScript)
		assert.Equal(t, (int)(KCfdDescriptorKeyNull), descriptorDataList[0].keyType)
		assert.Equal(t, "", descriptorDataList[0].pubkey)
		assert.Equal(t, "", descriptorDataList[0].extPubkey)
		assert.Equal(t, "", descriptorDataList[0].extPrivkey)
		assert.Equal(t, false, descriptorDataList[0].isMultisig)
		// 1
		assert.Equal(t, uint32(1), descriptorDataList[1].depth)
		assert.Equal(t, (int)(KCfdDescriptorScriptWsh), descriptorDataList[1].scriptType)
		assert.Equal(t, "0020fc5acc302aab97f821f9a61e1cc572e7968a603551e95d4ba12b51df6581482f", descriptorDataList[1].lockingScript)
		assert.Equal(t, "ex1ql3dvcvp24wtlsg0e5c0pe3tju7tg5cp428546jap9dga7evpfqhs0htdlf", descriptorDataList[1].address)
		assert.Equal(t, (int)(KCfdP2wsh), descriptorDataList[1].hashType)
		assert.Equal(t, "76a914c42e7ef92fdb603af844d064faad95db9bcdfd3d88ac", descriptorDataList[1].redeemScript)
		assert.Equal(t, (int)(KCfdDescriptorKeyNull), descriptorDataList[1].keyType)
		assert.Equal(t, "", descriptorDataList[1].pubkey)
		assert.Equal(t, "", descriptorDataList[1].extPubkey)
		assert.Equal(t, "", descriptorDataList[1].extPrivkey)
		assert.Equal(t, false, descriptorDataList[1].isMultisig)
		// 2
		assert.Equal(t, uint32(2), descriptorDataList[2].depth)
		assert.Equal(t, (int)(KCfdDescriptorScriptPkh), descriptorDataList[2].scriptType)
		assert.Equal(t, "76a914c42e7ef92fdb603af844d064faad95db9bcdfd3d88ac", descriptorDataList[2].lockingScript)
		assert.Equal(t, "QF9hGPQMVAPc8RxTHALgSvNPWEjGbL9bse", descriptorDataList[2].address)
		assert.Equal(t, (int)(KCfdP2pkh), descriptorDataList[2].hashType)
		assert.Equal(t, "", descriptorDataList[2].redeemScript)
		assert.Equal(t, (int)(KCfdDescriptorKeyPublic), descriptorDataList[2].keyType)
		assert.Equal(t, "02e493dbf1c10d80f3581e4904930b1404cc6c13900ee0758474fa94abe8c4cd13", descriptorDataList[2].pubkey)
		assert.Equal(t, "", descriptorDataList[2].extPubkey)
		assert.Equal(t, "", descriptorDataList[2].extPrivkey)
		assert.Equal(t, false, descriptorDataList[2].isMultisig)
	}
	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	// multisig (bitcoin)
	networkType = (int)(KCfdNetworkMainnet)
	descriptorDataList, multisigList, cfdRet = CfdGoParseDescriptor(handle,
		"wsh(multi(1,xpub661MyMwAqRbcFW31YEwpkMuc5THy2PSt5bDMsktWQcFF8syAmRUapSCGu8ED9W6oDMSgv6Zz8idoc4a6mr8BDzTJY47LJhkJ8UB7WEGuduB/1/0/*,xpub69H7F5d8KSRgmmdJg2KhpAK8SR3DjMwAdkxj3ZuxV27CprR9LgpeyGmXUbC6wb7ERfvrnKZjXoUmmDznezpbZb7ap6r1D3tgFxHmwMkQTPH/0/0/*))",
		networkType,
		"0")
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, 1, len(descriptorDataList))
	assert.Equal(t, 2, len(multisigList))
	if len(descriptorDataList) == 1 {
		assert.Equal(t, uint32(0), descriptorDataList[0].depth)
		assert.Equal(t, (int)(KCfdDescriptorScriptWsh), descriptorDataList[0].scriptType)
		assert.Equal(t, "002064969d8cdca2aa0bb72cfe88427612878db98a5f07f9a7ec6ec87b85e9f9208b", descriptorDataList[0].lockingScript)
		assert.Equal(t, "bc1qvjtfmrxu524qhdevl6yyyasjs7xmnzjlqlu60mrwepact60eyz9s9xjw0c", descriptorDataList[0].address)
		assert.Equal(t, (int)(KCfdP2wsh), descriptorDataList[0].hashType)
		assert.Equal(t, "51210205f8f73d8a553ad3287a506dbd53ed176cadeb200c8e4f7d68a001b1aed871062102c04c4e03921809fcbef9a26da2d62b19b2b4eb383b3e6cfaaef6370e7514477452ae", descriptorDataList[0].redeemScript)
		assert.Equal(t, (int)(KCfdDescriptorKeyNull), descriptorDataList[0].keyType)
		assert.Equal(t, "", descriptorDataList[0].pubkey)
		assert.Equal(t, "", descriptorDataList[0].extPubkey)
		assert.Equal(t, "", descriptorDataList[0].extPrivkey)
		assert.Equal(t, true, descriptorDataList[0].isMultisig)
	}
	if len(multisigList) == 2 {
		assert.Equal(t, (int)(KCfdDescriptorKeyBip32), multisigList[0].keyType)
		assert.Equal(t, "0205f8f73d8a553ad3287a506dbd53ed176cadeb200c8e4f7d68a001b1aed87106", multisigList[0].pubkey)
		assert.Equal(t, "xpub6BgWskLoyHmAUeKWgUXCGfDdCMRXseEjRCMEMvjkedmHpnvWtpXMaCRm8qcADw9einPR8o2c49ZpeHRZP4uYwGeMU2T63G7uf2Y1qJavrWQ", multisigList[0].extPubkey)
		assert.Equal(t, "", multisigList[0].extPrivkey)
		assert.Equal(t, (int)(KCfdDescriptorKeyBip32), multisigList[1].keyType)
		assert.Equal(t, "02c04c4e03921809fcbef9a26da2d62b19b2b4eb383b3e6cfaaef6370e75144774", multisigList[1].pubkey)
		assert.Equal(t, "xpub6EKMC2gSMfKgQJ3iNMZVNB4GLH1Dc4hNPah1iMbbztxdUPRo84MMcTgkPATWNRyzr7WifKrt5VvQi4GEqRwybCP1LHoXBKLN6cB15HuBKPE", multisigList[1].extPubkey)
		assert.Equal(t, "", multisigList[1].extPrivkey)
	}
	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdGoParseDescriptor test done.\n")
}

func TestCfdCreateRawTransaction(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	txHex, cfdRet := CfdGoInitializeConfidentialTx(handle, uint32(2), uint32(0))
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, "0200000000000000000000", txHex)

	if cfdRet == (int)(KCfdSuccess) {
		txHex, cfdRet = CfdGoAddConfidentialTxIn(
			handle, txHex,
			"7461b02405414d79e79a5050684a333c922c1136f4bdff5fb94b551394edebbd", 0,
			uint32(4294967295))
		assert.Equal(t, (int)(KCfdSuccess), cfdRet)
		assert.Equal(t, "020000000001bdebed9413554bb95fffbdf436112c923c334a6850509ae7794d410524b061740000000000ffffffff0000000000", txHex)
	}

	if cfdRet == (int)(KCfdSuccess) {
		txHex, cfdRet = CfdGoAddConfidentialTxIn(
			handle, txHex,
			"1497e1f146bc5fe00b6268ea16a7069ecb90a2a41a183446d5df8965d2356dc1", 1,
			uint32(4294967295))
		assert.Equal(t, (int)(KCfdSuccess), cfdRet)
		assert.Equal(t, "020000000002bdebed9413554bb95fffbdf436112c923c334a6850509ae7794d410524b061740000000000ffffffffc16d35d26589dfd54634181aa4a290cb9e06a716ea68620be05fbc46f1e197140100000000ffffffff0000000000", txHex)
	}

	if cfdRet == (int)(KCfdSuccess) {
		txHex, cfdRet = CfdGoAddConfidentialTxOut(
			handle, txHex,
			"ef47c42d34de1b06a02212e8061323f50d5f02ceed202f1cb375932aa299f751",
			int64(100000000), "",
			"CTEw7oSCUWDfmfhCEdsB3gsG7D9b4xLCZEq71H8JxRFeBu7yQN3CbSF6qT6J4F7qji4bq1jVSdVcqvRJ",
			"", "")
		assert.Equal(t, (int)(KCfdSuccess), cfdRet)
		assert.Equal(t, "020000000002bdebed9413554bb95fffbdf436112c923c334a6850509ae7794d410524b061740000000000ffffffffc16d35d26589dfd54634181aa4a290cb9e06a716ea68620be05fbc46f1e197140100000000ffffffff010151f799a22a9375b31c2f20edce025f0df5231306e81222a0061bde342dc447ef010000000005f5e10003a630456ab6d50b57981e085abced70e2816289ae2b49a44c2f471b205134c12b1976a914d08f5ba8874d36cf97d19379b370f1f23ba36d5888ac00000000", txHex)
	}

	if cfdRet == (int)(KCfdSuccess) {
		txHex, cfdRet = CfdGoAddConfidentialTxOut(
			handle, txHex,
			"6f1a4b6bd5571b5f08ab79c314dc6483f9b952af2f5ef206cd6f8e68eb1186f3",
			int64(1900500000), "",
			"2dxZw5iVZ6Pmqoc5Vn8gkUWDGB5dXuMBCmM", "", "")
		assert.Equal(t, (int)(KCfdSuccess), cfdRet)
		assert.Equal(t, "020000000002bdebed9413554bb95fffbdf436112c923c334a6850509ae7794d410524b061740000000000ffffffffc16d35d26589dfd54634181aa4a290cb9e06a716ea68620be05fbc46f1e197140100000000ffffffff020151f799a22a9375b31c2f20edce025f0df5231306e81222a0061bde342dc447ef010000000005f5e10003a630456ab6d50b57981e085abced70e2816289ae2b49a44c2f471b205134c12b1976a914d08f5ba8874d36cf97d19379b370f1f23ba36d5888ac01f38611eb688e6fcd06f25e2faf52b9f98364dc14c379ab085f1b57d56b4b1a6f010000000071475420001976a914fdd725970db682de970e7669646ed7afb8348ea188ac00000000", txHex)
	}

	if cfdRet == (int)(KCfdSuccess) {
		txHex, cfdRet = CfdGoAddConfidentialTxOut(
			handle, txHex,
			"6f1a4b6bd5571b5f08ab79c314dc6483f9b952af2f5ef206cd6f8e68eb1186f3",
			int64(500000), "", "", "", "")
		assert.Equal(t, (int)(KCfdSuccess), cfdRet)
		assert.Equal(t, "020000000002bdebed9413554bb95fffbdf436112c923c334a6850509ae7794d410524b061740000000000ffffffffc16d35d26589dfd54634181aa4a290cb9e06a716ea68620be05fbc46f1e197140100000000ffffffff030151f799a22a9375b31c2f20edce025f0df5231306e81222a0061bde342dc447ef010000000005f5e10003a630456ab6d50b57981e085abced70e2816289ae2b49a44c2f471b205134c12b1976a914d08f5ba8874d36cf97d19379b370f1f23ba36d5888ac01f38611eb688e6fcd06f25e2faf52b9f98364dc14c379ab085f1b57d56b4b1a6f010000000071475420001976a914fdd725970db682de970e7669646ed7afb8348ea188ac01f38611eb688e6fcd06f25e2faf52b9f98364dc14c379ab085f1b57d56b4b1a6f01000000000007a120000000000000", txHex)
	}

	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdCreateRawTransaction test done.\n")
}

func TestCfdGetTransaction(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	txHex := "0200000000020f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570000000000ffffffff0f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570100008000ffffffffd8bbe31bc590cbb6a47d2e53a956ec25d8890aefd60dcfc93efd34727554890b0683fe0819a4f9770c8a7cd5824e82975c825e017aff8ba0d6a5eb4959cf9c6f010000000023c346000004017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000003b947f6002200d8510dfcf8e2330c0795c771d1e6064daab2f274ac32a6e2708df9bfa893d17a914ef3e40882e17d6e477082fcafeb0f09dc32d377b87010bad521bafdac767421d45b71b29a349c7b2ca2a06b5d8e3b5898c91df2769ed010000000029b9270002cc645552109331726c0ffadccab21620dd7a5a33260c6ac7bd1c78b98cb1e35a1976a9146c22e209d36612e0d9d2a20b814d7d8648cc7a7788ac017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000000000c350000001cdb0ed311810e61036ac9255674101497850f5eee5e4320be07479c05473cbac010000000023c3460003ce4c4eac09fe317f365e45c00ffcf2e9639bc0fd792c10f72cdc173c4e5ed8791976a9149bdcb18911fa9faad6632ca43b81739082b0a19588ac00000000"

	count, cfdRet := CfdGoGetConfidentialTxInCount(handle, txHex)
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, uint32(2), count)

	count, cfdRet = CfdGoGetConfidentialTxOutCount(handle, txHex)
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, uint32(4), count)

	if cfdRet == (int)(KCfdSuccess) {
		txid, vout, sequence, scriptSig, txinRet := CfdGoGetConfidentialTxIn(handle, txHex, uint32(1))
		assert.Equal(t, (int)(KCfdSuccess), txinRet)
		assert.Equal(t, "57a15002d066ce52573d674df925c9bc0f1164849420705f2cfad8a68111230f", txid)
		assert.Equal(t, uint32(1), vout)
		assert.Equal(t, uint32(4294967295), sequence)
		assert.Equal(t, "", scriptSig)
	}

	if cfdRet == (int)(KCfdSuccess) {
		entropy, nonce, assetValue, tokenValue, assetRangeproof, tokenRangeproof, issueRet := CfdGoGetTxInIssuanceInfo(handle, txHex, uint32(1))
		assert.Equal(t, (int)(KCfdSuccess), issueRet)
		assert.Equal(t, "6f9ccf5949eba5d6a08bff7a015e825c97824e82d57c8a0c77f9a41908fe8306", entropy)
		assert.Equal(t, "0b8954757234fd3ec9cf0dd6ef0a89d825ec56a9532e7da4b6cb90c51be3bbd8", nonce)
		assert.Equal(t, "010000000023c34600", assetValue)
		assert.Equal(t, "", tokenValue)
		assert.Equal(t, "", assetRangeproof)
		assert.Equal(t, "", tokenRangeproof)
	}

	if cfdRet == (int)(KCfdSuccess) {
		asset, satoshiValue, valueCommitment, nonce, lockingScript, surjectionProof, rangeproof, txoutRet := CfdGoGetConfidentialTxOut(handle, txHex, uint32(3))
		assert.Equal(t, (int)(KCfdSuccess), txoutRet)
		assert.Equal(t, "accb7354c07974e00b32e4e5eef55078490141675592ac3610e6101831edb0cd", asset)
		assert.Equal(t, int64(600000000), satoshiValue)
		assert.Equal(t, "010000000023c34600", valueCommitment)
		assert.Equal(t, "03ce4c4eac09fe317f365e45c00ffcf2e9639bc0fd792c10f72cdc173c4e5ed879", nonce)
		assert.Equal(t, "76a9149bdcb18911fa9faad6632ca43b81739082b0a19588ac", lockingScript)
		assert.Equal(t, "", surjectionProof)
		assert.Equal(t, "", rangeproof)
	}

	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdGetTransaction test done.\n")
}

func TestCfdSetRawReissueAsset(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	txHex := "0200000000020f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570000000000ffffffff0f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570100000000ffffffff03017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000003b947f6002200d8510dfcf8e2330c0795c771d1e6064daab2f274ac32a6e2708df9bfa893d17a914ef3e40882e17d6e477082fcafeb0f09dc32d377b87010bad521bafdac767421d45b71b29a349c7b2ca2a06b5d8e3b5898c91df2769ed010000000029b9270002cc645552109331726c0ffadccab21620dd7a5a33260c6ac7bd1c78b98cb1e35a1976a9146c22e209d36612e0d9d2a20b814d7d8648cc7a7788ac017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000000000c350000000000000"

	asset, outTxHex, cfdRet := CfdGoSetRawReissueAsset(
		handle, txHex, "57a15002d066ce52573d674df925c9bc0f1164849420705f2cfad8a68111230f",
		uint32(1),
		int64(600000000), "0b8954757234fd3ec9cf0dd6ef0a89d825ec56a9532e7da4b6cb90c51be3bbd8",
		"6f9ccf5949eba5d6a08bff7a015e825c97824e82d57c8a0c77f9a41908fe8306",
		"CTExCoUri8VzkxbbhqzgsruWJ5zYtmoFXxCWtjiSLAzcMbpEWhHmDrZ66bAb41VsmSKnvJWrq2cfjUw9",
		"")
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, "accb7354c07974e00b32e4e5eef55078490141675592ac3610e6101831edb0cd", asset)
	assert.Equal(t, "0200000000020f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570000000000ffffffff0f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570100008000ffffffffd8bbe31bc590cbb6a47d2e53a956ec25d8890aefd60dcfc93efd34727554890b0683fe0819a4f9770c8a7cd5824e82975c825e017aff8ba0d6a5eb4959cf9c6f010000000023c346000004017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000003b947f6002200d8510dfcf8e2330c0795c771d1e6064daab2f274ac32a6e2708df9bfa893d17a914ef3e40882e17d6e477082fcafeb0f09dc32d377b87010bad521bafdac767421d45b71b29a349c7b2ca2a06b5d8e3b5898c91df2769ed010000000029b9270002cc645552109331726c0ffadccab21620dd7a5a33260c6ac7bd1c78b98cb1e35a1976a9146c22e209d36612e0d9d2a20b814d7d8648cc7a7788ac017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000000000c350000001cdb0ed311810e61036ac9255674101497850f5eee5e4320be07479c05473cbac010000000023c3460003ce4c4eac09fe317f365e45c00ffcf2e9639bc0fd792c10f72cdc173c4e5ed8791976a9149bdcb18911fa9faad6632ca43b81739082b0a19588ac00000000", outTxHex)

	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdSetRawReissueAsset test done.\n")
}

func TestCfdGetIssuanceBlindingKey(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	blindingKey, cfdRet := CfdGoGetIssuanceBlindingKey(
		handle, "ac2c1e4cce122139bb25abc50599e09738143cc4bc96e55f399a5e1e45d916a9",
		"57a15002d066ce52573d674df925c9bc0f1164849420705f2cfad8a68111230f", uint32(1))
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, "7d65c7970d836a878a1080399a3c11de39a8e82493e12b1ad154e383661fb77f", blindingKey)

	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdGetIssuanceBlindingKey test done.\n")
}

func TestCfdBlindTransaction(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	txHex := "0200000000020f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570000000000ffffffff0f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570100008000ffffffffd8bbe31bc590cbb6a47d2e53a956ec25d8890aefd60dcfc93efd34727554890b0683fe0819a4f9770c8a7cd5824e82975c825e017aff8ba0d6a5eb4959cf9c6f010000000023c346000004017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000003b947f6002200d8510dfcf8e2330c0795c771d1e6064daab2f274ac32a6e2708df9bfa893d17a914ef3e40882e17d6e477082fcafeb0f09dc32d377b87010bad521bafdac767421d45b71b29a349c7b2ca2a06b5d8e3b5898c91df2769ed010000000029b9270002cc645552109331726c0ffadccab21620dd7a5a33260c6ac7bd1c78b98cb1e35a1976a9146c22e209d36612e0d9d2a20b814d7d8648cc7a7788ac017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000000000c350000001cdb0ed311810e61036ac9255674101497850f5eee5e4320be07479c05473cbac010000000023c3460003ce4c4eac09fe317f365e45c00ffcf2e9639bc0fd792c10f72cdc173c4e5ed8791976a9149bdcb18911fa9faad6632ca43b81739082b0a19588ac00000000"

	blindHandle, cfdRet := CfdGoInitializeBlindTx(handle)
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)

	if cfdRet == (int)(KCfdSuccess) {
		cfdRet = CfdGoAddBlindTxInData(
			handle, blindHandle,
			"57a15002d066ce52573d674df925c9bc0f1164849420705f2cfad8a68111230f", uint32(0),
			"186c7f955149a5274b39e24b6a50d1d6479f552f6522d91f3a97d771f1c18179",
			"a10ecbe1be7a5f883d5d45d966e30dbc1beff5f21c55cec76cc21a2229116a9f",
			"ae0f46d1940f297c2dc3bbd82bf8ef6931a2431fbb05b3d3bc5df41af86ae808",
			int64(999637680), "", "")
		assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	}

	if cfdRet == (int)(KCfdSuccess) {
		cfdRet = CfdGoAddBlindTxInData(
			handle, blindHandle,
			"57a15002d066ce52573d674df925c9bc0f1164849420705f2cfad8a68111230f", uint32(1),
			"ed6927df918c89b5e3d8b5062acab2c749a3291bb7451d4267c7daaf1b52ad0b",
			"0b8954757234fd3ec9cf0dd6ef0a89d825ec56a9532e7da4b6cb90c51be3bbd8",
			"62e36e1f0fa4916b031648a6b6903083069fa587572a88b729250cde528cfd3b",
			int64(700000000),
			"7d65c7970d836a878a1080399a3c11de39a8e82493e12b1ad154e383661fb77f",
			"7d65c7970d836a878a1080399a3c11de39a8e82493e12b1ad154e383661fb77f")
		assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	}

	if cfdRet == (int)(KCfdSuccess) {
		cfdRet = CfdGoAddBlindTxOutData(
			handle, blindHandle, uint32(0),
			"02200d8510dfcf8e2330c0795c771d1e6064daab2f274ac32a6e2708df9bfa893d")
		assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	}

	if cfdRet == (int)(KCfdSuccess) {
		cfdRet = CfdGoAddBlindTxOutData(
			handle, blindHandle, uint32(1),
			"02cc645552109331726c0ffadccab21620dd7a5a33260c6ac7bd1c78b98cb1e35a")
		assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	}

	if cfdRet == (int)(KCfdSuccess) {
		cfdRet = CfdGoAddBlindTxOutData(
			handle, blindHandle, uint32(3),
			"03ce4c4eac09fe317f365e45c00ffcf2e9639bc0fd792c10f72cdc173c4e5ed879")
		assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	}

	if cfdRet == (int)(KCfdSuccess) {
		txHex, cfdRet = CfdGoFinalizeBlindTx(handle, blindHandle, txHex)
		assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	}

	CfdFreeBlindHandle(handle, blindHandle) // release

	// unblind test
	if cfdRet == (int)(KCfdSuccess) {
		asset, assetValue, aabf, avbf, token, tokenValue, tabf, tvbf, issueRet := CfdGoUnblindIssuance(
			handle, txHex, uint32(1),
			"7d65c7970d836a878a1080399a3c11de39a8e82493e12b1ad154e383661fb77f",
			"7d65c7970d836a878a1080399a3c11de39a8e82493e12b1ad154e383661fb77f")
		assert.Equal(t, (int)(KCfdSuccess), issueRet)
		assert.Equal(t, "accb7354c07974e00b32e4e5eef55078490141675592ac3610e6101831edb0cd", asset)
		assert.Equal(t, int64(600000000), assetValue)
		assert.Equal(t, "0000000000000000000000000000000000000000000000000000000000000000", aabf)
		assert.NotEqual(t, "0000000000000000000000000000000000000000000000000000000000000000", avbf)
		assert.Equal(t, "", token)
		assert.Equal(t, int64(0), tokenValue)
		assert.Equal(t, "", tabf)
		assert.Equal(t, "", tvbf)
	}

	if cfdRet == (int)(KCfdSuccess) {
		asset, value, abf, vbf, txoutRet := CfdGoUnblindTxOut(
			handle, txHex, uint32(0),
			"6a64f506be6e60b948987aa4d180d2ab05034a6a214146e06e28d4efe101d006")
		assert.Equal(t, (int)(KCfdSuccess), txoutRet)
		assert.Equal(t, "186c7f955149a5274b39e24b6a50d1d6479f552f6522d91f3a97d771f1c18179", asset)
		assert.Equal(t, int64(999587680), value)
		assert.NotEqual(t, "0000000000000000000000000000000000000000000000000000000000000000", abf)
		assert.NotEqual(t, "0000000000000000000000000000000000000000000000000000000000000000", vbf)
	}

	if cfdRet == (int)(KCfdSuccess) {
		asset, value, abf, vbf, txoutRet := CfdGoUnblindTxOut(
			handle, txHex, uint32(1),
			"94c85164605f589c4c572874f36b8301989c7fabfd44131297e95824d473681f")
		assert.Equal(t, (int)(KCfdSuccess), txoutRet)
		assert.Equal(t, "ed6927df918c89b5e3d8b5062acab2c749a3291bb7451d4267c7daaf1b52ad0b", asset)
		assert.Equal(t, int64(700000000), value)
		assert.NotEqual(t, "0000000000000000000000000000000000000000000000000000000000000000", abf)
		assert.NotEqual(t, "0000000000000000000000000000000000000000000000000000000000000000", vbf)
	}

	if cfdRet == (int)(KCfdSuccess) {
		asset, value, abf, vbf, txoutRet := CfdGoUnblindTxOut(
			handle, txHex, uint32(3),
			"0473d39aa6542e0c1bb6a2343b2319c3e92063dd019af4d47dbf50c460204f32")
		assert.Equal(t, (int)(KCfdSuccess), txoutRet)
		assert.Equal(t, "accb7354c07974e00b32e4e5eef55078490141675592ac3610e6101831edb0cd", asset)
		assert.Equal(t, int64(600000000), value)
		assert.NotEqual(t, "0000000000000000000000000000000000000000000000000000000000000000", abf)
		assert.NotEqual(t, "0000000000000000000000000000000000000000000000000000000000000000", vbf)
	}

	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdBlindTransaction test done.\n")
}

func TestCfdAddSignConfidentialTx(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	kTxData := "0200000000020f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570000000000ffffffff0f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570100008000ffffffffd8bbe31bc590cbb6a47d2e53a956ec25d8890aefd60dcfc93efd34727554890b0683fe0819a4f9770c8a7cd5824e82975c825e017aff8ba0d6a5eb4959cf9c6f010000000023c346000004017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000003b947f6002200d8510dfcf8e2330c0795c771d1e6064daab2f274ac32a6e2708df9bfa893d17a914ef3e40882e17d6e477082fcafeb0f09dc32d377b87010bad521bafdac767421d45b71b29a349c7b2ca2a06b5d8e3b5898c91df2769ed010000000029b9270002cc645552109331726c0ffadccab21620dd7a5a33260c6ac7bd1c78b98cb1e35a1976a9146c22e209d36612e0d9d2a20b814d7d8648cc7a7788ac017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000000000c350000001cdb0ed311810e61036ac9255674101497850f5eee5e4320be07479c05473cbac010000000023c3460003ce4c4eac09fe317f365e45c00ffcf2e9639bc0fd792c10f72cdc173c4e5ed8791976a9149bdcb18911fa9faad6632ca43b81739082b0a19588ac00000000"

	pubkey := "03f942716865bb9b62678d99aa34de4632249d066d99de2b5a2e542e54908450d6"
	privkey := "cU4KjNUT7GjHm7CkjRjG46SzLrXHXoH3ekXmqa2jTCFPMkQ64sw1"
	privkeyWifNetworkType := (int)(KCfdNetworkRegtest)
	txid := "57a15002d066ce52573d674df925c9bc0f1164849420705f2cfad8a68111230f"
	vout := uint32(0)
	txHex := ""
	sigHashType := (int)(KCfdSigHashAll)
	hashType := (int)(KCfdP2wpkh)
	isWitness := true
	if (hashType == (int)(KCfdP2pkh)) || (hashType == (int)(KCfdP2sh)) {
		isWitness = false
	}

	sighash, cfdRet := CfdGoCreateConfidentialSighash(
		handle, kTxData, txid, vout, hashType,
		pubkey, "", int64(13000000000000), "", sigHashType, false)
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, "c90939ef311f105806b401bcfa494921b8df297195fc125ebbd91a018c4066b9", sighash)

	signature, signRet := CfdGoCalculateEcSignature(
		handle, sighash, "", privkey, privkeyWifNetworkType, true)
	assert.Equal(t, (int)(KCfdSuccess), signRet)
	assert.Equal(t, "0268633a57723c6612ef217c49bdf804c632a14be2967c76afec4fd5781ad4c2131f358b2381a039c8c502959c64fbfeccf287be7dae710b4446968553aefbea", signature)

	// add signature
	txHex, signRet = CfdGoAddConfidentialTxDerSign(
		handle, kTxData, txid, vout, isWitness, signature, sigHashType, false, true)
	assert.Equal(t, (int)(KCfdSuccess), signRet)

	// add pubkey
	txHex, signRet = CfdGoAddConfidentialTxSign(
		handle, txHex, txid, vout, isWitness, pubkey, false)
	assert.Equal(t, (int)(KCfdSuccess), signRet)
	assert.Equal(t, "0200000001020f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570000000000ffffffff0f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570100008000ffffffffd8bbe31bc590cbb6a47d2e53a956ec25d8890aefd60dcfc93efd34727554890b0683fe0819a4f9770c8a7cd5824e82975c825e017aff8ba0d6a5eb4959cf9c6f010000000023c346000004017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000003b947f6002200d8510dfcf8e2330c0795c771d1e6064daab2f274ac32a6e2708df9bfa893d17a914ef3e40882e17d6e477082fcafeb0f09dc32d377b87010bad521bafdac767421d45b71b29a349c7b2ca2a06b5d8e3b5898c91df2769ed010000000029b9270002cc645552109331726c0ffadccab21620dd7a5a33260c6ac7bd1c78b98cb1e35a1976a9146c22e209d36612e0d9d2a20b814d7d8648cc7a7788ac017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000000000c350000001cdb0ed311810e61036ac9255674101497850f5eee5e4320be07479c05473cbac010000000023c3460003ce4c4eac09fe317f365e45c00ffcf2e9639bc0fd792c10f72cdc173c4e5ed8791976a9149bdcb18911fa9faad6632ca43b81739082b0a19588ac0000000000000247304402200268633a57723c6612ef217c49bdf804c632a14be2967c76afec4fd5781ad4c20220131f358b2381a039c8c502959c64fbfeccf287be7dae710b4446968553aefbea012103f942716865bb9b62678d99aa34de4632249d066d99de2b5a2e542e54908450d600000000000000000000000000", txHex)

	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdAddSignConfidentialTx test done.\n")
}

func TestCfdAddMultisigSignConfidentialTx(t *testing.T) {
	handle, ret := CfdGoCreateHandle()
	assert.Equal(t, (int)(KCfdSuccess), ret)

	kTxData := "0200000000020f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570000000000ffffffff0f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570100008000ffffffffd8bbe31bc590cbb6a47d2e53a956ec25d8890aefd60dcfc93efd34727554890b0683fe0819a4f9770c8a7cd5824e82975c825e017aff8ba0d6a5eb4959cf9c6f010000000023c346000004017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000003b947f6002200d8510dfcf8e2330c0795c771d1e6064daab2f274ac32a6e2708df9bfa893d17a914ef3e40882e17d6e477082fcafeb0f09dc32d377b87010bad521bafdac767421d45b71b29a349c7b2ca2a06b5d8e3b5898c91df2769ed010000000029b9270002cc645552109331726c0ffadccab21620dd7a5a33260c6ac7bd1c78b98cb1e35a1976a9146c22e209d36612e0d9d2a20b814d7d8648cc7a7788ac017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000000000c350000001cdb0ed311810e61036ac9255674101497850f5eee5e4320be07479c05473cbac010000000023c3460003ce4c4eac09fe317f365e45c00ffcf2e9639bc0fd792c10f72cdc173c4e5ed8791976a9149bdcb18911fa9faad6632ca43b81739082b0a19588ac00000000"

	txid := "57a15002d066ce52573d674df925c9bc0f1164849420705f2cfad8a68111230f"
	vout := uint32(0)

	pubkey1 := "02715ed9a5f16153c5216a6751b7d84eba32076f0b607550a58b209077ab7c30ad"
	privkey1 := "cRVLMWHogUo51WECRykTbeLNbm5c57iEpSegjdxco3oef6o5dbFi"
	pubkey2 := "02bfd7daa5d113fcbd8c2f374ae58cbb89cbed9570e898f1af5ff989457e2d4d71"
	privkey2 := "cQUTZ8VbWNYBEtrB7xwe41kqiKMQPRZshTvBHmkoJGaUfmS5pxzR"
	networkType := (int)(KCfdNetworkRegtest)
	sigHashType := (int)(KCfdSigHashAll)
	hashType := (int)(KCfdP2sh)

	// create multisig address
	pubkeys := []string{pubkey2, pubkey1}
	addr, multisigScript, _, cfdRet := CfdGoCreateMultisigScript(
		handle, networkType, hashType, pubkeys, uint32(2))
	assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	assert.Equal(t, "2MtG4TZaMXCNdEyUYAyJDraQRFwYC5j4S9U", addr)
	assert.Equal(t, "522102bfd7daa5d113fcbd8c2f374ae58cbb89cbed9570e898f1af5ff989457e2d4d712102715ed9a5f16153c5216a6751b7d84eba32076f0b607550a58b209077ab7c30ad52ae", multisigScript)

	// sign multisig
	multiSignHandle, muRet := CfdGoInitializeMultisigSign(handle)
	assert.Equal(t, (int)(KCfdSuccess), muRet)
	if cfdRet == (int)(KCfdSuccess) {
		satoshi := int64(13000000000000)
		sighash, sigRet0 := CfdGoCreateConfidentialSighash(handle, kTxData, txid, vout,
			hashType, "", multisigScript, satoshi, "", sigHashType, false)
		assert.Equal(t, (int)(KCfdSuccess), sigRet0)
		assert.Equal(t, "64878cbcd5c1805659d0747097cbf4b9ec5c187ebd80afa996c8fc95bd650b70", sighash)

		// user1
		signature1, sigRet1 := CfdGoCalculateEcSignature(
			handle, sighash, "", privkey1, networkType, true)
		assert.Equal(t, (int)(KCfdSuccess), sigRet1)

		sigRet1 = CfdGoAddMultisigSignDataToDer(
			handle, multiSignHandle, signature1, sigHashType, false, pubkey1)
		assert.Equal(t, (int)(KCfdSuccess), sigRet1)

		// user2
		signature2, sigRet2 := CfdGoCalculateEcSignature(
			handle, sighash, "", privkey2, networkType, true)
		assert.Equal(t, (int)(KCfdSuccess), sigRet2)

		sigRet2 = CfdGoAddMultisigSignDataToDer(
			handle, multiSignHandle, signature2, sigHashType, false, pubkey2)
		assert.Equal(t, (int)(KCfdSuccess), sigRet2)

		// generate
		txHex, muRet2 := CfdGoFinalizeElementsMultisigSign(
			handle, multiSignHandle, kTxData, txid, vout, hashType, "", multisigScript, true)
		assert.Equal(t, (int)(KCfdSuccess), muRet2)
		assert.Equal(t, "0200000000020f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a15700000000d90047304402206fc4cc7e489208a2f4d24f5d35466debab2ce7aa34b5d00e0a9426c9d63529cf02202ec744939ef0b4b629c7d87bc2d017714b52bb86dccb0fd0f10148f62b7a09ba01473044022073ea24720b24c736bcb305a5de2fd8117ca2f0a85d7da378fae5b90dc361d227022004c0088bf1b73a56ae5ec407cf9c330d7206ffbcd0c9bb1c72661726fd4990390147522102bfd7daa5d113fcbd8c2f374ae58cbb89cbed9570e898f1af5ff989457e2d4d712102715ed9a5f16153c5216a6751b7d84eba32076f0b607550a58b209077ab7c30ad52aeffffffff0f231181a6d8fa2c5f7020948464110fbcc925f94d673d5752ce66d00250a1570100008000ffffffffd8bbe31bc590cbb6a47d2e53a956ec25d8890aefd60dcfc93efd34727554890b0683fe0819a4f9770c8a7cd5824e82975c825e017aff8ba0d6a5eb4959cf9c6f010000000023c346000004017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000003b947f6002200d8510dfcf8e2330c0795c771d1e6064daab2f274ac32a6e2708df9bfa893d17a914ef3e40882e17d6e477082fcafeb0f09dc32d377b87010bad521bafdac767421d45b71b29a349c7b2ca2a06b5d8e3b5898c91df2769ed010000000029b9270002cc645552109331726c0ffadccab21620dd7a5a33260c6ac7bd1c78b98cb1e35a1976a9146c22e209d36612e0d9d2a20b814d7d8648cc7a7788ac017981c1f171d7973a1fd922652f559f47d6d1506a4be2394b27a54951957f6c1801000000000000c350000001cdb0ed311810e61036ac9255674101497850f5eee5e4320be07479c05473cbac010000000023c3460003ce4c4eac09fe317f365e45c00ffcf2e9639bc0fd792c10f72cdc173c4e5ed8791976a9149bdcb18911fa9faad6632ca43b81739082b0a19588ac00000000", txHex)

		cfdRet = CfdFreeMultisigSignHandle(handle, multiSignHandle)
		assert.Equal(t, (int)(KCfdSuccess), cfdRet)
	}

	if cfdRet != (int)(KCfdSuccess) {
		errStr, _ := CfdGoGetLastErrorMessage(handle)
		fmt.Print("[error message] " + errStr + "\n")
	}

	ret = CfdFreeHandle(handle)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestCfdAddMultisigSignConfidentialTx test done.\n")
}

// last test
/* comment out.
func TestFinalize(t *testing.T) {
	ret := CfdFinalize(false)
	assert.Equal(t, (int)(KCfdSuccess), ret)
	fmt.Print("TestFinalize test done.\n")
}
*/
