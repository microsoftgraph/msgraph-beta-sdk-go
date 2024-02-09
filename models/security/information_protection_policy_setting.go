package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type InformationProtectionPolicySetting struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewInformationProtectionPolicySetting instantiates a new InformationProtectionPolicySetting and sets the default values.
func NewInformationProtectionPolicySetting()(*InformationProtectionPolicySetting) {
    m := &InformationProtectionPolicySetting{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateInformationProtectionPolicySettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInformationProtectionPolicySettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionPolicySetting(), nil
}
// GetDefaultLabelId gets the defaultLabelId property value. The defaultLabelId property
// returns a *string when successful
func (m *InformationProtectionPolicySetting) GetDefaultLabelId()(*string) {
    val, err := m.GetBackingStore().Get("defaultLabelId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InformationProtectionPolicySetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultLabelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLabelId(val)
        }
        return nil
    }
    res["isDowngradeJustificationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDowngradeJustificationRequired(val)
        }
        return nil
    }
    res["isMandatory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMandatory(val)
        }
        return nil
    }
    res["moreInfoUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMoreInfoUrl(val)
        }
        return nil
    }
    return res
}
// GetIsDowngradeJustificationRequired gets the isDowngradeJustificationRequired property value. Exposes whether justification input is required on label downgrade.
// returns a *bool when successful
func (m *InformationProtectionPolicySetting) GetIsDowngradeJustificationRequired()(*bool) {
    val, err := m.GetBackingStore().Get("isDowngradeJustificationRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsMandatory gets the isMandatory property value. Exposes whether mandatory labeling is enabled.
// returns a *bool when successful
func (m *InformationProtectionPolicySetting) GetIsMandatory()(*bool) {
    val, err := m.GetBackingStore().Get("isMandatory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMoreInfoUrl gets the moreInfoUrl property value. Exposes the more information URL that can be configured by the administrator.
// returns a *string when successful
func (m *InformationProtectionPolicySetting) GetMoreInfoUrl()(*string) {
    val, err := m.GetBackingStore().Get("moreInfoUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InformationProtectionPolicySetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("defaultLabelId", m.GetDefaultLabelId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDowngradeJustificationRequired", m.GetIsDowngradeJustificationRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMandatory", m.GetIsMandatory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("moreInfoUrl", m.GetMoreInfoUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultLabelId sets the defaultLabelId property value. The defaultLabelId property
func (m *InformationProtectionPolicySetting) SetDefaultLabelId(value *string)() {
    err := m.GetBackingStore().Set("defaultLabelId", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDowngradeJustificationRequired sets the isDowngradeJustificationRequired property value. Exposes whether justification input is required on label downgrade.
func (m *InformationProtectionPolicySetting) SetIsDowngradeJustificationRequired(value *bool)() {
    err := m.GetBackingStore().Set("isDowngradeJustificationRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMandatory sets the isMandatory property value. Exposes whether mandatory labeling is enabled.
func (m *InformationProtectionPolicySetting) SetIsMandatory(value *bool)() {
    err := m.GetBackingStore().Set("isMandatory", value)
    if err != nil {
        panic(err)
    }
}
// SetMoreInfoUrl sets the moreInfoUrl property value. Exposes the more information URL that can be configured by the administrator.
func (m *InformationProtectionPolicySetting) SetMoreInfoUrl(value *string)() {
    err := m.GetBackingStore().Set("moreInfoUrl", value)
    if err != nil {
        panic(err)
    }
}
type InformationProtectionPolicySettingable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultLabelId()(*string)
    GetIsDowngradeJustificationRequired()(*bool)
    GetIsMandatory()(*bool)
    GetMoreInfoUrl()(*string)
    SetDefaultLabelId(value *string)()
    SetIsDowngradeJustificationRequired(value *bool)()
    SetIsMandatory(value *bool)()
    SetMoreInfoUrl(value *string)()
}
