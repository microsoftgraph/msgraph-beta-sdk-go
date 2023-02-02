package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The logicalTest property
    logicalTest ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The valueIfFalse property
    valueIfFalse ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The valueIfTrue property
    valueIfTrue ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["logicalTest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogicalTest(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["valueIfFalse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueIfFalse(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["valueIfTrue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueIfTrue(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    return res
}
// GetLogicalTest gets the logicalTest property value. The logicalTest property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody) GetLogicalTest()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.logicalTest
}
// GetValueIfFalse gets the valueIfFalse property value. The valueIfFalse property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody) GetValueIfFalse()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.valueIfFalse
}
// GetValueIfTrue gets the valueIfTrue property value. The valueIfTrue property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody) GetValueIfTrue()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.valueIfTrue
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("logicalTest", m.GetLogicalTest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("valueIfFalse", m.GetValueIfFalse())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("valueIfTrue", m.GetValueIfTrue())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLogicalTest sets the logicalTest property value. The logicalTest property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody) SetLogicalTest(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.logicalTest = value
}
// SetValueIfFalse sets the valueIfFalse property value. The valueIfFalse property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody) SetValueIfFalse(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.valueIfFalse = value
}
// SetValueIfTrue sets the valueIfTrue property value. The valueIfTrue property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphIfIfPostRequestBody) SetValueIfTrue(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.valueIfTrue = value
}
