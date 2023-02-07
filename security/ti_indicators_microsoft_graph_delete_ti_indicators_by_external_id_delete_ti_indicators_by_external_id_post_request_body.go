package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody 
type TiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The value property
    value []string
}
// NewTiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody instantiates a new TiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody and sets the default values.
func NewTiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody()(*TiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody) {
    m := &TiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *TiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody) GetValue()([]string) {
    return m.value
}
// Serialize serializes information the current object
func (m *TiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *TiIndicatorsMicrosoftGraphDeleteTiIndicatorsByExternalIdDeleteTiIndicatorsByExternalIdPostRequestBody) SetValue(value []string)() {
    m.value = value
}
