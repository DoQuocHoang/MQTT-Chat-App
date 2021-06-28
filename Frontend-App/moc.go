package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoFreePacked(ptr unsafe.Pointer) { std_core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_Moc_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return []byte(gs)
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type Person_ITF interface {
	std_core.QObject_ITF
	Person_PTR() *Person
}

func (ptr *Person) Person_PTR() *Person {
	return ptr
}

func (ptr *Person) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Person) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromPerson(ptr Person_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Person_PTR().Pointer()
	}
	return nil
}

func NewPersonFromPointer(ptr unsafe.Pointer) (n *Person) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Person)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Person:
			n = deduced

		case *std_core.QObject:
			n = &Person{QObject: *deduced}

		default:
			n = new(Person)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackPerson3419f1_Constructor
func callbackPerson3419f1_Constructor(ptr unsafe.Pointer) {
	this := NewPersonFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackPerson3419f1_Id
func callbackPerson3419f1_Id(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "id"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewPersonFromPointer(ptr).IdDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Person) ConnectId(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "id"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "id", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "id", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Person) DisconnectId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "id")
	}
}

func (ptr *Person) Id() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Person3419f1_Id(ptr.Pointer()))
	}
	return ""
}

func (ptr *Person) IdDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Person3419f1_IdDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackPerson3419f1_SetId
func callbackPerson3419f1_SetId(ptr unsafe.Pointer, id C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setId"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(id))
	} else {
		NewPersonFromPointer(ptr).SetIdDefault(cGoUnpackString(id))
	}
}

func (ptr *Person) ConnectSetId(f func(id string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setId"); signal != nil {
			f := func(id string) {
				(*(*func(string))(signal))(id)
				f(id)
			}
			qt.ConnectSignal(ptr.Pointer(), "setId", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setId", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Person) DisconnectSetId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setId")
	}
}

func (ptr *Person) SetId(id string) {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		C.Person3419f1_SetId(ptr.Pointer(), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))})
	}
}

func (ptr *Person) SetIdDefault(id string) {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		C.Person3419f1_SetIdDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))})
	}
}

//export callbackPerson3419f1_IdChanged
func callbackPerson3419f1_IdChanged(ptr unsafe.Pointer, id C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "idChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(id))
	}

}

func (ptr *Person) ConnectIdChanged(f func(id string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "idChanged") {
			C.Person3419f1_ConnectIdChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "idChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "idChanged"); signal != nil {
			f := func(id string) {
				(*(*func(string))(signal))(id)
				f(id)
			}
			qt.ConnectSignal(ptr.Pointer(), "idChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "idChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Person) DisconnectIdChanged() {
	if ptr.Pointer() != nil {
		C.Person3419f1_DisconnectIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "idChanged")
	}
}

func (ptr *Person) IdChanged(id string) {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		C.Person3419f1_IdChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))})
	}
}

//export callbackPerson3419f1_FirstName
func callbackPerson3419f1_FirstName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "firstName"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewPersonFromPointer(ptr).FirstNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Person) ConnectFirstName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "firstName"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "firstName", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstName", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Person) DisconnectFirstName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "firstName")
	}
}

func (ptr *Person) FirstName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Person3419f1_FirstName(ptr.Pointer()))
	}
	return ""
}

func (ptr *Person) FirstNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Person3419f1_FirstNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackPerson3419f1_SetFirstName
func callbackPerson3419f1_SetFirstName(ptr unsafe.Pointer, firstName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setFirstName"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(firstName))
	} else {
		NewPersonFromPointer(ptr).SetFirstNameDefault(cGoUnpackString(firstName))
	}
}

func (ptr *Person) ConnectSetFirstName(f func(firstName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFirstName"); signal != nil {
			f := func(firstName string) {
				(*(*func(string))(signal))(firstName)
				f(firstName)
			}
			qt.ConnectSignal(ptr.Pointer(), "setFirstName", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFirstName", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Person) DisconnectSetFirstName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFirstName")
	}
}

func (ptr *Person) SetFirstName(firstName string) {
	if ptr.Pointer() != nil {
		var firstNameC *C.char
		if firstName != "" {
			firstNameC = C.CString(firstName)
			defer C.free(unsafe.Pointer(firstNameC))
		}
		C.Person3419f1_SetFirstName(ptr.Pointer(), C.struct_Moc_PackedString{data: firstNameC, len: C.longlong(len(firstName))})
	}
}

func (ptr *Person) SetFirstNameDefault(firstName string) {
	if ptr.Pointer() != nil {
		var firstNameC *C.char
		if firstName != "" {
			firstNameC = C.CString(firstName)
			defer C.free(unsafe.Pointer(firstNameC))
		}
		C.Person3419f1_SetFirstNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: firstNameC, len: C.longlong(len(firstName))})
	}
}

//export callbackPerson3419f1_FirstNameChanged
func callbackPerson3419f1_FirstNameChanged(ptr unsafe.Pointer, firstName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "firstNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(firstName))
	}

}

func (ptr *Person) ConnectFirstNameChanged(f func(firstName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstNameChanged") {
			C.Person3419f1_ConnectFirstNameChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "firstNameChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstNameChanged"); signal != nil {
			f := func(firstName string) {
				(*(*func(string))(signal))(firstName)
				f(firstName)
			}
			qt.ConnectSignal(ptr.Pointer(), "firstNameChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstNameChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Person) DisconnectFirstNameChanged() {
	if ptr.Pointer() != nil {
		C.Person3419f1_DisconnectFirstNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstNameChanged")
	}
}

func (ptr *Person) FirstNameChanged(firstName string) {
	if ptr.Pointer() != nil {
		var firstNameC *C.char
		if firstName != "" {
			firstNameC = C.CString(firstName)
			defer C.free(unsafe.Pointer(firstNameC))
		}
		C.Person3419f1_FirstNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: firstNameC, len: C.longlong(len(firstName))})
	}
}

//export callbackPerson3419f1_LastName
func callbackPerson3419f1_LastName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "lastName"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewPersonFromPointer(ptr).LastNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Person) ConnectLastName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lastName"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lastName", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastName", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Person) DisconnectLastName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lastName")
	}
}

func (ptr *Person) LastName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Person3419f1_LastName(ptr.Pointer()))
	}
	return ""
}

func (ptr *Person) LastNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Person3419f1_LastNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackPerson3419f1_SetLastName
func callbackPerson3419f1_SetLastName(ptr unsafe.Pointer, lastName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setLastName"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(lastName))
	} else {
		NewPersonFromPointer(ptr).SetLastNameDefault(cGoUnpackString(lastName))
	}
}

func (ptr *Person) ConnectSetLastName(f func(lastName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLastName"); signal != nil {
			f := func(lastName string) {
				(*(*func(string))(signal))(lastName)
				f(lastName)
			}
			qt.ConnectSignal(ptr.Pointer(), "setLastName", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLastName", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Person) DisconnectSetLastName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLastName")
	}
}

func (ptr *Person) SetLastName(lastName string) {
	if ptr.Pointer() != nil {
		var lastNameC *C.char
		if lastName != "" {
			lastNameC = C.CString(lastName)
			defer C.free(unsafe.Pointer(lastNameC))
		}
		C.Person3419f1_SetLastName(ptr.Pointer(), C.struct_Moc_PackedString{data: lastNameC, len: C.longlong(len(lastName))})
	}
}

func (ptr *Person) SetLastNameDefault(lastName string) {
	if ptr.Pointer() != nil {
		var lastNameC *C.char
		if lastName != "" {
			lastNameC = C.CString(lastName)
			defer C.free(unsafe.Pointer(lastNameC))
		}
		C.Person3419f1_SetLastNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: lastNameC, len: C.longlong(len(lastName))})
	}
}

//export callbackPerson3419f1_LastNameChanged
func callbackPerson3419f1_LastNameChanged(ptr unsafe.Pointer, lastName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "lastNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(lastName))
	}

}

func (ptr *Person) ConnectLastNameChanged(f func(lastName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lastNameChanged") {
			C.Person3419f1_ConnectLastNameChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "lastNameChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lastNameChanged"); signal != nil {
			f := func(lastName string) {
				(*(*func(string))(signal))(lastName)
				f(lastName)
			}
			qt.ConnectSignal(ptr.Pointer(), "lastNameChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastNameChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Person) DisconnectLastNameChanged() {
	if ptr.Pointer() != nil {
		C.Person3419f1_DisconnectLastNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lastNameChanged")
	}
}

func (ptr *Person) LastNameChanged(lastName string) {
	if ptr.Pointer() != nil {
		var lastNameC *C.char
		if lastName != "" {
			lastNameC = C.CString(lastName)
			defer C.free(unsafe.Pointer(lastNameC))
		}
		C.Person3419f1_LastNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: lastNameC, len: C.longlong(len(lastName))})
	}
}

func Person_QRegisterMetaType() int {
	return int(int32(C.Person3419f1_Person3419f1_QRegisterMetaType()))
}

func (ptr *Person) QRegisterMetaType() int {
	return int(int32(C.Person3419f1_Person3419f1_QRegisterMetaType()))
}

func Person_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Person3419f1_Person3419f1_QRegisterMetaType2(typeNameC)))
}

func (ptr *Person) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Person3419f1_Person3419f1_QRegisterMetaType2(typeNameC)))
}

func Person_QmlRegisterType() int {
	return int(int32(C.Person3419f1_Person3419f1_QmlRegisterType()))
}

func (ptr *Person) QmlRegisterType() int {
	return int(int32(C.Person3419f1_Person3419f1_QmlRegisterType()))
}

func Person_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Person3419f1_Person3419f1_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Person) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Person3419f1_Person3419f1_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func Person_QmlRegisterUncreatableType(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	var messageC *C.char
	if message != "" {
		messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
	}
	return int(int32(C.Person3419f1_Person3419f1_QmlRegisterUncreatableType(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})))
}

func (ptr *Person) QmlRegisterUncreatableType(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	var messageC *C.char
	if message != "" {
		messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
	}
	return int(int32(C.Person3419f1_Person3419f1_QmlRegisterUncreatableType(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})))
}

func (ptr *Person) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Person3419f1___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Person) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Person3419f1___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Person) __children_newList() unsafe.Pointer {
	return C.Person3419f1___children_newList(ptr.Pointer())
}

func (ptr *Person) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Person3419f1___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Person) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Person3419f1___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Person) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Person3419f1___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Person) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Person3419f1___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Person) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Person3419f1___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Person) __findChildren_newList() unsafe.Pointer {
	return C.Person3419f1___findChildren_newList(ptr.Pointer())
}

func (ptr *Person) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Person3419f1___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Person) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Person3419f1___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Person) __findChildren_newList3() unsafe.Pointer {
	return C.Person3419f1___findChildren_newList3(ptr.Pointer())
}

func NewPerson(parent std_core.QObject_ITF) *Person {
	Person_QRegisterMetaType()
	tmpValue := NewPersonFromPointer(C.Person3419f1_NewPerson(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackPerson3419f1_DestroyPerson
func callbackPerson3419f1_DestroyPerson(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Person"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPersonFromPointer(ptr).DestroyPersonDefault()
	}
}

func (ptr *Person) ConnectDestroyPerson(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Person"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~Person", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Person", unsafe.Pointer(&f))
		}
	}
}

func (ptr *Person) DisconnectDestroyPerson() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Person")
	}
}

func (ptr *Person) DestroyPerson() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.Person3419f1_DestroyPerson(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *Person) DestroyPersonDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.Person3419f1_DestroyPersonDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackPerson3419f1_ChildEvent
func callbackPerson3419f1_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewPersonFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Person) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Person3419f1_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackPerson3419f1_ConnectNotify
func callbackPerson3419f1_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPersonFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Person) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Person3419f1_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPerson3419f1_CustomEvent
func callbackPerson3419f1_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewPersonFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Person) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Person3419f1_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackPerson3419f1_DeleteLater
func callbackPerson3419f1_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPersonFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Person) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.Person3419f1_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackPerson3419f1_Destroyed
func callbackPerson3419f1_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackPerson3419f1_DisconnectNotify
func callbackPerson3419f1_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPersonFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Person) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Person3419f1_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPerson3419f1_Event
func callbackPerson3419f1_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Person) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Person3419f1_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackPerson3419f1_EventFilter
func callbackPerson3419f1_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Person) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Person3419f1_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackPerson3419f1_ObjectNameChanged
func callbackPerson3419f1_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackPerson3419f1_TimerEvent
func callbackPerson3419f1_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewPersonFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Person) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Person3419f1_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type PersonModel_ITF interface {
	std_core.QAbstractListModel_ITF
	PersonModel_PTR() *PersonModel
}

func (ptr *PersonModel) PersonModel_PTR() *PersonModel {
	return ptr
}

func (ptr *PersonModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *PersonModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromPersonModel(ptr PersonModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.PersonModel_PTR().Pointer()
	}
	return nil
}

func NewPersonModelFromPointer(ptr unsafe.Pointer) (n *PersonModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(PersonModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *PersonModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &PersonModel{QAbstractListModel: *deduced}

		default:
			n = new(PersonModel)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *PersonModel) Init() { this.init() }

//export callbackPersonModel3419f1_Constructor
func callbackPersonModel3419f1_Constructor(ptr unsafe.Pointer) {
	this := NewPersonModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackPersonModel3419f1_AddPerson
func callbackPersonModel3419f1_AddPerson(ptr unsafe.Pointer, v0 unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addPerson"); signal != nil {
		(*(*func(*Person))(signal))(NewPersonFromPointer(v0))
	}

}

func (ptr *PersonModel) ConnectAddPerson(f func(v0 *Person)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addPerson"); signal != nil {
			f := func(v0 *Person) {
				(*(*func(*Person))(signal))(v0)
				f(v0)
			}
			qt.ConnectSignal(ptr.Pointer(), "addPerson", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addPerson", unsafe.Pointer(&f))
		}
	}
}

func (ptr *PersonModel) DisconnectAddPerson() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addPerson")
	}
}

func (ptr *PersonModel) AddPerson(v0 Person_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_AddPerson(ptr.Pointer(), PointerFromPerson(v0))
	}
}

//export callbackPersonModel3419f1_EditPerson
func callbackPersonModel3419f1_EditPerson(ptr unsafe.Pointer, row C.int, firstName C.struct_Moc_PackedString, lastName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "editPerson"); signal != nil {
		(*(*func(int, string, string))(signal))(int(int32(row)), cGoUnpackString(firstName), cGoUnpackString(lastName))
	}

}

func (ptr *PersonModel) ConnectEditPerson(f func(row int, firstName string, lastName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "editPerson"); signal != nil {
			f := func(row int, firstName string, lastName string) {
				(*(*func(int, string, string))(signal))(row, firstName, lastName)
				f(row, firstName, lastName)
			}
			qt.ConnectSignal(ptr.Pointer(), "editPerson", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "editPerson", unsafe.Pointer(&f))
		}
	}
}

func (ptr *PersonModel) DisconnectEditPerson() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "editPerson")
	}
}

func (ptr *PersonModel) EditPerson(row int, firstName string, lastName string) {
	if ptr.Pointer() != nil {
		var firstNameC *C.char
		if firstName != "" {
			firstNameC = C.CString(firstName)
			defer C.free(unsafe.Pointer(firstNameC))
		}
		var lastNameC *C.char
		if lastName != "" {
			lastNameC = C.CString(lastName)
			defer C.free(unsafe.Pointer(lastNameC))
		}
		C.PersonModel3419f1_EditPerson(ptr.Pointer(), C.int(int32(row)), C.struct_Moc_PackedString{data: firstNameC, len: C.longlong(len(firstName))}, C.struct_Moc_PackedString{data: lastNameC, len: C.longlong(len(lastName))})
	}
}

//export callbackPersonModel3419f1_RemovePerson
func callbackPersonModel3419f1_RemovePerson(ptr unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "removePerson"); signal != nil {
		(*(*func(int))(signal))(int(int32(row)))
	}

}

func (ptr *PersonModel) ConnectRemovePerson(f func(row int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removePerson"); signal != nil {
			f := func(row int) {
				(*(*func(int))(signal))(row)
				f(row)
			}
			qt.ConnectSignal(ptr.Pointer(), "removePerson", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removePerson", unsafe.Pointer(&f))
		}
	}
}

func (ptr *PersonModel) DisconnectRemovePerson() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removePerson")
	}
}

func (ptr *PersonModel) RemovePerson(row int) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_RemovePerson(ptr.Pointer(), C.int(int32(row)))
	}
}

//export callbackPersonModel3419f1_Roles
func callbackPersonModel3419f1_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__roles_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__roles_newList())
		for k, v := range NewPersonModelFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *PersonModel) ConnectRoles(f func() map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "roles"); signal != nil {
			f := func() map[int]*std_core.QByteArray {
				(*(*func() map[int]*std_core.QByteArray)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "roles", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "roles", unsafe.Pointer(&f))
		}
	}
}

func (ptr *PersonModel) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *PersonModel) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.PersonModel3419f1_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *PersonModel) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.PersonModel3419f1_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackPersonModel3419f1_SetRoles
func callbackPersonModel3419f1_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	} else {
		NewPersonModelFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	}
}

func (ptr *PersonModel) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRoles"); signal != nil {
			f := func(roles map[int]*std_core.QByteArray) {
				(*(*func(map[int]*std_core.QByteArray))(signal))(roles)
				f(roles)
			}
			qt.ConnectSignal(ptr.Pointer(), "setRoles", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", unsafe.Pointer(&f))
		}
	}
}

func (ptr *PersonModel) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *PersonModel) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *PersonModel) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackPersonModel3419f1_RolesChanged
func callbackPersonModel3419f1_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i, v := range tmpList.__rolesChanged_roles_keyList() {
				out[v] = tmpList.__rolesChanged_roles_atList(v, i)
			}
			return out
		}(roles))
	}

}

func (ptr *PersonModel) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.PersonModel3419f1_ConnectRolesChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "rolesChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rolesChanged"); signal != nil {
			f := func(roles map[int]*std_core.QByteArray) {
				(*(*func(map[int]*std_core.QByteArray))(signal))(roles)
				f(roles)
			}
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *PersonModel) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *PersonModel) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackPersonModel3419f1_People
func callbackPersonModel3419f1_People(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "people"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__people_newList())
			for _, v := range (*(*func() []*Person)(signal))() {
				tmpList.__people_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__people_newList())
		for _, v := range NewPersonModelFromPointer(ptr).PeopleDefault() {
			tmpList.__people_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *PersonModel) ConnectPeople(f func() []*Person) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "people"); signal != nil {
			f := func() []*Person {
				(*(*func() []*Person)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "people", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "people", unsafe.Pointer(&f))
		}
	}
}

func (ptr *PersonModel) DisconnectPeople() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "people")
	}
}

func (ptr *PersonModel) People() []*Person {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*Person {
			out := make([]*Person, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__people_atList(i)
			}
			return out
		}(C.PersonModel3419f1_People(ptr.Pointer()))
	}
	return make([]*Person, 0)
}

func (ptr *PersonModel) PeopleDefault() []*Person {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*Person {
			out := make([]*Person, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__people_atList(i)
			}
			return out
		}(C.PersonModel3419f1_PeopleDefault(ptr.Pointer()))
	}
	return make([]*Person, 0)
}

//export callbackPersonModel3419f1_SetPeople
func callbackPersonModel3419f1_SetPeople(ptr unsafe.Pointer, people C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setPeople"); signal != nil {
		(*(*func([]*Person))(signal))(func(l C.struct_Moc_PackedList) []*Person {
			out := make([]*Person, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setPeople_people_atList(i)
			}
			return out
		}(people))
	} else {
		NewPersonModelFromPointer(ptr).SetPeopleDefault(func(l C.struct_Moc_PackedList) []*Person {
			out := make([]*Person, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setPeople_people_atList(i)
			}
			return out
		}(people))
	}
}

func (ptr *PersonModel) ConnectSetPeople(f func(people []*Person)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPeople"); signal != nil {
			f := func(people []*Person) {
				(*(*func([]*Person))(signal))(people)
				f(people)
			}
			qt.ConnectSignal(ptr.Pointer(), "setPeople", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPeople", unsafe.Pointer(&f))
		}
	}
}

func (ptr *PersonModel) DisconnectSetPeople() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPeople")
	}
}

func (ptr *PersonModel) SetPeople(people []*Person) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_SetPeople(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__setPeople_people_newList())
			for _, v := range people {
				tmpList.__setPeople_people_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *PersonModel) SetPeopleDefault(people []*Person) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_SetPeopleDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__setPeople_people_newList())
			for _, v := range people {
				tmpList.__setPeople_people_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackPersonModel3419f1_PeopleChanged
func callbackPersonModel3419f1_PeopleChanged(ptr unsafe.Pointer, people C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "peopleChanged"); signal != nil {
		(*(*func([]*Person))(signal))(func(l C.struct_Moc_PackedList) []*Person {
			out := make([]*Person, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__peopleChanged_people_atList(i)
			}
			return out
		}(people))
	}

}

func (ptr *PersonModel) ConnectPeopleChanged(f func(people []*Person)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "peopleChanged") {
			C.PersonModel3419f1_ConnectPeopleChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "peopleChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "peopleChanged"); signal != nil {
			f := func(people []*Person) {
				(*(*func([]*Person))(signal))(people)
				f(people)
			}
			qt.ConnectSignal(ptr.Pointer(), "peopleChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "peopleChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *PersonModel) DisconnectPeopleChanged() {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_DisconnectPeopleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "peopleChanged")
	}
}

func (ptr *PersonModel) PeopleChanged(people []*Person) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_PeopleChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__peopleChanged_people_newList())
			for _, v := range people {
				tmpList.__peopleChanged_people_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func PersonModel_QRegisterMetaType() int {
	return int(int32(C.PersonModel3419f1_PersonModel3419f1_QRegisterMetaType()))
}

func (ptr *PersonModel) QRegisterMetaType() int {
	return int(int32(C.PersonModel3419f1_PersonModel3419f1_QRegisterMetaType()))
}

func PersonModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.PersonModel3419f1_PersonModel3419f1_QRegisterMetaType2(typeNameC)))
}

func (ptr *PersonModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.PersonModel3419f1_PersonModel3419f1_QRegisterMetaType2(typeNameC)))
}

func PersonModel_QmlRegisterType() int {
	return int(int32(C.PersonModel3419f1_PersonModel3419f1_QmlRegisterType()))
}

func (ptr *PersonModel) QmlRegisterType() int {
	return int(int32(C.PersonModel3419f1_PersonModel3419f1_QmlRegisterType()))
}

func PersonModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.PersonModel3419f1_PersonModel3419f1_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *PersonModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.PersonModel3419f1_PersonModel3419f1_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func PersonModel_QmlRegisterUncreatableType(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	var messageC *C.char
	if message != "" {
		messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
	}
	return int(int32(C.PersonModel3419f1_PersonModel3419f1_QmlRegisterUncreatableType(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})))
}

func (ptr *PersonModel) QmlRegisterUncreatableType(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	var messageC *C.char
	if message != "" {
		messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
	}
	return int(int32(C.PersonModel3419f1_PersonModel3419f1_QmlRegisterUncreatableType(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})))
}

func (ptr *PersonModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel3419f1_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.PersonModel3419f1_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *PersonModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel3419f1_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.PersonModel3419f1_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *PersonModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel3419f1_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.PersonModel3419f1_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *PersonModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.PersonModel3419f1___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *PersonModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.PersonModel3419f1___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *PersonModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.PersonModel3419f1___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *PersonModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.PersonModel3419f1___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *PersonModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel3419f1___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.PersonModel3419f1___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *PersonModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.PersonModel3419f1___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *PersonModel) __itemData_newList() unsafe.Pointer {
	return C.PersonModel3419f1___itemData_newList(ptr.Pointer())
}

func (ptr *PersonModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.PersonModel3419f1___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *PersonModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.PersonModel3419f1___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *PersonModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.PersonModel3419f1___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *PersonModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.PersonModel3419f1___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *PersonModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.PersonModel3419f1___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *PersonModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.PersonModel3419f1___match_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *PersonModel) __match_newList() unsafe.Pointer {
	return C.PersonModel3419f1___match_newList(ptr.Pointer())
}

func (ptr *PersonModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.PersonModel3419f1___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *PersonModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.PersonModel3419f1___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *PersonModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.PersonModel3419f1___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *PersonModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.PersonModel3419f1___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *PersonModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.PersonModel3419f1___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *PersonModel) __roleNames_newList() unsafe.Pointer {
	return C.PersonModel3419f1___roleNames_newList(ptr.Pointer())
}

func (ptr *PersonModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.PersonModel3419f1___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *PersonModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.PersonModel3419f1___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *PersonModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.PersonModel3419f1___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *PersonModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.PersonModel3419f1___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *PersonModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel3419f1_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.PersonModel3419f1_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *PersonModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel3419f1_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.PersonModel3419f1_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *PersonModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.PersonModel3419f1___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *PersonModel) __children_newList() unsafe.Pointer {
	return C.PersonModel3419f1___children_newList(ptr.Pointer())
}

func (ptr *PersonModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.PersonModel3419f1___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *PersonModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.PersonModel3419f1___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *PersonModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.PersonModel3419f1___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *PersonModel) __findChildren_newList() unsafe.Pointer {
	return C.PersonModel3419f1___findChildren_newList(ptr.Pointer())
}

func (ptr *PersonModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.PersonModel3419f1___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *PersonModel) __findChildren_newList3() unsafe.Pointer {
	return C.PersonModel3419f1___findChildren_newList3(ptr.Pointer())
}

func (ptr *PersonModel) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.PersonModel3419f1___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *PersonModel) __roles_newList() unsafe.Pointer {
	return C.PersonModel3419f1___roles_newList(ptr.Pointer())
}

func (ptr *PersonModel) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roles_keyList_atList(i)
			}
			return out
		}(C.PersonModel3419f1___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *PersonModel) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.PersonModel3419f1___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *PersonModel) __setRoles_roles_newList() unsafe.Pointer {
	return C.PersonModel3419f1___setRoles_roles_newList(ptr.Pointer())
}

func (ptr *PersonModel) __setRoles_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoles_roles_keyList_atList(i)
			}
			return out
		}(C.PersonModel3419f1___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *PersonModel) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.PersonModel3419f1___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *PersonModel) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.PersonModel3419f1___rolesChanged_roles_newList(ptr.Pointer())
}

func (ptr *PersonModel) __rolesChanged_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____rolesChanged_roles_keyList_atList(i)
			}
			return out
		}(C.PersonModel3419f1___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *PersonModel) __people_atList(i int) *Person {
	if ptr.Pointer() != nil {
		tmpValue := NewPersonFromPointer(C.PersonModel3419f1___people_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __people_setList(i Person_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___people_setList(ptr.Pointer(), PointerFromPerson(i))
	}
}

func (ptr *PersonModel) __people_newList() unsafe.Pointer {
	return C.PersonModel3419f1___people_newList(ptr.Pointer())
}

func (ptr *PersonModel) __setPeople_people_atList(i int) *Person {
	if ptr.Pointer() != nil {
		tmpValue := NewPersonFromPointer(C.PersonModel3419f1___setPeople_people_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __setPeople_people_setList(i Person_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___setPeople_people_setList(ptr.Pointer(), PointerFromPerson(i))
	}
}

func (ptr *PersonModel) __setPeople_people_newList() unsafe.Pointer {
	return C.PersonModel3419f1___setPeople_people_newList(ptr.Pointer())
}

func (ptr *PersonModel) __peopleChanged_people_atList(i int) *Person {
	if ptr.Pointer() != nil {
		tmpValue := NewPersonFromPointer(C.PersonModel3419f1___peopleChanged_people_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __peopleChanged_people_setList(i Person_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1___peopleChanged_people_setList(ptr.Pointer(), PointerFromPerson(i))
	}
}

func (ptr *PersonModel) __peopleChanged_people_newList() unsafe.Pointer {
	return C.PersonModel3419f1___peopleChanged_people_newList(ptr.Pointer())
}

func (ptr *PersonModel) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel3419f1_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____roles_keyList_newList() unsafe.Pointer {
	return C.PersonModel3419f1_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *PersonModel) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel3419f1_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.PersonModel3419f1_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *PersonModel) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel3419f1_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.PersonModel3419f1_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewPersonModel(parent std_core.QObject_ITF) *PersonModel {
	PersonModel_QRegisterMetaType()
	tmpValue := NewPersonModelFromPointer(C.PersonModel3419f1_NewPersonModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackPersonModel3419f1_DestroyPersonModel
func callbackPersonModel3419f1_DestroyPersonModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~PersonModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPersonModelFromPointer(ptr).DestroyPersonModelDefault()
	}
}

func (ptr *PersonModel) ConnectDestroyPersonModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~PersonModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~PersonModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~PersonModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *PersonModel) DisconnectDestroyPersonModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~PersonModel")
	}
}

func (ptr *PersonModel) DestroyPersonModel() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.PersonModel3419f1_DestroyPersonModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *PersonModel) DestroyPersonModelDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.PersonModel3419f1_DestroyPersonModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackPersonModel3419f1_DropMimeData
func callbackPersonModel3419f1_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_Flags
func callbackPersonModel3419f1_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewPersonModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *PersonModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.PersonModel3419f1_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackPersonModel3419f1_Index
func callbackPersonModel3419f1_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewPersonModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *PersonModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.PersonModel3419f1_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		qt.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel3419f1_Sibling
func callbackPersonModel3419f1_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewPersonModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *PersonModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.PersonModel3419f1_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		qt.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel3419f1_Buddy
func callbackPersonModel3419f1_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewPersonModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *PersonModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.PersonModel3419f1_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel3419f1_CanDropMimeData
func callbackPersonModel3419f1_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_CanFetchMore
func callbackPersonModel3419f1_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_ColumnCount
func callbackPersonModel3419f1_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewPersonModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *PersonModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel3419f1_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackPersonModel3419f1_ColumnsAboutToBeInserted
func callbackPersonModel3419f1_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel3419f1_ColumnsAboutToBeMoved
func callbackPersonModel3419f1_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackPersonModel3419f1_ColumnsAboutToBeRemoved
func callbackPersonModel3419f1_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel3419f1_ColumnsInserted
func callbackPersonModel3419f1_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel3419f1_ColumnsMoved
func callbackPersonModel3419f1_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackPersonModel3419f1_ColumnsRemoved
func callbackPersonModel3419f1_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel3419f1_Data
func callbackPersonModel3419f1_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewPersonModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *PersonModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.PersonModel3419f1_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel3419f1_DataChanged
func callbackPersonModel3419f1_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(signal))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackPersonModel3419f1_FetchMore
func callbackPersonModel3419f1_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewPersonModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *PersonModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackPersonModel3419f1_HasChildren
func callbackPersonModel3419f1_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_HeaderData
func callbackPersonModel3419f1_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewPersonModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *PersonModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.PersonModel3419f1_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel3419f1_HeaderDataChanged
func callbackPersonModel3419f1_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel3419f1_InsertColumns
func callbackPersonModel3419f1_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_InsertRows
func callbackPersonModel3419f1_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_ItemData
func callbackPersonModel3419f1_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*std_core.QModelIndex) map[int]*std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__itemData_newList())
		for k, v := range NewPersonModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *PersonModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.PersonModel3419f1_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackPersonModel3419f1_LayoutAboutToBeChanged
func callbackPersonModel3419f1_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackPersonModel3419f1_LayoutChanged
func callbackPersonModel3419f1_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackPersonModel3419f1_Match
func callbackPersonModel3419f1_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__match_newList())
		for _, v := range NewPersonModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *PersonModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.PersonModel3419f1_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackPersonModel3419f1_MimeData
func callbackPersonModel3419f1_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func([]*std_core.QModelIndex) *std_core.QMimeData)(signal))(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewPersonModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewPersonModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *PersonModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.PersonModel3419f1_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackPersonModel3419f1_MimeTypes
func callbackPersonModel3419f1_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewPersonModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *PersonModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.PersonModel3419f1_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackPersonModel3419f1_ModelAboutToBeReset
func callbackPersonModel3419f1_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackPersonModel3419f1_ModelReset
func callbackPersonModel3419f1_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackPersonModel3419f1_MoveColumns
func callbackPersonModel3419f1_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *PersonModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_MoveRows
func callbackPersonModel3419f1_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *PersonModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_Parent
func callbackPersonModel3419f1_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewPersonModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *PersonModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.PersonModel3419f1_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel3419f1_RemoveColumns
func callbackPersonModel3419f1_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_RemoveRows
func callbackPersonModel3419f1_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_ResetInternalData
func callbackPersonModel3419f1_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPersonModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *PersonModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackPersonModel3419f1_Revert
func callbackPersonModel3419f1_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPersonModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *PersonModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_RevertDefault(ptr.Pointer())
	}
}

//export callbackPersonModel3419f1_RoleNames
func callbackPersonModel3419f1_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewPersonModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *PersonModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.PersonModel3419f1_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackPersonModel3419f1_RowCount
func callbackPersonModel3419f1_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewPersonModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *PersonModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel3419f1_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackPersonModel3419f1_RowsAboutToBeInserted
func callbackPersonModel3419f1_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackPersonModel3419f1_RowsAboutToBeMoved
func callbackPersonModel3419f1_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackPersonModel3419f1_RowsAboutToBeRemoved
func callbackPersonModel3419f1_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel3419f1_RowsInserted
func callbackPersonModel3419f1_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel3419f1_RowsMoved
func callbackPersonModel3419f1_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackPersonModel3419f1_RowsRemoved
func callbackPersonModel3419f1_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel3419f1_SetData
func callbackPersonModel3419f1_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *PersonModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_SetHeaderData
func callbackPersonModel3419f1_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *PersonModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_SetItemData
func callbackPersonModel3419f1_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(signal))(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewPersonModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewPersonModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *PersonModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackPersonModel3419f1_Sort
func callbackPersonModel3419f1_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewPersonModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *PersonModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackPersonModel3419f1_Span
func callbackPersonModel3419f1_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewPersonModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *PersonModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.PersonModel3419f1_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		qt.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel3419f1_Submit
func callbackPersonModel3419f1_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *PersonModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackPersonModel3419f1_SupportedDragActions
func callbackPersonModel3419f1_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewPersonModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *PersonModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.PersonModel3419f1_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackPersonModel3419f1_SupportedDropActions
func callbackPersonModel3419f1_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewPersonModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *PersonModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.PersonModel3419f1_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackPersonModel3419f1_ChildEvent
func callbackPersonModel3419f1_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewPersonModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *PersonModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackPersonModel3419f1_ConnectNotify
func callbackPersonModel3419f1_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPersonModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *PersonModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPersonModel3419f1_CustomEvent
func callbackPersonModel3419f1_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewPersonModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *PersonModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackPersonModel3419f1_DeleteLater
func callbackPersonModel3419f1_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPersonModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *PersonModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.PersonModel3419f1_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackPersonModel3419f1_Destroyed
func callbackPersonModel3419f1_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackPersonModel3419f1_DisconnectNotify
func callbackPersonModel3419f1_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPersonModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *PersonModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPersonModel3419f1_Event
func callbackPersonModel3419f1_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *PersonModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_EventFilter
func callbackPersonModel3419f1_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *PersonModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PersonModel3419f1_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackPersonModel3419f1_ObjectNameChanged
func callbackPersonModel3419f1_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackPersonModel3419f1_TimerEvent
func callbackPersonModel3419f1_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewPersonModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *PersonModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel3419f1_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type QmlBridge_ITF interface {
	std_core.QObject_ITF
	QmlBridge_PTR() *QmlBridge
}

func (ptr *QmlBridge) QmlBridge_PTR() *QmlBridge {
	return ptr
}

func (ptr *QmlBridge) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlBridge) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlBridge(ptr QmlBridge_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlBridge_PTR().Pointer()
	}
	return nil
}

func NewQmlBridgeFromPointer(ptr unsafe.Pointer) (n *QmlBridge) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QmlBridge)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QmlBridge:
			n = deduced

		case *std_core.QObject:
			n = &QmlBridge{QObject: *deduced}

		default:
			n = new(QmlBridge)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackQmlBridge3419f1_Constructor
func callbackQmlBridge3419f1_Constructor(ptr unsafe.Pointer) {
	this := NewQmlBridgeFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackQmlBridge3419f1_SendToQml
func callbackQmlBridge3419f1_SendToQml(ptr unsafe.Pointer, source C.struct_Moc_PackedString, action C.struct_Moc_PackedString, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sendToQml"); signal != nil {
		(*(*func(string, string, string))(signal))(cGoUnpackString(source), cGoUnpackString(action), cGoUnpackString(data))
	}

}

func (ptr *QmlBridge) ConnectSendToQml(f func(source string, action string, data string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendToQml") {
			C.QmlBridge3419f1_ConnectSendToQml(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sendToQml")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendToQml"); signal != nil {
			f := func(source string, action string, data string) {
				(*(*func(string, string, string))(signal))(source, action, data)
				f(source, action, data)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendToQml", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendToQml", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectSendToQml() {
	if ptr.Pointer() != nil {
		C.QmlBridge3419f1_DisconnectSendToQml(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendToQml")
	}
}

func (ptr *QmlBridge) SendToQml(source string, action string, data string) {
	if ptr.Pointer() != nil {
		var sourceC *C.char
		if source != "" {
			sourceC = C.CString(source)
			defer C.free(unsafe.Pointer(sourceC))
		}
		var actionC *C.char
		if action != "" {
			actionC = C.CString(action)
			defer C.free(unsafe.Pointer(actionC))
		}
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.QmlBridge3419f1_SendToQml(ptr.Pointer(), C.struct_Moc_PackedString{data: sourceC, len: C.longlong(len(source))}, C.struct_Moc_PackedString{data: actionC, len: C.longlong(len(action))}, C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackQmlBridge3419f1_SendToGo
func callbackQmlBridge3419f1_SendToGo(ptr unsafe.Pointer, source C.struct_Moc_PackedString, action C.struct_Moc_PackedString, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sendToGo"); signal != nil {
		(*(*func(string, string, string))(signal))(cGoUnpackString(source), cGoUnpackString(action), cGoUnpackString(data))
	}

}

func (ptr *QmlBridge) ConnectSendToGo(f func(source string, action string, data string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sendToGo"); signal != nil {
			f := func(source string, action string, data string) {
				(*(*func(string, string, string))(signal))(source, action, data)
				f(source, action, data)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendToGo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendToGo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectSendToGo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sendToGo")
	}
}

func (ptr *QmlBridge) SendToGo(source string, action string, data string) {
	if ptr.Pointer() != nil {
		var sourceC *C.char
		if source != "" {
			sourceC = C.CString(source)
			defer C.free(unsafe.Pointer(sourceC))
		}
		var actionC *C.char
		if action != "" {
			actionC = C.CString(action)
			defer C.free(unsafe.Pointer(actionC))
		}
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.QmlBridge3419f1_SendToGo(ptr.Pointer(), C.struct_Moc_PackedString{data: sourceC, len: C.longlong(len(source))}, C.struct_Moc_PackedString{data: actionC, len: C.longlong(len(action))}, C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackQmlBridge3419f1_RegisterToGo
func callbackQmlBridge3419f1_RegisterToGo(ptr unsafe.Pointer, object unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "registerToGo"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(object))
	}

}

func (ptr *QmlBridge) ConnectRegisterToGo(f func(object *std_core.QObject)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "registerToGo"); signal != nil {
			f := func(object *std_core.QObject) {
				(*(*func(*std_core.QObject))(signal))(object)
				f(object)
			}
			qt.ConnectSignal(ptr.Pointer(), "registerToGo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "registerToGo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectRegisterToGo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "registerToGo")
	}
}

func (ptr *QmlBridge) RegisterToGo(object std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge3419f1_RegisterToGo(ptr.Pointer(), std_core.PointerFromQObject(object))
	}
}

//export callbackQmlBridge3419f1_DeregisterToGo
func callbackQmlBridge3419f1_DeregisterToGo(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "deregisterToGo"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

func (ptr *QmlBridge) ConnectDeregisterToGo(f func(objectName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "deregisterToGo"); signal != nil {
			f := func(objectName string) {
				(*(*func(string))(signal))(objectName)
				f(objectName)
			}
			qt.ConnectSignal(ptr.Pointer(), "deregisterToGo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deregisterToGo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDeregisterToGo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "deregisterToGo")
	}
}

func (ptr *QmlBridge) DeregisterToGo(objectName string) {
	if ptr.Pointer() != nil {
		var objectNameC *C.char
		if objectName != "" {
			objectNameC = C.CString(objectName)
			defer C.free(unsafe.Pointer(objectNameC))
		}
		C.QmlBridge3419f1_DeregisterToGo(ptr.Pointer(), C.struct_Moc_PackedString{data: objectNameC, len: C.longlong(len(objectName))})
	}
}

func QmlBridge_QRegisterMetaType() int {
	return int(int32(C.QmlBridge3419f1_QmlBridge3419f1_QRegisterMetaType()))
}

func (ptr *QmlBridge) QRegisterMetaType() int {
	return int(int32(C.QmlBridge3419f1_QmlBridge3419f1_QRegisterMetaType()))
}

func QmlBridge_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridge3419f1_QmlBridge3419f1_QRegisterMetaType2(typeNameC)))
}

func (ptr *QmlBridge) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridge3419f1_QmlBridge3419f1_QRegisterMetaType2(typeNameC)))
}

func QmlBridge_QmlRegisterType() int {
	return int(int32(C.QmlBridge3419f1_QmlBridge3419f1_QmlRegisterType()))
}

func (ptr *QmlBridge) QmlRegisterType() int {
	return int(int32(C.QmlBridge3419f1_QmlBridge3419f1_QmlRegisterType()))
}

func QmlBridge_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlBridge3419f1_QmlBridge3419f1_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlBridge3419f1_QmlBridge3419f1_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func QmlBridge_QmlRegisterUncreatableType(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	var messageC *C.char
	if message != "" {
		messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
	}
	return int(int32(C.QmlBridge3419f1_QmlBridge3419f1_QmlRegisterUncreatableType(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})))
}

func (ptr *QmlBridge) QmlRegisterUncreatableType(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	var messageC *C.char
	if message != "" {
		messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
	}
	return int(int32(C.QmlBridge3419f1_QmlBridge3419f1_QmlRegisterUncreatableType(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})))
}

func (ptr *QmlBridge) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge3419f1___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge3419f1___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __children_newList() unsafe.Pointer {
	return C.QmlBridge3419f1___children_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.QmlBridge3419f1___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge3419f1___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QmlBridge) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QmlBridge3419f1___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge3419f1___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge3419f1___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList() unsafe.Pointer {
	return C.QmlBridge3419f1___findChildren_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridge3419f1___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge3419f1___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList3() unsafe.Pointer {
	return C.QmlBridge3419f1___findChildren_newList3(ptr.Pointer())
}

func NewQmlBridge(parent std_core.QObject_ITF) *QmlBridge {
	QmlBridge_QRegisterMetaType()
	tmpValue := NewQmlBridgeFromPointer(C.QmlBridge3419f1_NewQmlBridge(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQmlBridge3419f1_DestroyQmlBridge
func callbackQmlBridge3419f1_DestroyQmlBridge(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QmlBridge"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQmlBridgeFromPointer(ptr).DestroyQmlBridgeDefault()
	}
}

func (ptr *QmlBridge) ConnectDestroyQmlBridge(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QmlBridge"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDestroyQmlBridge() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QmlBridge")
	}
}

func (ptr *QmlBridge) DestroyQmlBridge() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QmlBridge3419f1_DestroyQmlBridge(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QmlBridge) DestroyQmlBridgeDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QmlBridge3419f1_DestroyQmlBridgeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQmlBridge3419f1_ChildEvent
func callbackQmlBridge3419f1_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlBridge) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge3419f1_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQmlBridge3419f1_ConnectNotify
func callbackQmlBridge3419f1_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge3419f1_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge3419f1_CustomEvent
func callbackQmlBridge3419f1_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlBridge) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge3419f1_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQmlBridge3419f1_DeleteLater
func callbackQmlBridge3419f1_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQmlBridgeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QmlBridge) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QmlBridge3419f1_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQmlBridge3419f1_Destroyed
func callbackQmlBridge3419f1_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackQmlBridge3419f1_DisconnectNotify
func callbackQmlBridge3419f1_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge3419f1_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge3419f1_Event
func callbackQmlBridge3419f1_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QmlBridge) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QmlBridge3419f1_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQmlBridge3419f1_EventFilter
func callbackQmlBridge3419f1_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QmlBridge) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QmlBridge3419f1_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQmlBridge3419f1_ObjectNameChanged
func callbackQmlBridge3419f1_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQmlBridge3419f1_TimerEvent
func callbackQmlBridge3419f1_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlBridge) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge3419f1_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

func init() {
	qt.ItfMap["main.Person_ITF"] = Person{}
	qt.FuncMap["main.NewPerson"] = NewPerson
	qt.ItfMap["main.PersonModel_ITF"] = PersonModel{}
	qt.FuncMap["main.NewPersonModel"] = NewPersonModel
	qt.ItfMap["main.QmlBridge_ITF"] = QmlBridge{}
	qt.FuncMap["main.NewQmlBridge"] = NewQmlBridge
}
