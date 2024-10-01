package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomAppManagementConfiguration struct {
    AppManagementConfiguration
}
// NewCustomAppManagementConfiguration instantiates a new CustomAppManagementConfiguration and sets the default values.
func NewCustomAppManagementConfiguration()(*CustomAppManagementConfiguration) {
    m := &CustomAppManagementConfiguration{
        AppManagementConfiguration: *NewAppManagementConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.customAppManagementConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomAppManagementConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomAppManagementConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomAppManagementConfiguration(), nil
}
// GetApplicationRestrictions gets the applicationRestrictions property value. Restrictions applicable only to application objects that the policy applies to.
// returns a CustomAppManagementApplicationConfigurationable when successful
func (m *CustomAppManagementConfiguration) GetApplicationRestrictions()(CustomAppManagementApplicationConfigurationable) {
    val, err := m.GetBackingStore().Get("applicationRestrictions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CustomAppManagementApplicationConfigurationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomAppManagementConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppManagementConfiguration.GetFieldDeserializers()
    res["applicationRestrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomAppManagementApplicationConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationRestrictions(val.(CustomAppManagementApplicationConfigurationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CustomAppManagementConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AppManagementConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("applicationRestrictions", m.GetApplicationRestrictions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationRestrictions sets the applicationRestrictions property value. Restrictions applicable only to application objects that the policy applies to.
func (m *CustomAppManagementConfiguration) SetApplicationRestrictions(value CustomAppManagementApplicationConfigurationable)() {
    err := m.GetBackingStore().Set("applicationRestrictions", value)
    if err != nil {
        panic(err)
    }
}
type CustomAppManagementConfigurationable interface {
    AppManagementConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationRestrictions()(CustomAppManagementApplicationConfigurationable)
    SetApplicationRestrictions(value CustomAppManagementApplicationConfigurationable)()
}
