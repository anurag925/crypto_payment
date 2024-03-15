// Code generated by "enumer -type=AddressStatus -json -transform=snake -trimprefix=AddressStatus"; DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _AddressStatusName = "activeremoved"

var _AddressStatusIndex = [...]uint8{0, 6, 13}

const _AddressStatusLowerName = "activeremoved"

func (i AddressStatus) String() string {
	if i < 0 || i >= AddressStatus(len(_AddressStatusIndex)-1) {
		return fmt.Sprintf("AddressStatus(%d)", i)
	}
	return _AddressStatusName[_AddressStatusIndex[i]:_AddressStatusIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _AddressStatusNoOp() {
	var x [1]struct{}
	_ = x[AddressStatusActive-(0)]
	_ = x[AddressStatusRemoved-(1)]
}

var _AddressStatusValues = []AddressStatus{AddressStatusActive, AddressStatusRemoved}

var _AddressStatusNameToValueMap = map[string]AddressStatus{
	_AddressStatusName[0:6]:       AddressStatusActive,
	_AddressStatusLowerName[0:6]:  AddressStatusActive,
	_AddressStatusName[6:13]:      AddressStatusRemoved,
	_AddressStatusLowerName[6:13]: AddressStatusRemoved,
}

var _AddressStatusNames = []string{
	_AddressStatusName[0:6],
	_AddressStatusName[6:13],
}

// AddressStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AddressStatusString(s string) (AddressStatus, error) {
	if val, ok := _AddressStatusNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _AddressStatusNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AddressStatus values", s)
}

// AddressStatusValues returns all values of the enum
func AddressStatusValues() []AddressStatus {
	return _AddressStatusValues
}

// AddressStatusStrings returns a slice of all String values of the enum
func AddressStatusStrings() []string {
	strs := make([]string, len(_AddressStatusNames))
	copy(strs, _AddressStatusNames)
	return strs
}

// IsAAddressStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AddressStatus) IsAAddressStatus() bool {
	for _, v := range _AddressStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for AddressStatus
func (i AddressStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressStatus
func (i *AddressStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AddressStatus should be a string, got %s", data)
	}

	var err error
	*i, err = AddressStatusString(s)
	return err
}