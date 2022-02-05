package price

import "fmt"

type PriceInfo struct {
	Station string
	Detail  string
}

type PriceInfos []*PriceInfo

func (p *PriceInfo) String() string {
	return fmt.Sprintf("Station: %s, Detail: %s", p.Station, p.Detail)
}
