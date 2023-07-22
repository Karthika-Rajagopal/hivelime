package brotli


type cmdLutElement struct {
	insert_len_extra_bits byte
	copy_len_extra_bits   byte
	distance_code         int8
	context               byte
	insert_len_offset     uint16
	copy_len_offset       uint16
}

var kCmdLut = [numCommandSymbols]cmdLutElement{
	cmdLutElement{0x00, 0x00, 0, 0x00, 0x0000, 0x0002},
	cmdLutElement{0x00, 0x00, 0, 0x01, 0x0000, 0x0003},
	cmdLutElement{0x00, 0x00, 0, 0x02, 0x0000, 0x0004},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0000, 0x0005},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0000, 0x0006},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0000, 0x0007},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0000, 0x0008},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0000, 0x0009},
	cmdLutElement{0x00, 0x00, 0, 0x00, 0x0001, 0x0002},
	cmdLutElement{0x00, 0x00, 0, 0x01, 0x0001, 0x0003},
	cmdLutElement{0x00, 0x00, 0, 0x02, 0x0001, 0x0004},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0001, 0x0005},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0001, 0x0006},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0001, 0x0007},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0001, 0x0008},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0001, 0x0009},
	cmdLutElement{0x00, 0x00, 0, 0x00, 0x0002, 0x0002},
	cmdLutElement{0x00, 0x00, 0, 0x01, 0x0002, 0x0003},
	cmdLutElement{0x00, 0x00, 0, 0x02, 0x0002, 0x0004},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0002, 0x0005},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0002, 0x0006},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0002, 0x0007},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0002, 0x0008},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0002, 0x0009},
	cmdLutElement{0x00, 0x00, 0, 0x00, 0x0003, 0x0002},
	cmdLutElement{0x00, 0x00, 0, 0x01, 0x0003, 0x0003},
	cmdLutElement{0x00, 0x00, 0, 0x02, 0x0003, 0x0004},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0003, 0x0005},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0003, 0x0006},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0003, 0x0007},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0003, 0x0008},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0003, 0x0009},
	cmdLutElement{0x00, 0x00, 0, 0x00, 0x0004, 0x0002},
	cmdLutElement{0x00, 0x00, 0, 0x01, 0x0004, 0x0003},
	cmdLutElement{0x00, 0x00, 0, 0x02, 0x0004, 0x0004},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0004, 0x0005},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0004, 0x0006},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0004, 0x0007},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0004, 0x0008},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0004, 0x0009},
	cmdLutElement{0x00, 0x00, 0, 0x00, 0x0005, 0x0002},
	cmdLutElement{0x00, 0x00, 0, 0x01, 0x0005, 0x0003},
	cmdLutElement{0x00, 0x00, 0, 0x02, 0x0005, 0x0004},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0005, 0x0005},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0005, 0x0006},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0005, 0x0007},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0005, 0x0008},
	cmdLutElement{0x00, 0x00, 0, 0x03, 0x0005, 0x0009},
	cmdLutElement{0x01, 0x00, 0, 0x00, 0x0006, 0x0002},
	cmdLutElement{0x01, 0x00, 0, 0x01, 0x0006, 0x0003},
	cmdLutElement{0x01, 0x00, 0, 0x02, 0x0006, 0x0004},
	cmdLutElement{0x01, 0x00, 0, 0x03, 0x0006, 0x0005},
	cmdLutElement{0x01, 0x00, 0, 0x03, 0x0006, 0x0006},
	cmdLutElement{0x01, 0x00, 0, 0x03, 0x0006, 0x0007},
	cmdLutElement{0x01, 0x00, 0, 0x03, 0x0006, 0x0008},
	cmdLutElement{0x01, 0x00, 0, 0x03, 0x0006, 0x0009},
	cmdLutElement{0x01, 0x00, 0, 0x00, 0x0008, 0x0002},
	cmdLutElement{0x01, 0x00, 0, 0x01, 0x0008, 0x0003},
	cmdLutElement{0x01, 0x00, 0, 0x02, 0x0008, 0x0004},
	cmdLutElement{0x01, 0x00, 0, 0x03, 0x0008, 0x0005},
	cmdLutElement{0x01, 0x00, 0, 0x03, 0x0008, 0x0006},
	cmdLutElement{0x01, 0x00, 0, 0x03, 0x0008, 0x0007},
	cmdLutElement{0x01, 0x00, 0, 0x03, 0x0008, 0x0008},
	cmdLutElement{0x01, 0x00, 0, 0x03, 0x0008, 0x0009},
	cmdLutElement{0x00, 0x01, 0, 0x03, 0x0000, 0x000a},
	cmdLutElement{0x00, 0x01, 0, 0x03, 0x0000, 0x000c},
	cmdLutElement{0x00, 0x02, 0, 0x03, 0x0000, 0x000e},
	cmdLutElement{0x00, 0x02, 0, 0x03, 0x0000, 0x0012},
	cmdLutElement{0x00, 0x03, 0, 0x03, 0x0000, 0x0016},
	cmdLutElement{0x00, 0x03, 0, 0x03, 0x0000, 0x001e},
	cmdLutElement{0x00, 0x04, 0, 0x03, 0x0000, 0x0026},
	cmdLutElement{0x00, 0x04, 0, 0x03, 0x0000, 0x0036},
	cmdLutElement{0x00, 0x01, 0, 0x03, 0x0001, 0x000a},
	cmdLutElement{0x00, 0x01, 0, 0x03, 0x0001, 0x000c},
	cmdLutElement{0x00, 0x02, 0, 0x03, 0x0001, 0x000e},
	cmdLutElement{0x00, 0x02, 0, 0x03, 0x0001, 0x0012},
	cmdLutElement{0x00, 0x03, 0, 0x03, 0x0001, 0x0016},
	cmdLutElement{0x00, 0x03, 0, 0x03, 0x0001, 0x001e},
	cmdLutElement{0x00, 0x04, 0, 0x03, 0x0001, 0x0026},
	cmdLutElement{0x00, 0x04, 0, 0x03, 0x0001, 0x0036},
	cmdLutElement{0x00, 0x01, 0, 0x03, 0x0002, 0x000a},
	cmdLutElement{0x00, 0x01, 0, 0x03, 0x0002, 0x000c},
	cmdLutElement{0x00, 0x02, 0, 0x03, 0x0002, 0x000e},
	cmdLutElement{0x00, 0x02, 0, 0x03, 0x0002, 0x0012},
	cmdLutElement{0x00, 0x03, 0, 0x03, 0x0002, 0x0016},
	cmdLutElement{0x00, 0x03, 0, 0x03, 0x0002, 0x001e},
	cmdLutElement{0x00, 0x04, 0, 0x03, 0x0002, 0x0026},
	cmdLutElement{0x00, 0x04, 0, 0x03, 0x0002, 0x0036},
	cmdLutElement{0x00, 0x01, 0, 0x03, 0x0003, 0x000a},
	cmdLutElement{0x00, 0x01, 0, 0x03, 0x0003, 0x000c},
	cmdLutElement{0x00, 0x02, 0, 0x03, 0x0003, 0x000e},
	cmdLutElement{0x00, 0x02, 0, 0x03, 0x0003, 0x0012},
	cmdLutElement{0x00, 0x03, 0, 0x03, 0x0003, 0x0016},
	cmdLutElement{0x00, 0x03, 0, 0x03, 0x0003, 0x001e},
	cmdLutElement{0x00, 0x04, 0, 0x03, 0x0003, 0x0026},
	cmdLutElement{0x00, 0x04, 0, 0x03, 0x0003, 0x0036},
	cmdLutElement{0x00, 0x01, 0, 0x03, 0x0004, 0x000a},
	cmdLutElement{0x00, 0x01, 0, 0x03, 0x0004, 0x000c},
	cmdLutElement{0x00, 0x02, 0, 0x03, 0x0004, 0x000e},
	cmdLutElement{0x00, 0x02, 0, 0x03, 0x0004, 0x0012},
	cmdLutElement{0x00, 0x03, 0, 0x03, 0x0004, 0x0016},
	cmdLutElement{0x00, 0x03, 0, 0x03, 0x0004, 0x001e},
	cmdLutElement{0x00, 0x04, 0, 0x03, 0x0004, 0x0026},
	cmdLutElement{0x00, 0x04, 0, 0x03, 0x0004, 0x0036},
	cmdLutElement{0x00, 0x01, 0, 0x03, 0x0005, 0x000a},
	cmdLutElement{0x00, 0x01, 0, 0x03, 0x0005, 0x000c},
	cmdLutElement{0x00, 0x02, 0, 0x03, 0x0005, 0x000e},
	cmdLutElement{0x00, 0x02, 0, 0x03, 0x0005, 0x0012},
	cmdLutElement{0x00, 0x03, 0, 0x03, 0x0005, 0x0016},
	cmdLutElement{0x00, 0x03, 0, 0x03, 0x0005, 0x001e},
	cmdLutElement{0x00, 0x04, 0, 0x03, 0x0005, 0x0026},
	cmdLutElement{0x00, 0x04, 0, 0x03, 0x0005, 0x0036},
	cmdLutElement{0x01, 0x01, 0, 0x03, 0x0006, 0x000a},
	cmdLutElement{0x01, 0x01, 0, 0x03, 0x0006, 0x000c},
	cmdLutElement{0x01, 0x02, 0, 0x03, 0x0006, 0x000e},
	cmdLutElement{0x01, 0x02, 0, 0x03, 0x0006, 0x0012},
	cmdLutElement{0x01, 0x03, 0, 0x03, 0x0006, 0x0016},
	cmdLutElement{0x01, 0x03, 0, 0x03, 0x0006, 0x001e},
	cmdLutElement{0x01, 0x04, 0, 0x03, 0x0006, 0x0026},
	cmdLutElement{0x01, 0x04, 0, 0x03, 0x0006, 0x0036},
	cmdLutElement{0x01, 0x01, 0, 0x03, 0x0008, 0x000a},
	cmdLutElement{0x01, 0x01, 0, 0x03, 0x0008, 0x000c},
	cmdLutElement{0x01, 0x02, 0, 0x03, 0x0008, 0x000e},
	cmdLutElement{0x01, 0x02, 0, 0x03, 0x0008, 0x0012},
	cmdLutElement{0x01, 0x03, 0, 0x03, 0x0008, 0x0016},
	cmdLutElement{0x01, 0x03, 0, 0x03, 0x0008, 0x001e},
	cmdLutElement{0x01, 0x04, 0, 0x03, 0x0008, 0x0026},
	cmdLutElement{0x01, 0x04, 0, 0x03, 0x0008, 0x0036},
	cmdLutElement{0x00, 0x00, -1, 0x00, 0x0000, 0x0002},
	cmdLutElement{0x00, 0x00, -1, 0x01, 0x0000, 0x0003},
	cmdLutElement{0x00, 0x00, -1, 0x02, 0x0000, 0x0004},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0000, 0x0005},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0000, 0x0006},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0000, 0x0007},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0000, 0x0008},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0000, 0x0009},
	cmdLutElement{0x00, 0x00, -1, 0x00, 0x0001, 0x0002},
	cmdLutElement{0x00, 0x00, -1, 0x01, 0x0001, 0x0003},
	cmdLutElement{0x00, 0x00, -1, 0x02, 0x0001, 0x0004},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0001, 0x0005},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0001, 0x0006},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0001, 0x0007},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0001, 0x0008},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0001, 0x0009},
	cmdLutElement{0x00, 0x00, -1, 0x00, 0x0002, 0x0002},
	cmdLutElement{0x00, 0x00, -1, 0x01, 0x0002, 0x0003},
	cmdLutElement{0x00, 0x00, -1, 0x02, 0x0002, 0x0004},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0002, 0x0005},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0002, 0x0006},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0002, 0x0007},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0002, 0x0008},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0002, 0x0009},
	cmdLutElement{0x00, 0x00, -1, 0x00, 0x0003, 0x0002},
	cmdLutElement{0x00, 0x00, -1, 0x01, 0x0003, 0x0003},
	cmdLutElement{0x00, 0x00, -1, 0x02, 0x0003, 0x0004},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0003, 0x0005},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0003, 0x0006},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0003, 0x0007},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0003, 0x0008},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0003, 0x0009},
	cmdLutElement{0x00, 0x00, -1, 0x00, 0x0004, 0x0002},
	cmdLutElement{0x00, 0x00, -1, 0x01, 0x0004, 0x0003},
	cmdLutElement{0x00, 0x00, -1, 0x02, 0x0004, 0x0004},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0004, 0x0005},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0004, 0x0006},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0004, 0x0007},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0004, 0x0008},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0004, 0x0009},
	cmdLutElement{0x00, 0x00, -1, 0x00, 0x0005, 0x0002},
	cmdLutElement{0x00, 0x00, -1, 0x01, 0x0005, 0x0003},
	cmdLutElement{0x00, 0x00, -1, 0x02, 0x0005, 0x0004},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0005, 0x0005},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0005, 0x0006},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0005, 0x0007},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0005, 0x0008},
	cmdLutElement{0x00, 0x00, -1, 0x03, 0x0005, 0x0009},
	cmdLutElement{0x01, 0x00, -1, 0x00, 0x0006, 0x0002},
	cmdLutElement{0x01, 0x00, -1, 0x01, 0x0006, 0x0003},
	cmdLutElement{0x01, 0x00, -1, 0x02, 0x0006, 0x0004},
	cmdLutElement{0x01, 0x00, -1, 0x03, 0x0006, 0x0005},
	cmdLutElement{0x01, 0x00, -1, 0x03, 0x0006, 0x0006},
	cmdLutElement{0x01, 0x00, -1, 0x03, 0x0006, 0x0007},
	cmdLutElement{0x01, 0x00, -1, 0x03, 0x0006, 0x0008},
	cmdLutElement{0x01, 0x00, -1, 0x03, 0x0006, 0x0009},
	cmdLutElement{0x01, 0x00, -1, 0x00, 0x0008, 0x0002},
	cmdLutElement{0x01, 0x00, -1, 0x01, 0x0008, 0x0003},
	cmdLutElement{0x01, 0x00, -1, 0x02, 0x0008, 0x0004},
	cmdLutElement{0x01, 0x00, -1, 0x03, 0x0008, 0x0005},
	cmdLutElement{0x01, 0x00, -1, 0x03, 0x0008, 0x0006},
	cmdLutElement{0x01, 0x00, -1, 0x03, 0x0008, 0x0007},
	cmdLutElement{0x01, 0x00, -1, 0x03, 0x0008, 0x0008},
	cmdLutElement{0x01, 0x00, -1, 0x03, 0x0008, 0x0009},
	cmdLutElement{0x00, 0x01, -1, 0x03, 0x0000, 0x000a},
	cmdLutElement{0x00, 0x01, -1, 0x03, 0x0000, 0x000c},
	cmdLutElement{0x00, 0x02, -1, 0x03, 0x0000, 0x000e},
	cmdLutElement{0x00, 0x02, -1, 0x03, 0x0000, 0x0012},
	cmdLutElement{0x00, 0x03, -1, 0x03, 0x0000, 0x0016},
	cmdLutElement{0x00, 0x03, -1, 0x03, 0x0000, 0x001e},
	cmdLutElement{0x00, 0x04, -1, 0x03, 0x0000, 0x0026},
	cmdLutElement{0x00, 0x04, -1, 0x03, 0x0000, 0x0036},
	cmdLutElement{0x00, 0x01, -1, 0x03, 0x0001, 0x000a},
	cmdLutElement{0x00, 0x01, -1, 0x03, 0x0001, 0x000c},
	cmdLutElement{0x00, 0x02, -1, 0x03, 0x0001, 0x000e},
	cmdLutElement{0x00, 0x02, -1, 0x03, 0x0001, 0x0012},
	cmdLutElement{0x00, 0x03, -1, 0x03, 0x0001, 0x0016},
	cmdLutElement{0x00, 0x03, -1, 0x03, 0x0001, 0x001e},
	cmdLutElement{0x00, 0x04, -1, 0x03, 0x0001, 0x0026},
	cmdLutElement{0x00, 0x04, -1, 0x03, 0x0001, 0x0036},
	cmdLutElement{0x00, 0x01, -1, 0x03, 0x0002, 0x000a},
	cmdLutElement{0x00, 0x01, -1, 0x03, 0x0002, 0x000c},
	cmdLutElement{0x00, 0x02, -1, 0x03, 0x0002, 0x000e},
	cmdLutElement{0x00, 0x02, -1, 0x03, 0x0002, 0x0012},
	cmdLutElement{0x00, 0x03, -1, 0x03, 0x0002, 0x0016},
	cmdLutElement{0x00, 0x03, -1, 0x03, 0x0002, 0x001e},
	cmdLutElement{0x00, 0x04, -1, 0x03, 0x0002, 0x0026},
	cmdLutElement{0x00, 0x04, -1, 0x03, 0x0002, 0x0036},
	cmdLutElement{0x00, 0x01, -1, 0x03, 0x0003, 0x000a},
	cmdLutElement{0x00, 0x01, -1, 0x03, 0x0003, 0x000c},
	cmdLutElement{0x00, 0x02, -1, 0x03, 0x0003, 0x000e},
	cmdLutElement{0x00, 0x02, -1, 0x03, 0x0003, 0x0012},
	cmdLutElement{0x00, 0x03, -1, 0x03, 0x0003, 0x0016},
	cmdLutElement{0x00, 0x03, -1, 0x03, 0x0003, 0x001e},
	cmdLutElement{0x00, 0x04, -1, 0x03, 0x0003, 0x0026},
	cmdLutElement{0x00, 0x04, -1, 0x03, 0x0003, 0x0036},
	cmdLutElement{0x00, 0x01, -1, 0x03, 0x0004, 0x000a},
	cmdLutElement{0x00, 0x01, -1, 0x03, 0x0004, 0x000c},
	cmdLutElement{0x00, 0x02, -1, 0x03, 0x0004, 0x000e},
	cmdLutElement{0x00, 0x02, -1, 0x03, 0x0004, 0x0012},
	cmdLutElement{0x00, 0x03, -1, 0x03, 0x0004, 0x0016},
	cmdLutElement{0x00, 0x03, -1, 0x03, 0x0004, 0x001e},
	cmdLutElement{0x00, 0x04, -1, 0x03, 0x0004, 0x0026},
	cmdLutElement{0x00, 0x04, -1, 0x03, 0x0004, 0x0036},
	cmdLutElement{0x00, 0x01, -1, 0x03, 0x0005, 0x000a},
	cmdLutElement{0x00, 0x01, -1, 0x03, 0x0005, 0x000c},
	cmdLutElement{0x00, 0x02, -1, 0x03, 0x0005, 0x000e},
	cmdLutElement{0x00, 0x02, -1, 0x03, 0x0005, 0x0012},
	cmdLutElement{0x00, 0x03, -1, 0x03, 0x0005, 0x0016},
	cmdLutElement{0x00, 0x03, -1, 0x03, 0x0005, 0x001e},
	cmdLutElement{0x00, 0x04, -1, 0x03, 0x0005, 0x0026},
	cmdLutElement{0x00, 0x04, -1, 0x03, 0x0005, 0x0036},
	cmdLutElement{0x01, 0x01, -1, 0x03, 0x0006, 0x000a},
	cmdLutElement{0x01, 0x01, -1, 0x03, 0x0006, 0x000c},
	cmdLutElement{0x01, 0x02, -1, 0x03, 0x0006, 0x000e},
	cmdLutElement{0x01, 0x02, -1, 0x03, 0x0006, 0x0012},
	cmdLutElement{0x01, 0x03, -1, 0x03, 0x0006, 0x0016},
	cmdLutElement{0x01, 0x03, -1, 0x03, 0x0006, 0x001e},
	cmdLutElement{0x01, 0x04, -1, 0x03, 0x0006, 0x0026},
	cmdLutElement{0x01, 0x04, -1, 0x03, 0x0006, 0x0036},
	cmdLutElement{0x01, 0x01, -1, 0x03, 0x0008, 0x000a},
	cmdLutElement{0x01, 0x01, -1, 0x03, 0x0008, 0x000c},
	cmdLutElement{0x01, 0x02, -1, 0x03, 0x0008, 0x000e},
	cmdLutElement{0x01, 0x02, -1, 0x03, 0x0008, 0x0012},
	cmdLutElement{0x01, 0x03, -1, 0x03, 0x0008, 0x0016},
	cmdLutElement{0x01, 0x03, -1, 0x03, 0x0008, 0x001e},
	cmdLutElement{0x01, 0x04, -1, 0x03, 0x0008, 0x0026},
	cmdLutElement{0x01, 0x04, -1, 0x03, 0x0008, 0x0036},
	cmdLutElement{0x02, 0x00, -1, 0x00, 0x000a, 0x0002},
	cmdLutElement{0x02, 0x00, -1, 0x01, 0x000a, 0x0003},
	cmdLutElement{0x02, 0x00, -1, 0x02, 0x000a, 0x0004},
	cmdLutElement{0x02, 0x00, -1, 0x03, 0x000a, 0x0005},
	cmdLutElement{0x02, 0x00, -1, 0x03, 0x000a, 0x0006},
	cmdLutElement{0x02, 0x00, -1, 0x03, 0x000a, 0x0007},
	cmdLutElement{0x02, 0x00, -1, 0x03, 0x000a, 0x0008},
	cmdLutElement{0x02, 0x00, -1, 0x03, 0x000a, 0x0009},
	cmdLutElement{0x02, 0x00, -1, 0x00, 0x000e, 0x0002},
	cmdLutElement{0x02, 0x00, -1, 0x01, 0x000e, 0x0003},
	cmdLutElement{0x02, 0x00, -1, 0x02, 0x000e, 0x0004},
	cmdLutElement{0x02, 0x00, -1, 0x03, 0x000e, 0x0005},
	cmdLutElement{0x02, 0x00, -1, 0x03, 0x000e, 0x0006},
	cmdLutElement{0x02, 0x00, -1, 0x03, 0x000e, 0x0007},
	cmdLutElement{0x02, 0x00, -1, 0x03, 0x000e, 0x0008},
	cmdLutElement{0x02, 0x00, -1, 0x03, 0x000e, 0x0009},
	cmdLutElement{0x03, 0x00, -1, 0x00, 0x0012, 0x0002},
	cmdLutElement{0x03, 0x00, -1, 0x01, 0x0012, 0x0003},
	cmdLutElement{0x03, 0x00, -1, 0x02, 0x0012, 0x0004},
	cmdLutElement{0x03, 0x00, -1, 0x03, 0x0012, 0x0005},
	cmdLutElement{0x03, 0x00, -1, 0x03, 0x0012, 0x0006},
	cmdLutElement{0x03, 0x00, -1, 0x03, 0x0012, 0x0007},
	cmdLutElement{0x03, 0x00, -1, 0x03, 0x0012, 0x0008},
	cmdLutElement{0x03, 0x00, -1, 0x03, 0x0012, 0x0009},
	cmdLutElement{0x03, 0x00, -1, 0x00, 0x001a, 0x0002},
	cmdLutElement{0x03, 0x00, -1, 0x01, 0x001a, 0x0003},
	cmdLutElement{0x03, 0x00, -1, 0x02, 0x001a, 0x0004},
	cmdLutElement{0x03, 0x00, -1, 0x03, 0x001a, 0x0005},
	cmdLutElement{0x03, 0x00, -1, 0x03, 0x001a, 0x0006},
	cmdLutElement{0x03, 0x00, -1, 0x03, 0x001a, 0x0007},
	cmdLutElement{0x03, 0x00, -1, 0x03, 0x001a, 0x0008},
	cmdLutElement{0x03, 0x00, -1, 0x03, 0x001a, 0x0009},
	cmdLutElement{0x04, 0x00, -1, 0x00, 0x0022, 0x0002},
	cmdLutElement{0x04, 0x00, -1, 0x01, 0x0022, 0x0003},
	cmdLutElement{0x04, 0x00, -1, 0x02, 0x0022, 0x0004},
	cmdLutElement{0x04, 0x00, -1, 0x03, 0x0022, 0x0005},
	cmdLutElement{0x04, 0x00, -1, 0x03, 0x0022, 0x0006},
	cmdLutElement{0x04, 0x00, -1, 0x03, 0x0022, 0x0007},
	cmdLutElement{0x04, 0x00, -1, 0x03, 0x0022, 0x0008},
	cmdLutElement{0x04, 0x00, -1, 0x03, 0x0022, 0x0009},
	cmdLutElement{0x04, 0x00, -1, 0x00, 0x0032, 0x0002},
	cmdLutElement{0x04, 0x00, -1, 0x01, 0x0032, 0x0003},
	cmdLutElement{0x04, 0x00, -1, 0x02, 0x0032, 0x0004},
	cmdLutElement{0x04, 0x00, -1, 0x03, 0x0032, 0x0005},
	cmdLutElement{0x04, 0x00, -1, 0x03, 0x0032, 0x0006},
	cmdLutElement{0x04, 0x00, -1, 0x03, 0x0032, 0x0007},
	cmdLutElement{0x04, 0x00, -1, 0x03, 0x0032, 0x0008},
	cmdLutElement{0x04, 0x00, -1, 0x03, 0x0032, 0x0009},
	cmdLutElement{0x05, 0x00, -1, 0x00, 0x0042, 0x0002},
	cmdLutElement{0x05, 0x00, -1, 0x01, 0x0042, 0x0003},
	cmdLutElement{0x05, 0x00, -1, 0x02, 0x0042, 0x0004},
	cmdLutElement{0x05, 0x00, -1, 0x03, 0x0042, 0x0005},
	cmdLutElement{0x05, 0x00, -1, 0x03, 0x0042, 0x0006},
	cmdLutElement{0x05, 0x00, -1, 0x03, 0x0042, 0x0007},
	cmdLutElement{0x05, 0x00, -1, 0x03, 0x0042, 0x0008},
	cmdLutElement{0x05, 0x00, -1, 0x03, 0x0042, 0x0009},
	cmdLutElement{0x05, 0x00, -1, 0x00, 0x0062, 0x0002},
	cmdLutElement{0x05, 0x00, -1, 0x01, 0x0062, 0x0003},
	cmdLutElement{0x05, 0x00, -1, 0x02, 0x0062, 0x0004},
	cmdLutElement{0x05, 0x00, -1, 0x03, 0x0062, 0x0005},
	cmdLutElement{0x05, 0x00, -1, 0x03, 0x0062, 0x0006},
	cmdLutElement{0x05, 0x00, -1, 0x03, 0x0062, 0x0007},
	cmdLutElement{0x05, 0x00, -1, 0x03, 0x0062, 0x0008},
	cmdLutElement{0x05, 0x00, -1, 0x03, 0x0062, 0x0009},
	cmdLutElement{0x02, 0x01, -1, 0x03, 0x000a, 0x000a},
	cmdLutElement{0x02, 0x01, -1, 0x03, 0x000a, 0x000c},
	cmdLutElement{0x02, 0x02, -1, 0x03, 0x000a, 0x000e},
	cmdLutElement{0x02, 0x02, -1, 0x03, 0x000a, 0x0012},
	cmdLutElement{0x02, 0x03, -1, 0x03, 0x000a, 0x0016},
	cmdLutElement{0x02, 0x03, -1, 0x03, 0x000a, 0x001e},
	cmdLutElement{0x02, 0x04, -1, 0x03, 0x000a, 0x0026},
	cmdLutElement{0x02, 0x04, -1, 0x03, 0x000a, 0x0036},
	cmdLutElement{0x02, 0x01, -1, 0x03, 0x000e, 0x000a},
	cmdLutElement{0x02, 0x01, -1, 0x03, 0x000e, 0x000c},
	cmdLutElement{0x02, 0x02, -1, 0x03, 0x000e, 0x000e},
	cmdLutElement{0x02, 0x02, -1, 0x03, 0x000e, 0x0012},
	cmdLutElement{0x02, 0x03, -1, 0x03, 0x000e, 0x0016},
	cmdLutElement{0x02, 0x03, -1, 0x03, 0x000e, 0x001e},
	cmdLutElement{0x02, 0x04, -1, 0x03, 0x000e, 0x0026},
	cmdLutElement{0x02, 0x04, -1, 0x03, 0x000e, 0x0036},
	cmdLutElement{0x03, 0x01, -1, 0x03, 0x0012, 0x000a},
	cmdLutElement{0x03, 0x01, -1, 0x03, 0x0012, 0x000c},
	cmdLutElement{0x03, 0x02, -1, 0x03, 0x0012, 0x000e},
	cmdLutElement{0x03, 0x02, -1, 0x03, 0x0012, 0x0012},
	cmdLutElement{0x03, 0x03, -1, 0x03, 0x0012, 0x0016},
	cmdLutElement{0x03, 0x03, -1, 0x03, 0x0012, 0x001e},
	cmdLutElement{0x03, 0x04, -1, 0x03, 0x0012, 0x0026},
	cmdLutElement{0x03, 0x04, -1, 0x03, 0x0012, 0x0036},
	cmdLutElement{0x03, 0x01, -1, 0x03, 0x001a, 0x000a},
	cmdLutElement{0x03, 0x01, -1, 0x03, 0x001a, 0x000c},
	cmdLutElement{0x03, 0x02, -1, 0x03, 0x001a, 0x000e},
	cmdLutElement{0x03, 0x02, -1, 0x03, 0x001a, 0x0012},
	cmdLutElement{0x03, 0x03, -1, 0x03, 0x001a, 0x0016},
	cmdLutElement{0x03, 0x03, -1, 0x03, 0x001a, 0x001e},
	cmdLutElement{0x03, 0x04, -1, 0x03, 0x001a, 0x0026},
	cmdLutElement{0x03, 0x04, -1, 0x03, 0x001a, 0x0036},
	cmdLutElement{0x04, 0x01, -1, 0x03, 0x0022, 0x000a},
	cmdLutElement{0x04, 0x01, -1, 0x03, 0x0022, 0x000c},
	cmdLutElement{0x04, 0x02, -1, 0x03, 0x0022, 0x000e},
	cmdLutElement{0x04, 0x02, -1, 0x03, 0x0022, 0x0012},
	cmdLutElement{0x04, 0x03, -1, 0x03, 0x0022, 0x0016},
	cmdLutElement{0x04, 0x03, -1, 0x03, 0x0022, 0x001e},
	cmdLutElement{0x04, 0x04, -1, 0x03, 0x0022, 0x0026},
	cmdLutElement{0x04, 0x04, -1, 0x03, 0x0022, 0x0036},
	cmdLutElement{0x04, 0x01, -1, 0x03, 0x0032, 0x000a},
	cmdLutElement{0x04, 0x01, -1, 0x03, 0x0032, 0x000c},
	cmdLutElement{0x04, 0x02, -1, 0x03, 0x0032, 0x000e},
	cmdLutElement{0x04, 0x02, -1, 0x03, 0x0032, 0x0012},
	cmdLutElement{0x04, 0x03, -1, 0x03, 0x0032, 0x0016},
	cmdLutElement{0x04, 0x03, -1, 0x03, 0x0032, 0x001e},
	cmdLutElement{0x04, 0x04, -1, 0x03, 0x0032, 0x0026},
	cmdLutElement{0x04, 0x04, -1, 0x03, 0x0032, 0x0036},
	cmdLutElement{0x05, 0x01, -1, 0x03, 0x0042, 0x000a},
	cmdLutElement{0x05, 0x01, -1, 0x03, 0x0042, 0x000c},
	cmdLutElement{0x05, 0x02, -1, 0x03, 0x0042, 0x000e},
	cmdLutElement{0x05, 0x02, -1, 0x03, 0x0042, 0x0012},
	cmdLutElement{0x05, 0x03, -1, 0x03, 0x0042, 0x0016},
	cmdLutElement{0x05, 0x03, -1, 0x03, 0x0042, 0x001e},
	cmdLutElement{0x05, 0x04, -1, 0x03, 0x0042, 0x0026},
	cmdLutElement{0x05, 0x04, -1, 0x03, 0x0042, 0x0036},
	cmdLutElement{0x05, 0x01, -1, 0x03, 0x0062, 0x000a},
	cmdLutElement{0x05, 0x01, -1, 0x03, 0x0062, 0x000c},
	cmdLutElement{0x05, 0x02, -1, 0x03, 0x0062, 0x000e},
	cmdLutElement{0x05, 0x02, -1, 0x03, 0x0062, 0x0012},
	cmdLutElement{0x05, 0x03, -1, 0x03, 0x0062, 0x0016},
	cmdLutElement{0x05, 0x03, -1, 0x03, 0x0062, 0x001e},
	cmdLutElement{0x05, 0x04, -1, 0x03, 0x0062, 0x0026},
	cmdLutElement{0x05, 0x04, -1, 0x03, 0x0062, 0x0036},
	cmdLutElement{0x00, 0x05, -1, 0x03, 0x0000, 0x0046},
	cmdLutElement{0x00, 0x05, -1, 0x03, 0x0000, 0x0066},
	cmdLutElement{0x00, 0x06, -1, 0x03, 0x0000, 0x0086},
	cmdLutElement{0x00, 0x07, -1, 0x03, 0x0000, 0x00c6},
	cmdLutElement{0x00, 0x08, -1, 0x03, 0x0000, 0x0146},
	cmdLutElement{0x00, 0x09, -1, 0x03, 0x0000, 0x0246},
	cmdLutElement{0x00, 0x0a, -1, 0x03, 0x0000, 0x0446},
	cmdLutElement{0x00, 0x18, -1, 0x03, 0x0000, 0x0846},
	cmdLutElement{0x00, 0x05, -1, 0x03, 0x0001, 0x0046},
	cmdLutElement{0x00, 0x05, -1, 0x03, 0x0001, 0x0066},
	cmdLutElement{0x00, 0x06, -1, 0x03, 0x0001, 0x0086},
	cmdLutElement{0x00, 0x07, -1, 0x03, 0x0001, 0x00c6},
	cmdLutElement{0x00, 0x08, -1, 0x03, 0x0001, 0x0146},
	cmdLutElement{0x00, 0x09, -1, 0x03, 0x0001, 0x0246},
	cmdLutElement{0x00, 0x0a, -1, 0x03, 0x0001, 0x0446},
	cmdLutElement{0x00, 0x18, -1, 0x03, 0x0001, 0x0846},
	cmdLutElement{0x00, 0x05, -1, 0x03, 0x0002, 0x0046},
	cmdLutElement{0x00, 0x05, -1, 0x03, 0x0002, 0x0066},
	cmdLutElement{0x00, 0x06, -1, 0x03, 0x0002, 0x0086},
	cmdLutElement{0x00, 0x07, -1, 0x03, 0x0002, 0x00c6},
	cmdLutElement{0x00, 0x08, -1, 0x03, 0x0002, 0x0146},
	cmdLutElement{0x00, 0x09, -1, 0x03, 0x0002, 0x0246},
	cmdLutElement{0x00, 0x0a, -1, 0x03, 0x0002, 0x0446},
	cmdLutElement{0x00, 0x18, -1, 0x03, 0x0002, 0x0846},
	cmdLutElement{0x00, 0x05, -1, 0x03, 0x0003, 0x0046},
	cmdLutElement{0x00, 0x05, -1, 0x03, 0x0003, 0x0066},
	cmdLutElement{0x00, 0x06, -1, 0x03, 0x0003, 0x0086},
	cmdLutElement{0x00, 0x07, -1, 0x03, 0x0003, 0x00c6},
	cmdLutElement{0x00, 0x08, -1, 0x03, 0x0003, 0x0146},
	cmdLutElement{0x00, 0x09, -1, 0x03, 0x0003, 0x0246},
	cmdLutElement{0x00, 0x0a, -1, 0x03, 0x0003, 0x0446},
	cmdLutElement{0x00, 0x18, -1, 0x03, 0x0003, 0x0846},
	cmdLutElement{0x00, 0x05, -1, 0x03, 0x0004, 0x0046},
	cmdLutElement{0x00, 0x05, -1, 0x03, 0x0004, 0x0066},
	cmdLutElement{0x00, 0x06, -1, 0x03, 0x0004, 0x0086},
	cmdLutElement{0x00, 0x07, -1, 0x03, 0x0004, 0x00c6},
	cmdLutElement{0x00, 0x08, -1, 0x03, 0x0004, 0x0146},
	cmdLutElement{0x00, 0x09, -1, 0x03, 0x0004, 0x0246},
	cmdLutElement{0x00, 0x0a, -1, 0x03, 0x0004, 0x0446},
	cmdLutElement{0x00, 0x18, -1, 0x03, 0x0004, 0x0846},
	cmdLutElement{0x00, 0x05, -1, 0x03, 0x0005, 0x0046},
	cmdLutElement{0x00, 0x05, -1, 0x03, 0x0005, 0x0066},
	cmdLutElement{0x00, 0x06, -1, 0x03, 0x0005, 0x0086},
	cmdLutElement{0x00, 0x07, -1, 0x03, 0x0005, 0x00c6},
	cmdLutElement{0x00, 0x08, -1, 0x03, 0x0005, 0x0146},
	cmdLutElement{0x00, 0x09, -1, 0x03, 0x0005, 0x0246},
	cmdLutElement{0x00, 0x0a, -1, 0x03, 0x0005, 0x0446},
	cmdLutElement{0x00, 0x18, -1, 0x03, 0x0005, 0x0846},
	cmdLutElement{0x01, 0x05, -1, 0x03, 0x0006, 0x0046},
	cmdLutElement{0x01, 0x05, -1, 0x03, 0x0006, 0x0066},
	cmdLutElement{0x01, 0x06, -1, 0x03, 0x0006, 0x0086},
	cmdLutElement{0x01, 0x07, -1, 0x03, 0x0006, 0x00c6},
	cmdLutElement{0x01, 0x08, -1, 0x03, 0x0006, 0x0146},
	cmdLutElement{0x01, 0x09, -1, 0x03, 0x0006, 0x0246},
	cmdLutElement{0x01, 0x0a, -1, 0x03, 0x0006, 0x0446},
	cmdLutElement{0x01, 0x18, -1, 0x03, 0x0006, 0x0846},
	cmdLutElement{0x01, 0x05, -1, 0x03, 0x0008, 0x0046},
	cmdLutElement{0x01, 0x05, -1, 0x03, 0x0008, 0x0066},
	cmdLutElement{0x01, 0x06, -1, 0x03, 0x0008, 0x0086},
	cmdLutElement{0x01, 0x07, -1, 0x03, 0x0008, 0x00c6},
	cmdLutElement{0x01, 0x08, -1, 0x03, 0x0008, 0x0146},
	cmdLutElement{0x01, 0x09, -1, 0x03, 0x0008, 0x0246},
	cmdLutElement{0x01, 0x0a, -1, 0x03, 0x0008, 0x0446},
	cmdLutElement{0x01, 0x18, -1, 0x03, 0x0008, 0x0846},
	cmdLutElement{0x06, 0x00, -1, 0x00, 0x0082, 0x0002},
	cmdLutElement{0x06, 0x00, -1, 0x01, 0x0082, 0x0003},
	cmdLutElement{0x06, 0x00, -1, 0x02, 0x0082, 0x0004},
	cmdLutElement{0x06, 0x00, -1, 0x03, 0x0082, 0x0005},
	cmdLutElement{0x06, 0x00, -1, 0x03, 0x0082, 0x0006},
	cmdLutElement{0x06, 0x00, -1, 0x03, 0x0082, 0x0007},
	cmdLutElement{0x06, 0x00, -1, 0x03, 0x0082, 0x0008},
	cmdLutElement{0x06, 0x00, -1, 0x03, 0x0082, 0x0009},
	cmdLutElement{0x07, 0x00, -1, 0x00, 0x00c2, 0x0002},
	cmdLutElement{0x07, 0x00, -1, 0x01, 0x00c2, 0x0003},
	cmdLutElement{0x07, 0x00, -1, 0x02, 0x00c2, 0x0004},
	cmdLutElement{0x07, 0x00, -1, 0x03, 0x00c2, 0x0005},
	cmdLutElement{0x07, 0x00, -1, 0x03, 0x00c2, 0x0006},
	cmdLutElement{0x07, 0x00, -1, 0x03, 0x00c2, 0x0007},
	cmdLutElement{0x07, 0x00, -1, 0x03, 0x00c2, 0x0008},
	cmdLutElement{0x07, 0x00, -1, 0x03, 0x00c2, 0x0009},
	cmdLutElement{0x08, 0x00, -1, 0x00, 0x0142, 0x0002},
	cmdLutElement{0x08, 0x00, -1, 0x01, 0x0142, 0x0003},
	cmdLutElement{0x08, 0x00, -1, 0x02, 0x0142, 0x0004},
	cmdLutElement{0x08, 0x00, -1, 0x03, 0x0142, 0x0005},
	cmdLutElement{0x08, 0x00, -1, 0x03, 0x0142, 0x0006},
	cmdLutElement{0x08, 0x00, -1, 0x03, 0x0142, 0x0007},
	cmdLutElement{0x08, 0x00, -1, 0x03, 0x0142, 0x0008},
	cmdLutElement{0x08, 0x00, -1, 0x03, 0x0142, 0x0009},
	cmdLutElement{0x09, 0x00, -1, 0x00, 0x0242, 0x0002},
	cmdLutElement{0x09, 0x00, -1, 0x01, 0x0242, 0x0003},
	cmdLutElement{0x09, 0x00, -1, 0x02, 0x0242, 0x0004},
	cmdLutElement{0x09, 0x00, -1, 0x03, 0x0242, 0x0005},
	cmdLutElement{0x09, 0x00, -1, 0x03, 0x0242, 0x0006},
	cmdLutElement{0x09, 0x00, -1, 0x03, 0x0242, 0x0007},
	cmdLutElement{0x09, 0x00, -1, 0x03, 0x0242, 0x0008},
	cmdLutElement{0x09, 0x00, -1, 0x03, 0x0242, 0x0009},
	cmdLutElement{0x0a, 0x00, -1, 0x00, 0x0442, 0x0002},
	cmdLutElement{0x0a, 0x00, -1, 0x01, 0x0442, 0x0003},
	cmdLutElement{0x0a, 0x00, -1, 0x02, 0x0442, 0x0004},
	cmdLutElement{0x0a, 0x00, -1, 0x03, 0x0442, 0x0005},
	cmdLutElement{0x0a, 0x00, -1, 0x03, 0x0442, 0x0006},
	cmdLutElement{0x0a, 0x00, -1, 0x03, 0x0442, 0x0007},
	cmdLutElement{0x0a, 0x00, -1, 0x03, 0x0442, 0x0008},
	cmdLutElement{0x0a, 0x00, -1, 0x03, 0x0442, 0x0009},
	cmdLutElement{0x0c, 0x00, -1, 0x00, 0x0842, 0x0002},
	cmdLutElement{0x0c, 0x00, -1, 0x01, 0x0842, 0x0003},
	cmdLutElement{0x0c, 0x00, -1, 0x02, 0x0842, 0x0004},
	cmdLutElement{0x0c, 0x00, -1, 0x03, 0x0842, 0x0005},
	cmdLutElement{0x0c, 0x00, -1, 0x03, 0x0842, 0x0006},
	cmdLutElement{0x0c, 0x00, -1, 0x03, 0x0842, 0x0007},
	cmdLutElement{0x0c, 0x00, -1, 0x03, 0x0842, 0x0008},
	cmdLutElement{0x0c, 0x00, -1, 0x03, 0x0842, 0x0009},
	cmdLutElement{0x0e, 0x00, -1, 0x00, 0x1842, 0x0002},
	cmdLutElement{0x0e, 0x00, -1, 0x01, 0x1842, 0x0003},
	cmdLutElement{0x0e, 0x00, -1, 0x02, 0x1842, 0x0004},
	cmdLutElement{0x0e, 0x00, -1, 0x03, 0x1842, 0x0005},
	cmdLutElement{0x0e, 0x00, -1, 0x03, 0x1842, 0x0006},
	cmdLutElement{0x0e, 0x00, -1, 0x03, 0x1842, 0x0007},
	cmdLutElement{0x0e, 0x00, -1, 0x03, 0x1842, 0x0008},
	cmdLutElement{0x0e, 0x00, -1, 0x03, 0x1842, 0x0009},
	cmdLutElement{0x18, 0x00, -1, 0x00, 0x5842, 0x0002},
	cmdLutElement{0x18, 0x00, -1, 0x01, 0x5842, 0x0003},
	cmdLutElement{0x18, 0x00, -1, 0x02, 0x5842, 0x0004},
	cmdLutElement{0x18, 0x00, -1, 0x03, 0x5842, 0x0005},
	cmdLutElement{0x18, 0x00, -1, 0x03, 0x5842, 0x0006},
	cmdLutElement{0x18, 0x00, -1, 0x03, 0x5842, 0x0007},
	cmdLutElement{0x18, 0x00, -1, 0x03, 0x5842, 0x0008},
	cmdLutElement{0x18, 0x00, -1, 0x03, 0x5842, 0x0009},
	cmdLutElement{0x02, 0x05, -1, 0x03, 0x000a, 0x0046},
	cmdLutElement{0x02, 0x05, -1, 0x03, 0x000a, 0x0066},
	cmdLutElement{0x02, 0x06, -1, 0x03, 0x000a, 0x0086},
	cmdLutElement{0x02, 0x07, -1, 0x03, 0x000a, 0x00c6},
	cmdLutElement{0x02, 0x08, -1, 0x03, 0x000a, 0x0146},
	cmdLutElement{0x02, 0x09, -1, 0x03, 0x000a, 0x0246},
	cmdLutElement{0x02, 0x0a, -1, 0x03, 0x000a, 0x0446},
	cmdLutElement{0x02, 0x18, -1, 0x03, 0x000a, 0x0846},
	cmdLutElement{0x02, 0x05, -1, 0x03, 0x000e, 0x0046},
	cmdLutElement{0x02, 0x05, -1, 0x03, 0x000e, 0x0066},
	cmdLutElement{0x02, 0x06, -1, 0x03, 0x000e, 0x0086},
	cmdLutElement{0x02, 0x07, -1, 0x03, 0x000e, 0x00c6},
	cmdLutElement{0x02, 0x08, -1, 0x03, 0x000e, 0x0146},
	cmdLutElement{0x02, 0x09, -1, 0x03, 0x000e, 0x0246},
	cmdLutElement{0x02, 0x0a, -1, 0x03, 0x000e, 0x0446},
	cmdLutElement{0x02, 0x18, -1, 0x03, 0x000e, 0x0846},
	cmdLutElement{0x03, 0x05, -1, 0x03, 0x0012, 0x0046},
	cmdLutElement{0x03, 0x05, -1, 0x03, 0x0012, 0x0066},
	cmdLutElement{0x03, 0x06, -1, 0x03, 0x0012, 0x0086},
	cmdLutElement{0x03, 0x07, -1, 0x03, 0x0012, 0x00c6},
	cmdLutElement{0x03, 0x08, -1, 0x03, 0x0012, 0x0146},
	cmdLutElement{0x03, 0x09, -1, 0x03, 0x0012, 0x0246},
	cmdLutElement{0x03, 0x0a, -1, 0x03, 0x0012, 0x0446},
	cmdLutElement{0x03, 0x18, -1, 0x03, 0x0012, 0x0846},
	cmdLutElement{0x03, 0x05, -1, 0x03, 0x001a, 0x0046},
	cmdLutElement{0x03, 0x05, -1, 0x03, 0x001a, 0x0066},
	cmdLutElement{0x03, 0x06, -1, 0x03, 0x001a, 0x0086},
	cmdLutElement{0x03, 0x07, -1, 0x03, 0x001a, 0x00c6},
	cmdLutElement{0x03, 0x08, -1, 0x03, 0x001a, 0x0146},
	cmdLutElement{0x03, 0x09, -1, 0x03, 0x001a, 0x0246},
	cmdLutElement{0x03, 0x0a, -1, 0x03, 0x001a, 0x0446},
	cmdLutElement{0x03, 0x18, -1, 0x03, 0x001a, 0x0846},
	cmdLutElement{0x04, 0x05, -1, 0x03, 0x0022, 0x0046},
	cmdLutElement{0x04, 0x05, -1, 0x03, 0x0022, 0x0066},
	cmdLutElement{0x04, 0x06, -1, 0x03, 0x0022, 0x0086},
	cmdLutElement{0x04, 0x07, -1, 0x03, 0x0022, 0x00c6},
	cmdLutElement{0x04, 0x08, -1, 0x03, 0x0022, 0x0146},
	cmdLutElement{0x04, 0x09, -1, 0x03, 0x0022, 0x0246},
	cmdLutElement{0x04, 0x0a, -1, 0x03, 0x0022, 0x0446},
	cmdLutElement{0x04, 0x18, -1, 0x03, 0x0022, 0x0846},
	cmdLutElement{0x04, 0x05, -1, 0x03, 0x0032, 0x0046},
	cmdLutElement{0x04, 0x05, -1, 0x03, 0x0032, 0x0066},
	cmdLutElement{0x04, 0x06, -1, 0x03, 0x0032, 0x0086},
	cmdLutElement{0x04, 0x07, -1, 0x03, 0x0032, 0x00c6},
	cmdLutElement{0x04, 0x08, -1, 0x03, 0x0032, 0x0146},
	cmdLutElement{0x04, 0x09, -1, 0x03, 0x0032, 0x0246},
	cmdLutElement{0x04, 0x0a, -1, 0x03, 0x0032, 0x0446},
	cmdLutElement{0x04, 0x18, -1, 0x03, 0x0032, 0x0846},
	cmdLutElement{0x05, 0x05, -1, 0x03, 0x0042, 0x0046},
	cmdLutElement{0x05, 0x05, -1, 0x03, 0x0042, 0x0066},
	cmdLutElement{0x05, 0x06, -1, 0x03, 0x0042, 0x0086},
	cmdLutElement{0x05, 0x07, -1, 0x03, 0x0042, 0x00c6},
	cmdLutElement{0x05, 0x08, -1, 0x03, 0x0042, 0x0146},
	cmdLutElement{0x05, 0x09, -1, 0x03, 0x0042, 0x0246},
	cmdLutElement{0x05, 0x0a, -1, 0x03, 0x0042, 0x0446},
	cmdLutElement{0x05, 0x18, -1, 0x03, 0x0042, 0x0846},
	cmdLutElement{0x05, 0x05, -1, 0x03, 0x0062, 0x0046},
	cmdLutElement{0x05, 0x05, -1, 0x03, 0x0062, 0x0066},
	cmdLutElement{0x05, 0x06, -1, 0x03, 0x0062, 0x0086},
	cmdLutElement{0x05, 0x07, -1, 0x03, 0x0062, 0x00c6},
	cmdLutElement{0x05, 0x08, -1, 0x03, 0x0062, 0x0146},
	cmdLutElement{0x05, 0x09, -1, 0x03, 0x0062, 0x0246},
	cmdLutElement{0x05, 0x0a, -1, 0x03, 0x0062, 0x0446},
	cmdLutElement{0x05, 0x18, -1, 0x03, 0x0062, 0x0846},
	cmdLutElement{0x06, 0x01, -1, 0x03, 0x0082, 0x000a},
	cmdLutElement{0x06, 0x01, -1, 0x03, 0x0082, 0x000c},
	cmdLutElement{0x06, 0x02, -1, 0x03, 0x0082, 0x000e},
	cmdLutElement{0x06, 0x02, -1, 0x03, 0x0082, 0x0012},
	cmdLutElement{0x06, 0x03, -1, 0x03, 0x0082, 0x0016},
	cmdLutElement{0x06, 0x03, -1, 0x03, 0x0082, 0x001e},
	cmdLutElement{0x06, 0x04, -1, 0x03, 0x0082, 0x0026},
	cmdLutElement{0x06, 0x04, -1, 0x03, 0x0082, 0x0036},
	cmdLutElement{0x07, 0x01, -1, 0x03, 0x00c2, 0x000a},
	cmdLutElement{0x07, 0x01, -1, 0x03, 0x00c2, 0x000c},
	cmdLutElement{0x07, 0x02, -1, 0x03, 0x00c2, 0x000e},
	cmdLutElement{0x07, 0x02, -1, 0x03, 0x00c2, 0x0012},
	cmdLutElement{0x07, 0x03, -1, 0x03, 0x00c2, 0x0016},
	cmdLutElement{0x07, 0x03, -1, 0x03, 0x00c2, 0x001e},
	cmdLutElement{0x07, 0x04, -1, 0x03, 0x00c2, 0x0026},
	cmdLutElement{0x07, 0x04, -1, 0x03, 0x00c2, 0x0036},
	cmdLutElement{0x08, 0x01, -1, 0x03, 0x0142, 0x000a},
	cmdLutElement{0x08, 0x01, -1, 0x03, 0x0142, 0x000c},
	cmdLutElement{0x08, 0x02, -1, 0x03, 0x0142, 0x000e},
	cmdLutElement{0x08, 0x02, -1, 0x03, 0x0142, 0x0012},
	cmdLutElement{0x08, 0x03, -1, 0x03, 0x0142, 0x0016},
	cmdLutElement{0x08, 0x03, -1, 0x03, 0x0142, 0x001e},
	cmdLutElement{0x08, 0x04, -1, 0x03, 0x0142, 0x0026},
	cmdLutElement{0x08, 0x04, -1, 0x03, 0x0142, 0x0036},
	cmdLutElement{0x09, 0x01, -1, 0x03, 0x0242, 0x000a},
	cmdLutElement{0x09, 0x01, -1, 0x03, 0x0242, 0x000c},
	cmdLutElement{0x09, 0x02, -1, 0x03, 0x0242, 0x000e},
	cmdLutElement{0x09, 0x02, -1, 0x03, 0x0242, 0x0012},
	cmdLutElement{0x09, 0x03, -1, 0x03, 0x0242, 0x0016},
	cmdLutElement{0x09, 0x03, -1, 0x03, 0x0242, 0x001e},
	cmdLutElement{0x09, 0x04, -1, 0x03, 0x0242, 0x0026},
	cmdLutElement{0x09, 0x04, -1, 0x03, 0x0242, 0x0036},
	cmdLutElement{0x0a, 0x01, -1, 0x03, 0x0442, 0x000a},
	cmdLutElement{0x0a, 0x01, -1, 0x03, 0x0442, 0x000c},
	cmdLutElement{0x0a, 0x02, -1, 0x03, 0x0442, 0x000e},
	cmdLutElement{0x0a, 0x02, -1, 0x03, 0x0442, 0x0012},
	cmdLutElement{0x0a, 0x03, -1, 0x03, 0x0442, 0x0016},
	cmdLutElement{0x0a, 0x03, -1, 0x03, 0x0442, 0x001e},
	cmdLutElement{0x0a, 0x04, -1, 0x03, 0x0442, 0x0026},
	cmdLutElement{0x0a, 0x04, -1, 0x03, 0x0442, 0x0036},
	cmdLutElement{0x0c, 0x01, -1, 0x03, 0x0842, 0x000a},
	cmdLutElement{0x0c, 0x01, -1, 0x03, 0x0842, 0x000c},
	cmdLutElement{0x0c, 0x02, -1, 0x03, 0x0842, 0x000e},
	cmdLutElement{0x0c, 0x02, -1, 0x03, 0x0842, 0x0012},
	cmdLutElement{0x0c, 0x03, -1, 0x03, 0x0842, 0x0016},
	cmdLutElement{0x0c, 0x03, -1, 0x03, 0x0842, 0x001e},
	cmdLutElement{0x0c, 0x04, -1, 0x03, 0x0842, 0x0026},
	cmdLutElement{0x0c, 0x04, -1, 0x03, 0x0842, 0x0036},
	cmdLutElement{0x0e, 0x01, -1, 0x03, 0x1842, 0x000a},
	cmdLutElement{0x0e, 0x01, -1, 0x03, 0x1842, 0x000c},
	cmdLutElement{0x0e, 0x02, -1, 0x03, 0x1842, 0x000e},
	cmdLutElement{0x0e, 0x02, -1, 0x03, 0x1842, 0x0012},
	cmdLutElement{0x0e, 0x03, -1, 0x03, 0x1842, 0x0016},
	cmdLutElement{0x0e, 0x03, -1, 0x03, 0x1842, 0x001e},
	cmdLutElement{0x0e, 0x04, -1, 0x03, 0x1842, 0x0026},
	cmdLutElement{0x0e, 0x04, -1, 0x03, 0x1842, 0x0036},
	cmdLutElement{0x18, 0x01, -1, 0x03, 0x5842, 0x000a},
	cmdLutElement{0x18, 0x01, -1, 0x03, 0x5842, 0x000c},
	cmdLutElement{0x18, 0x02, -1, 0x03, 0x5842, 0x000e},
	cmdLutElement{0x18, 0x02, -1, 0x03, 0x5842, 0x0012},
	cmdLutElement{0x18, 0x03, -1, 0x03, 0x5842, 0x0016},
	cmdLutElement{0x18, 0x03, -1, 0x03, 0x5842, 0x001e},
	cmdLutElement{0x18, 0x04, -1, 0x03, 0x5842, 0x0026},
	cmdLutElement{0x18, 0x04, -1, 0x03, 0x5842, 0x0036},
	cmdLutElement{0x06, 0x05, -1, 0x03, 0x0082, 0x0046},
	cmdLutElement{0x06, 0x05, -1, 0x03, 0x0082, 0x0066},
	cmdLutElement{0x06, 0x06, -1, 0x03, 0x0082, 0x0086},
	cmdLutElement{0x06, 0x07, -1, 0x03, 0x0082, 0x00c6},
	cmdLutElement{0x06, 0x08, -1, 0x03, 0x0082, 0x0146},
	cmdLutElement{0x06, 0x09, -1, 0x03, 0x0082, 0x0246},
	cmdLutElement{0x06, 0x0a, -1, 0x03, 0x0082, 0x0446},
	cmdLutElement{0x06, 0x18, -1, 0x03, 0x0082, 0x0846},
	cmdLutElement{0x07, 0x05, -1, 0x03, 0x00c2, 0x0046},
	cmdLutElement{0x07, 0x05, -1, 0x03, 0x00c2, 0x0066},
	cmdLutElement{0x07, 0x06, -1, 0x03, 0x00c2, 0x0086},
	cmdLutElement{0x07, 0x07, -1, 0x03, 0x00c2, 0x00c6},
	cmdLutElement{0x07, 0x08, -1, 0x03, 0x00c2, 0x0146},
	cmdLutElement{0x07, 0x09, -1, 0x03, 0x00c2, 0x0246},
	cmdLutElement{0x07, 0x0a, -1, 0x03, 0x00c2, 0x0446},
	cmdLutElement{0x07, 0x18, -1, 0x03, 0x00c2, 0x0846},
	cmdLutElement{0x08, 0x05, -1, 0x03, 0x0142, 0x0046},
	cmdLutElement{0x08, 0x05, -1, 0x03, 0x0142, 0x0066},
	cmdLutElement{0x08, 0x06, -1, 0x03, 0x0142, 0x0086},
	cmdLutElement{0x08, 0x07, -1, 0x03, 0x0142, 0x00c6},
	cmdLutElement{0x08, 0x08, -1, 0x03, 0x0142, 0x0146},
	cmdLutElement{0x08, 0x09, -1, 0x03, 0x0142, 0x0246},
	cmdLutElement{0x08, 0x0a, -1, 0x03, 0x0142, 0x0446},
	cmdLutElement{0x08, 0x18, -1, 0x03, 0x0142, 0x0846},
	cmdLutElement{0x09, 0x05, -1, 0x03, 0x0242, 0x0046},
	cmdLutElement{0x09, 0x05, -1, 0x03, 0x0242, 0x0066},
	cmdLutElement{0x09, 0x06, -1, 0x03, 0x0242, 0x0086},
	cmdLutElement{0x09, 0x07, -1, 0x03, 0x0242, 0x00c6},
	cmdLutElement{0x09, 0x08, -1, 0x03, 0x0242, 0x0146},
	cmdLutElement{0x09, 0x09, -1, 0x03, 0x0242, 0x0246},
	cmdLutElement{0x09, 0x0a, -1, 0x03, 0x0242, 0x0446},
	cmdLutElement{0x09, 0x18, -1, 0x03, 0x0242, 0x0846},
	cmdLutElement{0x0a, 0x05, -1, 0x03, 0x0442, 0x0046},
	cmdLutElement{0x0a, 0x05, -1, 0x03, 0x0442, 0x0066},
	cmdLutElement{0x0a, 0x06, -1, 0x03, 0x0442, 0x0086},
	cmdLutElement{0x0a, 0x07, -1, 0x03, 0x0442, 0x00c6},
	cmdLutElement{0x0a, 0x08, -1, 0x03, 0x0442, 0x0146},
	cmdLutElement{0x0a, 0x09, -1, 0x03, 0x0442, 0x0246},
	cmdLutElement{0x0a, 0x0a, -1, 0x03, 0x0442, 0x0446},
	cmdLutElement{0x0a, 0x18, -1, 0x03, 0x0442, 0x0846},
	cmdLutElement{0x0c, 0x05, -1, 0x03, 0x0842, 0x0046},
	cmdLutElement{0x0c, 0x05, -1, 0x03, 0x0842, 0x0066},
	cmdLutElement{0x0c, 0x06, -1, 0x03, 0x0842, 0x0086},
	cmdLutElement{0x0c, 0x07, -1, 0x03, 0x0842, 0x00c6},
	cmdLutElement{0x0c, 0x08, -1, 0x03, 0x0842, 0x0146},
	cmdLutElement{0x0c, 0x09, -1, 0x03, 0x0842, 0x0246},
	cmdLutElement{0x0c, 0x0a, -1, 0x03, 0x0842, 0x0446},
	cmdLutElement{0x0c, 0x18, -1, 0x03, 0x0842, 0x0846},
	cmdLutElement{0x0e, 0x05, -1, 0x03, 0x1842, 0x0046},
	cmdLutElement{0x0e, 0x05, -1, 0x03, 0x1842, 0x0066},
	cmdLutElement{0x0e, 0x06, -1, 0x03, 0x1842, 0x0086},
	cmdLutElement{0x0e, 0x07, -1, 0x03, 0x1842, 0x00c6},
	cmdLutElement{0x0e, 0x08, -1, 0x03, 0x1842, 0x0146},
	cmdLutElement{0x0e, 0x09, -1, 0x03, 0x1842, 0x0246},
	cmdLutElement{0x0e, 0x0a, -1, 0x03, 0x1842, 0x0446},
	cmdLutElement{0x0e, 0x18, -1, 0x03, 0x1842, 0x0846},
	cmdLutElement{0x18, 0x05, -1, 0x03, 0x5842, 0x0046},
	cmdLutElement{0x18, 0x05, -1, 0x03, 0x5842, 0x0066},
	cmdLutElement{0x18, 0x06, -1, 0x03, 0x5842, 0x0086},
	cmdLutElement{0x18, 0x07, -1, 0x03, 0x5842, 0x00c6},
	cmdLutElement{0x18, 0x08, -1, 0x03, 0x5842, 0x0146},
	cmdLutElement{0x18, 0x09, -1, 0x03, 0x5842, 0x0246},
	cmdLutElement{0x18, 0x0a, -1, 0x03, 0x5842, 0x0446},
	cmdLutElement{0x18, 0x18, -1, 0x03, 0x5842, 0x0846},
}