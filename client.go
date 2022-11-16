package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/meshplus/pier/pkg/plugins"
	"os"
)

type Client struct {
	name string
}

var (
	_      plugins.Client = (*Client)(nil)
	logger                = hclog.New(&hclog.LoggerOptions{
		Name:   "client",
		Output: os.Stderr,
		Level:  hclog.Trace,
	})
	SealchainType = "sealchain"
)

//初始化插件
func (c *Client) Initialize(configPath string, extra []byte, mode string) error {

	return nil
}

//启动插件
func (c *Client) Start() error {
	return nil
}

//停止插件
func (c *Client) Stop() error {
	return nil
}

//将区块链上产生的跨链事件转化为标准的IBTP格式，Pier通过GetIBTP接口获取跨链请求再进行处理,
//返回一个事件通道，需要监听broker对应合约事件
func (c *Client) GetIBTPCh() chan *pb.IBTP {
	return nil
}

//区块链上信任根变更的信息传递给中继链,暂不实现
func (c *Client) GetUpdateMeta() chan *pb.UpdateMeta {
	return nil
}

//执行跨链请求 调用broker合约invokeInterchain方法,将入参传入合约
//构造一个插件中的pb结构返回
func (c *Client) SubmitIBTP(from string, index uint64, serviceID string, ibtpType pb.IBTP_Type, content *pb.Content, proof *pb.BxhProof, isEncrypted bool) (*pb.SubmitIBTPResponse, error) {

	InvokeInterchain(from, index, serviceID, uint64(ibtpType), content.Func, content.Args, uint64(proof.TxStatus), proof.MultiSign, isEncrypted)
	return nil, nil
}

//不明用处 暂不实现
func (c *Client) SubmitIBTPBatch(from []string, index []uint64, serviceID []string, ibtpType []pb.IBTP_Type, content []*pb.Content, proof []*pb.BxhProof, isEncrypted []bool) (*pb.SubmitIBTPResponse, error) {
	panic("implement me")
}

//调用broker合约InvokeReceipt方法
//构造一个插件中的pb结构返回
func (c *Client) SubmitReceipt(to string, index uint64, serviceID string, ibtpType pb.IBTP_Type, result *pb.Result, proof *pb.BxhProof) (*pb.SubmitIBTPResponse, error) {
	InvokeReceipt(serviceID, to, index, uint64(ibtpType), result.Data, uint64(proof.TxStatus), proof.MultiSign)
	return nil, nil
}

//不明用处 暂不实现
func (c *Client) SubmitReceiptBatch(to []string, index []uint64, serviceID []string, ibtpType []pb.IBTP_Type, result []*pb.Result, proof []*pb.BxhProof) (*pb.SubmitIBTPResponse, error) {
	panic("implement me")
}

//调用broker合约getOutMessage方法 查询历史跨链请求。查询键值中servicePair指定服务对，idx指定序号，查询结果为以Plugin负责的区块链作为来源链的跨链请求。
func (c *Client) GetOutMessage(servicePair string, idx uint64) (*pb.IBTP, error) {
	return nil, nil
}

//调用broker合约getInMessage方法 查询历史跨链回执。查询键值中servicePair指定服务对，idx指定序号，查询结果为以Plugin负责的区块链作为目的链的跨链回执。
func (c *Client) GetReceiptMessage(servicePair string, idx uint64) (*pb.IBTP, error) {
	return nil, nil
}

//调用合约getInnerMeta方法 是获取跨链回执相关的Meta信息的接口。以Plugin负责的区块链为目的链服务的一系列跨链请求的序号信息。
//如果Plugin负责A链服务，则A可能和多条链服务进行跨链，如A->B:3; A->C:5。
//返回的map中，key值为服务对（来源链服务ID+目的链服务ID），value对应已发送到该目的链的最新跨链请求的序号，如{A+B:3, A+C:5}。
func (c *Client) GetInMeta() (map[string]uint64, error) {
	return nil, nil
}

//调用broker合约GetOutMeta
func (c *Client) GetOutMeta() (map[string]uint64, error) {
	return nil, nil
}

//调用broker合约getCallbackMeta
func (c *Client) GetCallbackMeta() (map[string]uint64, error) {
	return nil, nil
}

//暂不实现
func (c *Client) GetDstRollbackMeta() (map[string]uint64, error) {
	panic("implement me")
}

//暂不实现
func (c *Client) GetDirectTransactionMeta(string) (uint64, uint64, uint64, error) {
	panic("implement me")
}

//暂不实现
func (c *Client) GetServices() ([]string, error) {
	panic("implement me")
}

//暂不实现
func (c *Client) GetChainID() (string, string, error) {
	panic("implement me")
}

//暂不实现
func (c *Client) GetAppchainInfo(chainID string) (string, []byte, string, error) {
	panic("implement me")
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) Type() string {
	return SealchainType
}

//暂不实现
func (c *Client) GetOffChainData(request *pb.GetDataRequest) (*pb.GetDataResponse, error) {
	panic("implement me")
}

//暂不实现
func (c *Client) GetOffChainDataReq() chan *pb.GetDataRequest {
	panic("implement me")
}

//暂不实现
func (c *Client) SubmitOffChainData(response *pb.GetDataResponse) error {
	panic("implement me")
}
