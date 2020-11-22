package helper

import (
	"demo/database"
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type ResponseDefault struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type OtpInput struct {
	PhoneNumber string  `json:"phone_number"`
	Pass        *string `json:"password"`
	Otp         string  `json:"otp"`
}

func GenerateOtp(phone string) (err error, statusCode int) {
	res, err, statusCode := models.CheckExistsUserFromPhone(phone)
	log.Println(res, err)
	if err != nil {
		return err, http.StatusForbidden
	}
	// Get time_life from env
	tl := utils.TimeLife
	// Validate Phone Number : Phone number must have prefix 84
	/*err = CheckPhoneNumberFormat(phone)
	if err != nil {
		return err, http.StatusBadRequest
	}*/

	phone = strings.Replace(phone, "0", "+84", 1)

	// Random n digit for otp
	result := getRandNum(utils.NumberOfCharacter)

	// Generate model
	var obj models.Otp = models.Otp{
		PhoneNumber: phone,
		Otp:         result,
		FailedTimes: 0,
		TimeLife:    time.Now().Add(time.Duration(tl) * time.Second),
		Type:        "registration",
	}
	// Send Otp to user
	err = SendSMS(obj)
	if err != nil {
		return err, http.StatusBadRequest
	}
	// Save to database
	_, err, statusCode = models.CreateOtp(obj)
	return
}

func CheckOtp(input OtpInput) (res interface{}, err error, statusCode int) {
	/*if input.DeviceID == "" {
		return nil, fmt.Errorf("Check your device information"), http.StatusForbidden
	} else {
		log.Println("device_id: ", input.DeviceID)
	}*/
	// Validate Phone Number : Phone number must have prefix 84
	/*err = CheckPhoneNumberFormat(input.PhoneNumber)
	if err != nil {
		return res, err, http.StatusBadRequest
	}*/

	// Validate Otp : Otp must have 4 digits
	err = CheckOtpFormat(input.Otp)
	if err != nil {
		return res, err, http.StatusBadRequest
	}

	// Get the newest Otp
	fOtp, err := models.GetOneOtp(input.PhoneNumber)
	if err != nil {
		log.Println("error when check OTP", err.Error())
		return models.Otp{}, fmt.Errorf("Something when wrong with your OTP"), http.StatusForbidden
	}

	// Get failed_times from env
	ft := utils.FailedTimes
	if fOtp.FailedTimes > ft {
		return res, fmt.Errorf("Vượt quá số lần nhập sai "), http.StatusForbidden
	}
	if fOtp.Otp != input.Otp {
		var count int = fOtp.FailedTimes + 1
		return models.IncreaseOtpFailedTime(fOtp.ID, count)
	} else {
		update, err := models.UpdateOtp(fOtp.ID)
		log.Println(update)
		if err != nil {
			return models.Otp{}, err, http.StatusForbidden
		}
		dbPublic, err := database.GetDatabase("default")
		if err != nil {
			return res, err, http.StatusInternalServerError
		}
		user := models.User{PhoneNumber: input.PhoneNumber, Password: input.Pass}
		tx := dbPublic.Create(&user)
		if tx.Error == nil {
			return res, nil, http.StatusOK
		} else {
			return res, tx.Error, http.StatusForbidden
		}
	}
}

func CheckPhoneNumberFormat(phone string) (err error) {
	re := regexp.MustCompile("([0-9]{10})$")
	if !re.MatchString(phone) {
		return fmt.Errorf("Số điện thoại không hợp lệ. Vui lòng kiểm tra lại")
	}
	return nil
}

func CheckOtpFormat(otp string) (err error) {
	re := regexp.MustCompile("([0-9]{4})$")
	if !re.MatchString(otp) {
		return fmt.Errorf("Otp wrong format!")
	}
	return nil
}

type SMS struct {
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"text"`
}

func SendSMS(otp models.Otp) (err error) {
	accountSid := "AC3aecabde3d69f5b3a0ebeb23abb1518a"
	urlstr := "https://api.twilio.com/2010-04-01/Accounts/AC3aecabde3d69f5b3a0ebeb23abb1518a/Messages.json"
	authToken := "24f652b8decc6d4e4432317e26ac0bfe"
	obj := SMS{
		From: "+12056221744",
		To:   otp.PhoneNumber,
		Body: "Xin chao! Ma code dang ky Travel Website cua ban la: " + otp.Otp,
	}
	msgData := url.Values{}
	msgData.Set("To", obj.To)
	log.Println(obj.To)
	msgData.Set("From", obj.From)
	msgData.Set("Body", obj.Body)
	msgDataReader := *strings.NewReader(msgData.Encode())

	req, err := http.NewRequest("POST", urlstr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	log.Println(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			return nil
		}
	} else {
		return fmt.Errorf("Tao ma otp that bai")
	}
	return nil
}

var letterRunes = []rune("0123456789")

func getRandNum(n int) string {
	b := make([]rune, n)
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//func getRandNum() (int, error) {
//	noc := utils.NumberOfCharacter
//	low := int(math.Pow10(noc - 1))
//	top := int(math.Pow10(noc)) - low
//	result := int(low) + rand.Intn(top)
//	return result, nil
//}
