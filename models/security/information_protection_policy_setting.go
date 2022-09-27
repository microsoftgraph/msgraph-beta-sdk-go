package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// InformationProtectionPolicySetting 
type InformationProtectionPolicySetting struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The defaultLabelId property
    defaultLabelId *string
    // Exposes whether justification input is required on label downgrade.
    isDowngradeJustificationRequired *bool
    // Exposes whether mandatory labeling is enabled.
    isMandatory *bool
    // Exposes the more information URL that can be configured by the administrator.
    moreInfoUrl *string
}
// NewInformationProtectionPolicySetting instantiates a new informationProtectionPolicySetting and sets the default values.
func NewInformationProtectionPolicySetting()(*InformationProtectionPolicySetting) {
    m := &InformationProtectionPolicySetting{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.security.informationProtectionPolicySetting";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateInformationProtectionPolicySettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionPolicySettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionPolicySetting(), nil
}
// GetDefaultLabelId gets the defaultLabelId property value. The defaultLabelId property
func (m *InformationProtectionPolicySetting) GetDefaultLabelId()(*string) {
    return m.defaultLabelId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionPolicySetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultLabelId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultLabelId)
    res["isDowngradeJustificationRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsDowngradeJustificationRequired)
    res["isMandatory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsMandatory)
    res["moreInfoUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMoreInfoUrl)
    return res
}
// GetIsDowngradeJustificationRequired gets the isDowngradeJustificationRequired property value. Exposes whether justification input is required on label downgrade.
func (m *InformationProtectionPolicySetting) GetIsDowngradeJustificationRequired()(*bool) {
    return m.isDowngradeJustificationRequired
}
// GetIsMandatory gets the isMandatory property value. Exposes whether mandatory labeling is enabled.
func (m *InformationProtectionPolicySetting) GetIsMandatory()(*bool) {
    return m.isMandatory
}
// GetMoreInfoUrl gets the moreInfoUrl property value. Exposes the more information URL that can be configured by the administrator.
func (m *InformationProtectionPolicySetting) GetMoreInfoUrl()(*string) {
    return m.moreInfoUrl
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
    m.defaultLabelId = value
}
// SetIsDowngradeJustificationRequired sets the isDowngradeJustificationRequired property value. Exposes whether justification input is required on label downgrade.
func (m *InformationProtectionPolicySetting) SetIsDowngradeJustificationRequired(value *bool)() {
    m.isDowngradeJustificationRequired = value
}
// SetIsMandatory sets the isMandatory property value. Exposes whether mandatory labeling is enabled.
func (m *InformationProtectionPolicySetting) SetIsMandatory(value *bool)() {
    m.isMandatory = value
}
// SetMoreInfoUrl sets the moreInfoUrl property value. Exposes the more information URL that can be configured by the administrator.
func (m *InformationProtectionPolicySetting) SetMoreInfoUrl(value *string)() {
    m.moreInfoUrl = value
}
