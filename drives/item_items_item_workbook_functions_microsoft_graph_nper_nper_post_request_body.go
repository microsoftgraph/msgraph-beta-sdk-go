package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The fv property
    fv ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The pmt property
    pmt ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The pv property
    pv ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The rate property
    rate ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The type property
    typeEscaped ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fv"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFv(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["pmt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPmt(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["pv"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPv(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["rate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRate(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    return res
}
// GetFv gets the fv property value. The fv property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) GetFv()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.fv
}
// GetPmt gets the pmt property value. The pmt property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) GetPmt()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.pmt
}
// GetPv gets the pv property value. The pv property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) GetPv()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.pv
}
// GetRate gets the rate property value. The rate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) GetRate()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.rate
}
// GetType gets the type property value. The type property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) GetType()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("fv", m.GetFv())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pmt", m.GetPmt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pv", m.GetPv())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rate", m.GetRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("type", m.GetType())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFv sets the fv property value. The fv property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) SetFv(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.fv = value
}
// SetPmt sets the pmt property value. The pmt property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) SetPmt(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.pmt = value
}
// SetPv sets the pv property value. The pv property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) SetPv(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.pv = value
}
// SetRate sets the rate property value. The rate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) SetRate(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.rate = value
}
// SetType sets the type property value. The type property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNperNperPostRequestBody) SetType(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.typeEscaped = value
}
