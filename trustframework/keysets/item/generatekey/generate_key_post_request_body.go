package generatekey

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GenerateKeyPostRequestBody provides operations to call the generateKey method.
type GenerateKeyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The exp property
    exp *int64
    // The kty property
    kty *string
    // The nbf property
    nbf *int64
    // The use property
    use *string
}
// NewGenerateKeyPostRequestBody instantiates a new generateKeyPostRequestBody and sets the default values.
func NewGenerateKeyPostRequestBody()(*GenerateKeyPostRequestBody) {
    m := &GenerateKeyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGenerateKeyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGenerateKeyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGenerateKeyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GenerateKeyPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExp gets the exp property value. The exp property
func (m *GenerateKeyPostRequestBody) GetExp()(*int64) {
    return m.exp
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GenerateKeyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["exp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExp(val)
        }
        return nil
    }
    res["kty"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKty(val)
        }
        return nil
    }
    res["nbf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNbf(val)
        }
        return nil
    }
    res["use"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUse(val)
        }
        return nil
    }
    return res
}
// GetKty gets the kty property value. The kty property
func (m *GenerateKeyPostRequestBody) GetKty()(*string) {
    return m.kty
}
// GetNbf gets the nbf property value. The nbf property
func (m *GenerateKeyPostRequestBody) GetNbf()(*int64) {
    return m.nbf
}
// GetUse gets the use property value. The use property
func (m *GenerateKeyPostRequestBody) GetUse()(*string) {
    return m.use
}
// Serialize serializes information the current object
func (m *GenerateKeyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("exp", m.GetExp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("kty", m.GetKty())
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
func (m *GenerateKeyPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExp sets the exp property value. The exp property
func (m *GenerateKeyPostRequestBody) SetExp(value *int64)() {
    m.exp = value
}
// SetKty sets the kty property value. The kty property
func (m *GenerateKeyPostRequestBody) SetKty(value *string)() {
    m.kty = value
}
// SetNbf sets the nbf property value. The nbf property
func (m *GenerateKeyPostRequestBody) SetNbf(value *int64)() {
    m.nbf = value
}
// SetUse sets the use property value. The use property
func (m *GenerateKeyPostRequestBody) SetUse(value *string)() {
    m.use = value
}
