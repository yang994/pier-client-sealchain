package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/meshplus/pier/pkg/plugins"
	"os"
)

type Client struct {
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

func (c *Client) Initialize(configPath string, extra []byte, mode string) error {

	return nil
}

func (c *Client) Start() error {
	return nil
}

func (c *Client) Stop() error {
	return nil
}

func (c *Client) GetIBTPCh() chan *pb.IBTP {
	return nil
}

func (c *Client) GetUpdateMeta() chan *pb.UpdateMeta {
	return nil
}

func (c *Client) SubmitIBTP(from string, index uint64, serviceID string, ibtpType pb.IBTP_Type, content *pb.Content, proof *pb.BxhProof, isEncrypted bool) (*pb.SubmitIBTPResponse, error) {

	return nil, nil
}

func (c *Client) SubmitIBTPBatch(from []string, index []uint64, serviceID []string, ibtpType []pb.IBTP_Type, content []*pb.Content, proof []*pb.BxhProof, isEncrypted []bool) (*pb.SubmitIBTPResponse, error) {
	return nil, nil
}

func (c *Client) SubmitReceipt(to string, index uint64, serviceID string, ibtpType pb.IBTP_Type, result *pb.Result, proof *pb.BxhProof) (*pb.SubmitIBTPResponse, error) {
	return nil, nil
}

func (c *Client) SubmitReceiptBatch(to []string, index []uint64, serviceID []string, ibtpType []pb.IBTP_Type, result []*pb.Result, proof []*pb.BxhProof) (*pb.SubmitIBTPResponse, error) {
	return nil, nil
}

func (c *Client) GetOutMessage(servicePair string, idx uint64) (*pb.IBTP, error) {
	return nil, nil
}

func (c *Client) GetReceiptMessage(servicePair string, idx uint64) (*pb.IBTP, error) {
	return nil, nil
}

func (c *Client) GetInMeta() (map[string]uint64, error) {
	return nil, nil
}

func (c *Client) GetOutMeta() (map[string]uint64, error) {
	return nil, nil
}

func (c *Client) GetCallbackMeta() (map[string]uint64, error) {
	return nil, nil
}

func (c *Client) GetDstRollbackMeta() (map[string]uint64, error) {
	return nil, nil
}

func (c *Client) GetDirectTransactionMeta(string) (uint64, uint64, uint64, error) {
	return 0, 0, 0, nil
}

func (c *Client) GetServices() ([]string, error) {
	return nil, nil
}

func (c *Client) GetChainID() (string, string, error) {
	return "", "", nil
}

func (c *Client) GetAppchainInfo(chainID string) (string, []byte, string, error) {
	return "", nil, "", nil
}

func (c *Client) Name() string {
	return ""
}

func (c *Client) Type() string {
	return SealchainType
}

func (c *Client) GetOffChainData(request *pb.GetDataRequest) (*pb.GetDataResponse, error) {
	return nil, nil
}

func (c *Client) GetOffChainDataReq() chan *pb.GetDataRequest {
	return nil
}

func (c *Client) SubmitOffChainData(response *pb.GetDataResponse) error {
	return nil
}
