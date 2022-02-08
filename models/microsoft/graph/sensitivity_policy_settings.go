package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SensitivityPolicySettings 
type SensitivityPolicySettings struct {
    Entity
    // 
    applicableTo *SensitivityLabelTarget;
    // 
    downgradeSensitivityRequiresJustification *bool;
    // 
    helpWebUrl *string;
    // 
    isMandatory *bool;
}
// NewSensitivityPolicySettings instantiates a new sensitivityPolicySettings and sets the default values.
func NewSensitivityPolicySettings()(*SensitivityPolicySettings) {
    m := &SensitivityPolicySettings{
        Entity: *NewEntity(),
    }
    return m
}
// GetApplicableTo gets the applicableTo property value. 
func (m *SensitivityPolicySettings) GetApplicableTo()(*SensitivityLabelTarget) {
    if m == nil {
        return nil
    } else {
        return m.applicableTo
    }
}
// GetDowngradeSensitivityRequiresJustification gets the downgradeSensitivityRequiresJustification property value. 
func (m *SensitivityPolicySettings) GetDowngradeSensitivityRequiresJustification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.downgradeSensitivityRequiresJustification
    }
}
// GetHelpWebUrl gets the helpWebUrl property value. 
func (m *SensitivityPolicySettings) GetHelpWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.helpWebUrl
    }
}
// GetIsMandatory gets the isMandatory property value. 
func (m *SensitivityPolicySettings) GetIsMandatory()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMandatory
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitivityPolicySettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicableTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivityLabelTarget)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableTo(val.(*SensitivityLabelTarget))
        }
        return nil
    }
    res["downgradeSensitivityRequiresJustification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDowngradeSensitivityRequiresJustification(val)
        }
        return nil
    }
    res["helpWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpWebUrl(val)
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
    return res
}
func (m *SensitivityPolicySettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SensitivityPolicySettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicableTo() != nil {
        cast := (*m.GetApplicableTo()).String()
        err = writer.WriteStringValue("applicableTo", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("downgradeSensitivityRequiresJustification", m.GetDowngradeSensitivityRequiresJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("helpWebUrl", m.GetHelpWebUrl())
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
    return nil
}
// SetApplicableTo sets the applicableTo property value. 
func (m *SensitivityPolicySettings) SetApplicableTo(value *SensitivityLabelTarget)() {
    if m != nil {
        m.applicableTo = value
    }
}
// SetDowngradeSensitivityRequiresJustification sets the downgradeSensitivityRequiresJustification property value. 
func (m *SensitivityPolicySettings) SetDowngradeSensitivityRequiresJustification(value *bool)() {
    if m != nil {
        m.downgradeSensitivityRequiresJustification = value
    }
}
// SetHelpWebUrl sets the helpWebUrl property value. 
func (m *SensitivityPolicySettings) SetHelpWebUrl(value *string)() {
    if m != nil {
        m.helpWebUrl = value
    }
}
// SetIsMandatory sets the isMandatory property value. 
func (m *SensitivityPolicySettings) SetIsMandatory(value *bool)() {
    if m != nil {
        m.isMandatory = value
    }
}
