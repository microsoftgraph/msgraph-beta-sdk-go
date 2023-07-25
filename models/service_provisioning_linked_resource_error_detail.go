package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceProvisioningLinkedResourceErrorDetail 
type ServiceProvisioningLinkedResourceErrorDetail struct {
    ServiceProvisioningResourceErrorDetail
}
// NewServiceProvisioningLinkedResourceErrorDetail instantiates a new serviceProvisioningLinkedResourceErrorDetail and sets the default values.
func NewServiceProvisioningLinkedResourceErrorDetail()(*ServiceProvisioningLinkedResourceErrorDetail) {
    m := &ServiceProvisioningLinkedResourceErrorDetail{
        ServiceProvisioningResourceErrorDetail: *NewServiceProvisioningResourceErrorDetail(),
    }
    return m
}
// CreateServiceProvisioningLinkedResourceErrorDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceProvisioningLinkedResourceErrorDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceProvisioningLinkedResourceErrorDetail(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceProvisioningLinkedResourceErrorDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ServiceProvisioningResourceErrorDetail.GetFieldDeserializers()
    res["propertyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyName(val)
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
// GetPropertyName gets the propertyName property value. The propertyName property
func (m *ServiceProvisioningLinkedResourceErrorDetail) GetPropertyName()(*string) {
    val, err := m.GetBackingStore().Get("propertyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTarget gets the target property value. The target property
func (m *ServiceProvisioningLinkedResourceErrorDetail) GetTarget()(*string) {
    val, err := m.GetBackingStore().Get("target")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ServiceProvisioningLinkedResourceErrorDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ServiceProvisioningResourceErrorDetail.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("propertyName", m.GetPropertyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPropertyName sets the propertyName property value. The propertyName property
func (m *ServiceProvisioningLinkedResourceErrorDetail) SetPropertyName(value *string)() {
    err := m.GetBackingStore().Set("propertyName", value)
    if err != nil {
        panic(err)
    }
}
// SetTarget sets the target property value. The target property
func (m *ServiceProvisioningLinkedResourceErrorDetail) SetTarget(value *string)() {
    err := m.GetBackingStore().Set("target", value)
    if err != nil {
        panic(err)
    }
}
// ServiceProvisioningLinkedResourceErrorDetailable 
type ServiceProvisioningLinkedResourceErrorDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceProvisioningResourceErrorDetailable
    GetPropertyName()(*string)
    GetTarget()(*string)
    SetPropertyName(value *string)()
    SetTarget(value *string)()
}
