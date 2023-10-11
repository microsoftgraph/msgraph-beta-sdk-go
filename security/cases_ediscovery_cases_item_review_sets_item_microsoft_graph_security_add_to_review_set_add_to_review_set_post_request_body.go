package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

// CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody 
type CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody instantiates a new CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody()(*CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody) {
    m := &CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetAdditionalDataOptions gets the additionalDataOptions property value. The additionalDataOptions property
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody) GetAdditionalDataOptions()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AdditionalDataOptions) {
    val, err := m.GetBackingStore().Get("additionalDataOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AdditionalDataOptions)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalDataOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParseAdditionalDataOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDataOptions(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AdditionalDataOptions))
        }
        return nil
    }
    res["search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoverySearchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoverySearchable))
        }
        return nil
    }
    return res
}
// GetSearch gets the search property value. The search property
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody) GetSearch()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoverySearchable) {
    val, err := m.GetBackingStore().Get("search")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoverySearchable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdditionalDataOptions() != nil {
        cast := (*m.GetAdditionalDataOptions()).String()
        err := writer.WriteStringValue("additionalDataOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("search", m.GetSearch())
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
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalDataOptions sets the additionalDataOptions property value. The additionalDataOptions property
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody) SetAdditionalDataOptions(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AdditionalDataOptions)() {
    err := m.GetBackingStore().Set("additionalDataOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetSearch sets the search property value. The search property
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBody) SetSearch(value i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoverySearchable)() {
    err := m.GetBackingStore().Set("search", value)
    if err != nil {
        panic(err)
    }
}
// CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBodyable 
type CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityAddToReviewSetAddToReviewSetPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalDataOptions()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AdditionalDataOptions)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetSearch()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoverySearchable)
    SetAdditionalDataOptions(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AdditionalDataOptions)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetSearch(value i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoverySearchable)()
}
