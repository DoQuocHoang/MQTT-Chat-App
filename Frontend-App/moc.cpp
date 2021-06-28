

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractListModel>
#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QHash>
#include <QMap>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QSize>
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


typedef QMap<qint32, QByteArray> type378cdd;
class Person3419f1: public QObject
{
Q_OBJECT
Q_PROPERTY(QString id READ id WRITE setId NOTIFY idChanged)
Q_PROPERTY(QString firstName READ firstName WRITE setFirstName NOTIFY firstNameChanged)
Q_PROPERTY(QString lastName READ lastName WRITE setLastName NOTIFY lastNameChanged)
public:
	Person3419f1(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Person3419f1_Person3419f1_QRegisterMetaType();Person3419f1_Person3419f1_QRegisterMetaTypes();callbackPerson3419f1_Constructor(this);};
	QString id() { return ({ Moc_PackedString tempVal = callbackPerson3419f1_Id(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setId(QString id) { QByteArray* t87ea5d = new QByteArray(id.toUtf8()); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d->prepend("WHITESPACE").constData()+10), t87ea5d->size()-10, t87ea5d };callbackPerson3419f1_SetId(this, idPacked); };
	void Signal_IdChanged(QString id) { QByteArray* t87ea5d = new QByteArray(id.toUtf8()); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d->prepend("WHITESPACE").constData()+10), t87ea5d->size()-10, t87ea5d };callbackPerson3419f1_IdChanged(this, idPacked); };
	QString firstName() { return ({ Moc_PackedString tempVal = callbackPerson3419f1_FirstName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFirstName(QString firstName) { QByteArray* te57915 = new QByteArray(firstName.toUtf8()); Moc_PackedString firstNamePacked = { const_cast<char*>(te57915->prepend("WHITESPACE").constData()+10), te57915->size()-10, te57915 };callbackPerson3419f1_SetFirstName(this, firstNamePacked); };
	void Signal_FirstNameChanged(QString firstName) { QByteArray* te57915 = new QByteArray(firstName.toUtf8()); Moc_PackedString firstNamePacked = { const_cast<char*>(te57915->prepend("WHITESPACE").constData()+10), te57915->size()-10, te57915 };callbackPerson3419f1_FirstNameChanged(this, firstNamePacked); };
	QString lastName() { return ({ Moc_PackedString tempVal = callbackPerson3419f1_LastName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setLastName(QString lastName) { QByteArray* t9b831e = new QByteArray(lastName.toUtf8()); Moc_PackedString lastNamePacked = { const_cast<char*>(t9b831e->prepend("WHITESPACE").constData()+10), t9b831e->size()-10, t9b831e };callbackPerson3419f1_SetLastName(this, lastNamePacked); };
	void Signal_LastNameChanged(QString lastName) { QByteArray* t9b831e = new QByteArray(lastName.toUtf8()); Moc_PackedString lastNamePacked = { const_cast<char*>(t9b831e->prepend("WHITESPACE").constData()+10), t9b831e->size()-10, t9b831e };callbackPerson3419f1_LastNameChanged(this, lastNamePacked); };
	 ~Person3419f1() { callbackPerson3419f1_DestroyPerson(this); };
	void childEvent(QChildEvent * event) { callbackPerson3419f1_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackPerson3419f1_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackPerson3419f1_CustomEvent(this, event); };
	void deleteLater() { callbackPerson3419f1_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackPerson3419f1_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackPerson3419f1_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackPerson3419f1_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackPerson3419f1_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackPerson3419f1_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackPerson3419f1_TimerEvent(this, event); };
	QString idDefault() { return _id; };
	void setIdDefault(QString p) { if (p != _id) { _id = p; idChanged(_id); } };
	QString firstNameDefault() { return _firstName; };
	void setFirstNameDefault(QString p) { if (p != _firstName) { _firstName = p; firstNameChanged(_firstName); } };
	QString lastNameDefault() { return _lastName; };
	void setLastNameDefault(QString p) { if (p != _lastName) { _lastName = p; lastNameChanged(_lastName); } };
signals:
	void idChanged(QString id);
	void firstNameChanged(QString firstName);
	void lastNameChanged(QString lastName);
public slots:
private:
	QString _id;
	QString _firstName;
	QString _lastName;
};

Q_DECLARE_METATYPE(Person3419f1*)


void Person3419f1_Person3419f1_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class PersonModel3419f1: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<Person3419f1*> people READ people WRITE setPeople NOTIFY peopleChanged)
public:
	PersonModel3419f1(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");PersonModel3419f1_PersonModel3419f1_QRegisterMetaType();PersonModel3419f1_PersonModel3419f1_QRegisterMetaTypes();callbackPersonModel3419f1_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackPersonModel3419f1_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackPersonModel3419f1_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue037c88 = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue037c88, tmpValue037c88->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackPersonModel3419f1_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue037c88 = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue037c88, tmpValue037c88->size() }; })); };
	QList<Person3419f1*> people() { return ({ QList<Person3419f1*>* tmpP = static_cast<QList<Person3419f1*>*>(callbackPersonModel3419f1_People(this)); QList<Person3419f1*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setPeople(QList<Person3419f1*> people) { callbackPersonModel3419f1_SetPeople(this, ({ QList<Person3419f1*>* tmpValueab3ccc = new QList<Person3419f1*>(people); Moc_PackedList { tmpValueab3ccc, tmpValueab3ccc->size() }; })); };
	void Signal_PeopleChanged(QList<Person3419f1*> people) { callbackPersonModel3419f1_PeopleChanged(this, ({ QList<Person3419f1*>* tmpValueab3ccc = new QList<Person3419f1*>(people); Moc_PackedList { tmpValueab3ccc, tmpValueab3ccc->size() }; })); };
	 ~PersonModel3419f1() { callbackPersonModel3419f1_DestroyPersonModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackPersonModel3419f1_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackPersonModel3419f1_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackPersonModel3419f1_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackPersonModel3419f1_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackPersonModel3419f1_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackPersonModel3419f1_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackPersonModel3419f1_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackPersonModel3419f1_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackPersonModel3419f1_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackPersonModel3419f1_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackPersonModel3419f1_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackPersonModel3419f1_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackPersonModel3419f1_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackPersonModel3419f1_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackPersonModel3419f1_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackPersonModel3419f1_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue037c88 = new QVector<int>(roles); Moc_PackedList { tmpValue037c88, tmpValue037c88->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackPersonModel3419f1_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackPersonModel3419f1_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackPersonModel3419f1_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackPersonModel3419f1_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackPersonModel3419f1_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackPersonModel3419f1_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackPersonModel3419f1_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackPersonModel3419f1_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackPersonModel3419f1_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackPersonModel3419f1_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackPersonModel3419f1_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValuee0adf2 = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValuee0adf2, tmpValuee0adf2->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackPersonModel3419f1_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackPersonModel3419f1_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackPersonModel3419f1_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackPersonModel3419f1_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackPersonModel3419f1_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackPersonModel3419f1_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackPersonModel3419f1_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackPersonModel3419f1_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackPersonModel3419f1_ResetInternalData(this); };
	void revert() { callbackPersonModel3419f1_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackPersonModel3419f1_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackPersonModel3419f1_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackPersonModel3419f1_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackPersonModel3419f1_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackPersonModel3419f1_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackPersonModel3419f1_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackPersonModel3419f1_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackPersonModel3419f1_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackPersonModel3419f1_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackPersonModel3419f1_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackPersonModel3419f1_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue037c88 = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue037c88, tmpValue037c88->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackPersonModel3419f1_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackPersonModel3419f1_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackPersonModel3419f1_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackPersonModel3419f1_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackPersonModel3419f1_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackPersonModel3419f1_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackPersonModel3419f1_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackPersonModel3419f1_CustomEvent(this, event); };
	void deleteLater() { callbackPersonModel3419f1_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackPersonModel3419f1_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackPersonModel3419f1_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackPersonModel3419f1_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackPersonModel3419f1_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackPersonModel3419f1_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackPersonModel3419f1_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<Person3419f1*> peopleDefault() { return _people; };
	void setPeopleDefault(QList<Person3419f1*> p) { if (p != _people) { _people = p; peopleChanged(_people); } };
signals:
	void rolesChanged(type378cdd roles);
	void peopleChanged(QList<Person3419f1*> people);
public slots:
	void addPerson(Person3419f1* v0) { callbackPersonModel3419f1_AddPerson(this, v0); };
	void editPerson(qint32 row, QString firstName, QString lastName) { QByteArray* te57915 = new QByteArray(firstName.toUtf8()); Moc_PackedString firstNamePacked = { const_cast<char*>(te57915->prepend("WHITESPACE").constData()+10), te57915->size()-10, te57915 };QByteArray* t9b831e = new QByteArray(lastName.toUtf8()); Moc_PackedString lastNamePacked = { const_cast<char*>(t9b831e->prepend("WHITESPACE").constData()+10), t9b831e->size()-10, t9b831e };callbackPersonModel3419f1_EditPerson(this, row, firstNamePacked, lastNamePacked); };
	void removePerson(qint32 row) { callbackPersonModel3419f1_RemovePerson(this, row); };
private:
	type378cdd _roles;
	QList<Person3419f1*> _people;
};

Q_DECLARE_METATYPE(PersonModel3419f1*)


void PersonModel3419f1_PersonModel3419f1_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<Person3419f1*>");
}

class QmlBridge3419f1: public QObject
{
Q_OBJECT
public:
	QmlBridge3419f1(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QmlBridge3419f1_QmlBridge3419f1_QRegisterMetaType();QmlBridge3419f1_QmlBridge3419f1_QRegisterMetaTypes();callbackQmlBridge3419f1_Constructor(this);};
	void Signal_SendToQml(QString source, QString action, QString data) { QByteArray* t828d33 = new QByteArray(source.toUtf8()); Moc_PackedString sourcePacked = { const_cast<char*>(t828d33->prepend("WHITESPACE").constData()+10), t828d33->size()-10, t828d33 };QByteArray* t34eb4c = new QByteArray(action.toUtf8()); Moc_PackedString actionPacked = { const_cast<char*>(t34eb4c->prepend("WHITESPACE").constData()+10), t34eb4c->size()-10, t34eb4c };QByteArray* ta17c9a = new QByteArray(data.toUtf8()); Moc_PackedString dataPacked = { const_cast<char*>(ta17c9a->prepend("WHITESPACE").constData()+10), ta17c9a->size()-10, ta17c9a };callbackQmlBridge3419f1_SendToQml(this, sourcePacked, actionPacked, dataPacked); };
	 ~QmlBridge3419f1() { callbackQmlBridge3419f1_DestroyQmlBridge(this); };
	void childEvent(QChildEvent * event) { callbackQmlBridge3419f1_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQmlBridge3419f1_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQmlBridge3419f1_CustomEvent(this, event); };
	void deleteLater() { callbackQmlBridge3419f1_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQmlBridge3419f1_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQmlBridge3419f1_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQmlBridge3419f1_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQmlBridge3419f1_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQmlBridge3419f1_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQmlBridge3419f1_TimerEvent(this, event); };
signals:
	void sendToQml(QString source, QString action, QString data);
public slots:
	void sendToGo(QString source, QString action, QString data) { QByteArray* t828d33 = new QByteArray(source.toUtf8()); Moc_PackedString sourcePacked = { const_cast<char*>(t828d33->prepend("WHITESPACE").constData()+10), t828d33->size()-10, t828d33 };QByteArray* t34eb4c = new QByteArray(action.toUtf8()); Moc_PackedString actionPacked = { const_cast<char*>(t34eb4c->prepend("WHITESPACE").constData()+10), t34eb4c->size()-10, t34eb4c };QByteArray* ta17c9a = new QByteArray(data.toUtf8()); Moc_PackedString dataPacked = { const_cast<char*>(ta17c9a->prepend("WHITESPACE").constData()+10), ta17c9a->size()-10, ta17c9a };callbackQmlBridge3419f1_SendToGo(this, sourcePacked, actionPacked, dataPacked); };
	void registerToGo(QObject* object) { callbackQmlBridge3419f1_RegisterToGo(this, object); };
	void deregisterToGo(QString objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQmlBridge3419f1_DeregisterToGo(this, objectNamePacked); };
private:
};

Q_DECLARE_METATYPE(QmlBridge3419f1*)


void QmlBridge3419f1_QmlBridge3419f1_QRegisterMetaTypes() {
}

void PersonModel3419f1_AddPerson(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<PersonModel3419f1*>(ptr), "addPerson", Q_ARG(Person3419f1*, static_cast<Person3419f1*>(v0)));
}

void PersonModel3419f1_EditPerson(void* ptr, int row, struct Moc_PackedString firstName, struct Moc_PackedString lastName)
{
	QMetaObject::invokeMethod(static_cast<PersonModel3419f1*>(ptr), "editPerson", Q_ARG(qint32, row), Q_ARG(QString, QString::fromUtf8(firstName.data, firstName.len)), Q_ARG(QString, QString::fromUtf8(lastName.data, lastName.len)));
}

void PersonModel3419f1_RemovePerson(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<PersonModel3419f1*>(ptr), "removePerson", Q_ARG(qint32, row));
}

struct Moc_PackedList PersonModel3419f1_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue744ee8 = new QMap<qint32, QByteArray>(static_cast<PersonModel3419f1*>(ptr)->roles()); Moc_PackedList { tmpValue744ee8, tmpValue744ee8->size() }; });
}

struct Moc_PackedList PersonModel3419f1_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue518509 = new QMap<qint32, QByteArray>(static_cast<PersonModel3419f1*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue518509, tmpValue518509->size() }; });
}

void PersonModel3419f1_SetRoles(void* ptr, void* roles)
{
	static_cast<PersonModel3419f1*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void PersonModel3419f1_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<PersonModel3419f1*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void PersonModel3419f1_ConnectRolesChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<PersonModel3419f1*>(ptr), static_cast<void (PersonModel3419f1::*)(QMap<qint32, QByteArray>)>(&PersonModel3419f1::rolesChanged), static_cast<PersonModel3419f1*>(ptr), static_cast<void (PersonModel3419f1::*)(QMap<qint32, QByteArray>)>(&PersonModel3419f1::Signal_RolesChanged), static_cast<Qt::ConnectionType>(t));
}

void PersonModel3419f1_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<PersonModel3419f1*>(ptr), static_cast<void (PersonModel3419f1::*)(QMap<qint32, QByteArray>)>(&PersonModel3419f1::rolesChanged), static_cast<PersonModel3419f1*>(ptr), static_cast<void (PersonModel3419f1::*)(QMap<qint32, QByteArray>)>(&PersonModel3419f1::Signal_RolesChanged));
}

void PersonModel3419f1_RolesChanged(void* ptr, void* roles)
{
	static_cast<PersonModel3419f1*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList PersonModel3419f1_People(void* ptr)
{
	return ({ QList<Person3419f1*>* tmpValue9eb7a2 = new QList<Person3419f1*>(static_cast<PersonModel3419f1*>(ptr)->people()); Moc_PackedList { tmpValue9eb7a2, tmpValue9eb7a2->size() }; });
}

struct Moc_PackedList PersonModel3419f1_PeopleDefault(void* ptr)
{
	return ({ QList<Person3419f1*>* tmpValue75ebaa = new QList<Person3419f1*>(static_cast<PersonModel3419f1*>(ptr)->peopleDefault()); Moc_PackedList { tmpValue75ebaa, tmpValue75ebaa->size() }; });
}

void PersonModel3419f1_SetPeople(void* ptr, void* people)
{
	static_cast<PersonModel3419f1*>(ptr)->setPeople(({ QList<Person3419f1*>* tmpP = static_cast<QList<Person3419f1*>*>(people); QList<Person3419f1*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void PersonModel3419f1_SetPeopleDefault(void* ptr, void* people)
{
	static_cast<PersonModel3419f1*>(ptr)->setPeopleDefault(({ QList<Person3419f1*>* tmpP = static_cast<QList<Person3419f1*>*>(people); QList<Person3419f1*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void PersonModel3419f1_ConnectPeopleChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<PersonModel3419f1*>(ptr), static_cast<void (PersonModel3419f1::*)(QList<Person3419f1*>)>(&PersonModel3419f1::peopleChanged), static_cast<PersonModel3419f1*>(ptr), static_cast<void (PersonModel3419f1::*)(QList<Person3419f1*>)>(&PersonModel3419f1::Signal_PeopleChanged), static_cast<Qt::ConnectionType>(t));
}

void PersonModel3419f1_DisconnectPeopleChanged(void* ptr)
{
	QObject::disconnect(static_cast<PersonModel3419f1*>(ptr), static_cast<void (PersonModel3419f1::*)(QList<Person3419f1*>)>(&PersonModel3419f1::peopleChanged), static_cast<PersonModel3419f1*>(ptr), static_cast<void (PersonModel3419f1::*)(QList<Person3419f1*>)>(&PersonModel3419f1::Signal_PeopleChanged));
}

void PersonModel3419f1_PeopleChanged(void* ptr, void* people)
{
	static_cast<PersonModel3419f1*>(ptr)->peopleChanged(({ QList<Person3419f1*>* tmpP = static_cast<QList<Person3419f1*>*>(people); QList<Person3419f1*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int PersonModel3419f1_PersonModel3419f1_QRegisterMetaType()
{
	return qRegisterMetaType<PersonModel3419f1*>();
}

int PersonModel3419f1_PersonModel3419f1_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<PersonModel3419f1*>(const_cast<const char*>(typeName));
}

int PersonModel3419f1_PersonModel3419f1_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<PersonModel3419f1>();
#else
	return 0;
#endif
}

int PersonModel3419f1_PersonModel3419f1_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<PersonModel3419f1>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int PersonModel3419f1_PersonModel3419f1_QmlRegisterUncreatableType(char* uri, int versionMajor, int versionMinor, char* qmlName, struct Moc_PackedString message)
{
#ifdef QT_QML_LIB
	return qmlRegisterUncreatableType<PersonModel3419f1>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName), QString::fromUtf8(message.data, message.len));
#else
	return 0;
#endif
}

int PersonModel3419f1_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* PersonModel3419f1_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int PersonModel3419f1_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* PersonModel3419f1_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int PersonModel3419f1_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* PersonModel3419f1_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* PersonModel3419f1___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* PersonModel3419f1___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* PersonModel3419f1___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* PersonModel3419f1___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int PersonModel3419f1___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void PersonModel3419f1___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* PersonModel3419f1___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* PersonModel3419f1___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* PersonModel3419f1___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList PersonModel3419f1___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

void* PersonModel3419f1___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* PersonModel3419f1___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* PersonModel3419f1___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* PersonModel3419f1___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* PersonModel3419f1___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* PersonModel3419f1___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* PersonModel3419f1___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* PersonModel3419f1___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* PersonModel3419f1___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* PersonModel3419f1___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* PersonModel3419f1___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* PersonModel3419f1___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList PersonModel3419f1___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7fc3bb = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue7fc3bb, tmpValue7fc3bb->size() }; });
}

void* PersonModel3419f1___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* PersonModel3419f1___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList PersonModel3419f1___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

int PersonModel3419f1_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* PersonModel3419f1_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int PersonModel3419f1_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* PersonModel3419f1_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* PersonModel3419f1___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* PersonModel3419f1___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* PersonModel3419f1___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* PersonModel3419f1___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* PersonModel3419f1___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* PersonModel3419f1___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* PersonModel3419f1___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* PersonModel3419f1___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* PersonModel3419f1___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* PersonModel3419f1___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList PersonModel3419f1___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValuecf92c1 = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValuecf92c1, tmpValuecf92c1->size() }; });
}

void* PersonModel3419f1___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* PersonModel3419f1___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList PersonModel3419f1___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValuecf92c1 = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValuecf92c1, tmpValuecf92c1->size() }; });
}

void* PersonModel3419f1___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void PersonModel3419f1___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* PersonModel3419f1___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList PersonModel3419f1___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValuecf92c1 = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValuecf92c1, tmpValuecf92c1->size() }; });
}

void* PersonModel3419f1___people_atList(void* ptr, int i)
{
	return ({Person3419f1* tmp = static_cast<QList<Person3419f1*>*>(ptr)->at(i); if (i == static_cast<QList<Person3419f1*>*>(ptr)->size()-1) { static_cast<QList<Person3419f1*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1___people_setList(void* ptr, void* i)
{
	static_cast<QList<Person3419f1*>*>(ptr)->append(static_cast<Person3419f1*>(i));
}

void* PersonModel3419f1___people_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Person3419f1*>();
}

void* PersonModel3419f1___setPeople_people_atList(void* ptr, int i)
{
	return ({Person3419f1* tmp = static_cast<QList<Person3419f1*>*>(ptr)->at(i); if (i == static_cast<QList<Person3419f1*>*>(ptr)->size()-1) { static_cast<QList<Person3419f1*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1___setPeople_people_setList(void* ptr, void* i)
{
	static_cast<QList<Person3419f1*>*>(ptr)->append(static_cast<Person3419f1*>(i));
}

void* PersonModel3419f1___setPeople_people_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Person3419f1*>();
}

void* PersonModel3419f1___peopleChanged_people_atList(void* ptr, int i)
{
	return ({Person3419f1* tmp = static_cast<QList<Person3419f1*>*>(ptr)->at(i); if (i == static_cast<QList<Person3419f1*>*>(ptr)->size()-1) { static_cast<QList<Person3419f1*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1___peopleChanged_people_setList(void* ptr, void* i)
{
	static_cast<QList<Person3419f1*>*>(ptr)->append(static_cast<Person3419f1*>(i));
}

void* PersonModel3419f1___peopleChanged_people_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Person3419f1*>();
}

int PersonModel3419f1_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* PersonModel3419f1_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int PersonModel3419f1_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* PersonModel3419f1_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int PersonModel3419f1_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PersonModel3419f1_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* PersonModel3419f1_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* PersonModel3419f1_NewPersonModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new PersonModel3419f1(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new PersonModel3419f1(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new PersonModel3419f1(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new PersonModel3419f1(static_cast<QWindow*>(parent));
	} else {
		return new PersonModel3419f1(static_cast<QObject*>(parent));
	}
}

void PersonModel3419f1_DestroyPersonModel(void* ptr)
{
	static_cast<PersonModel3419f1*>(ptr)->~PersonModel3419f1();
}

void PersonModel3419f1_DestroyPersonModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char PersonModel3419f1_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long PersonModel3419f1_FlagsDefault(void* ptr, void* index)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* PersonModel3419f1_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* PersonModel3419f1_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* PersonModel3419f1_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char PersonModel3419f1_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char PersonModel3419f1_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int PersonModel3419f1_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* PersonModel3419f1_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void PersonModel3419f1_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char PersonModel3419f1_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* PersonModel3419f1_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char PersonModel3419f1_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char PersonModel3419f1_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList PersonModel3419f1_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue3d38f9 = new QMap<int, QVariant>(static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue3d38f9, tmpValue3d38f9->size() }; });
}

struct Moc_PackedList PersonModel3419f1_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue68cdc3 = new QList<QModelIndex>(static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue68cdc3, tmpValue68cdc3->size() }; });
}

void* PersonModel3419f1_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString PersonModel3419f1_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray* t3895e7 = new QByteArray(static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8()); Moc_PackedString { const_cast<char*>(t3895e7->prepend("WHITESPACE").constData()+10), t3895e7->size()-10, t3895e7 }; });
}

char PersonModel3419f1_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char PersonModel3419f1_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* PersonModel3419f1_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char PersonModel3419f1_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char PersonModel3419f1_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void PersonModel3419f1_ResetInternalDataDefault(void* ptr)
{
	static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::resetInternalData();
}

void PersonModel3419f1_RevertDefault(void* ptr)
{
	static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList PersonModel3419f1_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue613226 = new QHash<int, QByteArray>(static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue613226, tmpValue613226->size() }; });
}

int PersonModel3419f1_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char PersonModel3419f1_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char PersonModel3419f1_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char PersonModel3419f1_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void PersonModel3419f1_SortDefault(void* ptr, int column, long long order)
{
	static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* PersonModel3419f1_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char PersonModel3419f1_SubmitDefault(void* ptr)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::submit();
}

long long PersonModel3419f1_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long PersonModel3419f1_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::supportedDropActions();
}

void PersonModel3419f1_ChildEventDefault(void* ptr, void* event)
{
	static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void PersonModel3419f1_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void PersonModel3419f1_CustomEventDefault(void* ptr, void* event)
{
	static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void PersonModel3419f1_DeleteLaterDefault(void* ptr)
{
	static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::deleteLater();
}

void PersonModel3419f1_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char PersonModel3419f1_EventDefault(void* ptr, void* e)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char PersonModel3419f1_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void PersonModel3419f1_TimerEventDefault(void* ptr, void* event)
{
	static_cast<PersonModel3419f1*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QmlBridge3419f1_ConnectSendToQml(void* ptr, long long t)
{
	QObject::connect(static_cast<QmlBridge3419f1*>(ptr), static_cast<void (QmlBridge3419f1::*)(QString, QString, QString)>(&QmlBridge3419f1::sendToQml), static_cast<QmlBridge3419f1*>(ptr), static_cast<void (QmlBridge3419f1::*)(QString, QString, QString)>(&QmlBridge3419f1::Signal_SendToQml), static_cast<Qt::ConnectionType>(t));
}

void QmlBridge3419f1_DisconnectSendToQml(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge3419f1*>(ptr), static_cast<void (QmlBridge3419f1::*)(QString, QString, QString)>(&QmlBridge3419f1::sendToQml), static_cast<QmlBridge3419f1*>(ptr), static_cast<void (QmlBridge3419f1::*)(QString, QString, QString)>(&QmlBridge3419f1::Signal_SendToQml));
}

void QmlBridge3419f1_SendToQml(void* ptr, struct Moc_PackedString source, struct Moc_PackedString action, struct Moc_PackedString data)
{
	static_cast<QmlBridge3419f1*>(ptr)->sendToQml(QString::fromUtf8(source.data, source.len), QString::fromUtf8(action.data, action.len), QString::fromUtf8(data.data, data.len));
}

void QmlBridge3419f1_SendToGo(void* ptr, struct Moc_PackedString source, struct Moc_PackedString action, struct Moc_PackedString data)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge3419f1*>(ptr), "sendToGo", Q_ARG(QString, QString::fromUtf8(source.data, source.len)), Q_ARG(QString, QString::fromUtf8(action.data, action.len)), Q_ARG(QString, QString::fromUtf8(data.data, data.len)));
}

void QmlBridge3419f1_RegisterToGo(void* ptr, void* object)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge3419f1*>(ptr), "registerToGo", Q_ARG(QObject*, static_cast<QObject*>(object)));
}

void QmlBridge3419f1_DeregisterToGo(void* ptr, struct Moc_PackedString objectName)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge3419f1*>(ptr), "deregisterToGo", Q_ARG(QString, QString::fromUtf8(objectName.data, objectName.len)));
}

int QmlBridge3419f1_QmlBridge3419f1_QRegisterMetaType()
{
	return qRegisterMetaType<QmlBridge3419f1*>();
}

int QmlBridge3419f1_QmlBridge3419f1_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QmlBridge3419f1*>(const_cast<const char*>(typeName));
}

int QmlBridge3419f1_QmlBridge3419f1_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge3419f1>();
#else
	return 0;
#endif
}

int QmlBridge3419f1_QmlBridge3419f1_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge3419f1>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int QmlBridge3419f1_QmlBridge3419f1_QmlRegisterUncreatableType(char* uri, int versionMajor, int versionMinor, char* qmlName, struct Moc_PackedString message)
{
#ifdef QT_QML_LIB
	return qmlRegisterUncreatableType<QmlBridge3419f1>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName), QString::fromUtf8(message.data, message.len));
#else
	return 0;
#endif
}

void* QmlBridge3419f1___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge3419f1___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge3419f1___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QmlBridge3419f1___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QmlBridge3419f1___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QmlBridge3419f1___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QmlBridge3419f1___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge3419f1___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge3419f1___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridge3419f1___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge3419f1___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge3419f1___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridge3419f1_NewQmlBridge(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QmlBridge3419f1(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge3419f1(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QmlBridge3419f1(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge3419f1(static_cast<QWindow*>(parent));
	} else {
		return new QmlBridge3419f1(static_cast<QObject*>(parent));
	}
}

void QmlBridge3419f1_DestroyQmlBridge(void* ptr)
{
	static_cast<QmlBridge3419f1*>(ptr)->~QmlBridge3419f1();
}

void QmlBridge3419f1_DestroyQmlBridgeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QmlBridge3419f1_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge3419f1*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlBridge3419f1_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge3419f1*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge3419f1_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge3419f1*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlBridge3419f1_DeleteLaterDefault(void* ptr)
{
	static_cast<QmlBridge3419f1*>(ptr)->QObject::deleteLater();
}

void QmlBridge3419f1_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge3419f1*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QmlBridge3419f1_EventDefault(void* ptr, void* e)
{
	return static_cast<QmlBridge3419f1*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QmlBridge3419f1_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QmlBridge3419f1*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void QmlBridge3419f1_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge3419f1*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString Person3419f1_Id(void* ptr)
{
	return ({ QByteArray* t0ce0f5 = new QByteArray(static_cast<Person3419f1*>(ptr)->id().toUtf8()); Moc_PackedString { const_cast<char*>(t0ce0f5->prepend("WHITESPACE").constData()+10), t0ce0f5->size()-10, t0ce0f5 }; });
}

struct Moc_PackedString Person3419f1_IdDefault(void* ptr)
{
	return ({ QByteArray* t0bbe94 = new QByteArray(static_cast<Person3419f1*>(ptr)->idDefault().toUtf8()); Moc_PackedString { const_cast<char*>(t0bbe94->prepend("WHITESPACE").constData()+10), t0bbe94->size()-10, t0bbe94 }; });
}

void Person3419f1_SetId(void* ptr, struct Moc_PackedString id)
{
	static_cast<Person3419f1*>(ptr)->setId(QString::fromUtf8(id.data, id.len));
}

void Person3419f1_SetIdDefault(void* ptr, struct Moc_PackedString id)
{
	static_cast<Person3419f1*>(ptr)->setIdDefault(QString::fromUtf8(id.data, id.len));
}

void Person3419f1_ConnectIdChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<Person3419f1*>(ptr), static_cast<void (Person3419f1::*)(QString)>(&Person3419f1::idChanged), static_cast<Person3419f1*>(ptr), static_cast<void (Person3419f1::*)(QString)>(&Person3419f1::Signal_IdChanged), static_cast<Qt::ConnectionType>(t));
}

void Person3419f1_DisconnectIdChanged(void* ptr)
{
	QObject::disconnect(static_cast<Person3419f1*>(ptr), static_cast<void (Person3419f1::*)(QString)>(&Person3419f1::idChanged), static_cast<Person3419f1*>(ptr), static_cast<void (Person3419f1::*)(QString)>(&Person3419f1::Signal_IdChanged));
}

void Person3419f1_IdChanged(void* ptr, struct Moc_PackedString id)
{
	static_cast<Person3419f1*>(ptr)->idChanged(QString::fromUtf8(id.data, id.len));
}

struct Moc_PackedString Person3419f1_FirstName(void* ptr)
{
	return ({ QByteArray* td2d56b = new QByteArray(static_cast<Person3419f1*>(ptr)->firstName().toUtf8()); Moc_PackedString { const_cast<char*>(td2d56b->prepend("WHITESPACE").constData()+10), td2d56b->size()-10, td2d56b }; });
}

struct Moc_PackedString Person3419f1_FirstNameDefault(void* ptr)
{
	return ({ QByteArray* te9c1fc = new QByteArray(static_cast<Person3419f1*>(ptr)->firstNameDefault().toUtf8()); Moc_PackedString { const_cast<char*>(te9c1fc->prepend("WHITESPACE").constData()+10), te9c1fc->size()-10, te9c1fc }; });
}

void Person3419f1_SetFirstName(void* ptr, struct Moc_PackedString firstName)
{
	static_cast<Person3419f1*>(ptr)->setFirstName(QString::fromUtf8(firstName.data, firstName.len));
}

void Person3419f1_SetFirstNameDefault(void* ptr, struct Moc_PackedString firstName)
{
	static_cast<Person3419f1*>(ptr)->setFirstNameDefault(QString::fromUtf8(firstName.data, firstName.len));
}

void Person3419f1_ConnectFirstNameChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<Person3419f1*>(ptr), static_cast<void (Person3419f1::*)(QString)>(&Person3419f1::firstNameChanged), static_cast<Person3419f1*>(ptr), static_cast<void (Person3419f1::*)(QString)>(&Person3419f1::Signal_FirstNameChanged), static_cast<Qt::ConnectionType>(t));
}

void Person3419f1_DisconnectFirstNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<Person3419f1*>(ptr), static_cast<void (Person3419f1::*)(QString)>(&Person3419f1::firstNameChanged), static_cast<Person3419f1*>(ptr), static_cast<void (Person3419f1::*)(QString)>(&Person3419f1::Signal_FirstNameChanged));
}

void Person3419f1_FirstNameChanged(void* ptr, struct Moc_PackedString firstName)
{
	static_cast<Person3419f1*>(ptr)->firstNameChanged(QString::fromUtf8(firstName.data, firstName.len));
}

struct Moc_PackedString Person3419f1_LastName(void* ptr)
{
	return ({ QByteArray* t2d35c4 = new QByteArray(static_cast<Person3419f1*>(ptr)->lastName().toUtf8()); Moc_PackedString { const_cast<char*>(t2d35c4->prepend("WHITESPACE").constData()+10), t2d35c4->size()-10, t2d35c4 }; });
}

struct Moc_PackedString Person3419f1_LastNameDefault(void* ptr)
{
	return ({ QByteArray* tdd9162 = new QByteArray(static_cast<Person3419f1*>(ptr)->lastNameDefault().toUtf8()); Moc_PackedString { const_cast<char*>(tdd9162->prepend("WHITESPACE").constData()+10), tdd9162->size()-10, tdd9162 }; });
}

void Person3419f1_SetLastName(void* ptr, struct Moc_PackedString lastName)
{
	static_cast<Person3419f1*>(ptr)->setLastName(QString::fromUtf8(lastName.data, lastName.len));
}

void Person3419f1_SetLastNameDefault(void* ptr, struct Moc_PackedString lastName)
{
	static_cast<Person3419f1*>(ptr)->setLastNameDefault(QString::fromUtf8(lastName.data, lastName.len));
}

void Person3419f1_ConnectLastNameChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<Person3419f1*>(ptr), static_cast<void (Person3419f1::*)(QString)>(&Person3419f1::lastNameChanged), static_cast<Person3419f1*>(ptr), static_cast<void (Person3419f1::*)(QString)>(&Person3419f1::Signal_LastNameChanged), static_cast<Qt::ConnectionType>(t));
}

void Person3419f1_DisconnectLastNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<Person3419f1*>(ptr), static_cast<void (Person3419f1::*)(QString)>(&Person3419f1::lastNameChanged), static_cast<Person3419f1*>(ptr), static_cast<void (Person3419f1::*)(QString)>(&Person3419f1::Signal_LastNameChanged));
}

void Person3419f1_LastNameChanged(void* ptr, struct Moc_PackedString lastName)
{
	static_cast<Person3419f1*>(ptr)->lastNameChanged(QString::fromUtf8(lastName.data, lastName.len));
}

int Person3419f1_Person3419f1_QRegisterMetaType()
{
	return qRegisterMetaType<Person3419f1*>();
}

int Person3419f1_Person3419f1_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Person3419f1*>(const_cast<const char*>(typeName));
}

int Person3419f1_Person3419f1_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Person3419f1>();
#else
	return 0;
#endif
}

int Person3419f1_Person3419f1_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Person3419f1>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int Person3419f1_Person3419f1_QmlRegisterUncreatableType(char* uri, int versionMajor, int versionMinor, char* qmlName, struct Moc_PackedString message)
{
#ifdef QT_QML_LIB
	return qmlRegisterUncreatableType<Person3419f1>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName), QString::fromUtf8(message.data, message.len));
#else
	return 0;
#endif
}

void* Person3419f1___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Person3419f1___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Person3419f1___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* Person3419f1___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void Person3419f1___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Person3419f1___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* Person3419f1___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Person3419f1___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Person3419f1___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Person3419f1___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Person3419f1___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Person3419f1___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Person3419f1_NewPerson(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Person3419f1(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Person3419f1(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Person3419f1(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Person3419f1(static_cast<QWindow*>(parent));
	} else {
		return new Person3419f1(static_cast<QObject*>(parent));
	}
}

void Person3419f1_DestroyPerson(void* ptr)
{
	static_cast<Person3419f1*>(ptr)->~Person3419f1();
}

void Person3419f1_DestroyPersonDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void Person3419f1_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Person3419f1*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Person3419f1_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Person3419f1*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Person3419f1_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Person3419f1*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Person3419f1_DeleteLaterDefault(void* ptr)
{
	static_cast<Person3419f1*>(ptr)->QObject::deleteLater();
}

void Person3419f1_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Person3419f1*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char Person3419f1_EventDefault(void* ptr, void* e)
{
	return static_cast<Person3419f1*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Person3419f1_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Person3419f1*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void Person3419f1_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Person3419f1*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
