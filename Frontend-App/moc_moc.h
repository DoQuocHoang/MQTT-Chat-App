/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.13.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <memory>
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#include <QtCore/QList>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.13.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_Person3419f1_t {
    QByteArrayData data[8];
    char stringdata0[79];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_Person3419f1_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_Person3419f1_t qt_meta_stringdata_Person3419f1 = {
    {
QT_MOC_LITERAL(0, 0, 12), // "Person3419f1"
QT_MOC_LITERAL(1, 13, 9), // "idChanged"
QT_MOC_LITERAL(2, 23, 0), // ""
QT_MOC_LITERAL(3, 24, 2), // "id"
QT_MOC_LITERAL(4, 27, 16), // "firstNameChanged"
QT_MOC_LITERAL(5, 44, 9), // "firstName"
QT_MOC_LITERAL(6, 54, 15), // "lastNameChanged"
QT_MOC_LITERAL(7, 70, 8) // "lastName"

    },
    "Person3419f1\0idChanged\0\0id\0firstNameChanged\0"
    "firstName\0lastNameChanged\0lastName"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_Person3419f1[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       3,   14, // methods
       3,   38, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       3,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   29,    2, 0x06 /* Public */,
       4,    1,   32,    2, 0x06 /* Public */,
       6,    1,   35,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    5,
    QMetaType::Void, QMetaType::QString,    7,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,
       5, QMetaType::QString, 0x00495103,
       7, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,

       0        // eod
};

void Person3419f1::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<Person3419f1 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->idChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->firstNameChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->lastNameChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (Person3419f1::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Person3419f1::idChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (Person3419f1::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Person3419f1::firstNameChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (Person3419f1::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Person3419f1::lastNameChanged)) {
                *result = 2;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<Person3419f1 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->id(); break;
        case 1: *reinterpret_cast< QString*>(_v) = _t->firstName(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->lastName(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<Person3419f1 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setId(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setFirstName(*reinterpret_cast< QString*>(_v)); break;
        case 2: _t->setLastName(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject Person3419f1::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_Person3419f1.data,
    qt_meta_data_Person3419f1,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *Person3419f1::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *Person3419f1::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_Person3419f1.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int Person3419f1::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 3)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 3)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 3;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 3;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void Person3419f1::idChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void Person3419f1::firstNameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void Person3419f1::lastNameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
struct qt_meta_stringdata_PersonModel3419f1_t {
    QByteArrayData data[16];
    char stringdata0[165];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_PersonModel3419f1_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_PersonModel3419f1_t qt_meta_stringdata_PersonModel3419f1 = {
    {
QT_MOC_LITERAL(0, 0, 17), // "PersonModel3419f1"
QT_MOC_LITERAL(1, 18, 12), // "rolesChanged"
QT_MOC_LITERAL(2, 31, 0), // ""
QT_MOC_LITERAL(3, 32, 10), // "type378cdd"
QT_MOC_LITERAL(4, 43, 5), // "roles"
QT_MOC_LITERAL(5, 49, 13), // "peopleChanged"
QT_MOC_LITERAL(6, 63, 20), // "QList<Person3419f1*>"
QT_MOC_LITERAL(7, 84, 6), // "people"
QT_MOC_LITERAL(8, 91, 9), // "addPerson"
QT_MOC_LITERAL(9, 101, 13), // "Person3419f1*"
QT_MOC_LITERAL(10, 115, 2), // "v0"
QT_MOC_LITERAL(11, 118, 10), // "editPerson"
QT_MOC_LITERAL(12, 129, 3), // "row"
QT_MOC_LITERAL(13, 133, 9), // "firstName"
QT_MOC_LITERAL(14, 143, 8), // "lastName"
QT_MOC_LITERAL(15, 152, 12) // "removePerson"

    },
    "PersonModel3419f1\0rolesChanged\0\0"
    "type378cdd\0roles\0peopleChanged\0"
    "QList<Person3419f1*>\0people\0addPerson\0"
    "Person3419f1*\0v0\0editPerson\0row\0"
    "firstName\0lastName\0removePerson"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_PersonModel3419f1[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       5,   14, // methods
       2,   58, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       2,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   39,    2, 0x06 /* Public */,
       5,    1,   42,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       8,    1,   45,    2, 0x0a /* Public */,
      11,    3,   48,    2, 0x0a /* Public */,
      15,    1,   55,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    QMetaType::Void, 0x80000000 | 6,    7,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 9,   10,
    QMetaType::Void, QMetaType::Int, QMetaType::QString, QMetaType::QString,   12,   13,   14,
    QMetaType::Void, QMetaType::Int,   12,

 // properties: name, type, flags
       4, 0x80000000 | 3, 0x0049510b,
       7, 0x80000000 | 6, 0x0049510b,

 // properties: notify_signal_id
       0,
       1,

       0        // eod
};

void PersonModel3419f1::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<PersonModel3419f1 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 1: _t->peopleChanged((*reinterpret_cast< QList<Person3419f1*>(*)>(_a[1]))); break;
        case 2: _t->addPerson((*reinterpret_cast< Person3419f1*(*)>(_a[1]))); break;
        case 3: _t->editPerson((*reinterpret_cast< qint32(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3]))); break;
        case 4: _t->removePerson((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<Person3419f1*> >(); break;
            }
            break;
        case 2:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< Person3419f1* >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (PersonModel3419f1::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&PersonModel3419f1::rolesChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (PersonModel3419f1::*)(QList<Person3419f1*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&PersonModel3419f1::peopleChanged)) {
                *result = 1;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<Person3419f1*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<PersonModel3419f1 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<Person3419f1*>*>(_v) = _t->people(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<PersonModel3419f1 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setPeople(*reinterpret_cast< QList<Person3419f1*>*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject PersonModel3419f1::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_PersonModel3419f1.data,
    qt_meta_data_PersonModel3419f1,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *PersonModel3419f1::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *PersonModel3419f1::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_PersonModel3419f1.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int PersonModel3419f1::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractListModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 5)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 5;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 5)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 5;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 2;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void PersonModel3419f1::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void PersonModel3419f1::peopleChanged(QList<Person3419f1*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
struct qt_meta_stringdata_QmlBridge3419f1_t {
    QByteArrayData data[11];
    char stringdata0[101];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QmlBridge3419f1_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QmlBridge3419f1_t qt_meta_stringdata_QmlBridge3419f1 = {
    {
QT_MOC_LITERAL(0, 0, 15), // "QmlBridge3419f1"
QT_MOC_LITERAL(1, 16, 9), // "sendToQml"
QT_MOC_LITERAL(2, 26, 0), // ""
QT_MOC_LITERAL(3, 27, 6), // "source"
QT_MOC_LITERAL(4, 34, 6), // "action"
QT_MOC_LITERAL(5, 41, 4), // "data"
QT_MOC_LITERAL(6, 46, 8), // "sendToGo"
QT_MOC_LITERAL(7, 55, 12), // "registerToGo"
QT_MOC_LITERAL(8, 68, 6), // "object"
QT_MOC_LITERAL(9, 75, 14), // "deregisterToGo"
QT_MOC_LITERAL(10, 90, 10) // "objectName"

    },
    "QmlBridge3419f1\0sendToQml\0\0source\0"
    "action\0data\0sendToGo\0registerToGo\0"
    "object\0deregisterToGo\0objectName"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QmlBridge3419f1[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       4,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       1,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    3,   34,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       6,    3,   41,    2, 0x0a /* Public */,
       7,    1,   48,    2, 0x0a /* Public */,
       9,    1,   51,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString, QMetaType::QString, QMetaType::QString,    3,    4,    5,

 // slots: parameters
    QMetaType::Void, QMetaType::QString, QMetaType::QString, QMetaType::QString,    3,    4,    5,
    QMetaType::Void, QMetaType::QObjectStar,    8,
    QMetaType::Void, QMetaType::QString,   10,

       0        // eod
};

void QmlBridge3419f1::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<QmlBridge3419f1 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->sendToQml((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3]))); break;
        case 1: _t->sendToGo((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3]))); break;
        case 2: _t->registerToGo((*reinterpret_cast< QObject*(*)>(_a[1]))); break;
        case 3: _t->deregisterToGo((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (QmlBridge3419f1::*)(QString , QString , QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QmlBridge3419f1::sendToQml)) {
                *result = 0;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject QmlBridge3419f1::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_QmlBridge3419f1.data,
    qt_meta_data_QmlBridge3419f1,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *QmlBridge3419f1::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QmlBridge3419f1::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QmlBridge3419f1.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int QmlBridge3419f1::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 4)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 4)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 4;
    }
    return _id;
}

// SIGNAL 0
void QmlBridge3419f1::sendToQml(QString _t1, QString _t2, QString _t3)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))), const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t2))), const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t3))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
