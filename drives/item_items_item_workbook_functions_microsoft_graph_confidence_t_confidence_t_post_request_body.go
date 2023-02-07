package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The alpha property
    alpha ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The size property
    size ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The standardDev property
    standardDev ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlpha gets the alpha property value. The alpha property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody) GetAlpha()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.alpha
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alpha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlpha(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["standardDev"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStandardDev(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    return res
}
// GetSize gets the size property value. The size property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody) GetSize()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.size
}
// GetStandardDev gets the standardDev property value. The standardDev property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody) GetStandardDev()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.standardDev
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("alpha", m.GetAlpha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("standardDev", m.GetStandardDev())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlpha sets the alpha property value. The alpha property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody) SetAlpha(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.alpha = value
}
// SetSize sets the size property value. The size property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody) SetSize(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.size = value
}
// SetStandardDev sets the standardDev property value. The standardDev property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphConfidence_TConfidence_TPostRequestBody) SetStandardDev(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.standardDev = value
}
