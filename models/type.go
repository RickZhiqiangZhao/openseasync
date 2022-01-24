package models

import "math/big"

type User struct {
	Id               string `json:"id" bson:"id"`                             // 用户 ID
	UserMetamaskID   string `json:"userMetamaskId" bson:"userMetamaskId"`     // 用户地址
	Username         string `json:"userName" bson:"userName"`                 // 用户头像
	AvatarUrl        string `json:"avatarUrl" bson:"avatarUrl"`               // 用户头像
	Discord          string `json:"discord" bson:"discord"`                   // discord链接
	Telegram         string `json:"telegram" bson:"telegram"`                 // telegram链接
	Instagram        string `json:"instagram" bson:"instagram"`               // instagram链接
	Twitter          string `json:"twitter" bson:"twitter"`                   // twitter链接
	PersonalPageLink string `json:"personalPageLink" bson:"personalPageLink"` // 个人链接
}

type Asset struct {
	Id                  int    `json:"id" bson:"id"`                           // NFT ID
	UserMetamaskID      string `json:"userMetamaskId" bson:"userMetamaskId"`   // 用户地址
	Name                string `json:"name" bson:"name"`                       // NFT作品标题
	CoverImageUrl       string `json:"coverImageUrl " bson:"coverImageUrl "`   // 封面图片
	CoverPreviewUrl     string `json:"coverPreviewUrl" bson:"coverPreviewUrl"` // NFT作品原图
	ThumbnailUrl        string `json:"thumbnailUrl" bson:"thumbnailUrl"`       // NFT作品缩略图
	FileUrl             string `json:"fileUrl" bson:"fileUrl"`
	Description         string `json:"description" bson:"description"`                 // NFT作品描述
	ContractAddress     string `json:"contractAddress" bson:"contractAddress"`         // 合约地址
	CollectibleTokenId  string `json:"collectibleTokenId" bson:"collectibleTokenId"`   // NFT token id
	NumSales            int    `json:"numSales" bson:"numSales"`                       // NFT售卖次数
	OwnerMetamaskId     string `json:"ownerMetamaskId" bson:"ownerMetamaskId"`         // NFT拥有者地址
	OwnerName           string `json:"ownerName" bson:"ownerName"`                     // NFT拥有者名字
	OwnerImgURL         string `json:"ownerImgUrl" bson:"ownerImgUrl"`                 // 拥有者头像
	CreatorMetamaskId   string `json:"creatorMetamaskId" bson:"creatorMetamaskId"`     // NFT创造者地址
	CreatorPersonalSite string `json:"creatorPersonalSite" bson:"creatorPersonalSite"` // NFT创造者个人网站
	CreatorName         string `json:"creatorName" bson:"creatorName"`                 // NFT创造者名字
	CreatorImgUrl       string `json:"creatorImgUrl" bson:"creatorImgUrl"`             // 创造者头像
	TokenMetadata       string `json:"tokenMetadata" bson:"tokenMetadata"`             // NFT元数据

	CollectionID string `json:"collectionId" bson:"collectionId"` // 集合 id
	RecordId     string `json:"recordId" bson:"recordId"`         // 售卖记录 id

	AssetsTopOwnerships []AssetsTopOwnership `json:"assetsTopOwnerships" bson:"assetsTopOwnerships"`
	Traits              []Trait              `json:"traits" bson:"traits"`

	SellOrders SellOrder `json:"sellOrders" bson:"sellOrders"`

	NumOfCopies int `json:"numOfCopies" bson:"numOfCopies"`
	TotalCopies int `json:"totalCopies " bson:"totalCopies "`

	Price string `json:"price" bson:"price"`

	LikesCount int `json:"likesCount" bson:"likesCount"`
	ViewsCount int `json:"viewsCount" bson:"viewsCount"`

	Status string `json:"status" bson:"status"`

	IsDelete    int8        `json:"isDelete" bson:"isDelete"`       // 是否删除 1删除 0未删除 默认为0
	RefreshTime int64       `json:"refreshTime" bson:"refreshTime"` // 刷新时间
	CreateDate  int64       `json:"createDate" bson:"createDate"`   // 创建时间
	StartTime   interface{} `json:"startTime" bson:"startTime"`     // 创建时间
	EndTime     interface{} `json:"endTime" bson:"endTime"`         // 结束时间
}

type Collection struct {
	ID                 string  `json:"id" bson:"id"`                                 // 集合ID
	UserMetamaskID     string  `json:"userMetamaskId" bson:"userMetamaskId"`         // 集合拥有者
	CollectionName     string  `json:"collectionName" bson:"collectionName"`         // 集合名称
	UserCoverUrl       string  `json:"userCoverUrl" bson:"userCoverUrl"`             // 集合背景图
	Description        string  `json:"description" bson:"description"`               // 集合描述
	CoverImageUrl      string  `json:"coverImageUrl " bson:"coverImageUrl "`         // 封面图片
	CoverLargeImageUrl string  `json:"coverLargeImageURL" bson:"coverLargeImageURL"` // 头像大图
	IsDelete           int8    `json:"isDelete" bson:"isDelete"`                     // 是否删除 1删除 0未删除 默认为0
	CreateDate         string  `json:"createDate" bson:"createDate"`                 // 集合创建时间
	RefreshTime        int64   `json:"refreshTime" bson:"refreshTime"`               // 刷新时间
	ItemsCount         int     `json:"itemsCount" bson:"itemsCount"`                 // 集合中NFT总数
	TotalVolume        float64 `json:"totalVolume" bson:"totalVolume"`               // 交易量
	FloorPrice         float64 `json:"floorPrice" bson:"floorPrice"`                 // 最低价格
	HighestPrice       float64 `json:"highestPrice" bson:"highestPrice"`             // 最高价格
	OwnersCount        int     `json:"ownersCount" bson:"ownersCount"`               // 所有NFT中属于自己的NFT个数 此地段可能是个big int, 所以采用string存储
	LikesCount         int     `json:"likesCount" bson:"likesCount"`
	ViewsCount         int     `json:"viewsCount" bson:"viewsCount"`
}

type Contract struct {
	Address      string `json:"address" bson:"address"`           // 合约地址
	ContractType string `json:"contractType" bson:"contractType"` // 合约类型 semi-fungible可替代 non-fungible 不可替代
	ContractName string `json:"contractName" bson:"contractName"` // 合约名字
	Symbol       string `json:"symbol" bson:"symbol"`             // 符号
	SchemaName   string `json:"schemaName" bson:"schemaName"`     // 合约类型
	TotalSupply  string `json:"totalSupply" bson:"totalSupply"`   // 总供应量
	Description  string `json:"description" bson:"description"`   // 合约描述
}

type SellOrder struct {
	CreateDate   string `json:"createDate" bson:"createDate"`     // 创建时间
	ClosingDate  string `json:"closingDate" bson:"closingDate"`   // 结束时间
	CurrentPrice string `json:"currentPrice" bson:"currentPrice"` // 当前价格

	PayTokenContract PayTokenContract `json:"payTokenContract" bson:"payTokenContract"` // 支付方式
}

type Orders struct {
	UUID              string           `json:"uuid" bson:"uuid"`
	Id                string           `json:"id" bson:"id"`                       // 订单hash
	CollectibleId     int              `json:"collectibleId" bson:"collectibleId"` // NFT id
	CreateDate        string           `json:"createDate" bson:"createDate"`       // 创建时间
	BidTime           string           `json:"bidTime" bson:"bidTime"`             // 投标时间
	ClosingDate       string           `json:"closingDate" bson:"closingDate"`     // 结束时间
	Price             string           `json:"price" bson:"price"`                 // 当前价格
	BasePrice         string           `json:"basePrice" bson:"basePrice"`         // 基础价格
	CurrentBounty     string           `json:"currentBounty" bson:"currentBounty"`
	AuctionUserId     string           `json:"auctionUserId" bson:"auctionUserId"`
	AuctionMetamaskId string           `json:"auctionMetamaskId" bson:"auctionMetamaskId"`
	AuctionUserName   string           `json:"auctionUserName" bson:"auctionUserName"`
	PayTokenContract  PayTokenContract `json:"payTokenContract" bson:"payTokenContract"` // 支付方式
	IsDelete          int8             `json:"isDelete" bson:"isDelete"`                 // 是否删除 1删除 0未删除 默认为0
}

type PayTokenContract struct {
	Symbol   string `json:"symbol" bson:"symbol"`
	ImageURL string `json:"imageURL" bson:"imageURL"`
	EthPrice string `json:"ethPrice" bson:"ethPrice"`
	UsdPrice string `json:"usdPrice" bson:"usdPrice"`
}

type Trait struct {
	//ID              int64  `json:"id" bson:"id"`                     // 主键
	UserMetamaskID  string      `json:"userMetamaskId" bson:"userMetamaskId"` // 用户地址
	ContractAddress string      `json:"_" bson:"contractAddress"`             // 合约地址
	TokenId         string      `json:"_" bson:"token_id"`                    // token id
	TraitType       string      `json:"traitType" bson:"traitType"`           // 特征类型
	Value           interface{} `json:"value" bson:"value"`                   // 特征值
	DisplayType     string      `json:"displayType" bson:"displayType"`
	MaxValue        interface{} `json:"maxValue" bson:"maxValue"`
	TraitCount      int         `json:"traitCount" bson:"traitCount"` // 数量
	OrderBy         string      `json:"orderBy" bson:"orderBy"`
	IsDelete        int8        `json:"isDelete" bson:"isDelete"`       // 是否删除 1删除 0未删除 默认为0
	RefreshTime     int64       `json:"refreshTime" bson:"refreshTime"` // 刷新时间
}

type AssetsTopOwnership struct {
	//ID              int64  `json:"id" bson:"id"`                           // 主键
	UserMetamaskID  string `json:"userMetamaskId" bson:"userMetamaskId"` // 用户地址
	ContractAddress string `json:"_" bson:"contractAddress"`             // 合约地址
	TokenId         string `json:"_" bson:"token_id"`                    // token id
	Owner           string `json:"owner" bson:"owner"`                   // 所有者地址
	ProfileImgUrl   string `json:"profileImgUrl" bson:"profileImgUrl"`   // 所有者头像
	Quantity        string `json:"quantity" bson:"quantity"`             // 数量
	IsDelete        int8   `json:"isDelete" bson:"isDelete"`             // 是否删除 1删除 0未删除 默认为0
	RefreshTime     int64  `json:"refreshTime" bson:"refreshTime"`       // 刷新时间
}

type ItemActivity struct {
	Id               int              `json:"id" bson:"id"`
	CollectibleId    int              `json:"collectibleId" bson:"collectibleId"`     // NFT id
	CollectibleName  string           `json:"collectibleName" bson:"collectibleName"` // NFT 名字
	CollectionId     string           `json:"collectionId" bson:"collectionId"`       // 集合ID
	CollectionName   string           `json:"collectionName" bson:"collectionName"`   // 集合名
	ContractAddress  string           `json:"contractAddress" bson:"contractAddress"`
	TokenId          string           `json:"tokenId" bson:"tokenId"`
	BidAmount        string           `json:"bidAmount" bson:"bidAmount"` // 投标金额
	CreateDate       int64            `json:"createDate" bson:"createDate"`
	Price            string           `json:"price" bson:"price"`                       // 成交价格ETH
	PriceInUsd       string           `json:"priceInUsd " bson:"priceInUsd "`           // 成交价格USD
	SellerId         int              `json:"sellerId" bson:"sellerId"`                 // 售卖者ID
	SellerMetamaskId string           `json:"sellerMetamaskId" bson:"sellerMetamaskId"` // 售卖者地址
	SellerName       string           `json:"sellerName" bson:"sellerName"`             // 售卖者名字
	SellerImgUrl     string           `json:"sellerImgUrl" bson:"sellerImgUrl"`         // 售卖者头像
	BuyerId          int              `json:"buyerId" bson:"buyerId"`                   // 购买者ID
	BuyerMetamaskId  string           `json:"buyerMetamaskId" bson:"buyerMetamaskId"`   // 购买者地址
	BuyerName        string           `json:"buyerName" bson:"buyerName"`               // 购买者名字
	BuyerImgURL      string           `json:"buyerImgUrl" bson:"buyerImgUrl"`           // 购买者头像
	Quantity         string           `json:"quantity" bson:"quantity"`                 // 数量
	IsDelete         int8             `json:"isDelete" bson:"isDelete"`                 // 是否删除 1删除 0未删除 默认为0
	TradeType        string           `json:"tradeType " bson:"tradeType"`              // 事件类型
	PayTokenContract PayTokenContract `json:"payTokenContract" bson:"payTokenContract"` // 支付方式
	Transaction      struct {
		BlockHash   string `json:"block_hash" bson:"block_hash"`
		BlockNumber string `json:"block_number" bson:"block_number"`
		FromAccount struct {
			User struct {
				Username interface{} `json:"username" bson:"username"`
			} `json:"user" bson:"user"`
			ProfileImgURL string `json:"profile_img_url" bson:"profile_img_url"`
			Address       string `json:"address" bson:"address"` // 支付人
			Config        string `json:"config" bson:"config"`
		} `json:"from_account" bson:"from_account"`
		ID        int    `json:"id" bson:"id"`
		Timestamp string `json:"timestamp" bson:"timestamp"`
		ToAccount struct {
			User          interface{} `json:"user" bson:"user"`
			ProfileImgURL string      `json:"profile_img_url" bson:"profile_img_url"`
			Address       string      `json:"address" bson:"address"` // 支付对象 合约地址
			Config        string      `json:"config" bson:"config"`
		} `json:"to_account" bson:"to_account"`
		TransactionHash  string `json:"transaction_hash" bson:"transaction_hash"`
		TransactionIndex string `json:"transaction_index" bson:"transaction_index"`
	} `json:"transaction" bson:"transaction"` // 支付的eth链上交易记录
}

type OwnerAsset struct {
	Assets []AutoAsset `json:"assets"`
}

type OwnerCollection struct {
	Collections []AutoCollection `json:"collections"`
}

type Event struct {
	AssetEvents []AutoEvent `json:"asset_events"`
}

type AutoAsset struct {
	ID                   int    `json:"id"`
	TokenID              string `json:"token_id"`
	NumSales             int    `json:"num_sales"`
	BackgroundColor      string `json:"background_color"`
	ImageURL             string `json:"image_url"`
	ImagePreviewURL      string `json:"image_preview_url"`
	ImageThumbnailURL    string `json:"image_thumbnail_url"`
	ImageOriginalURL     string `json:"image_original_url"`
	AnimationURL         string `json:"animation_url"`
	AnimationOriginalURL string `json:"animation_original_url"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	ExternalLink         string `json:"external_link"`
	AssetContract        struct {
		Address                     string `json:"address"`
		AssetContractType           string `json:"asset_contract_type"`
		CreatedDate                 string `json:"created_date"`
		Name                        string `json:"name"`
		NftVersion                  string `json:"nft_version"`
		OpenseaVersion              string `json:"opensea_version"`
		Owner                       int    `json:"owner"`
		SchemaName                  string `json:"schema_name"`
		Symbol                      string `json:"symbol"`
		TotalSupply                 string `json:"total_supply"`
		Description                 string `json:"description"`
		ExternalLink                string `json:"external_link"`
		ImageURL                    string `json:"image_url"`
		DefaultToFiat               bool   `json:"default_to_fiat"`
		DevBuyerFeeBasisPoints      int    `json:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints     int    `json:"dev_seller_fee_basis_points"`
		OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  int    `json:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints int    `json:"opensea_seller_fee_basis_points"`
		BuyerFeeBasisPoints         int    `json:"buyer_fee_basis_points"`
		SellerFeeBasisPoints        int    `json:"seller_fee_basis_points"`
		PayoutAddress               string `json:"payout_address"`
	} `json:"asset_contract"`
	Permalink  string `json:"permalink"`
	Collection struct {
		BannerImageURL          string `json:"banner_image_url"`
		ChatURL                 string `json:"chat_url"`
		CreatedDate             string `json:"created_date"`
		DefaultToFiat           bool   `json:"default_to_fiat"`
		Description             string `json:"description"`
		DevBuyerFeeBasisPoints  string `json:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints string `json:"dev_seller_fee_basis_points"`
		DiscordURL              string `json:"discord_url"`
		DisplayData             struct {
			CardDisplayStyle string `json:"card_display_style"`
		} `json:"display_data"`
		ExternalURL                 string `json:"external_url"`
		Featured                    bool   `json:"featured"`
		FeaturedImageURL            string `json:"featured_image_url"`
		Hidden                      bool   `json:"hidden"`
		SafelistRequestStatus       string `json:"safelist_request_status"`
		ImageURL                    string `json:"image_url"`
		IsSubjectToWhitelist        bool   `json:"is_subject_to_whitelist"`
		LargeImageURL               string `json:"large_image_url"`
		MediumUsername              string `json:"medium_username"`
		Name                        string `json:"name"`
		OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  string `json:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints string `json:"opensea_seller_fee_basis_points"`
		PayoutAddress               string `json:"payout_address"`
		RequireEmail                bool   `json:"require_email"`
		ShortDescription            string `json:"short_description"`
		Slug                        string `json:"slug"`
		TelegramURL                 string `json:"telegram_url"`
		TwitterUsername             string `json:"twitter_username"`
		InstagramUsername           string `json:"instagram_username"`
		WikiURL                     string `json:"wiki_url"`
	} `json:"collection"`
	Decimals      int    `json:"decimals"`
	TokenMetadata string `json:"token_metadata"`
	Owner         struct {
		User struct {
			Username string `json:"username"`
		} `json:"user"`
		ProfileImgURL string `json:"profile_img_url"`
		Address       string `json:"address"`
		Config        string `json:"config"`
	} `json:"owner"`
	LastSale struct {
		Asset struct {
			TokenID  string `json:"token_id"`
			Decimals int    `json:"decimals"`
		} `json:"asset"`
		AssetBundle    interface{} `json:"asset_bundle"`
		EventType      string      `json:"event_type"`
		EventTimestamp string      `json:"event_timestamp"`
		AuctionType    interface{} `json:"auction_type"`
		TotalPrice     string      `json:"total_price"`
		PaymentToken   struct {
			ID       int         `json:"id"`
			Symbol   string      `json:"symbol"`
			Address  string      `json:"address"`
			ImageURL string      `json:"image_url"`
			Name     interface{} `json:"name"`
			Decimals int         `json:"decimals"`
			EthPrice string      `json:"eth_price"`
			UsdPrice string      `json:"usd_price"`
		} `json:"payment_token"`
		Transaction struct {
			BlockHash   string `json:"block_hash"`
			BlockNumber string `json:"block_number"`
			FromAccount struct {
				User struct {
					Username interface{} `json:"username"`
				} `json:"user"`
				ProfileImgURL string `json:"profile_img_url"`
				Address       string `json:"address"`
				Config        string `json:"config"`
			} `json:"from_account"`
			ID        int    `json:"id"`
			Timestamp string `json:"timestamp"`
			ToAccount struct {
				User          interface{} `json:"user"`
				ProfileImgURL string      `json:"profile_img_url"`
				Address       string      `json:"address"`
				Config        string      `json:"config"`
			} `json:"to_account"`
			TransactionHash  string `json:"transaction_hash"`
			TransactionIndex string `json:"transaction_index"`
		} `json:"transaction"`
		CreatedDate string `json:"created_date"`
		Quantity    string `json:"quantity"`
	} `json:"last_sale"`
	SellOrders []struct {
		CreatedDate       string `json:"created_date"`
		ClosingDate       string `json:"closing_date"`
		ClosingExtendable bool   `json:"closing_extendable"`
		ExpirationTime    int    `json:"expiration_time"`
		ListingTime       int    `json:"listing_time"`
		OrderHash         string `json:"order_hash"`
		Metadata          struct {
			Asset struct {
				ID      string `json:"id"`
				Address string `json:"address"`
			} `json:"asset"`
			Schema string `json:"schema"`
		} `json:"metadata"`
		Exchange string `json:"exchange"`
		Maker    struct {
			User          int    `json:"user"`
			ProfileImgURL string `json:"profile_img_url"`
			Address       string `json:"address"`
			Config        string `json:"config"`
		} `json:"maker"`
		Taker struct {
			User          int    `json:"user"`
			ProfileImgURL string `json:"profile_img_url"`
			Address       string `json:"address"`
			Config        string `json:"config"`
		} `json:"taker"`
		CurrentPrice     string `json:"current_price"`
		CurrentBounty    string `json:"current_bounty"`
		BountyMultiple   string `json:"bounty_multiple"`
		MakerRelayerFee  string `json:"maker_relayer_fee"`
		TakerRelayerFee  string `json:"taker_relayer_fee"`
		MakerProtocolFee string `json:"maker_protocol_fee"`
		TakerProtocolFee string `json:"taker_protocol_fee"`
		MakerReferrerFee string `json:"maker_referrer_fee"`
		FeeRecipient     struct {
			User          int    `json:"user"`
			ProfileImgURL string `json:"profile_img_url"`
			Address       string `json:"address"`
			Config        string `json:"config"`
		} `json:"fee_recipient"`
		FeeMethod            int    `json:"fee_method"`
		Side                 int    `json:"side"`
		SaleKind             int    `json:"sale_kind"`
		Target               string `json:"target"`
		HowToCall            int    `json:"how_to_call"`
		Calldata             string `json:"calldata"`
		ReplacementPattern   string `json:"replacement_pattern"`
		StaticTarget         string `json:"static_target"`
		StaticExtradata      string `json:"static_extradata"`
		PaymentToken         string `json:"payment_token"`
		PaymentTokenContract struct {
			ID       int    `json:"id"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
			Name     string `json:"name"`
			Decimals int    `json:"decimals"`
			EthPrice string `json:"eth_price"`
			UsdPrice string `json:"usd_price"`
		} `json:"payment_token_contract"`
		BasePrice       string `json:"base_price"`
		Extra           string `json:"extra"`
		Quantity        string `json:"quantity"`
		Salt            string `json:"salt"`
		V               int    `json:"v"`
		R               string `json:"r"`
		S               string `json:"s"`
		ApprovedOnChain bool   `json:"approved_on_chain"`
		Cancelled       bool   `json:"cancelled"`
		Finalized       bool   `json:"finalized"`
		MarkedInvalid   bool   `json:"marked_invalid"`
		PrefixedHash    string `json:"prefixed_hash"`
	} `json:"sell_orders"`
	Orders []struct {
		CreatedDate       string `json:"created_date"`
		ClosingDate       string `json:"closing_date"`
		ClosingExtendable bool   `json:"closing_extendable"`
		ExpirationTime    int    `json:"expiration_time"`
		ListingTime       int    `json:"listing_time"`
		OrderHash         string `json:"order_hash"`
		Metadata          struct {
			Asset struct {
				ID       string `json:"id"`
				Address  string `json:"address"`
				Quantity string `json:"quantity"`
			} `json:"asset"`
			Schema string `json:"schema"`
		} `json:"metadata"`
		Exchange string `json:"exchange"`
		Maker    struct {
			User struct {
				Username string `json:"username"`
			} `json:"user"`
			ProfileImgURL string `json:"profile_img_url"`
			Address       string `json:"address"`
			Config        string `json:"config"`
		} `json:"maker"`
		Taker struct {
			User struct {
				Username string `json:"username"`
			} `json:"user"`
			ProfileImgURL string `json:"profile_img_url"`
			Address       string `json:"address"`
			Config        string `json:"config"`
		} `json:"taker"`
		CurrentPrice     string `json:"current_price"`
		CurrentBounty    string `json:"current_bounty"`
		BountyMultiple   string `json:"bounty_multiple"`
		MakerRelayerFee  string `json:"maker_relayer_fee"`
		TakerRelayerFee  string `json:"taker_relayer_fee"`
		MakerProtocolFee string `json:"maker_protocol_fee"`
		TakerProtocolFee string `json:"taker_protocol_fee"`
		MakerReferrerFee string `json:"maker_referrer_fee"`
		FeeRecipient     struct {
			User struct {
				Username string `json:"username"`
			} `json:"user"`
			ProfileImgURL string `json:"profile_img_url"`
			Address       string `json:"address"`
			Config        string `json:"config"`
		} `json:"fee_recipient"`
		FeeMethod            int    `json:"fee_method"`
		Side                 int    `json:"side"`
		SaleKind             int    `json:"sale_kind"`
		Target               string `json:"target"`
		HowToCall            int    `json:"how_to_call"`
		Calldata             string `json:"calldata"`
		ReplacementPattern   string `json:"replacement_pattern"`
		StaticTarget         string `json:"static_target"`
		StaticExtradata      string `json:"static_extradata"`
		PaymentToken         string `json:"payment_token"`
		PaymentTokenContract struct {
			ID       int    `json:"id"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
			Name     string `json:"name"`
			Decimals int    `json:"decimals"`
			EthPrice string `json:"eth_price"`
			UsdPrice string `json:"usd_price"`
		} `json:"payment_token_contract"`
		BasePrice       string `json:"base_price"`
		Extra           string `json:"extra"`
		Quantity        string `json:"quantity"`
		Salt            string `json:"salt"`
		V               int    `json:"v"`
		R               string `json:"r"`
		S               string `json:"s"`
		ApprovedOnChain bool   `json:"approved_on_chain"`
		Cancelled       bool   `json:"cancelled"`
		Finalized       bool   `json:"finalized"`
		MarkedInvalid   bool   `json:"marked_invalid"`
		PrefixedHash    string `json:"prefixed_hash"`
	} `json:"orders"`
	Creator struct {
		User struct {
			Username string `json:"username"`
		} `json:"user"`
		ProfileImgURL string `json:"profile_img_url"`
		Address       string `json:"address"`
		Config        string `json:"config"`
	} `json:"creator"`
	Traits []struct {
		TraitType   string      `json:"trait_type"`
		Value       interface{} `json:"value"`
		DisplayType string      `json:"display_type"`
		MaxValue    interface{} `json:"max_value"`
		TraitCount  int         `json:"trait_count"`
		Order       string      `json:"order"`
	} `json:"traits"`
	TopOwnerships []struct {
		Owner struct {
			User struct {
				Username string `json:"username"`
			} `json:"user"`
			ProfileImgURL string `json:"profile_img_url"`
			Address       string `json:"address"`
			Config        string `json:"config"`
		} `json:"owner"`
		Quantity string `json:"quantity"`
	} `json:"top_ownerships"`
	TopBid                  string `json:"top_bid"`
	ListingDate             string `json:"listing_date"`
	IsPresale               bool   `json:"is_presale"`
	TransferFeePaymentToken string `json:"transfer_fee_payment_token"`
	TransferFee             string `json:"transfer_fee"`
}

type AutoCollection struct {
	PrimaryAssetContracts []struct {
		Address                     string `json:"address"`
		AssetContractType           string `json:"asset_contract_type"`
		CreatedDate                 string `json:"created_date"`
		Name                        string `json:"name"`
		NftVersion                  string `json:"nft_version"`
		OpenseaVersion              string `json:"opensea_version"`
		Owner                       int    `json:"owner"`
		SchemaName                  string `json:"schema_name"`
		Symbol                      string `json:"symbol"`
		TotalSupply                 string `json:"total_supply"`
		Description                 string `json:"description"`
		ExternalLink                string `json:"external_link"`
		ImageURL                    string `json:"image_url"`
		DefaultToFiat               bool   `json:"default_to_fiat"`
		DevBuyerFeeBasisPoints      int    `json:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints     int    `json:"dev_seller_fee_basis_points"`
		OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  int    `json:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints int    `json:"opensea_seller_fee_basis_points"`
		BuyerFeeBasisPoints         int    `json:"buyer_fee_basis_points"`
		SellerFeeBasisPoints        int    `json:"seller_fee_basis_points"`
		PayoutAddress               string `json:"payout_address"`
	} `json:"primary_asset_contracts"`
	Traits struct {
	} `json:"traits"`
	Stats struct {
		OneDayVolume          float64 `json:"one_day_volume"`
		OneDayChange          float64 `json:"one_day_change"`
		OneDaySales           float64 `json:"one_day_sales"`
		OneDayAveragePrice    float64 `json:"one_day_average_price"`
		SevenDayVolume        float64 `json:"seven_day_volume"`
		SevenDayChange        float64 `json:"seven_day_change"`
		SevenDaySales         float64 `json:"seven_day_sales"`
		SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
		ThirtyDayVolume       float64 `json:"thirty_day_volume"`
		ThirtyDayChange       float64 `json:"thirty_day_change"`
		ThirtyDaySales        float64 `json:"thirty_day_sales"`
		ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
		TotalVolume           float64 `json:"total_volume"`
		TotalSales            float64 `json:"total_sales"`
		TotalSupply           float64 `json:"total_supply"`
		Count                 float64 `json:"count"`
		NumOwners             int     `json:"num_owners"`
		AveragePrice          float64 `json:"average_price"`
		NumReports            int     `json:"num_reports"`
		MarketCap             float64 `json:"market_cap"`
		FloorPrice            float64 `json:"floor_price"`
	} `json:"stats"`
	BannerImageURL          string `json:"banner_image_url"`
	ChatURL                 string `json:"chat_url"`
	CreatedDate             string `json:"created_date"`
	DefaultToFiat           bool   `json:"default_to_fiat"`
	Description             string `json:"description"`
	DevBuyerFeeBasisPoints  string `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints string `json:"dev_seller_fee_basis_points"`
	DiscordURL              string `json:"discord_url"`
	DisplayData             struct {
		CardDisplayStyle string `json:"card_display_style"`
	} `json:"display_data"`
	ExternalURL                 string  `json:"external_url"`
	Featured                    bool    `json:"featured"`
	FeaturedImageURL            string  `json:"featured_image_url"`
	Hidden                      bool    `json:"hidden"`
	SafelistRequestStatus       string  `json:"safelist_request_status"`
	ImageURL                    string  `json:"image_url"`
	IsSubjectToWhitelist        bool    `json:"is_subject_to_whitelist"`
	LargeImageURL               string  `json:"large_image_url"`
	MediumUsername              string  `json:"medium_username"`
	Name                        string  `json:"name"`
	OnlyProxiedTransfers        bool    `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string  `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string  `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               string  `json:"payout_address"`
	RequireEmail                bool    `json:"require_email"`
	ShortDescription            string  `json:"short_description"`
	Slug                        string  `json:"slug"`
	TelegramURL                 string  `json:"telegram_url"`
	TwitterUsername             string  `json:"twitter_username"`
	InstagramUsername           string  `json:"instagram_username"`
	WikiURL                     string  `json:"wiki_url"`
	OwnedAssetCount             big.Int `json:"owned_asset_count"`
}

type AutoEvent struct {
	ApprovedAccount interface{} `json:"approved_account"`
	Asset           struct {
		ID                   int         `json:"id"`
		TokenID              string      `json:"token_id"`
		NumSales             int         `json:"num_sales"`
		BackgroundColor      interface{} `json:"background_color"`
		ImageURL             string      `json:"image_url"`
		ImagePreviewURL      string      `json:"image_preview_url"`
		ImageThumbnailURL    string      `json:"image_thumbnail_url"`
		ImageOriginalURL     string      `json:"image_original_url"`
		AnimationURL         interface{} `json:"animation_url"`
		AnimationOriginalURL interface{} `json:"animation_original_url"`
		Name                 string      `json:"name"`
		Description          string      `json:"description"`
		ExternalLink         interface{} `json:"external_link"`
		AssetContract        struct {
			Address                     string      `json:"address"`
			AssetContractType           string      `json:"asset_contract_type"`
			CreatedDate                 string      `json:"created_date"`
			Name                        string      `json:"name"`
			NftVersion                  string      `json:"nft_version"`
			OpenseaVersion              interface{} `json:"opensea_version"`
			Owner                       int         `json:"owner"`
			SchemaName                  string      `json:"schema_name"`
			Symbol                      string      `json:"symbol"`
			TotalSupply                 string      `json:"total_supply"`
			Description                 string      `json:"description"`
			ExternalLink                interface{} `json:"external_link"`
			ImageURL                    string      `json:"image_url"`
			DefaultToFiat               bool        `json:"default_to_fiat"`
			DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
			DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
			OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
			OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
			OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
			BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
			SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
			PayoutAddress               string      `json:"payout_address"`
		} `json:"asset_contract"`
		Permalink  string `json:"permalink"`
		Collection struct {
			BannerImageURL          string      `json:"banner_image_url"`
			ChatURL                 interface{} `json:"chat_url"`
			CreatedDate             string      `json:"created_date"`
			DefaultToFiat           bool        `json:"default_to_fiat"`
			Description             string      `json:"description"`
			DevBuyerFeeBasisPoints  string      `json:"dev_buyer_fee_basis_points"`
			DevSellerFeeBasisPoints string      `json:"dev_seller_fee_basis_points"`
			DiscordURL              interface{} `json:"discord_url"`
			DisplayData             struct {
				CardDisplayStyle string `json:"card_display_style"`
			} `json:"display_data"`
			ExternalURL                 interface{} `json:"external_url"`
			Featured                    bool        `json:"featured"`
			FeaturedImageURL            string      `json:"featured_image_url"`
			Hidden                      bool        `json:"hidden"`
			SafelistRequestStatus       string      `json:"safelist_request_status"`
			ImageURL                    string      `json:"image_url"`
			IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
			LargeImageURL               string      `json:"large_image_url"`
			MediumUsername              interface{} `json:"medium_username"`
			Name                        string      `json:"name"`
			OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
			OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points"`
			OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points"`
			PayoutAddress               string      `json:"payout_address"`
			RequireEmail                bool        `json:"require_email"`
			ShortDescription            interface{} `json:"short_description"`
			Slug                        string      `json:"slug"`
			TelegramURL                 interface{} `json:"telegram_url"`
			TwitterUsername             interface{} `json:"twitter_username"`
			InstagramUsername           interface{} `json:"instagram_username"`
			WikiURL                     interface{} `json:"wiki_url"`
		} `json:"collection"`
		Decimals      int    `json:"decimals"`
		TokenMetadata string `json:"token_metadata"`
		Owner         struct {
			User struct {
				Username interface{} `json:"username"`
			} `json:"user"`
			ProfileImgURL string `json:"profile_img_url"`
			Address       string `json:"address"`
			Config        string `json:"config"`
		} `json:"owner"`
	} `json:"asset"`
	AssetBundle             interface{} `json:"asset_bundle"`
	AuctionType             interface{} `json:"auction_type"`
	BidAmount               string      `json:"bid_amount"`
	CollectionSlug          string      `json:"collection_slug"`
	ContractAddress         string      `json:"contract_address"`
	CreatedDate             string      `json:"created_date"`
	CustomEventName         interface{} `json:"custom_event_name"`
	DevFeePaymentEvent      interface{} `json:"dev_fee_payment_event"`
	DevSellerFeeBasisPoints int         `json:"dev_seller_fee_basis_points"`
	Duration                interface{} `json:"duration"`
	EndingPrice             interface{} `json:"ending_price"`
	EventType               string      `json:"event_type"`
	FromAccount             interface{} `json:"from_account"`
	ID                      int         `json:"id"`
	IsPrivate               bool        `json:"is_private"`
	OwnerAccount            interface{} `json:"owner_account"`
	PaymentToken            struct {
		ID       int         `json:"id"`
		Symbol   string      `json:"symbol"`
		Address  string      `json:"address"`
		ImageURL string      `json:"image_url"`
		Name     interface{} `json:"name"`
		Decimals int         `json:"decimals"`
		EthPrice string      `json:"eth_price"`
		UsdPrice interface{} `json:"usd_price"`
	} `json:"payment_token"`
	Quantity string `json:"quantity"`
	Seller   struct {
		User struct {
			Username string `json:"username"`
		} `json:"user"`
		ProfileImgURL string `json:"profile_img_url"`
		Address       string `json:"address"`
		Config        string `json:"config"`
	} `json:"seller"`
	StartingPrice interface{} `json:"starting_price"`
	ToAccount     interface{} `json:"to_account"`
	TotalPrice    string      `json:"total_price"`
	Transaction   struct {
		BlockHash   string `json:"block_hash" bson:"block_hash"`
		BlockNumber string `json:"block_number" bson:"block_number"`
		FromAccount struct {
			User struct {
				Username interface{} `json:"username" bson:"username"`
			} `json:"user" bson:"user"`
			ProfileImgURL string `json:"profile_img_url" bson:"profile_img_url"`
			Address       string `json:"address" bson:"address"` // 支付人
			Config        string `json:"config" bson:"config"`
		} `json:"from_account" bson:"from_account"`
		ID        int    `json:"id" bson:"id"`
		Timestamp string `json:"timestamp" bson:"timestamp"`
		ToAccount struct {
			User          interface{} `json:"user" bson:"user"`
			ProfileImgURL string      `json:"profile_img_url" bson:"profile_img_url"`
			Address       string      `json:"address" bson:"address"` // 支付对象 合约地址
			Config        string      `json:"config" bson:"config"`
		} `json:"to_account" bson:"to_account"`
		TransactionHash  string `json:"transaction_hash" bson:"transaction_hash"`
		TransactionIndex string `json:"transaction_index" bson:"transaction_index"`
	} `json:"transaction" bson:"transaction"` // 支付的eth链上交易记录
	WinnerAccount struct {
		User struct {
			Username string `json:"username"`
		} `json:"user"`
		ProfileImgURL string `json:"profile_img_url"`
		Address       string `json:"address"`
		Config        string `json:"config"`
	} `json:"winner_account"`
	ListingTime string `json:"listing_time"`
}

type Params struct {
	Page     int64   `form:"page" binding:"numeric,min=1"`
	PageSize int64   `form:"pageSize" binding:"numeric,min=1"`
	Status   int64   `form:"status" binding:"oneof=0 1"`
	SortBy   int64   `form:"sortBy" binding:"oneof=0 1 2 3 4 5 6"`
	MinPrice float64 `form:"minPrice" binding:"numeric,min=0"`
	MaxPrice float64 `form:"maxPrice" binding:"numeric,min=0"`
	Field    string  `form:"field"`
}
