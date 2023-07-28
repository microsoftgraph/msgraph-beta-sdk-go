package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceLevelAgreementRoot 
type ServiceLevelAgreementRoot struct {
    Entity
}
// NewServiceLevelAgreementRoot instantiates a new serviceLevelAgreementRoot and sets the default values.
func NewServiceLevelAgreementRoot()(*ServiceLevelAgreementRoot) {
    m := &ServiceLevelAgreementRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateServiceLevelAgreementRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceLevelAgreementRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceLevelAgreementRoot(), nil
}
// GetAzureADAuthentication gets the azureADAuthentication property value. Collects the Azure AD SLA attainment for each month for an Azure AD tenant.
func (m *ServiceLevelAgreementRoot) GetAzureADAuthentication()(AzureADAuthenticationable) {
    val, err := m.GetBackingStore().Get("azureADAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AzureADAuthenticationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceLevelAgreementRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureADAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAzureADAuthenticationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADAuthentication(val.(AzureADAuthenticationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ServiceLevelAgreementRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("azureADAuthentication", m.GetAzureADAuthentication())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureADAuthentication sets the azureADAuthentication property value. Collects the Azure AD SLA attainment for each month for an Azure AD tenant.
func (m *ServiceLevelAgreementRoot) SetAzureADAuthentication(value AzureADAuthenticationable)() {
    err := m.GetBackingStore().Set("azureADAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// ServiceLevelAgreementRootable 
type ServiceLevelAgreementRootable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureADAuthentication()(AzureADAuthenticationable)
    SetAzureADAuthentication(value AzureADAuthenticationable)()
}
