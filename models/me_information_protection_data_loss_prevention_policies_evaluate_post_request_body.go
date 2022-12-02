package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody provides operations to call the evaluate method.
type MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The evaluationInput property
    evaluationInput DlpEvaluationInputable
    // The notificationInfo property
    notificationInfo DlpNotificationable
    // The target property
    target *string
}
// NewMeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody instantiates a new MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody and sets the default values.
func NewMeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody()(*MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) {
    m := &MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEvaluationInput gets the evaluationInput property value. The evaluationInput property
func (m *MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetEvaluationInput()(DlpEvaluationInputable) {
    return m.evaluationInput
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["evaluationInput"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDlpEvaluationInputFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationInput(val.(DlpEvaluationInputable))
        }
        return nil
    }
    res["notificationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDlpNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationInfo(val.(DlpNotificationable))
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
func (m *MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetNotificationInfo()(DlpNotificationable) {
    return m.notificationInfo
}
// GetTarget gets the target property value. The target property
func (m *MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetTarget()(*string) {
    return m.target
}
// Serialize serializes information the current object
func (m *MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEvaluationInput sets the evaluationInput property value. The evaluationInput property
func (m *MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetEvaluationInput(value DlpEvaluationInputable)() {
    m.evaluationInput = value
}
// SetNotificationInfo sets the notificationInfo property value. The notificationInfo property
func (m *MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetNotificationInfo(value DlpNotificationable)() {
    m.notificationInfo = value
}
// SetTarget sets the target property value. The target property
func (m *MeInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetTarget(value *string)() {
    m.target = value
}
