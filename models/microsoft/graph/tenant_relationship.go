package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantRelationship provides operations to manage the tenantRelationship singleton.
type TenantRelationship struct {
    Entity
    // 
    delegatedAdminCustomers []DelegatedAdminCustomerable;
    // 
    delegatedAdminRelationships []DelegatedAdminRelationshipable;
    // The operations available to interact with the multi-tenant management platform.
    managedTenants ManagedTenantable;
}
// NewTenantRelationship instantiates a new tenantRelationship and sets the default values.
func NewTenantRelationship()(*TenantRelationship) {
    m := &TenantRelationship{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTenantRelationshipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantRelationshipFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTenantRelationship(), nil
}
// GetDelegatedAdminCustomers gets the delegatedAdminCustomers property value. 
func (m *TenantRelationship) GetDelegatedAdminCustomers()([]DelegatedAdminCustomerable) {
    if m == nil {
        return nil
    } else {
        return m.delegatedAdminCustomers
    }
}
// GetDelegatedAdminRelationships gets the delegatedAdminRelationships property value. 
func (m *TenantRelationship) GetDelegatedAdminRelationships()([]DelegatedAdminRelationshipable) {
    if m == nil {
        return nil
    } else {
        return m.delegatedAdminRelationships
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantRelationship) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["delegatedAdminCustomers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegatedAdminCustomerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminCustomerable, len(val))
            for i, v := range val {
                res[i] = v.(DelegatedAdminCustomerable)
            }
            m.SetDelegatedAdminCustomers(res)
        }
        return nil
    }
    res["delegatedAdminRelationships"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegatedAdminRelationshipFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminRelationshipable, len(val))
            for i, v := range val {
                res[i] = v.(DelegatedAdminRelationshipable)
            }
            m.SetDelegatedAdminRelationships(res)
        }
        return nil
    }
    res["managedTenants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedTenantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedTenants(val.(ManagedTenantable))
        }
        return nil
    }
    return res
}
// GetManagedTenants gets the managedTenants property value. The operations available to interact with the multi-tenant management platform.
func (m *TenantRelationship) GetManagedTenants()(ManagedTenantable) {
    if m == nil {
        return nil
    } else {
        return m.managedTenants
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("delegatedAdminCustomers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDelegatedAdminRelationships() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDelegatedAdminRelationships()))
        for i, v := range m.GetDelegatedAdminRelationships() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *TenantRelationship) SetDelegatedAdminCustomers(value []DelegatedAdminCustomerable)() {
    if m != nil {
        m.delegatedAdminCustomers = value
    }
}
// SetDelegatedAdminRelationships sets the delegatedAdminRelationships property value. 
func (m *TenantRelationship) SetDelegatedAdminRelationships(value []DelegatedAdminRelationshipable)() {
    if m != nil {
        m.delegatedAdminRelationships = value
    }
}
// SetManagedTenants sets the managedTenants property value. The operations available to interact with the multi-tenant management platform.
func (m *TenantRelationship) SetManagedTenants(value ManagedTenantable)() {
    if m != nil {
        m.managedTenants = value
    }
}
