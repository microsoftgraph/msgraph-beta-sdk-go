package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AuthenticationRequirementPolicy struct {
    additionalData map[string]interface{};
    detail *string;
    requirementProvider *RequirementProvider;
}
func NewAuthenticationRequirementPolicy()(*AuthenticationRequirementPolicy) {
    m := &AuthenticationRequirementPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AuthenticationRequirementPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AuthenticationRequirementPolicy) GetDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
func (m *AuthenticationRequirementPolicy) GetRequirementProvider()(*RequirementProvider) {
    if m == nil {
        return nil
    } else {
        return m.requirementProvider
    }
}
func (m *AuthenticationRequirementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDetail(val)
        return nil
    }
    res["requirementProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequirementProvider)
        if err != nil {
            return err
        }
        cast := val.(RequirementProvider)
        m.SetRequirementProvider(&cast)
        return nil
    }
    return res
}
func (m *AuthenticationRequirementPolicy) IsNil()(bool) {
    return m == nil
}
func (m *AuthenticationRequirementPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    if m.GetRequirementProvider() != nil {
        cast := m.GetRequirementProvider().String()
        err := writer.WriteStringValue("requirementProvider", &cast)
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
func (m *AuthenticationRequirementPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AuthenticationRequirementPolicy) SetDetail(value *string)() {
    m.detail = value
}
func (m *AuthenticationRequirementPolicy) SetRequirementProvider(value *RequirementProvider)() {
    m.requirementProvider = value
}
