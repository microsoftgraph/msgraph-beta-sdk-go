package compliance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody 
type EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ids property
    ids []string
}
// NewEdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody instantiates a new EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody and sets the default values.
func NewEdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody()(*EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) {
    m := &EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIds(res)
        }
        return nil
    }
    return res
}
// GetIds gets the ids property value. The ids property
func (m *EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) GetIds()([]string) {
    return m.ids
}
// Serialize serializes information the current object
func (m *EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIds() != nil {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
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
func (m *EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIds sets the ids property value. The ids property
func (m *EdiscoveryCasesItemCustodiansMicrosoftGraphEdiscoveryApplyHoldApplyHoldPostRequestBody) SetIds(value []string)() {
    m.ids = value
}
