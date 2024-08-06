package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
type ComanagedDevicesManagedDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ComanagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters the list of co-managed devices report
type ComanagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComanagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// ComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateDeviceEsim provides operations to call the activateDeviceEsim method.
// returns a *ComanagedDevicesItemActivateDeviceEsimRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*ComanagedDevicesItemActivateDeviceEsimRequestBuilder) {
    return NewComanagedDevicesItemActivateDeviceEsimRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
// returns a *ComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*ComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
// returns a *ComanagedDevicesItemBypassActivationLockRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*ComanagedDevicesItemBypassActivationLockRequestBuilder) {
    return NewComanagedDevicesItemBypassActivationLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChangeAssignments provides operations to call the changeAssignments method.
// returns a *ComanagedDevicesItemChangeAssignmentsRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ChangeAssignments()(*ComanagedDevicesItemChangeAssignmentsRequestBuilder) {
    return NewComanagedDevicesItemChangeAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
// returns a *ComanagedDevicesItemCleanWindowsDeviceRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ComanagedDevicesItemCleanWindowsDeviceRequestBuilder) {
    return NewComanagedDevicesItemCleanWindowsDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewComanagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ComanagedDevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewComanagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    m := &ComanagedDevicesManagedDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewComanagedDevicesManagedDeviceItemRequestBuilder instantiates a new ComanagedDevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewComanagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
// returns a *ComanagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*ComanagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewComanagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property comanagedDevices for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ComanagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ComanagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewComanagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Deprovision provides operations to call the deprovision method.
// returns a *ComanagedDevicesItemDeprovisionRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Deprovision()(*ComanagedDevicesItemDeprovisionRequestBuilder) {
    return NewComanagedDevicesItemDeprovisionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
// returns a *ComanagedDevicesItemDetectedAppsRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DetectedApps()(*ComanagedDevicesItemDetectedAppsRequestBuilder) {
    return NewComanagedDevicesItemDetectedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
// returns a *ComanagedDevicesItemDeviceCategoryRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ComanagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewComanagedDevicesItemDeviceCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
// returns a *ComanagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ComanagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewComanagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ComanagedDevicesItemDeviceConfigurationStatesRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ComanagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewComanagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceHealthScriptStates provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
// returns a *ComanagedDevicesItemDeviceHealthScriptStatesRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceHealthScriptStates()(*ComanagedDevicesItemDeviceHealthScriptStatesRequestBuilder) {
    return NewComanagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Disable provides operations to call the disable method.
// returns a *ComanagedDevicesItemDisableRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Disable()(*ComanagedDevicesItemDisableRequestBuilder) {
    return NewComanagedDevicesItemDisableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DisableLostMode provides operations to call the disableLostMode method.
// returns a *ComanagedDevicesItemDisableLostModeRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*ComanagedDevicesItemDisableLostModeRequestBuilder) {
    return NewComanagedDevicesItemDisableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableLostMode provides operations to call the enableLostMode method.
// returns a *ComanagedDevicesItemEnableLostModeRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*ComanagedDevicesItemEnableLostModeRequestBuilder) {
    return NewComanagedDevicesItemEnableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnrollNowAction provides operations to call the enrollNowAction method.
// returns a *ComanagedDevicesItemEnrollNowActionRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*ComanagedDevicesItemEnrollNowActionRequestBuilder) {
    return NewComanagedDevicesItemEnrollNowActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of co-managed devices report
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
// returns a *ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
// returns a *ComanagedDevicesItemGetCloudPcReviewStatusRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*ComanagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    return NewComanagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
// returns a *ComanagedDevicesItemGetFileVaultKeyRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*ComanagedDevicesItemGetFileVaultKeyRequestBuilder) {
    return NewComanagedDevicesItemGetFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
// returns a *ComanagedDevicesItemGetNonCompliantSettingsRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*ComanagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    return NewComanagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateDeviceAttestation provides operations to call the initiateDeviceAttestation method.
// returns a *ComanagedDevicesItemInitiateDeviceAttestationRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) InitiateDeviceAttestation()(*ComanagedDevicesItemInitiateDeviceAttestationRequestBuilder) {
    return NewComanagedDevicesItemInitiateDeviceAttestationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
// returns a *ComanagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*ComanagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewComanagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateOnDemandProactiveRemediation provides operations to call the initiateOnDemandProactiveRemediation method.
// returns a *ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) InitiateOnDemandProactiveRemediation()(*ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LocateDevice provides operations to call the locateDevice method.
// returns a *ComanagedDevicesItemLocateDeviceRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) LocateDevice()(*ComanagedDevicesItemLocateDeviceRequestBuilder) {
    return NewComanagedDevicesItemLocateDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
// returns a *ComanagedDevicesItemLogCollectionRequestsRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ComanagedDevicesItemLogCollectionRequestsRequestBuilder) {
    return NewComanagedDevicesItemLogCollectionRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
// returns a *ComanagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ComanagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewComanagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*ComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
// returns a *ComanagedDevicesItemOverrideComplianceStateRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*ComanagedDevicesItemOverrideComplianceStateRequestBuilder) {
    return NewComanagedDevicesItemOverrideComplianceStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property comanagedDevices in deviceManagement
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
// returns a *ComanagedDevicesItemPauseConfigurationRefreshRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) PauseConfigurationRefresh()(*ComanagedDevicesItemPauseConfigurationRefreshRequestBuilder) {
    return NewComanagedDevicesItemPauseConfigurationRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PlayLostModeSound provides operations to call the playLostModeSound method.
// returns a *ComanagedDevicesItemPlayLostModeSoundRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*ComanagedDevicesItemPlayLostModeSoundRequestBuilder) {
    return NewComanagedDevicesItemPlayLostModeSoundRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RebootNow provides operations to call the rebootNow method.
// returns a *ComanagedDevicesItemRebootNowRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RebootNow()(*ComanagedDevicesItemRebootNowRequestBuilder) {
    return NewComanagedDevicesItemRebootNowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecoverPasscode provides operations to call the recoverPasscode method.
// returns a *ComanagedDevicesItemRecoverPasscodeRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*ComanagedDevicesItemRecoverPasscodeRequestBuilder) {
    return NewComanagedDevicesItemRecoverPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reenable provides operations to call the reenable method.
// returns a *ComanagedDevicesItemReenableRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Reenable()(*ComanagedDevicesItemReenableRequestBuilder) {
    return NewComanagedDevicesItemReenableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteLock provides operations to call the remoteLock method.
// returns a *ComanagedDevicesItemRemoteLockRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RemoteLock()(*ComanagedDevicesItemRemoteLockRequestBuilder) {
    return NewComanagedDevicesItemRemoteLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
// returns a *ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
// returns a *ComanagedDevicesItemReprovisionCloudPcRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*ComanagedDevicesItemReprovisionCloudPcRequestBuilder) {
    return NewComanagedDevicesItemReprovisionCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
// returns a *ComanagedDevicesItemRequestRemoteAssistanceRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ComanagedDevicesItemRequestRemoteAssistanceRequestBuilder) {
    return NewComanagedDevicesItemRequestRemoteAssistanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetPasscode provides operations to call the resetPasscode method.
// returns a *ComanagedDevicesItemResetPasscodeRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*ComanagedDevicesItemResetPasscodeRequestBuilder) {
    return NewComanagedDevicesItemResetPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResizeCloudPc provides operations to call the resizeCloudPc method.
// returns a *ComanagedDevicesItemResizeCloudPcRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ResizeCloudPc()(*ComanagedDevicesItemResizeCloudPcRequestBuilder) {
    return NewComanagedDevicesItemResizeCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RestoreCloudPc provides operations to call the restoreCloudPc method.
// returns a *ComanagedDevicesItemRestoreCloudPcRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RestoreCloudPc()(*ComanagedDevicesItemRestoreCloudPcRequestBuilder) {
    return NewComanagedDevicesItemRestoreCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Retire provides operations to call the retire method.
// returns a *ComanagedDevicesItemRetireRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Retire()(*ComanagedDevicesItemRetireRequestBuilder) {
    return NewComanagedDevicesItemRetireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
// returns a *ComanagedDevicesItemRevokeAppleVppLicensesRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*ComanagedDevicesItemRevokeAppleVppLicensesRequestBuilder) {
    return NewComanagedDevicesItemRevokeAppleVppLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
// returns a *ComanagedDevicesItemRotateBitLockerKeysRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*ComanagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    return NewComanagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
// returns a *ComanagedDevicesItemRotateFileVaultKeyRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*ComanagedDevicesItemRotateFileVaultKeyRequestBuilder) {
    return NewComanagedDevicesItemRotateFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateLocalAdminPassword provides operations to call the rotateLocalAdminPassword method.
// returns a *ComanagedDevicesItemRotateLocalAdminPasswordRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RotateLocalAdminPassword()(*ComanagedDevicesItemRotateLocalAdminPasswordRequestBuilder) {
    return NewComanagedDevicesItemRotateLocalAdminPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
// returns a *ComanagedDevicesItemSecurityBaselineStatesRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*ComanagedDevicesItemSecurityBaselineStatesRequestBuilder) {
    return NewComanagedDevicesItemSecurityBaselineStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
// returns a *ComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*ComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
// returns a *ComanagedDevicesItemSetCloudPcReviewStatusRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*ComanagedDevicesItemSetCloudPcReviewStatusRequestBuilder) {
    return NewComanagedDevicesItemSetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetDeviceName provides operations to call the setDeviceName method.
// returns a *ComanagedDevicesItemSetDeviceNameRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*ComanagedDevicesItemSetDeviceNameRequestBuilder) {
    return NewComanagedDevicesItemSetDeviceNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShutDown provides operations to call the shutDown method.
// returns a *ComanagedDevicesItemShutDownRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ShutDown()(*ComanagedDevicesItemShutDownRequestBuilder) {
    return NewComanagedDevicesItemShutDownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncDevice provides operations to call the syncDevice method.
// returns a *ComanagedDevicesItemSyncDeviceRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SyncDevice()(*ComanagedDevicesItemSyncDeviceRequestBuilder) {
    return NewComanagedDevicesItemSyncDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property comanagedDevices for deviceManagement
// returns a *RequestInformation when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    return NewComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
// returns a *ComanagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*ComanagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder) {
    return NewComanagedDevicesItemUpdateWindowsDeviceAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
// returns a *ComanagedDevicesItemUsersRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Users()(*ComanagedDevicesItemUsersRequestBuilder) {
    return NewComanagedDevicesItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
// returns a *ComanagedDevicesItemWindowsDefenderScanRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ComanagedDevicesItemWindowsDefenderScanRequestBuilder) {
    return NewComanagedDevicesItemWindowsDefenderScanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
// returns a *ComanagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*ComanagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewComanagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
// returns a *ComanagedDevicesItemWindowsProtectionStateRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ComanagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewComanagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Wipe provides operations to call the wipe method.
// returns a *ComanagedDevicesItemWipeRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Wipe()(*ComanagedDevicesItemWipeRequestBuilder) {
    return NewComanagedDevicesItemWipeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ComanagedDevicesManagedDeviceItemRequestBuilder when successful
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    return NewComanagedDevicesManagedDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
