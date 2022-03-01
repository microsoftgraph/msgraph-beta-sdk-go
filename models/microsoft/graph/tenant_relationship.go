package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantRelationship 
type TenantRelationship struct {
    Entity
    // 
    delegatedAdminCustomers []DelegatedAdminCustomer;
    // 
    delegatedAdminRelationships []DelegatedAdminRelationship;
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
// GetDelegatedAdminCustomers gets the delegatedAdminCustomers property value. 
func (m *TenantRelationship) GetDelegatedAdminCustomers()([]DelegatedAdminCustomer) {
    if m == nil {
        return nil
    } else {
        return m.delegatedAdminCustomers
    }
}
// GetDelegatedAdminRelationships gets the delegatedAdminRelationships property value. 
func (m *TenantRelationship) GetDelegatedAdminRelationships()([]DelegatedAdminRelationship) {
    if m == nil {
        return nil
    } else {
        return m.delegatedAdminRelationships
    }
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
    res["delegatedAdminCustomers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDelegatedAdminCustomer() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminCustomer, len(val))
            for i, v := range val {
                res[i] = *(v.(*DelegatedAdminCustomer))
            }
            m.SetDelegatedAdminCustomers(res)
        }
        return nil
    }
    res["delegatedAdminRelationships"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDelegatedAdminRelationship() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminRelationship, len(val))
            for i, v := range val {
                res[i] = *(v.(*DelegatedAdminRelationship))
            }
            m.SetDelegatedAdminRelationships(res)
        }
        return nil
    }
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
    if m.GetDelegatedAdminCustomers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDelegatedAdminCustomers()))
        for i, v := range m.GetDelegatedAdminCustomers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("delegatedAdminCustomers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDelegatedAdminRelationships() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDelegatedAdminRelationships()))
        for i, v := range m.GetDelegatedAdminRelationships() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("delegatedAdminRelationships", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managedTenants", m.GetManagedTenants())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDelegatedAdminCustomers sets the delegatedAdminCustomers property value. 
func (m *TenantRelationship) SetDelegatedAdminCustomers(value []DelegatedAdminCustomer)() {
    if m != nil {
        m.delegatedAdminCustomers = value
    }
}
// SetDelegatedAdminRelationships sets the delegatedAdminRelationships property value. 
func (m *TenantRelationship) SetDelegatedAdminRelationships(value []DelegatedAdminRelationship)() {
    if m != nil {
        m.delegatedAdminRelationships = value
    }
}
// SetManagedTenants sets the managedTenants property value. The operations available to interact with the multi-tenant management platform.
func (m *TenantRelationship) SetManagedTenants(value *ManagedTenant)() {
    if m != nil {
        m.managedTenants = value
    }
}
