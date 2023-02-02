package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

// CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody 
type CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The purgeAreas property
    purgeAreas *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeAreas
    // The purgeType property
    purgeType *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeType
}
// NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody instantiates a new CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody()(*CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) {
    m := &CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["purgeAreas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParsePurgeAreas)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPurgeAreas(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeAreas))
        }
        return nil
    }
    res["purgeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParsePurgeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPurgeType(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeType))
        }
        return nil
    }
    return res
}
// GetPurgeAreas gets the purgeAreas property value. The purgeAreas property
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) GetPurgeAreas()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeAreas) {
    return m.purgeAreas
}
// GetPurgeType gets the purgeType property value. The purgeType property
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) GetPurgeType()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeType) {
    return m.purgeType
}
// Serialize serializes information the current object
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPurgeAreas() != nil {
        cast := (*m.GetPurgeAreas()).String()
        err := writer.WriteStringValue("purgeAreas", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPurgeType() != nil {
        cast := (*m.GetPurgeType()).String()
        err := writer.WriteStringValue("purgeType", &cast)
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
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPurgeAreas sets the purgeAreas property value. The purgeAreas property
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) SetPurgeAreas(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeAreas)() {
    m.purgeAreas = value
}
// SetPurgeType sets the purgeType property value. The purgeType property
func (m *CasesEdiscoveryCasesItemSearchesItemMicrosoftGraphSecurityPurgeDataPurgeDataPostRequestBody) SetPurgeType(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PurgeType)() {
    m.purgeType = value
}
