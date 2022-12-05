package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody provides operations to call the evaluateApplication method.
type MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contentInfo property
    contentInfo ContentInfoable
    // The labelingOptions property
    labelingOptions LabelingOptionsable
}
// NewMeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody instantiates a new MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody and sets the default values.
func NewMeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody()(*MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) {
    m := &MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) GetContentInfo()(ContentInfoable) {
    return m.contentInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(ContentInfoable))
        }
        return nil
    }
    res["labelingOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLabelingOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelingOptions(val.(LabelingOptionsable))
        }
        return nil
    }
    return res
}
// GetLabelingOptions gets the labelingOptions property value. The labelingOptions property
func (m *MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) GetLabelingOptions()(LabelingOptionsable) {
    return m.labelingOptions
}
// Serialize serializes information the current object
func (m *MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("labelingOptions", m.GetLabelingOptions())
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
func (m *MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) SetContentInfo(value ContentInfoable)() {
    m.contentInfo = value
}
// SetLabelingOptions sets the labelingOptions property value. The labelingOptions property
func (m *MeInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) SetLabelingOptions(value LabelingOptionsable)() {
    m.labelingOptions = value
}
