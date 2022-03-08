package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutboundSharedUserProfile provides operations to manage the directory singleton.
type OutboundSharedUserProfile struct {
    DirectoryObject
    // 
    tenants []TenantReferenceable;
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
// CreateOutboundSharedUserProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutboundSharedUserProfileFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOutboundSharedUserProfile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutboundSharedUserProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["tenants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantReferenceable, len(val))
            for i, v := range val {
                res[i] = v.(TenantReferenceable)
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
// GetTenants gets the tenants property value. 
func (m *OutboundSharedUserProfile) GetTenants()([]TenantReferenceable) {
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *OutboundSharedUserProfile) SetTenants(value []TenantReferenceable)() {
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
