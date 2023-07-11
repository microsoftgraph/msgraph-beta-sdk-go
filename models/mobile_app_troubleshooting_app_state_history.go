package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppTroubleshootingAppStateHistory history Item contained in the Mobile App Troubleshooting Event.
type MobileAppTroubleshootingAppStateHistory struct {
    MobileAppTroubleshootingHistoryItem
    // The OdataType property
    OdataType *string
}
// NewMobileAppTroubleshootingAppStateHistory instantiates a new mobileAppTroubleshootingAppStateHistory and sets the default values.
func NewMobileAppTroubleshootingAppStateHistory()(*MobileAppTroubleshootingAppStateHistory) {
    m := &MobileAppTroubleshootingAppStateHistory{
        MobileAppTroubleshootingHistoryItem: *NewMobileAppTroubleshootingHistoryItem(),
    }
    return m
}
// CreateMobileAppTroubleshootingAppStateHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppTroubleshootingAppStateHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppTroubleshootingAppStateHistory(), nil
}
// GetActionType gets the actionType property value. Defines the Action Types for an Intune Application.
func (m *MobileAppTroubleshootingAppStateHistory) GetActionType()(*MobileAppActionType) {
    val, err := m.GetBackingStore().Get("actionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MobileAppActionType)
    }
    return nil
}
// GetErrorCode gets the errorCode property value. Error code for the failure, empty if no failure.
func (m *MobileAppTroubleshootingAppStateHistory) GetErrorCode()(*string) {
    val, err := m.GetBackingStore().Get("errorCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppTroubleshootingAppStateHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppTroubleshootingHistoryItem.GetFieldDeserializers()
    res["actionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppActionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionType(val.(*MobileAppActionType))
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["runState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunState(val.(*RunState))
        }
        return nil
    }
    return res
}
// GetRunState gets the runState property value. Indicates the type of execution status of the device management script.
func (m *MobileAppTroubleshootingAppStateHistory) GetRunState()(*RunState) {
    val, err := m.GetBackingStore().Get("runState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RunState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobileAppTroubleshootingAppStateHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppTroubleshootingHistoryItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionType() != nil {
        cast := (*m.GetActionType()).String()
        err = writer.WriteStringValue("actionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    if m.GetRunState() != nil {
        cast := (*m.GetRunState()).String()
        err = writer.WriteStringValue("runState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionType sets the actionType property value. Defines the Action Types for an Intune Application.
func (m *MobileAppTroubleshootingAppStateHistory) SetActionType(value *MobileAppActionType)() {
    err := m.GetBackingStore().Set("actionType", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorCode sets the errorCode property value. Error code for the failure, empty if no failure.
func (m *MobileAppTroubleshootingAppStateHistory) SetErrorCode(value *string)() {
    err := m.GetBackingStore().Set("errorCode", value)
    if err != nil {
        panic(err)
    }
}
// SetRunState sets the runState property value. Indicates the type of execution status of the device management script.
func (m *MobileAppTroubleshootingAppStateHistory) SetRunState(value *RunState)() {
    err := m.GetBackingStore().Set("runState", value)
    if err != nil {
        panic(err)
    }
}
// MobileAppTroubleshootingAppStateHistoryable 
type MobileAppTroubleshootingAppStateHistoryable interface {
    MobileAppTroubleshootingHistoryItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionType()(*MobileAppActionType)
    GetErrorCode()(*string)
    GetRunState()(*RunState)
    SetActionType(value *MobileAppActionType)()
    SetErrorCode(value *string)()
    SetRunState(value *RunState)()
}
