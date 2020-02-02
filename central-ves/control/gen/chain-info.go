package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
)

type ChainInfoCategories struct {
	artisan.VirtualService
	List    artisan.Category
	Post    artisan.Category
	Inspect artisan.Category
	IdGroup artisan.Category
}

func DescribeChainInfoService(base string) artisan.ProposingService {
	var chainInfoModel = new(model.ChainInfo)
	var _chainInfoModel = new(model.ChainInfo)
	svc := &ChainInfoCategories{
		List: artisan.Ink().
			Path("chain_info-list").
			Method(artisan.POST, "ListChainInfos",
				artisan.QT("ListChainInfosRequest", model.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("chain_infos", _chainInfoModel)),
				),
			),
		Post: artisan.Ink().
			Path("chain_info").
			Method(artisan.POST, "PostChainInfo",
				artisan.Request(
					artisan.SPsC(&chainInfoModel.UserID, &chainInfoModel.Address, &chainInfoModel.ChainID),
				),
				artisan.Reply(
					codeField,
					artisan.Param("chain_info", &chainInfoModel),
				),
			),
		Inspect: artisan.Ink().Path("chain_info/:cid/inspect").
			Method(artisan.GET, "InspectChainInfo",
				artisan.Reply(
					codeField,
					artisan.Param("chain_info", &chainInfoModel),
				),
			),
		IdGroup: artisan.Ink().
			Path("chain_info/:cid").
			Method(artisan.GET, "GetChainInfo",
				artisan.Reply(
					codeField,
					artisan.Param("chain_info", &chainInfoModel),
				)).
			Method(artisan.PUT, "PutChainInfo",
				artisan.Request(
					&chainInfoModel.UserID, &chainInfoModel.Address, &chainInfoModel.ChainID),
			).
			Method(artisan.DELETE, "Delete"),
	}
	svc.Name("ChainInfoService").Base(base).
		UseModel(artisan.Model(artisan.Name("chain_info"), &chainInfoModel))
	return svc
}
