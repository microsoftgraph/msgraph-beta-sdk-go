package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAppIdentifier 
type WindowsAppIdentifier struct {
    MobileAppIdentifier
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identifier for an app, as specified in the app store.
    windowsAppId *string
}
// NewWindowsAppIdentifier instantiates a new WindowsAppIdentifier and sets the default values.
func NewWindowsAppIdentifier()(*WindowsAppIdentifier) {
    m := &WindowsAppIdentifier{
        MobileAppIdentifier: *NewMobileAppIdentifier(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsAppIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAppIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAppIdentifier(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsAppIdentifier) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAppIdentifier) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppIdentifier.GetFieldDeserializers()
    res["windowsAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsAppId(val)
        }
        return nil
    }
    return res
}
// GetWindowsAppId gets the windowsAppId property value. The identifier for an app, as specified in the app store.
func (m *WindowsAppIdentifier) GetWindowsAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.windowsAppId
    }
}
// Serialize serializes information the current object
func (m *WindowsAppIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppIdentifier.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("windowsAppId", m.GetWindowsAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsAppIdentifier) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetWindowsAppId sets the windowsAppId property value. The identifier for an app, as specified in the app store.
func (m *WindowsAppIdentifier) SetWindowsAppId(value *string)() {
    if m != nil {
        m.windowsAppId = value
    }
}
