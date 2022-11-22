package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesDirectorySynchronizationFeature 
type OnPremisesDirectorySynchronizationFeature struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The blockCloudObjectTakeoverThroughHardMatchEnabled property
    blockCloudObjectTakeoverThroughHardMatchEnabled *bool
    // The blockSoftMatchEnabled property
    blockSoftMatchEnabled *bool
    // The bypassDirSyncOverridesEnabled property
    bypassDirSyncOverridesEnabled *bool
    // The cloudPasswordPolicyForPasswordSyncedUsersEnabled property
    cloudPasswordPolicyForPasswordSyncedUsersEnabled *bool
    // The concurrentCredentialUpdateEnabled property
    concurrentCredentialUpdateEnabled *bool
    // The concurrentOrgIdProvisioningEnabled property
    concurrentOrgIdProvisioningEnabled *bool
    // The deviceWritebackEnabled property
    deviceWritebackEnabled *bool
    // The directoryExtensionsEnabled property
    directoryExtensionsEnabled *bool
    // The fopeConflictResolutionEnabled property
    fopeConflictResolutionEnabled *bool
    // The groupWriteBackEnabled property
    groupWriteBackEnabled *bool
    // The OdataType property
    odataType *string
    // The passwordSyncEnabled property
    passwordSyncEnabled *bool
    // The passwordWritebackEnabled property
    passwordWritebackEnabled *bool
    // The quarantineUponProxyAddressesConflictEnabled property
    quarantineUponProxyAddressesConflictEnabled *bool
    // The quarantineUponUpnConflictEnabled property
    quarantineUponUpnConflictEnabled *bool
    // The softMatchOnUpnEnabled property
    softMatchOnUpnEnabled *bool
    // The synchronizeUpnForManagedUsersEnabled property
    synchronizeUpnForManagedUsersEnabled *bool
    // The unifiedGroupWritebackEnabled property
    unifiedGroupWritebackEnabled *bool
    // The userForcePasswordChangeOnLogonEnabled property
    userForcePasswordChangeOnLogonEnabled *bool
    // The userWritebackEnabled property
    userWritebackEnabled *bool
}
// NewOnPremisesDirectorySynchronizationFeature instantiates a new onPremisesDirectorySynchronizationFeature and sets the default values.
func NewOnPremisesDirectorySynchronizationFeature()(*OnPremisesDirectorySynchronizationFeature) {
    m := &OnPremisesDirectorySynchronizationFeature{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnPremisesDirectorySynchronizationFeatureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesDirectorySynchronizationFeatureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesDirectorySynchronizationFeature(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesDirectorySynchronizationFeature) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBlockCloudObjectTakeoverThroughHardMatchEnabled gets the blockCloudObjectTakeoverThroughHardMatchEnabled property value. The blockCloudObjectTakeoverThroughHardMatchEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetBlockCloudObjectTakeoverThroughHardMatchEnabled()(*bool) {
    return m.blockCloudObjectTakeoverThroughHardMatchEnabled
}
// GetBlockSoftMatchEnabled gets the blockSoftMatchEnabled property value. The blockSoftMatchEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetBlockSoftMatchEnabled()(*bool) {
    return m.blockSoftMatchEnabled
}
// GetBypassDirSyncOverridesEnabled gets the bypassDirSyncOverridesEnabled property value. The bypassDirSyncOverridesEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetBypassDirSyncOverridesEnabled()(*bool) {
    return m.bypassDirSyncOverridesEnabled
}
// GetCloudPasswordPolicyForPasswordSyncedUsersEnabled gets the cloudPasswordPolicyForPasswordSyncedUsersEnabled property value. The cloudPasswordPolicyForPasswordSyncedUsersEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetCloudPasswordPolicyForPasswordSyncedUsersEnabled()(*bool) {
    return m.cloudPasswordPolicyForPasswordSyncedUsersEnabled
}
// GetConcurrentCredentialUpdateEnabled gets the concurrentCredentialUpdateEnabled property value. The concurrentCredentialUpdateEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetConcurrentCredentialUpdateEnabled()(*bool) {
    return m.concurrentCredentialUpdateEnabled
}
// GetConcurrentOrgIdProvisioningEnabled gets the concurrentOrgIdProvisioningEnabled property value. The concurrentOrgIdProvisioningEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetConcurrentOrgIdProvisioningEnabled()(*bool) {
    return m.concurrentOrgIdProvisioningEnabled
}
// GetDeviceWritebackEnabled gets the deviceWritebackEnabled property value. The deviceWritebackEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetDeviceWritebackEnabled()(*bool) {
    return m.deviceWritebackEnabled
}
// GetDirectoryExtensionsEnabled gets the directoryExtensionsEnabled property value. The directoryExtensionsEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetDirectoryExtensionsEnabled()(*bool) {
    return m.directoryExtensionsEnabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesDirectorySynchronizationFeature) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blockCloudObjectTakeoverThroughHardMatchEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBlockCloudObjectTakeoverThroughHardMatchEnabled)
    res["blockSoftMatchEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBlockSoftMatchEnabled)
    res["bypassDirSyncOverridesEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBypassDirSyncOverridesEnabled)
    res["cloudPasswordPolicyForPasswordSyncedUsersEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCloudPasswordPolicyForPasswordSyncedUsersEnabled)
    res["concurrentCredentialUpdateEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetConcurrentCredentialUpdateEnabled)
    res["concurrentOrgIdProvisioningEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetConcurrentOrgIdProvisioningEnabled)
    res["deviceWritebackEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDeviceWritebackEnabled)
    res["directoryExtensionsEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDirectoryExtensionsEnabled)
    res["fopeConflictResolutionEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetFopeConflictResolutionEnabled)
    res["groupWriteBackEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetGroupWriteBackEnabled)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["passwordSyncEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPasswordSyncEnabled)
    res["passwordWritebackEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPasswordWritebackEnabled)
    res["quarantineUponProxyAddressesConflictEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetQuarantineUponProxyAddressesConflictEnabled)
    res["quarantineUponUpnConflictEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetQuarantineUponUpnConflictEnabled)
    res["softMatchOnUpnEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSoftMatchOnUpnEnabled)
    res["synchronizeUpnForManagedUsersEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSynchronizeUpnForManagedUsersEnabled)
    res["unifiedGroupWritebackEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetUnifiedGroupWritebackEnabled)
    res["userForcePasswordChangeOnLogonEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetUserForcePasswordChangeOnLogonEnabled)
    res["userWritebackEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetUserWritebackEnabled)
    return res
}
// GetFopeConflictResolutionEnabled gets the fopeConflictResolutionEnabled property value. The fopeConflictResolutionEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetFopeConflictResolutionEnabled()(*bool) {
    return m.fopeConflictResolutionEnabled
}
// GetGroupWriteBackEnabled gets the groupWriteBackEnabled property value. The groupWriteBackEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetGroupWriteBackEnabled()(*bool) {
    return m.groupWriteBackEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnPremisesDirectorySynchronizationFeature) GetOdataType()(*string) {
    return m.odataType
}
// GetPasswordSyncEnabled gets the passwordSyncEnabled property value. The passwordSyncEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetPasswordSyncEnabled()(*bool) {
    return m.passwordSyncEnabled
}
// GetPasswordWritebackEnabled gets the passwordWritebackEnabled property value. The passwordWritebackEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetPasswordWritebackEnabled()(*bool) {
    return m.passwordWritebackEnabled
}
// GetQuarantineUponProxyAddressesConflictEnabled gets the quarantineUponProxyAddressesConflictEnabled property value. The quarantineUponProxyAddressesConflictEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetQuarantineUponProxyAddressesConflictEnabled()(*bool) {
    return m.quarantineUponProxyAddressesConflictEnabled
}
// GetQuarantineUponUpnConflictEnabled gets the quarantineUponUpnConflictEnabled property value. The quarantineUponUpnConflictEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetQuarantineUponUpnConflictEnabled()(*bool) {
    return m.quarantineUponUpnConflictEnabled
}
// GetSoftMatchOnUpnEnabled gets the softMatchOnUpnEnabled property value. The softMatchOnUpnEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetSoftMatchOnUpnEnabled()(*bool) {
    return m.softMatchOnUpnEnabled
}
// GetSynchronizeUpnForManagedUsersEnabled gets the synchronizeUpnForManagedUsersEnabled property value. The synchronizeUpnForManagedUsersEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetSynchronizeUpnForManagedUsersEnabled()(*bool) {
    return m.synchronizeUpnForManagedUsersEnabled
}
// GetUnifiedGroupWritebackEnabled gets the unifiedGroupWritebackEnabled property value. The unifiedGroupWritebackEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetUnifiedGroupWritebackEnabled()(*bool) {
    return m.unifiedGroupWritebackEnabled
}
// GetUserForcePasswordChangeOnLogonEnabled gets the userForcePasswordChangeOnLogonEnabled property value. The userForcePasswordChangeOnLogonEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetUserForcePasswordChangeOnLogonEnabled()(*bool) {
    return m.userForcePasswordChangeOnLogonEnabled
}
// GetUserWritebackEnabled gets the userWritebackEnabled property value. The userWritebackEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) GetUserWritebackEnabled()(*bool) {
    return m.userWritebackEnabled
}
// Serialize serializes information the current object
func (m *OnPremisesDirectorySynchronizationFeature) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("blockCloudObjectTakeoverThroughHardMatchEnabled", m.GetBlockCloudObjectTakeoverThroughHardMatchEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockSoftMatchEnabled", m.GetBlockSoftMatchEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("bypassDirSyncOverridesEnabled", m.GetBypassDirSyncOverridesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("cloudPasswordPolicyForPasswordSyncedUsersEnabled", m.GetCloudPasswordPolicyForPasswordSyncedUsersEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("concurrentCredentialUpdateEnabled", m.GetConcurrentCredentialUpdateEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("concurrentOrgIdProvisioningEnabled", m.GetConcurrentOrgIdProvisioningEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("deviceWritebackEnabled", m.GetDeviceWritebackEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("directoryExtensionsEnabled", m.GetDirectoryExtensionsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("fopeConflictResolutionEnabled", m.GetFopeConflictResolutionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("groupWriteBackEnabled", m.GetGroupWriteBackEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("passwordSyncEnabled", m.GetPasswordSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("passwordWritebackEnabled", m.GetPasswordWritebackEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("quarantineUponProxyAddressesConflictEnabled", m.GetQuarantineUponProxyAddressesConflictEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("quarantineUponUpnConflictEnabled", m.GetQuarantineUponUpnConflictEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("softMatchOnUpnEnabled", m.GetSoftMatchOnUpnEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("synchronizeUpnForManagedUsersEnabled", m.GetSynchronizeUpnForManagedUsersEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("unifiedGroupWritebackEnabled", m.GetUnifiedGroupWritebackEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("userForcePasswordChangeOnLogonEnabled", m.GetUserForcePasswordChangeOnLogonEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("userWritebackEnabled", m.GetUserWritebackEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesDirectorySynchronizationFeature) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBlockCloudObjectTakeoverThroughHardMatchEnabled sets the blockCloudObjectTakeoverThroughHardMatchEnabled property value. The blockCloudObjectTakeoverThroughHardMatchEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetBlockCloudObjectTakeoverThroughHardMatchEnabled(value *bool)() {
    m.blockCloudObjectTakeoverThroughHardMatchEnabled = value
}
// SetBlockSoftMatchEnabled sets the blockSoftMatchEnabled property value. The blockSoftMatchEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetBlockSoftMatchEnabled(value *bool)() {
    m.blockSoftMatchEnabled = value
}
// SetBypassDirSyncOverridesEnabled sets the bypassDirSyncOverridesEnabled property value. The bypassDirSyncOverridesEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetBypassDirSyncOverridesEnabled(value *bool)() {
    m.bypassDirSyncOverridesEnabled = value
}
// SetCloudPasswordPolicyForPasswordSyncedUsersEnabled sets the cloudPasswordPolicyForPasswordSyncedUsersEnabled property value. The cloudPasswordPolicyForPasswordSyncedUsersEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetCloudPasswordPolicyForPasswordSyncedUsersEnabled(value *bool)() {
    m.cloudPasswordPolicyForPasswordSyncedUsersEnabled = value
}
// SetConcurrentCredentialUpdateEnabled sets the concurrentCredentialUpdateEnabled property value. The concurrentCredentialUpdateEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetConcurrentCredentialUpdateEnabled(value *bool)() {
    m.concurrentCredentialUpdateEnabled = value
}
// SetConcurrentOrgIdProvisioningEnabled sets the concurrentOrgIdProvisioningEnabled property value. The concurrentOrgIdProvisioningEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetConcurrentOrgIdProvisioningEnabled(value *bool)() {
    m.concurrentOrgIdProvisioningEnabled = value
}
// SetDeviceWritebackEnabled sets the deviceWritebackEnabled property value. The deviceWritebackEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetDeviceWritebackEnabled(value *bool)() {
    m.deviceWritebackEnabled = value
}
// SetDirectoryExtensionsEnabled sets the directoryExtensionsEnabled property value. The directoryExtensionsEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetDirectoryExtensionsEnabled(value *bool)() {
    m.directoryExtensionsEnabled = value
}
// SetFopeConflictResolutionEnabled sets the fopeConflictResolutionEnabled property value. The fopeConflictResolutionEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetFopeConflictResolutionEnabled(value *bool)() {
    m.fopeConflictResolutionEnabled = value
}
// SetGroupWriteBackEnabled sets the groupWriteBackEnabled property value. The groupWriteBackEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetGroupWriteBackEnabled(value *bool)() {
    m.groupWriteBackEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesDirectorySynchronizationFeature) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPasswordSyncEnabled sets the passwordSyncEnabled property value. The passwordSyncEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetPasswordSyncEnabled(value *bool)() {
    m.passwordSyncEnabled = value
}
// SetPasswordWritebackEnabled sets the passwordWritebackEnabled property value. The passwordWritebackEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetPasswordWritebackEnabled(value *bool)() {
    m.passwordWritebackEnabled = value
}
// SetQuarantineUponProxyAddressesConflictEnabled sets the quarantineUponProxyAddressesConflictEnabled property value. The quarantineUponProxyAddressesConflictEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetQuarantineUponProxyAddressesConflictEnabled(value *bool)() {
    m.quarantineUponProxyAddressesConflictEnabled = value
}
// SetQuarantineUponUpnConflictEnabled sets the quarantineUponUpnConflictEnabled property value. The quarantineUponUpnConflictEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetQuarantineUponUpnConflictEnabled(value *bool)() {
    m.quarantineUponUpnConflictEnabled = value
}
// SetSoftMatchOnUpnEnabled sets the softMatchOnUpnEnabled property value. The softMatchOnUpnEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetSoftMatchOnUpnEnabled(value *bool)() {
    m.softMatchOnUpnEnabled = value
}
// SetSynchronizeUpnForManagedUsersEnabled sets the synchronizeUpnForManagedUsersEnabled property value. The synchronizeUpnForManagedUsersEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetSynchronizeUpnForManagedUsersEnabled(value *bool)() {
    m.synchronizeUpnForManagedUsersEnabled = value
}
// SetUnifiedGroupWritebackEnabled sets the unifiedGroupWritebackEnabled property value. The unifiedGroupWritebackEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetUnifiedGroupWritebackEnabled(value *bool)() {
    m.unifiedGroupWritebackEnabled = value
}
// SetUserForcePasswordChangeOnLogonEnabled sets the userForcePasswordChangeOnLogonEnabled property value. The userForcePasswordChangeOnLogonEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetUserForcePasswordChangeOnLogonEnabled(value *bool)() {
    m.userForcePasswordChangeOnLogonEnabled = value
}
// SetUserWritebackEnabled sets the userWritebackEnabled property value. The userWritebackEnabled property
func (m *OnPremisesDirectorySynchronizationFeature) SetUserWritebackEnabled(value *bool)() {
    m.userWritebackEnabled = value
}
