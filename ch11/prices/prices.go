package prices

import (
	"fmt"
	"go-learning/ch11/conversion"
	"go-learning/ch11/iomanagers"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64              `json:"tax_rate"`
	InputPrices       []float64            `json:"input_prices"`
	TaxIncludedPrices map[string]string    `json:"tax_included_prices"`
	IOManager         iomanagers.IOManager `json:"-"`
}

func NewTaxIncludedPriceJob(manager iomanagers.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: nil,
		IOManager:   manager,
	}
}

func (t *TaxIncludedPriceJob) LoadPriceData() error {
	lines, err := t.IOManager.ReadLines()
	if err != nil {
		return err
	}
	floatprices, err := conversion.StringToFloat(lines)
	if err != nil {
		return err
	}
	t.InputPrices = floatprices
	return nil
}

func (t *TaxIncludedPriceJob) Process() error {
	err := t.LoadPriceData()
	if err != nil {
		return err
	}
	result := make(map[string]string)
	for _, price := range t.InputPrices {
		newPrice := price * (1 + t.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", newPrice)
	}
	t.TaxIncludedPrices = result
	return t.IOManager.WriteResult(t)
}
