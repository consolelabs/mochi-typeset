// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package typeset

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Erc1155MetaData contains all meta data concerning the Erc1155 contract.
var Erc1155MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"DEFAULT_PAYLOAD_SIZE_LIMIT\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}],\"name\":\"FUNCTION_TYPE_SEND\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}],\"name\":\"FUNCTION_TYPE_SEND_BATCH\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"NO_EXTRA_GAS\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256[]\",\"name\":\"\",\"internalType\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"inputs\":[{\"type\":\"address[]\",\"name\":\"accounts\",\"internalType\":\"address[]\"},{\"type\":\"uint256[]\",\"name\":\"ids\",\"internalType\":\"uint256[]\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"delSigner\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"nativeFee\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"zroFee\",\"internalType\":\"uint256\"}],\"name\":\"estimateSendBatchFee\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_dstChainId\",\"internalType\":\"uint16\"},{\"type\":\"bytes\",\"name\":\"_toAddress\",\"internalType\":\"bytes\"},{\"type\":\"uint256[]\",\"name\":\"_tokenIds\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"_amounts\",\"internalType\":\"uint256[]\"},{\"type\":\"bool\",\"name\":\"_useZro\",\"internalType\":\"bool\"},{\"type\":\"bytes\",\"name\":\"_adapterParams\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"nativeFee\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"zroFee\",\"internalType\":\"uint256\"}],\"name\":\"estimateSendFee\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_dstChainId\",\"internalType\":\"uint16\"},{\"type\":\"bytes\",\"name\":\"_toAddress\",\"internalType\":\"bytes\"},{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"_useZro\",\"internalType\":\"bool\"},{\"type\":\"bytes\",\"name\":\"_adapterParams\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"exists\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"failedMessages\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"},{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"},{\"type\":\"uint64\",\"name\":\"\",\"internalType\":\"uint64\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"forceResumeReceive\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_srcChainId\",\"internalType\":\"uint16\"},{\"type\":\"bytes\",\"name\":\"_srcAddress\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"forwarder\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}],\"name\":\"getConfig\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_version\",\"internalType\":\"uint16\"},{\"type\":\"uint16\",\"name\":\"_chainId\",\"internalType\":\"uint16\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_configType\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getNonce\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}],\"name\":\"getTrustedRemoteAddress\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_remoteChainId\",\"internalType\":\"uint16\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"string\",\"name\":\"uri_\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"name_\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"symbol_\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"endpoint_\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"forwarder_\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isApprovedForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isSigner\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isTrustedRemote\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_srcChainId\",\"internalType\":\"uint16\"},{\"type\":\"bytes\",\"name\":\"_srcAddress\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractILayerZeroEndpointUpgradeable\"}],\"name\":\"lzEndpoint\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"lzReceive\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_srcChainId\",\"internalType\":\"uint16\"},{\"type\":\"bytes\",\"name\":\"_srcAddress\",\"internalType\":\"bytes\"},{\"type\":\"uint64\",\"name\":\"_nonce\",\"internalType\":\"uint64\"},{\"type\":\"bytes\",\"name\":\"_payload\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"minDstGasLookup\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"},{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"mint\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"none\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"signature\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"mint\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"none\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"signature\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"name\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"nonblockingLzReceive\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_srcChainId\",\"internalType\":\"uint16\"},{\"type\":\"bytes\",\"name\":\"_srcAddress\",\"internalType\":\"bytes\"},{\"type\":\"uint64\",\"name\":\"_nonce\",\"internalType\":\"uint64\"},{\"type\":\"bytes\",\"name\":\"_payload\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bytes4\",\"name\":\"\",\"internalType\":\"bytes4\"}],\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint256[]\",\"name\":\"\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"\",\"internalType\":\"uint256[]\"},{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bytes4\",\"name\":\"\",\"internalType\":\"bytes4\"}],\"name\":\"onERC1155Received\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"payloadSizeLimitLookup\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"precrime\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"retryMessage\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_srcChainId\",\"internalType\":\"uint16\"},{\"type\":\"bytes\",\"name\":\"_srcAddress\",\"internalType\":\"bytes\"},{\"type\":\"uint64\",\"name\":\"_nonce\",\"internalType\":\"uint64\"},{\"type\":\"bytes\",\"name\":\"_payload\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"safeBatchTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256[]\",\"name\":\"ids\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"amounts\",\"internalType\":\"uint256[]\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"safeTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"sendBatchFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"internalType\":\"address\"},{\"type\":\"uint16\",\"name\":\"_dstChainId\",\"internalType\":\"uint16\"},{\"type\":\"bytes\",\"name\":\"_toAddress\",\"internalType\":\"bytes\"},{\"type\":\"uint256[]\",\"name\":\"_tokenIds\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"_amounts\",\"internalType\":\"uint256[]\"},{\"type\":\"address\",\"name\":\"_refundAddress\",\"internalType\":\"addresspayable\"},{\"type\":\"address\",\"name\":\"_zroPaymentAddress\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"_adapterParams\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"sendFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"internalType\":\"address\"},{\"type\":\"uint16\",\"name\":\"_dstChainId\",\"internalType\":\"uint16\"},{\"type\":\"bytes\",\"name\":\"_toAddress\",\"internalType\":\"bytes\"},{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_refundAddress\",\"internalType\":\"addresspayable\"},{\"type\":\"address\",\"name\":\"_zroPaymentAddress\",\"internalType\":\"address\"},{\"type\":\"bytes\",\"name\":\"_adapterParams\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"approved\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setConfig\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_version\",\"internalType\":\"uint16\"},{\"type\":\"uint16\",\"name\":\"_chainId\",\"internalType\":\"uint16\"},{\"type\":\"uint256\",\"name\":\"_configType\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"_config\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setMinDstGas\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_dstChainId\",\"internalType\":\"uint16\"},{\"type\":\"uint16\",\"name\":\"_packetType\",\"internalType\":\"uint16\"},{\"type\":\"uint256\",\"name\":\"_minGas\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setNonce\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setPayloadSizeLimit\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_dstChainId\",\"internalType\":\"uint16\"},{\"type\":\"uint256\",\"name\":\"_size\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setPrecrime\",\"inputs\":[{\"type\":\"address\",\"name\":\"_precrime\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setReceiveVersion\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_version\",\"internalType\":\"uint16\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setSendVersion\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_version\",\"internalType\":\"uint16\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setSigner\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setTokenURI\",\"inputs\":[{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setTotalSupplyMax\",\"inputs\":[{\"type\":\"uint256[]\",\"name\":\"ids\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"amounts\",\"internalType\":\"uint256[]\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setTrustedRemote\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_srcChainId\",\"internalType\":\"uint16\"},{\"type\":\"bytes\",\"name\":\"_path\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setTrustedRemoteAddress\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_remoteChainId\",\"internalType\":\"uint16\"},{\"type\":\"bytes\",\"name\":\"_remoteAddress\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setUseCustomAdapterParams\",\"inputs\":[{\"type\":\"bool\",\"name\":\"_useCustomAdapterParams\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"supportsInterface\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"interfaceId\",\"internalType\":\"bytes4\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"symbol\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalSupply\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalSupplyMax\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"transferOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}],\"name\":\"trustedRemoteLookup\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"uri\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"useCustomAdapterParams\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"indexed\":true},{\"type\":\"address\",\"name\":\"operator\",\"indexed\":true},{\"type\":\"bool\",\"name\":\"approved\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"version\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageFailed\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_srcChainId\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"_srcAddress\",\"indexed\":false},{\"type\":\"uint64\",\"name\":\"_nonce\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"_payload\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"_reason\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceiveBatchFromChain\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_srcChainId\",\"indexed\":true},{\"type\":\"bytes\",\"name\":\"_srcAddress\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_toAddress\",\"indexed\":true},{\"type\":\"uint256[]\",\"name\":\"_tokenIds\",\"indexed\":false},{\"type\":\"uint256[]\",\"name\":\"_amounts\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceiveFromChain\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_srcChainId\",\"indexed\":true},{\"type\":\"bytes\",\"name\":\"_srcAddress\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_toAddress\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_tokenId\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"_amount\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RetryMessageSuccess\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_srcChainId\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"_srcAddress\",\"indexed\":false},{\"type\":\"uint64\",\"name\":\"_nonce\",\"indexed\":false},{\"type\":\"bytes32\",\"name\":\"_payloadHash\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendBatchToChain\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_dstChainId\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_from\",\"indexed\":true},{\"type\":\"bytes\",\"name\":\"_toAddress\",\"indexed\":true},{\"type\":\"uint256[]\",\"name\":\"_tokenIds\",\"indexed\":false},{\"type\":\"uint256[]\",\"name\":\"_amounts\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendToChain\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_dstChainId\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_from\",\"indexed\":true},{\"type\":\"bytes\",\"name\":\"_toAddress\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_tokenId\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"_amount\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMinDstGas\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_dstChainId\",\"indexed\":false},{\"type\":\"uint16\",\"name\":\"_type\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"_minDstGas\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetPrecrime\",\"inputs\":[{\"type\":\"address\",\"name\":\"precrime\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetTrustedRemote\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_remoteChainId\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"_path\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetTrustedRemoteAddress\",\"inputs\":[{\"type\":\"uint16\",\"name\":\"_remoteChainId\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"_remoteAddress\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetUseCustomAdapterParams\",\"inputs\":[{\"type\":\"bool\",\"name\":\"_useCustomAdapterParams\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferBatch\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"indexed\":true},{\"type\":\"address\",\"name\":\"from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"indexed\":true},{\"type\":\"uint256[]\",\"name\":\"ids\",\"indexed\":false},{\"type\":\"uint256[]\",\"name\":\"values\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferSingle\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"indexed\":true},{\"type\":\"address\",\"name\":\"from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"URI\",\"inputs\":[{\"type\":\"string\",\"name\":\"value\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"id\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UseSignature\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"none\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"signature\",\"indexed\":false}],\"anonymous\":false}]",
}

// Erc1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc1155MetaData.ABI instead.
var Erc1155ABI = Erc1155MetaData.ABI

// Erc1155 is an auto generated Go binding around an Ethereum contract.
type Erc1155 struct {
	Erc1155Caller     // Read-only binding to the contract
	Erc1155Transactor // Write-only binding to the contract
	Erc1155Filterer   // Log filterer for contract events
}

// Erc1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc1155Session struct {
	Contract     *Erc1155          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc1155CallerSession struct {
	Contract *Erc1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Erc1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc1155TransactorSession struct {
	Contract     *Erc1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Erc1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc1155Raw struct {
	Contract *Erc1155 // Generic contract binding to access the raw methods on
}

// Erc1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc1155CallerRaw struct {
	Contract *Erc1155Caller // Generic read-only contract binding to access the raw methods on
}

// Erc1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc1155TransactorRaw struct {
	Contract *Erc1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc1155 creates a new instance of Erc1155, bound to a specific deployed contract.
func NewErc1155(address common.Address, backend bind.ContractBackend) (*Erc1155, error) {
	contract, err := bindErc1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc1155{Erc1155Caller: Erc1155Caller{contract: contract}, Erc1155Transactor: Erc1155Transactor{contract: contract}, Erc1155Filterer: Erc1155Filterer{contract: contract}}, nil
}

// NewErc1155Caller creates a new read-only instance of Erc1155, bound to a specific deployed contract.
func NewErc1155Caller(address common.Address, caller bind.ContractCaller) (*Erc1155Caller, error) {
	contract, err := bindErc1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1155Caller{contract: contract}, nil
}

// NewErc1155Transactor creates a new write-only instance of Erc1155, bound to a specific deployed contract.
func NewErc1155Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc1155Transactor, error) {
	contract, err := bindErc1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1155Transactor{contract: contract}, nil
}

// NewErc1155Filterer creates a new log filterer instance of Erc1155, bound to a specific deployed contract.
func NewErc1155Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc1155Filterer, error) {
	contract, err := bindErc1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc1155Filterer{contract: contract}, nil
}

// bindErc1155 binds a generic wrapper to an already deployed contract.
func bindErc1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1155 *Erc1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc1155.Contract.Erc1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1155 *Erc1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1155.Contract.Erc1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1155 *Erc1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1155.Contract.Erc1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1155 *Erc1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1155 *Erc1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1155 *Erc1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1155.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_Erc1155 *Erc1155Caller) DEFAULTPAYLOADSIZELIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "DEFAULT_PAYLOAD_SIZE_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_Erc1155 *Erc1155Session) DEFAULTPAYLOADSIZELIMIT() (*big.Int, error) {
	return _Erc1155.Contract.DEFAULTPAYLOADSIZELIMIT(&_Erc1155.CallOpts)
}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_Erc1155 *Erc1155CallerSession) DEFAULTPAYLOADSIZELIMIT() (*big.Int, error) {
	return _Erc1155.Contract.DEFAULTPAYLOADSIZELIMIT(&_Erc1155.CallOpts)
}

// FUNCTIONTYPESEND is a free data retrieval call binding the contract method 0xaf3fb21c.
//
// Solidity: function FUNCTION_TYPE_SEND() view returns(uint16)
func (_Erc1155 *Erc1155Caller) FUNCTIONTYPESEND(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "FUNCTION_TYPE_SEND")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FUNCTIONTYPESEND is a free data retrieval call binding the contract method 0xaf3fb21c.
//
// Solidity: function FUNCTION_TYPE_SEND() view returns(uint16)
func (_Erc1155 *Erc1155Session) FUNCTIONTYPESEND() (uint16, error) {
	return _Erc1155.Contract.FUNCTIONTYPESEND(&_Erc1155.CallOpts)
}

// FUNCTIONTYPESEND is a free data retrieval call binding the contract method 0xaf3fb21c.
//
// Solidity: function FUNCTION_TYPE_SEND() view returns(uint16)
func (_Erc1155 *Erc1155CallerSession) FUNCTIONTYPESEND() (uint16, error) {
	return _Erc1155.Contract.FUNCTIONTYPESEND(&_Erc1155.CallOpts)
}

// FUNCTIONTYPESENDBATCH is a free data retrieval call binding the contract method 0x149e3e1f.
//
// Solidity: function FUNCTION_TYPE_SEND_BATCH() view returns(uint16)
func (_Erc1155 *Erc1155Caller) FUNCTIONTYPESENDBATCH(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "FUNCTION_TYPE_SEND_BATCH")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FUNCTIONTYPESENDBATCH is a free data retrieval call binding the contract method 0x149e3e1f.
//
// Solidity: function FUNCTION_TYPE_SEND_BATCH() view returns(uint16)
func (_Erc1155 *Erc1155Session) FUNCTIONTYPESENDBATCH() (uint16, error) {
	return _Erc1155.Contract.FUNCTIONTYPESENDBATCH(&_Erc1155.CallOpts)
}

// FUNCTIONTYPESENDBATCH is a free data retrieval call binding the contract method 0x149e3e1f.
//
// Solidity: function FUNCTION_TYPE_SEND_BATCH() view returns(uint16)
func (_Erc1155 *Erc1155CallerSession) FUNCTIONTYPESENDBATCH() (uint16, error) {
	return _Erc1155.Contract.FUNCTIONTYPESENDBATCH(&_Erc1155.CallOpts)
}

// NOEXTRAGAS is a free data retrieval call binding the contract method 0x44770515.
//
// Solidity: function NO_EXTRA_GAS() view returns(uint256)
func (_Erc1155 *Erc1155Caller) NOEXTRAGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "NO_EXTRA_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NOEXTRAGAS is a free data retrieval call binding the contract method 0x44770515.
//
// Solidity: function NO_EXTRA_GAS() view returns(uint256)
func (_Erc1155 *Erc1155Session) NOEXTRAGAS() (*big.Int, error) {
	return _Erc1155.Contract.NOEXTRAGAS(&_Erc1155.CallOpts)
}

// NOEXTRAGAS is a free data retrieval call binding the contract method 0x44770515.
//
// Solidity: function NO_EXTRA_GAS() view returns(uint256)
func (_Erc1155 *Erc1155CallerSession) NOEXTRAGAS() (*big.Int, error) {
	return _Erc1155.Contract.NOEXTRAGAS(&_Erc1155.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Erc1155.Contract.BalanceOf(&_Erc1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Erc1155.Contract.BalanceOf(&_Erc1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Erc1155 *Erc1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Erc1155 *Erc1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Erc1155.Contract.BalanceOfBatch(&_Erc1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Erc1155 *Erc1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Erc1155.Contract.BalanceOfBatch(&_Erc1155.CallOpts, accounts, ids)
}

// EstimateSendBatchFee is a free data retrieval call binding the contract method 0xb2535663.
//
// Solidity: function estimateSendBatchFee(uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, uint256[] _amounts, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Erc1155 *Erc1155Caller) EstimateSendBatchFee(opts *bind.CallOpts, _dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _amounts []*big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "estimateSendBatchFee", _dstChainId, _toAddress, _tokenIds, _amounts, _useZro, _adapterParams)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateSendBatchFee is a free data retrieval call binding the contract method 0xb2535663.
//
// Solidity: function estimateSendBatchFee(uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, uint256[] _amounts, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Erc1155 *Erc1155Session) EstimateSendBatchFee(_dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _amounts []*big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Erc1155.Contract.EstimateSendBatchFee(&_Erc1155.CallOpts, _dstChainId, _toAddress, _tokenIds, _amounts, _useZro, _adapterParams)
}

// EstimateSendBatchFee is a free data retrieval call binding the contract method 0xb2535663.
//
// Solidity: function estimateSendBatchFee(uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, uint256[] _amounts, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Erc1155 *Erc1155CallerSession) EstimateSendBatchFee(_dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _amounts []*big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Erc1155.Contract.EstimateSendBatchFee(&_Erc1155.CallOpts, _dstChainId, _toAddress, _tokenIds, _amounts, _useZro, _adapterParams)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0x8608e5f8.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, uint256 _amount, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Erc1155 *Erc1155Caller) EstimateSendFee(opts *bind.CallOpts, _dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _amount *big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "estimateSendFee", _dstChainId, _toAddress, _tokenId, _amount, _useZro, _adapterParams)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateSendFee is a free data retrieval call binding the contract method 0x8608e5f8.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, uint256 _amount, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Erc1155 *Erc1155Session) EstimateSendFee(_dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _amount *big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Erc1155.Contract.EstimateSendFee(&_Erc1155.CallOpts, _dstChainId, _toAddress, _tokenId, _amount, _useZro, _adapterParams)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0x8608e5f8.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, uint256 _amount, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Erc1155 *Erc1155CallerSession) EstimateSendFee(_dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _amount *big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Erc1155.Contract.EstimateSendFee(&_Erc1155.CallOpts, _dstChainId, _toAddress, _tokenId, _amount, _useZro, _adapterParams)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Erc1155 *Erc1155Caller) Exists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "exists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Erc1155 *Erc1155Session) Exists(id *big.Int) (bool, error) {
	return _Erc1155.Contract.Exists(&_Erc1155.CallOpts, id)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Erc1155 *Erc1155CallerSession) Exists(id *big.Int) (bool, error) {
	return _Erc1155.Contract.Exists(&_Erc1155.CallOpts, id)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Erc1155 *Erc1155Caller) FailedMessages(opts *bind.CallOpts, arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "failedMessages", arg0, arg1, arg2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Erc1155 *Erc1155Session) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Erc1155.Contract.FailedMessages(&_Erc1155.CallOpts, arg0, arg1, arg2)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Erc1155 *Erc1155CallerSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Erc1155.Contract.FailedMessages(&_Erc1155.CallOpts, arg0, arg1, arg2)
}

// Forwarder is a free data retrieval call binding the contract method 0xf645d4f9.
//
// Solidity: function forwarder() view returns(address)
func (_Erc1155 *Erc1155Caller) Forwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "forwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Forwarder is a free data retrieval call binding the contract method 0xf645d4f9.
//
// Solidity: function forwarder() view returns(address)
func (_Erc1155 *Erc1155Session) Forwarder() (common.Address, error) {
	return _Erc1155.Contract.Forwarder(&_Erc1155.CallOpts)
}

// Forwarder is a free data retrieval call binding the contract method 0xf645d4f9.
//
// Solidity: function forwarder() view returns(address)
func (_Erc1155 *Erc1155CallerSession) Forwarder() (common.Address, error) {
	return _Erc1155.Contract.Forwarder(&_Erc1155.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Erc1155 *Erc1155Caller) GetConfig(opts *bind.CallOpts, _version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "getConfig", _version, _chainId, arg2, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Erc1155 *Erc1155Session) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Erc1155.Contract.GetConfig(&_Erc1155.CallOpts, _version, _chainId, arg2, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Erc1155 *Erc1155CallerSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Erc1155.Contract.GetConfig(&_Erc1155.CallOpts, _version, _chainId, arg2, _configType)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint256)
func (_Erc1155 *Erc1155Caller) GetNonce(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "getNonce", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint256)
func (_Erc1155 *Erc1155Session) GetNonce(account common.Address) (*big.Int, error) {
	return _Erc1155.Contract.GetNonce(&_Erc1155.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint256)
func (_Erc1155 *Erc1155CallerSession) GetNonce(account common.Address) (*big.Int, error) {
	return _Erc1155.Contract.GetNonce(&_Erc1155.CallOpts, account)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Erc1155 *Erc1155Caller) GetTrustedRemoteAddress(opts *bind.CallOpts, _remoteChainId uint16) ([]byte, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "getTrustedRemoteAddress", _remoteChainId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Erc1155 *Erc1155Session) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _Erc1155.Contract.GetTrustedRemoteAddress(&_Erc1155.CallOpts, _remoteChainId)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Erc1155 *Erc1155CallerSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _Erc1155.Contract.GetTrustedRemoteAddress(&_Erc1155.CallOpts, _remoteChainId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Erc1155 *Erc1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Erc1155 *Erc1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Erc1155.Contract.IsApprovedForAll(&_Erc1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Erc1155 *Erc1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Erc1155.Contract.IsApprovedForAll(&_Erc1155.CallOpts, account, operator)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address account) view returns(bool)
func (_Erc1155 *Erc1155Caller) IsSigner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "isSigner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address account) view returns(bool)
func (_Erc1155 *Erc1155Session) IsSigner(account common.Address) (bool, error) {
	return _Erc1155.Contract.IsSigner(&_Erc1155.CallOpts, account)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address account) view returns(bool)
func (_Erc1155 *Erc1155CallerSession) IsSigner(account common.Address) (bool, error) {
	return _Erc1155.Contract.IsSigner(&_Erc1155.CallOpts, account)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Erc1155 *Erc1155Caller) IsTrustedRemote(opts *bind.CallOpts, _srcChainId uint16, _srcAddress []byte) (bool, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "isTrustedRemote", _srcChainId, _srcAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Erc1155 *Erc1155Session) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Erc1155.Contract.IsTrustedRemote(&_Erc1155.CallOpts, _srcChainId, _srcAddress)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Erc1155 *Erc1155CallerSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Erc1155.Contract.IsTrustedRemote(&_Erc1155.CallOpts, _srcChainId, _srcAddress)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Erc1155 *Erc1155Caller) LzEndpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "lzEndpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Erc1155 *Erc1155Session) LzEndpoint() (common.Address, error) {
	return _Erc1155.Contract.LzEndpoint(&_Erc1155.CallOpts)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Erc1155 *Erc1155CallerSession) LzEndpoint() (common.Address, error) {
	return _Erc1155.Contract.LzEndpoint(&_Erc1155.CallOpts)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Erc1155 *Erc1155Caller) MinDstGasLookup(opts *bind.CallOpts, arg0 uint16, arg1 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "minDstGasLookup", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Erc1155 *Erc1155Session) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _Erc1155.Contract.MinDstGasLookup(&_Erc1155.CallOpts, arg0, arg1)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Erc1155 *Erc1155CallerSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _Erc1155.Contract.MinDstGasLookup(&_Erc1155.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc1155 *Erc1155Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc1155 *Erc1155Session) Name() (string, error) {
	return _Erc1155.Contract.Name(&_Erc1155.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc1155 *Erc1155CallerSession) Name() (string, error) {
	return _Erc1155.Contract.Name(&_Erc1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc1155 *Erc1155Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc1155 *Erc1155Session) Owner() (common.Address, error) {
	return _Erc1155.Contract.Owner(&_Erc1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc1155 *Erc1155CallerSession) Owner() (common.Address, error) {
	return _Erc1155.Contract.Owner(&_Erc1155.CallOpts)
}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_Erc1155 *Erc1155Caller) PayloadSizeLimitLookup(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "payloadSizeLimitLookup", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_Erc1155 *Erc1155Session) PayloadSizeLimitLookup(arg0 uint16) (*big.Int, error) {
	return _Erc1155.Contract.PayloadSizeLimitLookup(&_Erc1155.CallOpts, arg0)
}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_Erc1155 *Erc1155CallerSession) PayloadSizeLimitLookup(arg0 uint16) (*big.Int, error) {
	return _Erc1155.Contract.PayloadSizeLimitLookup(&_Erc1155.CallOpts, arg0)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Erc1155 *Erc1155Caller) Precrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "precrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Erc1155 *Erc1155Session) Precrime() (common.Address, error) {
	return _Erc1155.Contract.Precrime(&_Erc1155.CallOpts)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Erc1155 *Erc1155CallerSession) Precrime() (common.Address, error) {
	return _Erc1155.Contract.Precrime(&_Erc1155.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc1155 *Erc1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc1155 *Erc1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc1155.Contract.SupportsInterface(&_Erc1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc1155 *Erc1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc1155.Contract.SupportsInterface(&_Erc1155.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc1155 *Erc1155Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc1155 *Erc1155Session) Symbol() (string, error) {
	return _Erc1155.Contract.Symbol(&_Erc1155.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc1155 *Erc1155CallerSession) Symbol() (string, error) {
	return _Erc1155.Contract.Symbol(&_Erc1155.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155Caller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "totalSupply", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155Session) TotalSupply(id *big.Int) (*big.Int, error) {
	return _Erc1155.Contract.TotalSupply(&_Erc1155.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155CallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _Erc1155.Contract.TotalSupply(&_Erc1155.CallOpts, id)
}

// TotalSupplyMax is a free data retrieval call binding the contract method 0x844e4244.
//
// Solidity: function totalSupplyMax(uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155Caller) TotalSupplyMax(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "totalSupplyMax", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyMax is a free data retrieval call binding the contract method 0x844e4244.
//
// Solidity: function totalSupplyMax(uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155Session) TotalSupplyMax(id *big.Int) (*big.Int, error) {
	return _Erc1155.Contract.TotalSupplyMax(&_Erc1155.CallOpts, id)
}

// TotalSupplyMax is a free data retrieval call binding the contract method 0x844e4244.
//
// Solidity: function totalSupplyMax(uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155CallerSession) TotalSupplyMax(id *big.Int) (*big.Int, error) {
	return _Erc1155.Contract.TotalSupplyMax(&_Erc1155.CallOpts, id)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Erc1155 *Erc1155Caller) TrustedRemoteLookup(opts *bind.CallOpts, arg0 uint16) ([]byte, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "trustedRemoteLookup", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Erc1155 *Erc1155Session) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Erc1155.Contract.TrustedRemoteLookup(&_Erc1155.CallOpts, arg0)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Erc1155 *Erc1155CallerSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Erc1155.Contract.TrustedRemoteLookup(&_Erc1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Erc1155 *Erc1155Caller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Erc1155 *Erc1155Session) Uri(arg0 *big.Int) (string, error) {
	return _Erc1155.Contract.Uri(&_Erc1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Erc1155 *Erc1155CallerSession) Uri(arg0 *big.Int) (string, error) {
	return _Erc1155.Contract.Uri(&_Erc1155.CallOpts, arg0)
}

// UseCustomAdapterParams is a free data retrieval call binding the contract method 0xed629c5c.
//
// Solidity: function useCustomAdapterParams() view returns(bool)
func (_Erc1155 *Erc1155Caller) UseCustomAdapterParams(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "useCustomAdapterParams")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseCustomAdapterParams is a free data retrieval call binding the contract method 0xed629c5c.
//
// Solidity: function useCustomAdapterParams() view returns(bool)
func (_Erc1155 *Erc1155Session) UseCustomAdapterParams() (bool, error) {
	return _Erc1155.Contract.UseCustomAdapterParams(&_Erc1155.CallOpts)
}

// UseCustomAdapterParams is a free data retrieval call binding the contract method 0xed629c5c.
//
// Solidity: function useCustomAdapterParams() view returns(bool)
func (_Erc1155 *Erc1155CallerSession) UseCustomAdapterParams() (bool, error) {
	return _Erc1155.Contract.UseCustomAdapterParams(&_Erc1155.CallOpts)
}

// DelSigner is a paid mutator transaction binding the contract method 0x883b524f.
//
// Solidity: function delSigner(address account) returns()
func (_Erc1155 *Erc1155Transactor) DelSigner(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "delSigner", account)
}

// DelSigner is a paid mutator transaction binding the contract method 0x883b524f.
//
// Solidity: function delSigner(address account) returns()
func (_Erc1155 *Erc1155Session) DelSigner(account common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.DelSigner(&_Erc1155.TransactOpts, account)
}

// DelSigner is a paid mutator transaction binding the contract method 0x883b524f.
//
// Solidity: function delSigner(address account) returns()
func (_Erc1155 *Erc1155TransactorSession) DelSigner(account common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.DelSigner(&_Erc1155.TransactOpts, account)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Erc1155 *Erc1155Transactor) ForceResumeReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "forceResumeReceive", _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Erc1155 *Erc1155Session) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.ForceResumeReceive(&_Erc1155.TransactOpts, _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Erc1155 *Erc1155TransactorSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.ForceResumeReceive(&_Erc1155.TransactOpts, _srcChainId, _srcAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xd6d0faee.
//
// Solidity: function initialize(string uri_, string name_, string symbol_, address endpoint_, address forwarder_) returns()
func (_Erc1155 *Erc1155Transactor) Initialize(opts *bind.TransactOpts, uri_ string, name_ string, symbol_ string, endpoint_ common.Address, forwarder_ common.Address) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "initialize", uri_, name_, symbol_, endpoint_, forwarder_)
}

// Initialize is a paid mutator transaction binding the contract method 0xd6d0faee.
//
// Solidity: function initialize(string uri_, string name_, string symbol_, address endpoint_, address forwarder_) returns()
func (_Erc1155 *Erc1155Session) Initialize(uri_ string, name_ string, symbol_ string, endpoint_ common.Address, forwarder_ common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.Initialize(&_Erc1155.TransactOpts, uri_, name_, symbol_, endpoint_, forwarder_)
}

// Initialize is a paid mutator transaction binding the contract method 0xd6d0faee.
//
// Solidity: function initialize(string uri_, string name_, string symbol_, address endpoint_, address forwarder_) returns()
func (_Erc1155 *Erc1155TransactorSession) Initialize(uri_ string, name_ string, symbol_ string, endpoint_ common.Address, forwarder_ common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.Initialize(&_Erc1155.TransactOpts, uri_, name_, symbol_, endpoint_, forwarder_)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Erc1155 *Erc1155Transactor) LzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "lzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Erc1155 *Erc1155Session) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.LzReceive(&_Erc1155.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Erc1155 *Erc1155TransactorSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.LzReceive(&_Erc1155.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// Mint is a paid mutator transaction binding the contract method 0x4a9eee69.
//
// Solidity: function mint(uint256 id, uint256 none, uint256 amount, bytes signature) returns()
func (_Erc1155 *Erc1155Transactor) Mint(opts *bind.TransactOpts, id *big.Int, none *big.Int, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "mint", id, none, amount, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x4a9eee69.
//
// Solidity: function mint(uint256 id, uint256 none, uint256 amount, bytes signature) returns()
func (_Erc1155 *Erc1155Session) Mint(id *big.Int, none *big.Int, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.Mint(&_Erc1155.TransactOpts, id, none, amount, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x4a9eee69.
//
// Solidity: function mint(uint256 id, uint256 none, uint256 amount, bytes signature) returns()
func (_Erc1155 *Erc1155TransactorSession) Mint(id *big.Int, none *big.Int, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.Mint(&_Erc1155.TransactOpts, id, none, amount, signature)
}

// Mint0 is a paid mutator transaction binding the contract method 0xb49c99b8.
//
// Solidity: function mint(address to, uint256 id, uint256 none, uint256 amount, bytes signature) returns()
func (_Erc1155 *Erc1155Transactor) Mint0(opts *bind.TransactOpts, to common.Address, id *big.Int, none *big.Int, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "mint0", to, id, none, amount, signature)
}

// Mint0 is a paid mutator transaction binding the contract method 0xb49c99b8.
//
// Solidity: function mint(address to, uint256 id, uint256 none, uint256 amount, bytes signature) returns()
func (_Erc1155 *Erc1155Session) Mint0(to common.Address, id *big.Int, none *big.Int, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.Mint0(&_Erc1155.TransactOpts, to, id, none, amount, signature)
}

// Mint0 is a paid mutator transaction binding the contract method 0xb49c99b8.
//
// Solidity: function mint(address to, uint256 id, uint256 none, uint256 amount, bytes signature) returns()
func (_Erc1155 *Erc1155TransactorSession) Mint0(to common.Address, id *big.Int, none *big.Int, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.Mint0(&_Erc1155.TransactOpts, to, id, none, amount, signature)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Erc1155 *Erc1155Transactor) NonblockingLzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "nonblockingLzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Erc1155 *Erc1155Session) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.NonblockingLzReceive(&_Erc1155.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Erc1155 *Erc1155TransactorSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.NonblockingLzReceive(&_Erc1155.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Erc1155 *Erc1155Transactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Erc1155 *Erc1155Session) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.OnERC1155BatchReceived(&_Erc1155.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Erc1155 *Erc1155TransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.OnERC1155BatchReceived(&_Erc1155.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Erc1155 *Erc1155Transactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Erc1155 *Erc1155Session) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.OnERC1155Received(&_Erc1155.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Erc1155 *Erc1155TransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.OnERC1155Received(&_Erc1155.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc1155 *Erc1155Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc1155 *Erc1155Session) RenounceOwnership() (*types.Transaction, error) {
	return _Erc1155.Contract.RenounceOwnership(&_Erc1155.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc1155 *Erc1155TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Erc1155.Contract.RenounceOwnership(&_Erc1155.TransactOpts)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Erc1155 *Erc1155Transactor) RetryMessage(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "retryMessage", _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Erc1155 *Erc1155Session) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.RetryMessage(&_Erc1155.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Erc1155 *Erc1155TransactorSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.RetryMessage(&_Erc1155.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Erc1155 *Erc1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Erc1155 *Erc1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SafeBatchTransferFrom(&_Erc1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Erc1155 *Erc1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SafeBatchTransferFrom(&_Erc1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Erc1155 *Erc1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Erc1155 *Erc1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SafeTransferFrom(&_Erc1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Erc1155 *Erc1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SafeTransferFrom(&_Erc1155.TransactOpts, from, to, id, amount, data)
}

// SendBatchFrom is a paid mutator transaction binding the contract method 0x4ab4e687.
//
// Solidity: function sendBatchFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, uint256[] _amounts, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Erc1155 *Erc1155Transactor) SendBatchFrom(opts *bind.TransactOpts, _from common.Address, _dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _amounts []*big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "sendBatchFrom", _from, _dstChainId, _toAddress, _tokenIds, _amounts, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendBatchFrom is a paid mutator transaction binding the contract method 0x4ab4e687.
//
// Solidity: function sendBatchFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, uint256[] _amounts, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Erc1155 *Erc1155Session) SendBatchFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _amounts []*big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SendBatchFrom(&_Erc1155.TransactOpts, _from, _dstChainId, _toAddress, _tokenIds, _amounts, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendBatchFrom is a paid mutator transaction binding the contract method 0x4ab4e687.
//
// Solidity: function sendBatchFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, uint256[] _amounts, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Erc1155 *Erc1155TransactorSession) SendBatchFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _amounts []*big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SendBatchFrom(&_Erc1155.TransactOpts, _from, _dstChainId, _toAddress, _tokenIds, _amounts, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendFrom is a paid mutator transaction binding the contract method 0x4db8226a.
//
// Solidity: function sendFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, uint256 _amount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Erc1155 *Erc1155Transactor) SendFrom(opts *bind.TransactOpts, _from common.Address, _dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _amount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "sendFrom", _from, _dstChainId, _toAddress, _tokenId, _amount, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendFrom is a paid mutator transaction binding the contract method 0x4db8226a.
//
// Solidity: function sendFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, uint256 _amount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Erc1155 *Erc1155Session) SendFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _amount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SendFrom(&_Erc1155.TransactOpts, _from, _dstChainId, _toAddress, _tokenId, _amount, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendFrom is a paid mutator transaction binding the contract method 0x4db8226a.
//
// Solidity: function sendFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, uint256 _amount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Erc1155 *Erc1155TransactorSession) SendFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _amount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SendFrom(&_Erc1155.TransactOpts, _from, _dstChainId, _toAddress, _tokenId, _amount, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc1155 *Erc1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc1155 *Erc1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc1155.Contract.SetApprovalForAll(&_Erc1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc1155 *Erc1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc1155.Contract.SetApprovalForAll(&_Erc1155.TransactOpts, operator, approved)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Erc1155 *Erc1155Transactor) SetConfig(opts *bind.TransactOpts, _version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setConfig", _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Erc1155 *Erc1155Session) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SetConfig(&_Erc1155.TransactOpts, _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Erc1155 *Erc1155TransactorSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SetConfig(&_Erc1155.TransactOpts, _version, _chainId, _configType, _config)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Erc1155 *Erc1155Transactor) SetMinDstGas(opts *bind.TransactOpts, _dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setMinDstGas", _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Erc1155 *Erc1155Session) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.SetMinDstGas(&_Erc1155.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Erc1155 *Erc1155TransactorSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.SetMinDstGas(&_Erc1155.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetNonce is a paid mutator transaction binding the contract method 0xcc8b84bf.
//
// Solidity: function setNonce(address account) returns()
func (_Erc1155 *Erc1155Transactor) SetNonce(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setNonce", account)
}

// SetNonce is a paid mutator transaction binding the contract method 0xcc8b84bf.
//
// Solidity: function setNonce(address account) returns()
func (_Erc1155 *Erc1155Session) SetNonce(account common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.SetNonce(&_Erc1155.TransactOpts, account)
}

// SetNonce is a paid mutator transaction binding the contract method 0xcc8b84bf.
//
// Solidity: function setNonce(address account) returns()
func (_Erc1155 *Erc1155TransactorSession) SetNonce(account common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.SetNonce(&_Erc1155.TransactOpts, account)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_Erc1155 *Erc1155Transactor) SetPayloadSizeLimit(opts *bind.TransactOpts, _dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setPayloadSizeLimit", _dstChainId, _size)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_Erc1155 *Erc1155Session) SetPayloadSizeLimit(_dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.SetPayloadSizeLimit(&_Erc1155.TransactOpts, _dstChainId, _size)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_Erc1155 *Erc1155TransactorSession) SetPayloadSizeLimit(_dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.SetPayloadSizeLimit(&_Erc1155.TransactOpts, _dstChainId, _size)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Erc1155 *Erc1155Transactor) SetPrecrime(opts *bind.TransactOpts, _precrime common.Address) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setPrecrime", _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Erc1155 *Erc1155Session) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.SetPrecrime(&_Erc1155.TransactOpts, _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Erc1155 *Erc1155TransactorSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.SetPrecrime(&_Erc1155.TransactOpts, _precrime)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Erc1155 *Erc1155Transactor) SetReceiveVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setReceiveVersion", _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Erc1155 *Erc1155Session) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Erc1155.Contract.SetReceiveVersion(&_Erc1155.TransactOpts, _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Erc1155 *Erc1155TransactorSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Erc1155.Contract.SetReceiveVersion(&_Erc1155.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Erc1155 *Erc1155Transactor) SetSendVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setSendVersion", _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Erc1155 *Erc1155Session) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Erc1155.Contract.SetSendVersion(&_Erc1155.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Erc1155 *Erc1155TransactorSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Erc1155.Contract.SetSendVersion(&_Erc1155.TransactOpts, _version)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address account) returns()
func (_Erc1155 *Erc1155Transactor) SetSigner(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setSigner", account)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address account) returns()
func (_Erc1155 *Erc1155Session) SetSigner(account common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.SetSigner(&_Erc1155.TransactOpts, account)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address account) returns()
func (_Erc1155 *Erc1155TransactorSession) SetSigner(account common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.SetSigner(&_Erc1155.TransactOpts, account)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0xe0df5b6f.
//
// Solidity: function setTokenURI(string uri) returns()
func (_Erc1155 *Erc1155Transactor) SetTokenURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setTokenURI", uri)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0xe0df5b6f.
//
// Solidity: function setTokenURI(string uri) returns()
func (_Erc1155 *Erc1155Session) SetTokenURI(uri string) (*types.Transaction, error) {
	return _Erc1155.Contract.SetTokenURI(&_Erc1155.TransactOpts, uri)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0xe0df5b6f.
//
// Solidity: function setTokenURI(string uri) returns()
func (_Erc1155 *Erc1155TransactorSession) SetTokenURI(uri string) (*types.Transaction, error) {
	return _Erc1155.Contract.SetTokenURI(&_Erc1155.TransactOpts, uri)
}

// SetTotalSupplyMax is a paid mutator transaction binding the contract method 0x26ff1979.
//
// Solidity: function setTotalSupplyMax(uint256[] ids, uint256[] amounts) returns()
func (_Erc1155 *Erc1155Transactor) SetTotalSupplyMax(opts *bind.TransactOpts, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setTotalSupplyMax", ids, amounts)
}

// SetTotalSupplyMax is a paid mutator transaction binding the contract method 0x26ff1979.
//
// Solidity: function setTotalSupplyMax(uint256[] ids, uint256[] amounts) returns()
func (_Erc1155 *Erc1155Session) SetTotalSupplyMax(ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.SetTotalSupplyMax(&_Erc1155.TransactOpts, ids, amounts)
}

// SetTotalSupplyMax is a paid mutator transaction binding the contract method 0x26ff1979.
//
// Solidity: function setTotalSupplyMax(uint256[] ids, uint256[] amounts) returns()
func (_Erc1155 *Erc1155TransactorSession) SetTotalSupplyMax(ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Erc1155.Contract.SetTotalSupplyMax(&_Erc1155.TransactOpts, ids, amounts)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _srcChainId, bytes _path) returns()
func (_Erc1155 *Erc1155Transactor) SetTrustedRemote(opts *bind.TransactOpts, _srcChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setTrustedRemote", _srcChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _srcChainId, bytes _path) returns()
func (_Erc1155 *Erc1155Session) SetTrustedRemote(_srcChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SetTrustedRemote(&_Erc1155.TransactOpts, _srcChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _srcChainId, bytes _path) returns()
func (_Erc1155 *Erc1155TransactorSession) SetTrustedRemote(_srcChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SetTrustedRemote(&_Erc1155.TransactOpts, _srcChainId, _path)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Erc1155 *Erc1155Transactor) SetTrustedRemoteAddress(opts *bind.TransactOpts, _remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setTrustedRemoteAddress", _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Erc1155 *Erc1155Session) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SetTrustedRemoteAddress(&_Erc1155.TransactOpts, _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Erc1155 *Erc1155TransactorSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SetTrustedRemoteAddress(&_Erc1155.TransactOpts, _remoteChainId, _remoteAddress)
}

// SetUseCustomAdapterParams is a paid mutator transaction binding the contract method 0xeab45d9c.
//
// Solidity: function setUseCustomAdapterParams(bool _useCustomAdapterParams) returns()
func (_Erc1155 *Erc1155Transactor) SetUseCustomAdapterParams(opts *bind.TransactOpts, _useCustomAdapterParams bool) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setUseCustomAdapterParams", _useCustomAdapterParams)
}

// SetUseCustomAdapterParams is a paid mutator transaction binding the contract method 0xeab45d9c.
//
// Solidity: function setUseCustomAdapterParams(bool _useCustomAdapterParams) returns()
func (_Erc1155 *Erc1155Session) SetUseCustomAdapterParams(_useCustomAdapterParams bool) (*types.Transaction, error) {
	return _Erc1155.Contract.SetUseCustomAdapterParams(&_Erc1155.TransactOpts, _useCustomAdapterParams)
}

// SetUseCustomAdapterParams is a paid mutator transaction binding the contract method 0xeab45d9c.
//
// Solidity: function setUseCustomAdapterParams(bool _useCustomAdapterParams) returns()
func (_Erc1155 *Erc1155TransactorSession) SetUseCustomAdapterParams(_useCustomAdapterParams bool) (*types.Transaction, error) {
	return _Erc1155.Contract.SetUseCustomAdapterParams(&_Erc1155.TransactOpts, _useCustomAdapterParams)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc1155 *Erc1155Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc1155 *Erc1155Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.TransferOwnership(&_Erc1155.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc1155 *Erc1155TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc1155.Contract.TransferOwnership(&_Erc1155.TransactOpts, newOwner)
}

// Erc1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Erc1155 contract.
type Erc1155ApprovalForAllIterator struct {
	Event *Erc1155ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155ApprovalForAll represents a ApprovalForAll event raised by the Erc1155 contract.
type Erc1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Erc1155 *Erc1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*Erc1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155ApprovalForAllIterator{contract: _Erc1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Erc1155 *Erc1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Erc1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155ApprovalForAll)
				if err := _Erc1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Erc1155 *Erc1155Filterer) ParseApprovalForAll(log types.Log) (*Erc1155ApprovalForAll, error) {
	event := new(Erc1155ApprovalForAll)
	if err := _Erc1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Erc1155 contract.
type Erc1155InitializedIterator struct {
	Event *Erc1155Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155Initialized represents a Initialized event raised by the Erc1155 contract.
type Erc1155Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Erc1155 *Erc1155Filterer) FilterInitialized(opts *bind.FilterOpts) (*Erc1155InitializedIterator, error) {

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Erc1155InitializedIterator{contract: _Erc1155.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Erc1155 *Erc1155Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Erc1155Initialized) (event.Subscription, error) {

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155Initialized)
				if err := _Erc1155.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Erc1155 *Erc1155Filterer) ParseInitialized(log types.Log) (*Erc1155Initialized, error) {
	event := new(Erc1155Initialized)
	if err := _Erc1155.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155MessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the Erc1155 contract.
type Erc1155MessageFailedIterator struct {
	Event *Erc1155MessageFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155MessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155MessageFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155MessageFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155MessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155MessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155MessageFailed represents a MessageFailed event raised by the Erc1155 contract.
type Erc1155MessageFailed struct {
	SrcChainId uint16
	SrcAddress []byte
	Nonce      uint64
	Payload    []byte
	Reason     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageFailed is a free log retrieval operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Erc1155 *Erc1155Filterer) FilterMessageFailed(opts *bind.FilterOpts) (*Erc1155MessageFailedIterator, error) {

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return &Erc1155MessageFailedIterator{contract: _Erc1155.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Erc1155 *Erc1155Filterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *Erc1155MessageFailed) (event.Subscription, error) {

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155MessageFailed)
				if err := _Erc1155.contract.UnpackLog(event, "MessageFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageFailed is a log parse operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Erc1155 *Erc1155Filterer) ParseMessageFailed(log types.Log) (*Erc1155MessageFailed, error) {
	event := new(Erc1155MessageFailed)
	if err := _Erc1155.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Erc1155 contract.
type Erc1155OwnershipTransferredIterator struct {
	Event *Erc1155OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155OwnershipTransferred represents a OwnershipTransferred event raised by the Erc1155 contract.
type Erc1155OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc1155 *Erc1155Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Erc1155OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155OwnershipTransferredIterator{contract: _Erc1155.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc1155 *Erc1155Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Erc1155OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155OwnershipTransferred)
				if err := _Erc1155.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc1155 *Erc1155Filterer) ParseOwnershipTransferred(log types.Log) (*Erc1155OwnershipTransferred, error) {
	event := new(Erc1155OwnershipTransferred)
	if err := _Erc1155.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155ReceiveBatchFromChainIterator is returned from FilterReceiveBatchFromChain and is used to iterate over the raw logs and unpacked data for ReceiveBatchFromChain events raised by the Erc1155 contract.
type Erc1155ReceiveBatchFromChainIterator struct {
	Event *Erc1155ReceiveBatchFromChain // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155ReceiveBatchFromChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155ReceiveBatchFromChain)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155ReceiveBatchFromChain)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155ReceiveBatchFromChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155ReceiveBatchFromChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155ReceiveBatchFromChain represents a ReceiveBatchFromChain event raised by the Erc1155 contract.
type Erc1155ReceiveBatchFromChain struct {
	SrcChainId uint16
	SrcAddress common.Hash
	ToAddress  common.Address
	TokenIds   []*big.Int
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReceiveBatchFromChain is a free log retrieval operation binding the contract event 0x1ae08edbbcd7baa8d064835de8593ce16b313414525ac89534e349f4da7926e4.
//
// Solidity: event ReceiveBatchFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256[] _tokenIds, uint256[] _amounts)
func (_Erc1155 *Erc1155Filterer) FilterReceiveBatchFromChain(opts *bind.FilterOpts, _srcChainId []uint16, _srcAddress [][]byte, _toAddress []common.Address) (*Erc1155ReceiveBatchFromChainIterator, error) {

	var _srcChainIdRule []interface{}
	for _, _srcChainIdItem := range _srcChainId {
		_srcChainIdRule = append(_srcChainIdRule, _srcChainIdItem)
	}
	var _srcAddressRule []interface{}
	for _, _srcAddressItem := range _srcAddress {
		_srcAddressRule = append(_srcAddressRule, _srcAddressItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "ReceiveBatchFromChain", _srcChainIdRule, _srcAddressRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155ReceiveBatchFromChainIterator{contract: _Erc1155.contract, event: "ReceiveBatchFromChain", logs: logs, sub: sub}, nil
}

// WatchReceiveBatchFromChain is a free log subscription operation binding the contract event 0x1ae08edbbcd7baa8d064835de8593ce16b313414525ac89534e349f4da7926e4.
//
// Solidity: event ReceiveBatchFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256[] _tokenIds, uint256[] _amounts)
func (_Erc1155 *Erc1155Filterer) WatchReceiveBatchFromChain(opts *bind.WatchOpts, sink chan<- *Erc1155ReceiveBatchFromChain, _srcChainId []uint16, _srcAddress [][]byte, _toAddress []common.Address) (event.Subscription, error) {

	var _srcChainIdRule []interface{}
	for _, _srcChainIdItem := range _srcChainId {
		_srcChainIdRule = append(_srcChainIdRule, _srcChainIdItem)
	}
	var _srcAddressRule []interface{}
	for _, _srcAddressItem := range _srcAddress {
		_srcAddressRule = append(_srcAddressRule, _srcAddressItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "ReceiveBatchFromChain", _srcChainIdRule, _srcAddressRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155ReceiveBatchFromChain)
				if err := _Erc1155.contract.UnpackLog(event, "ReceiveBatchFromChain", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceiveBatchFromChain is a log parse operation binding the contract event 0x1ae08edbbcd7baa8d064835de8593ce16b313414525ac89534e349f4da7926e4.
//
// Solidity: event ReceiveBatchFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256[] _tokenIds, uint256[] _amounts)
func (_Erc1155 *Erc1155Filterer) ParseReceiveBatchFromChain(log types.Log) (*Erc1155ReceiveBatchFromChain, error) {
	event := new(Erc1155ReceiveBatchFromChain)
	if err := _Erc1155.contract.UnpackLog(event, "ReceiveBatchFromChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155ReceiveFromChainIterator is returned from FilterReceiveFromChain and is used to iterate over the raw logs and unpacked data for ReceiveFromChain events raised by the Erc1155 contract.
type Erc1155ReceiveFromChainIterator struct {
	Event *Erc1155ReceiveFromChain // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155ReceiveFromChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155ReceiveFromChain)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155ReceiveFromChain)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155ReceiveFromChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155ReceiveFromChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155ReceiveFromChain represents a ReceiveFromChain event raised by the Erc1155 contract.
type Erc1155ReceiveFromChain struct {
	SrcChainId uint16
	SrcAddress common.Hash
	ToAddress  common.Address
	TokenId    *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReceiveFromChain is a free log retrieval operation binding the contract event 0x1bf64e58d19fc43de4c44b3d1bb1fae313979af831a7a39f3297564294329f0f.
//
// Solidity: event ReceiveFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256 _tokenId, uint256 _amount)
func (_Erc1155 *Erc1155Filterer) FilterReceiveFromChain(opts *bind.FilterOpts, _srcChainId []uint16, _srcAddress [][]byte, _toAddress []common.Address) (*Erc1155ReceiveFromChainIterator, error) {

	var _srcChainIdRule []interface{}
	for _, _srcChainIdItem := range _srcChainId {
		_srcChainIdRule = append(_srcChainIdRule, _srcChainIdItem)
	}
	var _srcAddressRule []interface{}
	for _, _srcAddressItem := range _srcAddress {
		_srcAddressRule = append(_srcAddressRule, _srcAddressItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "ReceiveFromChain", _srcChainIdRule, _srcAddressRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155ReceiveFromChainIterator{contract: _Erc1155.contract, event: "ReceiveFromChain", logs: logs, sub: sub}, nil
}

// WatchReceiveFromChain is a free log subscription operation binding the contract event 0x1bf64e58d19fc43de4c44b3d1bb1fae313979af831a7a39f3297564294329f0f.
//
// Solidity: event ReceiveFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256 _tokenId, uint256 _amount)
func (_Erc1155 *Erc1155Filterer) WatchReceiveFromChain(opts *bind.WatchOpts, sink chan<- *Erc1155ReceiveFromChain, _srcChainId []uint16, _srcAddress [][]byte, _toAddress []common.Address) (event.Subscription, error) {

	var _srcChainIdRule []interface{}
	for _, _srcChainIdItem := range _srcChainId {
		_srcChainIdRule = append(_srcChainIdRule, _srcChainIdItem)
	}
	var _srcAddressRule []interface{}
	for _, _srcAddressItem := range _srcAddress {
		_srcAddressRule = append(_srcAddressRule, _srcAddressItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "ReceiveFromChain", _srcChainIdRule, _srcAddressRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155ReceiveFromChain)
				if err := _Erc1155.contract.UnpackLog(event, "ReceiveFromChain", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceiveFromChain is a log parse operation binding the contract event 0x1bf64e58d19fc43de4c44b3d1bb1fae313979af831a7a39f3297564294329f0f.
//
// Solidity: event ReceiveFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256 _tokenId, uint256 _amount)
func (_Erc1155 *Erc1155Filterer) ParseReceiveFromChain(log types.Log) (*Erc1155ReceiveFromChain, error) {
	event := new(Erc1155ReceiveFromChain)
	if err := _Erc1155.contract.UnpackLog(event, "ReceiveFromChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155RetryMessageSuccessIterator is returned from FilterRetryMessageSuccess and is used to iterate over the raw logs and unpacked data for RetryMessageSuccess events raised by the Erc1155 contract.
type Erc1155RetryMessageSuccessIterator struct {
	Event *Erc1155RetryMessageSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155RetryMessageSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155RetryMessageSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155RetryMessageSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155RetryMessageSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155RetryMessageSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155RetryMessageSuccess represents a RetryMessageSuccess event raised by the Erc1155 contract.
type Erc1155RetryMessageSuccess struct {
	SrcChainId  uint16
	SrcAddress  []byte
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRetryMessageSuccess is a free log retrieval operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Erc1155 *Erc1155Filterer) FilterRetryMessageSuccess(opts *bind.FilterOpts) (*Erc1155RetryMessageSuccessIterator, error) {

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return &Erc1155RetryMessageSuccessIterator{contract: _Erc1155.contract, event: "RetryMessageSuccess", logs: logs, sub: sub}, nil
}

// WatchRetryMessageSuccess is a free log subscription operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Erc1155 *Erc1155Filterer) WatchRetryMessageSuccess(opts *bind.WatchOpts, sink chan<- *Erc1155RetryMessageSuccess) (event.Subscription, error) {

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155RetryMessageSuccess)
				if err := _Erc1155.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRetryMessageSuccess is a log parse operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Erc1155 *Erc1155Filterer) ParseRetryMessageSuccess(log types.Log) (*Erc1155RetryMessageSuccess, error) {
	event := new(Erc1155RetryMessageSuccess)
	if err := _Erc1155.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155SendBatchToChainIterator is returned from FilterSendBatchToChain and is used to iterate over the raw logs and unpacked data for SendBatchToChain events raised by the Erc1155 contract.
type Erc1155SendBatchToChainIterator struct {
	Event *Erc1155SendBatchToChain // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155SendBatchToChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155SendBatchToChain)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155SendBatchToChain)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155SendBatchToChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155SendBatchToChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155SendBatchToChain represents a SendBatchToChain event raised by the Erc1155 contract.
type Erc1155SendBatchToChain struct {
	DstChainId uint16
	From       common.Address
	ToAddress  common.Hash
	TokenIds   []*big.Int
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSendBatchToChain is a free log retrieval operation binding the contract event 0xddd15f7cfbd674ac2096d598f1650367f8a8bd72b4e3abd85591099ea3b57e33.
//
// Solidity: event SendBatchToChain(uint16 indexed _dstChainId, address indexed _from, bytes indexed _toAddress, uint256[] _tokenIds, uint256[] _amounts)
func (_Erc1155 *Erc1155Filterer) FilterSendBatchToChain(opts *bind.FilterOpts, _dstChainId []uint16, _from []common.Address, _toAddress [][]byte) (*Erc1155SendBatchToChainIterator, error) {

	var _dstChainIdRule []interface{}
	for _, _dstChainIdItem := range _dstChainId {
		_dstChainIdRule = append(_dstChainIdRule, _dstChainIdItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "SendBatchToChain", _dstChainIdRule, _fromRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155SendBatchToChainIterator{contract: _Erc1155.contract, event: "SendBatchToChain", logs: logs, sub: sub}, nil
}

// WatchSendBatchToChain is a free log subscription operation binding the contract event 0xddd15f7cfbd674ac2096d598f1650367f8a8bd72b4e3abd85591099ea3b57e33.
//
// Solidity: event SendBatchToChain(uint16 indexed _dstChainId, address indexed _from, bytes indexed _toAddress, uint256[] _tokenIds, uint256[] _amounts)
func (_Erc1155 *Erc1155Filterer) WatchSendBatchToChain(opts *bind.WatchOpts, sink chan<- *Erc1155SendBatchToChain, _dstChainId []uint16, _from []common.Address, _toAddress [][]byte) (event.Subscription, error) {

	var _dstChainIdRule []interface{}
	for _, _dstChainIdItem := range _dstChainId {
		_dstChainIdRule = append(_dstChainIdRule, _dstChainIdItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "SendBatchToChain", _dstChainIdRule, _fromRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155SendBatchToChain)
				if err := _Erc1155.contract.UnpackLog(event, "SendBatchToChain", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSendBatchToChain is a log parse operation binding the contract event 0xddd15f7cfbd674ac2096d598f1650367f8a8bd72b4e3abd85591099ea3b57e33.
//
// Solidity: event SendBatchToChain(uint16 indexed _dstChainId, address indexed _from, bytes indexed _toAddress, uint256[] _tokenIds, uint256[] _amounts)
func (_Erc1155 *Erc1155Filterer) ParseSendBatchToChain(log types.Log) (*Erc1155SendBatchToChain, error) {
	event := new(Erc1155SendBatchToChain)
	if err := _Erc1155.contract.UnpackLog(event, "SendBatchToChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155SendToChainIterator is returned from FilterSendToChain and is used to iterate over the raw logs and unpacked data for SendToChain events raised by the Erc1155 contract.
type Erc1155SendToChainIterator struct {
	Event *Erc1155SendToChain // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155SendToChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155SendToChain)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155SendToChain)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155SendToChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155SendToChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155SendToChain represents a SendToChain event raised by the Erc1155 contract.
type Erc1155SendToChain struct {
	DstChainId uint16
	From       common.Address
	ToAddress  common.Hash
	TokenId    *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSendToChain is a free log retrieval operation binding the contract event 0x968b0d61ebcf43e5d76ed87bd2c4ee2f22b4969b9f4ca49e3373c025eddd5eeb.
//
// Solidity: event SendToChain(uint16 indexed _dstChainId, address indexed _from, bytes indexed _toAddress, uint256 _tokenId, uint256 _amount)
func (_Erc1155 *Erc1155Filterer) FilterSendToChain(opts *bind.FilterOpts, _dstChainId []uint16, _from []common.Address, _toAddress [][]byte) (*Erc1155SendToChainIterator, error) {

	var _dstChainIdRule []interface{}
	for _, _dstChainIdItem := range _dstChainId {
		_dstChainIdRule = append(_dstChainIdRule, _dstChainIdItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "SendToChain", _dstChainIdRule, _fromRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155SendToChainIterator{contract: _Erc1155.contract, event: "SendToChain", logs: logs, sub: sub}, nil
}

// WatchSendToChain is a free log subscription operation binding the contract event 0x968b0d61ebcf43e5d76ed87bd2c4ee2f22b4969b9f4ca49e3373c025eddd5eeb.
//
// Solidity: event SendToChain(uint16 indexed _dstChainId, address indexed _from, bytes indexed _toAddress, uint256 _tokenId, uint256 _amount)
func (_Erc1155 *Erc1155Filterer) WatchSendToChain(opts *bind.WatchOpts, sink chan<- *Erc1155SendToChain, _dstChainId []uint16, _from []common.Address, _toAddress [][]byte) (event.Subscription, error) {

	var _dstChainIdRule []interface{}
	for _, _dstChainIdItem := range _dstChainId {
		_dstChainIdRule = append(_dstChainIdRule, _dstChainIdItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "SendToChain", _dstChainIdRule, _fromRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155SendToChain)
				if err := _Erc1155.contract.UnpackLog(event, "SendToChain", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSendToChain is a log parse operation binding the contract event 0x968b0d61ebcf43e5d76ed87bd2c4ee2f22b4969b9f4ca49e3373c025eddd5eeb.
//
// Solidity: event SendToChain(uint16 indexed _dstChainId, address indexed _from, bytes indexed _toAddress, uint256 _tokenId, uint256 _amount)
func (_Erc1155 *Erc1155Filterer) ParseSendToChain(log types.Log) (*Erc1155SendToChain, error) {
	event := new(Erc1155SendToChain)
	if err := _Erc1155.contract.UnpackLog(event, "SendToChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155SetMinDstGasIterator is returned from FilterSetMinDstGas and is used to iterate over the raw logs and unpacked data for SetMinDstGas events raised by the Erc1155 contract.
type Erc1155SetMinDstGasIterator struct {
	Event *Erc1155SetMinDstGas // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155SetMinDstGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155SetMinDstGas)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155SetMinDstGas)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155SetMinDstGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155SetMinDstGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155SetMinDstGas represents a SetMinDstGas event raised by the Erc1155 contract.
type Erc1155SetMinDstGas struct {
	DstChainId uint16
	Type       uint16
	MinDstGas  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMinDstGas is a free log retrieval operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Erc1155 *Erc1155Filterer) FilterSetMinDstGas(opts *bind.FilterOpts) (*Erc1155SetMinDstGasIterator, error) {

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return &Erc1155SetMinDstGasIterator{contract: _Erc1155.contract, event: "SetMinDstGas", logs: logs, sub: sub}, nil
}

// WatchSetMinDstGas is a free log subscription operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Erc1155 *Erc1155Filterer) WatchSetMinDstGas(opts *bind.WatchOpts, sink chan<- *Erc1155SetMinDstGas) (event.Subscription, error) {

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155SetMinDstGas)
				if err := _Erc1155.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMinDstGas is a log parse operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Erc1155 *Erc1155Filterer) ParseSetMinDstGas(log types.Log) (*Erc1155SetMinDstGas, error) {
	event := new(Erc1155SetMinDstGas)
	if err := _Erc1155.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155SetPrecrimeIterator is returned from FilterSetPrecrime and is used to iterate over the raw logs and unpacked data for SetPrecrime events raised by the Erc1155 contract.
type Erc1155SetPrecrimeIterator struct {
	Event *Erc1155SetPrecrime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155SetPrecrimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155SetPrecrime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155SetPrecrime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155SetPrecrimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155SetPrecrimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155SetPrecrime represents a SetPrecrime event raised by the Erc1155 contract.
type Erc1155SetPrecrime struct {
	Precrime common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPrecrime is a free log retrieval operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Erc1155 *Erc1155Filterer) FilterSetPrecrime(opts *bind.FilterOpts) (*Erc1155SetPrecrimeIterator, error) {

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return &Erc1155SetPrecrimeIterator{contract: _Erc1155.contract, event: "SetPrecrime", logs: logs, sub: sub}, nil
}

// WatchSetPrecrime is a free log subscription operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Erc1155 *Erc1155Filterer) WatchSetPrecrime(opts *bind.WatchOpts, sink chan<- *Erc1155SetPrecrime) (event.Subscription, error) {

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155SetPrecrime)
				if err := _Erc1155.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPrecrime is a log parse operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Erc1155 *Erc1155Filterer) ParseSetPrecrime(log types.Log) (*Erc1155SetPrecrime, error) {
	event := new(Erc1155SetPrecrime)
	if err := _Erc1155.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155SetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the Erc1155 contract.
type Erc1155SetTrustedRemoteIterator struct {
	Event *Erc1155SetTrustedRemote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155SetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155SetTrustedRemote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155SetTrustedRemote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155SetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155SetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155SetTrustedRemote represents a SetTrustedRemote event raised by the Erc1155 contract.
type Erc1155SetTrustedRemote struct {
	RemoteChainId uint16
	Path          []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Erc1155 *Erc1155Filterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*Erc1155SetTrustedRemoteIterator, error) {

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &Erc1155SetTrustedRemoteIterator{contract: _Erc1155.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Erc1155 *Erc1155Filterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *Erc1155SetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155SetTrustedRemote)
				if err := _Erc1155.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedRemote is a log parse operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Erc1155 *Erc1155Filterer) ParseSetTrustedRemote(log types.Log) (*Erc1155SetTrustedRemote, error) {
	event := new(Erc1155SetTrustedRemote)
	if err := _Erc1155.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155SetTrustedRemoteAddressIterator is returned from FilterSetTrustedRemoteAddress and is used to iterate over the raw logs and unpacked data for SetTrustedRemoteAddress events raised by the Erc1155 contract.
type Erc1155SetTrustedRemoteAddressIterator struct {
	Event *Erc1155SetTrustedRemoteAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155SetTrustedRemoteAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155SetTrustedRemoteAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155SetTrustedRemoteAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155SetTrustedRemoteAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155SetTrustedRemoteAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155SetTrustedRemoteAddress represents a SetTrustedRemoteAddress event raised by the Erc1155 contract.
type Erc1155SetTrustedRemoteAddress struct {
	RemoteChainId uint16
	RemoteAddress []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemoteAddress is a free log retrieval operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Erc1155 *Erc1155Filterer) FilterSetTrustedRemoteAddress(opts *bind.FilterOpts) (*Erc1155SetTrustedRemoteAddressIterator, error) {

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return &Erc1155SetTrustedRemoteAddressIterator{contract: _Erc1155.contract, event: "SetTrustedRemoteAddress", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemoteAddress is a free log subscription operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Erc1155 *Erc1155Filterer) WatchSetTrustedRemoteAddress(opts *bind.WatchOpts, sink chan<- *Erc1155SetTrustedRemoteAddress) (event.Subscription, error) {

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155SetTrustedRemoteAddress)
				if err := _Erc1155.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedRemoteAddress is a log parse operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Erc1155 *Erc1155Filterer) ParseSetTrustedRemoteAddress(log types.Log) (*Erc1155SetTrustedRemoteAddress, error) {
	event := new(Erc1155SetTrustedRemoteAddress)
	if err := _Erc1155.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155SetUseCustomAdapterParamsIterator is returned from FilterSetUseCustomAdapterParams and is used to iterate over the raw logs and unpacked data for SetUseCustomAdapterParams events raised by the Erc1155 contract.
type Erc1155SetUseCustomAdapterParamsIterator struct {
	Event *Erc1155SetUseCustomAdapterParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155SetUseCustomAdapterParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155SetUseCustomAdapterParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155SetUseCustomAdapterParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155SetUseCustomAdapterParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155SetUseCustomAdapterParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155SetUseCustomAdapterParams represents a SetUseCustomAdapterParams event raised by the Erc1155 contract.
type Erc1155SetUseCustomAdapterParams struct {
	UseCustomAdapterParams bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetUseCustomAdapterParams is a free log retrieval operation binding the contract event 0x1584ad594a70cbe1e6515592e1272a987d922b097ead875069cebe8b40c004a4.
//
// Solidity: event SetUseCustomAdapterParams(bool _useCustomAdapterParams)
func (_Erc1155 *Erc1155Filterer) FilterSetUseCustomAdapterParams(opts *bind.FilterOpts) (*Erc1155SetUseCustomAdapterParamsIterator, error) {

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "SetUseCustomAdapterParams")
	if err != nil {
		return nil, err
	}
	return &Erc1155SetUseCustomAdapterParamsIterator{contract: _Erc1155.contract, event: "SetUseCustomAdapterParams", logs: logs, sub: sub}, nil
}

// WatchSetUseCustomAdapterParams is a free log subscription operation binding the contract event 0x1584ad594a70cbe1e6515592e1272a987d922b097ead875069cebe8b40c004a4.
//
// Solidity: event SetUseCustomAdapterParams(bool _useCustomAdapterParams)
func (_Erc1155 *Erc1155Filterer) WatchSetUseCustomAdapterParams(opts *bind.WatchOpts, sink chan<- *Erc1155SetUseCustomAdapterParams) (event.Subscription, error) {

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "SetUseCustomAdapterParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155SetUseCustomAdapterParams)
				if err := _Erc1155.contract.UnpackLog(event, "SetUseCustomAdapterParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetUseCustomAdapterParams is a log parse operation binding the contract event 0x1584ad594a70cbe1e6515592e1272a987d922b097ead875069cebe8b40c004a4.
//
// Solidity: event SetUseCustomAdapterParams(bool _useCustomAdapterParams)
func (_Erc1155 *Erc1155Filterer) ParseSetUseCustomAdapterParams(log types.Log) (*Erc1155SetUseCustomAdapterParams, error) {
	event := new(Erc1155SetUseCustomAdapterParams)
	if err := _Erc1155.contract.UnpackLog(event, "SetUseCustomAdapterParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Erc1155 contract.
type Erc1155TransferBatchIterator struct {
	Event *Erc1155TransferBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155TransferBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155TransferBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155TransferBatch represents a TransferBatch event raised by the Erc1155 contract.
type Erc1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Erc1155 *Erc1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Erc1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155TransferBatchIterator{contract: _Erc1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Erc1155 *Erc1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *Erc1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155TransferBatch)
				if err := _Erc1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Erc1155 *Erc1155Filterer) ParseTransferBatch(log types.Log) (*Erc1155TransferBatch, error) {
	event := new(Erc1155TransferBatch)
	if err := _Erc1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Erc1155 contract.
type Erc1155TransferSingleIterator struct {
	Event *Erc1155TransferSingle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155TransferSingle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155TransferSingle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155TransferSingle represents a TransferSingle event raised by the Erc1155 contract.
type Erc1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Erc1155 *Erc1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Erc1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155TransferSingleIterator{contract: _Erc1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Erc1155 *Erc1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *Erc1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155TransferSingle)
				if err := _Erc1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Erc1155 *Erc1155Filterer) ParseTransferSingle(log types.Log) (*Erc1155TransferSingle, error) {
	event := new(Erc1155TransferSingle)
	if err := _Erc1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Erc1155 contract.
type Erc1155URIIterator struct {
	Event *Erc1155URI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155URI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155URI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155URI represents a URI event raised by the Erc1155 contract.
type Erc1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Erc1155 *Erc1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*Erc1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155URIIterator{contract: _Erc1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Erc1155 *Erc1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *Erc1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155URI)
				if err := _Erc1155.contract.UnpackLog(event, "URI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Erc1155 *Erc1155Filterer) ParseURI(log types.Log) (*Erc1155URI, error) {
	event := new(Erc1155URI)
	if err := _Erc1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155UseSignatureIterator is returned from FilterUseSignature and is used to iterate over the raw logs and unpacked data for UseSignature events raised by the Erc1155 contract.
type Erc1155UseSignatureIterator struct {
	Event *Erc1155UseSignature // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155UseSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155UseSignature)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155UseSignature)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155UseSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155UseSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155UseSignature represents a UseSignature event raised by the Erc1155 contract.
type Erc1155UseSignature struct {
	From      common.Address
	Id        *big.Int
	None      *big.Int
	Value     *big.Int
	Signature []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUseSignature is a free log retrieval operation binding the contract event 0xa53baea067b124fa4f5ff90acfbed6df1a24499f6a0de581d40e8ad128b3d167.
//
// Solidity: event UseSignature(address indexed from, uint256 id, uint256 none, uint256 value, bytes signature)
func (_Erc1155 *Erc1155Filterer) FilterUseSignature(opts *bind.FilterOpts, from []common.Address) (*Erc1155UseSignatureIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "UseSignature", fromRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155UseSignatureIterator{contract: _Erc1155.contract, event: "UseSignature", logs: logs, sub: sub}, nil
}

// WatchUseSignature is a free log subscription operation binding the contract event 0xa53baea067b124fa4f5ff90acfbed6df1a24499f6a0de581d40e8ad128b3d167.
//
// Solidity: event UseSignature(address indexed from, uint256 id, uint256 none, uint256 value, bytes signature)
func (_Erc1155 *Erc1155Filterer) WatchUseSignature(opts *bind.WatchOpts, sink chan<- *Erc1155UseSignature, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "UseSignature", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155UseSignature)
				if err := _Erc1155.contract.UnpackLog(event, "UseSignature", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUseSignature is a log parse operation binding the contract event 0xa53baea067b124fa4f5ff90acfbed6df1a24499f6a0de581d40e8ad128b3d167.
//
// Solidity: event UseSignature(address indexed from, uint256 id, uint256 none, uint256 value, bytes signature)
func (_Erc1155 *Erc1155Filterer) ParseUseSignature(log types.Log) (*Erc1155UseSignature, error) {
	event := new(Erc1155UseSignature)
	if err := _Erc1155.contract.UnpackLog(event, "UseSignature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
