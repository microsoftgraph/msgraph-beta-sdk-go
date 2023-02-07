package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

// ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody 
type ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The classificationResults property
    classificationResults []i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ClassificationResultable
    // The contentInfo property
    contentInfo i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable
}
// NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody instantiates a new ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody()(*ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) {
    m := &ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClassificationResults gets the classificationResults property value. The classificationResults property
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) GetClassificationResults()([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ClassificationResultable) {
    return m.classificationResults
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) GetContentInfo()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable) {
    return m.contentInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["classificationResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateClassificationResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ClassificationResultable, len(val))
            for i, v := range val {
                res[i] = v.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ClassificationResultable)
            }
            m.SetClassificationResults(res)
        }
        return nil
    }
    res["contentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateContentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClassificationResults sets the classificationResults property value. The classificationResults property
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) SetClassificationResults(value []i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ClassificationResultable)() {
    m.classificationResults = value
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) SetContentInfo(value i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable)() {
    m.contentInfo = value
}
