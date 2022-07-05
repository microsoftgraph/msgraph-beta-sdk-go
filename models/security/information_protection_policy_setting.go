package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// InformationProtectionPolicySetting provides operations to manage the collection of accessReview entities.
type InformationProtectionPolicySetting struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The defaultLabelId property
    defaultLabelId *string
    // The isDowngradeJustificationRequired property
    isDowngradeJustificationRequired *bool
    // The isMandatory property
    isMandatory *bool
    // The moreInfoUrl property
    moreInfoUrl *string
}
// NewInformationProtectionPolicySetting instantiates a new informationProtectionPolicySetting and sets the default values.
func NewInformationProtectionPolicySetting()(*InformationProtectionPolicySetting) {
    m := &InformationProtectionPolicySetting{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateInformationProtectionPolicySettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionPolicySettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionPolicySetting(), nil
}
// GetDefaultLabelId gets the defaultLabelId property value. The defaultLabelId property
func (m *InformationProtectionPolicySetting) GetDefaultLabelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLabelId
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetIsDowngradeJustificationRequired gets the isDowngradeJustificationRequired property value. The isDowngradeJustificationRequired property
func (m *InformationProtectionPolicySetting) GetIsDowngradeJustificationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDowngradeJustificationRequired
    }
}
// GetIsMandatory gets the isMandatory property value. The isMandatory property
func (m *InformationProtectionPolicySetting) GetIsMandatory()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMandatory
    }
}
// GetMoreInfoUrl gets the moreInfoUrl property value. The moreInfoUrl property
func (m *InformationProtectionPolicySetting) GetMoreInfoUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.moreInfoUrl
    }
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
    if m != nil {
        m.defaultLabelId = value
    }
}
// SetIsDowngradeJustificationRequired sets the isDowngradeJustificationRequired property value. The isDowngradeJustificationRequired property
func (m *InformationProtectionPolicySetting) SetIsDowngradeJustificationRequired(value *bool)() {
    if m != nil {
        m.isDowngradeJustificationRequired = value
    }
}
// SetIsMandatory sets the isMandatory property value. The isMandatory property
func (m *InformationProtectionPolicySetting) SetIsMandatory(value *bool)() {
    if m != nil {
        m.isMandatory = value
    }
}
// SetMoreInfoUrl sets the moreInfoUrl property value. The moreInfoUrl property
func (m *InformationProtectionPolicySetting) SetMoreInfoUrl(value *string)() {
    if m != nil {
        m.moreInfoUrl = value
    }
}
