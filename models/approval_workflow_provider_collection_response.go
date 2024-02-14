package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApprovalWorkflowProviderCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewApprovalWorkflowProviderCollectionResponse instantiates a new ApprovalWorkflowProviderCollectionResponse and sets the default values.
func NewApprovalWorkflowProviderCollectionResponse()(*ApprovalWorkflowProviderCollectionResponse) {
    m := &ApprovalWorkflowProviderCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateApprovalWorkflowProviderCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApprovalWorkflowProviderCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalWorkflowProviderCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApprovalWorkflowProviderCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalWorkflowProviderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApprovalWorkflowProviderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ApprovalWorkflowProviderable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []ApprovalWorkflowProviderable when successful
func (m *ApprovalWorkflowProviderCollectionResponse) GetValue()([]ApprovalWorkflowProviderable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ApprovalWorkflowProviderable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApprovalWorkflowProviderCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ApprovalWorkflowProviderCollectionResponse) SetValue(value []ApprovalWorkflowProviderable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type ApprovalWorkflowProviderCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ApprovalWorkflowProviderable)
    SetValue(value []ApprovalWorkflowProviderable)()
}
