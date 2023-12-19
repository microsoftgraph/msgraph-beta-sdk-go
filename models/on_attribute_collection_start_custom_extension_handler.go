package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnAttributeCollectionStartCustomExtensionHandler 
type OnAttributeCollectionStartCustomExtensionHandler struct {
    OnAttributeCollectionStartHandler
}
// NewOnAttributeCollectionStartCustomExtensionHandler instantiates a new onAttributeCollectionStartCustomExtensionHandler and sets the default values.
func NewOnAttributeCollectionStartCustomExtensionHandler()(*OnAttributeCollectionStartCustomExtensionHandler) {
    m := &OnAttributeCollectionStartCustomExtensionHandler{
        OnAttributeCollectionStartHandler: *NewOnAttributeCollectionStartHandler(),
    }
    odataTypeValue := "#microsoft.graph.onAttributeCollectionStartCustomExtensionHandler"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnAttributeCollectionStartCustomExtensionHandlerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnAttributeCollectionStartCustomExtensionHandlerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnAttributeCollectionStartCustomExtensionHandler(), nil
}
// GetConfiguration gets the configuration property value. The configuration property
func (m *OnAttributeCollectionStartCustomExtensionHandler) GetConfiguration()(CustomExtensionOverwriteConfigurationable) {
    val, err := m.GetBackingStore().Get("configuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CustomExtensionOverwriteConfigurationable)
    }
    return nil
}
// GetCustomExtension gets the customExtension property value. The customExtension property
func (m *OnAttributeCollectionStartCustomExtensionHandler) GetCustomExtension()(OnAttributeCollectionStartCustomExtensionable) {
    val, err := m.GetBackingStore().Get("customExtension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OnAttributeCollectionStartCustomExtensionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnAttributeCollectionStartCustomExtensionHandler) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnAttributeCollectionStartHandler.GetFieldDeserializers()
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomExtensionOverwriteConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(CustomExtensionOverwriteConfigurationable))
        }
        return nil
    }
    res["customExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnAttributeCollectionStartCustomExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtension(val.(OnAttributeCollectionStartCustomExtensionable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *OnAttributeCollectionStartCustomExtensionHandler) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnAttributeCollectionStartHandler.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("customExtension", m.GetCustomExtension())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfiguration sets the configuration property value. The configuration property
func (m *OnAttributeCollectionStartCustomExtensionHandler) SetConfiguration(value CustomExtensionOverwriteConfigurationable)() {
    err := m.GetBackingStore().Set("configuration", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomExtension sets the customExtension property value. The customExtension property
func (m *OnAttributeCollectionStartCustomExtensionHandler) SetCustomExtension(value OnAttributeCollectionStartCustomExtensionable)() {
    err := m.GetBackingStore().Set("customExtension", value)
    if err != nil {
        panic(err)
    }
}
// OnAttributeCollectionStartCustomExtensionHandlerable 
type OnAttributeCollectionStartCustomExtensionHandlerable interface {
    OnAttributeCollectionStartHandlerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfiguration()(CustomExtensionOverwriteConfigurationable)
    GetCustomExtension()(OnAttributeCollectionStartCustomExtensionable)
    SetConfiguration(value CustomExtensionOverwriteConfigurationable)()
    SetCustomExtension(value OnAttributeCollectionStartCustomExtensionable)()
}
