import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.3
import QtQuick.Window 2.1


Page {
    id: root
    visible: true
    title: chatWith
    property int margin: 11
    // minimumWidth: 600
    // minimumHeight: 450

    property string chatWith
    property string whoSend
    property string myID
    property string friendID
    property string sentByMe
    

    Connections {
        target: qmlBridge
        onSendToQml:
        {
            if (source == "Sub")
            {
                // chatWith = data
                // myID = action
                // friendID = data
                listView.model.insert(0, {"mesData": data })
                //title = chatWith
            }
            if (source == "FriendID")
            {
                // chatWith = data
                // myID = action
                // friendID = data
                listView.model.insert(0, {"mesData": chatWith })
                listView.model.insert(0, {"mesData": myID })
                listView.model.insert(0, {"mesData": friendID })
                //title = chatWith
            }
            if (source == "connect" && data != "")
            {
                
                listView.model.insert(0, {"mesData": "Connected"})
            }
            if (source == "receiveMes")
            {
                whoSend = friendID
                sentByMe = false
//              var row = {"mesData": data }
                listView.model.insert(0, {"mesData": whoSend + ": " + data, })
            }
            if (source == "mesSent"  && data != "")
            {
                //whoSend = myID
                sentByMe = true
                //listView.model.insert(0)
                listView.model.insert(0, {"mesData": data})
                //listView.model.append({"mesData": data, sentByMe: true}})


            }

        }
    }

    

    ColumnLayout {
        id: mainLayout
        anchors.fill: parent
        anchors.margins: margin
        GroupBox {
            id: messagesBox
            Layout.fillHeight: true
            Layout.fillWidth: true
            anchors.top: parent.top
            anchors.bottom: typingBox.top
            anchors.bottomMargin: margin

            ScrollView {
                id: scrollView
                anchors.fill: parent
                anchors.rightMargin: -12
                clip: true


//                Column {
//                        id: frMesDelegate
                        
//                        anchors.right: undefined
//                        //sentByMe ? parent.right : undefined
//                        spacing: 6

                        

//                        Row {
//                            id: messageRow2
//                            spacing: 6
//                            anchors.right: undefined

//                            Image {
//                                id: avatar
//                                source: "qrc:/" + model.author.replace(" ", "_") + ".png"
//                            }

//                            Rectangle {
//                                width: Math.min(messageText2.implicitWidth + 24,
//                                    listView.width - (avatar.width + messageRow2.spacing ))
//                                height: messageText2.implicitHeight + 24
//                                color: "steelblue"

//                                Label {
//                                    id: messageText2
//                                    text: mesData
//                                    color: "white"
//                                    anchors.fill: parent
//                                    anchors.margins: 12
//                                    wrapMode: Label.Wrap
//                                }

//                            }
//                        }

//                        Label {
//                            id: timestampText2
//                            text: Qt.formatDateTime(new Date(), "d MMM hh:mm")
//                            color: "lightgrey"
//                            anchors.right: anRight
//                        }
//                }

                ListView {
                    id: listView
                    Layout.fillWidth: true
                    Layout.fillHeight: true
                    Layout.margins: typingBox.leftPadding + typingMes.leftPadding
                    anchors.right: parent.right
                    anchors.left: parent.left
                    anchors.rightMargin: 17
                    verticalLayoutDirection: ListView.BottomToTop
                    spacing: 12
                    model: ListModel {
                        ListElement {
                            mesData: "Grey"
                        }
                        ListElement {
                            mesData: "Grey"
                        }
                        ListElement {
                            mesData: "Grey"
                        }

                    }
                    delegate: Column {
                        id: myMesDelegate
                        anchors.right: parent.right
                        //sentByMe ? parent.right : undefined
                        spacing: 6



                        Row {
                            id: messageRow
                            spacing: 6
                            anchors.right: parent.right



                            Rectangle {
                                id: rectangle
                                width: Math.min(messageText.implicitWidth + 24,
                                    listView.width)
                                height: messageText.implicitHeight + 24
                                color: "lightgrey"

                                TextEdit {
                                    id: messageText
                                    text: mesData
                                    color: "black"
                                    anchors.fill: parent
                                    anchors.margins: 12
                                    wrapMode: Label.Wrap
                                    selectByMouse: true
                                }

                            }
                        }

                        TextEdit {
                            id: timestampText
                            text: Qt.formatDateTime(new Date(), "d MMM hh:mm")
                            color: "lightgrey"
                            anchors.right: parent.right
                            selectByMouse: true

                        }
                }


//                        Loader {
//                        property bool sentByMe
//                        sourceComponent : myMesDelegate
//                        //sourceComponent: sentByMe == true ? myMesDelegate : frMesDelegate


//                        }



                }

            }
        }
        GroupBox {
            id: typingBox
            Layout.fillWidth: true
            anchors.bottom: parent.bottom

            RowLayout {
                id: typingLayout
                anchors.fill: parent
                TextField {
                    id: typingMes
                    placeholderText: "Type something..."
                    Layout.fillWidth: true
                    selectByMouse: true
                    onEditingFinished: {
                        qmlBridge.sendToGo("sendBtn", friendID, typingMes.text)

                        typingMes.text = ""
                    }
//                    focus: true

//                    Keys.onPressed: {
//                        if (event.key == Qt.Key_Enter) {
//                            qmlBridge.sendToGo("sendBtn", friendID, typingMes.text)

//                            typingMes.text = ""

//                        }
//                    }
                }
                Button {
                    id: sendBtn
                    text: "Send"
                    enabled: typingMes.text != ""
                    onClicked: {
                        //listView.model.sendMessage(chatWith, typingMes.text);
                        //whoSend = myID
                        //sentByMe = false
                        //listView.model.insert(0, {"mesData": typingMes.text + " - " + sentByMe})
                        qmlBridge.sendToGo("sendBtn", friendID, typingMes.text)


                        typingMes.text = ""
                    }
                }
            }
        }

    }
}



/*##^##
Designer {
    D{i:0;autoSize:true;height:480;width:640}
}
##^##*/
