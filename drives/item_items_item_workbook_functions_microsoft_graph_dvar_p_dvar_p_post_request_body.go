package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The criteria property
    criteria ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The database property
    database ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The field property
    field ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCriteria gets the criteria property value. The criteria property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody) GetCriteria()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.criteria
}
// GetDatabase gets the database property value. The database property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody) GetDatabase()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.database
}
// GetField gets the field property value. The field property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody) GetField()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.field
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["criteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["database"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatabase(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["field"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetField(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("criteria", m.GetCriteria())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("database", m.GetDatabase())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("field", m.GetField())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCriteria sets the criteria property value. The criteria property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody) SetCriteria(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.criteria = value
}
// SetDatabase sets the database property value. The database property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody) SetDatabase(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.database = value
}
// SetField sets the field property value. The field property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDvarPDvarPPostRequestBody) SetField(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.field = value
}
