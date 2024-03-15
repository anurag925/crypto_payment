// Code generated by "enumer -type=TransactionStatus -json -transform=snake -trimprefix=TransactionStatus"; DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _TransactionStatusName = "createdpendingcompletedrefundedrejectedcanceled"

var _TransactionStatusIndex = [...]uint8{0, 7, 14, 23, 31, 39, 47}

const _TransactionStatusLowerName = "createdpendingcompletedrefundedrejectedcanceled"

func (i TransactionStatus) String() string {
	if i < 0 || i >= TransactionStatus(len(_TransactionStatusIndex)-1) {
		return fmt.Sprintf("TransactionStatus(%d)", i)
	}
	return _TransactionStatusName[_TransactionStatusIndex[i]:_TransactionStatusIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TransactionStatusNoOp() {
	var x [1]struct{}
	_ = x[TransactionStatusCreated-(0)]
	_ = x[TransactionStatusPending-(1)]
	_ = x[TransactionStatusCompleted-(2)]
	_ = x[TransactionStatusRefunded-(3)]
	_ = x[TransactionStatusRejected-(4)]
	_ = x[TransactionStatusCanceled-(5)]
}

var _TransactionStatusValues = []TransactionStatus{TransactionStatusCreated, TransactionStatusPending, TransactionStatusCompleted, TransactionStatusRefunded, TransactionStatusRejected, TransactionStatusCanceled}

var _TransactionStatusNameToValueMap = map[string]TransactionStatus{
	_TransactionStatusName[0:7]:        TransactionStatusCreated,
	_TransactionStatusLowerName[0:7]:   TransactionStatusCreated,
	_TransactionStatusName[7:14]:       TransactionStatusPending,
	_TransactionStatusLowerName[7:14]:  TransactionStatusPending,
	_TransactionStatusName[14:23]:      TransactionStatusCompleted,
	_TransactionStatusLowerName[14:23]: TransactionStatusCompleted,
	_TransactionStatusName[23:31]:      TransactionStatusRefunded,
	_TransactionStatusLowerName[23:31]: TransactionStatusRefunded,
	_TransactionStatusName[31:39]:      TransactionStatusRejected,
	_TransactionStatusLowerName[31:39]: TransactionStatusRejected,
	_TransactionStatusName[39:47]:      TransactionStatusCanceled,
	_TransactionStatusLowerName[39:47]: TransactionStatusCanceled,
}

var _TransactionStatusNames = []string{
	_TransactionStatusName[0:7],
	_TransactionStatusName[7:14],
	_TransactionStatusName[14:23],
	_TransactionStatusName[23:31],
	_TransactionStatusName[31:39],
	_TransactionStatusName[39:47],
}

// TransactionStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TransactionStatusString(s string) (TransactionStatus, error) {
	if val, ok := _TransactionStatusNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _TransactionStatusNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to TransactionStatus values", s)
}

// TransactionStatusValues returns all values of the enum
func TransactionStatusValues() []TransactionStatus {
	return _TransactionStatusValues
}

// TransactionStatusStrings returns a slice of all String values of the enum
func TransactionStatusStrings() []string {
	strs := make([]string, len(_TransactionStatusNames))
	copy(strs, _TransactionStatusNames)
	return strs
}

// IsATransactionStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i TransactionStatus) IsATransactionStatus() bool {
	for _, v := range _TransactionStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for TransactionStatus
func (i TransactionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionStatus
func (i *TransactionStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TransactionStatus should be a string, got %s", data)
	}

	var err error
	*i, err = TransactionStatusString(s)
	return err
}
