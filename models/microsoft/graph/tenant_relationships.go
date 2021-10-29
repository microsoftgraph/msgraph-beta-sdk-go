package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TenantRelationships struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The operations available to interact with the multi-tenant management platform.
    managedTenants *ManagedTenant;
}
// Instantiates a new tenantRelationships and sets the default values.
func NewTenantRelationships()(*TenantRelationships) {
    m := &TenantRelationships{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantRelationships) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the managedTenants property value. The operations available to interact with the multi-tenant management platform.
func (m *TenantRelationships) GetManagedTenants()(*ManagedTenant) {
    if m == nil {
        return nil
    } else {
        return m.managedTenants
    }
}
// The deserialization information for the current model
func (m *TenantRelationships) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["managedTenants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedTenant() })
        if err != nil {
            return err
        }
        m.SetManagedTenants(val.(*ManagedTenant))
        return nil
    }
    return res
}
func (m *TenantRelationships) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TenantRelationships) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("managedTenants", m.GetManagedTenants())
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
func (m *TenantRelationships) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the managedTenants property value. The operations available to interact with the multi-tenant management platform.
// Parameters:
//  - value : Value to set for the managedTenants property.
func (m *TenantRelationships) SetManagedTenants(value *ManagedTenant)() {
    m.managedTenants = value
}
