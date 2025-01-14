package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OnOtpSendCustomExtensionHandler struct {
    OnOtpSendHandler
}
// NewOnOtpSendCustomExtensionHandler instantiates a new OnOtpSendCustomExtensionHandler and sets the default values.
func NewOnOtpSendCustomExtensionHandler()(*OnOtpSendCustomExtensionHandler) {
    m := &OnOtpSendCustomExtensionHandler{
        OnOtpSendHandler: *NewOnOtpSendHandler(),
    }
    odataTypeValue := "#microsoft.graph.onOtpSendCustomExtensionHandler"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnOtpSendCustomExtensionHandlerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnOtpSendCustomExtensionHandlerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnOtpSendCustomExtensionHandler(), nil
}
// GetConfiguration gets the configuration property value. Configuration regarding properties of the custom extension that are can be overwritten for the onEmailOtpSendListener event listener.
// returns a CustomExtensionOverwriteConfigurationable when successful
func (m *OnOtpSendCustomExtensionHandler) GetConfiguration()(CustomExtensionOverwriteConfigurationable) {
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
// returns a OnOtpSendCustomExtensionable when successful
func (m *OnOtpSendCustomExtensionHandler) GetCustomExtension()(OnOtpSendCustomExtensionable) {
    val, err := m.GetBackingStore().Get("customExtension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OnOtpSendCustomExtensionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OnOtpSendCustomExtensionHandler) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnOtpSendHandler.GetFieldDeserializers()
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
        val, err := n.GetObjectValue(CreateOnOtpSendCustomExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtension(val.(OnOtpSendCustomExtensionable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *OnOtpSendCustomExtensionHandler) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnOtpSendHandler.Serialize(writer)
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
// SetConfiguration sets the configuration property value. Configuration regarding properties of the custom extension that are can be overwritten for the onEmailOtpSendListener event listener.
func (m *OnOtpSendCustomExtensionHandler) SetConfiguration(value CustomExtensionOverwriteConfigurationable)() {
    err := m.GetBackingStore().Set("configuration", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomExtension sets the customExtension property value. The customExtension property
func (m *OnOtpSendCustomExtensionHandler) SetCustomExtension(value OnOtpSendCustomExtensionable)() {
    err := m.GetBackingStore().Set("customExtension", value)
    if err != nil {
        panic(err)
    }
}
type OnOtpSendCustomExtensionHandlerable interface {
    OnOtpSendHandlerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfiguration()(CustomExtensionOverwriteConfigurationable)
    GetCustomExtension()(OnOtpSendCustomExtensionable)
    SetConfiguration(value CustomExtensionOverwriteConfigurationable)()
    SetCustomExtension(value OnOtpSendCustomExtensionable)()
}
