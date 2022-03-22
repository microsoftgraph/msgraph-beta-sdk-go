package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppRelationshipState describes the installation status details of the child app in the context of UPN and device id.
type MobileAppRelationshipState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The corresponding device id.
    deviceId *string;
    // The error code for install or uninstall failures of target app.
    errorCode *int32;
    // The install state of the app of target app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
    installState *ResultantAppState;
    // The install state detail of the app. Possible values are: noAdditionalDetails, dependencyFailedToInstall, dependencyWithRequirementsNotMet, dependencyPendingReboot, dependencyWithAutoInstallDisabled, supersededAppUninstallFailed, supersededAppUninstallPendingReboot, removingSupersededApps, iosAppStoreUpdateFailedToInstall, vppAppHasUpdateAvailable, userRejectedUpdate, uninstallPendingReboot, supersedingAppsDetected, supersededAppsDetected, seeInstallErrorCode, autoInstallDisabled, managedAppNoLongerPresent, userRejectedInstall, userIsNotLoggedIntoAppStore, untargetedSupersedingAppsDetected, appRemovedBySupersedence, seeUninstallErrorCode, pendingReboot, installingDependencies, contentDownloaded, supersedingAppsNotApplicable, powerShellScriptRequirementNotMet, registryRequirementNotMet, fileSystemRequirementNotMet, platformNotApplicable, minimumCpuSpeedNotMet, minimumLogicalProcessorCountNotMet, minimumPhysicalMemoryNotMet, minimumOsVersionNotMet, minimumDiskSpaceNotMet, processorArchitectureNotApplicable.
    installStateDetail *ResultantAppStateDetail;
    // The collection of source mobile app's ids.
    sourceIds []string;
    // The related target app's display name.
    targetDisplayName *string;
    // The related target app's id.
    targetId *string;
    // The last sync time of the target app.
    targetLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewMobileAppRelationshipState instantiates a new mobileAppRelationshipState and sets the default values.
func NewMobileAppRelationshipState()(*MobileAppRelationshipState) {
    m := &MobileAppRelationshipState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMobileAppRelationshipStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppRelationshipStateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMobileAppRelationshipState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppRelationshipState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceId gets the deviceId property value. The corresponding device id.
func (m *MobileAppRelationshipState) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetErrorCode gets the errorCode property value. The error code for install or uninstall failures of target app.
func (m *MobileAppRelationshipState) GetErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppRelationshipState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["installState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallState(val.(*ResultantAppState))
        }
        return nil
    }
    res["installStateDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppStateDetail)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallStateDetail(val.(*ResultantAppStateDetail))
        }
        return nil
    }
    res["sourceIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSourceIds(res)
        }
        return nil
    }
    res["targetDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDisplayName(val)
        }
        return nil
    }
    res["targetId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetId(val)
        }
        return nil
    }
    res["targetLastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetLastSyncDateTime(val)
        }
        return nil
    }
    return res
}
// GetInstallState gets the installState property value. The install state of the app of target app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
func (m *MobileAppRelationshipState) GetInstallState()(*ResultantAppState) {
    if m == nil {
        return nil
    } else {
        return m.installState
    }
}
// GetInstallStateDetail gets the installStateDetail property value. The install state detail of the app. Possible values are: noAdditionalDetails, dependencyFailedToInstall, dependencyWithRequirementsNotMet, dependencyPendingReboot, dependencyWithAutoInstallDisabled, supersededAppUninstallFailed, supersededAppUninstallPendingReboot, removingSupersededApps, iosAppStoreUpdateFailedToInstall, vppAppHasUpdateAvailable, userRejectedUpdate, uninstallPendingReboot, supersedingAppsDetected, supersededAppsDetected, seeInstallErrorCode, autoInstallDisabled, managedAppNoLongerPresent, userRejectedInstall, userIsNotLoggedIntoAppStore, untargetedSupersedingAppsDetected, appRemovedBySupersedence, seeUninstallErrorCode, pendingReboot, installingDependencies, contentDownloaded, supersedingAppsNotApplicable, powerShellScriptRequirementNotMet, registryRequirementNotMet, fileSystemRequirementNotMet, platformNotApplicable, minimumCpuSpeedNotMet, minimumLogicalProcessorCountNotMet, minimumPhysicalMemoryNotMet, minimumOsVersionNotMet, minimumDiskSpaceNotMet, processorArchitectureNotApplicable.
func (m *MobileAppRelationshipState) GetInstallStateDetail()(*ResultantAppStateDetail) {
    if m == nil {
        return nil
    } else {
        return m.installStateDetail
    }
}
// GetSourceIds gets the sourceIds property value. The collection of source mobile app's ids.
func (m *MobileAppRelationshipState) GetSourceIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sourceIds
    }
}
// GetTargetDisplayName gets the targetDisplayName property value. The related target app's display name.
func (m *MobileAppRelationshipState) GetTargetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetDisplayName
    }
}
// GetTargetId gets the targetId property value. The related target app's id.
func (m *MobileAppRelationshipState) GetTargetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetId
    }
}
// GetTargetLastSyncDateTime gets the targetLastSyncDateTime property value. The last sync time of the target app.
func (m *MobileAppRelationshipState) GetTargetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.targetLastSyncDateTime
    }
}
// Serialize serializes information the current object
func (m *MobileAppRelationshipState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    if m.GetInstallState() != nil {
        cast := (*m.GetInstallState()).String()
        err := writer.WriteStringValue("installState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInstallStateDetail() != nil {
        cast := (*m.GetInstallStateDetail()).String()
        err := writer.WriteStringValue("installStateDetail", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSourceIds() != nil {
        err := writer.WriteCollectionOfStringValues("sourceIds", m.GetSourceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetDisplayName", m.GetTargetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetId", m.GetTargetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("targetLastSyncDateTime", m.GetTargetLastSyncDateTime())
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
func (m *MobileAppRelationshipState) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceId sets the deviceId property value. The corresponding device id.
func (m *MobileAppRelationshipState) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetErrorCode sets the errorCode property value. The error code for install or uninstall failures of target app.
func (m *MobileAppRelationshipState) SetErrorCode(value *int32)() {
    if m != nil {
        m.errorCode = value
    }
}
// SetInstallState sets the installState property value. The install state of the app of target app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
func (m *MobileAppRelationshipState) SetInstallState(value *ResultantAppState)() {
    if m != nil {
        m.installState = value
    }
}
// SetInstallStateDetail sets the installStateDetail property value. The install state detail of the app. Possible values are: noAdditionalDetails, dependencyFailedToInstall, dependencyWithRequirementsNotMet, dependencyPendingReboot, dependencyWithAutoInstallDisabled, supersededAppUninstallFailed, supersededAppUninstallPendingReboot, removingSupersededApps, iosAppStoreUpdateFailedToInstall, vppAppHasUpdateAvailable, userRejectedUpdate, uninstallPendingReboot, supersedingAppsDetected, supersededAppsDetected, seeInstallErrorCode, autoInstallDisabled, managedAppNoLongerPresent, userRejectedInstall, userIsNotLoggedIntoAppStore, untargetedSupersedingAppsDetected, appRemovedBySupersedence, seeUninstallErrorCode, pendingReboot, installingDependencies, contentDownloaded, supersedingAppsNotApplicable, powerShellScriptRequirementNotMet, registryRequirementNotMet, fileSystemRequirementNotMet, platformNotApplicable, minimumCpuSpeedNotMet, minimumLogicalProcessorCountNotMet, minimumPhysicalMemoryNotMet, minimumOsVersionNotMet, minimumDiskSpaceNotMet, processorArchitectureNotApplicable.
func (m *MobileAppRelationshipState) SetInstallStateDetail(value *ResultantAppStateDetail)() {
    if m != nil {
        m.installStateDetail = value
    }
}
// SetSourceIds sets the sourceIds property value. The collection of source mobile app's ids.
func (m *MobileAppRelationshipState) SetSourceIds(value []string)() {
    if m != nil {
        m.sourceIds = value
    }
}
// SetTargetDisplayName sets the targetDisplayName property value. The related target app's display name.
func (m *MobileAppRelationshipState) SetTargetDisplayName(value *string)() {
    if m != nil {
        m.targetDisplayName = value
    }
}
// SetTargetId sets the targetId property value. The related target app's id.
func (m *MobileAppRelationshipState) SetTargetId(value *string)() {
    if m != nil {
        m.targetId = value
    }
}
// SetTargetLastSyncDateTime sets the targetLastSyncDateTime property value. The last sync time of the target app.
func (m *MobileAppRelationshipState) SetTargetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.targetLastSyncDateTime = value
    }
}
