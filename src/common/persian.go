package common

import (
	"log"
	"regexp"
)

const IranianMobileNumberPattern string = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func IranianMobileNumberValidate(mobileNumber string) bool {
	res, err := regexp.MatchString(IranianMobileNumberPattern, mobileNumber)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}
