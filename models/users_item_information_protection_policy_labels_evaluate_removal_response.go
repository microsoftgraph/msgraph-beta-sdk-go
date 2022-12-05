package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponse provides operations to call the evaluateRemoval method.
type UsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []InformationProtectionActionable
}
// NewUsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponse instantiates a new UsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponse and sets the default values.
func NewUsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponse()(*UsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponse) {
    m := &UsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateUsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInformationProtectionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InformationProtectionActionable, len(val))
            for i, v := range val {
                res[i] = v.(InformationProtectionActionable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *UsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponse) GetValue()([]InformationProtectionActionable) {
    return m.value
}
// Serialize serializes information the current object
func (m *UsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponse) SetValue(value []InformationProtectionActionable)() {
    m.value = value
}
