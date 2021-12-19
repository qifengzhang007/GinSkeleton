package error_record

import "goskeleton/app/global/variable"

// ErrorDeal 记录错误
func ErrorDeal(err error) error {
	if err != nil {
		variable.ZapLog.Error(err.Error())
	}
	return err
}
