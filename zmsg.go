// *** DO NOT MODIFY ***
// AUTOGENERATED BY go generate from msg_generate.go

package dns

//import (
//"encoding/base64"
//"net"
//)

// pack*() functions

func (rr *A) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := packHeader(rr.Hdr, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packDataA(rr.A, msg, off)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *AAAA) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := packHeader(rr.Hdr, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packDataAAAA(rr.AAAA, msg, off)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *L32) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := packHeader(rr.Hdr, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packUint16(rr.Preference, msg, off, len(msg))
	if err != nil {
		return off, err
	}
	off, err = packDataA(rr.Locator32, msg, off)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

func (rr *MX) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {
	off, err := packHeader(rr.Hdr, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	headerEnd := off
	off, err = packUint16(rr.Preference, msg, off, len(msg))
	if err != nil {
		return off, err
	}
	off, err = PackDomainName(rr.Mx, msg, off, compression, compress)
	if err != nil {
		return off, err
	}
	rr.Header().Rdlength = uint16(off - headerEnd)
	return off, nil
}

// unpack*() functions

func unpackA(h RR_Header, msg []byte, off int) (RR, int, error) {
	if dynamicUpdate(h) {
		return nil, off, nil
	}
	var err error
	rr := new(A)
	lenmsg := len(msg)
	_ = lenmsg

	rr.Hdr = h

	rr.A, off, err = unpackDataA(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackAAAA(h RR_Header, msg []byte, off int) (RR, int, error) {
	if dynamicUpdate(h) {
		return nil, off, nil
	}
	var err error
	rr := new(AAAA)
	lenmsg := len(msg)
	_ = lenmsg

	rr.Hdr = h

	rr.AAAA, off, err = unpackDataAAAA(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackL32(h RR_Header, msg []byte, off int) (RR, int, error) {
	if dynamicUpdate(h) {
		return nil, off, nil
	}
	var err error
	rr := new(L32)
	lenmsg := len(msg)
	_ = lenmsg

	rr.Hdr = h

	rr.Preference, off, err = unpackUint16(msg, off, lenmsg)
	if err != nil {
		return rr, off, err
	}
	if off == lenmsg {
		return rr, off, nil
	}
	rr.Locator32, off, err = unpackDataA(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}

func unpackMX(h RR_Header, msg []byte, off int) (RR, int, error) {
	if dynamicUpdate(h) {
		return nil, off, nil
	}
	var err error
	rr := new(MX)
	lenmsg := len(msg)
	_ = lenmsg

	rr.Hdr = h

	rr.Preference, off, err = unpackUint16(msg, off, lenmsg)
	if err != nil {
		return rr, off, err
	}
	if off == lenmsg {
		return rr, off, nil
	}
	rr.Mx, off, err = UnpackDomainName(msg, off)
	if err != nil {
		return rr, off, err
	}
	return rr, off, nil
}
