package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody provides operations to call the evaluateClassificationResults method.
type InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The classificationResults property
    classificationResults []ClassificationResultable
    // The contentInfo property
    contentInfo ContentInfoable
}
// NewInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody instantiates a new InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody and sets the default values.
func NewInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody()(*InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) {
    m := &InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClassificationResults gets the classificationResults property value. The classificationResults property
func (m *InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetClassificationResults()([]ClassificationResultable) {
    return m.classificationResults
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetContentInfo()(ContentInfoable) {
    return m.contentInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["classificationResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClassificationResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClassificationResultable, len(val))
            for i, v := range val {
                res[i] = v.(ClassificationResultable)
            }
            m.SetClassificationResults(res)
        }
        return nil
    }
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
    return res
}
// Serialize serializes information the current object
func (m *InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClassificationResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClassificationResults()))
        for i, v := range m.GetClassificationResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("classificationResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
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
func (m *InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClassificationResults sets the classificationResults property value. The classificationResults property
func (m *InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) SetClassificationResults(value []ClassificationResultable)() {
    m.classificationResults = value
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *InformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) SetContentInfo(value ContentInfoable)() {
    m.contentInfo = value
}
