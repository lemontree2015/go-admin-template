package captcha

import (
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"sscmgroup.com/app/config"
	"sscmgroup.com/app/model"
	"sscmgroup.com/app/module/captcha/store"
	"time"
)

var capStore base64Captcha.Store

type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

func DriverStringFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverString = base64Captcha.NewDriverString(46, 140, 2, 2, 4, "234567890abcdefghjkmnpqrstuvwxyz", &color.RGBA{240, 240, 246, 246}, []string{"wqy-microhei.ttc"})
	driver := e.DriverString.ConvertFonts()
	cap := base64Captcha.NewCaptcha(driver, capStore)
	return cap.Generate()
}

func DriverDigitFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	driver := e.DriverDigit
	cap := base64Captcha.NewCaptcha(driver, capStore)
	return cap.Generate()
}

// Verify 校验验证码
func Verify(id, code string, clear bool) bool {
	return capStore.Verify(id, code, clear)
}

func InitCaptcha() {
	cfg := config.Conf.Captcha
	if cfg.Store == "redis" {
		capStore = store.NewRedisStoreWithCli(model.RedisClient("user_cache"), time.Duration(cfg.TTL)*time.Second, cfg.RedisPrefix)
		//captcha.SetCustomStore(store.NewRedisStoreWithCli(model.RedisClient("user_cache"), time.Duration(cfg.TTL)*time.Second, logger.Logger, cfg.RedisPrefix))
	}
}
