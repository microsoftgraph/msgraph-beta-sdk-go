package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalCreationConditionSet 
type ServicePrincipalCreationConditionSet struct {
    Entity
}
// NewServicePrincipalCreationConditionSet instantiates a new servicePrincipalCreationConditionSet and sets the default values.
func NewServicePrincipalCreationConditionSet()(*ServicePrincipalCreationConditionSet) {
    m := &ServicePrincipalCreationConditionSet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateServicePrincipalCreationConditionSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalCreationConditionSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalCreationConditionSet(), nil
}
// GetApplicationIds gets the applicationIds property value. The applicationIds property
func (m *ServicePrincipalCreationConditionSet) GetApplicationIds()([]string) {
    val, err := m.GetBackingStore().Get("applicationIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetApplicationPublisherIds gets the applicationPublisherIds property value. The applicationPublisherIds property
func (m *ServicePrincipalCreationConditionSet) GetApplicationPublisherIds()([]string) {
    val, err := m.GetBackingStore().Get("applicationPublisherIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetApplicationsFromVerifiedPublisherOnly gets the applicationsFromVerifiedPublisherOnly property value. The applicationsFromVerifiedPublisherOnly property
func (m *ServicePrincipalCreationConditionSet) GetApplicationsFromVerifiedPublisherOnly()(*bool) {
    val, err := m.GetBackingStore().Get("applicationsFromVerifiedPublisherOnly")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplicationTenantIds gets the applicationTenantIds property value. The applicationTenantIds property
func (m *ServicePrincipalCreationConditionSet) GetApplicationTenantIds()([]string) {
    val, err := m.GetBackingStore().Get("applicationTenantIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetCertifiedApplicationsOnly gets the certifiedApplicationsOnly property value. The certifiedApplicationsOnly property
func (m *ServicePrincipalCreationConditionSet) GetCertifiedApplicationsOnly()(*bool) {
    val, err := m.GetBackingStore().Get("certifiedApplicationsOnly")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalCreationConditionSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetApplicationIds(res)
        }
        return nil
    }
    res["applicationPublisherIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetApplicationPublisherIds(res)
        }
        return nil
    }
    res["applicationsFromVerifiedPublisherOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationsFromVerifiedPublisherOnly(val)
        }
        return nil
    }
    res["applicationTenantIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetApplicationTenantIds(res)
        }
        return nil
    }
    res["certifiedApplicationsOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertifiedApplicationsOnly(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ServicePrincipalCreationConditionSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicationIds() != nil {
        err = writer.WriteCollectionOfStringValues("applicationIds", m.GetApplicationIds())
        if err != nil {
            return err
        }
    }
    if m.GetApplicationPublisherIds() != nil {
        err = writer.WriteCollectionOfStringValues("applicationPublisherIds", m.GetApplicationPublisherIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationsFromVerifiedPublisherOnly", m.GetApplicationsFromVerifiedPublisherOnly())
        if err != nil {
            return err
        }
    }
    if m.GetApplicationTenantIds() != nil {
        err = writer.WriteCollectionOfStringValues("applicationTenantIds", m.GetApplicationTenantIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("certifiedApplicationsOnly", m.GetCertifiedApplicationsOnly())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationIds sets the applicationIds property value. The applicationIds property
func (m *ServicePrincipalCreationConditionSet) SetApplicationIds(value []string)() {
    err := m.GetBackingStore().Set("applicationIds", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationPublisherIds sets the applicationPublisherIds property value. The applicationPublisherIds property
func (m *ServicePrincipalCreationConditionSet) SetApplicationPublisherIds(value []string)() {
    err := m.GetBackingStore().Set("applicationPublisherIds", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationsFromVerifiedPublisherOnly sets the applicationsFromVerifiedPublisherOnly property value. The applicationsFromVerifiedPublisherOnly property
func (m *ServicePrincipalCreationConditionSet) SetApplicationsFromVerifiedPublisherOnly(value *bool)() {
    err := m.GetBackingStore().Set("applicationsFromVerifiedPublisherOnly", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationTenantIds sets the applicationTenantIds property value. The applicationTenantIds property
func (m *ServicePrincipalCreationConditionSet) SetApplicationTenantIds(value []string)() {
    err := m.GetBackingStore().Set("applicationTenantIds", value)
    if err != nil {
        panic(err)
    }
}
// SetCertifiedApplicationsOnly sets the certifiedApplicationsOnly property value. The certifiedApplicationsOnly property
func (m *ServicePrincipalCreationConditionSet) SetCertifiedApplicationsOnly(value *bool)() {
    err := m.GetBackingStore().Set("certifiedApplicationsOnly", value)
    if err != nil {
        panic(err)
    }
}
// ServicePrincipalCreationConditionSetable 
type ServicePrincipalCreationConditionSetable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationIds()([]string)
    GetApplicationPublisherIds()([]string)
    GetApplicationsFromVerifiedPublisherOnly()(*bool)
    GetApplicationTenantIds()([]string)
    GetCertifiedApplicationsOnly()(*bool)
    SetApplicationIds(value []string)()
    SetApplicationPublisherIds(value []string)()
    SetApplicationsFromVerifiedPublisherOnly(value *bool)()
    SetApplicationTenantIds(value []string)()
    SetCertifiedApplicationsOnly(value *bool)()
}
