package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppIntentAndState 
type MobileAppIntentAndState struct {
    Entity
    // Device identifier created or collected by Intune.
    managedDeviceIdentifier *string
    // The list of payload intents and states for the tenant.
    mobileAppList []MobileAppIntentAndStateDetailable
    // Identifier for the user that tried to enroll the device.
    userId *string
}
// NewMobileAppIntentAndState instantiates a new MobileAppIntentAndState and sets the default values.
func NewMobileAppIntentAndState()(*MobileAppIntentAndState) {
    m := &MobileAppIntentAndState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppIntentAndStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppIntentAndStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppIntentAndState(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppIntentAndState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["managedDeviceIdentifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagedDeviceIdentifier)
    res["mobileAppList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppIntentAndStateDetailFromDiscriminatorValue , m.SetMobileAppList)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    return res
}
// GetManagedDeviceIdentifier gets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppIntentAndState) GetManagedDeviceIdentifier()(*string) {
    return m.managedDeviceIdentifier
}
// GetMobileAppList gets the mobileAppList property value. The list of payload intents and states for the tenant.
func (m *MobileAppIntentAndState) GetMobileAppList()([]MobileAppIntentAndStateDetailable) {
    return m.mobileAppList
}
// GetUserId gets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppIntentAndState) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *MobileAppIntentAndState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("managedDeviceIdentifier", m.GetManagedDeviceIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppList() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMobileAppList())
        err = writer.WriteCollectionOfObjectValues("mobileAppList", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetManagedDeviceIdentifier sets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppIntentAndState) SetManagedDeviceIdentifier(value *string)() {
    m.managedDeviceIdentifier = value
}
// SetMobileAppList sets the mobileAppList property value. The list of payload intents and states for the tenant.
func (m *MobileAppIntentAndState) SetMobileAppList(value []MobileAppIntentAndStateDetailable)() {
    m.mobileAppList = value
}
// SetUserId sets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppIntentAndState) SetUserId(value *string)() {
    m.userId = value
}
