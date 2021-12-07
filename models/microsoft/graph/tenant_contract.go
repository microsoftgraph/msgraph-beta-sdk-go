package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantContract 
type TenantContract struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The type of relationship that exists between the managing entity and tenant. Optional. Read-only.
    contractType *int32;
    // The default domain name for the tenant. Required. Read-only.
    defaultDomainName *string;
    // The display name for the tenant. Optional. Read-only.
    displayName *string;
}
// NewTenantContract instantiates a new tenantContract and sets the default values.
func NewTenantContract()(*TenantContract) {
    m := &TenantContract{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantContract) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContractType gets the contractType property value. The type of relationship that exists between the managing entity and tenant. Optional. Read-only.
func (m *TenantContract) GetContractType()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.contractType
    }
}
// GetDefaultDomainName gets the defaultDomainName property value. The default domain name for the tenant. Required. Read-only.
func (m *TenantContract) GetDefaultDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDomainName
    }
}
// GetDisplayName gets the displayName property value. The display name for the tenant. Optional. Read-only.
func (m *TenantContract) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantContract) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contractType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContractType(val)
        }
        return nil
    }
    res["defaultDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDomainName(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
func (m *TenantContract) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TenantContract) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("contractType", m.GetContractType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultDomainName", m.GetDefaultDomainName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
func (m *TenantContract) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContractType sets the contractType property value. The type of relationship that exists between the managing entity and tenant. Optional. Read-only.
func (m *TenantContract) SetContractType(value *int32)() {
    if m != nil {
        m.contractType = value
    }
}
// SetDefaultDomainName sets the defaultDomainName property value. The default domain name for the tenant. Required. Read-only.
func (m *TenantContract) SetDefaultDomainName(value *string)() {
    if m != nil {
        m.defaultDomainName = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the tenant. Optional. Read-only.
func (m *TenantContract) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
