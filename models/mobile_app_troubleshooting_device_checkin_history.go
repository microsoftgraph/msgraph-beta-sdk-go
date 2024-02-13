package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppTroubleshootingDeviceCheckinHistory history Item contained in the Mobile App Troubleshooting Event.
type MobileAppTroubleshootingDeviceCheckinHistory struct {
    MobileAppTroubleshootingHistoryItem
}
// NewMobileAppTroubleshootingDeviceCheckinHistory instantiates a new MobileAppTroubleshootingDeviceCheckinHistory and sets the default values.
func NewMobileAppTroubleshootingDeviceCheckinHistory()(*MobileAppTroubleshootingDeviceCheckinHistory) {
    m := &MobileAppTroubleshootingDeviceCheckinHistory{
        MobileAppTroubleshootingHistoryItem: *NewMobileAppTroubleshootingHistoryItem(),
    }
    return m
}
// CreateMobileAppTroubleshootingDeviceCheckinHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMobileAppTroubleshootingDeviceCheckinHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppTroubleshootingDeviceCheckinHistory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MobileAppTroubleshootingDeviceCheckinHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppTroubleshootingHistoryItem.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MobileAppTroubleshootingDeviceCheckinHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppTroubleshootingHistoryItem.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type MobileAppTroubleshootingDeviceCheckinHistoryable interface {
    MobileAppTroubleshootingHistoryItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
