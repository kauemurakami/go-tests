// UNIT TEST
package addresses

import (
	"testing"
)

type test_scenario struct {
	address_inserted string
	expected_return  string
}

func TestAddressType(t *testing.T) {

	scenarios_of_Test := []test_scenario{
		{"Street Abc", "Street"},
		{"Avenue xyz", "Avenue"},
		{"Road 138", "Road"},
		{"Square park", "Invalid type"},
		{"HiGhway dbo", "Highway"},
		{"", "Invalid type"},
	}

	for _, scenario := range scenarios_of_Test {
		receivedAddressType := AddressType(scenario.address_inserted)

		if receivedAddressType != scenario.expected_return {
			// método do param t, ele chama um erro no seu test
			// o erro sera logado no terminal e será considerado que
			// quebrou ou não está fazendo o que esperamos
			t.Errorf("Received type invalid, wait %s and receive %s",
				scenario.expected_return,
				receivedAddressType,
			)
		}
	}

}

// before refactor
// package addresses

// import (
// 	"testing"
// )

// func TestAddressType(t *testing.T) {
// 	address_to_test := "Avenue Paulista"
// 	expected_address_type := "Avenue"
// 	receivedAddressType := AddressType(address_to_test)

// 	if receivedAddressType != expected_address_type {
// 		// método do param t, ele chama um erro no seu test
// 		// o erro sera logado no terminal e será considerado que
// 		// quebrou ou não está fazendo o que esperamos
// 		t.Errorf("Received type invalid, wait %s and receive %s",
// 			expected_address_type,
// 			receivedAddressType,
// 		)
// 	}
// }
