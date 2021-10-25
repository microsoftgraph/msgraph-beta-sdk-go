package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DlpEvaluationInput struct {
    accessScope *AccessScope;
    additionalData map[string]interface{};
    currentLabel *CurrentLabel;
    discoveredSensitiveTypes []DiscoveredSensitiveType;
}
func NewDlpEvaluationInput()(*DlpEvaluationInput) {
    m := &DlpEvaluationInput{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DlpEvaluationInput) GetAccessScope()(*AccessScope) {
    if m == nil {
        return nil
    } else {
        return m.accessScope
    }
}
func (m *DlpEvaluationInput) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DlpEvaluationInput) GetCurrentLabel()(*CurrentLabel) {
    if m == nil {
        return nil
    } else {
        return m.currentLabel
    }
}
func (m *DlpEvaluationInput) GetDiscoveredSensitiveTypes()([]DiscoveredSensitiveType) {
    if m == nil {
        return nil
    } else {
        return m.discoveredSensitiveTypes
    }
}
func (m *DlpEvaluationInput) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessScope)
        if err != nil {
            return err
        }
        cast := val.(AccessScope)
        m.SetAccessScope(&cast)
        return nil
    }
    res["currentLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrentLabel() })
        if err != nil {
            return err
        }
        m.SetCurrentLabel(val.(*CurrentLabel))
        return nil
    }
    res["discoveredSensitiveTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDiscoveredSensitiveType() })
        if err != nil {
            return err
        }
        res := make([]DiscoveredSensitiveType, len(val))
        for i, v := range val {
            res[i] = *(v.(*DiscoveredSensitiveType))
        }
        m.SetDiscoveredSensitiveTypes(res)
        return nil
    }
    return res
}
func (m *DlpEvaluationInput) IsNil()(bool) {
    return m == nil
}
func (m *DlpEvaluationInput) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAccessScope() != nil {
        cast := m.GetAccessScope().String()
        err := writer.WriteStringValue("accessScope", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("currentLabel", m.GetCurrentLabel())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDiscoveredSensitiveTypes()))
        for i, v := range m.GetDiscoveredSensitiveTypes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("discoveredSensitiveTypes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DlpEvaluationInput) SetAccessScope(value *AccessScope)() {
    m.accessScope = value
}
func (m *DlpEvaluationInput) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DlpEvaluationInput) SetCurrentLabel(value *CurrentLabel)() {
    m.currentLabel = value
}
func (m *DlpEvaluationInput) SetDiscoveredSensitiveTypes(value []DiscoveredSensitiveType)() {
    m.discoveredSensitiveTypes = value
}
