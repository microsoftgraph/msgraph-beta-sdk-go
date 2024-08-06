package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
type ManagedDevicesManagedDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters the list of managed devices.
type ManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// ManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateDeviceEsim provides operations to call the activateDeviceEsim method.
// returns a *ManagedDevicesItemActivateDeviceEsimRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*ManagedDevicesItemActivateDeviceEsimRequestBuilder) {
    return NewManagedDevicesItemActivateDeviceEsimRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
// returns a *ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
// returns a *ManagedDevicesItemBypassActivationLockRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*ManagedDevicesItemBypassActivationLockRequestBuilder) {
    return NewManagedDevicesItemBypassActivationLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChangeAssignments provides operations to call the changeAssignments method.
// returns a *ManagedDevicesItemChangeAssignmentsRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ChangeAssignments()(*ManagedDevicesItemChangeAssignmentsRequestBuilder) {
    return NewManagedDevicesItemChangeAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
// returns a *ManagedDevicesItemCleanWindowsDeviceRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ManagedDevicesItemCleanWindowsDeviceRequestBuilder) {
    return NewManagedDevicesItemCleanWindowsDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesManagedDeviceItemRequestBuilder) {
    m := &ManagedDevicesManagedDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
// returns a *ManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*ManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property managedDevices for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Deprovision provides operations to call the deprovision method.
// returns a *ManagedDevicesItemDeprovisionRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Deprovision()(*ManagedDevicesItemDeprovisionRequestBuilder) {
    return NewManagedDevicesItemDeprovisionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
// returns a *ManagedDevicesItemDetectedAppsRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DetectedApps()(*ManagedDevicesItemDetectedAppsRequestBuilder) {
    return NewManagedDevicesItemDetectedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
// returns a *ManagedDevicesItemDeviceCategoryRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ManagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewManagedDevicesItemDeviceCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
// returns a *ManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ManagedDevicesItemDeviceConfigurationStatesRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ManagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewManagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceHealthScriptStates provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
// returns a *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceHealthScriptStates()(*ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) {
    return NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Disable provides operations to call the disable method.
// returns a *ManagedDevicesItemDisableRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Disable()(*ManagedDevicesItemDisableRequestBuilder) {
    return NewManagedDevicesItemDisableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DisableLostMode provides operations to call the disableLostMode method.
// returns a *ManagedDevicesItemDisableLostModeRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*ManagedDevicesItemDisableLostModeRequestBuilder) {
    return NewManagedDevicesItemDisableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableLostMode provides operations to call the enableLostMode method.
// returns a *ManagedDevicesItemEnableLostModeRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*ManagedDevicesItemEnableLostModeRequestBuilder) {
    return NewManagedDevicesItemEnableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnrollNowAction provides operations to call the enrollNowAction method.
// returns a *ManagedDevicesItemEnrollNowActionRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*ManagedDevicesItemEnrollNowActionRequestBuilder) {
    return NewManagedDevicesItemEnrollNowActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of managed devices.
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
// returns a *ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
// returns a *ManagedDevicesItemGetCloudPcReviewStatusRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*ManagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    return NewManagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
// returns a *ManagedDevicesItemGetFileVaultKeyRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*ManagedDevicesItemGetFileVaultKeyRequestBuilder) {
    return NewManagedDevicesItemGetFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
// returns a *ManagedDevicesItemGetNonCompliantSettingsRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*ManagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    return NewManagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateDeviceAttestation provides operations to call the initiateDeviceAttestation method.
// returns a *ManagedDevicesItemInitiateDeviceAttestationRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) InitiateDeviceAttestation()(*ManagedDevicesItemInitiateDeviceAttestationRequestBuilder) {
    return NewManagedDevicesItemInitiateDeviceAttestationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
// returns a *ManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*ManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateOnDemandProactiveRemediation provides operations to call the initiateOnDemandProactiveRemediation method.
// returns a *ManagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) InitiateOnDemandProactiveRemediation()(*ManagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewManagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LocateDevice provides operations to call the locateDevice method.
// returns a *ManagedDevicesItemLocateDeviceRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) LocateDevice()(*ManagedDevicesItemLocateDeviceRequestBuilder) {
    return NewManagedDevicesItemLocateDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
// returns a *ManagedDevicesItemLogCollectionRequestsRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ManagedDevicesItemLogCollectionRequestsRequestBuilder) {
    return NewManagedDevicesItemLogCollectionRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
// returns a *ManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*ManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
// returns a *ManagedDevicesItemOverrideComplianceStateRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*ManagedDevicesItemOverrideComplianceStateRequestBuilder) {
    return NewManagedDevicesItemOverrideComplianceStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property managedDevices in deviceManagement
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
// returns a *ManagedDevicesItemPauseConfigurationRefreshRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) PauseConfigurationRefresh()(*ManagedDevicesItemPauseConfigurationRefreshRequestBuilder) {
    return NewManagedDevicesItemPauseConfigurationRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PlayLostModeSound provides operations to call the playLostModeSound method.
// returns a *ManagedDevicesItemPlayLostModeSoundRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*ManagedDevicesItemPlayLostModeSoundRequestBuilder) {
    return NewManagedDevicesItemPlayLostModeSoundRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RebootNow provides operations to call the rebootNow method.
// returns a *ManagedDevicesItemRebootNowRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) RebootNow()(*ManagedDevicesItemRebootNowRequestBuilder) {
    return NewManagedDevicesItemRebootNowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecoverPasscode provides operations to call the recoverPasscode method.
// returns a *ManagedDevicesItemRecoverPasscodeRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*ManagedDevicesItemRecoverPasscodeRequestBuilder) {
    return NewManagedDevicesItemRecoverPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reenable provides operations to call the reenable method.
// returns a *ManagedDevicesItemReenableRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Reenable()(*ManagedDevicesItemReenableRequestBuilder) {
    return NewManagedDevicesItemReenableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteLock provides operations to call the remoteLock method.
// returns a *ManagedDevicesItemRemoteLockRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) RemoteLock()(*ManagedDevicesItemRemoteLockRequestBuilder) {
    return NewManagedDevicesItemRemoteLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
// returns a *ManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*ManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
// returns a *ManagedDevicesItemReprovisionCloudPcRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*ManagedDevicesItemReprovisionCloudPcRequestBuilder) {
    return NewManagedDevicesItemReprovisionCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
// returns a *ManagedDevicesItemRequestRemoteAssistanceRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ManagedDevicesItemRequestRemoteAssistanceRequestBuilder) {
    return NewManagedDevicesItemRequestRemoteAssistanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetPasscode provides operations to call the resetPasscode method.
// returns a *ManagedDevicesItemResetPasscodeRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*ManagedDevicesItemResetPasscodeRequestBuilder) {
    return NewManagedDevicesItemResetPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResizeCloudPc provides operations to call the resizeCloudPc method.
// returns a *ManagedDevicesItemResizeCloudPcRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ResizeCloudPc()(*ManagedDevicesItemResizeCloudPcRequestBuilder) {
    return NewManagedDevicesItemResizeCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RestoreCloudPc provides operations to call the restoreCloudPc method.
// returns a *ManagedDevicesItemRestoreCloudPcRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) RestoreCloudPc()(*ManagedDevicesItemRestoreCloudPcRequestBuilder) {
    return NewManagedDevicesItemRestoreCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Retire provides operations to call the retire method.
// returns a *ManagedDevicesItemRetireRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Retire()(*ManagedDevicesItemRetireRequestBuilder) {
    return NewManagedDevicesItemRetireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
// returns a *ManagedDevicesItemRevokeAppleVppLicensesRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*ManagedDevicesItemRevokeAppleVppLicensesRequestBuilder) {
    return NewManagedDevicesItemRevokeAppleVppLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
// returns a *ManagedDevicesItemRotateBitLockerKeysRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*ManagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    return NewManagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
// returns a *ManagedDevicesItemRotateFileVaultKeyRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*ManagedDevicesItemRotateFileVaultKeyRequestBuilder) {
    return NewManagedDevicesItemRotateFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateLocalAdminPassword provides operations to call the rotateLocalAdminPassword method.
// returns a *ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) RotateLocalAdminPassword()(*ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder) {
    return NewManagedDevicesItemRotateLocalAdminPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
// returns a *ManagedDevicesItemSecurityBaselineStatesRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*ManagedDevicesItemSecurityBaselineStatesRequestBuilder) {
    return NewManagedDevicesItemSecurityBaselineStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
// returns a *ManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*ManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
// returns a *ManagedDevicesItemSetCloudPcReviewStatusRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*ManagedDevicesItemSetCloudPcReviewStatusRequestBuilder) {
    return NewManagedDevicesItemSetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetDeviceName provides operations to call the setDeviceName method.
// returns a *ManagedDevicesItemSetDeviceNameRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*ManagedDevicesItemSetDeviceNameRequestBuilder) {
    return NewManagedDevicesItemSetDeviceNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShutDown provides operations to call the shutDown method.
// returns a *ManagedDevicesItemShutDownRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ShutDown()(*ManagedDevicesItemShutDownRequestBuilder) {
    return NewManagedDevicesItemShutDownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncDevice provides operations to call the syncDevice method.
// returns a *ManagedDevicesItemSyncDeviceRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) SyncDevice()(*ManagedDevicesItemSyncDeviceRequestBuilder) {
    return NewManagedDevicesItemSyncDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property managedDevices for deviceManagement
// returns a *RequestInformation when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedDevicesItemTriggerConfigurationManagerActionRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*ManagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    return NewManagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
// returns a *ManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*ManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder) {
    return NewManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
// returns a *ManagedDevicesItemUsersRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Users()(*ManagedDevicesItemUsersRequestBuilder) {
    return NewManagedDevicesItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
// returns a *ManagedDevicesItemWindowsDefenderScanRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ManagedDevicesItemWindowsDefenderScanRequestBuilder) {
    return NewManagedDevicesItemWindowsDefenderScanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
// returns a *ManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*ManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
// returns a *ManagedDevicesItemWindowsProtectionStateRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ManagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewManagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Wipe provides operations to call the wipe method.
// returns a *ManagedDevicesItemWipeRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Wipe()(*ManagedDevicesItemWipeRequestBuilder) {
    return NewManagedDevicesItemWipeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManagedDevicesManagedDeviceItemRequestBuilder when successful
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesManagedDeviceItemRequestBuilder) {
    return NewManagedDevicesManagedDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
