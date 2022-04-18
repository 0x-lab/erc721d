package model

const (
	NftSourceZero     NftSource = 0
	NftSourceMint     NftSource = 1 //mint
	NftSourceTransfer NftSource = 2 //transfer
)

type NftStatus int32

const (
	NftStatusZero   NftStatus = 0
	NftStatusNormal NftStatus = 1
	NftStatusStop   NftStatus = 2 //stop
	NftStatusList   NftStatus = 3 //list
)

type Nft struct {
	TokenId    int64     `bson:"tokenid"`
	ActivityIndex   int64    `bson:"activityindex"`
	BlockIndex uint64    `bson:"blockindex"`
	Address    string    `bson:"address"`
	Source     NftSource `bson:"source"` // 1:mint 2:transfer
	From       string    `bson:"from"`
	Status     NftStatus `bson:"status"` //0:income 1:stop（on list）
	Income     string    `bson:"income"`
	Balance    string    `bson:"balance"`
	HaveTime   uint64    `bson:"havetime"`
}

const (
	PostOpensea   = "opensea"
	PostLooksrare = "looksrare"
)

type ActivityType int32

const (
	ActivityTypeZero       ActivityType = 0
	ActivityTypeMint       ActivityType = 1 //mint
	ActivityTypeRoyalties  ActivityType = 2 //royalties
	ActivityTypeTransfer   ActivityType = 3 //transfer
	ActivityTypeList       ActivityType = 4 //list
	ActivityTypeCancelList ActivityType = 5 //CancelList
	ActivityTypeListStart  ActivityType = 6 //list start
	ActivityTypeListEnd    ActivityType = 7 //List end
	ActivityTypeClaim      ActivityType = 8 //claim
	ActivityTypeSale       ActivityType = 9 //sale
)

type ListPlatform int32

const (
	ListPlatformZero      ListPlatform = 0
	ListPlatformOpensea   ListPlatform = 1 //opensea
	ListPlatformLooksrare ListPlatform = 2 //looksrare
)

