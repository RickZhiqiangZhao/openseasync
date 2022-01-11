package models

import "math/big"

type Asset struct {
	//ID                int64  `json:"id" bson:"id"`                                   // 主键
	UserAddress       string `json:"user_address" bson:"user_address"`               // 用户地址
	Title             string `json:"title" bson:"title"`                             // NFT作品标题
	ImageURL          string `json:"image_url" bson:"image_url"`                     // NFT作品图片
	ImagePreviewURL   string `json:"image_preview_url" bson:"image_preview_url"`     // NFT作品原图
	ImageThumbnailURL string `json:"image_thumbnail_url" bson:"image_thumbnail_url"` // NFT作品缩略图
	Description       string `json:"description" bson:"description"`                 // NFT作品描述
	ContractAddress   string `json:"contract_address" bson:"contract_address"`       // 合约地址
	TokenId           string `json:"token_id" bson:"token_id"`                       // NFT token id
	NumSales          int    `json:"num_sales" bson:"num_sales"`                     // NFT售卖次数
	Owner             string `json:"owner" bson:"owner"`                             // NFT拥有者
	OwnerImgURL       string `json:"owner_img_url" bson:"owner_img_url"`             // 拥有者头像
	Creator           string `json:"creator" bson:"creator"`                         // NFT创造者
	CreatorImgURL     string `json:"creator_img_url" bson:"creator_img_url"`         // 创造者头像
	TokenMetadata     string `json:"token_metadata" bson:"token_metadata"`           // NFT元数据

	Slug string `json:"slug" bson:"slug"` // 集合唯一标识符号

	//Contract            Contract             `json:"contract" bson:"contract"`
	//Collection          Collection           `json:"collection" bson:"collection"`
	AssetsTopOwnerships []AssetsTopOwnership `json:"assets_top_ownerships" bson:"assets_top_ownerships"`
	Traits              []Trait              `json:"traits" bson:"traits"`

	SellOrders SellOrder `json:"sell_orders" bson:"sell_orders"`

	IsDelete    int8 `json:"is_delete" bson:"is_delete"`       // 是否删除 1删除 0未删除 默认为0
	RefreshTime int  `json:"refresh_time" bson:"refresh_time"` // 刷新时间
}

type Collection struct {
	//ID              int64   `json:"id" bson:"id"`                               // 主键
	Slug            string  `json:"slug" bson:"slug"`                           // 集合唯一标识符
	UserAddress     string  `json:"user_address" bson:"user_address"`           // 集合拥有者
	Name            string  `json:"name" bson:"name"`                           // 集合名称
	BannerImageURL  string  `json:"banner_image_url" bson:"banner_image_url"`   // 集合背景图
	Description     string  `json:"description" bson:"description"`             // 集合描述
	ImageURL        string  `json:"image_url" bson:"image_url"`                 // 集合头像
	LargeImageURL   string  `json:"large_image_url" bson:"large_image_url"`     // 头像大图
	IsDelete        int8    `json:"is_delete" bson:"is_delete"`                 // 是否删除 1删除 0未删除 默认为0
	CreateDate      string  `json:"create_date" bson:"create_date"`             // 集合创建时间
	RefreshTime     int     `json:"refresh_time" bson:"refresh_time"`           // 刷新时间
	NumOwners       int     `json:"num_owners" bson:"num_owners"`               // 集合中输入自己的NFT个数
	TotalSupply     int     `json:"total_supply" bson:"total_supply"`           // 集合中NFT总数
	TotalVolume     float64 `json:"total_volume" bson:"total_volume"`           // 交易量
	FloorPrice      float64 `json:"floor_price"`                                // 最低价格
	OwnedAssetCount string  `json:"owned_asset_count" bson:"owned_asset_count"` // 所有NFT中属于自己的NFT个数 此地段可能是个big int, 所以采用string存储
}

type Contract struct {
	//ID           int64  `json:"id" bson:"id"`                       // 主键
	Address      string `json:"address" bson:"address"`             // 合约地址
	ContractType string `json:"contract_type" bson:"contract_type"` // 合约类型 semi-fungible可替代 non-fungible 不可替代
	ContractName string `json:"contract_name" bson:"contract_name"` // 合约名字
	Symbol       string `json:"symbol" bson:"symbol"`               // 符号
	SchemaName   string `json:"schema_name" bson:"schema_name"`     // 合约类型
	TotalSupply  string `json:"total_supply" bson:"total_supply"`   // 总供应量
	Description  string `json:"description" bson:"description"`     // 合约描述
}

type SellOrder struct {
	CreateDate   string `json:"create_date" bson:"create_date"`     // 创建时间
	ClosingDate  string `json:"closing_date" bson:"closing_date"`   // 结束时间
	CurrentPrice string `json:"current_price" bson:"current_price"` // 当前价格

	PayTokenContract PayTokenContract `json:"pay_token_contract" bson:"pay_token_contract"` // 支付方式
}

type PayTokenContract struct {
	Symbol   string `json:"symbol" bson:"symbol"`
	ImageURL string `json:"image_url" bson:"image_url"`
	EthPrice string `json:"eth_price" bson:"eth_price"`
	UsdPrice string `json:"usd_price" bson:"usd_price"`
}

type Trait struct {
	//ID              int64  `json:"id" bson:"id"`                     // 主键
	UserAddress     string `json:"user_address" bson:"user_address"` // 用户地址
	ContractAddress string `json:"_" bson:"contract_address"`        // 合约地址
	TokenId         string `json:"_" bson:"token_id"`                // token id
	TraitType       string `json:"trait_type" bson:"trait_type"`     // 特征类型
	Value           string `json:"value" bson:"value"`               // 特征值
	DisplayType     string `json:"display_type" bson:"display_type"`
	MaxValue        int    `json:"max_value" bson:"max_value"`
	TraitCount      int    `json:"trait_count" bson:"trait_count"` // 数量
	OrderBy         string `json:"order_by" bson:"order_by"`
	IsDelete        int8   `json:"is_delete" bson:"is_delete"`       // 是否删除 1删除 0未删除 默认为0
	RefreshTime     int    `json:"refresh_time" bson:"refresh_time"` // 刷新时间
}

type AssetsTopOwnership struct {
	//ID              int64  `json:"id" bson:"id"`                           // 主键
	UserAddress     string `json:"user_address" bson:"user_address"`       // 用户地址
	ContractAddress string `json:"_" bson:"contract_address"`              // 合约地址
	TokenId         string `json:"_" bson:"token_id"`                      // token id
	Owner           string `json:"owner" bson:"owner"`                     // 所有者地址
	ProfileImgURL   string `json:"profile_img_url" bson:"profile_img_url"` // 所有者头像
	Quantity        string `json:"quantity" bson:"quantity"`               // 数量
	IsDelete        int8   `json:"is_delete" bson:"is_delete"`             // 是否删除 1删除 0未删除 默认为0
	RefreshTime     int    `json:"refresh_time" bson:"refresh_time"`       // 刷新时间
}

type ItemActivity struct {
	Id                  int    `json:"id"`
	UserAddress         string `json:"user_address" bson:"user_address"`
	ContractAddress     string `json:"contract_address" bson:"contract_address"`
	TokenId             string `json:"token_id" bson:"token_id"`
	BidAmount           string `json:"bid_amount" bson:"bid_amount"` // 投标金额
	CreateDate          string `json:"create_date" bson:"create_date"`
	TotalPrice          string `json:"total_price" bson:"total_price"`                       // 成交价格
	Seller              string `json:"seller" bson:"seller"`                                 // 售卖者地址
	SellerProfileImgURL string `json:"seller_profile_img_url" bson:"seller_profile_img_url"` // 售卖者头像
	Winner              string `json:"winner" bson:"winner"`                                 // 购买者地址
	WinnerProfileImgURL string `json:"winner_profile_img_url" bson:"winner_profile_img_url"` // 购买者头像
	IsDelete            int8   `json:"is_delete" bson:"is_delete"`                           // 是否删除 1删除 0未删除 默认为0
	EventType           string `json:"event_type" bson:"event_type"`                         // 事件类型

	Transaction struct {
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
	Creator struct {
		User struct {
			Username string `json:"username"`
		} `json:"user"`
		ProfileImgURL string `json:"profile_img_url"`
		Address       string `json:"address"`
		Config        string `json:"config"`
	} `json:"creator"`
	Traits []struct {
		TraitType   string `json:"trait_type"`
		Value       string `json:"value"`
		DisplayType string `json:"display_type"`
		MaxValue    int    `json:"max_value"`
		TraitCount  int    `json:"trait_count"`
		Order       string `json:"order"`
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
		UsdPrice string      `json:"usd_price"`
	} `json:"payment_token"`
	Quantity string `json:"quantity"`
	Seller   struct {
		User          interface{} `json:"user"`
		ProfileImgURL string      `json:"profile_img_url"`
		Address       string      `json:"address"`
		Config        string      `json:"config"`
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
			Username interface{} `json:"username"`
		} `json:"user"`
		ProfileImgURL string `json:"profile_img_url"`
		Address       string `json:"address"`
		Config        string `json:"config"`
	} `json:"winner_account"`
	ListingTime string `json:"listing_time"`
}
