package informationprotection

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody 
type MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The buffer property
    buffer []byte
    // The labelId property
    labelId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}
// NewMicrosoftGraphEncryptBufferEncryptBufferPostRequestBody instantiates a new MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody and sets the default values.
func NewMicrosoftGraphEncryptBufferEncryptBufferPostRequestBody()(*MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody) {
    m := &MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftGraphEncryptBufferEncryptBufferPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphEncryptBufferEncryptBufferPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphEncryptBufferEncryptBufferPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuffer gets the buffer property value. The buffer property
func (m *MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody) GetBuffer()([]byte) {
    return m.buffer
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["buffer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuffer(val)
        }
        return nil
    }
    res["labelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelId(val)
        }
        return nil
    }
    return res
}
// GetLabelId gets the labelId property value. The labelId property
func (m *MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody) GetLabelId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.labelId
}
// Serialize serializes information the current object
func (m *MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("buffer", m.GetBuffer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("labelId", m.GetLabelId())
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
func (m *MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuffer sets the buffer property value. The buffer property
func (m *MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody) SetBuffer(value []byte)() {
    m.buffer = value
}
// SetLabelId sets the labelId property value. The labelId property
func (m *MicrosoftGraphEncryptBufferEncryptBufferPostRequestBody) SetLabelId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.labelId = value
}
