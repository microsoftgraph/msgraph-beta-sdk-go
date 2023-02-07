package informationprotection

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody 
type DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The evaluationInput property
    evaluationInput ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpEvaluationInputable
    // The notificationInfo property
    notificationInfo ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpNotificationable
    // The target property
    target *string
}
// NewDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody instantiates a new DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody and sets the default values.
func NewDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody()(*DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) {
    m := &DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEvaluationInput gets the evaluationInput property value. The evaluationInput property
func (m *DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) GetEvaluationInput()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpEvaluationInputable) {
    return m.evaluationInput
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["evaluationInput"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDlpEvaluationInputFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationInput(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpEvaluationInputable))
        }
        return nil
    }
    res["notificationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDlpNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationInfo(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpNotificationable))
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
// GetNotificationInfo gets the notificationInfo property value. The notificationInfo property
func (m *DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) GetNotificationInfo()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpNotificationable) {
    return m.notificationInfo
}
// GetTarget gets the target property value. The target property
func (m *DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) GetTarget()(*string) {
    return m.target
}
// Serialize serializes information the current object
func (m *DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("evaluationInput", m.GetEvaluationInput())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notificationInfo", m.GetNotificationInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target", m.GetTarget())
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
func (m *DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEvaluationInput sets the evaluationInput property value. The evaluationInput property
func (m *DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) SetEvaluationInput(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpEvaluationInputable)() {
    m.evaluationInput = value
}
// SetNotificationInfo sets the notificationInfo property value. The notificationInfo property
func (m *DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) SetNotificationInfo(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpNotificationable)() {
    m.notificationInfo = value
}
// SetTarget sets the target property value. The target property
func (m *DataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) SetTarget(value *string)() {
    m.target = value
}
