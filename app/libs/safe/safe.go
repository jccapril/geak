package safe

import (
	"geak/libs/log"
	"go.uber.org/zap"
)

func Go(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error("go routine panic",zap.Any("err",err))
			}
		}()
		fn()
	}()
}