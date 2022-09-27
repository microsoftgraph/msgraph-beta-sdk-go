package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalCreationPolicy 
type ServicePrincipalCreationPolicy struct {
    PolicyBase
    // The excludes property
    excludes []ServicePrincipalCreationConditionSetable
    // The includes property
    includes []ServicePrincipalCreationConditionSetable
    // The isBuiltIn property
    isBuiltIn *bool
}
// NewServicePrincipalCreationPolicy instantiates a new ServicePrincipalCreationPolicy and sets the default values.
func NewServicePrincipalCreationPolicy()(*ServicePrincipalCreationPolicy) {
    m := &ServicePrincipalCreationPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    odataTypeValue := "#microsoft.graph.servicePrincipalCreationPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateServicePrincipalCreationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalCreationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalCreationPolicy(), nil
}
// GetExcludes gets the excludes property value. The excludes property
func (m *ServicePrincipalCreationPolicy) GetExcludes()([]ServicePrincipalCreationConditionSetable) {
    return m.excludes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalCreationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["excludes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateServicePrincipalCreationConditionSetFromDiscriminatorValue , m.SetExcludes)
    res["includes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateServicePrincipalCreationConditionSetFromDiscriminatorValue , m.SetIncludes)
    res["isBuiltIn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsBuiltIn)
    return res
}
// GetIncludes gets the includes property value. The includes property
func (m *ServicePrincipalCreationPolicy) GetIncludes()([]ServicePrincipalCreationConditionSetable) {
    return m.includes
}
// GetIsBuiltIn gets the isBuiltIn property value. The isBuiltIn property
func (m *ServicePrincipalCreationPolicy) GetIsBuiltIn()(*bool) {
    return m.isBuiltIn
}
// Serialize serializes information the current object
func (m *ServicePrincipalCreationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExcludes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExcludes())
        err = writer.WriteCollectionOfObjectValues("excludes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncludes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetIncludes())
        err = writer.WriteCollectionOfObjectValues("includes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBuiltIn", m.GetIsBuiltIn())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExcludes sets the excludes property value. The excludes property
func (m *ServicePrincipalCreationPolicy) SetExcludes(value []ServicePrincipalCreationConditionSetable)() {
    m.excludes = value
}
// SetIncludes sets the includes property value. The includes property
func (m *ServicePrincipalCreationPolicy) SetIncludes(value []ServicePrincipalCreationConditionSetable)() {
    m.includes = value
}
// SetIsBuiltIn sets the isBuiltIn property value. The isBuiltIn property
func (m *ServicePrincipalCreationPolicy) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
