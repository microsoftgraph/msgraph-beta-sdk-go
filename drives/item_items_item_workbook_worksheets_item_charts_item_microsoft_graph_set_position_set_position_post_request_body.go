package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody 
type ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The endCell property
    endCell ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The startCell property
    startCell ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody()(*ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEndCell gets the endCell property value. The endCell property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody) GetEndCell()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.endCell
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endCell"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndCell(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["startCell"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartCell(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    return res
}
// GetStartCell gets the startCell property value. The startCell property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody) GetStartCell()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.startCell
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("endCell", m.GetEndCell())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("startCell", m.GetStartCell())
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
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEndCell sets the endCell property value. The endCell property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody) SetEndCell(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.endCell = value
}
// SetStartCell sets the startCell property value. The startCell property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemMicrosoftGraphSetPositionSetPositionPostRequestBody) SetStartCell(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.startCell = value
}
