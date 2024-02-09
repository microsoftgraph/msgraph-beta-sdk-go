package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppPolicySetItem a class containing the properties used for mobile app PolicySetItem.
type MobileAppPolicySetItem struct {
    PolicySetItem
}
// NewMobileAppPolicySetItem instantiates a new MobileAppPolicySetItem and sets the default values.
func NewMobileAppPolicySetItem()(*MobileAppPolicySetItem) {
    m := &MobileAppPolicySetItem{
        PolicySetItem: *NewPolicySetItem(),
    }
    odataTypeValue := "#microsoft.graph.mobileAppPolicySetItem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMobileAppPolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMobileAppPolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppPolicySetItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MobileAppPolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicySetItem.GetFieldDeserializers()
    res["intent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInstallIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntent(val.(*InstallIntent))
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppAssignmentSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(MobileAppAssignmentSettingsable))
        }
        return nil
    }
    return res
}
// GetIntent gets the intent property value. Possible values for the install intent chosen by the admin.
// returns a *InstallIntent when successful
func (m *MobileAppPolicySetItem) GetIntent()(*InstallIntent) {
    val, err := m.GetBackingStore().Get("intent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*InstallIntent)
    }
    return nil
}
// GetSettings gets the settings property value. Settings of the MobileAppPolicySetItem.
// returns a MobileAppAssignmentSettingsable when successful
func (m *MobileAppPolicySetItem) GetSettings()(MobileAppAssignmentSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MobileAppAssignmentSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobileAppPolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicySetItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIntent() != nil {
        cast := (*m.GetIntent()).String()
        err = writer.WriteStringValue("intent", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIntent sets the intent property value. Possible values for the install intent chosen by the admin.
func (m *MobileAppPolicySetItem) SetIntent(value *InstallIntent)() {
    err := m.GetBackingStore().Set("intent", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. Settings of the MobileAppPolicySetItem.
func (m *MobileAppPolicySetItem) SetSettings(value MobileAppAssignmentSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
type MobileAppPolicySetItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicySetItemable
    GetIntent()(*InstallIntent)
    GetSettings()(MobileAppAssignmentSettingsable)
    SetIntent(value *InstallIntent)()
    SetSettings(value MobileAppAssignmentSettingsable)()
}
