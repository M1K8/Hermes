package aries

type Generator struct {
	baseUrl string
}

func (g *Generator) GetBaseUrl() string {
	return g.baseUrl
}

func (g *Generator) GetStockUrl(ticker string, stop, t_stop, pt int) (string, error) {
	return "", nil
}

func (g *Generator) GetOptionsUrl(ticker, expiry, strike string, stop, t_stop, pt int) (string, error) {
	return "", nil
}
