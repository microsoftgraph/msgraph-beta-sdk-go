package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SecurityBaselineContributingPolicy struct {
    additionalData map[string]interface{};
    displayName *string;
    sourceId *string;
    sourceType *SecurityBaselinePolicySourceType;
}
func NewSecurityBaselineContributingPolicy()(*SecurityBaselineContributingPolicy) {
    m := &SecurityBaselineContributingPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SecurityBaselineContributingPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SecurityBaselineContributingPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *SecurityBaselineContributingPolicy) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
func (m *SecurityBaselineContributingPolicy) GetSourceType()(*SecurityBaselinePolicySourceType) {
    if m == nil {
        return nil
    } else {
        return m.sourceType
    }
}
func (m *SecurityBaselineContributingPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["sourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSourceId(val)
        return nil
    }
    res["sourceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityBaselinePolicySourceType)
        if err != nil {
            return err
        }
        cast := val.(SecurityBaselinePolicySourceType)
        m.SetSourceType(&cast)
        return nil
    }
    return res
}
func (m *SecurityBaselineContributingPolicy) IsNil()(bool) {
    return m == nil
}
func (m *SecurityBaselineContributingPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceId", m.GetSourceId())
        if err != nil {
            return err
        }
    }
    if m.GetSourceType() != nil {
        cast := m.GetSourceType().String()
        err := writer.WriteStringValue("sourceType", &cast)
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
func (m *SecurityBaselineContributingPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SecurityBaselineContributingPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *SecurityBaselineContributingPolicy) SetSourceId(value *string)() {
    m.sourceId = value
}
func (m *SecurityBaselineContributingPolicy) SetSourceType(value *SecurityBaselinePolicySourceType)() {
    m.sourceType = value
}
