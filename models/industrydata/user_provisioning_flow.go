package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserProvisioningFlow struct {
    ProvisioningFlow
}
// NewUserProvisioningFlow instantiates a new UserProvisioningFlow and sets the default values.
func NewUserProvisioningFlow()(*UserProvisioningFlow) {
    m := &UserProvisioningFlow{
        ProvisioningFlow: *NewProvisioningFlow(),
    }
    odataTypeValue := "#microsoft.graph.industryData.userProvisioningFlow"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUserProvisioningFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserProvisioningFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserProvisioningFlow(), nil
}
// GetCreateUnmatchedUsers gets the createUnmatchedUsers property value. A boolean choice indicating whether unmatched users should be created or ignored.
// returns a *bool when successful
func (m *UserProvisioningFlow) GetCreateUnmatchedUsers()(*bool) {
    val, err := m.GetBackingStore().Get("createUnmatchedUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCreationOptions gets the creationOptions property value. The different management choices for the new users to be provisioned.
// returns a UserCreationOptionsable when successful
func (m *UserProvisioningFlow) GetCreationOptions()(UserCreationOptionsable) {
    val, err := m.GetBackingStore().Get("creationOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserCreationOptionsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserProvisioningFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningFlow.GetFieldDeserializers()
    res["createUnmatchedUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreateUnmatchedUsers(val)
        }
        return nil
    }
    res["creationOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserCreationOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationOptions(val.(UserCreationOptionsable))
        }
        return nil
    }
    res["managementOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserManagementOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementOptions(val.(UserManagementOptionsable))
        }
        return nil
    }
    return res
}
// GetManagementOptions gets the managementOptions property value. The managementOptions property
// returns a UserManagementOptionsable when successful
func (m *UserProvisioningFlow) GetManagementOptions()(UserManagementOptionsable) {
    val, err := m.GetBackingStore().Get("managementOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserManagementOptionsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserProvisioningFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningFlow.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("createUnmatchedUsers", m.GetCreateUnmatchedUsers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("creationOptions", m.GetCreationOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managementOptions", m.GetManagementOptions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreateUnmatchedUsers sets the createUnmatchedUsers property value. A boolean choice indicating whether unmatched users should be created or ignored.
func (m *UserProvisioningFlow) SetCreateUnmatchedUsers(value *bool)() {
    err := m.GetBackingStore().Set("createUnmatchedUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetCreationOptions sets the creationOptions property value. The different management choices for the new users to be provisioned.
func (m *UserProvisioningFlow) SetCreationOptions(value UserCreationOptionsable)() {
    err := m.GetBackingStore().Set("creationOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementOptions sets the managementOptions property value. The managementOptions property
func (m *UserProvisioningFlow) SetManagementOptions(value UserManagementOptionsable)() {
    err := m.GetBackingStore().Set("managementOptions", value)
    if err != nil {
        panic(err)
    }
}
type UserProvisioningFlowable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningFlowable
    GetCreateUnmatchedUsers()(*bool)
    GetCreationOptions()(UserCreationOptionsable)
    GetManagementOptions()(UserManagementOptionsable)
    SetCreateUnmatchedUsers(value *bool)()
    SetCreationOptions(value UserCreationOptionsable)()
    SetManagementOptions(value UserManagementOptionsable)()
}
