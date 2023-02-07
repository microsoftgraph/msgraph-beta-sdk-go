package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The endPeriod property
    endPeriod ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The nper property
    nper ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The pv property
    pv ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The rate property
    rate ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The startPeriod property
    startPeriod ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The type property
    typeEscaped ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEndPeriod gets the endPeriod property value. The endPeriod property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) GetEndPeriod()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.endPeriod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndPeriod(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["nper"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNper(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
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
    res["startPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartPeriod(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
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
// GetNper gets the nper property value. The nper property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) GetNper()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.nper
}
// GetPv gets the pv property value. The pv property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) GetPv()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.pv
}
// GetRate gets the rate property value. The rate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) GetRate()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.rate
}
// GetStartPeriod gets the startPeriod property value. The startPeriod property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) GetStartPeriod()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.startPeriod
}
// GetType gets the type property value. The type property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) GetType()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("endPeriod", m.GetEndPeriod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("nper", m.GetNper())
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
        err := writer.WriteObjectValue("startPeriod", m.GetStartPeriod())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEndPeriod sets the endPeriod property value. The endPeriod property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) SetEndPeriod(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.endPeriod = value
}
// SetNper sets the nper property value. The nper property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) SetNper(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.nper = value
}
// SetPv sets the pv property value. The pv property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) SetPv(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.pv = value
}
// SetRate sets the rate property value. The rate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) SetRate(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.rate = value
}
// SetStartPeriod sets the startPeriod property value. The startPeriod property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) SetStartPeriod(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.startPeriod = value
}
// SetType sets the type property value. The type property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCumPrincCumPrincPostRequestBody) SetType(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.typeEscaped = value
}
