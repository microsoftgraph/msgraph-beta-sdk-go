package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceProvisioningResourceErrorDetailCollectionResponse 
type ServiceProvisioningResourceErrorDetailCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewServiceProvisioningResourceErrorDetailCollectionResponse instantiates a new ServiceProvisioningResourceErrorDetailCollectionResponse and sets the default values.
func NewServiceProvisioningResourceErrorDetailCollectionResponse()(*ServiceProvisioningResourceErrorDetailCollectionResponse) {
    m := &ServiceProvisioningResourceErrorDetailCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateServiceProvisioningResourceErrorDetailCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceProvisioningResourceErrorDetailCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceProvisioningResourceErrorDetailCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceProvisioningResourceErrorDetailCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceProvisioningResourceErrorDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceProvisioningResourceErrorDetailable, len(val))
            for i, v := range val {
                res[i] = v.(ServiceProvisioningResourceErrorDetailable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ServiceProvisioningResourceErrorDetailCollectionResponse) GetValue()([]ServiceProvisioningResourceErrorDetailable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ServiceProvisioningResourceErrorDetailable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ServiceProvisioningResourceErrorDetailCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ServiceProvisioningResourceErrorDetailCollectionResponse) SetValue(value []ServiceProvisioningResourceErrorDetailable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// ServiceProvisioningResourceErrorDetailCollectionResponseable 
type ServiceProvisioningResourceErrorDetailCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ServiceProvisioningResourceErrorDetailable)
    SetValue(value []ServiceProvisioningResourceErrorDetailable)()
}
