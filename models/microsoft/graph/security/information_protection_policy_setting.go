package security

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// InformationProtectionPolicySetting provides operations to manage the compliance singleton.
type InformationProtectionPolicySetting struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    defaultLabelId *string;
    // 
    isDowngradeJustificationRequired *bool;
    // 
    isMandatory *bool;
    // 
    moreInfoUrl *string;
}
// NewInformationProtectionPolicySetting instantiates a new informationProtectionPolicySetting and sets the default values.
func NewInformationProtectionPolicySetting()(*InformationProtectionPolicySetting) {
    m := &InformationProtectionPolicySetting{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// CreateInformationProtectionPolicySettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionPolicySettingFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewInformationProtectionPolicySetting(), nil
}
// GetDefaultLabelId gets the defaultLabelId property value. 
func (m *InformationProtectionPolicySetting) GetDefaultLabelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLabelId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionPolicySetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultLabelId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLabelId(val)
        }
        return nil
    }
    res["isDowngradeJustificationRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDowngradeJustificationRequired(val)
        }
        return nil
    }
    res["isMandatory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMandatory(val)
        }
        return nil
    }
    res["moreInfoUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetIsDowngradeJustificationRequired gets the isDowngradeJustificationRequired property value. 
func (m *InformationProtectionPolicySetting) GetIsDowngradeJustificationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDowngradeJustificationRequired
    }
}
// GetIsMandatory gets the isMandatory property value. 
func (m *InformationProtectionPolicySetting) GetIsMandatory()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMandatory
    }
}
// GetMoreInfoUrl gets the moreInfoUrl property value. 
func (m *InformationProtectionPolicySetting) GetMoreInfoUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.moreInfoUrl
    }
}
func (m *InformationProtectionPolicySetting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InformationProtectionPolicySetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetDefaultLabelId sets the defaultLabelId property value. 
func (m *InformationProtectionPolicySetting) SetDefaultLabelId(value *string)() {
    if m != nil {
        m.defaultLabelId = value
    }
}
// SetIsDowngradeJustificationRequired sets the isDowngradeJustificationRequired property value. 
func (m *InformationProtectionPolicySetting) SetIsDowngradeJustificationRequired(value *bool)() {
    if m != nil {
        m.isDowngradeJustificationRequired = value
    }
}
// SetIsMandatory sets the isMandatory property value. 
func (m *InformationProtectionPolicySetting) SetIsMandatory(value *bool)() {
    if m != nil {
        m.isMandatory = value
    }
}
// SetMoreInfoUrl sets the moreInfoUrl property value. 
func (m *InformationProtectionPolicySetting) SetMoreInfoUrl(value *string)() {
    if m != nil {
        m.moreInfoUrl = value
    }
}
