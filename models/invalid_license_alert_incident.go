package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InvalidLicenseAlertIncident 
type InvalidLicenseAlertIncident struct {
    UnifiedRoleManagementAlertIncident
}
// NewInvalidLicenseAlertIncident instantiates a new InvalidLicenseAlertIncident and sets the default values.
func NewInvalidLicenseAlertIncident()(*InvalidLicenseAlertIncident) {
    m := &InvalidLicenseAlertIncident{
        UnifiedRoleManagementAlertIncident: *NewUnifiedRoleManagementAlertIncident(),
    }
    odataTypeValue := "#microsoft.graph.invalidLicenseAlertIncident"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInvalidLicenseAlertIncidentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInvalidLicenseAlertIncidentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInvalidLicenseAlertIncident(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InvalidLicenseAlertIncident) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementAlertIncident.GetFieldDeserializers()
    res["tenantLicenseStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantLicenseStatus(val)
        }
        return nil
    }
    return res
}
// GetTenantLicenseStatus gets the tenantLicenseStatus property value. The tenantLicenseStatus property
func (m *InvalidLicenseAlertIncident) GetTenantLicenseStatus()(*string) {
    val, err := m.GetBackingStore().Get("tenantLicenseStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InvalidLicenseAlertIncident) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementAlertIncident.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("tenantLicenseStatus", m.GetTenantLicenseStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTenantLicenseStatus sets the tenantLicenseStatus property value. The tenantLicenseStatus property
func (m *InvalidLicenseAlertIncident) SetTenantLicenseStatus(value *string)() {
    err := m.GetBackingStore().Set("tenantLicenseStatus", value)
    if err != nil {
        panic(err)
    }
}
// InvalidLicenseAlertIncidentable 
type InvalidLicenseAlertIncidentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertIncidentable
    GetTenantLicenseStatus()(*string)
    SetTenantLicenseStatus(value *string)()
}
