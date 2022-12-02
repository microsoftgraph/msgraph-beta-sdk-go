package models

import (
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
    res["blockCloudObjectTakeoverThroughHardMatchEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockCloudObjectTakeoverThroughHardMatchEnabled(val)
        }
        return nil
    }
    res["blockSoftMatchEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockSoftMatchEnabled(val)
        }
        return nil
    }
    res["bypassDirSyncOverridesEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBypassDirSyncOverridesEnabled(val)
        }
        return nil
    }
    res["cloudPasswordPolicyForPasswordSyncedUsersEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPasswordPolicyForPasswordSyncedUsersEnabled(val)
        }
        return nil
    }
    res["concurrentCredentialUpdateEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConcurrentCredentialUpdateEnabled(val)
        }
        return nil
    }
    res["concurrentOrgIdProvisioningEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConcurrentOrgIdProvisioningEnabled(val)
        }
        return nil
    }
    res["deviceWritebackEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceWritebackEnabled(val)
        }
        return nil
    }
    res["directoryExtensionsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryExtensionsEnabled(val)
        }
        return nil
    }
    res["fopeConflictResolutionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFopeConflictResolutionEnabled(val)
        }
        return nil
    }
    res["groupWriteBackEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupWriteBackEnabled(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["passwordSyncEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordSyncEnabled(val)
        }
        return nil
    }
    res["passwordWritebackEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordWritebackEnabled(val)
        }
        return nil
    }
    res["quarantineUponProxyAddressesConflictEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuarantineUponProxyAddressesConflictEnabled(val)
        }
        return nil
    }
    res["quarantineUponUpnConflictEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuarantineUponUpnConflictEnabled(val)
        }
        return nil
    }
    res["softMatchOnUpnEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftMatchOnUpnEnabled(val)
        }
        return nil
    }
    res["synchronizeUpnForManagedUsersEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizeUpnForManagedUsersEnabled(val)
        }
        return nil
    }
    res["unifiedGroupWritebackEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnifiedGroupWritebackEnabled(val)
        }
        return nil
    }
    res["userForcePasswordChangeOnLogonEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserForcePasswordChangeOnLogonEnabled(val)
        }
        return nil
    }
    res["userWritebackEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserWritebackEnabled(val)
        }
        return nil
    }
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
