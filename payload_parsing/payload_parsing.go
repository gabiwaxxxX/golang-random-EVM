package payload_parsing

import (
	"fmt"
)

// fct that pars a ethereum payload and return the function string and the args

// first 4 bytes are signature bytes 0x1a3bef2b is the signature of the function
// cut other in 32 bytes chunks
func ParsePayload(payload string) (string, []string) {
	// remove 0x
	payload = payload[2:]
	// get the function signature
	functionSignature := payload[:8]
	// get the args
	args := payload[8:]
	// split args in 32 bytes chunks

	argsArray := SplitArgs(args)
	MappingPositionArgs(argsArray)
	// fmt.Println("================================")

	return functionSignature, argsArray
}

// 0x00  00000000000000000000000000000000000000000000000000000000000000c0
// 0x20  000000000000000000000000000000000000000000000002c3b728349d274000
// 0x40  000000000000000000000000000000000000000000000000000000003b9aca00
// 0x60  00000000000000000000000000000000000000000000000000000000635a10ea
// 0x80  000000000000000000000000bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c
// 0xa0  000000000000000000000000000000000000000000000000000000000000000b
// 0xc0  0000000000000000000000000000000000000000000000000000000000000003
// 0xe0  0000000000000000000000000000000000000000000000000000000f387c0000
// 0x100 000000000000000000000000e5e2fcc3387460a7aa2a680d4d01301a74ec3c9a
// 0x120 0000000000000000000000003b6c2af5828ffba097bf009681b4e98a25ed8408
// 0x140 0000000000000000000000000000000000000000000000000000000e752c0000
// 0x160 00000000000000000000000091350ccfab95b60893fd840d6c127c02feb99554
// 0x180 00000000000000000000000077d10315f11332193e6ae98359bf37d415b28565
// 0x1a0 0000000000000000000000000000000000000000000000000000000f387c0000
// 0x1c0 00000000000000000000000034bb330f0fb9b81d61118cb86fdb7d8a7740f19e
// 0x1e0 000000000000000000000000994e40bb23bc879d18d40f6e32c62a62d517f56c
// function mapping from 32 bytes chunks position in word of 32 bytes and the args
func MappingPositionArgs(args []string) map[int]string {
	var mapping = make(map[int]string)
	for i, arg := range args {

		fmt.Println(bytesWordNumber(i), " ", arg)
		// get the first 32 bytes
		mapping[i] = arg
	}
	return mapping
}

func bytesWordNumber(i int) string {
	// (i+1) * 32 bytes in hex format
	return fmt.Sprintf("%x", (i+1)*32)
}

// split args in 32 bytes chunks
func SplitArgs(args string) []string {
	var argsArray []string
	for len(args) > 0 {
		if len(args) >= 64 {
			argsArray = append(argsArray, args[:64])
			args = args[64:]
		} else {
			argsArray = append(argsArray, args)
			args = ""
		}
	}
	return argsArray
}

//presentation
