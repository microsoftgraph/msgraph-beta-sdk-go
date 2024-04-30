package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcServicePlan struct {
    Entity
}
// NewCloudPcServicePlan instantiates a new CloudPcServicePlan and sets the default values.
func NewCloudPcServicePlan()(*CloudPcServicePlan) {
    m := &CloudPcServicePlan{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcServicePlanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcServicePlanFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcServicePlan(), nil
}
// GetDisplayName gets the displayName property value. The name for the service plan. Read-only.
// returns a *string when successful
func (m *CloudPcServicePlan) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcServicePlan) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["provisioningType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcProvisioningType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningType(val.(*CloudPcProvisioningType))
        }
        return nil
    }
    res["ramInGB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamInGB(val)
        }
        return nil
    }
    res["storageInGB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageInGB(val)
        }
        return nil
    }
    res["supportedSolution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcManagementService)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedSolution(val.(*CloudPcManagementService))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcServicePlanType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*CloudPcServicePlanType))
        }
        return nil
    }
    res["userProfileInGB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserProfileInGB(val)
        }
        return nil
    }
    res["vCpuCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVCpuCount(val)
        }
        return nil
    }
    return res
}
// GetProvisioningType gets the provisioningType property value. Specifies the type of license used when provisioning Cloud PCs. By default, the license type is dedicated. Possible values are: dedicated, shared, unknownFutureValue, sharedByUser, sharedByEntraGroup. You must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: sharedByUser, sharedByEntraGroup. The shared member is deprecated and will stop returning on April 30, 2027; going forward, use the sharedByUser member.
// returns a *CloudPcProvisioningType when successful
func (m *CloudPcServicePlan) GetProvisioningType()(*CloudPcProvisioningType) {
    val, err := m.GetBackingStore().Get("provisioningType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcProvisioningType)
    }
    return nil
}
// GetRamInGB gets the ramInGB property value. The size of the RAM in GB. Read-only.
// returns a *int32 when successful
func (m *CloudPcServicePlan) GetRamInGB()(*int32) {
    val, err := m.GetBackingStore().Get("ramInGB")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetStorageInGB gets the storageInGB property value. The size of the OS Disk in GB. Read-only.
// returns a *int32 when successful
func (m *CloudPcServicePlan) GetStorageInGB()(*int32) {
    val, err := m.GetBackingStore().Get("storageInGB")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSupportedSolution gets the supportedSolution property value. The supportedSolution property
// returns a *CloudPcManagementService when successful
func (m *CloudPcServicePlan) GetSupportedSolution()(*CloudPcManagementService) {
    val, err := m.GetBackingStore().Get("supportedSolution")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcManagementService)
    }
    return nil
}
// GetTypeEscaped gets the type property value. The type of the service plan. Possible values are: enterprise, business, unknownFutureValue. Read-only.
// returns a *CloudPcServicePlanType when successful
func (m *CloudPcServicePlan) GetTypeEscaped()(*CloudPcServicePlanType) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcServicePlanType)
    }
    return nil
}
// GetUserProfileInGB gets the userProfileInGB property value. The size of the user profile disk in GB. Read-only.
// returns a *int32 when successful
func (m *CloudPcServicePlan) GetUserProfileInGB()(*int32) {
    val, err := m.GetBackingStore().Get("userProfileInGB")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetVCpuCount gets the vCpuCount property value. The number of vCPUs. Read-only.
// returns a *int32 when successful
func (m *CloudPcServicePlan) GetVCpuCount()(*int32) {
    val, err := m.GetBackingStore().Get("vCpuCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcServicePlan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningType() != nil {
        cast := (*m.GetProvisioningType()).String()
        err = writer.WriteStringValue("provisioningType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ramInGB", m.GetRamInGB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("storageInGB", m.GetStorageInGB())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedSolution() != nil {
        cast := (*m.GetSupportedSolution()).String()
        err = writer.WriteStringValue("supportedSolution", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("userProfileInGB", m.GetUserProfileInGB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("vCpuCount", m.GetVCpuCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name for the service plan. Read-only.
func (m *CloudPcServicePlan) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisioningType sets the provisioningType property value. Specifies the type of license used when provisioning Cloud PCs. By default, the license type is dedicated. Possible values are: dedicated, shared, unknownFutureValue, sharedByUser, sharedByEntraGroup. You must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: sharedByUser, sharedByEntraGroup. The shared member is deprecated and will stop returning on April 30, 2027; going forward, use the sharedByUser member.
func (m *CloudPcServicePlan) SetProvisioningType(value *CloudPcProvisioningType)() {
    err := m.GetBackingStore().Set("provisioningType", value)
    if err != nil {
        panic(err)
    }
}
// SetRamInGB sets the ramInGB property value. The size of the RAM in GB. Read-only.
func (m *CloudPcServicePlan) SetRamInGB(value *int32)() {
    err := m.GetBackingStore().Set("ramInGB", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageInGB sets the storageInGB property value. The size of the OS Disk in GB. Read-only.
func (m *CloudPcServicePlan) SetStorageInGB(value *int32)() {
    err := m.GetBackingStore().Set("storageInGB", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedSolution sets the supportedSolution property value. The supportedSolution property
func (m *CloudPcServicePlan) SetSupportedSolution(value *CloudPcManagementService)() {
    err := m.GetBackingStore().Set("supportedSolution", value)
    if err != nil {
        panic(err)
    }
}
// SetTypeEscaped sets the type property value. The type of the service plan. Possible values are: enterprise, business, unknownFutureValue. Read-only.
func (m *CloudPcServicePlan) SetTypeEscaped(value *CloudPcServicePlanType)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// SetUserProfileInGB sets the userProfileInGB property value. The size of the user profile disk in GB. Read-only.
func (m *CloudPcServicePlan) SetUserProfileInGB(value *int32)() {
    err := m.GetBackingStore().Set("userProfileInGB", value)
    if err != nil {
        panic(err)
    }
}
// SetVCpuCount sets the vCpuCount property value. The number of vCPUs. Read-only.
func (m *CloudPcServicePlan) SetVCpuCount(value *int32)() {
    err := m.GetBackingStore().Set("vCpuCount", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcServicePlanable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetProvisioningType()(*CloudPcProvisioningType)
    GetRamInGB()(*int32)
    GetStorageInGB()(*int32)
    GetSupportedSolution()(*CloudPcManagementService)
    GetTypeEscaped()(*CloudPcServicePlanType)
    GetUserProfileInGB()(*int32)
    GetVCpuCount()(*int32)
    SetDisplayName(value *string)()
    SetProvisioningType(value *CloudPcProvisioningType)()
    SetRamInGB(value *int32)()
    SetStorageInGB(value *int32)()
    SetSupportedSolution(value *CloudPcManagementService)()
    SetTypeEscaped(value *CloudPcServicePlanType)()
    SetUserProfileInGB(value *int32)()
    SetVCpuCount(value *int32)()
}
