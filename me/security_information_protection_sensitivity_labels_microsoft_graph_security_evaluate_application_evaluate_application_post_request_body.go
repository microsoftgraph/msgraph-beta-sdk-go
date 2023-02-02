package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

// SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody 
type SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The contentInfo property
    contentInfo i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable
    // The labelingOptions property
    labelingOptions i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.LabelingOptionsable
}
// NewSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody instantiates a new SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody and sets the default values.
func NewSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody()(*SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody) {
    m := &SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody) GetContentInfo()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable) {
    return m.contentInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["labelingOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateLabelingOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelingOptions(val.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.LabelingOptionsable))
        }
        return nil
    }
    return res
}
// GetLabelingOptions gets the labelingOptions property value. The labelingOptions property
func (m *SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody) GetLabelingOptions()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.LabelingOptionsable) {
    return m.labelingOptions
}
// Serialize serializes information the current object
func (m *SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody) SetContentInfo(value i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable)() {
    m.contentInfo = value
}
// SetLabelingOptions sets the labelingOptions property value. The labelingOptions property
func (m *SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBody) SetLabelingOptions(value i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.LabelingOptionsable)() {
    m.labelingOptions = value
}
