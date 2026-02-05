package decoder

import (
	"encoding/binary"
	"fmt"
)

type Reader struct {
	data []byte
	pos  int
}

func NewReader(b []byte) *Reader {
	return &Reader{data: b}
}

func (r *Reader) read(n int) ([]byte, error) {
	if r.pos+n > len(r.data) {
		return nil, fmt.Errorf("Not enough data")
	}
	out := r.data[r.pos : r.pos+n]
	r.pos += n 
	return out,nil
}

func (r *Reader) ReadUint32() (uint32,error) {
	b,err := r.read(4)
	if err != nil {
		return  0,err
	}
	return  binary.LittleEndian.Uint32(b),nil
}

func (r *Reader) ReadVarInt() (uint64, error){
	prefix, err := r.read(1)
	if err != nil {
		return 0, err
	}

	switch prefix[0] {
	case 0xfd:
		b,err := r.read(2)
		if err != nil {
			return 0, err
		}
		return uint64(binary.LittleEndian.Uint16(b)), nil

	case 0xfe:
		b, err := r.read(4)
		if err != nil {
			return 0, err
		}
		return uint64(binary.LittleEndian.Uint32(b)), nil

	case 0xff:
		b, err := r.read(8)
		if err != nil {
			return 0, err
		}
		return binary.LittleEndian.Uint64(b), nil

	default:
		return uint64(prefix[0]), nil
	}
		
	
}

func (r *Reader) ReadUint64()(uint64, error){
	b,err := r.read(8)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b),nil
	
}
