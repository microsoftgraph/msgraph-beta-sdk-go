package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalCreationConditionSet casts the previous resource to application.
type ServicePrincipalCreationConditionSet struct {
    Entity
    // The applicationIds property
    applicationIds []string
    // The applicationPublisherIds property
    applicationPublisherIds []string
    // The applicationsFromVerifiedPublisherOnly property
    applicationsFromVerifiedPublisherOnly *bool
    // The applicationTenantIds property
    applicationTenantIds []string
    // The certifiedApplicationsOnly property
    certifiedApplicationsOnly *bool
}
// NewServicePrincipalCreationConditionSet instantiates a new servicePrincipalCreationConditionSet and sets the default values.
func NewServicePrincipalCreationConditionSet()(*ServicePrincipalCreationConditionSet) {
    m := &ServicePrincipalCreationConditionSet{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.servicePrincipalCreationConditionSet";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateServicePrincipalCreationConditionSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalCreationConditionSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalCreationConditionSet(), nil
}
// GetApplicationIds gets the applicationIds property value. The applicationIds property
func (m *ServicePrincipalCreationConditionSet) GetApplicationIds()([]string) {
    return m.applicationIds
}
// GetApplicationPublisherIds gets the applicationPublisherIds property value. The applicationPublisherIds property
func (m *ServicePrincipalCreationConditionSet) GetApplicationPublisherIds()([]string) {
    return m.applicationPublisherIds
}
// GetApplicationsFromVerifiedPublisherOnly gets the applicationsFromVerifiedPublisherOnly property value. The applicationsFromVerifiedPublisherOnly property
func (m *ServicePrincipalCreationConditionSet) GetApplicationsFromVerifiedPublisherOnly()(*bool) {
    return m.applicationsFromVerifiedPublisherOnly
}
// GetApplicationTenantIds gets the applicationTenantIds property value. The applicationTenantIds property
func (m *ServicePrincipalCreationConditionSet) GetApplicationTenantIds()([]string) {
    return m.applicationTenantIds
}
// GetCertifiedApplicationsOnly gets the certifiedApplicationsOnly property value. The certifiedApplicationsOnly property
func (m *ServicePrincipalCreationConditionSet) GetCertifiedApplicationsOnly()(*bool) {
    return m.certifiedApplicationsOnly
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalCreationConditionSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetApplicationIds)
    res["applicationPublisherIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetApplicationPublisherIds)
    res["applicationsFromVerifiedPublisherOnly"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetApplicationsFromVerifiedPublisherOnly)
    res["applicationTenantIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetApplicationTenantIds)
    res["certifiedApplicationsOnly"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCertifiedApplicationsOnly)
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
    m.applicationIds = value
}
// SetApplicationPublisherIds sets the applicationPublisherIds property value. The applicationPublisherIds property
func (m *ServicePrincipalCreationConditionSet) SetApplicationPublisherIds(value []string)() {
    m.applicationPublisherIds = value
}
// SetApplicationsFromVerifiedPublisherOnly sets the applicationsFromVerifiedPublisherOnly property value. The applicationsFromVerifiedPublisherOnly property
func (m *ServicePrincipalCreationConditionSet) SetApplicationsFromVerifiedPublisherOnly(value *bool)() {
    m.applicationsFromVerifiedPublisherOnly = value
}
// SetApplicationTenantIds sets the applicationTenantIds property value. The applicationTenantIds property
func (m *ServicePrincipalCreationConditionSet) SetApplicationTenantIds(value []string)() {
    m.applicationTenantIds = value
}
// SetCertifiedApplicationsOnly sets the certifiedApplicationsOnly property value. The certifiedApplicationsOnly property
func (m *ServicePrincipalCreationConditionSet) SetCertifiedApplicationsOnly(value *bool)() {
    m.certifiedApplicationsOnly = value
}
