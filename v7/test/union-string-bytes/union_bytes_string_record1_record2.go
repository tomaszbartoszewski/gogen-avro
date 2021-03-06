// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
)

type UnionBytesStringRecord1Record2TypeEnum int

const (
	UnionBytesStringRecord1Record2TypeEnumBytes UnionBytesStringRecord1Record2TypeEnum = 0

	UnionBytesStringRecord1Record2TypeEnumString UnionBytesStringRecord1Record2TypeEnum = 1

	UnionBytesStringRecord1Record2TypeEnumRecord1 UnionBytesStringRecord1Record2TypeEnum = 2

	UnionBytesStringRecord1Record2TypeEnumRecord2 UnionBytesStringRecord1Record2TypeEnum = 3
)

type UnionBytesStringRecord1Record2 struct {
	Bytes     []byte
	String    string
	Record1   *Record1
	Record2   *Record2
	UnionType UnionBytesStringRecord1Record2TypeEnum
}

func writeUnionBytesStringRecord1Record2(r *UnionBytesStringRecord1Record2, w io.Writer) error {

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionBytesStringRecord1Record2TypeEnumBytes:
		return vm.WriteBytes(r.Bytes, w)
	case UnionBytesStringRecord1Record2TypeEnumString:
		return vm.WriteString(r.String, w)
	case UnionBytesStringRecord1Record2TypeEnumRecord1:
		return writeRecord1(r.Record1, w)
	case UnionBytesStringRecord1Record2TypeEnumRecord2:
		return writeRecord2(r.Record2, w)
	}
	return fmt.Errorf("invalid value for *UnionBytesStringRecord1Record2")
}

func NewUnionBytesStringRecord1Record2() *UnionBytesStringRecord1Record2 {
	return &UnionBytesStringRecord1Record2{}
}

func (_ *UnionBytesStringRecord1Record2) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionBytesStringRecord1Record2) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionBytesStringRecord1Record2) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionBytesStringRecord1Record2) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionBytesStringRecord1Record2) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionBytesStringRecord1Record2) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionBytesStringRecord1Record2) SetLong(v int64) {
	r.UnionType = (UnionBytesStringRecord1Record2TypeEnum)(v)
}
func (r *UnionBytesStringRecord1Record2) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Bytes{Target: (&r.Bytes)}
	case 1:
		return &types.String{Target: (&r.String)}
	case 2:
		r.Record1 = NewRecord1()
		return r.Record1
	case 3:
		r.Record2 = NewRecord2()
		return r.Record2
	}
	panic("Unknown field index")
}
func (_ *UnionBytesStringRecord1Record2) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionBytesStringRecord1Record2) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionBytesStringRecord1Record2) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionBytesStringRecord1Record2) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionBytesStringRecord1Record2) Finalize()                {}

func (r *UnionBytesStringRecord1Record2) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionBytesStringRecord1Record2TypeEnumBytes:
		return json.Marshal(map[string]interface{}{"bytes": r.Bytes})
	case UnionBytesStringRecord1Record2TypeEnumString:
		return json.Marshal(map[string]interface{}{"string": r.String})
	case UnionBytesStringRecord1Record2TypeEnumRecord1:
		return json.Marshal(map[string]interface{}{"record1": r.Record1})
	case UnionBytesStringRecord1Record2TypeEnumRecord2:
		return json.Marshal(map[string]interface{}{"record2": r.Record2})
	}
	return nil, fmt.Errorf("invalid value for *UnionBytesStringRecord1Record2")
}

func (r *UnionBytesStringRecord1Record2) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["bytes"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.Bytes)
	}
	if value, ok := fields["string"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.String)
	}
	if value, ok := fields["record1"]; ok {
		r.UnionType = 2
		return json.Unmarshal([]byte(value), &r.Record1)
	}
	if value, ok := fields["record2"]; ok {
		r.UnionType = 3
		return json.Unmarshal([]byte(value), &r.Record2)
	}
	return fmt.Errorf("invalid value for *UnionBytesStringRecord1Record2")
}
