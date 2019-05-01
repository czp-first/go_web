// 9.6 加密和解密数据
// 有的时候，想把一些敏感数据加密后存储起来，在将来的某个时候，随需将他们解密出来，此时应该选用对称加密算法来满足需求

// base64加解密
// 如果web应用足够简单，数据的安全性没有那么严格的要求，那么可以采用一种比较简单的加密方法是base64，这种方式实现起来比较简单，go的base64包已经很好的支持了这个
// 例：1.go

// 高级加解密
// go的crypto里面支持对称加密的高级加解密包有：
//	crypto/ases包：AES(Advanced Encryption Standard), 又称Rijndael加密法，是美国联邦政府采用的一种区块加密标准
//	crypto/des包：DEA(Data Encryption Algorithm)，是一种对称加密算法，是目前使用最广泛的密钥系统，特别是在保护金融数据的安全中

// 因为这两种算法使用方法类似，所以在此，仅用ases包为例来讲解他们的使用
// 例：2.go
// 上面通过调用函数ases.NewCipher(参数key必须是16、24或者32位的[]byte，分别对应AES-128，AES-192或者AES-256算法)，返回了一个cipher.Block接口，这个接口实现了三个功能
type Block interface {
	// BlockSize returns the cipher's block size
	BlockSize() int

	// Encrypt encrypts the first block in src into dst
	// Dst and src may point at the same memory
	Encrypt(dst, src []byte)

	// Decrypt decrypts the first block in src into dst
	// Dst and src may point at the same memory
	Decrypt(dst, src []byte)
	
}