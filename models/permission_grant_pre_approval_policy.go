package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PermissionGrantPreApprovalPolicy struct {
    DirectoryObject
}
// NewPermissionGrantPreApprovalPolicy instantiates a new PermissionGrantPreApprovalPolicy and sets the default values.
func NewPermissionGrantPreApprovalPolicy()(*PermissionGrantPreApprovalPolicy) {
    m := &PermissionGrantPreApprovalPolicy{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.permissionGrantPreApprovalPolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePermissionGrantPreApprovalPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePermissionGrantPreApprovalPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPermissionGrantPreApprovalPolicy(), nil
}
// GetConditions gets the conditions property value. A list of condition sets describing the conditions under which the permission to grant consent for the app has been preapproved.
// returns a []PreApprovalDetailable when successful
func (m *PermissionGrantPreApprovalPolicy) GetConditions()([]PreApprovalDetailable) {
    val, err := m.GetBackingStore().Get("conditions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PreApprovalDetailable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PermissionGrantPreApprovalPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["conditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePreApprovalDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PreApprovalDetailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PreApprovalDetailable)
                }
            }
            m.SetConditions(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PermissionGrantPreApprovalPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConditions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConditions()))
        for i, v := range m.GetConditions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("conditions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConditions sets the conditions property value. A list of condition sets describing the conditions under which the permission to grant consent for the app has been preapproved.
func (m *PermissionGrantPreApprovalPolicy) SetConditions(value []PreApprovalDetailable)() {
    err := m.GetBackingStore().Set("conditions", value)
    if err != nil {
        panic(err)
    }
}
type PermissionGrantPreApprovalPolicyable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConditions()([]PreApprovalDetailable)
    SetConditions(value []PreApprovalDetailable)()
}
