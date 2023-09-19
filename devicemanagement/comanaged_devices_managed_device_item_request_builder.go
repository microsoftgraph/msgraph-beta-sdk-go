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
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*ComanagedDevicesItemActivateDeviceEsimRequestBuilder) {
    return NewComanagedDevicesItemActivateDeviceEsimRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*ComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*ComanagedDevicesItemBypassActivationLockRequestBuilder) {
    return NewComanagedDevicesItemBypassActivationLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ComanagedDevicesItemCleanWindowsDeviceRequestBuilder) {
    return NewComanagedDevicesItemCleanWindowsDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewComanagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewComanagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    m := &ComanagedDevicesManagedDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewComanagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewComanagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*ComanagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewComanagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property comanagedDevices for deviceManagement
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeleteUserFromSharedAppleDevice provides operations to call the deleteUserFromSharedAppleDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ComanagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewComanagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Deprovision provides operations to call the deprovision method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Deprovision()(*ComanagedDevicesItemDeprovisionRequestBuilder) {
    return NewComanagedDevicesItemDeprovisionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DetectedApps()(*ComanagedDevicesItemDetectedAppsRequestBuilder) {
    return NewComanagedDevicesItemDetectedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ComanagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewComanagedDevicesItemDeviceCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ComanagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewComanagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ComanagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewComanagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceHealthScriptStates provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceHealthScriptStates()(*ComanagedDevicesItemDeviceHealthScriptStatesRequestBuilder) {
    return NewComanagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Disable provides operations to call the disable method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Disable()(*ComanagedDevicesItemDisableRequestBuilder) {
    return NewComanagedDevicesItemDisableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DisableLostMode provides operations to call the disableLostMode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*ComanagedDevicesItemDisableLostModeRequestBuilder) {
    return NewComanagedDevicesItemDisableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableLostMode provides operations to call the enableLostMode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*ComanagedDevicesItemEnableLostModeRequestBuilder) {
    return NewComanagedDevicesItemEnableLostModeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnrollNowAction provides operations to call the enrollNowAction method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*ComanagedDevicesItemEnrollNowActionRequestBuilder) {
    return NewComanagedDevicesItemEnrollNowActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of co-managed devices report
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*ComanagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    return NewComanagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*ComanagedDevicesItemGetFileVaultKeyRequestBuilder) {
    return NewComanagedDevicesItemGetFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*ComanagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    return NewComanagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*ComanagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewComanagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InitiateOnDemandProactiveRemediation provides operations to call the initiateOnDemandProactiveRemediation method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) InitiateOnDemandProactiveRemediation()(*ComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewComanagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LocateDevice provides operations to call the locateDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) LocateDevice()(*ComanagedDevicesItemLocateDeviceRequestBuilder) {
    return NewComanagedDevicesItemLocateDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ComanagedDevicesItemLogCollectionRequestsRequestBuilder) {
    return NewComanagedDevicesItemLogCollectionRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ComanagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewComanagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*ComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*ComanagedDevicesItemOverrideComplianceStateRequestBuilder) {
    return NewComanagedDevicesItemOverrideComplianceStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property comanagedDevices in deviceManagement
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// PlayLostModeSound provides operations to call the playLostModeSound method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*ComanagedDevicesItemPlayLostModeSoundRequestBuilder) {
    return NewComanagedDevicesItemPlayLostModeSoundRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RebootNow provides operations to call the rebootNow method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RebootNow()(*ComanagedDevicesItemRebootNowRequestBuilder) {
    return NewComanagedDevicesItemRebootNowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecoverPasscode provides operations to call the recoverPasscode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*ComanagedDevicesItemRecoverPasscodeRequestBuilder) {
    return NewComanagedDevicesItemRecoverPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reenable provides operations to call the reenable method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Reenable()(*ComanagedDevicesItemReenableRequestBuilder) {
    return NewComanagedDevicesItemReenableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteLock provides operations to call the remoteLock method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RemoteLock()(*ComanagedDevicesItemRemoteLockRequestBuilder) {
    return NewComanagedDevicesItemRemoteLockRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*ComanagedDevicesItemReprovisionCloudPcRequestBuilder) {
    return NewComanagedDevicesItemReprovisionCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ComanagedDevicesItemRequestRemoteAssistanceRequestBuilder) {
    return NewComanagedDevicesItemRequestRemoteAssistanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetPasscode provides operations to call the resetPasscode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*ComanagedDevicesItemResetPasscodeRequestBuilder) {
    return NewComanagedDevicesItemResetPasscodeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResizeCloudPc provides operations to call the resizeCloudPc method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ResizeCloudPc()(*ComanagedDevicesItemResizeCloudPcRequestBuilder) {
    return NewComanagedDevicesItemResizeCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RestoreCloudPc provides operations to call the restoreCloudPc method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RestoreCloudPc()(*ComanagedDevicesItemRestoreCloudPcRequestBuilder) {
    return NewComanagedDevicesItemRestoreCloudPcRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Retire provides operations to call the retire method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Retire()(*ComanagedDevicesItemRetireRequestBuilder) {
    return NewComanagedDevicesItemRetireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*ComanagedDevicesItemRevokeAppleVppLicensesRequestBuilder) {
    return NewComanagedDevicesItemRevokeAppleVppLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*ComanagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    return NewComanagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*ComanagedDevicesItemRotateFileVaultKeyRequestBuilder) {
    return NewComanagedDevicesItemRotateFileVaultKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RotateLocalAdminPassword provides operations to call the rotateLocalAdminPassword method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RotateLocalAdminPassword()(*ComanagedDevicesItemRotateLocalAdminPasswordRequestBuilder) {
    return NewComanagedDevicesItemRotateLocalAdminPasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*ComanagedDevicesItemSecurityBaselineStatesRequestBuilder) {
    return NewComanagedDevicesItemSecurityBaselineStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*ComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*ComanagedDevicesItemSetCloudPcReviewStatusRequestBuilder) {
    return NewComanagedDevicesItemSetCloudPcReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetDeviceName provides operations to call the setDeviceName method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*ComanagedDevicesItemSetDeviceNameRequestBuilder) {
    return NewComanagedDevicesItemSetDeviceNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShutDown provides operations to call the shutDown method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ShutDown()(*ComanagedDevicesItemShutDownRequestBuilder) {
    return NewComanagedDevicesItemShutDownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncDevice provides operations to call the syncDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SyncDevice()(*ComanagedDevicesItemSyncDeviceRequestBuilder) {
    return NewComanagedDevicesItemSyncDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property comanagedDevices for deviceManagement
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the list of co-managed devices report
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property comanagedDevices in deviceManagement
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// TriggerConfigurationManagerAction provides operations to call the triggerConfigurationManagerAction method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    return NewComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*ComanagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder) {
    return NewComanagedDevicesItemUpdateWindowsDeviceAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Users()(*ComanagedDevicesItemUsersRequestBuilder) {
    return NewComanagedDevicesItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ComanagedDevicesItemWindowsDefenderScanRequestBuilder) {
    return NewComanagedDevicesItemWindowsDefenderScanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*ComanagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewComanagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ComanagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewComanagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Wipe provides operations to call the wipe method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Wipe()(*ComanagedDevicesItemWipeRequestBuilder) {
    return NewComanagedDevicesItemWipeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    return NewComanagedDevicesManagedDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
