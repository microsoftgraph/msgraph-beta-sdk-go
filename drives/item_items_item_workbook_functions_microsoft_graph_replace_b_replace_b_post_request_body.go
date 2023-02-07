package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The newText property
    newText ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The numBytes property
    numBytes ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The oldText property
    oldText ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The startNum property
    startNum ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewText(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["numBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumBytes(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["oldText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldText(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["startNum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartNum(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    return res
}
// GetNewText gets the newText property value. The newText property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) GetNewText()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.newText
}
// GetNumBytes gets the numBytes property value. The numBytes property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) GetNumBytes()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.numBytes
}
// GetOldText gets the oldText property value. The oldText property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) GetOldText()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.oldText
}
// GetStartNum gets the startNum property value. The startNum property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) GetStartNum()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.startNum
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("newText", m.GetNewText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("numBytes", m.GetNumBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("oldText", m.GetOldText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("startNum", m.GetStartNum())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewText sets the newText property value. The newText property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) SetNewText(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.newText = value
}
// SetNumBytes sets the numBytes property value. The numBytes property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) SetNumBytes(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.numBytes = value
}
// SetOldText sets the oldText property value. The oldText property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) SetOldText(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.oldText = value
}
// SetStartNum sets the startNum property value. The startNum property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceBReplaceBPostRequestBody) SetStartNum(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.startNum = value
}
