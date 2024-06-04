package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EndpointPrivilegeManagementProvisioningStatus endpoint privilege management (EPM) tenant provisioning status contains tenant level license and onboarding state information.
type EndpointPrivilegeManagementProvisioningStatus struct {
    Entity
}
// NewEndpointPrivilegeManagementProvisioningStatus instantiates a new EndpointPrivilegeManagementProvisioningStatus and sets the default values.
func NewEndpointPrivilegeManagementProvisioningStatus()(*EndpointPrivilegeManagementProvisioningStatus) {
    m := &EndpointPrivilegeManagementProvisioningStatus{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEndpointPrivilegeManagementProvisioningStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEndpointPrivilegeManagementProvisioningStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEndpointPrivilegeManagementProvisioningStatus(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EndpointPrivilegeManagementProvisioningStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["licenseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseType(val.(*LicenseType))
        }
        return nil
    }
    res["onboardedToMicrosoftManagedPlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardedToMicrosoftManagedPlatform(val)
        }
        return nil
    }
    return res
}
// GetLicenseType gets the licenseType property value. Indicates whether tenant has a valid Intune Endpoint Privilege Management license. Possible value are : 0 - notPaid, 1 - paid, 2 - trial. See LicenseType enum for more details. Default notPaid .
// returns a *LicenseType when successful
func (m *EndpointPrivilegeManagementProvisioningStatus) GetLicenseType()(*LicenseType) {
    val, err := m.GetBackingStore().Get("licenseType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LicenseType)
    }
    return nil
}
// GetOnboardedToMicrosoftManagedPlatform gets the onboardedToMicrosoftManagedPlatform property value. Indicates whether tenant is onboarded to Microsoft Managed Platform - Cloud (MMPC). When set to true, implies tenant is onboarded and when set to false, implies tenant is not onboarded. Default set to false.
// returns a *bool when successful
func (m *EndpointPrivilegeManagementProvisioningStatus) GetOnboardedToMicrosoftManagedPlatform()(*bool) {
    val, err := m.GetBackingStore().Get("onboardedToMicrosoftManagedPlatform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EndpointPrivilegeManagementProvisioningStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLicenseType() != nil {
        cast := (*m.GetLicenseType()).String()
        err = writer.WriteStringValue("licenseType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onboardedToMicrosoftManagedPlatform", m.GetOnboardedToMicrosoftManagedPlatform())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLicenseType sets the licenseType property value. Indicates whether tenant has a valid Intune Endpoint Privilege Management license. Possible value are : 0 - notPaid, 1 - paid, 2 - trial. See LicenseType enum for more details. Default notPaid .
func (m *EndpointPrivilegeManagementProvisioningStatus) SetLicenseType(value *LicenseType)() {
    err := m.GetBackingStore().Set("licenseType", value)
    if err != nil {
        panic(err)
    }
}
// SetOnboardedToMicrosoftManagedPlatform sets the onboardedToMicrosoftManagedPlatform property value. Indicates whether tenant is onboarded to Microsoft Managed Platform - Cloud (MMPC). When set to true, implies tenant is onboarded and when set to false, implies tenant is not onboarded. Default set to false.
func (m *EndpointPrivilegeManagementProvisioningStatus) SetOnboardedToMicrosoftManagedPlatform(value *bool)() {
    err := m.GetBackingStore().Set("onboardedToMicrosoftManagedPlatform", value)
    if err != nil {
        panic(err)
    }
}
type EndpointPrivilegeManagementProvisioningStatusable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLicenseType()(*LicenseType)
    GetOnboardedToMicrosoftManagedPlatform()(*bool)
    SetLicenseType(value *LicenseType)()
    SetOnboardedToMicrosoftManagedPlatform(value *bool)()
}
