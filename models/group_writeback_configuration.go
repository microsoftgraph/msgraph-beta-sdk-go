package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupWritebackConfiguration 
type GroupWritebackConfiguration struct {
    WritebackConfiguration
}
// NewGroupWritebackConfiguration instantiates a new groupWritebackConfiguration and sets the default values.
func NewGroupWritebackConfiguration()(*GroupWritebackConfiguration) {
    m := &GroupWritebackConfiguration{
        WritebackConfiguration: *NewWritebackConfiguration(),
    }
    return m
}
// CreateGroupWritebackConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupWritebackConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupWritebackConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupWritebackConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WritebackConfiguration.GetFieldDeserializers()
    res["onPremisesGroupType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesGroupType(val)
        }
        return nil
    }
    return res
}
// GetOnPremisesGroupType gets the onPremisesGroupType property value. Indicates the target on-premises group type the cloud object is written back as. Nullable. The possible values are: universalDistributionGroup, universalSecurityGroup, universalMailEnabledSecurityGroup.If the cloud group is a unified (Microsoft 365) group, this property can be one of the following: universalDistributionGroup, universalSecurityGroup, universalMailEnabledSecurityGroup. Microsoft Entra security groups can be written back as universalSecurityGroup. If isEnabled or the NewUnifiedGroupWritebackDefault group setting is true but this property isn't explicitly configured: Microsoft 365 groups are written back as universalDistributionGroup by defaultSecurity groups are written back as universalSecurityGroup by default
func (m *GroupWritebackConfiguration) GetOnPremisesGroupType()(*string) {
    val, err := m.GetBackingStore().Get("onPremisesGroupType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupWritebackConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WritebackConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("onPremisesGroupType", m.GetOnPremisesGroupType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOnPremisesGroupType sets the onPremisesGroupType property value. Indicates the target on-premises group type the cloud object is written back as. Nullable. The possible values are: universalDistributionGroup, universalSecurityGroup, universalMailEnabledSecurityGroup.If the cloud group is a unified (Microsoft 365) group, this property can be one of the following: universalDistributionGroup, universalSecurityGroup, universalMailEnabledSecurityGroup. Microsoft Entra security groups can be written back as universalSecurityGroup. If isEnabled or the NewUnifiedGroupWritebackDefault group setting is true but this property isn't explicitly configured: Microsoft 365 groups are written back as universalDistributionGroup by defaultSecurity groups are written back as universalSecurityGroup by default
func (m *GroupWritebackConfiguration) SetOnPremisesGroupType(value *string)() {
    err := m.GetBackingStore().Set("onPremisesGroupType", value)
    if err != nil {
        panic(err)
    }
}
// GroupWritebackConfigurationable 
type GroupWritebackConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WritebackConfigurationable
    GetOnPremisesGroupType()(*string)
    SetOnPremisesGroupType(value *string)()
}
