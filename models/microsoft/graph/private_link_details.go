package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivateLinkDetails 
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
// NewPrivateLinkDetails instantiates a new privateLinkDetails and sets the default values.
func NewPrivateLinkDetails()(*PrivateLinkDetails) {
    m := &PrivateLinkDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivateLinkDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPolicyId gets the policyId property value. 
func (m *PrivateLinkDetails) GetPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyId
    }
}
// GetPolicyName gets the policyName property value. 
func (m *PrivateLinkDetails) GetPolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyName
    }
}
// GetPolicyTenantId gets the policyTenantId property value. 
func (m *PrivateLinkDetails) GetPolicyTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyTenantId
    }
}
// GetResourceId gets the resourceId property value. 
func (m *PrivateLinkDetails) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivateLinkDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["policyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
        }
        return nil
    }
    res["policyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyName(val)
        }
        return nil
    }
    res["policyTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyTenantId(val)
        }
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    return res
}
func (m *PrivateLinkDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivateLinkDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPolicyId sets the policyId property value. 
func (m *PrivateLinkDetails) SetPolicyId(value *string)() {
    m.policyId = value
}
// SetPolicyName sets the policyName property value. 
func (m *PrivateLinkDetails) SetPolicyName(value *string)() {
    m.policyName = value
}
// SetPolicyTenantId sets the policyTenantId property value. 
func (m *PrivateLinkDetails) SetPolicyTenantId(value *string)() {
    m.policyTenantId = value
}
// SetResourceId sets the resourceId property value. 
func (m *PrivateLinkDetails) SetResourceId(value *string)() {
    m.resourceId = value
}
