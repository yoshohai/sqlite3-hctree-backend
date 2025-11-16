package sqlite3_backend

/*
#cgo CFLAGS: -std=gnu99

#cgo CFLAGS: -DHAVE_USLEEP=1

#cgo CFLAGS: -DSQLITE_ALLOW_URI_AUTHORITY
#cgo CFLAGS: -DSQLITE_ENABLE_DBPAGE_VTAB
#cgo CFLAGS: -DSQLITE_ENABLE_DBSTAT_VTAB
#cgo CFLAGS: -DSQLITE_ENABLE_BYTECODE_VTAB

#cgo CFLAGS: -DSQLITE_OMIT_DECLTYPE
#cgo CFLAGS: -DSQLITE_OMIT_SHARED_CACHE
#cgo CFLAGS: -DSQLITE_OMIT_DEPRECATED

#cgo LDFLAGS: -lm

#include "sqlite3.h"
*/
import "C"
import "unsafe"

type Backend struct{}

// Database operations

func (b *Backend) OpenReadWrite() int {
	return C.SQLITE_OPEN_READWRITE
}

func (b *Backend) OpenCreate() int {
	return C.SQLITE_OPEN_CREATE
}

func (b *Backend) OpenNoMutex() int {
	return C.SQLITE_OPEN_NOMUTEX
}

func (b *Backend) OpenFullMutex() int {
	return C.SQLITE_OPEN_FULLMUTEX
}

func (b *Backend) OpenURI() int {
	return C.SQLITE_OPEN_URI
}

func (b *Backend) OpenReadOnly() int {
	return C.SQLITE_OPEN_READONLY
}

func (b *Backend) OpenMemory() int {
	return C.SQLITE_OPEN_MEMORY
}

func (b *Backend) OpenExtendedResultCode() int {
	return C.SQLITE_OPEN_EXRESCODE
}

func (b *Backend) OpenV2(filename unsafe.Pointer, ppDb unsafe.Pointer, flags int, zVfs unsafe.Pointer) int {
	return int(C.sqlite3_open_v2(
		(*C.char)(filename),
		(**C.sqlite3)(ppDb),
		C.int(flags),
		(*C.char)(zVfs),
	))
}

func (b *Backend) CloseV2(db unsafe.Pointer) int {
	return int(C.sqlite3_close_v2((*C.sqlite3)(db)))
}

func (b *Backend) Exec(db unsafe.Pointer, sql unsafe.Pointer) int {
	return int(C.sqlite3_exec((*C.sqlite3)(db), (*C.char)(sql), nil, nil, nil))
}

// Statement operations

func (b *Backend) PrepareV2(db unsafe.Pointer, zSql unsafe.Pointer, ppStmt unsafe.Pointer) int {
	return int(C.sqlite3_prepare_v2(
		(*C.sqlite3)(db),
		(*C.char)(zSql),
		-1,
		(**C.sqlite3_stmt)(ppStmt),
		nil,
	))
}

func (b *Backend) Step(stmt unsafe.Pointer) int {
	return int(C.sqlite3_step((*C.sqlite3_stmt)(stmt)))
}

func (b *Backend) Reset(stmt unsafe.Pointer) int {
	return int(C.sqlite3_reset((*C.sqlite3_stmt)(stmt)))
}

func (b *Backend) Finalize(stmt unsafe.Pointer) int {
	return int(C.sqlite3_finalize((*C.sqlite3_stmt)(stmt)))
}

func (b *Backend) BindInt64(stmt unsafe.Pointer, index int, value int64) int {
	return int(C.sqlite3_bind_int64(
		(*C.sqlite3_stmt)(stmt),
		C.int(index),
		C.sqlite3_int64(value),
	))
}

func (b *Backend) BindDouble(stmt unsafe.Pointer, index int, value float64) int {
	return int(C.sqlite3_bind_double(
		(*C.sqlite3_stmt)(stmt),
		C.int(index),
		C.double(value),
	))
}

func (b *Backend) BindText(stmt unsafe.Pointer, index int, value unsafe.Pointer, n int) int {
	return int(C.sqlite3_bind_text(
		(*C.sqlite3_stmt)(stmt),
		C.int(index),
		(*C.char)(value),
		C.int(n),
		C.SQLITE_TRANSIENT,
	))
}

func (b *Backend) BindNull(stmt unsafe.Pointer, index int) int {
	return int(C.sqlite3_bind_null((*C.sqlite3_stmt)(stmt), C.int(index)))
}

// Column operations

func (b *Backend) ColumnCount(stmt unsafe.Pointer) int {
	return int(C.sqlite3_column_count((*C.sqlite3_stmt)(stmt)))
}

func (b *Backend) ColumnName(stmt unsafe.Pointer, i int) string {
	return C.GoString(C.sqlite3_column_name((*C.sqlite3_stmt)(stmt), C.int(i)))
}

func (b *Backend) ColumnType(stmt unsafe.Pointer, i int) int {
	return int(C.sqlite3_column_type((*C.sqlite3_stmt)(stmt), C.int(i)))
}

func (b *Backend) ColumnDouble(stmt unsafe.Pointer, i int) float64 {
	return float64(C.sqlite3_column_double((*C.sqlite3_stmt)(stmt), C.int(i)))
}

func (b *Backend) ColumnInt64(stmt unsafe.Pointer, i int) int64 {
	return int64(C.sqlite3_column_int64((*C.sqlite3_stmt)(stmt), C.int(i)))
}

func (b *Backend) ColumnText(stmt unsafe.Pointer, i int) string {
	cStmt := (*C.sqlite3_stmt)(stmt)
	n := int(C.sqlite3_column_bytes(cStmt, C.int(i)))
	cText := C.sqlite3_column_text(cStmt, C.int(i))
	return C.GoStringN((*C.char)(unsafe.Pointer(cText)), C.int(n))
}

func (b *Backend) ColumnBytes(stmt unsafe.Pointer, i int) int {
	return int(C.sqlite3_column_bytes((*C.sqlite3_stmt)(stmt), C.int(i)))
}

// Error operations

func (b *Backend) ErrMsg(db unsafe.Pointer) string {
	return C.GoString(C.sqlite3_errmsg((*C.sqlite3)(db)))
}

func (b *Backend) ErrStr(rc int) string {
	return C.GoString(C.sqlite3_errstr(C.int(rc)))
}

// Type conversions

func (b *Backend) CharPtr(p unsafe.Pointer) unsafe.Pointer {
	return p
}

func (b *Backend) StringData(s string) unsafe.Pointer {
	return unsafe.Pointer(unsafe.StringData(s))
}

// Constants

func (b *Backend) ResultOk() int {
	return C.SQLITE_OK
}

func (b *Backend) ResultRow() int {
	return C.SQLITE_ROW
}

func (b *Backend) ResultDone() int {
	return C.SQLITE_DONE
}

func (b *Backend) ResultNull() int {
	return C.SQLITE_NULL
}
