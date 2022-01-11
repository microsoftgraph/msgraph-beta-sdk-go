package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivateLinkDetails 
type PrivateLinkDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The unique identifier for the Private Link policy.
    policyId *string;
    // The name of the Private Link policy in Azure AD.
    policyName *string;
    // The tenant identifier of the Azure AD tenant the Private Link policy belongs to.
    policyTenantId *string;
    // The Azure Resource Manager (ARM) path for the Private Link policy resource.
    resourceId *string;
}
// NewPrivateLinkDetails instantiates a new privateLinkDetails and sets the default values.
func NewPrivateLinkDetails()(*PrivateLinkDetails) {
    m := &PrivateLinkDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivateLinkDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPolicyId gets the policyId property value. The unique identifier for the Private Link policy.
func (m *PrivateLinkDetails) GetPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyId
    }
}
// GetPolicyName gets the policyName property value. The name of the Private Link policy in Azure AD.
func (m *PrivateLinkDetails) GetPolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyName
    }
}
// GetPolicyTenantId gets the policyTenantId property value. The tenant identifier of the Azure AD tenant the Private Link policy belongs to.
func (m *PrivateLinkDetails) GetPolicyTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyTenantId
    }
}
// GetResourceId gets the resourceId property value. The Azure Resource Manager (ARM) path for the Private Link policy resource.
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivateLinkDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPolicyId sets the policyId property value. The unique identifier for the Private Link policy.
func (m *PrivateLinkDetails) SetPolicyId(value *string)() {
    if m != nil {
        m.policyId = value
    }
}
// SetPolicyName sets the policyName property value. The name of the Private Link policy in Azure AD.
func (m *PrivateLinkDetails) SetPolicyName(value *string)() {
    if m != nil {
        m.policyName = value
    }
}
// SetPolicyTenantId sets the policyTenantId property value. The tenant identifier of the Azure AD tenant the Private Link policy belongs to.
func (m *PrivateLinkDetails) SetPolicyTenantId(value *string)() {
    if m != nil {
        m.policyTenantId = value
    }
}
// SetResourceId sets the resourceId property value. The Azure Resource Manager (ARM) path for the Private Link policy resource.
func (m *PrivateLinkDetails) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
