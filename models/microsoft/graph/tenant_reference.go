package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantReference 
type TenantReference struct {
    DirectoryObject
    // 
    tenantId *string;
}
// NewTenantReference instantiates a new tenantReference and sets the default values.
func NewTenantReference()(*TenantReference) {
    m := &TenantReference{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// GetTenantId gets the tenantId property value. 
func (m *TenantReference) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantReference) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
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
func (m *TenantReference) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TenantReference) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTenantId sets the tenantId property value. 
func (m *TenantReference) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
