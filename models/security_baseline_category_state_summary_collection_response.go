package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineCategoryStateSummaryCollectionResponse 
type SecurityBaselineCategoryStateSummaryCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewSecurityBaselineCategoryStateSummaryCollectionResponse instantiates a new SecurityBaselineCategoryStateSummaryCollectionResponse and sets the default values.
func NewSecurityBaselineCategoryStateSummaryCollectionResponse()(*SecurityBaselineCategoryStateSummaryCollectionResponse) {
    m := &SecurityBaselineCategoryStateSummaryCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSecurityBaselineCategoryStateSummaryCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityBaselineCategoryStateSummaryCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityBaselineCategoryStateSummaryCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityBaselineCategoryStateSummaryCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityBaselineCategoryStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityBaselineCategoryStateSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(SecurityBaselineCategoryStateSummaryable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *SecurityBaselineCategoryStateSummaryCollectionResponse) GetValue()([]SecurityBaselineCategoryStateSummaryable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecurityBaselineCategoryStateSummaryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SecurityBaselineCategoryStateSummaryCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SecurityBaselineCategoryStateSummaryCollectionResponse) SetValue(value []SecurityBaselineCategoryStateSummaryable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// SecurityBaselineCategoryStateSummaryCollectionResponseable 
type SecurityBaselineCategoryStateSummaryCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]SecurityBaselineCategoryStateSummaryable)
    SetValue(value []SecurityBaselineCategoryStateSummaryable)()
}
