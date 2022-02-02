package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutboundSharedUserProfile 
type OutboundSharedUserProfile struct {
    DirectoryObject
    // 
    tenants []TenantReference;
    // 
    userId *string;
}
// NewOutboundSharedUserProfile instantiates a new outboundSharedUserProfile and sets the default values.
func NewOutboundSharedUserProfile()(*OutboundSharedUserProfile) {
    m := &OutboundSharedUserProfile{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// GetTenants gets the tenants property value. 
func (m *OutboundSharedUserProfile) GetTenants()([]TenantReference) {
    if m == nil {
        return nil
    } else {
        return m.tenants
    }
}
// GetUserId gets the userId property value. 
func (m *OutboundSharedUserProfile) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutboundSharedUserProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["tenants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTenantReference() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantReference, len(val))
            for i, v := range val {
                res[i] = *(v.(*TenantReference))
            }
            m.SetTenants(res)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *OutboundSharedUserProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OutboundSharedUserProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetTenants() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTenants()))
        for i, v := range m.GetTenants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tenants", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTenants sets the tenants property value. 
func (m *OutboundSharedUserProfile) SetTenants(value []TenantReference)() {
    if m != nil {
        m.tenants = value
    }
}
// SetUserId sets the userId property value. 
func (m *OutboundSharedUserProfile) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
