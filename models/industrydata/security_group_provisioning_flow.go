package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecurityGroupProvisioningFlow struct {
    ProvisioningFlow
}
// NewSecurityGroupProvisioningFlow instantiates a new SecurityGroupProvisioningFlow and sets the default values.
func NewSecurityGroupProvisioningFlow()(*SecurityGroupProvisioningFlow) {
    m := &SecurityGroupProvisioningFlow{
        ProvisioningFlow: *NewProvisioningFlow(),
    }
    odataTypeValue := "#microsoft.graph.industryData.securityGroupProvisioningFlow"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSecurityGroupProvisioningFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityGroupProvisioningFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityGroupProvisioningFlow(), nil
}
// GetCreationOptions gets the creationOptions property value. The creationOptions property
// returns a SecurityGroupCreationOptionsable when successful
func (m *SecurityGroupProvisioningFlow) GetCreationOptions()(SecurityGroupCreationOptionsable) {
    val, err := m.GetBackingStore().Get("creationOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SecurityGroupCreationOptionsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecurityGroupProvisioningFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningFlow.GetFieldDeserializers()
    res["creationOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityGroupCreationOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationOptions(val.(SecurityGroupCreationOptionsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *SecurityGroupProvisioningFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningFlow.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("creationOptions", m.GetCreationOptions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreationOptions sets the creationOptions property value. The creationOptions property
func (m *SecurityGroupProvisioningFlow) SetCreationOptions(value SecurityGroupCreationOptionsable)() {
    err := m.GetBackingStore().Set("creationOptions", value)
    if err != nil {
        panic(err)
    }
}
type SecurityGroupProvisioningFlowable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningFlowable
    GetCreationOptions()(SecurityGroupCreationOptionsable)
    SetCreationOptions(value SecurityGroupCreationOptionsable)()
}
