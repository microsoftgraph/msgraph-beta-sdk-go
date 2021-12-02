package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InboundSharedUserProfile 
type InboundSharedUserProfile struct {
    DirectoryObject
    // 
    displayName *string;
    // 
    homeTenantId *string;
    // 
    userId *string;
    // 
    userPrincipalName *string;
}
// NewInboundSharedUserProfile instantiates a new inboundSharedUserProfile and sets the default values.
func NewInboundSharedUserProfile()(*InboundSharedUserProfile) {
    m := &InboundSharedUserProfile{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// GetDisplayName gets the displayName property value. 
func (m *InboundSharedUserProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetHomeTenantId gets the homeTenantId property value. 
func (m *InboundSharedUserProfile) GetHomeTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantId
    }
}
// GetUserId gets the userId property value. 
func (m *InboundSharedUserProfile) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. 
func (m *InboundSharedUserProfile) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InboundSharedUserProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
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
    res["homeTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantId(val)
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
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *InboundSharedUserProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InboundSharedUserProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("homeTenantId", m.GetHomeTenantId())
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
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. 
func (m *InboundSharedUserProfile) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetHomeTenantId sets the homeTenantId property value. 
func (m *InboundSharedUserProfile) SetHomeTenantId(value *string)() {
    if m != nil {
        m.homeTenantId = value
    }
}
// SetUserId sets the userId property value. 
func (m *InboundSharedUserProfile) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. 
func (m *InboundSharedUserProfile) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
