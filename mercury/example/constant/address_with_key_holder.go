package constant

const (
	TEST_ADDRESS0                    = "ckt1qyq28wze3cw48ek9az0g4jmtfs6d8td38u4s6hp2s0"
	TEST_PUBKEY0                     = "0xa3b8598e1d53e6c5e89e8acb6b4c34d3adb13f2b"
	TEST_ADDRESS1                    = "ckt1qyq27z6pccncqlaamnh8ttapwn260egnt67ss2cwvz"
	TEST_PUBKEY1                     = "0xaf0b41c627807fbddcee75afa174d5a7e5135ebd"
	TEST_ADDRESS2                    = "ckt1qyqqtg06h75ymw098r3w0l3u4xklsj04tnsqctqrmc"
	TEST_PUBKEY2                     = "0x05a1fabfa84db9e538e2e7fe3ca9adf849f55ce0"
	TEST_ADDRESS3                    = "ckt1qyqzqfj8lmx9h8vvhk62uut8us844v0yh2hsnqvvgc"
	TEST_PUBKEY3                     = "0x202647fecc5b9d8cbdb4ae7167e40f5ab1e4baaf"
	TEST_ADDRESS4                    = "ckt1qyqg88ccqm59ksxp85788pnqg4rkejdgcg2qxcu2qf"
	TEST_PUBKEY4                     = "0x839f1806e85b40c13d3c73866045476cc9a8c214"
	TEST_KEY0                        = "6fc935dad260867c749cf1ba6602d5f5ed7fb1131f1beb65be2d342e912eaafe"
	TEST_KEY1                        = "9d8ca87d75d150692211fa62b0d30de4d1ee6c530d5678b40b8cedacf0750d0f"
	TEST_KEY2                        = "88a09e06735d89452552e359a052315ab5130dc2e4d864ae3eed21d6505b2f67"
	TEST_KEY3                        = "2d4cf0546a1dc93092ad56f2e18fbe6e41ee477d9dec0575cf43b69740ce9f74"
	TEST_KEY4                        = "5e46fdbb6ffd86d232080dc71f24b60df2a119e0102ca45a7c165472de14c104"
	CEX_ADDRESS                      = "ckt1qyqz50qxup4cw2dvg0exg8y9m5e90dykuh6q72rj0h"
	QUERY_TRANSACTION_ADDRESS_PUBKEY = "0x2a3c06e06b8729ac43f2641c85dd3257b496e5f4"
	CEX_KEY                          = "b274e3a1c8ece62367c7165ec9bca18112ae1386a67ccb95e7acd384af017cbf"
	QUERY_TRANSACTION_ADDRESS        = "ckt1qyqp5nlkxkvwgwhee4pryj4mwetl4pyut0psq2h92t"
	QUERY_TRANSACTION_KEY_PUBKEY     = "0x1a4ff63598e43af9cd42324abb7657fa849c5bc3"
	QUERY_TRANSACTION_KEY            = "385b57e3fedf89e5b553a3274e7039f7be742040a5af98303de29aff61b05c2c"
)

func GetKey(address string) string {
	switch address {
	case TEST_ADDRESS0:
		return TEST_KEY0
	case TEST_ADDRESS1:
		return TEST_KEY1
	case TEST_ADDRESS2:
		return TEST_KEY2
	case TEST_ADDRESS3:
		return TEST_KEY3
	case TEST_ADDRESS4:
		return TEST_KEY4
	case CEX_ADDRESS:
		return CEX_KEY
	case QUERY_TRANSACTION_ADDRESS:
		return QUERY_TRANSACTION_KEY
	default:
		return ""
	}
}
