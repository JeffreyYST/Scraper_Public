package models

import (
	"Scraper/common"
	"context"
	"time"
)

type Record struct {

	//tableName struct{} `json:"-",pg:"user_account"`

	IDInv                  int    `json:"idInv"`
	StrDesc                string `json:"strDesc"`
	CodeNew                string `json:"codeNew"`
	CodeComplete           string `json:"codeComplete"`
	StrInvImgURL           string `json:"strInvImgUrl"`
	IDInvImg               int    `json:"idInvImg"`
	TypeInvImg             string `json:"typeInvImg"`
	N4Qty                  int    `json:"n4Qty"`
	IDColorDefault         int    `json:"idColorDefault"`
	TypeImgDefault         string `json:"typeImgDefault"`
	HasExtendedDescription int    `json:"hasExtendedDescription"`
	InstantCheckout        bool   `json:"instantCheckout"`
	MDisplaySalePrice      string `json:"mDisplaySalePrice"`
	MInvSalePrice          string `json:"mInvSalePrice"`
	NSalePct               int    `json:"nSalePct"`
	NTier1Qty              int    `json:"nTier1Qty"`
	NTier2Qty              int    `json:"nTier2Qty"`
	NTier3Qty              int    `json:"nTier3Qty"`
	NTire1DisplayPrice     string `json:"nTier1DisplayPrice"`
	NTire2DisplayPrice     string `json:"nTier2DisplayPrice"`
	NTire3DisplayPrice     string `json:"nTier3DisplayPrice"`
	NTire1InvPrice         string `json:"nTier1InvPrice"`
	NTire2InvPrice         string `json:"nTier2InvPrice"`
	NTire3InvPrice         string `json:"nTier3InvPrice"`
	IDColor                int    `json:"idColor"`
	StrCategory            string `json:"strCategory"`
	StrStorename           string `json:"strStorename"`
	IDCurrencyStore        int    `json:"idCurrencyStore"`
	MMinBuy                string `json:"mMinBuy"`
	StrSellerUsername      string `json:"strSellerUsername"`
	N4SellerFeedbackScore  int    `json:"n4SellerFeedbackScore"`
	StrSellerCountryName   string `json:"strSellerCountryName"`
	StrSellerCountryCode   string `json:"strSellerCountryCode"`
	StrColor               string `json:"strColor"`
	//LOL string `json:"-"`
}

type Listings struct {
	Total_Count int `json:"total_count"`
	ColorId     int `json:"idColor"`
	RPP         int `json:"rpp"`
	PageIndex   int `json:"pi"`

	Records []Record `json:"list"`

	ReturnCode     int    `json:"returnCode"`
	ReturnMessage  string `json:"returnMessage"`
	ErrorTicket    int    `json:"errorTicket"`
	ProcessingTime int    `json:"processingTime"`
}

func InsertRecord(itemId string, r Record) error {
	dbpool := common.GetDBPool()

	_, err := dbpool.Exec(context.Background(),
		"INSERT INTO listings(item_id, inv_id, code_new, code_complete, quantity, color_id, display_sale_price, category, store_name, store_currency_id, min_buy, seller_username, seller_feedback_score, seller_country_name, seller_country_code, description, time_stamp) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)", 
		itemId,
		r.IDInv,
		r.CodeNew,
		r.CodeComplete,
		r.N4Qty,
		r.IDColor,
		r.MDisplaySalePrice,
		r.StrCategory,
		r.StrStorename,
		r.IDCurrencyStore,
		r.MMinBuy,
		r.StrSellerUsername,
		r.N4SellerFeedbackScore,
		r.StrSellerCountryName,
		r.StrSellerCountryCode,
		r.StrDesc,
		time.Now())

	return err 
}

