package models

import "math/big"

type OwnerAsset struct {
	Assets []AutoAsset `json:"assets"`
}

type OwnerCollection struct {
	Collections []AutoCollection `json:"collections"`
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
		FloorPrice            int     `json:"floor_price"`
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
