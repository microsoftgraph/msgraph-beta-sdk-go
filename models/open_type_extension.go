package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OpenTypeExtension 
type OpenTypeExtension struct {
    Extension
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A unique text identifier for an open type data extension. Required.
    extensionName *string
}
// NewOpenTypeExtension instantiates a new OpenTypeExtension and sets the default values.
func NewOpenTypeExtension()(*OpenTypeExtension) {
    m := &OpenTypeExtension{
        Extension: *NewExtension(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOpenTypeExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOpenTypeExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOpenTypeExtension(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OpenTypeExtension) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExtensionName gets the extensionName property value. A unique text identifier for an open type data extension. Required.
func (m *OpenTypeExtension) GetExtensionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OpenTypeExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Extension.GetFieldDeserializers()
    res["extensionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *OpenTypeExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Extension.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("extensionName", m.GetExtensionName())
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
func (m *OpenTypeExtension) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExtensionName sets the extensionName property value. A unique text identifier for an open type data extension. Required.
func (m *OpenTypeExtension) SetExtensionName(value *string)() {
    if m != nil {
        m.extensionName = value
    }
}
