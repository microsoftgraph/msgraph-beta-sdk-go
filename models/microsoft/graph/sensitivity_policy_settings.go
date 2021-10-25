package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SensitivityPolicySettings struct {
    Entity
    applicableTo *SensitivityLabelTarget;
    downgradeSensitivityRequiresJustification *bool;
    helpWebUrl *string;
    isMandatory *bool;
}
func NewSensitivityPolicySettings()(*SensitivityPolicySettings) {
    m := &SensitivityPolicySettings{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SensitivityPolicySettings) GetApplicableTo()(*SensitivityLabelTarget) {
    if m == nil {
        return nil
    } else {
        return m.applicableTo
    }
}
func (m *SensitivityPolicySettings) GetDowngradeSensitivityRequiresJustification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.downgradeSensitivityRequiresJustification
    }
}
func (m *SensitivityPolicySettings) GetHelpWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.helpWebUrl
    }
}
func (m *SensitivityPolicySettings) GetIsMandatory()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMandatory
    }
}
func (m *SensitivityPolicySettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicableTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivityLabelTarget)
        if err != nil {
            return err
        }
        cast := val.(SensitivityLabelTarget)
        m.SetApplicableTo(&cast)
        return nil
    }
    res["downgradeSensitivityRequiresJustification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDowngradeSensitivityRequiresJustification(val)
        return nil
    }
    res["helpWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHelpWebUrl(val)
        return nil
    }
    res["isMandatory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMandatory(val)
        return nil
    }
    return res
}
func (m *SensitivityPolicySettings) IsNil()(bool) {
    return m == nil
}
func (m *SensitivityPolicySettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicableTo() != nil {
        cast := m.GetApplicableTo().String()
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
func (m *SensitivityPolicySettings) SetApplicableTo(value *SensitivityLabelTarget)() {
    m.applicableTo = value
}
func (m *SensitivityPolicySettings) SetDowngradeSensitivityRequiresJustification(value *bool)() {
    m.downgradeSensitivityRequiresJustification = value
}
func (m *SensitivityPolicySettings) SetHelpWebUrl(value *string)() {
    m.helpWebUrl = value
}
func (m *SensitivityPolicySettings) SetIsMandatory(value *bool)() {
    m.isMandatory = value
}
