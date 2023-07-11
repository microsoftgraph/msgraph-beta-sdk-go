package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppTroubleshootingAppUpdateHistory history Item contained in the Mobile App Troubleshooting Event.
type MobileAppTroubleshootingAppUpdateHistory struct {
    MobileAppTroubleshootingHistoryItem
    // The OdataType property
    OdataType *string
}
// NewMobileAppTroubleshootingAppUpdateHistory instantiates a new mobileAppTroubleshootingAppUpdateHistory and sets the default values.
func NewMobileAppTroubleshootingAppUpdateHistory()(*MobileAppTroubleshootingAppUpdateHistory) {
    m := &MobileAppTroubleshootingAppUpdateHistory{
        MobileAppTroubleshootingHistoryItem: *NewMobileAppTroubleshootingHistoryItem(),
    }
    return m
}
// CreateMobileAppTroubleshootingAppUpdateHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppTroubleshootingAppUpdateHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppTroubleshootingAppUpdateHistory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppTroubleshootingAppUpdateHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppTroubleshootingHistoryItem.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MobileAppTroubleshootingAppUpdateHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppTroubleshootingHistoryItem.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// MobileAppTroubleshootingAppUpdateHistoryable 
type MobileAppTroubleshootingAppUpdateHistoryable interface {
    MobileAppTroubleshootingHistoryItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
