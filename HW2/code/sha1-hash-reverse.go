package main

// NCCU Information Security HW2 - PW cracker
// Mar 11, 2022
// 109971014 林翰陽

// Goroutine reference:
// https://javasgl.github.io/goroutinue-limit/

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
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math"
	"runtime"
	"sync"
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

func toNumberSystem26(n int, rtn *[]byte) {
	if n < 26 {
		*rtn = append(*rtn, byte(n)+0x61)
	} else {
		toNumberSystem26(n/26-1, rtn)
		toNumberSystem26(n%26, rtn)
	}
}

func reverseSha1Hash(hash_value *[]byte) []byte {
	var total int = 0

	// // UTF-8 code a~z = 0x61 ~ 0x7a

	// // LIMITED to 10 characters

	for n := 1; n <= 10; n = n + 1 {
		total = total + int(math.Pow(26, float64(n)))
	}

	var ans []byte
	is_done := false

	var wg sync.WaitGroup
	limiter := make(chan bool, runtime.NumCPU())

	sp := 0

	for {
		pre_sp := sp

		if sp+1000 <= total {
			sp = sp + 1000
		} else {
			sp = total
		}

		if is_done {
			break
		}

		wg.Add(1)
		limiter <- true
		// fmt.Printf("calculating range: %d ~ %d\n", pre_sp, sp)

		go func(start_count int, end_count int, done *bool, limiter chan bool, ans *[]byte) {
			for j := start_count; j < end_count; j = j + 1 {
				var byte_arr []byte
				toNumberSystem26(j, &byte_arr)

				guess_hash_value := calcSha1(&byte_arr)
				// fmt.Printf("NUM=%d -> 26 NUM = %s, HASH:%s\n", j, string(byte_arr[:]), hex.EncodeToString(guess_hash_value))
				res := bytes.Compare(guess_hash_value, *hash_value)

				if res == 0 {
					*done = true
					*ans = byte_arr
					break
				}
			}
			defer wg.Done()
			<-limiter
		}(pre_sp, sp, &is_done, limiter, &ans)

		if sp == total {
			break
		}
	}

	wg.Wait()

	if is_done {
		return ans
	} else {
		return nil
	}
}

func main() {
	var sha1_hash_value string

	fmt.Println("Please Input Original SHA-1 hash value: ")

	// INPUT
	fmt.Scanln(&sha1_hash_value)
	// sha1_hash_value = "d134875d6721c9ce683e1a30cbad9de1354a65a1"

	hash_byte_arr_data, err := hex.DecodeString(sha1_hash_value)
	if err != nil {
		// Not the hex string. STOP PROGRAM
		panic(err)
	}
	// timer start
	start := time.Now()

	fmt.Printf("Calculating... Please Wait\n\n")
	answer := reverseSha1Hash(&hash_byte_arr_data)

	// timer end
	t := time.Now()

	// calculate spend time
	elapsed := t.Sub(start)

	fmt.Print("Spend Time: ")
	fmt.Println(elapsed)

	fmt.Printf("The hash original value is:\n%s\n\n\n\n", string(answer[:]))

	fmt.Println("Press the Enter Key to terminate the console screen!")
	fmt.Scanln() // wait for Enter Key
}
