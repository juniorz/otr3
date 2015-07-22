package otr3

import "encoding/binary"

func (c *akeContext) genDataMsg(tlvsBytes []byte) []byte {
	msgHeader := dhMessage{
		protocolVersion:     c.protocolVersion(),
		needInstanceTag:     c.needInstanceTag(),
		senderInstanceTag:   uint32(0),
		receiverInstanceTag: uint32(0),
	}

	var tlvs []tlv
	index := 0
	for index < len(tlvsBytes) {
		atlv := tlv{}
		atlv.tlvType = binary.BigEndian.Uint16(tlvsBytes[index : index+2])
		atlv.tlvLength = binary.BigEndian.Uint16(tlvsBytes[index+2 : index+4])
		endOfTLV := index + 4 + int(atlv.tlvLength)
		atlv.tlvValue = tlvsBytes[index+4 : endOfTLV]
		tlvs = append(tlvs, atlv)
		index = endOfTLV
	}

	dataMessage := dataMsg{
		dhMessage: msgHeader,
		//TODO: implement IGNORE_UNREADABLE
		flag: 0x00,

		//TODO after key management
		senderKeyID:    c.ourKeyID - 1,
		recipientKeyID: c.theirKeyID,
		//TODO after key management
		y:          c.ourCurrentDHKeys.pub,
		topHalfCtr: [8]byte{},
		//tlv is properly formatted
		dataMsgEncrypted: []byte{},
		//TODO after key management
		authenticator:   [20]byte{},
		oldRevealKeyMAC: []byte{},
	}

	return dataMessage.serialize()
}