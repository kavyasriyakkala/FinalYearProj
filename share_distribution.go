package main

// Message struct for sending shares between devices
type Message struct {
	SenderID string
	Data     struct {
		Share *Share
		HMAC  []byte
	}
}

// Distribute shares to peers (without using GroupKey during this phase)
func DistributeShares(device *Device, shares []*Share) {
	i := 0 // Use an integer index to access shares
	for _, peer := range device.Peers {
		// Generate HMAC for each share to ensure integrity (skip GroupKey use here)
		shareBytes := append(shares[i].X.Bytes(), shares[i].Y.Bytes()...)
		hmac := GenerateHMAC(shareBytes) // Modify to not rely on GroupKey during distribution
		peer.ReceiveShare(Message{
			SenderID: device.ID,
			Data: struct {
				Share *Share
				HMAC  []byte
			}{
				Share: shares[i],
				HMAC:  hmac,
			},
		})
		i++ // Increment index to move to the next share
	}
}

// Receive share and verify its integrity
func (device *Device) ReceiveShare(msg Message) {
	data := msg.Data

	shareBytes := append(data.Share.X.Bytes(), data.Share.Y.Bytes()...)
	if VerifyHMAC(shareBytes, data.HMAC) { // Skip GroupKey use here too
		println("Device", device.ID, "received valid share from", msg.SenderID)
		device.Share = data.Share
	} else {
		println("Device", device.ID, "received invalid share from", msg.SenderID)
	}
}
