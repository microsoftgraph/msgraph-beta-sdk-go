package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dates property
    dates ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The guess property
    guess ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The values property
    values ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDates gets the dates property value. The dates property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody) GetDates()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.dates
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDates(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["guess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuess(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    return res
}
// GetGuess gets the guess property value. The guess property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody) GetGuess()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.guess
}
// GetValues gets the values property value. The values property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody) GetValues()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.values
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("dates", m.GetDates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("guess", m.GetGuess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("values", m.GetValues())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDates sets the dates property value. The dates property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody) SetDates(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.dates = value
}
// SetGuess sets the guess property value. The guess property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody) SetGuess(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.guess = value
}
// SetValues sets the values property value. The values property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphXirrXirrPostRequestBody) SetValues(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.values = value
}
