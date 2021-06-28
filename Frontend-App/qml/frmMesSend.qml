import QtQuick 2.13
import QtQuick.Controls 2.13

Column {
                        id: myMesDelegate
                        //anchors.right: parent.right

                        //sentByMe ? parent.right : undefined
                        spacing: 6



                        Row {
                            id: messageRow
                            spacing: 6
                            anchors.right: parent.right



                            Rectangle {
                                id: rectangle
                                width: messageText.implicitWidth + 24
//                                    Math.min(messageText.implicitWidth + 24,
//                                    myMesDelegate.width)
                                height: messageText.implicitHeight + 24
                                color: "lightgrey"

                                Label {
                                    id: messageText
                                    text: mesData
                                    color: "black"
                                    anchors.fill: parent
                                    anchors.margins: 12
                                    wrapMode: Label.Wrap
                                }

                            }
                        }

                        Label {
                            id: timestampText
                            text: Qt.formatDateTime(new Date(), "d MMM hh:mm")
                            color: "lightgrey"
                            anchors.right: parent.right
                        }
                }
