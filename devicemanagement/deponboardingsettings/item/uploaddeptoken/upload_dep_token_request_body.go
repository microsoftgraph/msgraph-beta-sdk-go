package uploaddeptoken

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UploadDepTokenRequestBody provides operations to call the uploadDepToken method.
type UploadDepTokenRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The appleId property
    appleId *string;
    // The depToken property
    depToken *string;
}
// NewUploadDepTokenRequestBody instantiates a new uploadDepTokenRequestBody and sets the default values.
func NewUploadDepTokenRequestBody()(*UploadDepTokenRequestBody) {
    m := &UploadDepTokenRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUploadDepTokenRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUploadDepTokenRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUploadDepTokenRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UploadDepTokenRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppleId gets the appleId property value. The appleId property
func (m *UploadDepTokenRequestBody) GetAppleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appleId
    }
}
// GetDepToken gets the depToken property value. The depToken property
func (m *UploadDepTokenRequestBody) GetDepToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.depToken
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UploadDepTokenRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appleId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleId(val)
        }
        return nil
    }
    res["depToken"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepToken(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UploadDepTokenRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appleId", m.GetAppleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("depToken", m.GetDepToken())
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
func (m *UploadDepTokenRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppleId sets the appleId property value. The appleId property
func (m *UploadDepTokenRequestBody) SetAppleId(value *string)() {
    if m != nil {
        m.appleId = value
    }
}
// SetDepToken sets the depToken property value. The depToken property
func (m *UploadDepTokenRequestBody) SetDepToken(value *string)() {
    if m != nil {
        m.depToken = value
    }
}
