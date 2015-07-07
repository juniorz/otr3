package otr3

import "math/big"

var (
	p         *big.Int // prime field
	pMinusTwo *big.Int
	q         *big.Int
	g1        *big.Int // group generator
)

func init() {
	p, _ = new(big.Int).SetString(
		"FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD1"+
			"29024E088A67CC74020BBEA63B139B22514A08798E3404DD"+
			"EF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245"+
			"E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7ED"+
			"EE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3D"+
			"C2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F"+
			"83655D23DCA3AD961C62F356208552BB9ED529077096966D"+
			"670C354E4ABC9804F1746C08CA237327FFFFFFFFFFFFFFFF", 16)

	q, _ = new(big.Int).SetString(
		"7FFFFFFFFFFFFFFFE487ED5110B4611A62633145C06E0E68"+
			"948127044533E63A0105DF531D89CD9128A5043CC71A026E"+
			"F7CA8CD9E69D218D98158536F92F8A1BA7F09AB6B6A8E122"+
			"F242DABB312F3F637A262174D31BF6B585FFAE5B7A035BF6"+
			"F71C35FDAD44CFD2D74F9208BE258FF324943328F6722D9E"+
			"E1003E5C50B1DF82CC6D241B0E2AE9CD348B1FD47E9267AF"+
			"C1B2AE91EE51D6CB0E3179AB1042A95DCF6A9483B84B4B36"+
			"B3861AA7255E4C0278BA36046511B993FFFFFFFFFFFFFFFF", 16)

	pMinusTwo = new(big.Int).Sub(p, new(big.Int).SetInt64(2))
	g1 = new(big.Int).SetInt64(2)
}

func isGroupElement(n *big.Int) bool {
	return g1.Cmp(n) != 1 && pMinusTwo.Cmp(n) != -1
}
