package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse instantiates a new ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse and sets the default values.
func NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse()(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse) {
    m := &ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateInformationProtectionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionActionable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []InformationProtectionActionable when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse) GetValue()([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionActionable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionActionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponse) SetValue(value []i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionActionable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionActionable)
    SetValue(value []i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionActionable)()
}
