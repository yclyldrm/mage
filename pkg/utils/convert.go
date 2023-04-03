package utils

import (
	"mage/pkg/models"
	"strconv"
	"time"
)

func ConvertHexToInt(hexStr string) (int64, error) {
	return strconv.ParseInt(hexStr, 16, 64)
}

func ConvertTimetoString(timeData time.Time) string {
	return timeData.Format(models.TIMEFORMAT)
}
