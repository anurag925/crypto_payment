package impl

import (
	"context"
	"fmt"
	"github.com/anurag925/crypto_payment/utils/logger"
	"net/http"
	"net/url"
	"strconv"

	"github.com/parnurzeal/gorequest"
)

const exchangeRateUrl = "https://api.coinbase.com/v2/exchange-rates?currency="

type coinbaseLibImpl struct {
}

func NewCoinbaseLib() *coinbaseLibImpl {
	return &coinbaseLibImpl{}
}

type USDTEuroExchangeRate struct {
	Data struct {
		Currency string `json:"currency"`
		Rates    struct {
			USDT string `json:"USDT" redis:"USDT"`
		} `json:"rates" redis:"Rates"`
	} `json:"data"`
}

type AllExchangeRates struct {
	Data struct {
		Currency string `json:"currency"`
		Rates    Rates  `json:"rates"`
	} `json:"data"`
}

type Rates map[string]string

// type Rates struct {
// 	AED     string `json:"AED" redis:"AED"`
// 	AFN     string `json:"AFN" redis:"AFN"`
// 	ALL     string `json:"ALL" redis:"ALL"`
// 	AMD     string `json:"AMD" redis:"AMD"`
// 	ANG     string `json:"ANG" redis:"ANG"`
// 	AOA     string `json:"AOA" redis:"AOA"`
// 	ARS     string `json:"ARS" redis:"ARS"`
// 	AWG     string `json:"AWG" redis:"AWG"`
// 	AZN     string `json:"AZN" redis:"AZN"`
// 	BAM     string `json:"BAM" redis:"BAM"`
// 	BBD     string `json:"BBD" redis:"BBD"`
// 	BDT     string `json:"BDT" redis:"BDT"`
// 	BGN     string `json:"BGN" redis:"BGN"`
// 	BHD     string `json:"BHD" redis:"BHD"`
// 	BIF     string `json:"BIF" redis:"BIF"`
// 	BMD     string `json:"BMD" redis:"BMD"`
// 	BND     string `json:"BND" redis:"BND"`
// 	BOB     string `json:"BOB" redis:"BOB"`
// 	BRL     string `json:"BRL" redis:"BRL"`
// 	BSD     string `json:"BSD" redis:"BSD"`
// 	BTN     string `json:"BTN" redis:"BTN"`
// 	BWP     string `json:"BWP" redis:"BWP"`
// 	BYN     string `json:"BYN" redis:"BYN"`
// 	BYR     string `json:"BYR" redis:"BYR"`
// 	BZD     string `json:"BZD" redis:"BZD"`
// 	CAD     string `json:"CAD" redis:"CAD"`
// 	CDF     string `json:"CDF" redis:"CDF"`
// 	CHF     string `json:"CHF" redis:"CHF"`
// 	CLF     string `json:"CLF" redis:"CLF"`
// 	CLP     string `json:"CLP" redis:"CLP"`
// 	CNY     string `json:"CNY" redis:"CNY"`
// 	COP     string `json:"COP" redis:"COP"`
// 	CRC     string `json:"CRC" redis:"CRC"`
// 	CUC     string `json:"CUC" redis:"CUC"`
// 	CVE     string `json:"CVE" redis:"CVE"`
// 	CZK     string `json:"CZK" redis:"CZK"`
// 	DJF     string `json:"DJF" redis:"DJF"`
// 	DKK     string `json:"DKK" redis:"DKK"`
// 	DOP     string `json:"DOP" redis:"DOP"`
// 	DZD     string `json:"DZD" redis:"DZD"`
// 	EGP     string `json:"EGP" redis:"EGP"`
// 	ETB     string `json:"ETB" redis:"ETB"`
// 	EUR     string `json:"EUR" redis:"EUR"`
// 	FJD     string `json:"FJD" redis:"FJD"`
// 	FKP     string `json:"FKP" redis:"FKP"`
// 	GBP     string `json:"GBP" redis:"GBP"`
// 	GEL     string `json:"GEL" redis:"GEL"`
// 	GHS     string `json:"GHS" redis:"GHS"`
// 	GIP     string `json:"GIP" redis:"GIP"`
// 	GMD     string `json:"GMD" redis:"GMD"`
// 	GNF     string `json:"GNF" redis:"GNF"`
// 	GTQ     string `json:"GTQ" redis:"GTQ"`
// 	GYD     string `json:"GYD" redis:"GYD"`
// 	HKD     string `json:"HKD" redis:"HKD"`
// 	HNL     string `json:"HNL" redis:"HNL"`
// 	HRK     string `json:"HRK" redis:"HRK"`
// 	HTG     string `json:"HTG" redis:"HTG"`
// 	HUF     string `json:"HUF" redis:"HUF"`
// 	IDR     string `json:"IDR" redis:"IDR"`
// 	ILS     string `json:"ILS" redis:"ILS"`
// 	INR     string `json:"INR" redis:"INR"`
// 	IQD     string `json:"IQD" redis:"IQD"`
// 	ISK     string `json:"ISK" redis:"ISK"`
// 	JMD     string `json:"JMD" redis:"JMD"`
// 	JOD     string `json:"JOD" redis:"JOD"`
// 	JPY     string `json:"JPY" redis:"JPY"`
// 	KES     string `json:"KES" redis:"KES"`
// 	KGS     string `json:"KGS" redis:"KGS"`
// 	KHR     string `json:"KHR" redis:"KHR"`
// 	KMF     string `json:"KMF" redis:"KMF"`
// 	KRW     string `json:"KRW" redis:"KRW"`
// 	KWD     string `json:"KWD" redis:"KWD"`
// 	KYD     string `json:"KYD" redis:"KYD"`
// 	KZT     string `json:"KZT" redis:"KZT"`
// 	LAK     string `json:"LAK" redis:"LAK"`
// 	LBP     string `json:"LBP" redis:"LBP"`
// 	LKR     string `json:"LKR" redis:"LKR"`
// 	LRD     string `json:"LRD" redis:"LRD"`
// 	LSL     string `json:"LSL" redis:"LSL"`
// 	LYD     string `json:"LYD" redis:"LYD"`
// 	MAD     string `json:"MAD" redis:"MAD"`
// 	MDL     string `json:"MDL" redis:"MDL"`
// 	MGA     string `json:"MGA" redis:"MGA"`
// 	MKD     string `json:"MKD" redis:"MKD"`
// 	MMK     string `json:"MMK" redis:"MMK"`
// 	MNT     string `json:"MNT" redis:"MNT"`
// 	MOP     string `json:"MOP" redis:"MOP"`
// 	MRO     string `json:"MRO" redis:"MRO"`
// 	MUR     string `json:"MUR" redis:"MUR"`
// 	MVR     string `json:"MVR" redis:"MVR"`
// 	MWK     string `json:"MWK" redis:"MWK"`
// 	MXN     string `json:"MXN" redis:"MXN"`
// 	MYR     string `json:"MYR" redis:"MYR"`
// 	MZN     string `json:"MZN" redis:"MZN"`
// 	NAD     string `json:"NAD" redis:"NAD"`
// 	NGN     string `json:"NGN" redis:"NGN"`
// 	NIO     string `json:"NIO" redis:"NIO"`
// 	NOK     string `json:"NOK" redis:"NOK"`
// 	NPR     string `json:"NPR" redis:"NPR"`
// 	NZD     string `json:"NZD" redis:"NZD"`
// 	OMR     string `json:"OMR" redis:"OMR"`
// 	PAB     string `json:"PAB" redis:"PAB"`
// 	PEN     string `json:"PEN" redis:"PEN"`
// 	PGK     string `json:"PGK" redis:"PGK"`
// 	PHP     string `json:"PHP" redis:"PHP"`
// 	PKR     string `json:"PKR" redis:"PKR"`
// 	PLN     string `json:"PLN" redis:"PLN"`
// 	PYG     string `json:"PYG" redis:"PYG"`
// 	QAR     string `json:"QAR" redis:"QAR"`
// 	RON     string `json:"RON" redis:"RON"`
// 	RSD     string `json:"RSD" redis:"RSD"`
// 	RUB     string `json:"RUB" redis:"RUB"`
// 	RWF     string `json:"RWF" redis:"RWF"`
// 	SAR     string `json:"SAR" redis:"SAR"`
// 	SBD     string `json:"SBD" redis:"SBD"`
// 	SCR     string `json:"SCR" redis:"SCR"`
// 	SDG     string `json:"SDG" redis:"SDG"`
// 	SEK     string `json:"SEK" redis:"SEK"`
// 	SHP     string `json:"SHP" redis:"SHP"`
// 	SKK     string `json:"SKK" redis:"SKK"`
// 	SLL     string `json:"SLL" redis:"SLL"`
// 	SOS     string `json:"SOS" redis:"SOS"`
// 	SRD     string `json:"SRD" redis:"SRD"`
// 	STD     string `json:"STD" redis:"STD"`
// 	SVC     string `json:"SVC" redis:"SVC"`
// 	SZL     string `json:"SZL" redis:"SZL"`
// 	THB     string `json:"THB" redis:"THB"`
// 	TJS     string `json:"TJS" redis:"TJS"`
// 	TMT     string `json:"TMT" redis:"TMT"`
// 	TND     string `json:"TND" redis:"TND"`
// 	TOP     string `json:"TOP" redis:"TOP"`
// 	TRY     string `json:"TRY" redis:"TRY"`
// 	TTD     string `json:"TTD" redis:"TTD"`
// 	TWD     string `json:"TWD" redis:"TWD"`
// 	TZS     string `json:"TZS" redis:"TZS"`
// 	UAH     string `json:"UAH" redis:"UAH"`
// 	UGX     string `json:"UGX" redis:"UGX"`
// 	UYU     string `json:"UYU" redis:"UYU"`
// 	UZS     string `json:"UZS" redis:"UZS"`
// 	VES     string `json:"VES" redis:"VES"`
// 	VND     string `json:"VND" redis:"VND"`
// 	VUV     string `json:"VUV" redis:"VUV"`
// 	WST     string `json:"WST" redis:"WST"`
// 	XAF     string `json:"XAF" redis:"XAF"`
// 	XAG     string `json:"XAG" redis:"XAG"`
// 	XAU     string `json:"XAU" redis:"XAU"`
// 	XCD     string `json:"XCD" redis:"XCD"`
// 	XOF     string `json:"XOF" redis:"XOF"`
// 	XPD     string `json:"XPD" redis:"XPD"`
// 	XPF     string `json:"XPF" redis:"XPF"`
// 	XPT     string `json:"XPT" redis:"XPT"`
// 	YER     string `json:"YER" redis:"YER"`
// 	ZAR     string `json:"ZAR" redis:"ZAR"`
// 	ZMK     string `json:"ZMK" redis:"ZMK"`
// 	ZMW     string `json:"ZMW" redis:"ZMW"`
// 	JEP     string `json:"JEP" redis:"JEP"`
// 	GGP     string `json:"GGP" redis:"GGP"`
// 	IMP     string `json:"IMP" redis:"IMP"`
// 	CNH     string `json:"CNH" redis:"CNH"`
// 	EEK     string `json:"EEK" redis:"EEK"`
// 	LTL     string `json:"LTL" redis:"LTL"`
// 	LVL     string `json:"LVL" redis:"LVL"`
// 	TMM     string `json:"TMM" redis:"TMM"`
// 	ZWD     string `json:"ZWD" redis:"ZWD"`
// 	VEF     string `json:"VEF" redis:"VEF"`
// 	SGD     string `json:"SGD" redis:"SGD"`
// 	AUD     string `json:"AUD" redis:"AUD"`
// 	USD     string `json:"USD" redis:"USD"`
// 	BTC     string `json:"BTC" redis:"BTC"`
// 	BCH     string `json:"BCH" redis:"BCH"`
// 	BSV     string `json:"BSV" redis:"BSV"`
// 	ETH     string `json:"ETH" redis:"ETH"`
// 	ETH2    string `json:"ETH2" redis:"ETH2"`
// 	ETC     string `json:"ETC" redis:"ETC"`
// 	LTC     string `json:"LTC" redis:"LTC"`
// 	ZRX     string `json:"ZRX" redis:"ZRX"`
// 	USDC    string `json:"USDC" redis:"USDC"`
// 	BAT     string `json:"BAT" redis:"BAT"`
// 	LOOM    string `json:"LOOM" redis:"LOOM"`
// 	MANA    string `json:"MANA" redis:"MANA"`
// 	KNC     string `json:"KNC" redis:"KNC"`
// 	LINK    string `json:"LINK" redis:"LINK"`
// 	DNT     string `json:"DNT" redis:"DNT"`
// 	MKR     string `json:"MKR" redis:"MKR"`
// 	CVC     string `json:"CVC" redis:"CVC"`
// 	OMG     string `json:"OMG" redis:"OMG"`
// 	GNT     string `json:"GNT" redis:"GNT"`
// 	DAI     string `json:"DAI" redis:"DAI"`
// 	SNT     string `json:"SNT" redis:"SNT"`
// 	ZEC     string `json:"ZEC" redis:"ZEC"`
// 	XRP     string `json:"XRP" redis:"XRP"`
// 	REP     string `json:"REP" redis:"REP"`
// 	XLM     string `json:"XLM" redis:"XLM"`
// 	EOS     string `json:"EOS" redis:"EOS"`
// 	XTZ     string `json:"XTZ" redis:"XTZ"`
// 	ALGO    string `json:"ALGO" redis:"ALGO"`
// 	DASH    string `json:"DASH" redis:"DASH"`
// 	ATOM    string `json:"ATOM" redis:"ATOM"`
// 	OXT     string `json:"OXT" redis:"OXT"`
// 	COMP    string `json:"COMP" redis:"COMP"`
// 	ENJ     string `json:"ENJ" redis:"ENJ"`
// 	REPV2   string `json:"REPV2" redis:"REPV2"`
// 	BAND    string `json:"BAND" redis:"BAND"`
// 	NMR     string `json:"NMR" redis:"NMR"`
// 	CGLD    string `json:"CGLD" redis:"CGLD"`
// 	UMA     string `json:"UMA" redis:"UMA"`
// 	LRC     string `json:"LRC" redis:"LRC"`
// 	YFI     string `json:"YFI" redis:"YFI"`
// 	UNI     string `json:"UNI" redis:"UNI"`
// 	BAL     string `json:"BAL" redis:"BAL"`
// 	REN     string `json:"REN" redis:"REN"`
// 	WBTC    string `json:"WBTC" redis:"WBTC"`
// 	NU      string `json:"NU" redis:"NU"`
// 	YFII    string `json:"YFII" redis:"YFII"`
// 	FIL     string `json:"FIL" redis:"FIL"`
// 	AAVE    string `json:"AAVE" redis:"AAVE"`
// 	BNT     string `json:"BNT" redis:"BNT"`
// 	GRT     string `json:"GRT" redis:"GRT"`
// 	SNX     string `json:"SNX" redis:"SNX"`
// 	STORJ   string `json:"STORJ" redis:"STORJ"`
// 	SUSHI   string `json:"SUSHI" redis:"SUSHI"`
// 	MATIC   string `json:"MATIC" redis:"MATIC"`
// 	SKL     string `json:"SKL" redis:"SKL"`
// 	ADA     string `json:"ADA" redis:"ADA"`
// 	ANKR    string `json:"ANKR" redis:"ANKR"`
// 	CRV     string `json:"CRV" redis:"CRV"`
// 	ICP     string `json:"ICP" redis:"ICP"`
// 	NKN     string `json:"NKN" redis:"NKN"`
// 	OGN     string `json:"OGN" redis:"OGN"`
// 	OneINCH string `json:"1INCH" redis:"OneINCH"`
// 	USDT    string `json:"USDT" redis:"USDT"`
// 	FORTH   string `json:"FORTH" redis:"FORTH"`
// 	CTSI    string `json:"CTSI" redis:"CTSI"`
// 	TRB     string `json:"TRB" redis:"TRB"`
// 	POLY    string `json:"POLY" redis:"POLY"`
// 	MIR     string `json:"MIR" redis:"MIR"`
// 	RLC     string `json:"RLC" redis:"RLC"`
// 	DOT     string `json:"DOT" redis:"DOT"`
// 	SOL     string `json:"SOL" redis:"SOL"`
// 	DOGE    string `json:"DOGE" redis:"DOGE"`
// 	MLN     string `json:"MLN" redis:"MLN"`
// 	GTC     string `json:"GTC" redis:"GTC"`
// 	AMP     string `json:"AMP" redis:"AMP"`
// 	SHIB    string `json:"SHIB" redis:"SHIB"`
// 	CHZ     string `json:"CHZ" redis:"CHZ"`
// 	KEEP    string `json:"KEEP" redis:"KEEP"`
// 	LPT     string `json:"LPT" redis:"LPT"`
// 	QNT     string `json:"QNT" redis:"QNT"`
// 	BOND    string `json:"BOND" redis:"BOND"`
// 	RLY     string `json:"RLY" redis:"RLY"`
// 	CLV     string `json:"CLV" redis:"CLV"`
// 	FARM    string `json:"FARM" redis:"FARM"`
// 	MASK    string `json:"MASK" redis:"MASK"`
// 	ANT     string `json:"ANT" redis:"ANT"`
// 	FET     string `json:"FET" redis:"FET"`
// 	PAX     string `json:"PAX" redis:"PAX"`
// 	ACH     string `json:"ACH" redis:"ACH"`
// 	ASM     string `json:"ASM" redis:"ASM"`
// 	PLA     string `json:"PLA" redis:"PLA"`
// 	RAI     string `json:"RAI" redis:"RAI"`
// 	TRIBE   string `json:"TRIBE" redis:"TRIBE"`
// 	ORN     string `json:"ORN" redis:"ORN"`
// 	IOTX    string `json:"IOTX" redis:"IOTX"`
// 	UST     string `json:"UST" redis:"UST"`
// 	QUICK   string `json:"QUICK" redis:"QUICK"`
// 	AXS     string `json:"AXS" redis:"AXS"`
// 	REQ     string `json:"REQ" redis:"REQ"`
// 	WLUNA   string `json:"WLUNA" redis:"WLUNA"`
// 	TRU     string `json:"TRU" redis:"TRU"`
// 	RAD     string `json:"RAD" redis:"RAD"`
// 	COTI    string `json:"COTI" redis:"COTI"`
// 	DDX     string `json:"DDX" redis:"DDX"`
// 	SUKU    string `json:"SUKU" redis:"SUKU"`
// 	RGT     string `json:"RGT" redis:"RGT"`
// 	XYO     string `json:"XYO" redis:"XYO"`
// 	ZEN     string `json:"ZEN" redis:"ZEN"`
// 	AST     string `json:"AST" redis:"AST"`
// 	AUCTION string `json:"AUCTION" redis:"AUCTION"`
// 	BUSD    string `json:"BUSD" redis:"BUSD"`
// 	JASMY   string `json:"JASMY" redis:"JASMY"`
// 	WCFG    string `json:"WCFG" redis:"WCFG"`
// 	BTRST   string `json:"BTRST" redis:"BTRST"`
// 	AGLD    string `json:"AGLD" redis:"AGLD"`
// 	AVAX    string `json:"AVAX" redis:"AVAX"`
// 	FX      string `json:"FX" redis:"FX"`
// 	TRAC    string `json:"TRAC" redis:"TRAC"`
// 	LCX     string `json:"LCX" redis:"LCX"`
// 	ARPA    string `json:"ARPA" redis:"ARPA"`
// 	BADGER  string `json:"BADGER" redis:"BADGER"`
// 	KRL     string `json:"KRL" redis:"KRL"`
// 	PERP    string `json:"PERP" redis:"PERP"`
// 	RARI    string `json:"RARI" redis:"RARI"`
// 	DESO    string `json:"DESO" redis:"DESO"`
// 	API3    string `json:"API3" redis:"API3"`
// 	NCT     string `json:"NCT" redis:"NCT"`
// 	SHPING  string `json:"SHPING" redis:"SHPING"`
// 	UPI     string `json:"UPI" redis:"UPI"`
// 	CRO     string `json:"CRO" redis:"CRO"`
// 	MTL     string `json:"MTL" redis:"MTL"`
// 	ABT     string `json:"ABT" redis:"ABT"`
// 	CVX     string `json:"CVX" redis:"CVX"`
// 	AVT     string `json:"AVT" redis:"AVT"`
// 	MDT     string `json:"MDT" redis:"MDT"`
// 	VGX     string `json:"VGX" redis:"VGX"`
// 	ALCX    string `json:"ALCX" redis:"ALCX"`
// 	COVAL   string `json:"COVAL" redis:"COVAL"`
// 	FOX     string `json:"FOX" redis:"FOX"`
// 	MUSD    string `json:"MUSD" redis:"MUSD"`
// 	CELR    string `json:"CELR" redis:"CELR"`
// 	GALA    string `json:"GALA" redis:"GALA"`
// 	POWR    string `json:"POWR" redis:"POWR"`
// 	GYEN    string `json:"GYEN" redis:"GYEN"`
// 	ALICE   string `json:"ALICE" redis:"ALICE"`
// 	INV     string `json:"INV" redis:"INV"`
// 	LQTY    string `json:"LQTY" redis:"LQTY"`
// 	PRO     string `json:"PRO" redis:"PRO"`
// 	SPELL   string `json:"SPELL" redis:"SPELL"`
// 	ENS     string `json:"ENS" redis:"ENS"`
// 	DIA     string `json:"DIA" redis:"DIA"`
// 	BLZ     string `json:"BLZ" redis:"BLZ"`
// 	CTX     string `json:"CTX" redis:"CTX"`
// 	ERN     string `json:"ERN" redis:"ERN"`
// 	IDEX    string `json:"IDEX" redis:"IDEX"`
// 	MCO2    string `json:"MCO2" redis:"MCO2"`
// 	POLS    string `json:"POLS" redis:"POLS"`
// 	SUPER   string `json:"SUPER" redis:"SUPER"`
// 	UNFI    string `json:"UNFI" redis:"UNFI"`
// 	STX     string `json:"STX" redis:"STX"`
// 	KSM     string `json:"KSM" redis:"KSM"`
// 	GODS    string `json:"GODS" redis:"GODS"`
// 	IMX     string `json:"IMX" redis:"IMX"`
// 	RBN     string `json:"RBN" redis:"RBN"`
// 	BICO    string `json:"BICO" redis:"BICO"`
// 	GFI     string `json:"GFI" redis:"GFI"`
// 	ATA     string `json:"ATA" redis:"ATA"`
// 	GLM     string `json:"GLM" redis:"GLM"`
// 	MPL     string `json:"MPL" redis:"MPL"`
// 	PLU     string `json:"PLU" redis:"PLU"`
// 	SWFTC   string `json:"SWFTC" redis:"SWFTC"`
// 	SAND    string `json:"SAND" redis:"SAND"`
// 	OCEAN   string `json:"OCEAN" redis:"OCEAN"`
// 	GNO     string `json:"GNO" redis:"GNO"`
// 	FIDA    string `json:"FIDA" redis:"FIDA"`
// 	ORCA    string `json:"ORCA" redis:"ORCA"`
// 	CRPT    string `json:"CRPT" redis:"CRPT"`
// 	QSP     string `json:"QSP" redis:"QSP"`
// 	RNDR    string `json:"RNDR" redis:"RNDR"`
// 	NEST    string `json:"NEST" redis:"NEST"`
// 	PRQ     string `json:"PRQ" redis:"PRQ"`
// 	HOPR    string `json:"HOPR" redis:"HOPR"`
// 	JUP     string `json:"JUP" redis:"JUP"`
// 	MATH    string `json:"MATH" redis:"MATH"`
// 	SYN     string `json:"SYN" redis:"SYN"`
// 	AIOZ    string `json:"AIOZ" redis:"AIOZ"`
// 	WAMPL   string `json:"WAMPL" redis:"WAMPL"`
// 	AERGO   string `json:"AERGO" redis:"AERGO"`
// 	INDEX   string `json:"INDEX" redis:"INDEX"`
// 	TONE    string `json:"TONE" redis:"TONE"`
// 	HIGH    string `json:"HIGH" redis:"HIGH"`
// 	GUSD    string `json:"GUSD" redis:"GUSD"`
// 	FLOW    string `json:"FLOW" redis:"FLOW"`
// 	ROSE    string `json:"ROSE" redis:"ROSE"`
// 	OP      string `json:"OP" redis:"OP"`
// 	APE     string `json:"APE" redis:"APE"`
// 	MINA    string `json:"MINA" redis:"MINA"`
// 	MUSE    string `json:"MUSE" redis:"MUSE"`
// 	SYLO    string `json:"SYLO" redis:"SYLO"`
// 	CBETH   string `json:"CBETH" redis:"CBETH"`
// 	DREP    string `json:"DREP" redis:"DREP"`
// 	ELA     string `json:"ELA" redis:"ELA"`
// 	FORT    string `json:"FORT" redis:"FORT"`
// 	ALEPH   string `json:"ALEPH" redis:"ALEPH"`
// 	DEXT    string `json:"DEXT" redis:"DEXT"`
// 	FIS     string `json:"FIS" redis:"FIS"`
// 	BIT     string `json:"BIT" redis:"BIT"`
// 	GMT     string `json:"GMT" redis:"GMT"`
// 	GST     string `json:"GST" redis:"GST"`
// 	MEDIA   string `json:"MEDIA" redis:"MEDIA"`
// 	C98     string `json:"C98" redis:"C98"`
// 	ARB     string `json:"ARB" redis:"ARB"`
// 	TIME    string `json:"TIME" redis:"TIME"`
// 	RPL     string `json:"RPL" redis:"RPL"`
// 	MXC     string `json:"MXC" redis:"MXC"`
// 	HBAR    string `json:"HBAR" redis:"HBAR"`
// 	KAVA    string `json:"KAVA" redis:"KAVA"`
// 	SPA     string `json:"SPA" redis:"SPA"`
// 	EGLD    string `json:"EGLD" redis:"EGLD"`
// 	GHST    string `json:"GHST" redis:"GHST"`
// 	NEAR    string `json:"NEAR" redis:"NEAR"`
// 	INJ     string `json:"INJ" redis:"INJ"`
// 	AUDIO   string `json:"AUDIO" redis:"AUDIO"`
// 	MONA    string `json:"MONA" redis:"MONA"`
// 	TVK     string `json:"TVK" redis:"TVK"`
// 	POND    string `json:"POND" redis:"POND"`
// 	DYP     string `json:"DYP" redis:"DYP"`
// 	LDO     string `json:"LDO" redis:"LDO"`
// 	LIT     string `json:"LIT" redis:"LIT"`
// 	XMON    string `json:"XMON" redis:"XMON"`
// 	ILV     string `json:"ILV" redis:"ILV"`
// 	PUNDIX  string `json:"PUNDIX" redis:"PUNDIX"`
// 	PYR     string `json:"PYR" redis:"PYR"`
// 	PNG     string `json:"PNG" redis:"PNG"`
// 	METIS   string `json:"METIS" redis:"METIS"`
// 	RARE    string `json:"RARE" redis:"RARE"`
// 	QI      string `json:"QI" redis:"QI"`
// 	MSOL    string `json:"MSOL" redis:"MSOL"`
// 	OSMO    string `json:"OSMO" redis:"OSMO"`
// 	XCN     string `json:"XCN" redis:"XCN"`
// 	DAR     string `json:"DAR" redis:"DAR"`
// 	MAGIC   string `json:"MAGIC" redis:"MAGIC"`
// 	AURORA  string `json:"AURORA" redis:"AURORA"`
// 	BOBA    string `json:"BOBA" redis:"BOBA"`
// 	VOXEL   string `json:"VOXEL" redis:"VOXEL"`
// 	OOKI    string `json:"OOKI" redis:"OOKI"`
// 	MULTI   string `json:"MULTI" redis:"MULTI"`
// 	LOKA    string `json:"LOKA" redis:"LOKA"`
// 	T       string `json:"T" redis:"T"`
// 	MNDE    string `json:"MNDE" redis:"MNDE"`
// 	STG     string `json:"STG" redis:"STG"`
// 	GAL     string `json:"GAL" redis:"GAL"`
// 	Num00   string `json:"00" redis:"Num00"`
// 	HFT     string `json:"HFT" redis:"HFT"`
// 	DIMO    string `json:"DIMO" redis:"DIMO"`
// 	EUROC   string `json:"EUROC" redis:"EUROC"`
// 	BLUR    string `json:"BLUR" redis:"BLUR"`
// 	APT     string `json:"APT" redis:"APT"`
// 	WAXL    string `json:"WAXL" redis:"WAXL"`
// 	LSETH   string `json:"LSETH" redis:"LSETH"`
// 	SUI     string `json:"SUI" redis:"SUI"`
// 	AXL     string `json:"AXL" redis:"AXL"`
// 	ACS     string `json:"ACS" redis:"ACS"`
// 	FLR     string `json:"FLR" redis:"FLR"`
// 	PRIME   string `json:"PRIME" redis:"PRIME"`
// }

func (l coinbaseLibImpl) GetUSDTtoEUROExchangeRate(ctx context.Context, currency string) (float64, error) {
	response := USDTEuroExchangeRate{}
	coinbaseUrl := exchangeRateUrl + url.QueryEscape(currency)
	res, _, errs := gorequest.New().Get(coinbaseUrl).
		Retry(3, 3, http.StatusBadGateway, http.StatusInternalServerError).EndStruct(&response)
	if len(errs) > 0 {
		logger.Error(ctx, "error when get exchange rate from coinbase", "errors", errs)
		return 0, errs[0]
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("response is not success with code %d", res.StatusCode)
	}
	rate, err := strconv.ParseFloat(response.Data.Rates.USDT, 64)
	if err != nil {
		return 0, err
	}
	logger.Info(ctx, "get exchange rate from coinbase", "rate", rate)
	return rate, nil
}

func (l coinbaseLibImpl) AllExchangeRates(ctx context.Context, currency string) (Rates, error) {
	response := AllExchangeRates{}
	coinbaseUrl := exchangeRateUrl + url.QueryEscape(currency)
	res, _, errs := gorequest.New().Get(coinbaseUrl).
		Retry(3, 3, http.StatusBadGateway, http.StatusInternalServerError).EndStruct(&response)
	if len(errs) > 0 {
		logger.Error(ctx, "error when get exchange rate from coinbase", "errors", errs)
		return Rates{}, errs[0]
	}
	if res.StatusCode != http.StatusOK {
		return Rates{}, fmt.Errorf("response is not success with code %d", res.StatusCode)
	}
	logger.Info(ctx, "get exchange rate from coinbase", "response", response)
	return response.Data.Rates, nil
}
