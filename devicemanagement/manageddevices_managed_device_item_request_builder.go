package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
type ManageddevicesManagedDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManageddevicesManagedDeviceItemRequestBuilderGetQueryParameters the list of managed devices.
type ManageddevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// ManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateDeviceEsim provides operations to call the activateDeviceEsim method.
// returns a *ManageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*ManageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder) {
    return NewManageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*ManageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewManageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
// returns a *ManageddevicesItemBypassactivationlockBypassActivationLockRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*ManageddevicesItemBypassactivationlockBypassActivationLockRequestBuilder) {
    return NewManageddevicesItemBypassactivationlockBypassActivationLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
// returns a *ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) {
    return NewManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewManageddevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManageddevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewManageddevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesManagedDeviceItemRequestBuilder) {
    m := &ManageddevicesManagedDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManageddevicesManagedDeviceItemRequestBuilder instantiates a new ManageddevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewManageddevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
// returns a *ManageddevicesItemCreatedevicelogcollectionrequestCreateDeviceLogCollectionRequestRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*ManageddevicesItemCreatedevicelogcollectionrequestCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewManageddevicesItemCreatedevicelogcollectionrequestCreateDeviceLogCollectionRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property managedDevices for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeleteUserFromSharedAppleDevice provides operations to call the deleteUserFromSharedAppleDevice method.
// returns a *ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Deprovision provides operations to call the deprovision method.
// returns a *ManageddevicesItemDeprovisionRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Deprovision()(*ManageddevicesItemDeprovisionRequestBuilder) {
    return NewManageddevicesItemDeprovisionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemDetectedappsDetectedAppsRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) DetectedApps()(*ManageddevicesItemDetectedappsDetectedAppsRequestBuilder) {
    return NewManageddevicesItemDetectedappsDetectedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemDevicecategoryDeviceCategoryRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ManageddevicesItemDevicecategoryDeviceCategoryRequestBuilder) {
    return NewManageddevicesItemDevicecategoryDeviceCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) {
    return NewManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilder) {
    return NewManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceHealthScriptStates provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemDevicehealthscriptstatesDeviceHealthScriptStatesRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) DeviceHealthScriptStates()(*ManageddevicesItemDevicehealthscriptstatesDeviceHealthScriptStatesRequestBuilder) {
    return NewManageddevicesItemDevicehealthscriptstatesDeviceHealthScriptStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Disable provides operations to call the disable method.
// returns a *ManageddevicesItemDisableRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Disable()(*ManageddevicesItemDisableRequestBuilder) {
    return NewManageddevicesItemDisableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DisableLostMode provides operations to call the disableLostMode method.
// returns a *ManageddevicesItemDisablelostmodeDisableLostModeRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*ManageddevicesItemDisablelostmodeDisableLostModeRequestBuilder) {
    return NewManageddevicesItemDisablelostmodeDisableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableLostMode provides operations to call the enableLostMode method.
// returns a *ManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*ManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder) {
    return NewManageddevicesItemEnablelostmodeEnableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnrollNowAction provides operations to call the enrollNowAction method.
// returns a *ManageddevicesItemEnrollnowactionEnrollNowActionRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*ManageddevicesItemEnrollnowactionEnrollNowActionRequestBuilder) {
    return NewManageddevicesItemEnrollnowactionEnrollNowActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of managed devices.
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
// GetCloudPcRemoteActionResults provides operations to call the getCloudPcRemoteActionResults method.
// returns a *ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
// returns a *ManageddevicesItemGetcloudpcreviewstatusGetCloudPcReviewStatusRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*ManageddevicesItemGetcloudpcreviewstatusGetCloudPcReviewStatusRequestBuilder) {
    return NewManageddevicesItemGetcloudpcreviewstatusGetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
// returns a *ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) {
    return NewManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
// returns a *ManageddevicesItemGetnoncompliantsettingsGetNonCompliantSettingsRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*ManageddevicesItemGetnoncompliantsettingsGetNonCompliantSettingsRequestBuilder) {
    return NewManageddevicesItemGetnoncompliantsettingsGetNonCompliantSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateDeviceAttestation provides operations to call the initiateDeviceAttestation method.
// returns a *ManageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) InitiateDeviceAttestation()(*ManageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder) {
    return NewManageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
// returns a *ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateOnDemandProactiveRemediation provides operations to call the initiateOnDemandProactiveRemediation method.
// returns a *ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) InitiateOnDemandProactiveRemediation()(*ManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LocateDevice provides operations to call the locateDevice method.
// returns a *ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) LocateDevice()(*ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) {
    return NewManageddevicesItemLocatedeviceLocateDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) {
    return NewManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
// returns a *ManageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ManageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewManageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*ManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
// returns a *ManageddevicesItemOverridecompliancestateOverrideComplianceStateRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*ManageddevicesItemOverridecompliancestateOverrideComplianceStateRequestBuilder) {
    return NewManageddevicesItemOverridecompliancestateOverrideComplianceStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property managedDevices in deviceManagement
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
// PauseConfigurationRefresh provides operations to call the pauseConfigurationRefresh method.
// returns a *ManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) PauseConfigurationRefresh()(*ManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshRequestBuilder) {
    return NewManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PlayLostModeSound provides operations to call the playLostModeSound method.
// returns a *ManageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*ManageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder) {
    return NewManageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RebootNow provides operations to call the rebootNow method.
// returns a *ManageddevicesItemRebootnowRebootNowRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RebootNow()(*ManageddevicesItemRebootnowRebootNowRequestBuilder) {
    return NewManageddevicesItemRebootnowRebootNowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecoverPasscode provides operations to call the recoverPasscode method.
// returns a *ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*ManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) {
    return NewManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reenable provides operations to call the reenable method.
// returns a *ManageddevicesItemReenableRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Reenable()(*ManageddevicesItemReenableRequestBuilder) {
    return NewManageddevicesItemReenableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteLock provides operations to call the remoteLock method.
// returns a *ManageddevicesItemRemotelockRemoteLockRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RemoteLock()(*ManageddevicesItemRemotelockRemoteLockRequestBuilder) {
    return NewManageddevicesItemRemotelockRemoteLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
// returns a *ManageddevicesItemRemovedevicefirmwareconfigurationinterfacemanagementRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*ManageddevicesItemRemovedevicefirmwareconfigurationinterfacemanagementRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewManageddevicesItemRemovedevicefirmwareconfigurationinterfacemanagementRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
// returns a *ManageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*ManageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder) {
    return NewManageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
// returns a *ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) {
    return NewManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetPasscode provides operations to call the resetPasscode method.
// returns a *ManageddevicesItemResetpasscodeResetPasscodeRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*ManageddevicesItemResetpasscodeResetPasscodeRequestBuilder) {
    return NewManageddevicesItemResetpasscodeResetPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResizeCloudPc provides operations to call the resizeCloudPc method.
// returns a *ManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ResizeCloudPc()(*ManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder) {
    return NewManageddevicesItemResizecloudpcResizeCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RestoreCloudPc provides operations to call the restoreCloudPc method.
// returns a *ManageddevicesItemRestorecloudpcRestoreCloudPcRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RestoreCloudPc()(*ManageddevicesItemRestorecloudpcRestoreCloudPcRequestBuilder) {
    return NewManageddevicesItemRestorecloudpcRestoreCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Retire provides operations to call the retire method.
// returns a *ManageddevicesItemRetireRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Retire()(*ManageddevicesItemRetireRequestBuilder) {
    return NewManageddevicesItemRetireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
// returns a *ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*ManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder) {
    return NewManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
// returns a *ManageddevicesItemRotatebitlockerkeysRotateBitLockerKeysRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*ManageddevicesItemRotatebitlockerkeysRotateBitLockerKeysRequestBuilder) {
    return NewManageddevicesItemRotatebitlockerkeysRotateBitLockerKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
// returns a *ManageddevicesItemRotatefilevaultkeyRotateFileVaultKeyRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*ManageddevicesItemRotatefilevaultkeyRotateFileVaultKeyRequestBuilder) {
    return NewManageddevicesItemRotatefilevaultkeyRotateFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateLocalAdminPassword provides operations to call the rotateLocalAdminPassword method.
// returns a *ManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) RotateLocalAdminPassword()(*ManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder) {
    return NewManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemSecuritybaselinestatesSecurityBaselineStatesRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*ManageddevicesItemSecuritybaselinestatesSecurityBaselineStatesRequestBuilder) {
    return NewManageddevicesItemSecuritybaselinestatesSecurityBaselineStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
// returns a *ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*ManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
// returns a *ManageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*ManageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder) {
    return NewManageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetDeviceName provides operations to call the setDeviceName method.
// returns a *ManageddevicesItemSetdevicenameSetDeviceNameRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*ManageddevicesItemSetdevicenameSetDeviceNameRequestBuilder) {
    return NewManageddevicesItemSetdevicenameSetDeviceNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShutDown provides operations to call the shutDown method.
// returns a *ManageddevicesItemShutdownShutDownRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ShutDown()(*ManageddevicesItemShutdownShutDownRequestBuilder) {
    return NewManageddevicesItemShutdownShutDownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncDevice provides operations to call the syncDevice method.
// returns a *ManageddevicesItemSyncdeviceSyncDeviceRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) SyncDevice()(*ManageddevicesItemSyncdeviceSyncDeviceRequestBuilder) {
    return NewManageddevicesItemSyncdeviceSyncDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property managedDevices for deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of managed devices.
// returns a *RequestInformation when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property managedDevices in deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// TriggerConfigurationManagerAction provides operations to call the triggerConfigurationManagerAction method.
// returns a *ManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*ManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder) {
    return NewManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
// returns a *ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder) {
    return NewManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemUsersRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Users()(*ManageddevicesItemUsersRequestBuilder) {
    return NewManageddevicesItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
// returns a *ManageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ManageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilder) {
    return NewManageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
// returns a *ManageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*ManageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewManageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
// returns a *ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) {
    return NewManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Wipe provides operations to call the wipe method.
// returns a *ManageddevicesItemWipeRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) Wipe()(*ManageddevicesItemWipeRequestBuilder) {
    return NewManageddevicesItemWipeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddevicesManagedDeviceItemRequestBuilder when successful
func (m *ManageddevicesManagedDeviceItemRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesManagedDeviceItemRequestBuilder) {
    return NewManageddevicesManagedDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
