package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The degFreedom property
    degFreedom ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The probability property
    probability ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDegFreedom gets the degFreedom property value. The degFreedom property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody) GetDegFreedom()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.degFreedom
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["degFreedom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDegFreedom(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["probability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProbability(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    return res
}
// GetProbability gets the probability property value. The probability property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody) GetProbability()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.probability
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("degFreedom", m.GetDegFreedom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("probability", m.GetProbability())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDegFreedom sets the degFreedom property value. The degFreedom property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody) SetDegFreedom(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.degFreedom = value
}
// SetProbability sets the probability property value. The probability property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphChiSq_Inv_RTChiSq_Inv_RTPostRequestBody) SetProbability(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.probability = value
}
