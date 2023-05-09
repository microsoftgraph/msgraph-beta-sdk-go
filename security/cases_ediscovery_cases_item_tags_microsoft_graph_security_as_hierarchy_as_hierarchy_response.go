package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

// CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponse 
type CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewCasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponse instantiates a new CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponse and sets the default values.
func NewCasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponse()(*CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponse) {
    m := &CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryReviewTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewTagable, len(val))
            for i, v := range val {
                res[i] = v.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewTagable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponse) GetValue()([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewTagable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewTagable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponse) SetValue(value []i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewTagable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponseable 
type CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewTagable)
    SetValue(value []i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewTagable)()
}
