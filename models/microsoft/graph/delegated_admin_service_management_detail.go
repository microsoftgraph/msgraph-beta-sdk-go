package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DelegatedAdminServiceManagementDetail 
type DelegatedAdminServiceManagementDetail struct {
    Entity
    // 
    serviceId *string;
    // 
    serviceManagementUrl *string;
    // 
    serviceName *string;
}
// NewDelegatedAdminServiceManagementDetail instantiates a new delegatedAdminServiceManagementDetail and sets the default values.
func NewDelegatedAdminServiceManagementDetail()(*DelegatedAdminServiceManagementDetail) {
    m := &DelegatedAdminServiceManagementDetail{
        Entity: *NewEntity(),
    }
    return m
}
// GetServiceId gets the serviceId property value. 
func (m *DelegatedAdminServiceManagementDetail) GetServiceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceId
    }
}
// GetServiceManagementUrl gets the serviceManagementUrl property value. 
func (m *DelegatedAdminServiceManagementDetail) GetServiceManagementUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceManagementUrl
    }
}
// GetServiceName gets the serviceName property value. 
func (m *DelegatedAdminServiceManagementDetail) GetServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminServiceManagementDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["serviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceId(val)
        }
        return nil
    }
    res["serviceManagementUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceManagementUrl(val)
        }
        return nil
    }
    res["serviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceName(val)
        }
        return nil
    }
    return res
}
func (m *DelegatedAdminServiceManagementDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DelegatedAdminServiceManagementDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("serviceId", m.GetServiceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceManagementUrl", m.GetServiceManagementUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceName", m.GetServiceName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetServiceId sets the serviceId property value. 
func (m *DelegatedAdminServiceManagementDetail) SetServiceId(value *string)() {
    if m != nil {
        m.serviceId = value
    }
}
// SetServiceManagementUrl sets the serviceManagementUrl property value. 
func (m *DelegatedAdminServiceManagementDetail) SetServiceManagementUrl(value *string)() {
    if m != nil {
        m.serviceManagementUrl = value
    }
}
// SetServiceName sets the serviceName property value. 
func (m *DelegatedAdminServiceManagementDetail) SetServiceName(value *string)() {
    if m != nil {
        m.serviceName = value
    }
}
