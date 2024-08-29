package models

// JsonInternalResponse define a estrutura dos dados de resposta interna para a cotação de câmbio.
// Cada campo representa uma informação específica sobre a cotação do dólar em relação ao real.
type JsonInternalResponse struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

// JsonExternalResponse define a estrutura para a resposta externa da API de cotação de câmbio.
// A resposta externa encapsula a estrutura interna JsonInternalResponse sob a chave "USDBRL".
type JsonExternalResponse struct {
	JsonInternalResponse `json:"USDBRL"`
}

// ContextKey é um tipo personalizado para chaves usadas em contextos.
// Isso ajuda a evitar colisões de chave ao armazenar e recuperar valores do contexto.
type ContextKey string

// ExchangeRateKey é a chave usada para armazenar a estrutura JsonExternalResponse no contexto.
const ExchangeRateKey ContextKey = "exchangeRate"
