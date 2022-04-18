package main

import (
	"common/model"
	"common/tools"
	"math/big"
)

var nftMap = map[int64]*model.Nft{}
var openMap = map[int64]*model.Nft{}

func InitNft()  {
	nftMap = map[int64]*model.Nft{}
	openMap = map[int64]*model.Nft{}
}

/*
* @Description  read balance
*/
func getBalance(tokenId int64) string {
	return nftMap[tokenId].Balance
}

/*
* @Description  Possession of the read address
*/
func findClaimToken(address string, unix int64) (balance *big.Int, tokens []int64) {
	balance = big.NewInt(0)
	for _, nft := range nftMap {
		if nft.Address != address {
			continue
		}
		if tools.CanClaim(int64(nft.HaveTime), unix) < 0 {
			continue
		}
		b := tools.ToBig(nft.Balance)
		if b.Cmp(big.NewInt(0)) < 1 {
			continue
		}
		balance = tools.AddBig(nft.Balance, balance)
		tokens = append(tokens, nft.TokenId)
	}
	return
}


func shareIncome(value *big.Int)  {
	s := tools.DivBigBig(value,int64(len(openMap)))
	for _,n := range openMap{
		changeIncome(n,0,s)
	}
}

func changeIncome(n *model.Nft,mode int32,s *big.Int)  {
	if mode == 0{
		n.Income = tools.AddBig(n.Income,s).String()
		n.Balance = tools.AddBig(n.Balance,s).String()
	}else{
		n.Income = tools.SubBig(n.Income,s).String()
		n.Balance = tools.SubBig(n.Balance,s).String()
	}
}

/*
* @Description Active data processing
 */
func doActivity(activity *model.Activity)  {
	switch activity.Type {
	case model.ActivityTypeMint:
		income := tools.DivBigBig(tools.MulBig(activity.Income,75),100)
		oneIncome := tools.DivBigBig(income,int64(len(activity.TokenIds)))
		for _,tokenId := range activity.TokenIds{
			mintNft(&model.Nft{
				TokenId:tokenId,
				Address:tools.HexForDb(activity.Mint.To),
				BlockIndex:activity.Mint.BlockIndex,
				Source:model.NftSourceMint,
				From:"0x0000000000000000000000000000000000000000",
				Status:model.NftStatusNormal,
				HaveTime:uint64(activity.Time),
			},oneIncome)
		}
	case model.ActivityTypeRoyalties:
		income := tools.DivBigBig(tools.MulBig(activity.Income,55),100)
		royaltiesNft(income)
	case model.ActivityTypeTransfer:
		if activity.Transfer.To == activity.Transfer.From {
			return
		}
		transferNft(activity.TokenIds[0],tools.HexForDb(activity.Transfer.To),activity.Time)
	case model.ActivityTypeList:
		listNft(activity.TokenIds[0])
	case model.ActivityTypeCancelList:
		cancelListNft(activity.TokenIds[0])
	case model.ActivityTypeListStart:
		startListNft(activity.TokenIds[0])
	case model.ActivityTypeListEnd:
		endListNft(activity.TokenIds[0])
	case model.ActivityTypeClaim:
		if activity.Status != model.ActivityStatusVerifyed{
			return
		}
		for _,tokenId := range activity.TokenIds{
			claimNft(tokenId)
		}
	}
}

func mintNft(nft *model.Nft,income *big.Int)  {
	nft.Income =income.String()
	nft.Balance = nft.Income
	nft.Status = model.NftStatusNormal
	nftMap[nft.TokenId] = nft
	openMap[nft.TokenId] = nft
}

func royaltiesNft(income *big.Int)  {
	shareIncome(income)
}

func listNft(tokenId int64)  {
	nftMap[tokenId].Status = model.NftStatusList
}

func startListNft(tokenId int64)  {
	if nftMap[tokenId].Status != model.NftStatusStop {
		nftMap[tokenId].Status = model.NftStatusStop
	}
	delete(openMap,tokenId)
}

func endListNft(tokenId int64)  {
	openMap[tokenId] = nftMap[tokenId]
	nftMap[tokenId].Status = model.NftStatusNormal
}

func claimNft(tokenId int64)  {
	nftMap[tokenId].Balance = big.NewInt(0).String()
}

func cancelListNft(tokenId int64)  {
	openMap[tokenId] = nftMap[tokenId]
	nftMap[tokenId].Status = model.NftStatusNormal
}

func transferNft(tokenId int64,to string,haveTime int)  {
	nowNft := nftMap[tokenId]
	half := tools.DivBig(nowNft.Balance,2)
	balance := tools.SubBig(nowNft.Balance,half)
	delete(openMap,tokenId)
	shareIncome(half)
	nftMap[tokenId].HaveTime = uint64(haveTime)
	nftMap[tokenId].Income =  balance.String()
	nftMap[tokenId].Balance = nftMap[tokenId].Income
	nftMap[tokenId].Address = to
	nftMap[tokenId].Status = model.NftStatusNormal
	openMap[tokenId] = nftMap[tokenId]
}
