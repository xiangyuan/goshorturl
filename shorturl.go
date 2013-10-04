package main

/*
1)将长网址md5生成32位签名串,分为4段, 每段8个字节;
2)对这四段循环处理, 取8个字节, 将他看成16进制串与0x3fffffff(30位1)与操作, 即超过30位的忽略处理;
3)这30位分成6段, 每5位的数字作为字母表的索引取得特定字符, 依次进行获得6位字符串;
4)总的md5串可以获得4个6位串; 取里面的任意一个就可作为这个长url的短url地址;
这种算法,虽然会生成4个,但是仍然存在重复几率,下面的算法一和三,就是这种的实现.
*/
import (
	// "bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	// "strings"
)

const (
	signSalt = "Ak47"
	//	base32   = "abcdefghijklmnopqrstuvwxyz01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//base62 = "abcdefghijklmnopqrstuvwxyz01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base32 = "bcdefghjkmnpqrstuvwxyz0123456789"
)

func main() {
	url := "www.google.com"
	h := md5.New()
	io.WriteString(h, url+signSalt)
	hx := h.Sum(nil)
	var result [4]string
	s := hex.EncodeToString(hx)
	length := len(s) / 8
	for i := 0; i < length; i++ {
		subhex := s[i*8 : i*8+8]
		// fmt.Println(strconv.QuoteToASCII("0x" + subhex))
		//hexstr := strconv.QuoteToASCII("0x" + subhex)
		bit, _ := strconv.ParseInt(subhex, 16, 64)
		// fmt.Println(bit & 0x3FFFFFFF)
		var chHex = 0x3FFFFFFF & bit
		var outchars = ""
		for i := 0; i < 6; i++ {
			//base 62 3D base32 1F
			//比base值少2
			val := 0x0000001F & chHex
			outchars += base32[val : val+1]
			chHex = chHex >> 5
		}
		result[i] = outchars
	}
	fmt.Println(result)
	// fmt.Println(hex.EncodeToString(hx))
	// fmt.Println(strings.ToUpper(base32))

}
