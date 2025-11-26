package types

type Input struct {
	Ip string `json:"ip"`
}

type Output struct {
	Response string `json:"response"`
}

type SSLScanInput struct {
	Ip string `json:"ip"`
}

type SSLScanOutput struct {
	Response string `json:"response"`
}
