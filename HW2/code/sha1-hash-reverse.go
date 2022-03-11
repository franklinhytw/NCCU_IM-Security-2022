package main

/***********************************************************************************
// PW cracker (密碼破解工具開發)

// •Input : SHA-1 hash value
// •Output: SHA-1 input, Time for breaking SHA-1
// PW domain P={a,b,c,d,….,z} (26 lowercase English letters ) in UTF-8 encoding
//        First: check SHA-1(nccu) is
// a786f899fefd42d46031741fed7157e059f517fa

// Second: input the following hash value and find the input of SHA-1 (and also the time used)
// e410b808a7f76c6890c9ecacf2b564ec98204fdb

// Please give your source code, execution file, and your answer
**********************************************************************************/

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func calcSha1(hash_value *[]byte) []byte {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	h := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write(*hash_value)

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	return bs
}

func reverseSha1Hash(hash_value *[]byte) []byte {
	var counter int = 0
	// UTF-8 code a~z = 0x61 ~ 0x7a

	// LIMITED to 10 characters
	for i := 1; i <= 3; i = i + 1 {
		byte_arr := make([]byte, i)

		// init array to set 'a'
		for j := range byte_arr {
			byte_arr[j] = 0x61
		}

		var offset byte = 0
		var shift_pointer int = 1

		for {
			byte_arr[i-1] = 0x61 + offset

			/* Try to hash */
			fmt.Println(string(byte_arr[:]))
			counter = counter + 1
			// guess_hash_value := calcSha1(&byte_arr)
			// res := bytes.Compare(guess_hash_value, byte_arr)

			// if res == 0 {
			// 	return byte_arr
			// }
			/***************/

			if offset >= 25 {
				if i == 1 {
					break
				}

				var escape bool = false

				for s := 1; s <= shift_pointer; s = s + 1 {
					fmt.Printf("s=%d, i=%d\n", s, i)
					if s == i {
						escape = true
						break
					}

					if byte_arr[i-1-s]+1 <= 0x7a {
						byte_arr[i-1-s] = byte_arr[i-1-s] + 1
						break
					} else {
						// init array to set 'a'
						for idx := 1; idx <= shift_pointer; idx = idx + 1 {
							fmt.Printf("idx=%d, shift_pointer=%d\n", idx, shift_pointer)
							byte_arr[i-idx] = 0x61
						}

						// next index
						shift_pointer = shift_pointer + 1
					}
				}

				if escape {
					break
				}

				// RESET TO ZERO
				offset = 0
			} else {
				offset = offset + 1
			}
		}
	}
	fmt.Printf("RESULT:%d\n", counter)

	return nil
}

func main() {
	// var sha1_hash_value string

	// fmt.Println("Please Input Original SHA-1 hash value: ")

	// // INPUT
	// fmt.Scanln(&sha1_hash_value)

	hash_byte_arr_data := []byte{0x00}
	// hash_byte_arr_data, err := hex.DecodeString(sha1_hash_value)
	// if err != nil {
	// 	// Not the hex string. STOP PROGRAM
	// 	panic(err)
	// }

	// timer start
	start := time.Now()

	answer := reverseSha1Hash(&hash_byte_arr_data)

	// timer end
	t := time.Now()

	// calculate spend time
	elapsed := t.Sub(start)

	fmt.Println(elapsed)

	fmt.Printf("The hash original value is:\n%s\n", answer)
}
