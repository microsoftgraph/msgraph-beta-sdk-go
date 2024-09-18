package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AppManagementApplicationConfiguration struct {
    AppManagementConfiguration
}
// NewAppManagementApplicationConfiguration instantiates a new AppManagementApplicationConfiguration and sets the default values.
func NewAppManagementApplicationConfiguration()(*AppManagementApplicationConfiguration) {
    m := &AppManagementApplicationConfiguration{
        AppManagementConfiguration: *NewAppManagementConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.appManagementApplicationConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAppManagementApplicationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppManagementApplicationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppManagementApplicationConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppManagementApplicationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppManagementConfiguration.GetFieldDeserializers()
    res["identifierUris"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentifierUriConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifierUris(val.(IdentifierUriConfigurationable))
        }
        return nil
    }
    return res
}
// GetIdentifierUris gets the identifierUris property value. Configuration object for restrictions on identifierUris property for an application
// returns a IdentifierUriConfigurationable when successful
func (m *AppManagementApplicationConfiguration) GetIdentifierUris()(IdentifierUriConfigurationable) {
    val, err := m.GetBackingStore().Get("identifierUris")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentifierUriConfigurationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AppManagementApplicationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AppManagementConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("identifierUris", m.GetIdentifierUris())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentifierUris sets the identifierUris property value. Configuration object for restrictions on identifierUris property for an application
func (m *AppManagementApplicationConfiguration) SetIdentifierUris(value IdentifierUriConfigurationable)() {
    err := m.GetBackingStore().Set("identifierUris", value)
    if err != nil {
        panic(err)
    }
}
type AppManagementApplicationConfigurationable interface {
    AppManagementConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentifierUris()(IdentifierUriConfigurationable)
    SetIdentifierUris(value IdentifierUriConfigurationable)()
}
