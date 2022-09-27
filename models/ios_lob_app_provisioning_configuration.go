package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosLobAppProvisioningConfiguration this topic provides descriptions of the declared methods, properties and relationships exposed by the iOS Lob App Provisioning Configuration resource.
type IosLobAppProvisioningConfiguration struct {
    Entity
    // The associated group assignments for IosLobAppProvisioningConfiguration.
    assignments []IosLobAppProvisioningConfigurationAssignmentable
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Admin provided description of the Device Configuration.
    description *string
    // The list of device installation states for this mobile app configuration.
    deviceStatuses []ManagedDeviceMobileAppConfigurationDeviceStatusable
    // Admin provided name of the device configuration.
    displayName *string
    // Optional profile expiration date and time.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The associated group assignments.
    groupAssignments []MobileAppProvisioningConfigGroupAssignmentable
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Payload. (UTF8 encoded byte array)
    payload []byte
    // Payload file name (.mobileprovision
    payloadFileName *string
    // List of Scope Tags for this iOS LOB app provisioning configuration entity.
    roleScopeTagIds []string
    // The list of user installation states for this mobile app configuration.
    userStatuses []ManagedDeviceMobileAppConfigurationUserStatusable
    // Version of the device configuration.
    version *int32
}
// NewIosLobAppProvisioningConfiguration instantiates a new iosLobAppProvisioningConfiguration and sets the default values.
func NewIosLobAppProvisioningConfiguration()(*IosLobAppProvisioningConfiguration) {
    m := &IosLobAppProvisioningConfiguration{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.iosLobAppProvisioningConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosLobAppProvisioningConfiguration(), nil
}
// GetAssignments gets the assignments property value. The associated group assignments for IosLobAppProvisioningConfiguration.
func (m *IosLobAppProvisioningConfiguration) GetAssignments()([]IosLobAppProvisioningConfigurationAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. DateTime the object was created.
func (m *IosLobAppProvisioningConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Admin provided description of the Device Configuration.
func (m *IosLobAppProvisioningConfiguration) GetDescription()(*string) {
    return m.description
}
// GetDeviceStatuses gets the deviceStatuses property value. The list of device installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) GetDeviceStatuses()([]ManagedDeviceMobileAppConfigurationDeviceStatusable) {
    return m.deviceStatuses
}
// GetDisplayName gets the displayName property value. Admin provided name of the device configuration.
func (m *IosLobAppProvisioningConfiguration) GetDisplayName()(*string) {
    return m.displayName
}
// GetExpirationDateTime gets the expirationDateTime property value. Optional profile expiration date and time.
func (m *IosLobAppProvisioningConfiguration) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosLobAppProvisioningConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIosLobAppProvisioningConfigurationAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["deviceStatuses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationDeviceStatusFromDiscriminatorValue , m.SetDeviceStatuses)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["expirationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetExpirationDateTime)
    res["groupAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppProvisioningConfigGroupAssignmentFromDiscriminatorValue , m.SetGroupAssignments)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["payload"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetPayload)
    res["payloadFileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPayloadFileName)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["userStatuses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationUserStatusFromDiscriminatorValue , m.SetUserStatuses)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetVersion)
    return res
}
// GetGroupAssignments gets the groupAssignments property value. The associated group assignments.
func (m *IosLobAppProvisioningConfiguration) GetGroupAssignments()([]MobileAppProvisioningConfigGroupAssignmentable) {
    return m.groupAssignments
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *IosLobAppProvisioningConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPayload gets the payload property value. Payload. (UTF8 encoded byte array)
func (m *IosLobAppProvisioningConfiguration) GetPayload()([]byte) {
    return m.payload
}
// GetPayloadFileName gets the payloadFileName property value. Payload file name (.mobileprovision
func (m *IosLobAppProvisioningConfiguration) GetPayloadFileName()(*string) {
    return m.payloadFileName
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this iOS LOB app provisioning configuration entity.
func (m *IosLobAppProvisioningConfiguration) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetUserStatuses gets the userStatuses property value. The list of user installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) GetUserStatuses()([]ManagedDeviceMobileAppConfigurationUserStatusable) {
    return m.userStatuses
}
// GetVersion gets the version property value. Version of the device configuration.
func (m *IosLobAppProvisioningConfiguration) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *IosLobAppProvisioningConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStatuses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceStatuses())
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetGroupAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGroupAssignments())
        err = writer.WriteCollectionOfObjectValues("groupAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payloadFileName", m.GetPayloadFileName())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetUserStatuses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserStatuses())
        err = writer.WriteCollectionOfObjectValues("userStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The associated group assignments for IosLobAppProvisioningConfiguration.
func (m *IosLobAppProvisioningConfiguration) SetAssignments(value []IosLobAppProvisioningConfigurationAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. DateTime the object was created.
func (m *IosLobAppProvisioningConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Admin provided description of the Device Configuration.
func (m *IosLobAppProvisioningConfiguration) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceStatuses sets the deviceStatuses property value. The list of device installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) SetDeviceStatuses(value []ManagedDeviceMobileAppConfigurationDeviceStatusable)() {
    m.deviceStatuses = value
}
// SetDisplayName sets the displayName property value. Admin provided name of the device configuration.
func (m *IosLobAppProvisioningConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExpirationDateTime sets the expirationDateTime property value. Optional profile expiration date and time.
func (m *IosLobAppProvisioningConfiguration) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetGroupAssignments sets the groupAssignments property value. The associated group assignments.
func (m *IosLobAppProvisioningConfiguration) SetGroupAssignments(value []MobileAppProvisioningConfigGroupAssignmentable)() {
    m.groupAssignments = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *IosLobAppProvisioningConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPayload sets the payload property value. Payload. (UTF8 encoded byte array)
func (m *IosLobAppProvisioningConfiguration) SetPayload(value []byte)() {
    m.payload = value
}
// SetPayloadFileName sets the payloadFileName property value. Payload file name (.mobileprovision
func (m *IosLobAppProvisioningConfiguration) SetPayloadFileName(value *string)() {
    m.payloadFileName = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this iOS LOB app provisioning configuration entity.
func (m *IosLobAppProvisioningConfiguration) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetUserStatuses sets the userStatuses property value. The list of user installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) SetUserStatuses(value []ManagedDeviceMobileAppConfigurationUserStatusable)() {
    m.userStatuses = value
}
// SetVersion sets the version property value. Version of the device configuration.
func (m *IosLobAppProvisioningConfiguration) SetVersion(value *int32)() {
    m.version = value
}
