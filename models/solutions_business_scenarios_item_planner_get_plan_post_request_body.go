package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody provides operations to call the getPlan method.
type SolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The target property
    target BusinessScenarioTaskTargetBaseable
}
// NewSolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody instantiates a new SolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody and sets the default values.
func NewSolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody()(*SolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody) {
    m := &SolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSolutionsBusinessScenariosItemPlannerGetPlanPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSolutionsBusinessScenariosItemPlannerGetPlanPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBusinessScenarioTaskTargetBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(BusinessScenarioTaskTargetBaseable))
        }
        return nil
    }
    return res
}
// GetTarget gets the target property value. The target property
func (m *SolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody) GetTarget()(BusinessScenarioTaskTargetBaseable) {
    return m.target
}
// Serialize serializes information the current object
func (m *SolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetTarget sets the target property value. The target property
func (m *SolutionsBusinessScenariosItemPlannerGetPlanPostRequestBody) SetTarget(value BusinessScenarioTaskTargetBaseable)() {
    m.target = value
}
