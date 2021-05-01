package flyff

func NewFFPacketIn_Join(reader *FPReader) *FFPacketIn_Join {
	p := new(FFPacketIn_Join)
	p.Read(reader)
	return p
}

func (p *FFPacketIn_Join) Read(reader *FPReader) {
	_ = reader.ReadUInt32()
	_ = reader.ReadUInt32()
	p.AuthKey = reader.ReadUInt32()
	_ = reader.ReadUInt32()
	_ = reader.ReadUInt32()
	_ = reader.ReadUInt32()
	_ = reader.ReadUInt32()
	p.Slot = (uint32)(reader.ReadByte())
}
