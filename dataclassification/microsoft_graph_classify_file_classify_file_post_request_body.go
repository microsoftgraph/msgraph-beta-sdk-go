package dataclassification

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftGraphClassifyFileClassifyFilePostRequestBody 
type MicrosoftGraphClassifyFileClassifyFilePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The file property
    file []byte
    // The sensitiveTypeIds property
    sensitiveTypeIds []string
}
// NewMicrosoftGraphClassifyFileClassifyFilePostRequestBody instantiates a new MicrosoftGraphClassifyFileClassifyFilePostRequestBody and sets the default values.
func NewMicrosoftGraphClassifyFileClassifyFilePostRequestBody()(*MicrosoftGraphClassifyFileClassifyFilePostRequestBody) {
    m := &MicrosoftGraphClassifyFileClassifyFilePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftGraphClassifyFileClassifyFilePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphClassifyFileClassifyFilePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphClassifyFileClassifyFilePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphClassifyFileClassifyFilePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphClassifyFileClassifyFilePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["file"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val)
        }
        return nil
    }
    res["sensitiveTypeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSensitiveTypeIds(res)
        }
        return nil
    }
    return res
}
// GetFile gets the file property value. The file property
func (m *MicrosoftGraphClassifyFileClassifyFilePostRequestBody) GetFile()([]byte) {
    return m.file
}
// GetSensitiveTypeIds gets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *MicrosoftGraphClassifyFileClassifyFilePostRequestBody) GetSensitiveTypeIds()([]string) {
    return m.sensitiveTypeIds
}
// Serialize serializes information the current object
func (m *MicrosoftGraphClassifyFileClassifyFilePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeIds() != nil {
        err := writer.WriteCollectionOfStringValues("sensitiveTypeIds", m.GetSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphClassifyFileClassifyFilePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFile sets the file property value. The file property
func (m *MicrosoftGraphClassifyFileClassifyFilePostRequestBody) SetFile(value []byte)() {
    m.file = value
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *MicrosoftGraphClassifyFileClassifyFilePostRequestBody) SetSensitiveTypeIds(value []string)() {
    m.sensitiveTypeIds = value
}
