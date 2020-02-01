
package control

import (
    "github.com/Myriad-Dreamin/minimum-lib/controller"
    "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
    "github.com/Myriad-Dreamin/go-ves/central-ves/model/db-layer"

)

var _ controller.MContext


type ChainInfoService interface {
    ChainInfoServiceSignatureXXX() interface{}
    ListChainInfos(c controller.MContext)
    PostChainInfo(c controller.MContext)
    InspectChainInfo(c controller.MContext)
    GetChainInfo(c controller.MContext)
    PutChainInfo(c controller.MContext)
    Delete(c controller.MContext)

}
type ListChainInfosRequest = gorm_crud_dao.Filter

type ListChainInfosReply struct {
    Code int `json:"code" form:"code"`
    ChainInfos []dblayer.ChainInfo `form:"chain_infos" json:"chain_infos"`
}

type PostChainInfoRequest struct {

}

type PostChainInfoReply struct {
    Code int `form:"code" json:"code"`
    ChainInfo *dblayer.ChainInfo `json:"chain_info" form:"chain_info"`
}

type InspectChainInfoReply struct {
    Code int `json:"code" form:"code"`
    ChainInfo *dblayer.ChainInfo `json:"chain_info" form:"chain_info"`
}

type GetChainInfoReply struct {
    Code int `json:"code" form:"code"`
    ChainInfo *dblayer.ChainInfo `json:"chain_info" form:"chain_info"`
}

type PutChainInfoRequest struct {

}
func PSerializeListChainInfosReply(_code int, _chainInfos []dblayer.ChainInfo) *ListChainInfosReply {

    return &ListChainInfosReply{
        Code: _code,
        ChainInfos: _chainInfos,
    }
}
func SerializeListChainInfosReply(_code int, _chainInfos []dblayer.ChainInfo) ListChainInfosReply {

    return ListChainInfosReply{
        Code: _code,
        ChainInfos: _chainInfos,
    }
}
func _packSerializeListChainInfosReply(_code int, _chainInfos []dblayer.ChainInfo) ListChainInfosReply {

    return ListChainInfosReply{
        Code: _code,
        ChainInfos: _chainInfos,
    }
}
func PackSerializeListChainInfosReply(_code []int, _chainInfos [][]dblayer.ChainInfo) (pack []ListChainInfosReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListChainInfosReply(_code[i], _chainInfos[i]))
	}
	return
}
func PSerializePostChainInfoRequest() *PostChainInfoRequest {

    return &PostChainInfoRequest{

    }
}
func SerializePostChainInfoRequest() PostChainInfoRequest {

    return PostChainInfoRequest{

    }
}
func _packSerializePostChainInfoRequest() PostChainInfoRequest {

    return PostChainInfoRequest{

    }
}
func PackSerializePostChainInfoRequest() (pack []PostChainInfoRequest) {
	return
}
func PSerializePostChainInfoReply(_code int, _chainInfo *dblayer.ChainInfo) *PostChainInfoReply {

    return &PostChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func SerializePostChainInfoReply(_code int, _chainInfo *dblayer.ChainInfo) PostChainInfoReply {

    return PostChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func _packSerializePostChainInfoReply(_code int, _chainInfo *dblayer.ChainInfo) PostChainInfoReply {

    return PostChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func PackSerializePostChainInfoReply(_code []int, _chainInfo []*dblayer.ChainInfo) (pack []PostChainInfoReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostChainInfoReply(_code[i], _chainInfo[i]))
	}
	return
}
func PSerializeInspectChainInfoReply(_code int, _chainInfo *dblayer.ChainInfo) *InspectChainInfoReply {

    return &InspectChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func SerializeInspectChainInfoReply(_code int, _chainInfo *dblayer.ChainInfo) InspectChainInfoReply {

    return InspectChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func _packSerializeInspectChainInfoReply(_code int, _chainInfo *dblayer.ChainInfo) InspectChainInfoReply {

    return InspectChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func PackSerializeInspectChainInfoReply(_code []int, _chainInfo []*dblayer.ChainInfo) (pack []InspectChainInfoReply) {
	for i := range _code {
		pack = append(pack, _packSerializeInspectChainInfoReply(_code[i], _chainInfo[i]))
	}
	return
}
func PSerializeGetChainInfoReply(_code int, _chainInfo *dblayer.ChainInfo) *GetChainInfoReply {

    return &GetChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func SerializeGetChainInfoReply(_code int, _chainInfo *dblayer.ChainInfo) GetChainInfoReply {

    return GetChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func _packSerializeGetChainInfoReply(_code int, _chainInfo *dblayer.ChainInfo) GetChainInfoReply {

    return GetChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func PackSerializeGetChainInfoReply(_code []int, _chainInfo []*dblayer.ChainInfo) (pack []GetChainInfoReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetChainInfoReply(_code[i], _chainInfo[i]))
	}
	return
}
func PSerializePutChainInfoRequest() *PutChainInfoRequest {

    return &PutChainInfoRequest{

    }
}
func SerializePutChainInfoRequest() PutChainInfoRequest {

    return PutChainInfoRequest{

    }
}
func _packSerializePutChainInfoRequest() PutChainInfoRequest {

    return PutChainInfoRequest{

    }
}
func PackSerializePutChainInfoRequest() (pack []PutChainInfoRequest) {
	return
}
