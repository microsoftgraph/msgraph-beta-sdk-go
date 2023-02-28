package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalCreationPolicy 
type ServicePrincipalCreationPolicy struct {
    PolicyBase
}
// NewServicePrincipalCreationPolicy instantiates a new ServicePrincipalCreationPolicy and sets the default values.
func NewServicePrincipalCreationPolicy()(*ServicePrincipalCreationPolicy) {
    m := &ServicePrincipalCreationPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    odataTypeValue := "#microsoft.graph.servicePrincipalCreationPolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateServicePrincipalCreationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalCreationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalCreationPolicy(), nil
}
// GetExcludes gets the excludes property value. The excludes property
func (m *ServicePrincipalCreationPolicy) GetExcludes()([]ServicePrincipalCreationConditionSetable) {
    val, err := m.GetBackingStore().Get("excludes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ServicePrincipalCreationConditionSetable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalCreationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["excludes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicePrincipalCreationConditionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePrincipalCreationConditionSetable, len(val))
            for i, v := range val {
                res[i] = v.(ServicePrincipalCreationConditionSetable)
            }
            m.SetExcludes(res)
        }
        return nil
    }
    res["includes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicePrincipalCreationConditionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePrincipalCreationConditionSetable, len(val))
            for i, v := range val {
                res[i] = v.(ServicePrincipalCreationConditionSetable)
            }
            m.SetIncludes(res)
        }
        return nil
    }
    res["isBuiltIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBuiltIn(val)
        }
        return nil
    }
    return res
}
// GetIncludes gets the includes property value. The includes property
func (m *ServicePrincipalCreationPolicy) GetIncludes()([]ServicePrincipalCreationConditionSetable) {
    val, err := m.GetBackingStore().Get("includes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ServicePrincipalCreationConditionSetable)
    }
    return nil
}
// GetIsBuiltIn gets the isBuiltIn property value. The isBuiltIn property
func (m *ServicePrincipalCreationPolicy) GetIsBuiltIn()(*bool) {
    val, err := m.GetBackingStore().Get("isBuiltIn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ServicePrincipalCreationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExcludes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExcludes()))
        for i, v := range m.GetExcludes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("excludes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncludes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludes()))
        for i, v := range m.GetIncludes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    err := m.GetBackingStore().Set("excludes", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludes sets the includes property value. The includes property
func (m *ServicePrincipalCreationPolicy) SetIncludes(value []ServicePrincipalCreationConditionSetable)() {
    err := m.GetBackingStore().Set("includes", value)
    if err != nil {
        panic(err)
    }
}
// SetIsBuiltIn sets the isBuiltIn property value. The isBuiltIn property
func (m *ServicePrincipalCreationPolicy) SetIsBuiltIn(value *bool)() {
    err := m.GetBackingStore().Set("isBuiltIn", value)
    if err != nil {
        panic(err)
    }
}
// ServicePrincipalCreationPolicyable 
type ServicePrincipalCreationPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyBaseable
    GetExcludes()([]ServicePrincipalCreationConditionSetable)
    GetIncludes()([]ServicePrincipalCreationConditionSetable)
    GetIsBuiltIn()(*bool)
    SetExcludes(value []ServicePrincipalCreationConditionSetable)()
    SetIncludes(value []ServicePrincipalCreationConditionSetable)()
    SetIsBuiltIn(value *bool)()
}
