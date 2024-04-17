// Code support by protoc-gen-gogo. DO NOT EDIT.
// source: artela/evm/v1/evm.proto

package support

import (
	fmt "fmt"
	cosmossdk_io_math "cosmossdk.io/math"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this support file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the EVM module parameters
type Params struct {
	// evm_denom represents the token denomination used to run the EVM states
	// transitions.
	EvmDenom string `protobuf:"bytes,1,opt,name=evm_denom,json=evmDenom,proto3" json:"evm_denom,omitempty" yaml:"evm_denom"`
	// enable_create toggles states transitions that use the vm.Create function
	EnableCreate bool `protobuf:"varint,2,opt,name=enable_create,json=enableCreate,proto3" json:"enable_create,omitempty" yaml:"enable_create"`
	// enable_call toggles states transitions that use the vm.Call function
	EnableCall bool `protobuf:"varint,3,opt,name=enable_call,json=enableCall,proto3" json:"enable_call,omitempty" yaml:"enable_call"`
	// extra_eips defines the additional EIPs for the vm.Config
	ExtraEIPs []int64 `protobuf:"varint,4,rep,packed,name=extra_eips,json=extraEips,proto3" json:"extra_eips,omitempty" yaml:"extra_eips"`
	// chain_config defines the EVM chain configuration parameters
	ChainConfig ChainConfig `protobuf:"bytes,5,opt,name=chain_config,json=chainConfig,proto3" json:"chain_config" yaml:"chain_config"`
	// allow_unprotected_txs defines if replay-protected (i.e non EIP155
	// signed) transactions can be executed on the states machine.
	AllowUnprotectedTxs bool `protobuf:"varint,6,opt,name=allow_unprotected_txs,json=allowUnprotectedTxs,proto3" json:"allow_unprotected_txs,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95fb7abfbae4d4d, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEvmDenom() string {
	if m != nil {
		return m.EvmDenom
	}
	return ""
}

func (m *Params) GetEnableCreate() bool {
	if m != nil {
		return m.EnableCreate
	}
	return false
}

func (m *Params) GetEnableCall() bool {
	if m != nil {
		return m.EnableCall
	}
	return false
}

func (m *Params) GetExtraEIPs() []int64 {
	if m != nil {
		return m.ExtraEIPs
	}
	return nil
}

func (m *Params) GetChainConfig() ChainConfig {
	if m != nil {
		return m.ChainConfig
	}
	return ChainConfig{}
}

func (m *Params) GetAllowUnprotectedTxs() bool {
	if m != nil {
		return m.AllowUnprotectedTxs
	}
	return false
}

// ChainConfig defines the Ethereum ChainConfig parameters using *sdk.Int values
// instead of *big.Int.
type ChainConfig struct {
	// homestead_block switch (nil no fork, 0 = already homestead)
	HomesteadBlock *cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=homestead_block,json=homesteadBlock,proto3,customtype=cosmossdk_io_math.Int" json:"homestead_block,omitempty" yaml:"homestead_block"`
	// dao_fork_block corresponds to TheDAO hard-fork switch block (nil no fork)
	DAOForkBlock *cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=dao_fork_block,json=daoForkBlock,proto3,customtype=cosmossdk_io_math.Int" json:"dao_fork_block,omitempty" yaml:"dao_fork_block"`
	// dao_fork_support defines whether the nodes supports or opposes the DAO hard-fork
	DAOForkSupport bool `protobuf:"varint,3,opt,name=dao_fork_support,json=daoForkSupport,proto3" json:"dao_fork_support,omitempty" yaml:"dao_fork_support"`
	// eip150_block: EIP150 implements the Gas price changes
	// (https://github.com/ethereum/EIPs/issues/150) EIP150 HF block (nil no fork)
	EIP150Block *cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=eip150_block,json=eip150Block,proto3,customtype=cosmossdk_io_math.Int" json:"eip150_block,omitempty" yaml:"eip150_block"`
	// eip150_hash: EIP150 HF hash (needed for header only clients as only gas pricing changed)
	EIP150Hash string `protobuf:"bytes,5,opt,name=eip150_hash,json=eip150Hash,proto3" json:"eip150_hash,omitempty" yaml:"byzantium_block"`
	// eip155_block: EIP155Block HF block
	EIP155Block *cosmossdk_io_math.Int `protobuf:"bytes,6,opt,name=eip155_block,json=eip155Block,proto3,customtype=cosmossdk_io_math.Int" json:"eip155_block,omitempty" yaml:"eip155_block"`
	// eip158_block: EIP158 HF block
	EIP158Block *cosmossdk_io_math.Int `protobuf:"bytes,7,opt,name=eip158_block,json=eip158Block,proto3,customtype=cosmossdk_io_math.Int" json:"eip158_block,omitempty" yaml:"eip158_block"`
	// byzantium_block: Byzantium switch block (nil no fork, 0 = already on byzantium)
	ByzantiumBlock *cosmossdk_io_math.Int `protobuf:"bytes,8,opt,name=byzantium_block,json=byzantiumBlock,proto3,customtype=cosmossdk_io_math.Int" json:"byzantium_block,omitempty" yaml:"byzantium_block"`
	// constantinople_block: Constantinople switch block (nil no fork, 0 = already activated)
	ConstantinopleBlock *cosmossdk_io_math.Int `protobuf:"bytes,9,opt,name=constantinople_block,json=constantinopleBlock,proto3,customtype=cosmossdk_io_math.Int" json:"constantinople_block,omitempty" yaml:"constantinople_block"`
	// petersburg_block: Petersburg switch block (nil same as Constantinople)
	PetersburgBlock *cosmossdk_io_math.Int `protobuf:"bytes,10,opt,name=petersburg_block,json=petersburgBlock,proto3,customtype=cosmossdk_io_math.Int" json:"petersburg_block,omitempty" yaml:"petersburg_block"`
	// istanbul_block: Istanbul switch block (nil no fork, 0 = already on istanbul)
	IstanbulBlock *cosmossdk_io_math.Int `protobuf:"bytes,11,opt,name=istanbul_block,json=istanbulBlock,proto3,customtype=cosmossdk_io_math.Int" json:"istanbul_block,omitempty" yaml:"istanbul_block"`
	// muir_glacier_block: Eip-2384 (bomb delay) switch block (nil no fork, 0 = already activated)
	MuirGlacierBlock *cosmossdk_io_math.Int `protobuf:"bytes,12,opt,name=muir_glacier_block,json=muirGlacierBlock,proto3,customtype=cosmossdk_io_math.Int" json:"muir_glacier_block,omitempty" yaml:"muir_glacier_block"`
	// berlin_block: Berlin switch block (nil = no fork, 0 = already on berlin)
	BerlinBlock *cosmossdk_io_math.Int `protobuf:"bytes,13,opt,name=berlin_block,json=berlinBlock,proto3,customtype=cosmossdk_io_math.Int" json:"berlin_block,omitempty" yaml:"berlin_block"`
	// london_block: London switch block (nil = no fork, 0 = already on london)
	LondonBlock *cosmossdk_io_math.Int `protobuf:"bytes,17,opt,name=london_block,json=londonBlock,proto3,customtype=cosmossdk_io_math.Int" json:"london_block,omitempty" yaml:"london_block"`
	// arrow_glacier_block: Eip-4345 (bomb delay) switch block (nil = no fork, 0 = already activated)
	ArrowGlacierBlock *cosmossdk_io_math.Int `protobuf:"bytes,18,opt,name=arrow_glacier_block,json=arrowGlacierBlock,proto3,customtype=cosmossdk_io_math.Int" json:"arrow_glacier_block,omitempty" yaml:"arrow_glacier_block"`
	// gray_glacier_block: EIP-5133 (bomb delay) switch block (nil = no fork, 0 = already activated)
	GrayGlacierBlock *cosmossdk_io_math.Int `protobuf:"bytes,20,opt,name=gray_glacier_block,json=grayGlacierBlock,proto3,customtype=cosmossdk_io_math.Int" json:"gray_glacier_block,omitempty" yaml:"gray_glacier_block"`
	// merge_netsplit_block: Virtual fork after The Merge to use as a network splitter
	MergeNetsplitBlock *cosmossdk_io_math.Int `protobuf:"bytes,21,opt,name=merge_netsplit_block,json=mergeNetsplitBlock,proto3,customtype=cosmossdk_io_math.Int" json:"merge_netsplit_block,omitempty" yaml:"merge_netsplit_block"`
	// shanghai_block switch block (nil = no fork, 0 = already on shanghai)
	ShanghaiBlock *cosmossdk_io_math.Int `protobuf:"bytes,22,opt,name=shanghai_block,json=shanghaiBlock,proto3,customtype=cosmossdk_io_math.Int" json:"shanghai_block,omitempty" yaml:"shanghai_block"`
	// cancun_block switch block (nil = no fork, 0 = already on cancun)
	CancunBlock *cosmossdk_io_math.Int `protobuf:"bytes,23,opt,name=cancun_block,json=cancunBlock,proto3,customtype=cosmossdk_io_math.Int" json:"cancun_block,omitempty" yaml:"cancun_block"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95fb7abfbae4d4d, []int{1}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

func (m *ChainConfig) GetDAOForkSupport() bool {
	if m != nil {
		return m.DAOForkSupport
	}
	return false
}

func (m *ChainConfig) GetEIP150Hash() string {
	if m != nil {
		return m.EIP150Hash
	}
	return ""
}

// State represents a single Storage key value pair item.
type State struct {
	// key is the stored key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value is the stored value for the given key
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95fb7abfbae4d4d, []int{2}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_State.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return m.Size()
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *State) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// TransactionLogs define the logs support from a txs execution
// with a given hash. It it used for import/export data as transactions are not
// persisted on blockchain states after an upgrade.
type TransactionLogs struct {
	// hash of the txs
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// logs is an array of Logs for the given txs hash
	Logs []*Log `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (m *TransactionLogs) Reset()         { *m = TransactionLogs{} }
func (m *TransactionLogs) String() string { return proto.CompactTextString(m) }
func (*TransactionLogs) ProtoMessage()    {}
func (*TransactionLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95fb7abfbae4d4d, []int{3}
}
func (m *TransactionLogs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionLogs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionLogs.Merge(m, src)
}
func (m *TransactionLogs) XXX_Size() int {
	return m.Size()
}
func (m *TransactionLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionLogs.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionLogs proto.InternalMessageInfo

func (m *TransactionLogs) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *TransactionLogs) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

// Log represents an protobuf compatible Ethereum Log that defines a contract
// log event. These events are support by the LOG opcode and stored/indexed by
// the node.
//
// NOTE: address, topics and data are consensus fields. The rest of the fields
// are derived, i.e. filled in by the nodes, but not secured by consensus.
type Log struct {
	// address of the contract that support the event
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// topics is a list of topics provided by the contract.
	Topics []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	// data which is supplied by the contract, usually ABI-encoded
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// block_number of the block in which the txs was included
	BlockNumber uint64 `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"blockNumber"`
	// tx_hash is the txs hash
	TxHash string `protobuf:"bytes,5,opt,name=tx_hash,json=txHash,proto3" json:"transactionHash"`
	// tx_index of the txs in the block
	TxIndex uint64 `protobuf:"varint,6,opt,name=tx_index,json=txIndex,proto3" json:"transactionIndex"`
	// block_hash of the block in which the txs was included
	BlockHash string `protobuf:"bytes,7,opt,name=block_hash,json=blockHash,proto3" json:"blockHash"`
	// index of the log in the block
	Index uint64 `protobuf:"varint,8,opt,name=index,proto3" json:"logIndex"`
	// removed is true if this log was reverted due to a chain
	// reorganisation. You must pay attention to this field if you receive logs
	// through a filter query.
	Removed bool `protobuf:"varint,9,opt,name=removed,proto3" json:"removed,omitempty"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95fb7abfbae4d4d, []int{4}
}
func (m *Log) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Log.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return m.Size()
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Log) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *Log) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Log) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *Log) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *Log) GetTxIndex() uint64 {
	if m != nil {
		return m.TxIndex
	}
	return 0
}

func (m *Log) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *Log) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Log) GetRemoved() bool {
	if m != nil {
		return m.Removed
	}
	return false
}

// TxResult stores results of Tx execution.
type TxResult struct {
	// contract_address contains the ethereum address of the created contract (if
	// any). If the states transition is an evm.Call, the contract address will be
	// empty.
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty" yaml:"contract_address"`
	// bloom represents the bloom filter bytes
	Bloom []byte `protobuf:"bytes,2,opt,name=bloom,proto3" json:"bloom,omitempty"`
	// tx_logs contains the txs hash and the proto-compatible ethereum
	// logs.
	TxLogs TransactionLogs `protobuf:"bytes,3,opt,name=tx_logs,json=txLogs,proto3" json:"tx_logs" yaml:"tx_logs"`
	// ret defines the bytes from the execution.
	Ret []byte `protobuf:"bytes,4,opt,name=ret,proto3" json:"ret,omitempty"`
	// reverted flag is set to true when the call has been reverted
	Reverted bool `protobuf:"varint,5,opt,name=reverted,proto3" json:"reverted,omitempty"`
	// gas_used notes the amount of gas consumed while execution
	GasUsed uint64 `protobuf:"varint,6,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
}

func (m *TxResult) Reset()         { *m = TxResult{} }
func (m *TxResult) String() string { return proto.CompactTextString(m) }
func (*TxResult) ProtoMessage()    {}
func (*TxResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95fb7abfbae4d4d, []int{5}
}
func (m *TxResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResult.Merge(m, src)
}
func (m *TxResult) XXX_Size() int {
	return m.Size()
}
func (m *TxResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResult.DiscardUnknown(m)
}

var xxx_messageInfo_TxResult proto.InternalMessageInfo

// AccessTuple is the element type of an access list.
type AccessTuple struct {
	// address is a hex formatted ethereum address
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// storage_keys are hex formatted hashes of the storage keys
	StorageKeys []string `protobuf:"bytes,2,rep,name=storage_keys,json=storageKeys,proto3" json:"storageKeys"`
}

func (m *AccessTuple) Reset()         { *m = AccessTuple{} }
func (m *AccessTuple) String() string { return proto.CompactTextString(m) }
func (*AccessTuple) ProtoMessage()    {}
func (*AccessTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95fb7abfbae4d4d, []int{6}
}
func (m *AccessTuple) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessTuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessTuple.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessTuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTuple.Merge(m, src)
}
func (m *AccessTuple) XXX_Size() int {
	return m.Size()
}
func (m *AccessTuple) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTuple.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTuple proto.InternalMessageInfo

// TraceConfig holds extra parameters to trace functions.
type TraceConfig struct {
	// tracer is a custom javascript tracer
	Tracer string `protobuf:"bytes,1,opt,name=tracer,proto3" json:"tracer,omitempty"`
	// timeout overrides the default timeout of 5 seconds for JavaScript-based tracing
	// calls
	Timeout string `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// reexec defines the number of blocks the tracer is willing to go back
	Reexec uint64 `protobuf:"varint,3,opt,name=reexec,proto3" json:"reexec,omitempty"`
	// disable_stack switches stack capture
	DisableStack bool `protobuf:"varint,5,opt,name=disable_stack,json=disableStack,proto3" json:"disableStack"`
	// disable_storage switches storage capture
	DisableStorage bool `protobuf:"varint,6,opt,name=disable_storage,json=disableStorage,proto3" json:"disableStorage"`
	// debug can be used to print output during capture end
	Debug bool `protobuf:"varint,8,opt,name=debug,proto3" json:"debug,omitempty"`
	// limit defines the maximum length of output, but zero means unlimited
	Limit int32 `protobuf:"varint,9,opt,name=limit,proto3" json:"limit,omitempty"`
	// overrides can be used to execute a trace using future fork rules
	Overrides *ChainConfig `protobuf:"bytes,10,opt,name=overrides,proto3" json:"overrides,omitempty"`
	// enable_memory switches memory capture
	EnableMemory bool `protobuf:"varint,11,opt,name=enable_memory,json=enableMemory,proto3" json:"enableMemory"`
	// enable_return_data switches the capture of return data
	EnableReturnData bool `protobuf:"varint,12,opt,name=enable_return_data,json=enableReturnData,proto3" json:"enableReturnData"`
	// tracer_json_config configures the tracer using a JSON string
	TracerJsonConfig string `protobuf:"bytes,13,opt,name=tracer_json_config,json=tracerJsonConfig,proto3" json:"tracerConfig"`
}

func (m *TraceConfig) Reset()         { *m = TraceConfig{} }
func (m *TraceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceConfig) ProtoMessage()    {}
func (*TraceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95fb7abfbae4d4d, []int{7}
}
func (m *TraceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceConfig.Merge(m, src)
}
func (m *TraceConfig) XXX_Size() int {
	return m.Size()
}
func (m *TraceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceConfig proto.InternalMessageInfo

func (m *TraceConfig) GetTracer() string {
	if m != nil {
		return m.Tracer
	}
	return ""
}

func (m *TraceConfig) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

func (m *TraceConfig) GetReexec() uint64 {
	if m != nil {
		return m.Reexec
	}
	return 0
}

func (m *TraceConfig) GetDisableStack() bool {
	if m != nil {
		return m.DisableStack
	}
	return false
}

func (m *TraceConfig) GetDisableStorage() bool {
	if m != nil {
		return m.DisableStorage
	}
	return false
}

func (m *TraceConfig) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

func (m *TraceConfig) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TraceConfig) GetOverrides() *ChainConfig {
	if m != nil {
		return m.Overrides
	}
	return nil
}

func (m *TraceConfig) GetEnableMemory() bool {
	if m != nil {
		return m.EnableMemory
	}
	return false
}

func (m *TraceConfig) GetEnableReturnData() bool {
	if m != nil {
		return m.EnableReturnData
	}
	return false
}

func (m *TraceConfig) GetTracerJsonConfig() string {
	if m != nil {
		return m.TracerJsonConfig
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "artela.evm.v1.Params")
	proto.RegisterType((*ChainConfig)(nil), "artela.evm.v1.ChainConfig")
	proto.RegisterType((*State)(nil), "artela.evm.v1.State")
	proto.RegisterType((*TransactionLogs)(nil), "artela.evm.v1.TransactionLogs")
	proto.RegisterType((*Log)(nil), "artela.evm.v1.Log")
	proto.RegisterType((*TxResult)(nil), "artela.evm.v1.TxResult")
	proto.RegisterType((*AccessTuple)(nil), "artela.evm.v1.AccessTuple")
	proto.RegisterType((*TraceConfig)(nil), "artela.evm.v1.TraceConfig")
}

func init() { proto.RegisterFile("artela/evm/v1/evm.proto", fileDescriptor_c95fb7abfbae4d4d) }

var fileDescriptor_c95fb7abfbae4d4d = []byte{
	// 1605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0x4d, 0x6f, 0x23, 0xb7,
	0x19, 0xb6, 0xad, 0xb1, 0x3d, 0xa2, 0xbe, 0xc6, 0xb4, 0xd6, 0xab, 0xec, 0x02, 0x1e, 0x63, 0x0e,
	0x81, 0x0f, 0x59, 0x2b, 0x76, 0x60, 0x74, 0x91, 0xa2, 0x05, 0xac, 0x5d, 0x27, 0xb1, 0xbb, 0x49,
	0x16, 0x5c, 0x07, 0x05, 0x72, 0x19, 0x50, 0x33, 0xcc, 0x68, 0xa2, 0x99, 0xa1, 0x40, 0x72, 0xb4,
	0x52, 0xdb, 0x73, 0x91, 0x63, 0xff, 0x40, 0x8b, 0xfe, 0x9c, 0xa0, 0xa7, 0x1c, 0x8b, 0x1e, 0x06,
	0x85, 0xf7, 0xe6, 0xa3, 0x7e, 0x41, 0xc1, 0x0f, 0x7d, 0xda, 0x08, 0x6a, 0x9d, 0x34, 0xcf, 0xfb,
	0xf1, 0x3c, 0xe4, 0xcb, 0x97, 0x22, 0x09, 0x9e, 0x62, 0x26, 0x48, 0x82, 0xdb, 0x64, 0x98, 0xb6,
	0x87, 0xa7, 0xf2, 0xe7, 0x64, 0xc0, 0xa8, 0xa0, 0xb0, 0xa6, 0x1d, 0x27, 0xd2, 0x32, 0x3c, 0x7d,
	0xd6, 0x8c, 0x68, 0x44, 0x95, 0xa7, 0x2d, 0xbf, 0x74, 0x90, 0xf7, 0xd7, 0x12, 0xd8, 0x79, 0x8b,
	0x19, 0x4e, 0x39, 0x3c, 0x05, 0x65, 0x32, 0x4c, 0xfd, 0x90, 0x64, 0x34, 0x6d, 0x6d, 0x1e, 0x6d,
	0x1e, 0x97, 0x3b, 0xcd, 0x49, 0xe1, 0x3a, 0x63, 0x9c, 0x26, 0x9f, 0x7b, 0x33, 0x97, 0x87, 0x6c,
	0x32, 0x4c, 0x5f, 0xcb, 0x4f, 0xf8, 0x3b, 0x50, 0x23, 0x19, 0xee, 0x26, 0xc4, 0x0f, 0x18, 0xc1,
	0x82, 0xb4, 0xb6, 0x8e, 0x36, 0x8f, 0xed, 0x4e, 0x6b, 0x52, 0xb8, 0x4d, 0x93, 0xb6, 0xe8, 0xf6,
	0x50, 0x55, 0xe3, 0x57, 0x0a, 0xc2, 0xdf, 0x80, 0xca, 0xd4, 0x8f, 0x93, 0xa4, 0x55, 0x52, 0xc9,
	0x07, 0x93, 0xc2, 0x85, 0xcb, 0xc9, 0x38, 0x49, 0x3c, 0x04, 0x4c, 0x2a, 0x4e, 0x12, 0x78, 0x01,
	0x00, 0x19, 0x09, 0x86, 0x7d, 0x12, 0x0f, 0x78, 0xcb, 0x3a, 0x2a, 0x1d, 0x97, 0x3a, 0xde, 0x6d,
	0xe1, 0x96, 0x2f, 0xa5, 0xf5, 0xf2, 0xea, 0x2d, 0x9f, 0x14, 0xee, 0x9e, 0x21, 0x99, 0x05, 0x7a,
	0xa8, 0xac, 0xc0, 0x65, 0x3c, 0xe0, 0xf0, 0x7b, 0x50, 0x0d, 0x7a, 0x38, 0xce, 0xfc, 0x80, 0x66,
	0x3f, 0xc4, 0x51, 0x6b, 0xfb, 0x68, 0xf3, 0xb8, 0x72, 0xf6, 0xec, 0x64, 0xa9, 0x68, 0x27, 0xaf,
	0x64, 0xc8, 0x2b, 0x15, 0xd1, 0x79, 0xfe, 0x73, 0xe1, 0x6e, 0x4c, 0x0a, 0x77, 0x5f, 0xf3, 0x2e,
	0x66, 0x7b, 0xa8, 0x12, 0xcc, 0x23, 0xe1, 0x19, 0x78, 0x82, 0x93, 0x84, 0xbe, 0xf7, 0xf3, 0x4c,
	0x56, 0x99, 0x04, 0x82, 0x84, 0xbe, 0x18, 0xf1, 0xd6, 0x8e, 0x9c, 0x21, 0xda, 0x57, 0xce, 0xef,
	0xe6, 0xbe, 0x9b, 0x11, 0xf7, 0xfe, 0xb1, 0x07, 0x2a, 0x0b, 0x6a, 0x30, 0x05, 0x8d, 0x1e, 0x4d,
	0x09, 0x17, 0x04, 0x87, 0x7e, 0x37, 0xa1, 0x41, 0xdf, 0xac, 0xc9, 0xeb, 0xff, 0x14, 0xee, 0xc7,
	0x51, 0x2c, 0x7a, 0x79, 0xf7, 0x24, 0xa0, 0x69, 0x3b, 0xa0, 0x3c, 0xa5, 0xdc, 0xfc, 0xbc, 0xe0,
	0x61, 0xbf, 0x2d, 0xc6, 0x03, 0xc2, 0x4f, 0xae, 0x32, 0x31, 0x29, 0xdc, 0x03, 0x3d, 0xd8, 0x15,
	0x2a, 0x0f, 0xd5, 0x67, 0x96, 0x8e, 0x34, 0xc0, 0x31, 0xa8, 0x87, 0x98, 0xfa, 0x3f, 0x50, 0xd6,
	0x37, 0x6a, 0x5b, 0x4a, 0xed, 0xdd, 0xff, 0xaf, 0x76, 0x5b, 0xb8, 0xd5, 0xd7, 0x17, 0xdf, 0x7e,
	0x41, 0x59, 0x5f, 0x71, 0x4e, 0x0a, 0xf7, 0x89, 0x56, 0x5f, 0x66, 0xf6, 0x50, 0x35, 0xc4, 0x74,
	0x16, 0x06, 0xff, 0x08, 0x9c, 0x59, 0x00, 0xcf, 0x07, 0x03, 0xca, 0x84, 0x69, 0x85, 0x17, 0xb7,
	0x85, 0x5b, 0x37, 0x94, 0xef, 0xb4, 0x67, 0x52, 0xb8, 0x4f, 0x57, 0x48, 0x4d, 0x8e, 0x87, 0xea,
	0x86, 0xd6, 0x84, 0x42, 0x0e, 0xaa, 0x24, 0x1e, 0x9c, 0x9e, 0x7f, 0x6a, 0x66, 0x64, 0xa9, 0x19,
	0xbd, 0x7d, 0xd4, 0x8c, 0x2a, 0x97, 0x57, 0x6f, 0x4f, 0xcf, 0x3f, 0x9d, 0x4e, 0xc8, 0xac, 0xfd,
	0x22, 0xad, 0x87, 0x2a, 0x1a, 0xea, 0xd9, 0x5c, 0x01, 0x03, 0xfd, 0x1e, 0xe6, 0x3d, 0xd5, 0x56,
	0xe5, 0xce, 0xf1, 0x6d, 0xe1, 0x02, 0xcd, 0xf4, 0x15, 0xe6, 0xbd, 0xf9, 0xba, 0x74, 0xc7, 0x7f,
	0xc2, 0x99, 0x88, 0xf3, 0x74, 0xca, 0x05, 0x74, 0xb2, 0x8c, 0x9a, 0x8d, 0xff, 0xdc, 0x8c, 0x7f,
	0x67, 0xed, 0xf1, 0x9f, 0x3f, 0x34, 0xfe, 0xf3, 0xe5, 0xf1, 0xeb, 0x98, 0x99, 0xe8, 0x4b, 0x23,
	0xba, 0xbb, 0xb6, 0xe8, 0xcb, 0x87, 0x44, 0x5f, 0x2e, 0x8b, 0xea, 0x18, 0xd9, 0xec, 0x2b, 0x95,
	0x68, 0xd9, 0xeb, 0x37, 0xfb, 0xbd, 0xa2, 0xd6, 0x67, 0x16, 0x2d, 0xf7, 0x17, 0xd0, 0x0c, 0x68,
	0xc6, 0x85, 0xb4, 0x65, 0x74, 0x90, 0x10, 0xa3, 0x59, 0x56, 0x9a, 0x57, 0x8f, 0xd2, 0x7c, 0x6e,
	0xfe, 0x0d, 0x1e, 0xe0, 0xf3, 0xd0, 0xfe, 0xb2, 0x59, 0xab, 0x0f, 0x80, 0x33, 0x20, 0x82, 0x30,
	0xde, 0xcd, 0x59, 0x64, 0x94, 0x81, 0x52, 0xbe, 0x7c, 0x94, 0xb2, 0xd9, 0x07, 0xab, 0x5c, 0x1e,
	0x6a, 0xcc, 0x4d, 0x5a, 0xf1, 0x47, 0x50, 0x8f, 0xe5, 0x30, 0xba, 0x79, 0x62, 0xf4, 0x2a, 0x4a,
	0xef, 0xd5, 0xa3, 0xf4, 0xcc, 0x66, 0x5e, 0x66, 0xf2, 0x50, 0x6d, 0x6a, 0xd0, 0x5a, 0x39, 0x80,
	0x69, 0x1e, 0x33, 0x3f, 0x4a, 0x70, 0x10, 0x13, 0x66, 0xf4, 0xaa, 0x4a, 0xef, 0xcb, 0x47, 0xe9,
	0x7d, 0xa4, 0xf5, 0xee, 0xb3, 0x79, 0xc8, 0x91, 0xc6, 0x2f, 0xb5, 0x4d, 0xcb, 0x86, 0xa0, 0xda,
	0x25, 0x2c, 0x89, 0x33, 0x23, 0x58, 0x53, 0x82, 0x17, 0x8f, 0x12, 0x34, 0x7d, 0xba, 0xc8, 0xe3,
	0xa1, 0x8a, 0x86, 0x33, 0x95, 0x84, 0x66, 0x21, 0x9d, 0xaa, 0xec, 0xad, 0xaf, 0xb2, 0xc8, 0xe3,
	0xa1, 0x8a, 0x86, 0x5a, 0x65, 0x04, 0xf6, 0x31, 0x63, 0xf4, 0xfd, 0x4a, 0x0d, 0xa1, 0x12, 0xfb,
	0xea, 0x51, 0x62, 0xcf, 0xb4, 0xd8, 0x03, 0x74, 0x1e, 0xda, 0x53, 0xd6, 0xa5, 0x2a, 0xe6, 0x00,
	0x46, 0x0c, 0x8f, 0x57, 0x84, 0x9b, 0xeb, 0x2f, 0xde, 0x7d, 0x36, 0x0f, 0x39, 0xd2, 0xb8, 0x24,
	0xfb, 0x67, 0xd0, 0x4c, 0x09, 0x8b, 0x88, 0x9f, 0x11, 0xc1, 0x07, 0x49, 0x2c, 0x8c, 0xf0, 0x93,
	0xf5, 0xf7, 0xe3, 0x43, 0x7c, 0x1e, 0x82, 0xca, 0xfc, 0x8d, 0xb1, 0xce, 0x36, 0x07, 0xef, 0xe1,
	0x2c, 0xea, 0xe1, 0xd8, 0xc8, 0x1e, 0xac, 0xbf, 0x39, 0x96, 0x99, 0x3c, 0x54, 0x9b, 0x1a, 0x66,
	0xfd, 0x13, 0xe0, 0x2c, 0xc8, 0xa7, 0xfd, 0xf3, 0x74, 0xfd, 0xfe, 0x59, 0xe4, 0x91, 0xd7, 0x0f,
	0x05, 0x95, 0xca, 0xb5, 0x65, 0xd7, 0x9d, 0xc6, 0xb5, 0x65, 0x37, 0x1c, 0xe7, 0xda, 0xb2, 0x1d,
	0x67, 0xef, 0xda, 0xb2, 0xf7, 0x9d, 0x26, 0xaa, 0x8d, 0x69, 0x42, 0xfd, 0xe1, 0x67, 0x3a, 0x09,
	0x55, 0xc8, 0x7b, 0xcc, 0xcd, 0x7f, 0x24, 0xaa, 0x07, 0x58, 0xe0, 0x64, 0xcc, 0x4d, 0xa9, 0x90,
	0xa3, 0x0b, 0xb8, 0x70, 0x6a, 0xb7, 0xc1, 0xf6, 0x3b, 0x21, 0x6f, 0x6d, 0x0e, 0x28, 0xf5, 0xc9,
	0x58, 0xdf, 0x46, 0x90, 0xfc, 0x84, 0x4d, 0xb0, 0x3d, 0xc4, 0x49, 0xae, 0xaf, 0x7f, 0x65, 0xa4,
	0x81, 0xf7, 0x35, 0x68, 0xdc, 0x30, 0x9c, 0x71, 0x1c, 0x88, 0x98, 0x66, 0x6f, 0x68, 0xc4, 0x21,
	0x04, 0x96, 0x3a, 0x15, 0x75, 0xae, 0xfa, 0x86, 0x1f, 0x03, 0x2b, 0xa1, 0x11, 0x6f, 0x6d, 0x1d,
	0x95, 0x8e, 0x2b, 0x67, 0x70, 0xe5, 0x02, 0xf6, 0x86, 0x46, 0x48, 0xf9, 0xbd, 0x7f, 0x6d, 0x81,
	0xd2, 0x1b, 0x1a, 0xc1, 0x16, 0xd8, 0xc5, 0x61, 0xc8, 0x08, 0xe7, 0x86, 0x66, 0x0a, 0xe1, 0x01,
	0xd8, 0x11, 0x74, 0x10, 0x07, 0x9a, 0xab, 0x8c, 0x0c, 0x92, 0xaa, 0x21, 0x16, 0x58, 0x5d, 0x2a,
	0xaa, 0x48, 0x7d, 0xc3, 0x33, 0x50, 0x55, 0xd3, 0xf2, 0xb3, 0x3c, 0xed, 0x12, 0xa6, 0xee, 0x06,
	0x56, 0xa7, 0x71, 0x57, 0xb8, 0x15, 0x65, 0xff, 0x46, 0x99, 0xd1, 0x22, 0x80, 0x9f, 0x80, 0x5d,
	0x31, 0x5a, 0x3c, 0xd6, 0xf7, 0xef, 0x0a, 0xb7, 0x21, 0xe6, 0x73, 0x94, 0xa7, 0x36, 0xda, 0x11,
	0x23, 0x75, 0x7a, 0xb7, 0x81, 0x2d, 0x46, 0x7e, 0x9c, 0x85, 0x64, 0xa4, 0x4e, 0x6e, 0xab, 0xd3,
	0xbc, 0x2b, 0x5c, 0x67, 0x21, 0xfc, 0x4a, 0xfa, 0xd0, 0xae, 0x18, 0xa9, 0x0f, 0xf8, 0x09, 0x00,
	0x7a, 0x48, 0x4a, 0x41, 0x9f, 0xbb, 0xb5, 0xbb, 0xc2, 0x2d, 0x2b, 0xab, 0xe2, 0x9e, 0x7f, 0x42,
	0x0f, 0x6c, 0x6b, 0x6e, 0x5b, 0x71, 0x57, 0xef, 0x0a, 0xd7, 0x4e, 0x68, 0xa4, 0x39, 0xb5, 0x4b,
	0x96, 0x8a, 0x91, 0x94, 0x0e, 0x49, 0xa8, 0x8e, 0x36, 0x1b, 0x4d, 0xa1, 0xf7, 0xd3, 0x16, 0xb0,
	0x6f, 0x46, 0x88, 0xf0, 0x3c, 0x11, 0xf0, 0x0b, 0xe0, 0x04, 0x34, 0x13, 0x0c, 0x07, 0xc2, 0x5f,
	0x2a, 0x6d, 0xe7, 0xf9, 0xfc, 0x98, 0x59, 0x8d, 0xf0, 0x50, 0x63, 0x6a, 0xba, 0x30, 0xf5, 0x6f,
	0x82, 0xed, 0x6e, 0x42, 0x69, 0xaa, 0xda, 0xa0, 0x8a, 0x34, 0x80, 0xdf, 0xaa, 0xaa, 0xa9, 0x25,
	0x2e, 0xa9, 0x3b, 0xf6, 0xe1, 0xca, 0x12, 0xaf, 0x34, 0x49, 0xe7, 0xc0, 0xdc, 0xb3, 0xeb, 0x5a,
	0xd8, 0x24, 0x7b, 0xb2, 0xb0, 0xaa, 0x89, 0x1c, 0x50, 0x62, 0x44, 0xa8, 0x15, 0xab, 0x22, 0xf9,
	0x09, 0x9f, 0x01, 0x9b, 0x91, 0x21, 0x61, 0x82, 0x84, 0x6a, 0x65, 0x6c, 0x34, 0xc3, 0xf0, 0x23,
	0x60, 0x47, 0x98, 0xfb, 0x39, 0x27, 0xa1, 0x5e, 0x06, 0xb4, 0x1b, 0x61, 0xfe, 0x1d, 0x27, 0xe1,
	0xe7, 0xd6, 0x4f, 0xff, 0x74, 0x37, 0x3c, 0x0c, 0x2a, 0x17, 0x41, 0x40, 0x38, 0xbf, 0xc9, 0x07,
	0x09, 0xf9, 0x95, 0xf6, 0x3a, 0x03, 0x55, 0x2e, 0x28, 0xc3, 0x11, 0xf1, 0xfb, 0x64, 0x6c, 0x9a,
	0x4c, 0xb7, 0x8c, 0xb1, 0xff, 0x81, 0x8c, 0x39, 0x5a, 0x04, 0x46, 0xe2, 0xef, 0x16, 0xa8, 0xdc,
	0x30, 0x1c, 0x10, 0x73, 0xb7, 0x97, 0x8d, 0x2a, 0x21, 0x33, 0x12, 0x06, 0x49, 0x6d, 0x11, 0xa7,
	0x84, 0xe6, 0xc2, 0xec, 0xa4, 0x29, 0x94, 0x19, 0x8c, 0x90, 0x11, 0x09, 0x54, 0x0d, 0x2d, 0x64,
	0x10, 0x3c, 0x07, 0xb5, 0x30, 0xe6, 0xea, 0x95, 0xc4, 0x05, 0x0e, 0xfa, 0x7a, 0xfa, 0x1d, 0xe7,
	0xae, 0x70, 0xab, 0xc6, 0xf1, 0x4e, 0xda, 0xd1, 0x12, 0x82, 0xbf, 0x05, 0x8d, 0x79, 0x9a, 0x1a,
	0xad, 0x7e, 0x9a, 0x74, 0xe0, 0x5d, 0xe1, 0xd6, 0x67, 0xa1, 0xca, 0x83, 0x56, 0xb0, 0x5c, 0xe6,
	0x90, 0x74, 0xf3, 0x48, 0x75, 0x9e, 0x8d, 0x34, 0x90, 0xd6, 0x24, 0x4e, 0x63, 0xa1, 0x3a, 0x6d,
	0x1b, 0x69, 0x00, 0x5f, 0x82, 0x32, 0x1d, 0x12, 0xc6, 0xe2, 0x90, 0x70, 0x75, 0xc9, 0xf9, 0xd5,
	0x27, 0x16, 0x9a, 0x07, 0xcb, 0x99, 0x99, 0xe7, 0x5f, 0x4a, 0x52, 0xca, 0xc6, 0xea, 0xca, 0x62,
	0x66, 0xa6, 0x1d, 0x5f, 0x2b, 0x3b, 0x5a, 0x42, 0xb0, 0x03, 0xa0, 0x49, 0x63, 0x44, 0xe4, 0x2c,
	0xf3, 0xd5, 0xce, 0xaf, 0xaa, 0x5c, 0xb5, 0xff, 0xb4, 0x17, 0x29, 0xe7, 0x6b, 0x2c, 0x30, 0xba,
	0x67, 0x81, 0xbf, 0x07, 0x50, 0x2f, 0x88, 0xff, 0x23, 0xa7, 0xb3, 0x07, 0xa2, 0xbe, 0x51, 0x28,
	0x7d, 0xed, 0x35, 0x63, 0x76, 0x34, 0xba, 0xe6, 0xd4, 0xcc, 0xe2, 0xda, 0xb2, 0x2d, 0x67, 0xfb,
	0xda, 0xb2, 0x77, 0x1d, 0x7b, 0x56, 0x3c, 0x33, 0x0b, 0xb4, 0x3f, 0xc5, 0x0b, 0xc3, 0xeb, 0x5c,
	0xfd, 0x7c, 0x7b, 0xb8, 0xf9, 0xcb, 0xed, 0xe1, 0xe6, 0x7f, 0x6f, 0x0f, 0x37, 0xff, 0xf6, 0xe1,
	0x70, 0xe3, 0x97, 0x0f, 0x87, 0x1b, 0xff, 0xfe, 0x70, 0xb8, 0xf1, 0x7d, 0x7b, 0xe1, 0x58, 0xd0,
	0x65, 0x7b, 0x91, 0x11, 0xf1, 0x9e, 0xb2, 0xbe, 0x81, 0xf2, 0xc9, 0x3f, 0x52, 0x6f, 0x7f, 0x75,
	0x46, 0x74, 0x77, 0xd4, 0xb3, 0xfe, 0xb3, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x6e, 0xd7,
	0xa7, 0x16, 0x10, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AllowUnprotectedTxs {
		i--
		if m.AllowUnprotectedTxs {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.ChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.ExtraEIPs) > 0 {
		dAtA3 := make([]byte, len(m.ExtraEIPs)*10)
		var j2 int
		for _, num1 := range m.ExtraEIPs {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintEvm(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0x22
	}
	if m.EnableCall {
		i--
		if m.EnableCall {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.EnableCreate {
		i--
		if m.EnableCreate {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.EvmDenom) > 0 {
		i -= len(m.EvmDenom)
		copy(dAtA[i:], m.EvmDenom)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.EvmDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CancunBlock != nil {
		{
			size := m.CancunBlock.Size()
			i -= size
			if _, err := m.CancunBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xba
	}
	if m.ShanghaiBlock != nil {
		{
			size := m.ShanghaiBlock.Size()
			i -= size
			if _, err := m.ShanghaiBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb2
	}
	if m.MergeNetsplitBlock != nil {
		{
			size := m.MergeNetsplitBlock.Size()
			i -= size
			if _, err := m.MergeNetsplitBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	if m.GrayGlacierBlock != nil {
		{
			size := m.GrayGlacierBlock.Size()
			i -= size
			if _, err := m.GrayGlacierBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	if m.ArrowGlacierBlock != nil {
		{
			size := m.ArrowGlacierBlock.Size()
			i -= size
			if _, err := m.ArrowGlacierBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if m.LondonBlock != nil {
		{
			size := m.LondonBlock.Size()
			i -= size
			if _, err := m.LondonBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if m.BerlinBlock != nil {
		{
			size := m.BerlinBlock.Size()
			i -= size
			if _, err := m.BerlinBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6a
	}
	if m.MuirGlacierBlock != nil {
		{
			size := m.MuirGlacierBlock.Size()
			i -= size
			if _, err := m.MuirGlacierBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if m.IstanbulBlock != nil {
		{
			size := m.IstanbulBlock.Size()
			i -= size
			if _, err := m.IstanbulBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.PetersburgBlock != nil {
		{
			size := m.PetersburgBlock.Size()
			i -= size
			if _, err := m.PetersburgBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.ConstantinopleBlock != nil {
		{
			size := m.ConstantinopleBlock.Size()
			i -= size
			if _, err := m.ConstantinopleBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.ByzantiumBlock != nil {
		{
			size := m.ByzantiumBlock.Size()
			i -= size
			if _, err := m.ByzantiumBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.EIP158Block != nil {
		{
			size := m.EIP158Block.Size()
			i -= size
			if _, err := m.EIP158Block.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.EIP155Block != nil {
		{
			size := m.EIP155Block.Size()
			i -= size
			if _, err := m.EIP155Block.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.EIP150Hash) > 0 {
		i -= len(m.EIP150Hash)
		copy(dAtA[i:], m.EIP150Hash)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.EIP150Hash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.EIP150Block != nil {
		{
			size := m.EIP150Block.Size()
			i -= size
			if _, err := m.EIP150Block.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.DAOForkSupport {
		i--
		if m.DAOForkSupport {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.DAOForkBlock != nil {
		{
			size := m.DAOForkBlock.Size()
			i -= size
			if _, err := m.DAOForkBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.HomesteadBlock != nil {
		{
			size := m.HomesteadBlock.Size()
			i -= size
			if _, err := m.HomesteadBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *State) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *State) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *State) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TransactionLogs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionLogs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionLogs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Logs) > 0 {
		for iNdEx := len(m.Logs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Logs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvm(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Log) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Log) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Log) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Removed {
		i--
		if m.Removed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.Index != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x40
	}
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0x3a
	}
	if m.TxIndex != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.TxIndex))
		i--
		dAtA[i] = 0x30
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.BlockNumber != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Topics) > 0 {
		for iNdEx := len(m.Topics) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Topics[iNdEx])
			copy(dAtA[i:], m.Topics[iNdEx])
			i = encodeVarintEvm(dAtA, i, uint64(len(m.Topics[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasUsed != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x30
	}
	if m.Reverted {
		i--
		if m.Reverted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Ret) > 0 {
		i -= len(m.Ret)
		copy(dAtA[i:], m.Ret)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Ret)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.TxLogs.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Bloom) > 0 {
		i -= len(m.Bloom)
		copy(dAtA[i:], m.Bloom)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Bloom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AccessTuple) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessTuple) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessTuple) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StorageKeys) > 0 {
		for iNdEx := len(m.StorageKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.StorageKeys[iNdEx])
			copy(dAtA[i:], m.StorageKeys[iNdEx])
			i = encodeVarintEvm(dAtA, i, uint64(len(m.StorageKeys[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TraceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TracerJsonConfig) > 0 {
		i -= len(m.TracerJsonConfig)
		copy(dAtA[i:], m.TracerJsonConfig)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.TracerJsonConfig)))
		i--
		dAtA[i] = 0x6a
	}
	if m.EnableReturnData {
		i--
		if m.EnableReturnData {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x60
	}
	if m.EnableMemory {
		i--
		if m.EnableMemory {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if m.Overrides != nil {
		{
			size, err := m.Overrides.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.Limit != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x48
	}
	if m.Debug {
		i--
		if m.Debug {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.DisableStorage {
		i--
		if m.DisableStorage {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.DisableStack {
		i--
		if m.DisableStack {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Reexec != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.Reexec))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Timeout) > 0 {
		i -= len(m.Timeout)
		copy(dAtA[i:], m.Timeout)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Timeout)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Tracer) > 0 {
		i -= len(m.Tracer)
		copy(dAtA[i:], m.Tracer)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Tracer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvm(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EvmDenom)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.EnableCreate {
		n += 2
	}
	if m.EnableCall {
		n += 2
	}
	if len(m.ExtraEIPs) > 0 {
		l = 0
		for _, e := range m.ExtraEIPs {
			l += sovEvm(uint64(e))
		}
		n += 1 + sovEvm(uint64(l)) + l
	}
	l = m.ChainConfig.Size()
	n += 1 + l + sovEvm(uint64(l))
	if m.AllowUnprotectedTxs {
		n += 2
	}
	return n
}

func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HomesteadBlock != nil {
		l = m.HomesteadBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.DAOForkBlock != nil {
		l = m.DAOForkBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.DAOForkSupport {
		n += 2
	}
	if m.EIP150Block != nil {
		l = m.EIP150Block.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.EIP150Hash)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.EIP155Block != nil {
		l = m.EIP155Block.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.EIP158Block != nil {
		l = m.EIP158Block.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.ByzantiumBlock != nil {
		l = m.ByzantiumBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.ConstantinopleBlock != nil {
		l = m.ConstantinopleBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.PetersburgBlock != nil {
		l = m.PetersburgBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.IstanbulBlock != nil {
		l = m.IstanbulBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.MuirGlacierBlock != nil {
		l = m.MuirGlacierBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.BerlinBlock != nil {
		l = m.BerlinBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.LondonBlock != nil {
		l = m.LondonBlock.Size()
		n += 2 + l + sovEvm(uint64(l))
	}
	if m.ArrowGlacierBlock != nil {
		l = m.ArrowGlacierBlock.Size()
		n += 2 + l + sovEvm(uint64(l))
	}
	if m.GrayGlacierBlock != nil {
		l = m.GrayGlacierBlock.Size()
		n += 2 + l + sovEvm(uint64(l))
	}
	if m.MergeNetsplitBlock != nil {
		l = m.MergeNetsplitBlock.Size()
		n += 2 + l + sovEvm(uint64(l))
	}
	if m.ShanghaiBlock != nil {
		l = m.ShanghaiBlock.Size()
		n += 2 + l + sovEvm(uint64(l))
	}
	if m.CancunBlock != nil {
		l = m.CancunBlock.Size()
		n += 2 + l + sovEvm(uint64(l))
	}
	return n
}

func (m *State) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	return n
}

func (m *TransactionLogs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovEvm(uint64(l))
		}
	}
	return n
}

func (m *Log) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if len(m.Topics) > 0 {
		for _, s := range m.Topics {
			l = len(s)
			n += 1 + l + sovEvm(uint64(l))
		}
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovEvm(uint64(m.BlockNumber))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.TxIndex != 0 {
		n += 1 + sovEvm(uint64(m.TxIndex))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovEvm(uint64(m.Index))
	}
	if m.Removed {
		n += 2
	}
	return n
}

func (m *TxResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Bloom)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = m.TxLogs.Size()
	n += 1 + l + sovEvm(uint64(l))
	l = len(m.Ret)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.Reverted {
		n += 2
	}
	if m.GasUsed != 0 {
		n += 1 + sovEvm(uint64(m.GasUsed))
	}
	return n
}

func (m *AccessTuple) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if len(m.StorageKeys) > 0 {
		for _, s := range m.StorageKeys {
			l = len(s)
			n += 1 + l + sovEvm(uint64(l))
		}
	}
	return n
}

func (m *TraceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tracer)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Timeout)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.Reexec != 0 {
		n += 1 + sovEvm(uint64(m.Reexec))
	}
	if m.DisableStack {
		n += 2
	}
	if m.DisableStorage {
		n += 2
	}
	if m.Debug {
		n += 2
	}
	if m.Limit != 0 {
		n += 1 + sovEvm(uint64(m.Limit))
	}
	if m.Overrides != nil {
		l = m.Overrides.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.EnableMemory {
		n += 2
	}
	if m.EnableReturnData {
		n += 2
	}
	l = len(m.TracerJsonConfig)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	return n
}

func sovEvm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvm(x uint64) (n int) {
	return sovEvm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvmDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EvmDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableCreate", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableCreate = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableCall", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableCall = bool(v != 0)
		case 4:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvm
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ExtraEIPs = append(m.ExtraEIPs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvm
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEvm
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEvm
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ExtraEIPs) == 0 {
					m.ExtraEIPs = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEvm
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ExtraEIPs = append(m.ExtraEIPs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraEIPs", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowUnprotectedTxs", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowUnprotectedTxs = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HomesteadBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.HomesteadBlock = &v
			if err := m.HomesteadBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DAOForkBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.DAOForkBlock = &v
			if err := m.DAOForkBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DAOForkSupport", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DAOForkSupport = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EIP150Block", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.EIP150Block = &v
			if err := m.EIP150Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EIP150Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EIP150Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EIP155Block", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.EIP155Block = &v
			if err := m.EIP155Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EIP158Block", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.EIP158Block = &v
			if err := m.EIP158Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ByzantiumBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.ByzantiumBlock = &v
			if err := m.ByzantiumBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConstantinopleBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.ConstantinopleBlock = &v
			if err := m.ConstantinopleBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PetersburgBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.PetersburgBlock = &v
			if err := m.PetersburgBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IstanbulBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.IstanbulBlock = &v
			if err := m.IstanbulBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MuirGlacierBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.MuirGlacierBlock = &v
			if err := m.MuirGlacierBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BerlinBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.BerlinBlock = &v
			if err := m.BerlinBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LondonBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.LondonBlock = &v
			if err := m.LondonBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArrowGlacierBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.ArrowGlacierBlock = &v
			if err := m.ArrowGlacierBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrayGlacierBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.GrayGlacierBlock = &v
			if err := m.GrayGlacierBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MergeNetsplitBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.MergeNetsplitBlock = &v
			if err := m.MergeNetsplitBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShanghaiBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.ShanghaiBlock = &v
			if err := m.ShanghaiBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 23:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CancunBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Int
			m.CancunBlock = &v
			if err := m.CancunBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *State) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: State: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: State: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TransactionLogs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TransactionLogs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionLogs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, &Log{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Log) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Log: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Log: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topics", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topics = append(m.Topics, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIndex", wireType)
			}
			m.TxIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Removed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Removed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TxResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TxResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bloom", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bloom = append(m.Bloom[:0], dAtA[iNdEx:postIndex]...)
			if m.Bloom == nil {
				m.Bloom = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxLogs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxLogs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ret = append(m.Ret[:0], dAtA[iNdEx:postIndex]...)
			if m.Ret == nil {
				m.Ret = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reverted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Reverted = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AccessTuple) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AccessTuple: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessTuple: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StorageKeys = append(m.StorageKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TraceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TraceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tracer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tracer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timeout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reexec", wireType)
			}
			m.Reexec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Reexec |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableStack", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableStack = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableStorage", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableStorage = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Debug", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Debug = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Overrides", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Overrides == nil {
				m.Overrides = &ChainConfig{}
			}
			if err := m.Overrides.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableMemory", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableMemory = bool(v != 0)
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableReturnData", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableReturnData = bool(v != 0)
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TracerJsonConfig", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TracerJsonConfig = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvm
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvm = fmt.Errorf("proto: unexpected end of group")
)
