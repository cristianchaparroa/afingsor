package afingsor

type Code string

const (
	FingerPrintOk                 Code = 0x00
	FingerPrintPacketRecieveerr   Code = 0x01
	FingerPrintNoFinger           Code = 0x02
	FingerPrintImageFail          Code = 0x03
	FingerPrintImageMess          Code = 0x06
	FingerPrintFeatureFail        Code = 0x07
	FingerPrintNoMatch            Code = 0x08
	FingerPrintNotFound           Code = 0x09
	FingerPrintEnrollMismatch     Code = 0x0A
	FingerPrintBadLocation        Code = 0x0B
	FingerPrintDBRangeFail        Code = 0x0C
	FingerPrintUploadFeatureFail  Code = 0x0D
	FingerPrintPacketResponseFail Code = 0x0E
	FingerPrintUploadFail         Code = 0x0F
	FingerPrintDeleteFail         Code = 0x10
	FingerPrintDBClearFail        Code = 0x11
	FingerPrintPassFail           Code = 0x13
	FingerPrintInvalidImage       Code = 0x15
	FingerPrintFlasherr           Code = 0x18
	FingerPrintInvalidREG         Code = 0x1A
	FingerPrintAddrCode           Code = 0x20
	FingerPrintPassVerify         Code = 0x21
	FingerPrintSTARTCode          Code = 0xEF01
	FingerPrintCommandPacket      Code = 0x1
	FingerPrintDataPacket         Code = 0x2
	FingerPrintACKPacket          Code = 0x7
	FingerPrintEndDataPacket      Code = 0x8
	FingerPrintTimeout            Code = 0xFF
	FingerPrintBadPacket          Code = 0xFE
	FingerPrintGetImage           Code = 0x01
	FingerPrintImage2TZ           Code = 0x02
	FingerPrintRegModel           Code = 0x05
	FingerPrintStore              Code = 0x06
	FingerPrintLoad               Code = 0x07
	FingerPrintUpload             Code = 0x08
	FingerPrintDelete             Code = 0x0C
	FingerPrintEmpty              Code = 0x0D
	FingerPrintVerifyPassword     Code = 0x13
	FingerPrintHiSpeedSearch      Code = 0x1B
	FingerPrintTemplateCount      Code = 0x1D

	//	DEFAULTTIMEOUT FingerPrintDEBUG in milliseconds
	DEFAULTTIMEOUT Code = 5000
)
