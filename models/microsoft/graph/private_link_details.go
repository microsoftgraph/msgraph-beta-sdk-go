package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrivateLinkDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    policyId *string;
    // 
    policyName *string;
    // 
    policyTenantId *string;
    // 
    resourceId *string;
}
// Instantiates a new privateLinkDetails and sets the default values.
func NewPrivateLinkDetails()(*PrivateLinkDetails) {
    m := &PrivateLinkDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivateLinkDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the policyId property value. 
func (m *PrivateLinkDetails) GetPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyId
    }
}
// Gets the policyName property value. 
func (m *PrivateLinkDetails) GetPolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyName
    }
}
// Gets the policyTenantId property value. 
func (m *PrivateLinkDetails) GetPolicyTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyTenantId
    }
}
// Gets the resourceId property value. 
func (m *PrivateLinkDetails) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// The deserialization information for the current model
func (m *PrivateLinkDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["policyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyId(val)
        return nil
    }
    res["policyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyName(val)
        return nil
    }
    res["policyTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyTenantId(val)
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceId(val)
        return nil
    }
    return res
}
func (m *PrivateLinkDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrivateLinkDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("policyId", m.GetPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyName", m.GetPolicyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyTenantId", m.GetPolicyTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceId", m.GetResourceId())
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
func (m *PrivateLinkDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the policyId property value. 
// Parameters:
//  - value : Value to set for the policyId property.
func (m *PrivateLinkDetails) SetPolicyId(value *string)() {
    m.policyId = value
}
// Sets the policyName property value. 
// Parameters:
//  - value : Value to set for the policyName property.
func (m *PrivateLinkDetails) SetPolicyName(value *string)() {
    m.policyName = value
}
// Sets the policyTenantId property value. 
// Parameters:
//  - value : Value to set for the policyTenantId property.
func (m *PrivateLinkDetails) SetPolicyTenantId(value *string)() {
    m.policyTenantId = value
}
// Sets the resourceId property value. 
// Parameters:
//  - value : Value to set for the resourceId property.
func (m *PrivateLinkDetails) SetResourceId(value *string)() {
    m.resourceId = value
}
