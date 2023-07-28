package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssociatedBranch 
type AssociatedBranch struct {
    Association
}
// NewAssociatedBranch instantiates a new associatedBranch and sets the default values.
func NewAssociatedBranch()(*AssociatedBranch) {
    m := &AssociatedBranch{
        Association: *NewAssociation(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.associatedBranch"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAssociatedBranchFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssociatedBranchFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssociatedBranch(), nil
}
// GetBranchId gets the branchId property value. Identifier for the branch.
func (m *AssociatedBranch) GetBranchId()(*string) {
    val, err := m.GetBackingStore().Get("branchId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssociatedBranch) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Association.GetFieldDeserializers()
    res["branchId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranchId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AssociatedBranch) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Association.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("branchId", m.GetBranchId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBranchId sets the branchId property value. Identifier for the branch.
func (m *AssociatedBranch) SetBranchId(value *string)() {
    err := m.GetBackingStore().Set("branchId", value)
    if err != nil {
        panic(err)
    }
}
// AssociatedBranchable 
type AssociatedBranchable interface {
    Associationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBranchId()(*string)
    SetBranchId(value *string)()
}
