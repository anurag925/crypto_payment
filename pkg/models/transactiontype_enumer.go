// Code generated by "enumer -type=TransactionType -json -transform=snake -trimprefix=TransactionType"; DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _TransactionTypeName = "creditdebit"

var _TransactionTypeIndex = [...]uint8{0, 6, 11}

const _TransactionTypeLowerName = "creditdebit"

func (i TransactionType) String() string {
	if i < 0 || i >= TransactionType(len(_TransactionTypeIndex)-1) {
		return fmt.Sprintf("TransactionType(%d)", i)
	}
	return _TransactionTypeName[_TransactionTypeIndex[i]:_TransactionTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TransactionTypeNoOp() {
	var x [1]struct{}
	_ = x[TransactionTypeCredit-(0)]
	_ = x[TransactionTypeDebit-(1)]
}

var _TransactionTypeValues = []TransactionType{TransactionTypeCredit, TransactionTypeDebit}

var _TransactionTypeNameToValueMap = map[string]TransactionType{
	_TransactionTypeName[0:6]:       TransactionTypeCredit,
	_TransactionTypeLowerName[0:6]:  TransactionTypeCredit,
	_TransactionTypeName[6:11]:      TransactionTypeDebit,
	_TransactionTypeLowerName[6:11]: TransactionTypeDebit,
}

var _TransactionTypeNames = []string{
	_TransactionTypeName[0:6],
	_TransactionTypeName[6:11],
}

// TransactionTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TransactionTypeString(s string) (TransactionType, error) {
	if val, ok := _TransactionTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _TransactionTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to TransactionType values", s)
}

// TransactionTypeValues returns all values of the enum
func TransactionTypeValues() []TransactionType {
	return _TransactionTypeValues
}

// TransactionTypeStrings returns a slice of all String values of the enum
func TransactionTypeStrings() []string {
	strs := make([]string, len(_TransactionTypeNames))
	copy(strs, _TransactionTypeNames)
	return strs
}

// IsATransactionType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i TransactionType) IsATransactionType() bool {
	for _, v := range _TransactionTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for TransactionType
func (i TransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionType
func (i *TransactionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TransactionType should be a string, got %s", data)
	}

	var err error
	*i, err = TransactionTypeString(s)
	return err
}