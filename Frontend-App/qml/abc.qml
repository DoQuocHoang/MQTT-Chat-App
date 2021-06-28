import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.3

ApplicationWindow {
    visible: true
    title: "Basic layouts"
    property int margin: 11
    minimumWidth: 600
    minimumHeight: 450

    ColumnLayout {
        id: mainLayout
        anchors.fill: parent
        anchors.margins: margin
        // ScrollView {

        //         clip: true
        //         ScrollBar.vertical.policy: ScrollBar.AlwaysOn
        //         Layout.fillHeight: true
        //     Layout.fillWidth: true
        //     anchors.top: parent.top
        //     anchors.bottom: typingBox.top
        //     anchors.bottomMargin: margin
        GroupBox {
            id: messagesBox
            Layout.fillHeight: true
            Layout.fillWidth: true
            anchors.top: parent.top
            anchors.bottom: typingBox.top
            anchors.bottomMargin: margin

            ScrollView {
                width: parent.width+12
                height: parent.height 
                // anchors.bottom: typingBox.top
                //anchors.margins: margin
                clip: true
                // ScrollBar.active
                ScrollBar.vertical.policy: ScrollBar.AlwaysOn
                //ScrollBar.verticalLayoutDirection:
                y: parent.height
                transform: Scale {yScale: -1}
                

                ListView {
                    id: mesListView
                    Layout.fillWidth: true
                    Layout.fillHeight: true
                    anchors.left: parent.left
                    anchors.right: parent.right
                    anchors.rightMargin: 17
                    // displayMarginBeginning: 40
                    // displayMarginEnd: 40

                    rotation: 180
                    // ScrollBar.vertical: ScrollBar {
                    // //ScrollBar.vertical.policy: ScrollBar.AlwaysOn
                    // height: parent.height
                    // y: parent.height
                    // x: parent.width
                    // transform: Scale {xScale: -1; yScale: -1}
                    // }
                    
                    //verticalLayoutDirection: ListView.TopToBottom
                    spacing: 12
                    model: ListModel{}
                    delegate: Row {
                        //id: mesIDCount
                        readonly property bool sentByMe: mesIDCount % 2 == 1

                        anchors.right: sentByMe ? parent.right : undefined
                        spacing: 6
                        rotation: 180
                        // Rectangle {
                        //     id: avatar
                        //     width: height
                        //     height: parent.height
                        //     color: "grey"
                        //     visible: !sentByMe
                        // }

                        TextField {
                            width: 80
                            height: 40
                            //color: sentByMe ? "lightgrey" : "steelblue"
                            Layout.fillWidth: true


                            text: mesData


                            readOnly: true
                            selectByMouse: true
                            
                            background: Rectangle {
                                
                                border.color: "lightgray"
                                radius: 10
                                color: "lightgray"
                            }
                        }
                    }

                    
                }

//                    delegate: TextField {
//                    id: mesField
//                    Layout.fillWidth: true
//                    text: ""
//                    readOnly: true
                    //anchors.rightMargin: 11

                    //Layout.alignment: Qt.AlignRight
                    //anchors.left: messagesBox.left
                    //width: messagesBox.width*0.7
                    //width: mesField.text.contentWidth

//                    property int count: 0

//                    background: Rectangle {
//                        implicitWidth: 200
//                        implicitHeight: 40
//                        border.color: "lightgray"
//                        radius: 10
//                        color: "yellow"
//                    }
//                    onTextChanged: {
//                        count++
//                        if(count % 2 == 0){
//                            anchors.left = undefined
//                            anchors.right = parent.right
//                            horizontalAlignment = Qt.AlignRight
//                        } else {
//                            anchors.right = undefined
//                            anchors.left = parent.left

//                            horizontalAlignment = Qt.AlignLeft
//                        }

//                        y += mesField.height

//                    }
                //} //textfield

                    
//                } // listview

                property int mesIDCount: 0

                Connections {
                    target: qmlBridge
                    onSendToQml:
                    {
                        if (source == "mesContent" && action == "txt")
                        {
                            //mesIDCount++
                            //mesField.text = data
                            //mesField.width = mesField.text.contentWidth
                            var row = {"mesData": data }
                            mesListView.model.insert(0, row)
                            // mesListView.model.append({"mesData": data})
                        }
                                
                    }
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
                }
                Button {
                    id: sendBtn
                    text: "Send"
                    onClicked: qmlBridge.sendToGo("sendBtn", "click", typingMes.text)
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
