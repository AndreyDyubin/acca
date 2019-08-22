// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package services

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type clientTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("acca").
func (v *clientTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("clients").
func (v *clientTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *clientTableType) Columns() []string {
	return []string{"client_id", "access_token", "created_at"}
}

// NewStruct makes a new struct for that view or table.
func (v *clientTableType) NewStruct() reform.Struct {
	return new(Client)
}

// NewRecord makes a new record for that table.
func (v *clientTableType) NewRecord() reform.Record {
	return new(Client)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *clientTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ClientTable represents clients view or table in SQL database.
var ClientTable = &clientTableType{
	s: parse.StructInfo{Type: "Client", SQLSchema: "acca", SQLName: "clients", Fields: []parse.FieldInfo{{Name: "ClientID", Type: "int64", Column: "client_id"}, {Name: "AccessToken", Type: "string", Column: "access_token"}, {Name: "CreatedAt", Type: "time.Time", Column: "created_at"}}, PKFieldIndex: 0},
	z: new(Client).Values(),
}

// String returns a string representation of this struct or record.
func (s Client) String() string {
	res := make([]string, 3)
	res[0] = "ClientID: " + reform.Inspect(s.ClientID, true)
	res[1] = "AccessToken: " + reform.Inspect(s.AccessToken, true)
	res[2] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Client) Values() []interface{} {
	return []interface{}{
		s.ClientID,
		s.AccessToken,
		s.CreatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Client) Pointers() []interface{} {
	return []interface{}{
		&s.ClientID,
		&s.AccessToken,
		&s.CreatedAt,
	}
}

// View returns View object for that struct.
func (s *Client) View() reform.View {
	return ClientTable
}

// Table returns Table object for that record.
func (s *Client) Table() reform.Table {
	return ClientTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Client) PKValue() interface{} {
	return s.ClientID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Client) PKPointer() interface{} {
	return &s.ClientID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Client) HasPK() bool {
	return s.ClientID != ClientTable.z[ClientTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Client) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ClientID = int64(i64)
	} else {
		s.ClientID = pk.(int64)
	}
}

// check interfaces
var (
	_ reform.View   = ClientTable
	_ reform.Struct = (*Client)(nil)
	_ reform.Table  = ClientTable
	_ reform.Record = (*Client)(nil)
	_ fmt.Stringer  = (*Client)(nil)
)

func init() {
	parse.AssertUpToDate(&ClientTable.s, new(Client))
}