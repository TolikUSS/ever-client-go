package domain

import (
	"encoding/json"
	"math/big"
)

const (
	// SignNone ...
	SignNone SignerType = "None"

	// SignExternal ...
	SignExternal SignerType = "External"

	// SignKeys ...
	SignKeys SignerType = "Keys"

	// SignSigningBox ...
	SignSigningBox SignerType = "SigningBox"

	// TypeS ...
	TypeS AbiType = "Serialized"

	// TypeH ...
	TypeH AbiType = "Handle"

	// MessageBodyTypeInput - Message contains the input of the ABI function.
	MessageBodyTypeInput MessageBodyType = "Input"

	// MessageBodyTypeOutput - Message contains the output of the ABI function.
	MessageBodyTypeOutput MessageBodyType = "Output"

	// MessageBodyTypeInternalOutput - Message contains the input of the imported ABI function.
	//Occurs when contract sends an internal message to other contract.
	MessageBodyTypeInternalOutput MessageBodyType = "InternalOutput"

	// MessageBodyTypeEvent - Message contains the input of the ABI event.
	MessageBodyTypeEvent MessageBodyType = "Event"

	// StateInitSourceTypeMessage ...
	StateInitSourceTypeMessage StateInitSourceType = "Message"

	// StateInitSourceTypeStateInit ...
	StateInitSourceTypeStateInit StateInitSourceType = "StateInit"

	// StateInitSourceTypeTvc ...
	StateInitSourceTypeTvc StateInitSourceType = "Tvc"

	// MessageSourceTypeEncoded ...
	MessageSourceTypeEncoded MessageSourceType = "Encoded"

	// MessageSourceTypeEncodingParams ...
	MessageSourceTypeEncodingParams MessageSourceType = "EncodingParams"
)

// AbiErrorCode ...
var AbiErrorCode map[string]int

type (
	// AbiHandle ...
	AbiHandle int

	// SignerType ...
	SignerType string

	// AbiType ...
	AbiType string

	// MessageBodyType ...
	MessageBodyType string

	// StateInitSourceType ...
	StateInitSourceType string

	// MessageSourceType ...
	MessageSourceType string

	// FunctionHeader ...
	FunctionHeader struct {
		Expire int      `json:"expire,omitempty"`
		Time   *big.Int `json:"time,omitempty"`
		PubKey string   `json:"pubkey,omitempty"`
	}

	// CallSet ...
	CallSet struct {
		FunctionName string          `json:"function_name"`
		Header       *FunctionHeader `json:"header,omitempty"`
		Input        interface{}     `json:"input,omitempty"`
	}

	// DeploySet ...
	DeploySet struct {
		Tvc           string      `json:"tvc"`
		WorkchainID   int         `json:"workchain_id"`
		InitialData   interface{} `json:"initial_data"`
		InitialPubKey string      `json:"initial_pubkey,omitempty"`
	}

	// SignerNone No keys are provided. Creates an unsigned message.
	SignerNone struct {
		Type SignerType `json:"type"`
	}

	// SignerExternal Only public key is provided to generate unsigned message and data_to_sign which can be signed later.
	SignerExternal struct {
		Type      SignerType `json:"type"`
		PublicKey string     `json:"public_key"`
	}

	// SignerKeys Key pair is provided for signing
	SignerKeys struct {
		Type SignerType `json:"type"`
		Keys *KeyPair   `json:"keys"`
	}

	// SignerSigningBox Signing Box interface is provided for signing, allows Dapps to sign messages using external APIs, such as HSM, cold wallet, etc.
	SignerSigningBox struct {
		Type   SignerType       `json:"type"`
		Handle SigningBoxHandle `json:"handle"`
	}

	//Abi ...
	Abi struct {
		Type  AbiType     `json:"type"`
		Value interface{} `json:"value"`
	}

	// AbiParam ...
	AbiParam struct {
		Name       string      `json:"name"`
		Type       string      `json:"type"`
		Components []*AbiParam `json:"components,omitempty"`
	}

	// AbiEvent ...
	AbiEvent struct {
		Name   string      `json:"name"`
		Inputs []*AbiParam `json:"inputs"`
		ID     string      `json:"id,omitempty"`
	}

	// AbiData ...
	AbiData struct {
		Key        *big.Int    `json:"key"`
		Name       string      `json:"name"`
		Type       string      `json:"type"`
		Components []*AbiParam `json:"components,omitempty"`
	}

	// AbiFunctions ...
	AbiFunctions struct {
		Name    string      `json:"name"`
		Inputs  []*AbiParam `json:"inputs"`
		Outputs []*AbiParam `json:"outputs"`
		ID      string      `json:"id,omitempty"`
	}

	// AbiContract ...
	AbiContract struct {
		AbiVersion int             `json:"ABI version,omitempty"`
		Header     []string        `json:"header,omitempty"`
		Functions  []*AbiFunctions `json:"functions,omitempty"`
		Events     []*AbiEvent     `json:"events,omitempty"`
		Data       []*AbiData      `json:"data,omitempty"`
	}

	// StateInitSourceM Deploy message.
	StateInitSourceM struct {
		Type   StateInitSourceType `json:"type"`
		Source interface{}         `json:"source"` //MessageSource
	}

	// StateInitSourceSI State init data
	StateInitSourceSI struct {
		Type    StateInitSourceType `json:"type"`
		Code    string              `json:"code"`
		Data    string              `json:"data"`
		Library string              `json:"library,omitempty"`
	}

	// StateInitSourceT Content of the TVC file. Encoded in base64.
	StateInitSourceT struct {
		Type       StateInitSourceType `json:"type"`
		Tvc        string              `json:"tvc"`
		PublicKey  string              `json:"public_key,omitempty"`
		InitParams *StateInitParams    `json:"init_params,omitempty"`
	}

	// MessageSourceEncoded ...
	MessageSourceEncoded struct {
		Type    MessageSourceType `json:"type"`
		Message string            `json:"message,omitempty"` //omitempty?s
		Abi     *Abi              `json:"abi,omitempty"`
	}

	// MessageSourceEncodingParams ...
	MessageSourceEncodingParams struct {
		Type MessageSourceType `json:"type"`
		*ParamsOfEncodeMessage
	}

	// StateInitParams ...
	StateInitParams struct {
		Abi   Abi         `json:"abi"`
		Value interface{} `json:"value"`
	}

	// ParamsOfEncodeMessageBody ...
	ParamsOfEncodeMessageBody struct {
		Abi                *Abi        `json:"abi"`
		CallSet            CallSet     `json:"call_set"`
		IsInternal         bool        `json:"is_internal"`
		Signer             interface{} `json:"signer"`
		ProcessingTryIndex int         `json:"processing_try_index"`
	}

	// ResultOfEncodeMessageBody ...
	ResultOfEncodeMessageBody struct {
		Body       string `json:"body"`
		DataToSign string `json:"data_to_sign,omitempty"`
	}

	// ParamsOfAttachSignatureToMessageBody ...
	ParamsOfAttachSignatureToMessageBody struct {
		Abi       *Abi   `json:"abi"`
		PublicKey string `json:"public_key"`
		Message   string `json:"message"`
		Signature string `json:"signature"`
	}

	// ResultOfAttachSignatureToMessageBody ...
	ResultOfAttachSignatureToMessageBody struct {
		Body string `json:"body"`
	}

	// ParamsOfEncodeMessage ...
	ParamsOfEncodeMessage struct {
		Type               MessageSourceType `json:"type,omitempty"`
		Abi                *Abi              `json:"abi"`
		Address            string            `json:"address,omitempty"`
		DeploySet          *DeploySet        `json:"deploy_set,omitempty"`
		CallSet            *CallSet          `json:"call_set,omitempty"`
		Signer             interface{}       `json:"signer"`
		ProcessingTryIndex int               `json:"processing_try_index"`
	}

	// ResultOfEncodeMessage ...
	ResultOfEncodeMessage struct {
		Message    string `json:"message"`
		DataToSign string `json:"data_to_sign,omitempty"`
		Address    string `json:"address"`
		MessageID  string `json:"message_id"`
	}

	// ParamsOfEncodeInternalMessage ...
	ParamsOfEncodeInternalMessage struct {
		Abi        *Abi       `json:"abi,omitempty"`
		Address    string     `json:"address,omitempty"`
		SrcAddress string     `json:"src_address,omitempty"`
		DeploySet  *DeploySet `json:"deploy_set,omitempty"`
		CallSet    *CallSet   `json:"call_set,omitempty"`
		Value      string     `json:"value"`
		Bounce     bool       `json:"bounce"`     //? Default is true.
		EnableIhr  bool       `json:"enable_ihr"` //? Default is false.
	}

	// ResultOfEncodeInternalMessage ...
	ResultOfEncodeInternalMessage struct {
		Message   string `json:"message"`
		Address   string `json:"address"`
		MessageID string `json:"message_id"`
	}

	// ParamsOfAttachSignature ...
	ParamsOfAttachSignature struct {
		Abi       *Abi   `json:"abi"`
		PublicKey string `json:"public_key"`
		Message   string `json:"message"`
		Signature string `json:"signature"`
	}

	// ResultOfAttachSignature ...
	ResultOfAttachSignature struct {
		Message   string `json:"message"`
		MessageID string `json:"message_id"`
	}

	// ParamsOfDecodeMessage ...
	ParamsOfDecodeMessage struct {
		Abi     *Abi   `json:"abi"`
		Message string `json:"message"`
	}

	// DecodedMessageBody ...
	DecodedMessageBody struct {
		BodyType MessageBodyType `json:"body_type"`
		Name     string          `json:"name"`
		Value    json.RawMessage `json:"value,omitempty"`
		Header   *FunctionHeader `json:"header,omitempty"`
	}

	// ParamsOfDecodeMessageBody ...
	ParamsOfDecodeMessageBody struct {
		Abi        *Abi   `json:"abi"`
		Body       string `json:"body"`
		IsInternal bool   `json:"is_internal"`
	}

	// ParamsOfEncodeAccount ...
	ParamsOfEncodeAccount struct {
		StateInit   interface{}   `json:"state_init"`
		Balance     *big.Int      `json:"balance,omitempty"`
		LastTransLt *big.Int      `json:"last_trans_lt,omitempty"`
		LastPaid    int           `json:"last_paid,omitempty"`
		BocCache    *BocCacheType `json:"boc_cache,omitempty"`
	}

	// ResultOfEncodeAccount ...
	ResultOfEncodeAccount struct {
		Account string `json:"account"`
		ID      string `json:"id"`
	}

	//AbiUseCase ...
	AbiUseCase interface {
		EncodeMessageBody(*ParamsOfEncodeMessageBody) (*ResultOfEncodeMessageBody, error)
		AttachSignatureToMessageBody(*ParamsOfAttachSignatureToMessageBody) (*ResultOfAttachSignatureToMessageBody, error)
		EncodeMessage(*ParamsOfEncodeMessage) (*ResultOfEncodeMessage, error)
		EncodeInternalMessage(*ParamsOfEncodeInternalMessage) (*ResultOfEncodeInternalMessage, error)
		AttachSignature(*ParamsOfAttachSignature) (*ResultOfAttachSignature, error)
		DecodeMessage(*ParamsOfDecodeMessage) (*DecodedMessageBody, error)
		DecodeMessageBody(*ParamsOfDecodeMessageBody) (*DecodedMessageBody, error)
		EncodeAccount(*ParamsOfEncodeAccount) (*ResultOfEncodeAccount, error)
	}
)

func init() {
	AbiErrorCode = map[string]int{
		"RequiredAddressMissingForEncodeMessage":    301,
		"RequiredCallSetMissingForEncodeMessage":    302,
		"InvalidJson":                               303,
		"InvalidMessage":                            304,
		"EncodeDeployMessageFailed":                 305,
		"EncodeRunMessageFailed":                    306,
		"AttachSignatureFailed":                     307,
		"InvalidTvcImage":                           308,
		"RequiredPublicKeyMissingForFunctionHeader": 309,
		"InvalidSigner":                             310,
		"InvalidAbi":                                311,
		"InvalidFunctionId":                         312,
	}
}

// NewAbiContract Abi type Contract
func NewAbiContract(value *AbiContract) *Abi {
	return &Abi{Type: "Contract", Value: value}
}

// NewAbiJSON Abi type Json
func NewAbiJSON(value string) *Abi {
	return &Abi{Type: "Json", Value: value}
}

// NewAbiHandle Abi type Handle
func NewAbiHandle(value AbiHandle) *Abi {
	return &Abi{Type: "Handle", Value: value}
}

// NewAbiSerialized Abi type Serialized
func NewAbiSerialized(value *AbiContract) *Abi {
	return &Abi{Type: "Serialized", Value: value}
}

// NewSignerNone Signer type None
func NewSignerNone() *SignerNone {
	return &SignerNone{Type: SignNone}
}

// NewSignerExternal Signer type External
func NewSignerExternal(publicKey string) *SignerExternal {
	return &SignerExternal{Type: SignExternal, PublicKey: publicKey}
}

// NewSignerKeys Signer type Keys
func NewSignerKeys(keys *KeyPair) *SignerKeys {
	return &SignerKeys{Type: SignKeys, Keys: keys}
}

// NewSignerSigningBox Signer type SigningBox
func NewSignerSigningBox(handle SigningBoxHandle) *SignerSigningBox {
	return &SignerSigningBox{Type: SignSigningBox, Handle: handle}
}

// StateInitSourceMessage ..
func StateInitSourceMessageEncoded(source interface{}) *StateInitSourceM {
	return &StateInitSourceM{Type: StateInitSourceTypeMessage, Source: source}
}

// StateInitSourceStateInit ...
func StateInitSourceStateInit(code, data, library string) *StateInitSourceSI {
	return &StateInitSourceSI{Type: StateInitSourceTypeStateInit, Code: code, Data: data, Library: library}
}

// StateInitTvc ...
func StateInitTvc(tvc, publicKey string, initParams *StateInitParams) *StateInitSourceT {
	return &StateInitSourceT{Type: StateInitSourceTypeTvc, PublicKey: publicKey, InitParams: initParams}
}

// NewMessageSourceEncoded ...
func NewMessageSourceEncoded(message string, abi *Abi) *MessageSourceEncoded {
	return &MessageSourceEncoded{Type: MessageSourceTypeEncoded, Message: message, Abi: abi}
}

// NewMessageSourceEncodingParams ...
func NewMessageSourceEncodingParams(params *ParamsOfEncodeMessage) *MessageSourceEncodingParams {
	return &MessageSourceEncodingParams{Type: MessageSourceTypeEncodingParams, ParamsOfEncodeMessage: params}
}
