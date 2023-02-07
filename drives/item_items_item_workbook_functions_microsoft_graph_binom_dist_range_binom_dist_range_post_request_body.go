package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The numberS property
    numberS ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The numberS2 property
    numberS2 ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The probabilityS property
    probabilityS ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
    // The trials property
    trials ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["numberS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberS(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["numberS2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberS2(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["probabilityS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProbabilityS(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    res["trials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrials(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable))
        }
        return nil
    }
    return res
}
// GetNumberS gets the numberS property value. The numberS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) GetNumberS()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.numberS
}
// GetNumberS2 gets the numberS2 property value. The numberS2 property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) GetNumberS2()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.numberS2
}
// GetProbabilityS gets the probabilityS property value. The probabilityS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) GetProbabilityS()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.probabilityS
}
// GetTrials gets the trials property value. The trials property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) GetTrials()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable) {
    return m.trials
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("numberS", m.GetNumberS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("numberS2", m.GetNumberS2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("probabilityS", m.GetProbabilityS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("trials", m.GetTrials())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNumberS sets the numberS property value. The numberS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) SetNumberS(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.numberS = value
}
// SetNumberS2 sets the numberS2 property value. The numberS2 property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) SetNumberS2(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.numberS2 = value
}
// SetProbabilityS sets the probabilityS property value. The probabilityS property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) SetProbabilityS(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.probabilityS = value
}
// SetTrials sets the trials property value. The trials property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphBinom_Dist_RangeBinom_Dist_RangePostRequestBody) SetTrials(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Jsonable)() {
    m.trials = value
}
