package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody provides operations to call the deleteTiIndicatorsByExternalId method.
type TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value property
    value []string
}
// NewTiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody instantiates a new TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody and sets the default values.
func NewTiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody()(*TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody) {
    m := &TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody) GetValue()([]string) {
    return m.value
}
// Serialize serializes information the current object
func (m *TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetValue() != nil {
        err := writer.WriteCollectionOfStringValues("value", m.GetValue())
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
func (m *TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *TiIndicatorsDeleteTiIndicatorsByExternalIdPostRequestBody) SetValue(value []string)() {
    m.value = value
}
