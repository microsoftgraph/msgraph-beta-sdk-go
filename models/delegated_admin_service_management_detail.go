package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DelegatedAdminServiceManagementDetail 
type DelegatedAdminServiceManagementDetail struct {
    Entity
    // The serviceId property
    serviceId *string
    // The URL of the management portal for the managed service. Read-only.
    serviceManagementUrl *string
    // The name of a managed service. Read-only.
    serviceName *string
}
// NewDelegatedAdminServiceManagementDetail instantiates a new delegatedAdminServiceManagementDetail and sets the default values.
func NewDelegatedAdminServiceManagementDetail()(*DelegatedAdminServiceManagementDetail) {
    m := &DelegatedAdminServiceManagementDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDelegatedAdminServiceManagementDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminServiceManagementDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDelegatedAdminServiceManagementDetail(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminServiceManagementDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["serviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceId(val)
        }
        return nil
    }
    res["serviceManagementUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceManagementUrl(val)
        }
        return nil
    }
    res["serviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetServiceId gets the serviceId property value. The serviceId property
func (m *DelegatedAdminServiceManagementDetail) GetServiceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceId
    }
}
// GetServiceManagementUrl gets the serviceManagementUrl property value. The URL of the management portal for the managed service. Read-only.
func (m *DelegatedAdminServiceManagementDetail) GetServiceManagementUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceManagementUrl
    }
}
// GetServiceName gets the serviceName property value. The name of a managed service. Read-only.
func (m *DelegatedAdminServiceManagementDetail) GetServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceName
    }
}
// Serialize serializes information the current object
func (m *DelegatedAdminServiceManagementDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetServiceId sets the serviceId property value. The serviceId property
func (m *DelegatedAdminServiceManagementDetail) SetServiceId(value *string)() {
    if m != nil {
        m.serviceId = value
    }
}
// SetServiceManagementUrl sets the serviceManagementUrl property value. The URL of the management portal for the managed service. Read-only.
func (m *DelegatedAdminServiceManagementDetail) SetServiceManagementUrl(value *string)() {
    if m != nil {
        m.serviceManagementUrl = value
    }
}
// SetServiceName sets the serviceName property value. The name of a managed service. Read-only.
func (m *DelegatedAdminServiceManagementDetail) SetServiceName(value *string)() {
    if m != nil {
        m.serviceName = value
    }
}
