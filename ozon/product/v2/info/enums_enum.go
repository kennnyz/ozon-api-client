// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.5
// Revision: b9e7d1ac24b2b7f6a5b451fa3d21706ffd8d79e2
// Build Date: 2023-01-30T01:49:43Z
// Built By: goreleaser

package info

import (
	"fmt"
	"strings"
)

const (
	// ListResponseResultItemCurrencyCodeRUB is a ListResponseResultItemCurrencyCode of type RUB.
	ListResponseResultItemCurrencyCodeRUB ListResponseResultItemCurrencyCode = iota + 1
	// ListResponseResultItemCurrencyCodeBYN is a ListResponseResultItemCurrencyCode of type BYN.
	ListResponseResultItemCurrencyCodeBYN
	// ListResponseResultItemCurrencyCodeKZT is a ListResponseResultItemCurrencyCode of type KZT.
	ListResponseResultItemCurrencyCodeKZT
	// ListResponseResultItemCurrencyCodeEUR is a ListResponseResultItemCurrencyCode of type EUR.
	ListResponseResultItemCurrencyCodeEUR
	// ListResponseResultItemCurrencyCodeUSD is a ListResponseResultItemCurrencyCode of type USD.
	ListResponseResultItemCurrencyCodeUSD
	// ListResponseResultItemCurrencyCodeCNY is a ListResponseResultItemCurrencyCode of type CNY.
	ListResponseResultItemCurrencyCodeCNY
)

var ErrInvalidListResponseResultItemCurrencyCode = fmt.Errorf("not a valid ListResponseResultItemCurrencyCode, try [%s]", strings.Join(_ListResponseResultItemCurrencyCodeNames, ", "))

const _ListResponseResultItemCurrencyCodeName = "RUBBYNKZTEURUSDCNY"

var _ListResponseResultItemCurrencyCodeNames = []string{
	_ListResponseResultItemCurrencyCodeName[0:3],
	_ListResponseResultItemCurrencyCodeName[3:6],
	_ListResponseResultItemCurrencyCodeName[6:9],
	_ListResponseResultItemCurrencyCodeName[9:12],
	_ListResponseResultItemCurrencyCodeName[12:15],
	_ListResponseResultItemCurrencyCodeName[15:18],
}

// ListResponseResultItemCurrencyCodeNames returns a list of possible string values of ListResponseResultItemCurrencyCode.
func ListResponseResultItemCurrencyCodeNames() []string {
	tmp := make([]string, len(_ListResponseResultItemCurrencyCodeNames))
	copy(tmp, _ListResponseResultItemCurrencyCodeNames)
	return tmp
}

var _ListResponseResultItemCurrencyCodeMap = map[ListResponseResultItemCurrencyCode]string{
	ListResponseResultItemCurrencyCodeRUB: _ListResponseResultItemCurrencyCodeName[0:3],
	ListResponseResultItemCurrencyCodeBYN: _ListResponseResultItemCurrencyCodeName[3:6],
	ListResponseResultItemCurrencyCodeKZT: _ListResponseResultItemCurrencyCodeName[6:9],
	ListResponseResultItemCurrencyCodeEUR: _ListResponseResultItemCurrencyCodeName[9:12],
	ListResponseResultItemCurrencyCodeUSD: _ListResponseResultItemCurrencyCodeName[12:15],
	ListResponseResultItemCurrencyCodeCNY: _ListResponseResultItemCurrencyCodeName[15:18],
}

// String implements the Stringer interface.
func (x ListResponseResultItemCurrencyCode) String() string {
	if str, ok := _ListResponseResultItemCurrencyCodeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ListResponseResultItemCurrencyCode(%d)", x)
}

var _ListResponseResultItemCurrencyCodeValue = map[string]ListResponseResultItemCurrencyCode{
	_ListResponseResultItemCurrencyCodeName[0:3]:   ListResponseResultItemCurrencyCodeRUB,
	_ListResponseResultItemCurrencyCodeName[3:6]:   ListResponseResultItemCurrencyCodeBYN,
	_ListResponseResultItemCurrencyCodeName[6:9]:   ListResponseResultItemCurrencyCodeKZT,
	_ListResponseResultItemCurrencyCodeName[9:12]:  ListResponseResultItemCurrencyCodeEUR,
	_ListResponseResultItemCurrencyCodeName[12:15]: ListResponseResultItemCurrencyCodeUSD,
	_ListResponseResultItemCurrencyCodeName[15:18]: ListResponseResultItemCurrencyCodeCNY,
}

// ParseListResponseResultItemCurrencyCode attempts to convert a string to a ListResponseResultItemCurrencyCode.
func ParseListResponseResultItemCurrencyCode(name string) (ListResponseResultItemCurrencyCode, error) {
	if x, ok := _ListResponseResultItemCurrencyCodeValue[name]; ok {
		return x, nil
	}
	return ListResponseResultItemCurrencyCode(0), fmt.Errorf("%s is %w", name, ErrInvalidListResponseResultItemCurrencyCode)
}

// MarshalText implements the text marshaller method.
func (x ListResponseResultItemCurrencyCode) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ListResponseResultItemCurrencyCode) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseListResponseResultItemCurrencyCode(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// ListResponseResultItemPriceIndexesPriceIndexWITHOUTINDEX is a ListResponseResultItemPriceIndexesPriceIndex of type WITHOUT_INDEX.
	ListResponseResultItemPriceIndexesPriceIndexWITHOUTINDEX ListResponseResultItemPriceIndexesPriceIndex = iota + 1
	// ListResponseResultItemPriceIndexesPriceIndexPROFIT is a ListResponseResultItemPriceIndexesPriceIndex of type PROFIT.
	ListResponseResultItemPriceIndexesPriceIndexPROFIT
	// ListResponseResultItemPriceIndexesPriceIndexAVGPROFIT is a ListResponseResultItemPriceIndexesPriceIndex of type AVG_PROFIT.
	ListResponseResultItemPriceIndexesPriceIndexAVGPROFIT
	// ListResponseResultItemPriceIndexesPriceIndexNONPROFIT is a ListResponseResultItemPriceIndexesPriceIndex of type NON_PROFIT.
	ListResponseResultItemPriceIndexesPriceIndexNONPROFIT
)

var ErrInvalidListResponseResultItemPriceIndexesPriceIndex = fmt.Errorf("not a valid ListResponseResultItemPriceIndexesPriceIndex, try [%s]", strings.Join(_ListResponseResultItemPriceIndexesPriceIndexNames, ", "))

const _ListResponseResultItemPriceIndexesPriceIndexName = "WITHOUT_INDEXPROFITAVG_PROFITNON_PROFIT"

var _ListResponseResultItemPriceIndexesPriceIndexNames = []string{
	_ListResponseResultItemPriceIndexesPriceIndexName[0:13],
	_ListResponseResultItemPriceIndexesPriceIndexName[13:19],
	_ListResponseResultItemPriceIndexesPriceIndexName[19:29],
	_ListResponseResultItemPriceIndexesPriceIndexName[29:39],
}

// ListResponseResultItemPriceIndexesPriceIndexNames returns a list of possible string values of ListResponseResultItemPriceIndexesPriceIndex.
func ListResponseResultItemPriceIndexesPriceIndexNames() []string {
	tmp := make([]string, len(_ListResponseResultItemPriceIndexesPriceIndexNames))
	copy(tmp, _ListResponseResultItemPriceIndexesPriceIndexNames)
	return tmp
}

var _ListResponseResultItemPriceIndexesPriceIndexMap = map[ListResponseResultItemPriceIndexesPriceIndex]string{
	ListResponseResultItemPriceIndexesPriceIndexWITHOUTINDEX: _ListResponseResultItemPriceIndexesPriceIndexName[0:13],
	ListResponseResultItemPriceIndexesPriceIndexPROFIT:       _ListResponseResultItemPriceIndexesPriceIndexName[13:19],
	ListResponseResultItemPriceIndexesPriceIndexAVGPROFIT:    _ListResponseResultItemPriceIndexesPriceIndexName[19:29],
	ListResponseResultItemPriceIndexesPriceIndexNONPROFIT:    _ListResponseResultItemPriceIndexesPriceIndexName[29:39],
}

// String implements the Stringer interface.
func (x ListResponseResultItemPriceIndexesPriceIndex) String() string {
	if str, ok := _ListResponseResultItemPriceIndexesPriceIndexMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ListResponseResultItemPriceIndexesPriceIndex(%d)", x)
}

var _ListResponseResultItemPriceIndexesPriceIndexValue = map[string]ListResponseResultItemPriceIndexesPriceIndex{
	_ListResponseResultItemPriceIndexesPriceIndexName[0:13]:  ListResponseResultItemPriceIndexesPriceIndexWITHOUTINDEX,
	_ListResponseResultItemPriceIndexesPriceIndexName[13:19]: ListResponseResultItemPriceIndexesPriceIndexPROFIT,
	_ListResponseResultItemPriceIndexesPriceIndexName[19:29]: ListResponseResultItemPriceIndexesPriceIndexAVGPROFIT,
	_ListResponseResultItemPriceIndexesPriceIndexName[29:39]: ListResponseResultItemPriceIndexesPriceIndexNONPROFIT,
}

// ParseListResponseResultItemPriceIndexesPriceIndex attempts to convert a string to a ListResponseResultItemPriceIndexesPriceIndex.
func ParseListResponseResultItemPriceIndexesPriceIndex(name string) (ListResponseResultItemPriceIndexesPriceIndex, error) {
	if x, ok := _ListResponseResultItemPriceIndexesPriceIndexValue[name]; ok {
		return x, nil
	}
	return ListResponseResultItemPriceIndexesPriceIndex(0), fmt.Errorf("%s is %w", name, ErrInvalidListResponseResultItemPriceIndexesPriceIndex)
}

// MarshalText implements the text marshaller method.
func (x ListResponseResultItemPriceIndexesPriceIndex) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ListResponseResultItemPriceIndexesPriceIndex) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseListResponseResultItemPriceIndexesPriceIndex(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
