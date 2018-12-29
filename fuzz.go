// +build gofuzz

package branca

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

var b, _ = NewBranca([]byte("abcdefghabcdefghabcdefghabcdefgh"))

// ////////////////////////////////////////////////////////////////////////////////// //

func FuzzEncode(data []byte) int {
	_, err := b.Encode(data)

	if err != nil {
		return 1
	}

	return 0
}

func FuzzDecode(data []byte) int {
	_, err := b.Decode(data)

	if err != nil {
		return 1
	}

	return 0
}