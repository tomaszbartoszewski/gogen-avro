// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type Record2 struct {
	Intfield int32 `json:"intfield"`
}

const Record2AvroCRC64Fingerprint = "S\n_\xab\xc7-\xf9\x9c"

func NewRecord2() *Record2 {
	return &Record2{}
}

func DeserializeRecord2(r io.Reader) (*Record2, error) {
	t := NewRecord2()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeRecord2FromSchema(r io.Reader, schema string) (*Record2, error) {
	t := NewRecord2()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeRecord2(r *Record2, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Intfield, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Record2) Serialize(w io.Writer) error {
	return writeRecord2(r, w)
}

func (r *Record2) Schema() string {
	return "{\"fields\":[{\"name\":\"intfield\",\"type\":\"int\"}],\"name\":\"record2\",\"type\":\"record\"}"
}

func (r *Record2) SchemaName() string {
	return "record2"
}

func (_ *Record2) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Record2) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Record2) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Record2) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Record2) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Record2) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Record2) SetString(v string)   { panic("Unsupported operation") }
func (_ *Record2) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Record2) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Int{Target: &r.Intfield}
	}
	panic("Unknown field index")
}

func (r *Record2) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Record2) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *Record2) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Record2) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Record2) Finalize()                        {}

func (_ *Record2) AvroCRC64Fingerprint() []byte {
	return []byte(Record2AvroCRC64Fingerprint)
}
