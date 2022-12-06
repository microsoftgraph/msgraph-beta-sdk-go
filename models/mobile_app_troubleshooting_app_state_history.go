package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppTroubleshootingAppStateHistory 
type MobileAppTroubleshootingAppStateHistory struct {
    MobileAppTroubleshootingHistoryItem
    // Defines the Action Types for an Intune Application.
    actionType *MobileAppActionType
    // Error code for the failure, empty if no failure.
    errorCode *string
    // Indicates the type of execution status of the device management script.
    runState *RunState
}
// NewMobileAppTroubleshootingAppStateHistory instantiates a new MobileAppTroubleshootingAppStateHistory and sets the default values.
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
    return m.actionType
}
// GetErrorCode gets the errorCode property value. Error code for the failure, empty if no failure.
func (m *MobileAppTroubleshootingAppStateHistory) GetErrorCode()(*string) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppTroubleshootingAppStateHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppTroubleshootingHistoryItem.GetFieldDeserializers()
    res["actionType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMobileAppActionType , m.SetActionType)
    res["errorCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetErrorCode)
    res["runState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRunState , m.SetRunState)
    return res
}
// GetRunState gets the runState property value. Indicates the type of execution status of the device management script.
func (m *MobileAppTroubleshootingAppStateHistory) GetRunState()(*RunState) {
    return m.runState
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
    m.actionType = value
}
// SetErrorCode sets the errorCode property value. Error code for the failure, empty if no failure.
func (m *MobileAppTroubleshootingAppStateHistory) SetErrorCode(value *string)() {
    m.errorCode = value
}
// SetRunState sets the runState property value. Indicates the type of execution status of the device management script.
func (m *MobileAppTroubleshootingAppStateHistory) SetRunState(value *RunState)() {
    m.runState = value
}
