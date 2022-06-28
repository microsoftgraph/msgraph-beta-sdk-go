package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacAppIdentifier 
type MacAppIdentifier struct {
    MobileAppIdentifier
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identifier for an app, as specified in the app store.
    bundleId *string
}
// NewMacAppIdentifier instantiates a new MacAppIdentifier and sets the default values.
func NewMacAppIdentifier()(*MacAppIdentifier) {
    m := &MacAppIdentifier{
        MobileAppIdentifier: *NewMobileAppIdentifier(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacAppIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacAppIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacAppIdentifier(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacAppIdentifier) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBundleId gets the bundleId property value. The identifier for an app, as specified in the app store.
func (m *MacAppIdentifier) GetBundleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bundleId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacAppIdentifier) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppIdentifier.GetFieldDeserializers()
    res["bundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MacAppIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppIdentifier.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("bundleId", m.GetBundleId())
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
func (m *MacAppIdentifier) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBundleId sets the bundleId property value. The identifier for an app, as specified in the app store.
func (m *MacAppIdentifier) SetBundleId(value *string)() {
    if m != nil {
        m.bundleId = value
    }
}
