package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DelegatedAdminCustomer 
type DelegatedAdminCustomer struct {
    Entity
    // 
    displayName *string;
    // 
    serviceManagementDetails []DelegatedAdminServiceManagementDetailable;
    // 
    tenantId *string;
}
// NewDelegatedAdminCustomer instantiates a new delegatedAdminCustomer and sets the default values.
func NewDelegatedAdminCustomer()(*DelegatedAdminCustomer) {
    m := &DelegatedAdminCustomer{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDelegatedAdminCustomerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminCustomerFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDelegatedAdminCustomer(), nil
}
// GetDisplayName gets the displayName property value. 
func (m *DelegatedAdminCustomer) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminCustomer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["serviceManagementDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegatedAdminServiceManagementDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminServiceManagementDetailable, len(val))
            for i, v := range val {
                res[i] = v.(DelegatedAdminServiceManagementDetailable)
            }
            m.SetServiceManagementDetails(res)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetServiceManagementDetails gets the serviceManagementDetails property value. 
func (m *DelegatedAdminCustomer) GetServiceManagementDetails()([]DelegatedAdminServiceManagementDetailable) {
    if m == nil {
        return nil
    } else {
        return m.serviceManagementDetails
    }
}
// GetTenantId gets the tenantId property value. 
func (m *DelegatedAdminCustomer) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Serialize serializes information the current object
func (m *DelegatedAdminCustomer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetServiceManagementDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetServiceManagementDetails()))
        for i, v := range m.GetServiceManagementDetails() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("serviceManagementDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. 
func (m *DelegatedAdminCustomer) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetServiceManagementDetails sets the serviceManagementDetails property value. 
func (m *DelegatedAdminCustomer) SetServiceManagementDetails(value []DelegatedAdminServiceManagementDetailable)() {
    if m != nil {
        m.serviceManagementDetails = value
    }
}
// SetTenantId sets the tenantId property value. 
func (m *DelegatedAdminCustomer) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
