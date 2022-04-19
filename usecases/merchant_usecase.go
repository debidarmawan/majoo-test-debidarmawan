package usecases

import (
	"encoding/json"
	"majoo-test-debidarmawan/models"
	"majoo-test-debidarmawan/repositories"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
	"golang.org/x/exp/slices"
)

type MerchantUseCase interface {
	MerchantOmzet(request models.MerchantOmzet) chan models.Result
	MerchantsOutletOmzet(request models.MerchantOutletOmzet) chan models.Result
}

type merchantUsecase struct {
	merchantRepo repositories.MerchantRepo
}

func NewMerchantUseCase(merchantRepo repositories.MerchantRepo) MerchantUseCase {
	return &merchantUsecase{
		merchantRepo: merchantRepo,
	}
}

func (mu *merchantUsecase) MerchantOmzet(request models.MerchantOmzet) chan models.Result {
	output := make(chan models.Result)
	go func() {
		var now time.Time
		if request.Period == "" {
			now = time.Now()
		} else {
			now, _ = time.Parse("2006-01", request.Period)
		}
		year, month, _ := now.Date()
		startDate := time.Date(year, month, 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
		endDate := time.Date(year, month+1, 0, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
		result := <-mu.merchantRepo.MerchantOmzet(request.UserID, startDate, endDate, request.Limit, request.Page)
		if result.StatusCode == fiber.StatusOK {
			var (
				resData      map[string]interface{}
				results      models.MerchantOmzetResponse
				transactions []models.Transaction
			)
			resByte, _ := json.Marshal(result.Data)
			_ = json.Unmarshal(resByte, &resData)
			if resData != nil {
				merchantByte, _ := json.Marshal(resData["merchant_data"])
				results.MerchantID = uint64(gjson.Get(string(merchantByte), "id").Int())
				results.MerchantName = gjson.Get(string(merchantByte), "merchant_name").Str

				transactionByte, _ := json.Marshal(resData["transactions"])
				_ = json.Unmarshal(transactionByte, &transactions)

				if len(transactions) > 0 {
					for _, data := range transactions {
						var omzetData models.DailyOmzet
						dateFormated := data.CreatedAt.Format("2006-01-02")
						idx := slices.IndexFunc(results.DailyOmzet, func(x models.DailyOmzet) bool { return x.Date == dateFormated })
						if idx < 0 {
							omzetData.Date = dateFormated
							omzetData.Omzet = int64(data.BillTotal)
							results.DailyOmzet = append(results.DailyOmzet, omzetData)
						} else {
							results.DailyOmzet[idx].Omzet += int64(data.BillTotal)
						}
					}
				} else {
					results.DailyOmzet = make([]models.DailyOmzet, 0)
				}
			}
			result.Data = results
		}
		output <- result
	}()
	return output
}

func (mu *merchantUsecase) MerchantsOutletOmzet(request models.MerchantOutletOmzet) chan models.Result {
	output := make(chan models.Result)
	go func() {
		var now time.Time
		if request.Period == "" {
			now = time.Now()
		} else {
			now, _ = time.Parse("2006-01", request.Period)
		}
		year, month, _ := now.Date()
		startDate := time.Date(year, month, 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
		endDate := time.Date(year, month+1, 0, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
		result := <-mu.merchantRepo.MerchantsOutletOmzet(request.UserID, request.OutletID, startDate, endDate, request.Limit, request.Page)
		if result.StatusCode == fiber.StatusOK {
			var (
				resData      map[string]interface{}
				results      models.MerchantOutletOmzetResponse
				outlets      []models.OutletData
				transactions []models.Transaction
			)
			resByte, _ := json.Marshal(result.Data)
			_ = json.Unmarshal(resByte, &resData)
			if resData != nil {
				merchantByte, _ := json.Marshal(resData["merchant_data"])
				results.MerchantID = uint64(gjson.Get(string(merchantByte), "id").Int())
				results.MerchantName = gjson.Get(string(merchantByte), "merchant_name").Str

				transactionByte, _ := json.Marshal(resData["transactions"])
				_ = json.Unmarshal(transactionByte, &transactions)

				outletsByte, _ := json.Marshal(resData["outlets"])
				_ = json.Unmarshal(outletsByte, &outlets)

				if len(outlets) > 0 {
					for _, o := range outlets {
						var outlet models.OutletData
						outlet.ID = o.ID
						outlet.OutletName = o.OutletName

						if len(transactions) > 0 {
							for _, data := range transactions {
								if data.OutletID == o.ID {
									var omzetData models.DailyOmzet
									dateFormated := data.CreatedAt.Format("2006-01-02")
									idx := slices.IndexFunc(outlet.DailyOmzet, func(x models.DailyOmzet) bool { return x.Date == dateFormated })
									if idx < 0 {
										omzetData.Date = dateFormated
										omzetData.Omzet = int64(data.BillTotal)
										outlet.DailyOmzet = append(outlet.DailyOmzet, omzetData)
									} else {
										outlet.DailyOmzet[idx].Omzet += int64(data.BillTotal)
									}
								}
							}
						} else {
							outlet.DailyOmzet = make([]models.DailyOmzet, 0)
						}
						results.OutletsData = append(results.OutletsData, outlet)
					}
				} else {
					results.OutletsData = make([]models.OutletData, 0)
				}
			}
			result.Data = results
		}
		output <- result
	}()
	return output
}
