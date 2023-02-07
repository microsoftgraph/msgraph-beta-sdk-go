package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The basis property
    basis ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The maturity property
    maturity ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The pr property
    pr ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The redemption property
    redemption ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The settlement property
    settlement ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBasis gets the basis property value. The basis property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) GetBasis()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.basis
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["basis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBasis(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["maturity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaturity(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["pr"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPr(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["redemption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedemption(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["settlement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettlement(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    return res
}
// GetMaturity gets the maturity property value. The maturity property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) GetMaturity()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.maturity
}
// GetPr gets the pr property value. The pr property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) GetPr()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.pr
}
// GetRedemption gets the redemption property value. The redemption property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) GetRedemption()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.redemption
}
// GetSettlement gets the settlement property value. The settlement property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) GetSettlement()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.settlement
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("basis", m.GetBasis())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("maturity", m.GetMaturity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pr", m.GetPr())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("redemption", m.GetRedemption())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("settlement", m.GetSettlement())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBasis sets the basis property value. The basis property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) SetBasis(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.basis = value
}
// SetMaturity sets the maturity property value. The maturity property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) SetMaturity(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.maturity = value
}
// SetPr sets the pr property value. The pr property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) SetPr(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.pr = value
}
// SetRedemption sets the redemption property value. The redemption property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) SetRedemption(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.redemption = value
}
// SetSettlement sets the settlement property value. The settlement property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphDiscDiscPostRequestBody) SetSettlement(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.settlement = value
}
