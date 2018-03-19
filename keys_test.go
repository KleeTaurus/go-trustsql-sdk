package trustsql

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestGeneratePairkey(t *testing.T) {
	keyPair := GeneratePairkey()

	/*
		log.Printf("Private Key: %s, len: %d\n", base64Encode(keyPair.PrivateKey), len(keyPair.PrivateKey))
		log.Printf("Public key : %s, len: %d\n", base64Encode(keyPair.PublicKey), len(keyPair.PublicKey))
		log.Printf("Address    : %s, len: %d\n", keyPair.GetAddress(), len(keyPair.GetAddress()))
	*/

	if len(keyPair.PrivateKey) != 32 {
		t.Errorf("Incorrect length of the private key, it should be 32 bytes\n")
	}

	if len(keyPair.PublicKey) != 65 {
		t.Errorf("Incorrect length of the public key, it should be 65 bytes\n")
	}

	if len(keyPair.GenerateAddrByPubkey()) != 34 && len(keyPair.GenerateAddrByPubkey()) != 33 {
		t.Errorf("Incorrect length of the address, it should be 34 or 33 bytes\n")
	}
}

func TestPublicKeyAndAddress(t *testing.T) {

	keyPairs := [][]string{
		[]string{"cmcXYkcVkglRwvzfolcC1dh5vGQQNPPe2mF9p5rPbZ4=", "BFtXfRfxD5eIeeLr3j643FBVL6/wXa+/z8STliv+OAUd8rgRfDhsGL2DAgn/bfYPQSw0OdUNziJjpCr7OhgLSsw=", "12RFccDGioxkb9zDsxjqGXH3XDnsz813Wo"},
		[]string{"0QGhrwF23Hboaad84UCCJr2V3b9VUF1GZyB7d65piIE=", "BPNO+HvYDKhk//o7d8nm9D6+NUR1WOHmGCqI5hFSbbN/DX3+hD8EMdiz6HDtqWGsXyO3zowLvWS2TfouGFVUtmM=", "1Bg1FB5fEeZ8W5s5nQRKWzDD2R1hzJ3NRH"},
		[]string{"RIC/B0zuBcrra4azng7tMe1OWVq0K8Ex/VhRUYvdW/M=", "BAHlSqteXyj4YCzLTJN5845h52a9MBKxxBkNMF4mb1RicMmFpjn+wcbNx4UpJlFK3E0Hyo5lu5Na9dzXqb6wWJA=", "16Xe1h3hsq9L111JHDEiJvrTRvfnaisGNV"},
		[]string{"Gktkhkq3VU42TvSJeM5576P+T9+iIX2JKz+LMau9lHA=", "BF7g77UDffh8z5PdnWHtCd7e2DmzOJ7DlSDbczYhOHd/rcDXnokr40JsPGbnwuNbfxnyePCK0V3SIG2gdRdaolc=", "1HrAiM3LuZUwnW1CkiUgmad3DTBoT6aVLy"},
		[]string{"ZaZAUBs6fvS+wg2/hFMfoWrOWxrIINFsPfNILTw+mvY=", "BIs9GHWKH54ZIPJF/IMyd0JPYiDt/bwmFMkwPoJYRAg5gzxyGdv8d0RZFjlwWMOW3ni2O6S3AiI2YFXWm11mEaw=", "1LVDjvZEXVCELeHCsizkW3NtXTgYoW69bt"},
		[]string{"kH26vlkPiK1AwCURQjToFrgDMjQC/BC50wRaoXb6Txc=", "BGS3V7dT7zH62GnKlfurLpBYhOP+jlQGJ1AoBhj70wieMfZ0jbZWNa7uqwAuJtelF5/Cgtm0FdVwjgLw/5cIaxU=", "1FR89NGuA8efNvSo5VfBazubaYRsD2AevQ"},
		[]string{"K2+/D1mWoyNwcFHlP8NpyMyZg38xzJ7spm9x+otzPP0=", "BFEeLPHlp1j0KrYi6VgnIYhZ9rFLg3WT8eBiWMBpgx5hSR0oroiKRwFj7tCsHe/i1brDp2bagg6oCJQHUva4iyw=", "13uTZqjfLLTK48StN1m43XjU6FwoTpmR4N"},
		[]string{"VyBLy1Jx1UtGKHPfvop4ucxhGZe+vumrTMPYVhbty9E=", "BF8jtXslMkf895UwrapXwwU5HepQiKu9MpeI4mtGp151AG7/uuTJ6N1M8uMI8sv1YUSfuo45OA3w7nk6Ul6l6OA=", "15ejR3mugnaZQTX6Dq8du8Xz5F2oEPHF5C"},
		[]string{"mma4fzgsdarKLqslGVm5h7XA3eFZDLgOOqrCdT45Y+k=", "BPzMhh5NwXO3Y8nl1jVJdnIjrd8XkwkPhitCElVjIdnfsZZnjn1UhjUJvxHQb4YOB3h5bM5r33Uvu2/vu86Ar14=", "18dLi1ScUuCnujqWJR1CiijimgUpEzNNWE"},
		[]string{"XP69NcpSmRvf+XnPzBc9OAF9jnRGkHlqMx6BEiC5sNk=", "BNVQXYROK30AbitAuLGKpLF8qaNwatU42t0C3CLavyKl/1n7Bnb27xFX3h5X0cP4X6IclcSbnwJp8VJYMQXHGW4=", "1BCcGj2dteMdKrRR5gyzF9bKkAeAqxhgky"}}

	for i := range keyPairs {
		privateKey := keyPairs[i][0]
		publicKey := keyPairs[i][1]
		address := keyPairs[i][2]

		derivedPublicKey, err := GeneratePubkeyByPrvkey(base64Decode(privateKey))
		if err != nil {
			log.Panic(err)
		}
		if publicKey != base64Encode(derivedPublicKey) {
			t.Errorf("Incorrect derived public key, target: %s, derived: %s\n", publicKey, base64Encode(derivedPublicKey))
		}

		keyPair := &KeyPair{base64Decode(privateKey), base64Decode(publicKey)}
		if address != string(keyPair.GenerateAddrByPubkey()) {
			t.Errorf("Incorrect derived address, target: %s, derived: %s\n", address, keyPair.GenerateAddrByPubkey())
		}
	}
}

func base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func base64Decode(input string) []byte {
	src, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		log.Panic(err)
	}

	return src
}
