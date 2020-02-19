
package control

import (
    "github.com/Myriad-Dreamin/minimum-lib/controller"
    "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
    "github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"

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
    ChainInfos []database.ChainInfo `json:"chain_infos" form:"chain_infos"`
}

type PostChainInfoRequest struct {
    UserId uint `json:"user_id" form:"user_id"`
    Address string `json:"address" form:"address"`
    ChainId uint64 `json:"chain_id" form:"chain_id"`
}

type PostChainInfoReply struct {
    Code int `json:"code" form:"code"`
    ChainInfo *database.ChainInfo `json:"chain_info" form:"chain_info"`
}

type InspectChainInfoReply struct {
    Code int `json:"code" form:"code"`
    ChainInfo *database.ChainInfo `json:"chain_info" form:"chain_info"`
}

type GetChainInfoReply struct {
    Code int `json:"code" form:"code"`
    ChainInfo *database.ChainInfo `form:"chain_info" json:"chain_info"`
}

type PutChainInfoRequest struct {

}
func PSerializeListChainInfosReply(_code int, _chainInfos []database.ChainInfo) *ListChainInfosReply {

    return &ListChainInfosReply{
        Code: _code,
        ChainInfos: _chainInfos,
    }
}
func SerializeListChainInfosReply(_code int, _chainInfos []database.ChainInfo) ListChainInfosReply {

    return ListChainInfosReply{
        Code: _code,
        ChainInfos: _chainInfos,
    }
}
func _packSerializeListChainInfosReply(_code int, _chainInfos []database.ChainInfo) ListChainInfosReply {

    return ListChainInfosReply{
        Code: _code,
        ChainInfos: _chainInfos,
    }
}
func PackSerializeListChainInfosReply(_code []int, _chainInfos [][]database.ChainInfo) (pack []ListChainInfosReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListChainInfosReply(_code[i], _chainInfos[i]))
	}
	return
}
func PSerializePostChainInfoRequest(chain_info *database.ChainInfo) *PostChainInfoRequest {

    return &PostChainInfoRequest{
        UserId: chain_info.UserID,
        Address: chain_info.Address,
        ChainId: chain_info.ChainID,
    }
}
func SerializePostChainInfoRequest(chain_info *database.ChainInfo) PostChainInfoRequest {

    return PostChainInfoRequest{
        UserId: chain_info.UserID,
        Address: chain_info.Address,
        ChainId: chain_info.ChainID,
    }
}
func _packSerializePostChainInfoRequest(chain_info *database.ChainInfo) PostChainInfoRequest {

    return PostChainInfoRequest{
        UserId: chain_info.UserID,
        Address: chain_info.Address,
        ChainId: chain_info.ChainID,
    }
}
func PackSerializePostChainInfoRequest(chain_info []*database.ChainInfo) (pack []PostChainInfoRequest) {
	for i := range chain_info {
		pack = append(pack, _packSerializePostChainInfoRequest(chain_info[i]))
	}
	return
}
func PSerializePostChainInfoReply(_code int, _chainInfo *database.ChainInfo) *PostChainInfoReply {

    return &PostChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func SerializePostChainInfoReply(_code int, _chainInfo *database.ChainInfo) PostChainInfoReply {

    return PostChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func _packSerializePostChainInfoReply(_code int, _chainInfo *database.ChainInfo) PostChainInfoReply {

    return PostChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func PackSerializePostChainInfoReply(_code []int, _chainInfo []*database.ChainInfo) (pack []PostChainInfoReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostChainInfoReply(_code[i], _chainInfo[i]))
	}
	return
}
func PSerializeInspectChainInfoReply(_code int, _chainInfo *database.ChainInfo) *InspectChainInfoReply {

    return &InspectChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func SerializeInspectChainInfoReply(_code int, _chainInfo *database.ChainInfo) InspectChainInfoReply {

    return InspectChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func _packSerializeInspectChainInfoReply(_code int, _chainInfo *database.ChainInfo) InspectChainInfoReply {

    return InspectChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func PackSerializeInspectChainInfoReply(_code []int, _chainInfo []*database.ChainInfo) (pack []InspectChainInfoReply) {
	for i := range _code {
		pack = append(pack, _packSerializeInspectChainInfoReply(_code[i], _chainInfo[i]))
	}
	return
}
func PSerializeGetChainInfoReply(_code int, _chainInfo *database.ChainInfo) *GetChainInfoReply {

    return &GetChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func SerializeGetChainInfoReply(_code int, _chainInfo *database.ChainInfo) GetChainInfoReply {

    return GetChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func _packSerializeGetChainInfoReply(_code int, _chainInfo *database.ChainInfo) GetChainInfoReply {

    return GetChainInfoReply{
        Code: _code,
        ChainInfo: _chainInfo,
    }
}
func PackSerializeGetChainInfoReply(_code []int, _chainInfo []*database.ChainInfo) (pack []GetChainInfoReply) {
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
