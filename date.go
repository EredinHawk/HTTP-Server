package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// ScanDate получает параметры из URL строки и возвращает time.Duration в днях
func ScanDate(r *http.Request) (time.Time, error) {
	birthYear,  errYear  := strconv.Atoi(r.PathValue("year"))
	birthMonth, errMonth := strconv.Atoi(r.PathValue("month"))
	birthDay,   errDay   := strconv.Atoi(r.PathValue("day"))

	if errYear != nil || errMonth != nil || errDay != nil || birthYear > time.Now().Year() || birthYear <= 0 ||
	   birthMonth <= 0 || birthMonth > 12 || birthDay > 31 || birthDay <= 0 {
		return time.Time{}, fmt.Errorf("недопустимые значения в параметрах пути 'year' - %v, 'month' - %v, 'day' - %v", birthYear, birthMonth, birthDay)
	}

	return time.Date(birthYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC), nil
}

// AgeString возвращает расчет возраста в виде строки.
// Параметр birthday - день рождения.
func AgeString(birthday time.Time) string {
	age    := time.Now().Sub(birthday)
	years  := int(age.Hours() / 24 / 365)
	months := (int(age.Hours()/24) - (years * 365)) / 31
	days := int(age.Hours()/24) - (years * 365) - (months * 31)

	return fmt.Sprintf("%v лет %v месяцев %v дней\n", years, months, days)
}
