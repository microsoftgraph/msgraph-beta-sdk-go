package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CustomTaskExtensionCallbackConfiguration 
type CustomTaskExtensionCallbackConfiguration struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionCallbackConfiguration
}
// NewCustomTaskExtensionCallbackConfiguration instantiates a new CustomTaskExtensionCallbackConfiguration and sets the default values.
func NewCustomTaskExtensionCallbackConfiguration()(*CustomTaskExtensionCallbackConfiguration) {
    m := &CustomTaskExtensionCallbackConfiguration{
        CustomExtensionCallbackConfiguration: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewCustomExtensionCallbackConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.customTaskExtensionCallbackConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomTaskExtensionCallbackConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomTaskExtensionCallbackConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomTaskExtensionCallbackConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomTaskExtensionCallbackConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomExtensionCallbackConfiguration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *CustomTaskExtensionCallbackConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomExtensionCallbackConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// CustomTaskExtensionCallbackConfigurationable 
type CustomTaskExtensionCallbackConfigurationable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionCallbackConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
