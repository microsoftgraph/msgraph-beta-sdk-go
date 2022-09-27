package encryptbuffer

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EncryptBufferPostRequestBody provides operations to call the encryptBuffer method.
type EncryptBufferPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The buffer property
    buffer []byte
    // The labelId property
    labelId *string
}
// NewEncryptBufferPostRequestBody instantiates a new encryptBufferPostRequestBody and sets the default values.
func NewEncryptBufferPostRequestBody()(*EncryptBufferPostRequestBody) {
    m := &EncryptBufferPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEncryptBufferPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEncryptBufferPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEncryptBufferPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EncryptBufferPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBuffer gets the buffer property value. The buffer property
func (m *EncryptBufferPostRequestBody) GetBuffer()([]byte) {
    return m.buffer
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EncryptBufferPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["buffer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetBuffer)
    res["labelId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLabelId)
    return res
}
// GetLabelId gets the labelId property value. The labelId property
func (m *EncryptBufferPostRequestBody) GetLabelId()(*string) {
    return m.labelId
}
// Serialize serializes information the current object
func (m *EncryptBufferPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("buffer", m.GetBuffer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("labelId", m.GetLabelId())
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
func (m *EncryptBufferPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBuffer sets the buffer property value. The buffer property
func (m *EncryptBufferPostRequestBody) SetBuffer(value []byte)() {
    m.buffer = value
}
// SetLabelId sets the labelId property value. The labelId property
func (m *EncryptBufferPostRequestBody) SetLabelId(value *string)() {
    m.labelId = value
}
