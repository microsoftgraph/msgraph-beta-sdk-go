package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AuthenticationRequirementPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    detail *string;
    // 
    requirementProvider *RequirementProvider;
}
// Instantiates a new authenticationRequirementPolicy and sets the default values.
func NewAuthenticationRequirementPolicy()(*AuthenticationRequirementPolicy) {
    m := &AuthenticationRequirementPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationRequirementPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the detail property value. 
func (m *AuthenticationRequirementPolicy) GetDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// Gets the requirementProvider property value. 
func (m *AuthenticationRequirementPolicy) GetRequirementProvider()(*RequirementProvider) {
    if m == nil {
        return nil
    } else {
        return m.requirementProvider
    }
}
// The deserialization information for the current model
func (m *AuthenticationRequirementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val)
        }
        return nil
    }
    res["requirementProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequirementProvider)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RequirementProvider)
            m.SetRequirementProvider(&cast)
        }
        return nil
    }
    return res
}
func (m *AuthenticationRequirementPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AuthenticationRequirementPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the detail property value. 
// Parameters:
//  - value : Value to set for the detail property.
func (m *AuthenticationRequirementPolicy) SetDetail(value *string)() {
    m.detail = value
}
// Sets the requirementProvider property value. 
// Parameters:
//  - value : Value to set for the requirementProvider property.
func (m *AuthenticationRequirementPolicy) SetRequirementProvider(value *RequirementProvider)() {
    m.requirementProvider = value
}
