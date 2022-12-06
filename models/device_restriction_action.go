package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceRestrictionAction 
type DeviceRestrictionAction struct {
    DlpActionInfo
    // The message property
    message *string
    // The restrictionAction property
    restrictionAction *RestrictionAction
    // The triggers property
    triggers []RestrictionTrigger
}
// NewDeviceRestrictionAction instantiates a new DeviceRestrictionAction and sets the default values.
func NewDeviceRestrictionAction()(*DeviceRestrictionAction) {
    m := &DeviceRestrictionAction{
        DlpActionInfo: *NewDlpActionInfo(),
    }
    return m
}
// CreateDeviceRestrictionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceRestrictionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceRestrictionAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceRestrictionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DlpActionInfo.GetFieldDeserializers()
    res["message"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMessage)
    res["restrictionAction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRestrictionAction , m.SetRestrictionAction)
    res["triggers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseRestrictionTrigger , m.SetTriggers)
    return res
}
// GetMessage gets the message property value. The message property
func (m *DeviceRestrictionAction) GetMessage()(*string) {
    return m.message
}
// GetRestrictionAction gets the restrictionAction property value. The restrictionAction property
func (m *DeviceRestrictionAction) GetRestrictionAction()(*RestrictionAction) {
    return m.restrictionAction
}
// GetTriggers gets the triggers property value. The triggers property
func (m *DeviceRestrictionAction) GetTriggers()([]RestrictionTrigger) {
    return m.triggers
}
// Serialize serializes information the current object
func (m *DeviceRestrictionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DlpActionInfo.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    if m.GetRestrictionAction() != nil {
        cast := (*m.GetRestrictionAction()).String()
        err = writer.WriteStringValue("restrictionAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTriggers() != nil {
        err = writer.WriteCollectionOfStringValues("triggers", SerializeRestrictionTrigger(m.GetTriggers()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMessage sets the message property value. The message property
func (m *DeviceRestrictionAction) SetMessage(value *string)() {
    m.message = value
}
// SetRestrictionAction sets the restrictionAction property value. The restrictionAction property
func (m *DeviceRestrictionAction) SetRestrictionAction(value *RestrictionAction)() {
    m.restrictionAction = value
}
// SetTriggers sets the triggers property value. The triggers property
func (m *DeviceRestrictionAction) SetTriggers(value []RestrictionTrigger)() {
    m.triggers = value
}
