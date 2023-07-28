package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceProvisioningResourceError 
type ServiceProvisioningResourceError struct {
    ServiceProvisioningError
}
// NewServiceProvisioningResourceError instantiates a new serviceProvisioningResourceError and sets the default values.
func NewServiceProvisioningResourceError()(*ServiceProvisioningResourceError) {
    m := &ServiceProvisioningResourceError{
        ServiceProvisioningError: *NewServiceProvisioningError(),
    }
    odataTypeValue := "#microsoft.graph.serviceProvisioningResourceError"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateServiceProvisioningResourceErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceProvisioningResourceErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceProvisioningResourceError(), nil
}
// GetErrors gets the errors property value. The errors property
func (m *ServiceProvisioningResourceError) GetErrors()([]ServiceProvisioningResourceErrorDetailable) {
    val, err := m.GetBackingStore().Get("errors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ServiceProvisioningResourceErrorDetailable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceProvisioningResourceError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ServiceProvisioningError.GetFieldDeserializers()
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceProvisioningResourceErrorDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceProvisioningResourceErrorDetailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ServiceProvisioningResourceErrorDetailable)
                }
            }
            m.SetErrors(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ServiceProvisioningResourceError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ServiceProvisioningError.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetErrors()))
        for i, v := range m.GetErrors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("errors", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetErrors sets the errors property value. The errors property
func (m *ServiceProvisioningResourceError) SetErrors(value []ServiceProvisioningResourceErrorDetailable)() {
    err := m.GetBackingStore().Set("errors", value)
    if err != nil {
        panic(err)
    }
}
// ServiceProvisioningResourceErrorable 
type ServiceProvisioningResourceErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceProvisioningErrorable
    GetErrors()([]ServiceProvisioningResourceErrorDetailable)
    SetErrors(value []ServiceProvisioningResourceErrorDetailable)()
}
