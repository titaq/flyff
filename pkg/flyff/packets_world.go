package flyff

// FPWelcome
const FPWelcomeCmdID = 0x00000000

type FPWelcome struct {
	FPWriter
	ID uint32
}

func (f *FPWelcome) Serialize() []byte {
	return f.
		Initialize().
		WriteUInt32(FPWelcomeCmdID).
		WriteInt32((int32)(f.ID)).
		finalize()
}
