package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ItemManageddevicesManagedDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemManageddevicesManagedDeviceItemRequestBuilderGetQueryParameters the managed devices associated with the user.
type ItemManageddevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManageddevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// ItemManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateDeviceEsim provides operations to call the activateDeviceEsim method.
// returns a *ItemManageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*ItemManageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder) {
    return NewItemManageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*ItemManageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewItemManageddevicesItemAssignmentfilterevaluationstatusdetailsAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
// returns a *ItemManageddevicesItemBypassactivationlockBypassActivationLockRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*ItemManageddevicesItemBypassactivationlockBypassActivationLockRequestBuilder) {
    return NewItemManageddevicesItemBypassactivationlockBypassActivationLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
// returns a *ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) {
    return NewItemManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemManageddevicesManagedDeviceItemRequestBuilderInternal instantiates a new ItemManageddevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewItemManageddevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesManagedDeviceItemRequestBuilder) {
    m := &ItemManageddevicesManagedDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemManageddevicesManagedDeviceItemRequestBuilder instantiates a new ItemManageddevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewItemManageddevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
// returns a *ItemManageddevicesItemCreatedevicelogcollectionrequestCreateDeviceLogCollectionRequestRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*ItemManageddevicesItemCreatedevicelogcollectionrequestCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewItemManageddevicesItemCreatedevicelogcollectionrequestCreateDeviceLogCollectionRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property managedDevices for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ItemManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ItemManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewItemManageddevicesItemDeleteuserfromsharedappledeviceDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Deprovision provides operations to call the deprovision method.
// returns a *ItemManageddevicesItemDeprovisionRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Deprovision()(*ItemManageddevicesItemDeprovisionRequestBuilder) {
    return NewItemManageddevicesItemDeprovisionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemDetectedappsDetectedAppsRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) DetectedApps()(*ItemManageddevicesItemDetectedappsDetectedAppsRequestBuilder) {
    return NewItemManageddevicesItemDetectedappsDetectedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemDevicecategoryDeviceCategoryRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ItemManageddevicesItemDevicecategoryDeviceCategoryRequestBuilder) {
    return NewItemManageddevicesItemDevicecategoryDeviceCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ItemManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilder) {
    return NewItemManageddevicesItemDevicecompliancepolicystatesDeviceCompliancePolicyStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilder) {
    return NewItemManageddevicesItemDeviceconfigurationstatesDeviceConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceHealthScriptStates provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemDevicehealthscriptstatesDeviceHealthScriptStatesRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) DeviceHealthScriptStates()(*ItemManageddevicesItemDevicehealthscriptstatesDeviceHealthScriptStatesRequestBuilder) {
    return NewItemManageddevicesItemDevicehealthscriptstatesDeviceHealthScriptStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Disable provides operations to call the disable method.
// returns a *ItemManageddevicesItemDisableRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Disable()(*ItemManageddevicesItemDisableRequestBuilder) {
    return NewItemManageddevicesItemDisableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DisableLostMode provides operations to call the disableLostMode method.
// returns a *ItemManageddevicesItemDisablelostmodeDisableLostModeRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*ItemManageddevicesItemDisablelostmodeDisableLostModeRequestBuilder) {
    return NewItemManageddevicesItemDisablelostmodeDisableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableLostMode provides operations to call the enableLostMode method.
// returns a *ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder) {
    return NewItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnrollNowAction provides operations to call the enrollNowAction method.
// returns a *ItemManageddevicesItemEnrollnowactionEnrollNowActionRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*ItemManageddevicesItemEnrollnowactionEnrollNowActionRequestBuilder) {
    return NewItemManageddevicesItemEnrollnowactionEnrollNowActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the managed devices associated with the user.
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
// returns a *ItemManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*ItemManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewItemManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
// returns a *ItemManageddevicesItemGetcloudpcreviewstatusGetCloudPcReviewStatusRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*ItemManageddevicesItemGetcloudpcreviewstatusGetCloudPcReviewStatusRequestBuilder) {
    return NewItemManageddevicesItemGetcloudpcreviewstatusGetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
// returns a *ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) {
    return NewItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
// returns a *ItemManageddevicesItemGetnoncompliantsettingsGetNonCompliantSettingsRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*ItemManageddevicesItemGetnoncompliantsettingsGetNonCompliantSettingsRequestBuilder) {
    return NewItemManageddevicesItemGetnoncompliantsettingsGetNonCompliantSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateDeviceAttestation provides operations to call the initiateDeviceAttestation method.
// returns a *ItemManageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) InitiateDeviceAttestation()(*ItemManageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder) {
    return NewItemManageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
// returns a *ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateOnDemandProactiveRemediation provides operations to call the initiateOnDemandProactiveRemediation method.
// returns a *ItemManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) InitiateOnDemandProactiveRemediation()(*ItemManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewItemManageddevicesItemInitiateondemandproactiveremediationInitiateOnDemandProactiveRemediationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LocateDevice provides operations to call the locateDevice method.
// returns a *ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) LocateDevice()(*ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) {
    return NewItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) {
    return NewItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
// returns a *ItemManageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ItemManageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewItemManageddevicesItemLogoutsharedappledeviceactiveuserLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
// returns a *ItemManageddevicesItemOverridecompliancestateOverrideComplianceStateRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*ItemManageddevicesItemOverridecompliancestateOverrideComplianceStateRequestBuilder) {
    return NewItemManageddevicesItemOverridecompliancestateOverrideComplianceStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property managedDevices in users
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ItemManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
// returns a *ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) PauseConfigurationRefresh()(*ItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshRequestBuilder) {
    return NewItemManageddevicesItemPauseconfigurationrefreshPauseConfigurationRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PlayLostModeSound provides operations to call the playLostModeSound method.
// returns a *ItemManageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*ItemManageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilder) {
    return NewItemManageddevicesItemPlaylostmodesoundPlayLostModeSoundRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RebootNow provides operations to call the rebootNow method.
// returns a *ItemManageddevicesItemRebootnowRebootNowRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RebootNow()(*ItemManageddevicesItemRebootnowRebootNowRequestBuilder) {
    return NewItemManageddevicesItemRebootnowRebootNowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecoverPasscode provides operations to call the recoverPasscode method.
// returns a *ItemManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*ItemManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilder) {
    return NewItemManageddevicesItemRecoverpasscodeRecoverPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reenable provides operations to call the reenable method.
// returns a *ItemManageddevicesItemReenableRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Reenable()(*ItemManageddevicesItemReenableRequestBuilder) {
    return NewItemManageddevicesItemReenableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteLock provides operations to call the remoteLock method.
// returns a *ItemManageddevicesItemRemotelockRemoteLockRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RemoteLock()(*ItemManageddevicesItemRemotelockRemoteLockRequestBuilder) {
    return NewItemManageddevicesItemRemotelockRemoteLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
// returns a *ItemManageddevicesItemRemovedevicefirmwareconfigurationinterfacemanagementRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*ItemManageddevicesItemRemovedevicefirmwareconfigurationinterfacemanagementRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewItemManageddevicesItemRemovedevicefirmwareconfigurationinterfacemanagementRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
// returns a *ItemManageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*ItemManageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder) {
    return NewItemManageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
// returns a *ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) {
    return NewItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetPasscode provides operations to call the resetPasscode method.
// returns a *ItemManageddevicesItemResetpasscodeResetPasscodeRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*ItemManageddevicesItemResetpasscodeResetPasscodeRequestBuilder) {
    return NewItemManageddevicesItemResetpasscodeResetPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResizeCloudPc provides operations to call the resizeCloudPc method.
// returns a *ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ResizeCloudPc()(*ItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilder) {
    return NewItemManageddevicesItemResizecloudpcResizeCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RestoreCloudPc provides operations to call the restoreCloudPc method.
// returns a *ItemManageddevicesItemRestorecloudpcRestoreCloudPcRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RestoreCloudPc()(*ItemManageddevicesItemRestorecloudpcRestoreCloudPcRequestBuilder) {
    return NewItemManageddevicesItemRestorecloudpcRestoreCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Retire provides operations to call the retire method.
// returns a *ItemManageddevicesItemRetireRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Retire()(*ItemManageddevicesItemRetireRequestBuilder) {
    return NewItemManageddevicesItemRetireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
// returns a *ItemManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*ItemManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilder) {
    return NewItemManageddevicesItemRevokeapplevpplicensesRevokeAppleVppLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
// returns a *ItemManageddevicesItemRotatebitlockerkeysRotateBitLockerKeysRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*ItemManageddevicesItemRotatebitlockerkeysRotateBitLockerKeysRequestBuilder) {
    return NewItemManageddevicesItemRotatebitlockerkeysRotateBitLockerKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
// returns a *ItemManageddevicesItemRotatefilevaultkeyRotateFileVaultKeyRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*ItemManageddevicesItemRotatefilevaultkeyRotateFileVaultKeyRequestBuilder) {
    return NewItemManageddevicesItemRotatefilevaultkeyRotateFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateLocalAdminPassword provides operations to call the rotateLocalAdminPassword method.
// returns a *ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) RotateLocalAdminPassword()(*ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder) {
    return NewItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemSecuritybaselinestatesSecurityBaselineStatesRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*ItemManageddevicesItemSecuritybaselinestatesSecurityBaselineStatesRequestBuilder) {
    return NewItemManageddevicesItemSecuritybaselinestatesSecurityBaselineStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
// returns a *ItemManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*ItemManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewItemManageddevicesItemSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
// returns a *ItemManageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*ItemManageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilder) {
    return NewItemManageddevicesItemSetcloudpcreviewstatusSetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetDeviceName provides operations to call the setDeviceName method.
// returns a *ItemManageddevicesItemSetdevicenameSetDeviceNameRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*ItemManageddevicesItemSetdevicenameSetDeviceNameRequestBuilder) {
    return NewItemManageddevicesItemSetdevicenameSetDeviceNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShutDown provides operations to call the shutDown method.
// returns a *ItemManageddevicesItemShutdownShutDownRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ShutDown()(*ItemManageddevicesItemShutdownShutDownRequestBuilder) {
    return NewItemManageddevicesItemShutdownShutDownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncDevice provides operations to call the syncDevice method.
// returns a *ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) SyncDevice()(*ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder) {
    return NewItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property managedDevices for users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the managed devices associated with the user.
// returns a *RequestInformation when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedDevices in users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ItemManageddevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder) {
    return NewItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
// returns a *ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*ItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder) {
    return NewItemManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemUsersRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Users()(*ItemManageddevicesItemUsersRequestBuilder) {
    return NewItemManageddevicesItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
// returns a *ItemManageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ItemManageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilder) {
    return NewItemManageddevicesItemWindowsdefenderscanWindowsDefenderScanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
// returns a *ItemManageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*ItemManageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewItemManageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilder) {
    return NewItemManageddevicesItemWindowsprotectionstateWindowsProtectionStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Wipe provides operations to call the wipe method.
// returns a *ItemManageddevicesItemWipeRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) Wipe()(*ItemManageddevicesItemWipeRequestBuilder) {
    return NewItemManageddevicesItemWipeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManageddevicesManagedDeviceItemRequestBuilder when successful
func (m *ItemManageddevicesManagedDeviceItemRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesManagedDeviceItemRequestBuilder) {
    return NewItemManageddevicesManagedDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
