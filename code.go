package afingsor

const (
	FingerPrintOk                 = 0x00
	FingerPrintPacketRecieveerr   = 0x01
	FingerPrintNoFinger           = 0x02
	FingerPrintImageFail          = 0x03
	FingerPrintImageMess          = 0x06
	FingerPrintFeatureFail        = 0x07
	FingerPrintNoMatch            = 0x08
	FingerPrintNotFound           = 0x09
	FingerPrintEnrollMismatch     = 0x0A
	FingerPrintBadLocation        = 0x0B
	FingerPrintDBRangeFail        = 0x0C
	FingerPrintUploadFeatureFail  = 0x0D
	FingerPrintPacketResponseFail = 0x0E
	FingerPrintUploadFail         = 0x0F
	FingerPrintDeleteFail         = 0x10
	FingerPrintDBClearFail        = 0x11
	FingerPrintPassFail           = 0x13
	FingerPrintInvalidImage       = 0x15
	FingerPrintFlasherr           = 0x18
	FingerPrintInvalidREG         = 0x1A
	FingerPrintAddr               = 0x20
	FingerPrintPassVerify         = 0x21
	FingerPrintStartCode          = 0xEF01
	FingerPrintCommandPacket      = 0x1
	FingerPrintDataPacket         = 0x2
	FingerPrintACKPacket          = 0x7
	FingerPrintEndDataPacket      = 0x8
	FingerPrintTimeout            = 0xFF
	FingerPrintBadPacket          = 0xFE
	FingerPrintGetImage           = 0x01
	FingerPrintImage2TZ           = 0x02
	FingerPrintRegModel           = 0x05
	FingerPrintStore              = 0x06
	FingerPrintLoad               = 0x07
	FingerPrintUpload             = 0x08
	FingerPrintDelete             = 0x0C
	FingerPrintEmpty              = 0x0D
	FingerPrintVerifyPassword     = 0x13
	FingerPrintHiSpeedSearch      = 0x1B
	FingerPrintTemplateCount      = 0x1D
	DEFAULTTIMEOUT                = 5000 //	DEFAULTTIMEOUT FingerPrintDEBUG in milliseconds
)
