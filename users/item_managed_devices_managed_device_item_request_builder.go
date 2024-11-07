package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ItemManagedDevicesManagedDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters the managed devices associated with the user.
type ItemManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateDeviceEsim provides operations to call the activateDeviceEsim method.
// returns a *ItemManagedDevicesItemActivateDeviceEsimRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*ItemManagedDevicesItemActivateDeviceEsimRequestBuilder) {
    return NewItemManagedDevicesItemActivateDeviceEsimRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
// returns a *ItemManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*ItemManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewItemManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
// returns a *ItemManagedDevicesItemBypassActivationLockRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*ItemManagedDevicesItemBypassActivationLockRequestBuilder) {
    return NewItemManagedDevicesItemBypassActivationLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChangeAssignments provides operations to call the changeAssignments method.
// returns a *ItemManagedDevicesItemChangeAssignmentsRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ChangeAssignments()(*ItemManagedDevicesItemChangeAssignmentsRequestBuilder) {
    return NewItemManagedDevicesItemChangeAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
// returns a *ItemManagedDevicesItemCleanWindowsDeviceRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ItemManagedDevicesItemCleanWindowsDeviceRequestBuilder) {
    return NewItemManagedDevicesItemCleanWindowsDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ItemManagedDevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesManagedDeviceItemRequestBuilder) {
    m := &ItemManagedDevicesManagedDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemManagedDevicesManagedDeviceItemRequestBuilder instantiates a new ItemManagedDevicesManagedDeviceItemRequestBuilder and sets the default values.
func NewItemManagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
// returns a *ItemManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*ItemManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewItemManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property managedDevices for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ItemManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ItemManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewItemManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Deprovision provides operations to call the deprovision method.
// returns a *ItemManagedDevicesItemDeprovisionRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Deprovision()(*ItemManagedDevicesItemDeprovisionRequestBuilder) {
    return NewItemManagedDevicesItemDeprovisionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
// returns a *ItemManagedDevicesItemDetectedAppsRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DetectedApps()(*ItemManagedDevicesItemDetectedAppsRequestBuilder) {
    return NewItemManagedDevicesItemDetectedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
// returns a *ItemManagedDevicesItemDeviceCategoryRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ItemManagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewItemManagedDevicesItemDeviceCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ItemManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewItemManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManagedDevicesItemDeviceConfigurationStatesRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ItemManagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewItemManagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceHealthScriptStates provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManagedDevicesItemDeviceHealthScriptStatesRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceHealthScriptStates()(*ItemManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) {
    return NewItemManagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Disable provides operations to call the disable method.
// returns a *ItemManagedDevicesItemDisableRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Disable()(*ItemManagedDevicesItemDisableRequestBuilder) {
    return NewItemManagedDevicesItemDisableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DisableLostMode provides operations to call the disableLostMode method.
// returns a *ItemManagedDevicesItemDisableLostModeRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*ItemManagedDevicesItemDisableLostModeRequestBuilder) {
    return NewItemManagedDevicesItemDisableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableLostMode provides operations to call the enableLostMode method.
// returns a *ItemManagedDevicesItemEnableLostModeRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*ItemManagedDevicesItemEnableLostModeRequestBuilder) {
    return NewItemManagedDevicesItemEnableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnrollNowAction provides operations to call the enrollNowAction method.
// returns a *ItemManagedDevicesItemEnrollNowActionRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*ItemManagedDevicesItemEnrollNowActionRequestBuilder) {
    return NewItemManagedDevicesItemEnrollNowActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the managed devices associated with the user.
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
// returns a *ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
// returns a *ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    return NewItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
// returns a *ItemManagedDevicesItemGetFileVaultKeyRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*ItemManagedDevicesItemGetFileVaultKeyRequestBuilder) {
    return NewItemManagedDevicesItemGetFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
// returns a *ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    return NewItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateDeviceAttestation provides operations to call the initiateDeviceAttestation method.
// returns a *ItemManagedDevicesItemInitiateDeviceAttestationRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) InitiateDeviceAttestation()(*ItemManagedDevicesItemInitiateDeviceAttestationRequestBuilder) {
    return NewItemManagedDevicesItemInitiateDeviceAttestationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
// returns a *ItemManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*ItemManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewItemManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateOnDemandProactiveRemediation provides operations to call the initiateOnDemandProactiveRemediation method.
// returns a *ItemManagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) InitiateOnDemandProactiveRemediation()(*ItemManagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewItemManagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LocateDevice provides operations to call the locateDevice method.
// returns a *ItemManagedDevicesItemLocateDeviceRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) LocateDevice()(*ItemManagedDevicesItemLocateDeviceRequestBuilder) {
    return NewItemManagedDevicesItemLocateDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
// returns a *ItemManagedDevicesItemLogCollectionRequestsRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ItemManagedDevicesItemLogCollectionRequestsRequestBuilder) {
    return NewItemManagedDevicesItemLogCollectionRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
// returns a *ItemManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ItemManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewItemManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*ItemManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewItemManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
// returns a *ItemManagedDevicesItemOverrideComplianceStateRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*ItemManagedDevicesItemOverrideComplianceStateRequestBuilder) {
    return NewItemManagedDevicesItemOverrideComplianceStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property managedDevices in users
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
// returns a *ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) PauseConfigurationRefresh()(*ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder) {
    return NewItemManagedDevicesItemPauseConfigurationRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PlayLostModeSound provides operations to call the playLostModeSound method.
// returns a *ItemManagedDevicesItemPlayLostModeSoundRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*ItemManagedDevicesItemPlayLostModeSoundRequestBuilder) {
    return NewItemManagedDevicesItemPlayLostModeSoundRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RebootNow provides operations to call the rebootNow method.
// returns a *ItemManagedDevicesItemRebootNowRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RebootNow()(*ItemManagedDevicesItemRebootNowRequestBuilder) {
    return NewItemManagedDevicesItemRebootNowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecoverPasscode provides operations to call the recoverPasscode method.
// returns a *ItemManagedDevicesItemRecoverPasscodeRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*ItemManagedDevicesItemRecoverPasscodeRequestBuilder) {
    return NewItemManagedDevicesItemRecoverPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reenable provides operations to call the reenable method.
// returns a *ItemManagedDevicesItemReenableRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Reenable()(*ItemManagedDevicesItemReenableRequestBuilder) {
    return NewItemManagedDevicesItemReenableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteLock provides operations to call the remoteLock method.
// returns a *ItemManagedDevicesItemRemoteLockRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RemoteLock()(*ItemManagedDevicesItemRemoteLockRequestBuilder) {
    return NewItemManagedDevicesItemRemoteLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
// returns a *ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
// returns a *ItemManagedDevicesItemRequestRemoteAssistanceRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ItemManagedDevicesItemRequestRemoteAssistanceRequestBuilder) {
    return NewItemManagedDevicesItemRequestRemoteAssistanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetPasscode provides operations to call the resetPasscode method.
// returns a *ItemManagedDevicesItemResetPasscodeRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*ItemManagedDevicesItemResetPasscodeRequestBuilder) {
    return NewItemManagedDevicesItemResetPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Retire provides operations to call the retire method.
// returns a *ItemManagedDevicesItemRetireRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Retire()(*ItemManagedDevicesItemRetireRequestBuilder) {
    return NewItemManagedDevicesItemRetireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
// returns a *ItemManagedDevicesItemRevokeAppleVppLicensesRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*ItemManagedDevicesItemRevokeAppleVppLicensesRequestBuilder) {
    return NewItemManagedDevicesItemRevokeAppleVppLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
// returns a *ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    return NewItemManagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
// returns a *ItemManagedDevicesItemRotateFileVaultKeyRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*ItemManagedDevicesItemRotateFileVaultKeyRequestBuilder) {
    return NewItemManagedDevicesItemRotateFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateLocalAdminPassword provides operations to call the rotateLocalAdminPassword method.
// returns a *ItemManagedDevicesItemRotateLocalAdminPasswordRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RotateLocalAdminPassword()(*ItemManagedDevicesItemRotateLocalAdminPasswordRequestBuilder) {
    return NewItemManagedDevicesItemRotateLocalAdminPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
// returns a *ItemManagedDevicesItemSecurityBaselineStatesRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*ItemManagedDevicesItemSecurityBaselineStatesRequestBuilder) {
    return NewItemManagedDevicesItemSecurityBaselineStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
// returns a *ItemManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*ItemManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewItemManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
// returns a *ItemManagedDevicesItemSetCloudPcReviewStatusRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*ItemManagedDevicesItemSetCloudPcReviewStatusRequestBuilder) {
    return NewItemManagedDevicesItemSetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetDeviceName provides operations to call the setDeviceName method.
// returns a *ItemManagedDevicesItemSetDeviceNameRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*ItemManagedDevicesItemSetDeviceNameRequestBuilder) {
    return NewItemManagedDevicesItemSetDeviceNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShutDown provides operations to call the shutDown method.
// returns a *ItemManagedDevicesItemShutDownRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ShutDown()(*ItemManagedDevicesItemShutDownRequestBuilder) {
    return NewItemManagedDevicesItemShutDownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncDevice provides operations to call the syncDevice method.
// returns a *ItemManagedDevicesItemSyncDeviceRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) SyncDevice()(*ItemManagedDevicesItemSyncDeviceRequestBuilder) {
    return NewItemManagedDevicesItemSyncDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property managedDevices for users
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManagedDevicesItemTriggerConfigurationManagerActionRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*ItemManagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    return NewItemManagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
// returns a *ItemManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*ItemManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder) {
    return NewItemManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
// returns a *ItemManagedDevicesItemUsersRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Users()(*ItemManagedDevicesItemUsersRequestBuilder) {
    return NewItemManagedDevicesItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
// returns a *ItemManagedDevicesItemWindowsDefenderScanRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ItemManagedDevicesItemWindowsDefenderScanRequestBuilder) {
    return NewItemManagedDevicesItemWindowsDefenderScanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
// returns a *ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
// returns a *ItemManagedDevicesItemWindowsProtectionStateRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ItemManagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewItemManagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Wipe provides operations to call the wipe method.
// returns a *ItemManagedDevicesItemWipeRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Wipe()(*ItemManagedDevicesItemWipeRequestBuilder) {
    return NewItemManagedDevicesItemWipeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManagedDevicesManagedDeviceItemRequestBuilder when successful
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesManagedDeviceItemRequestBuilder) {
    return NewItemManagedDevicesManagedDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
