package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantRelationship 
type TenantRelationship struct {
    Entity
    // The operations available to interact with the multi-tenant management platform.
    managedTenants *ManagedTenant;
}
// NewTenantRelationship instantiates a new tenantRelationship and sets the default values.
func NewTenantRelationship()(*TenantRelationship) {
    m := &TenantRelationship{
        Entity: *NewEntity(),
    }
    return m
}
// GetManagedTenants gets the managedTenants property value. The operations available to interact with the multi-tenant management platform.
func (m *TenantRelationship) GetManagedTenants()(*ManagedTenant) {
    if m == nil {
        return nil
    } else {
        return m.managedTenants
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantRelationship) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["managedTenants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedTenant() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedTenants(val.(*ManagedTenant))
        }
        return nil
    }
    return res
}
func (m *TenantRelationship) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TenantRelationship) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("managedTenants", m.GetManagedTenants())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetManagedTenants sets the managedTenants property value. The operations available to interact with the multi-tenant management platform.
func (m *TenantRelationship) SetManagedTenants(value *ManagedTenant)() {
    if m != nil {
        m.managedTenants = value
    }
}
