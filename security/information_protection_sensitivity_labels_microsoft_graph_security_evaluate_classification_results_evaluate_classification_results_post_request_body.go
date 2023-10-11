package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

// InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody 
type InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody instantiates a new InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody and sets the default values.
func NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody()(*InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) {
    m := &InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetClassificationResults gets the classificationResults property value. The classificationResults property
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) GetClassificationResults()([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ClassificationResultable) {
    val, err := m.GetBackingStore().Get("classificationResults")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ClassificationResultable)
    }
    return nil
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) GetContentInfo()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable) {
    val, err := m.GetBackingStore().Get("contentInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["classificationResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateClassificationResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ClassificationResultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ClassificationResultable)
                }
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
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClassificationResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClassificationResults()))
        for i, v := range m.GetClassificationResults() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetClassificationResults sets the classificationResults property value. The classificationResults property
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) SetClassificationResults(value []i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ClassificationResultable)() {
    err := m.GetBackingStore().Set("classificationResults", value)
    if err != nil {
        panic(err)
    }
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBody) SetContentInfo(value i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable)() {
    err := m.GetBackingStore().Set("contentInfo", value)
    if err != nil {
        panic(err)
    }
}
// InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBodyable 
type InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetClassificationResults()([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ClassificationResultable)
    GetContentInfo()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetClassificationResults(value []i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ClassificationResultable)()
    SetContentInfo(value i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable)()
}
