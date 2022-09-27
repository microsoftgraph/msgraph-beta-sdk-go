package uploadsecret

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UploadSecretPostRequestBody provides operations to call the uploadSecret method.
type UploadSecretPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The exp property
    exp *int64
    // The k property
    k *string
    // The nbf property
    nbf *int64
    // The use property
    use *string
}
// NewUploadSecretPostRequestBody instantiates a new uploadSecretPostRequestBody and sets the default values.
func NewUploadSecretPostRequestBody()(*UploadSecretPostRequestBody) {
    m := &UploadSecretPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUploadSecretPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUploadSecretPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUploadSecretPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UploadSecretPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExp gets the exp property value. The exp property
func (m *UploadSecretPostRequestBody) GetExp()(*int64) {
    return m.exp
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UploadSecretPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["exp"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetExp)
    res["k"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetK)
    res["nbf"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetNbf)
    res["use"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUse)
    return res
}
// GetK gets the k property value. The k property
func (m *UploadSecretPostRequestBody) GetK()(*string) {
    return m.k
}
// GetNbf gets the nbf property value. The nbf property
func (m *UploadSecretPostRequestBody) GetNbf()(*int64) {
    return m.nbf
}
// GetUse gets the use property value. The use property
func (m *UploadSecretPostRequestBody) GetUse()(*string) {
    return m.use
}
// Serialize serializes information the current object
func (m *UploadSecretPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("exp", m.GetExp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("k", m.GetK())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("nbf", m.GetNbf())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("use", m.GetUse())
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
func (m *UploadSecretPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExp sets the exp property value. The exp property
func (m *UploadSecretPostRequestBody) SetExp(value *int64)() {
    m.exp = value
}
// SetK sets the k property value. The k property
func (m *UploadSecretPostRequestBody) SetK(value *string)() {
    m.k = value
}
// SetNbf sets the nbf property value. The nbf property
func (m *UploadSecretPostRequestBody) SetNbf(value *int64)() {
    m.nbf = value
}
// SetUse sets the use property value. The use property
func (m *UploadSecretPostRequestBody) SetUse(value *string)() {
    m.use = value
}
