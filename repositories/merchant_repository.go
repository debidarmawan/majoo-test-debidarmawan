package repositories

import (
	"fmt"
	"majoo-test-debidarmawan/config"
	"majoo-test-debidarmawan/models"
	"net/http"
)

type MerchantRepo interface {
	MerchantOmzet(userID uint64, startDate string, endDate string, limit int, page int) chan models.Result
	MerchantsOutletOmzet(userID uint64, outletID uint64, startDate string, endDate string, limit int, page int) chan models.Result
}

type merchantRepo struct {
	dbConn *config.DbConnection
}

func NewMerchantRepo(dbConn *config.DbConnection) MerchantRepo {
	return &merchantRepo{
		dbConn: dbConn,
	}
}

func (mr *merchantRepo) MerchantOmzet(userID uint64, startDate string, endDate string, limit int, page int) chan models.Result {
	output := make(chan models.Result)
	go func() {
		var (
			merchant     models.Merchant
			transactions []models.Transaction
			offset       int
		)
		err := mr.dbConn.MajooDB.Select("id", "merchant_name").Where("user_id = ?", userID).First(&merchant).Error
		if err != nil {
			fmt.Println(err)
			output <- models.Result{StatusCode: http.StatusInternalServerError, Error: err, Data: nil, Message: "Internal Server Error"}
			return
		}

		if page == 0 {
			page = 1
		}
		if limit == 0 {
			limit = 10
		}

		offset = (page - 1) * limit

		trxErr := mr.dbConn.MajooDB.Select("bill_total", "created_at").
			Where("created_at BETWEEN ? AND ?", startDate, endDate).
			Where("merchant_id = ?", merchant.ID).
			Order("created_at asc").
			Offset(offset).
			Limit(limit).
			Find(&transactions).Error
		if trxErr != nil {
			fmt.Println(trxErr)
			output <- models.Result{StatusCode: http.StatusInternalServerError, Error: trxErr, Data: nil, Message: "Internal Server Error"}
			return
		}

		respData := map[string]interface{}{
			"merchant_data": merchant,
			"transactions":  transactions,
		}

		output <- models.Result{StatusCode: http.StatusOK, Data: respData, Error: nil, Message: "Success"}
	}()
	return output
}

func (mr *merchantRepo) MerchantsOutletOmzet(userID uint64, outletID uint64, startDate string, endDate string, limit int, page int) chan models.Result {
	output := make(chan models.Result)
	go func() {
		var (
			merchant     models.Merchant
			outlets      []models.Outlet
			transactions []models.Transaction
			offset       int
			outletIDs    []int
			outletErr    error
		)
		err := mr.dbConn.MajooDB.Select("id", "merchant_name").Where("user_id = ?", userID).First(&merchant).Error
		if err != nil {
			fmt.Println(err)
			output <- models.Result{StatusCode: http.StatusInternalServerError, Error: err, Data: nil, Message: "Internal Server Error"}
			return
		}

		if outletID != 0 {
			outletErr = mr.dbConn.MajooDB.Select("id", "outlet_name").Where("merchant_id", merchant.ID).Where("id", outletID).Find(&outlets).Error
		} else {
			outletErr = mr.dbConn.MajooDB.Select("id", "outlet_name").Where("merchant_id", merchant.ID).Find(&outlets).Error
		}

		if outletErr != nil {
			fmt.Println(outletErr)
			output <- models.Result{StatusCode: http.StatusInternalServerError, Error: outletErr, Data: nil, Message: "Internal Server Error"}
			return
		}

		if len(outlets) > 0 {
			for _, o := range outlets {
				outletIDs = append(outletIDs, int(o.ID))
			}
		}

		if page == 0 {
			page = 1
		}
		if limit == 0 {
			limit = 10
		}

		offset = (page - 1) * limit

		trxErr := mr.dbConn.MajooDB.Select("outlet_id", "bill_total", "created_at").
			Where("created_at BETWEEN ? AND ?", startDate, endDate).
			Where("merchant_id = ?", merchant.ID).
			Where("outlet_id IN ?", outletIDs).
			Order("created_at asc").
			Offset(offset).
			Limit(limit).
			Find(&transactions).Error
		if trxErr != nil {
			fmt.Println(trxErr)
			output <- models.Result{StatusCode: http.StatusInternalServerError, Error: trxErr, Data: nil, Message: "Internal Server Error"}
			return
		}

		respData := map[string]interface{}{
			"merchant_data": merchant,
			"outlets":       outlets,
			"transactions":  transactions,
		}

		output <- models.Result{StatusCode: http.StatusOK, Data: respData, Error: nil, Message: "Success"}
	}()
	return output
}
