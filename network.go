package main

type Message struct {
    SenderID string
    Data     interface{}
}

func SendShare(sender *Device, receiver *Device) {
    // Simulate network delay if necessary
    receiver.ReceiveShare(Message{
        SenderID: sender.ID,
        Data: struct {
            Share *Share
            HMAC  []byte
        }{
            Share: sender.Share,
            HMAC:  sender.HMAC,
        },
    })
}

func (device *Device) ReceiveShare(msg Message) {
    data := msg.Data.(struct {
        Share *Share
        HMAC  []byte
    })
    // Verify HMAC using pre-shared key
    shareBytes := append(data.Share.X.Bytes(), data.Share.Y.Bytes()...)
    if VerifyHMAC(shareBytes, data.HMAC) {
        device.Peers[msg.SenderID] = &Device{
            ID:    msg.SenderID,
            Share: data.Share,
        }
        println("Device", device.ID, "received valid share from", msg.SenderID)
    } else {
        println("Device", device.ID, "received invalid share from", msg.SenderID)
    }
}
