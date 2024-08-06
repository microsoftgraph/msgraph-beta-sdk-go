package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceAssignmentItem represents the application or configuration included in the ChangeAssignments action execution or result. For action execution, it represents the application or configuration intended to be uninstalled or removed on the managed device. For action result, it represents the live reporting data for this application or configuration regarding its removal or restoration process.
type DeviceAssignmentItem struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceAssignmentItem instantiates a new DeviceAssignmentItem and sets the default values.
func NewDeviceAssignmentItem()(*DeviceAssignmentItem) {
    m := &DeviceAssignmentItem{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceAssignmentItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceAssignmentItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAssignmentItem(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeviceAssignmentItem) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAssignmentItemActionIntent gets the assignmentItemActionIntent property value. A list of possible assignment item action intent values on the application or configuration when executing this action on the managed device. For example, if the application or configuration is intended to be removed on the managed device, then the intent value is remove, and if the application or configuration already under removal through previous actions and is now intended to be restored on the managed device, then the intent value is restore
// returns a *DeviceAssignmentItemIntent when successful
func (m *DeviceAssignmentItem) GetAssignmentItemActionIntent()(*DeviceAssignmentItemIntent) {
    val, err := m.GetBackingStore().Get("assignmentItemActionIntent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceAssignmentItemIntent)
    }
    return nil
}
// GetAssignmentItemActionStatus gets the assignmentItemActionStatus property value. A list of possible assignment item action status values for the application or configuration regarding their executed action on the managed device. For example, a configuration included in the deviceAssignmentItems list has just been executed the action. Its status starts with inProgress until it's successfully removed to reflect as removed status or failed to be removed to reflect as error status on the managed device. Similar status change happens for restoration process
// returns a *DeviceAssignmentItemStatus when successful
func (m *DeviceAssignmentItem) GetAssignmentItemActionStatus()(*DeviceAssignmentItemStatus) {
    val, err := m.GetBackingStore().Get("assignmentItemActionStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceAssignmentItemStatus)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DeviceAssignmentItem) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetErrorCode gets the errorCode property value. The error code for the application or configuration regarding the failed executed action on the managed device. Read-Only. Returned in the action result. 0 is default value and indicates no failure. Valid values -9.22337203685478E+18 to 9.22337203685478E+18. This property is read-only.
// returns a *int64 when successful
func (m *DeviceAssignmentItem) GetErrorCode()(*int64) {
    val, err := m.GetBackingStore().Get("errorCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceAssignmentItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentItemActionIntent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAssignmentItemIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentItemActionIntent(val.(*DeviceAssignmentItemIntent))
        }
        return nil
    }
    res["assignmentItemActionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAssignmentItemStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentItemActionStatus(val.(*DeviceAssignmentItemStatus))
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["intentActionMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntentActionMessage(val)
        }
        return nil
    }
    res["itemDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemDisplayName(val)
        }
        return nil
    }
    res["itemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemId(val)
        }
        return nil
    }
    res["itemSubTypeDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemSubTypeDisplayName(val)
        }
        return nil
    }
    res["itemType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAssignmentItemType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemType(val.(*DeviceAssignmentItemType))
        }
        return nil
    }
    res["lastActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
    return res
}
// GetIntentActionMessage gets the intentActionMessage property value. The intent action message for the application or configuration regarding the executed action on the managed device. When the action is on error, this property provides message on the reason of failure. When the action is in progress, this property provides message on what's being processed on the device. Read-Only. Returned in the action result. Can be null. Max length is 1500. This property is read-only.
// returns a *string when successful
func (m *DeviceAssignmentItem) GetIntentActionMessage()(*string) {
    val, err := m.GetBackingStore().Get("intentActionMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetItemDisplayName gets the itemDisplayName property value. The item displayName name for the application or configuration. Read-Only. Returned in the action result. Default value is null. The property value cannot be modified and is automatically populated with the action result. Max length is 200. This property is read-only.
// returns a *string when successful
func (m *DeviceAssignmentItem) GetItemDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("itemDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetItemId gets the itemId property value. The unique identifier for the application or configuration. ItemId is required property which needs to be set in the action POST request parameter for the DeviceAssignmentItem intended to remove. Max length is 40
// returns a *string when successful
func (m *DeviceAssignmentItem) GetItemId()(*string) {
    val, err := m.GetBackingStore().Get("itemId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetItemSubTypeDisplayName gets the itemSubTypeDisplayName property value. Indicates the specific type for the application or configuration. For example, unknown, application, appConfiguration, exploitProtection, bitLocker, deviceControl, microsoftEdgeBaseline, attackSurfaceReductionRulesConfigMgr, endpointDetectionandResponse, windowsUpdateforBusiness, microsoftDefenderFirewallRules, applicationControl, microsoftDefenderAntivirusexclusions, microsoftDefenderAntivirus, wiredNetwork, derivedPersonalIdentityVerificationCredential, windowsHealthMonitoring, extensions, mxProfileZebraOnly, deviceFirmwareConfigurationInterface, deliveryOptimization, identityProtection, kiosk, overrideGroupPolicy, domainJoinPreview, pkcsImportedCertificate, networkBoundary, endpointProtection, microsoftDefenderAtpWindows10Desktop, sharedMultiUserDevice, deviceFeatures, secureAssessmentEducation, wiFiImport, editionUpgradeAndModeSwitch, vpn, custom, softwareUpdates, deviceRestrictionsWindows10Team, email, trustedCertificate, scepCertificate, emailSamsungKnoxOnly, pkcsCertificate, deviceRestrictions, wiFi, settingsCatalog. Read-Only. Returned in the action result. Default value is null. The property value cannot be modified and is automatically populated with the action result. Max length is 200. This property is read-only.
// returns a *string when successful
func (m *DeviceAssignmentItem) GetItemSubTypeDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("itemSubTypeDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetItemType gets the itemType property value. A list of possible device assignment item types to execute this action on the managed device. Device assignment item represents existing assigned Intune resource such as application or configuration. Currently supported device assignment item types are Application, DeviceConfiguration, DeviceManagementConfigurationPolicy and MobileAppConfiguration
// returns a *DeviceAssignmentItemType when successful
func (m *DeviceAssignmentItem) GetItemType()(*DeviceAssignmentItemType) {
    val, err := m.GetBackingStore().Get("itemType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceAssignmentItemType)
    }
    return nil
}
// GetLastActionDateTime gets the lastActionDateTime property value. The date and time when the application or configuration was initiated an action execution. Read-Only. Returned in the action result. The property value cannot be modified and is automatically populated when the action is initiated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2025 would look like this: '2025-01-01T00:00:00Z'. This property is read-only.
// returns a *Time when successful
func (m *DeviceAssignmentItem) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastActionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the application or configuration was last modified because of either action execution or status change. Read-Only. Returned in the action result. The property value cannot be modified and is automatically populated when the action is initiated or the device has a status change. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2025 would look like this: '2025-01-01T00:00:00Z'. This property is read-only.
// returns a *Time when successful
func (m *DeviceAssignmentItem) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DeviceAssignmentItem) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceAssignmentItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentItemActionIntent() != nil {
        cast := (*m.GetAssignmentItemActionIntent()).String()
        err := writer.WriteStringValue("assignmentItemActionIntent", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentItemActionStatus() != nil {
        cast := (*m.GetAssignmentItemActionStatus()).String()
        err := writer.WriteStringValue("assignmentItemActionStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("itemId", m.GetItemId())
        if err != nil {
            return err
        }
    }
    if m.GetItemType() != nil {
        cast := (*m.GetItemType()).String()
        err := writer.WriteStringValue("itemType", &cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAssignmentItem) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentItemActionIntent sets the assignmentItemActionIntent property value. A list of possible assignment item action intent values on the application or configuration when executing this action on the managed device. For example, if the application or configuration is intended to be removed on the managed device, then the intent value is remove, and if the application or configuration already under removal through previous actions and is now intended to be restored on the managed device, then the intent value is restore
func (m *DeviceAssignmentItem) SetAssignmentItemActionIntent(value *DeviceAssignmentItemIntent)() {
    err := m.GetBackingStore().Set("assignmentItemActionIntent", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentItemActionStatus sets the assignmentItemActionStatus property value. A list of possible assignment item action status values for the application or configuration regarding their executed action on the managed device. For example, a configuration included in the deviceAssignmentItems list has just been executed the action. Its status starts with inProgress until it's successfully removed to reflect as removed status or failed to be removed to reflect as error status on the managed device. Similar status change happens for restoration process
func (m *DeviceAssignmentItem) SetAssignmentItemActionStatus(value *DeviceAssignmentItemStatus)() {
    err := m.GetBackingStore().Set("assignmentItemActionStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeviceAssignmentItem) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetErrorCode sets the errorCode property value. The error code for the application or configuration regarding the failed executed action on the managed device. Read-Only. Returned in the action result. 0 is default value and indicates no failure. Valid values -9.22337203685478E+18 to 9.22337203685478E+18. This property is read-only.
func (m *DeviceAssignmentItem) SetErrorCode(value *int64)() {
    err := m.GetBackingStore().Set("errorCode", value)
    if err != nil {
        panic(err)
    }
}
// SetIntentActionMessage sets the intentActionMessage property value. The intent action message for the application or configuration regarding the executed action on the managed device. When the action is on error, this property provides message on the reason of failure. When the action is in progress, this property provides message on what's being processed on the device. Read-Only. Returned in the action result. Can be null. Max length is 1500. This property is read-only.
func (m *DeviceAssignmentItem) SetIntentActionMessage(value *string)() {
    err := m.GetBackingStore().Set("intentActionMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetItemDisplayName sets the itemDisplayName property value. The item displayName name for the application or configuration. Read-Only. Returned in the action result. Default value is null. The property value cannot be modified and is automatically populated with the action result. Max length is 200. This property is read-only.
func (m *DeviceAssignmentItem) SetItemDisplayName(value *string)() {
    err := m.GetBackingStore().Set("itemDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetItemId sets the itemId property value. The unique identifier for the application or configuration. ItemId is required property which needs to be set in the action POST request parameter for the DeviceAssignmentItem intended to remove. Max length is 40
func (m *DeviceAssignmentItem) SetItemId(value *string)() {
    err := m.GetBackingStore().Set("itemId", value)
    if err != nil {
        panic(err)
    }
}
// SetItemSubTypeDisplayName sets the itemSubTypeDisplayName property value. Indicates the specific type for the application or configuration. For example, unknown, application, appConfiguration, exploitProtection, bitLocker, deviceControl, microsoftEdgeBaseline, attackSurfaceReductionRulesConfigMgr, endpointDetectionandResponse, windowsUpdateforBusiness, microsoftDefenderFirewallRules, applicationControl, microsoftDefenderAntivirusexclusions, microsoftDefenderAntivirus, wiredNetwork, derivedPersonalIdentityVerificationCredential, windowsHealthMonitoring, extensions, mxProfileZebraOnly, deviceFirmwareConfigurationInterface, deliveryOptimization, identityProtection, kiosk, overrideGroupPolicy, domainJoinPreview, pkcsImportedCertificate, networkBoundary, endpointProtection, microsoftDefenderAtpWindows10Desktop, sharedMultiUserDevice, deviceFeatures, secureAssessmentEducation, wiFiImport, editionUpgradeAndModeSwitch, vpn, custom, softwareUpdates, deviceRestrictionsWindows10Team, email, trustedCertificate, scepCertificate, emailSamsungKnoxOnly, pkcsCertificate, deviceRestrictions, wiFi, settingsCatalog. Read-Only. Returned in the action result. Default value is null. The property value cannot be modified and is automatically populated with the action result. Max length is 200. This property is read-only.
func (m *DeviceAssignmentItem) SetItemSubTypeDisplayName(value *string)() {
    err := m.GetBackingStore().Set("itemSubTypeDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetItemType sets the itemType property value. A list of possible device assignment item types to execute this action on the managed device. Device assignment item represents existing assigned Intune resource such as application or configuration. Currently supported device assignment item types are Application, DeviceConfiguration, DeviceManagementConfigurationPolicy and MobileAppConfiguration
func (m *DeviceAssignmentItem) SetItemType(value *DeviceAssignmentItemType)() {
    err := m.GetBackingStore().Set("itemType", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. The date and time when the application or configuration was initiated an action execution. Read-Only. Returned in the action result. The property value cannot be modified and is automatically populated when the action is initiated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2025 would look like this: '2025-01-01T00:00:00Z'. This property is read-only.
func (m *DeviceAssignmentItem) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastActionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the application or configuration was last modified because of either action execution or status change. Read-Only. Returned in the action result. The property value cannot be modified and is automatically populated when the action is initiated or the device has a status change. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2025 would look like this: '2025-01-01T00:00:00Z'. This property is read-only.
func (m *DeviceAssignmentItem) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceAssignmentItem) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type DeviceAssignmentItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentItemActionIntent()(*DeviceAssignmentItemIntent)
    GetAssignmentItemActionStatus()(*DeviceAssignmentItemStatus)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetErrorCode()(*int64)
    GetIntentActionMessage()(*string)
    GetItemDisplayName()(*string)
    GetItemId()(*string)
    GetItemSubTypeDisplayName()(*string)
    GetItemType()(*DeviceAssignmentItemType)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    SetAssignmentItemActionIntent(value *DeviceAssignmentItemIntent)()
    SetAssignmentItemActionStatus(value *DeviceAssignmentItemStatus)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetErrorCode(value *int64)()
    SetIntentActionMessage(value *string)()
    SetItemDisplayName(value *string)()
    SetItemId(value *string)()
    SetItemSubTypeDisplayName(value *string)()
    SetItemType(value *DeviceAssignmentItemType)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
}
