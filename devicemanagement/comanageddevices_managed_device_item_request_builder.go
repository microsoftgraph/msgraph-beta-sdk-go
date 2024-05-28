package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanageddevicesManagedDeviceItemRequestBuilder provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
type ComanageddevicesManagedDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ComanageddevicesManagedDeviceItemRequestBuilderGetQueryParameters the list of co-managed devices report
type ComanageddevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ComanageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComanageddevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// ComanageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateDeviceEsim provides operations to call the activateDeviceEsim method.
// returns a *ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder) {
    return NewComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
// returns a *ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*ComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewComanageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
// returns a *ComanageddevicesItemBypassactivationlockBypassActivationLockRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*ComanageddevicesItemBypassactivationlockBypassActivationLockRequestBuilder) {
    return NewComanageddevicesItemBypassactivationlockBypassActivationLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
// returns a *ComanageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ComanageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) {
    return NewComanageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewComanageddevicesManagedDeviceItemRequestBuilderInternal instantiates a new ComanageddevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewComanageddevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesManagedDeviceItemRequestBuilder) {
    m := &ComanageddevicesManagedDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewComanageddevicesManagedDeviceItemRequestBuilder instantiates a new ComanageddevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewComanageddevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanageddevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
// returns a *ComanageddevicesItemCreatedevicelogcollectionrequestCreateDeviceLogCollectionRequestRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*ComanageddevicesItemCreatedevicelogcollectionrequestCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewComanageddevicesItemCreatedevicelogcollectionrequestCreateDeviceLogCollectionRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property comanagedDevices for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComanageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ComanageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ComanageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewComanageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Deprovision provides operations to call the deprovision method.
// returns a *ComanageddevicesItemDeprovisionRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) Deprovision()(*ComanageddevicesItemDeprovisionRequestBuilder) {
    return NewComanageddevicesItemDeprovisionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
// returns a *ComanageddevicesItemDetectedappsDetectedAppsRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) DetectedApps()(*ComanageddevicesItemDetectedappsDetectedAppsRequestBuilder) {
    return NewComanageddevicesItemDetectedappsDetectedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
// returns a *ComanageddevicesItemDevicecategoryDeviceCategoryRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ComanageddevicesItemDevicecategoryDeviceCategoryRequestBuilder) {
    return NewComanageddevicesItemDevicecategoryDeviceCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
// returns a *ComanageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ComanageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) {
    return NewComanageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilder) {
    return NewComanageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceHealthScriptStates provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
// returns a *ComanageddevicesItemDevicehealthscriptstatesDeviceHealthScriptStatesRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) DeviceHealthScriptStates()(*ComanageddevicesItemDevicehealthscriptstatesDeviceHealthScriptStatesRequestBuilder) {
    return NewComanageddevicesItemDevicehealthscriptstatesDeviceHealthScriptStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Disable provides operations to call the disable method.
// returns a *ComanageddevicesItemDisableRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) Disable()(*ComanageddevicesItemDisableRequestBuilder) {
    return NewComanageddevicesItemDisableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DisableLostMode provides operations to call the disableLostMode method.
// returns a *ComanageddevicesItemDisablelostmodeDisableLostModeRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*ComanageddevicesItemDisablelostmodeDisableLostModeRequestBuilder) {
    return NewComanageddevicesItemDisablelostmodeDisableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableLostMode provides operations to call the enableLostMode method.
// returns a *ComanageddevicesItemEnablelostmodeEnableLostModeRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*ComanageddevicesItemEnablelostmodeEnableLostModeRequestBuilder) {
    return NewComanageddevicesItemEnablelostmodeEnableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnrollNowAction provides operations to call the enrollNowAction method.
// returns a *ComanageddevicesItemEnrollnowactionEnrollNowActionRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*ComanageddevicesItemEnrollnowactionEnrollNowActionRequestBuilder) {
    return NewComanageddevicesItemEnrollnowactionEnrollNowActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of co-managed devices report
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
// returns a *ComanageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*ComanageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewComanageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
// returns a *ComanageddevicesItemGetcloudpcreviewstatusGetCloudPcReviewStatusRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*ComanageddevicesItemGetcloudpcreviewstatusGetCloudPcReviewStatusRequestBuilder) {
    return NewComanageddevicesItemGetcloudpcreviewstatusGetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
// returns a *ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) {
    return NewComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
// returns a *ComanageddevicesItemGetnoncompliantsettingsGetNonCompliantSettingsRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*ComanageddevicesItemGetnoncompliantsettingsGetNonCompliantSettingsRequestBuilder) {
    return NewComanageddevicesItemGetnoncompliantsettingsGetNonCompliantSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateDeviceAttestation provides operations to call the initiateDeviceAttestation method.
// returns a *ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) InitiateDeviceAttestation()(*ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder) {
    return NewComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
// returns a *ComanageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*ComanageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewComanageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateOnDemandProactiveRemediation provides operations to call the initiateOnDemandProactiveRemediation method.
// returns a *ComanageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) InitiateOnDemandProactiveRemediation()(*ComanageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewComanageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LocateDevice provides operations to call the locateDevice method.
// returns a *ComanageddevicesItemLocatedeviceLocateDeviceRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) LocateDevice()(*ComanageddevicesItemLocatedeviceLocateDeviceRequestBuilder) {
    return NewComanageddevicesItemLocatedeviceLocateDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
// returns a *ComanageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ComanageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) {
    return NewComanageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
// returns a *ComanageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ComanageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewComanageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ComanageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*ComanageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewComanageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
// returns a *ComanageddevicesItemOverridecompliancestateOverrideComplianceStateRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*ComanageddevicesItemOverridecompliancestateOverrideComplianceStateRequestBuilder) {
    return NewComanageddevicesItemOverridecompliancestateOverrideComplianceStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property comanagedDevices in deviceManagement
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ComanageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
// returns a *ComanageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) PauseConfigurationRefresh()(*ComanageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshRequestBuilder) {
    return NewComanageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PlayLostModeSound provides operations to call the playLostModeSound method.
// returns a *ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*ComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder) {
    return NewComanageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RebootNow provides operations to call the rebootNow method.
// returns a *ComanageddevicesItemRebootnowRebootNowRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) RebootNow()(*ComanageddevicesItemRebootnowRebootNowRequestBuilder) {
    return NewComanageddevicesItemRebootnowRebootNowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecoverPasscode provides operations to call the recoverPasscode method.
// returns a *ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*ComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) {
    return NewComanageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reenable provides operations to call the reenable method.
// returns a *ComanageddevicesItemReenableRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) Reenable()(*ComanageddevicesItemReenableRequestBuilder) {
    return NewComanageddevicesItemReenableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteLock provides operations to call the remoteLock method.
// returns a *ComanageddevicesItemRemotelockRemoteLockRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) RemoteLock()(*ComanageddevicesItemRemotelockRemoteLockRequestBuilder) {
    return NewComanageddevicesItemRemotelockRemoteLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
// returns a *ComanageddevicesItemRemovedevicefirmwareconfigurationinterfacemanagementRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*ComanageddevicesItemRemovedevicefirmwareconfigurationinterfacemanagementRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewComanageddevicesItemRemovedevicefirmwareconfigurationinterfacemanagementRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
// returns a *ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder) {
    return NewComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
// returns a *ComanageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ComanageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) {
    return NewComanageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetPasscode provides operations to call the resetPasscode method.
// returns a *ComanageddevicesItemResetpasscodeResetPasscodeRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*ComanageddevicesItemResetpasscodeResetPasscodeRequestBuilder) {
    return NewComanageddevicesItemResetpasscodeResetPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResizeCloudPc provides operations to call the resizeCloudPc method.
// returns a *ComanageddevicesItemResizecloudpcResizeCloudPcRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) ResizeCloudPc()(*ComanageddevicesItemResizecloudpcResizeCloudPcRequestBuilder) {
    return NewComanageddevicesItemResizecloudpcResizeCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RestoreCloudPc provides operations to call the restoreCloudPc method.
// returns a *ComanageddevicesItemRestorecloudpcRestoreCloudPcRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) RestoreCloudPc()(*ComanageddevicesItemRestorecloudpcRestoreCloudPcRequestBuilder) {
    return NewComanageddevicesItemRestorecloudpcRestoreCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Retire provides operations to call the retire method.
// returns a *ComanageddevicesItemRetireRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) Retire()(*ComanageddevicesItemRetireRequestBuilder) {
    return NewComanageddevicesItemRetireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
// returns a *ComanageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*ComanageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder) {
    return NewComanageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
// returns a *ComanageddevicesItemRotatebitlockerkeysRotateBitLockerKeysRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*ComanageddevicesItemRotatebitlockerkeysRotateBitLockerKeysRequestBuilder) {
    return NewComanageddevicesItemRotatebitlockerkeysRotateBitLockerKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
// returns a *ComanageddevicesItemRotatefilevaultkeyRotateFileVaultKeyRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*ComanageddevicesItemRotatefilevaultkeyRotateFileVaultKeyRequestBuilder) {
    return NewComanageddevicesItemRotatefilevaultkeyRotateFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateLocalAdminPassword provides operations to call the rotateLocalAdminPassword method.
// returns a *ComanageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) RotateLocalAdminPassword()(*ComanageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder) {
    return NewComanageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
// returns a *ComanageddevicesItemSecuritybaselinestatesSecurityBaselineStatesRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*ComanageddevicesItemSecuritybaselinestatesSecurityBaselineStatesRequestBuilder) {
    return NewComanageddevicesItemSecuritybaselinestatesSecurityBaselineStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
// returns a *ComanageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*ComanageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewComanageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
// returns a *ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*ComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder) {
    return NewComanageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetDeviceName provides operations to call the setDeviceName method.
// returns a *ComanageddevicesItemSetdevicenameSetDeviceNameRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*ComanageddevicesItemSetdevicenameSetDeviceNameRequestBuilder) {
    return NewComanageddevicesItemSetdevicenameSetDeviceNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShutDown provides operations to call the shutDown method.
// returns a *ComanageddevicesItemShutdownShutDownRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) ShutDown()(*ComanageddevicesItemShutdownShutDownRequestBuilder) {
    return NewComanageddevicesItemShutdownShutDownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncDevice provides operations to call the syncDevice method.
// returns a *ComanageddevicesItemSyncdeviceSyncDeviceRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) SyncDevice()(*ComanageddevicesItemSyncdeviceSyncDeviceRequestBuilder) {
    return NewComanageddevicesItemSyncdeviceSyncDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property comanagedDevices for deviceManagement
// returns a *RequestInformation when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of co-managed devices report
// returns a *RequestInformation when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property comanagedDevices in deviceManagement
// returns a *RequestInformation when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ComanageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*ComanageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder) {
    return NewComanageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
// returns a *ComanageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*ComanageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder) {
    return NewComanageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
// returns a *ComanageddevicesItemUsersRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) Users()(*ComanageddevicesItemUsersRequestBuilder) {
    return NewComanageddevicesItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
// returns a *ComanageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ComanageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilder) {
    return NewComanageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
// returns a *ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
// returns a *ComanageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ComanageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) {
    return NewComanageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Wipe provides operations to call the wipe method.
// returns a *ComanageddevicesItemWipeRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) Wipe()(*ComanageddevicesItemWipeRequestBuilder) {
    return NewComanageddevicesItemWipeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ComanageddevicesManagedDeviceItemRequestBuilder when successful
func (m *ComanageddevicesManagedDeviceItemRequestBuilder) WithUrl(rawUrl string)(*ComanageddevicesManagedDeviceItemRequestBuilder) {
    return NewComanageddevicesManagedDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
