package solutions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// BusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody 
type BusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The target property
    target ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskTargetBaseable
}
// NewBusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody instantiates a new BusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody and sets the default values.
func NewBusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody()(*BusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody) {
    m := &BusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessScenarioTaskTargetBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskTargetBaseable))
        }
        return nil
    }
    return res
}
// GetTarget gets the target property value. The target property
func (m *BusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody) GetTarget()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskTargetBaseable) {
    return m.target
}
// Serialize serializes information the current object
func (m *BusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("target", m.GetTarget())
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
func (m *BusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTarget sets the target property value. The target property
func (m *BusinessScenariosItemPlannerMicrosoftGraphGetPlanGetPlanPostRequestBody) SetTarget(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskTargetBaseable)() {
    m.target = value
}
