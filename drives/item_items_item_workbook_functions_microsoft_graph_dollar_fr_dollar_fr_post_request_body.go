package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The decimalDollar property
    decimalDollar ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The fraction property
    fraction ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDecimalDollar gets the decimalDollar property value. The decimalDollar property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody) GetDecimalDollar()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.decimalDollar
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["decimalDollar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecimalDollar(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["fraction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFraction(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    return res
}
// GetFraction gets the fraction property value. The fraction property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody) GetFraction()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.fraction
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("decimalDollar", m.GetDecimalDollar())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("fraction", m.GetFraction())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDecimalDollar sets the decimalDollar property value. The decimalDollar property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody) SetDecimalDollar(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.decimalDollar = value
}
// SetFraction sets the fraction property value. The fraction property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDollarFrDollarFrPostRequestBody) SetFraction(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.fraction = value
}
