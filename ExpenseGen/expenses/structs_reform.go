// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package expenses

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type expenseTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *expenseTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("ExpenseTable").
func (v *expenseTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *expenseTableType) Columns() []string {
	return []string{"id", "description", "typeofaccount", "amount", "created_on", "updated_on"}
}

// NewStruct makes a new struct for that view or table.
func (v *expenseTableType) NewStruct() reform.Struct {
	return new(Expense)
}

// NewRecord makes a new record for that table.
func (v *expenseTableType) NewRecord() reform.Record {
	return new(Expense)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *expenseTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ExpenseTable represents ExpenseTable view or table in SQL database.
var ExpenseTable = &expenseTableType{
	s: parse.StructInfo{Type: "Expense", SQLSchema: "", SQLName: "ExpenseTable", Fields: []parse.FieldInfo{{Name: "Id", Type: "int", Column: "id"}, {Name: "Description", Type: "string", Column: "description"}, {Name: "Type", Type: "string", Column: "typeofaccount"}, {Name: "Amount", Type: "float64", Column: "amount"}, {Name: "CreatedOn", Type: "time.Time", Column: "created_on"}, {Name: "UpdatedOn", Type: "time.Time", Column: "updated_on"}}, PKFieldIndex: 0},
	z: new(Expense).Values(),
}

// String returns a string representation of this struct or record.
func (s Expense) String() string {
	res := make([]string, 6)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "Description: " + reform.Inspect(s.Description, true)
	res[2] = "Type: " + reform.Inspect(s.Type, true)
	res[3] = "Amount: " + reform.Inspect(s.Amount, true)
	res[4] = "CreatedOn: " + reform.Inspect(s.CreatedOn, true)
	res[5] = "UpdatedOn: " + reform.Inspect(s.UpdatedOn, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Expense) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.Description,
		s.Type,
		s.Amount,
		s.CreatedOn,
		s.UpdatedOn,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Expense) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.Description,
		&s.Type,
		&s.Amount,
		&s.CreatedOn,
		&s.UpdatedOn,
	}
}

// View returns View object for that struct.
func (s *Expense) View() reform.View {
	return ExpenseTable
}

// Table returns Table object for that record.
func (s *Expense) Table() reform.Table {
	return ExpenseTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Expense) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Expense) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Expense) HasPK() bool {
	return s.Id != ExpenseTable.z[ExpenseTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Expense) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.Id = int(i64)
	} else {
		s.Id = pk.(int)
	}
}

// check interfaces
var (
	_ reform.View   = ExpenseTable
	_ reform.Struct = (*Expense)(nil)
	_ reform.Table  = ExpenseTable
	_ reform.Record = (*Expense)(nil)
	_ fmt.Stringer  = (*Expense)(nil)
)

func init() {
	parse.AssertUpToDate(&ExpenseTable.s, new(Expense))
}
