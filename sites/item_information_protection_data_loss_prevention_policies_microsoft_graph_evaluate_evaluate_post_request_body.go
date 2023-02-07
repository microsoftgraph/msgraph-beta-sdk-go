package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody 
type ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The evaluationInput property
    evaluationInput ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpEvaluationInputable
    // The notificationInfo property
    notificationInfo ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpNotificationable
    // The target property
    target *string
}
// NewItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody instantiates a new ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody and sets the default values.
func NewItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody()(*ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) {
    m := &ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEvaluationInput gets the evaluationInput property value. The evaluationInput property
func (m *ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) GetEvaluationInput()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpEvaluationInputable) {
    return m.evaluationInput
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) GetNotificationInfo()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpNotificationable) {
    return m.notificationInfo
}
// GetTarget gets the target property value. The target property
func (m *ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) GetTarget()(*string) {
    return m.target
}
// Serialize serializes information the current object
func (m *ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEvaluationInput sets the evaluationInput property value. The evaluationInput property
func (m *ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) SetEvaluationInput(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpEvaluationInputable)() {
    m.evaluationInput = value
}
// SetNotificationInfo sets the notificationInfo property value. The notificationInfo property
func (m *ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) SetNotificationInfo(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpNotificationable)() {
    m.notificationInfo = value
}
// SetTarget sets the target property value. The target property
func (m *ItemInformationProtectionDataLossPreventionPoliciesMicrosoftGraphEvaluateEvaluatePostRequestBody) SetTarget(value *string)() {
    m.target = value
}
