package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeManagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type MeManagedDevicesManagedDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters the managed devices associated with the user.
type MeManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// MeManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateDeviceEsim provides operations to call the activateDeviceEsim method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*MeManagedDevicesItemActivateDeviceEsimRequestBuilder) {
    return NewMeManagedDevicesItemActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewMeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetailsById provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetailsById(id string)(*MeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignmentFilterEvaluationStatusDetails%2Did"] = id
    }
    return NewMeManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*MeManagedDevicesItemBypassActivationLockRequestBuilder) {
    return NewMeManagedDevicesItemBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*MeManagedDevicesItemCleanWindowsDeviceRequestBuilder) {
    return NewMeManagedDevicesItemCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMeManagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewMeManagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeManagedDevicesManagedDeviceItemRequestBuilder) {
    m := &MeManagedDevicesManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeManagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewMeManagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeManagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeManagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedDevices for me
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*MeManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewMeManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the managed devices associated with the user.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property managedDevices in me
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *MeManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateRemoteHelpSession provides operations to call the createRemoteHelpSession method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) CreateRemoteHelpSession()(*MeManagedDevicesItemCreateRemoteHelpSessionRequestBuilder) {
    return NewMeManagedDevicesItemCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property managedDevices for me
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeleteUserFromSharedAppleDevice provides operations to call the deleteUserFromSharedAppleDevice method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*MeManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewMeManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Deprovision provides operations to call the deprovision method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) Deprovision()(*MeManagedDevicesItemDeprovisionRequestBuilder) {
    return NewMeManagedDevicesItemDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) DetectedApps()(*MeManagedDevicesItemDetectedAppsRequestBuilder) {
    return NewMeManagedDevicesItemDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedAppsById provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) DetectedAppsById(id string)(*MeManagedDevicesItemDetectedAppsDetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp%2Did"] = id
    }
    return NewMeManagedDevicesItemDetectedAppsDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*MeManagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewMeManagedDevicesItemDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*MeManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewMeManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStatesById provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*MeManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = id
    }
    return NewMeManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*MeManagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewMeManagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationStatesById provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*MeManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState%2Did"] = id
    }
    return NewMeManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Disable provides operations to call the disable method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) Disable()(*MeManagedDevicesItemDisableRequestBuilder) {
    return NewMeManagedDevicesItemDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DisableLostMode provides operations to call the disableLostMode method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*MeManagedDevicesItemDisableLostModeRequestBuilder) {
    return NewMeManagedDevicesItemDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnableLostMode provides operations to call the enableLostMode method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*MeManagedDevicesItemEnableLostModeRequestBuilder) {
    return NewMeManagedDevicesItemEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndRemoteHelpSession provides operations to call the endRemoteHelpSession method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) EndRemoteHelpSession()(*MeManagedDevicesItemEndRemoteHelpSessionRequestBuilder) {
    return NewMeManagedDevicesItemEndRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollNowAction provides operations to call the enrollNowAction method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*MeManagedDevicesItemEnrollNowActionRequestBuilder) {
    return NewMeManagedDevicesItemEnrollNowActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the managed devices associated with the user.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
// GetCloudPcRemoteActionResults provides operations to call the getCloudPcRemoteActionResults method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*MeManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewMeManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*MeManagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    return NewMeManagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*MeManagedDevicesItemGetFileVaultKeyRequestBuilder) {
    return NewMeManagedDevicesItemGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*MeManagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    return NewMeManagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty provides operations to call the getOemWarranty method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) GetOemWarranty()(*MeManagedDevicesItemGetOemWarrantyRequestBuilder) {
    return NewMeManagedDevicesItemGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*MeManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewMeManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocateDevice provides operations to call the locateDevice method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) LocateDevice()(*MeManagedDevicesItemLocateDeviceRequestBuilder) {
    return NewMeManagedDevicesItemLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*MeManagedDevicesItemLogCollectionRequestsRequestBuilder) {
    return NewMeManagedDevicesItemLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequestsById provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequestsById(id string)(*MeManagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceLogCollectionResponse%2Did"] = id
    }
    return NewMeManagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*MeManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewMeManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*MeManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewMeManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStatesById provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStatesById(id string)(*MeManagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState%2Did"] = id
    }
    return NewMeManagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*MeManagedDevicesItemOverrideComplianceStateRequestBuilder) {
    return NewMeManagedDevicesItemOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property managedDevices in me
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *MeManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
// PlayLostModeSound provides operations to call the playLostModeSound method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*MeManagedDevicesItemPlayLostModeSoundRequestBuilder) {
    return NewMeManagedDevicesItemPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RebootNow provides operations to call the rebootNow method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) RebootNow()(*MeManagedDevicesItemRebootNowRequestBuilder) {
    return NewMeManagedDevicesItemRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecoverPasscode provides operations to call the recoverPasscode method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*MeManagedDevicesItemRecoverPasscodeRequestBuilder) {
    return NewMeManagedDevicesItemRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reenable provides operations to call the reenable method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) Reenable()(*MeManagedDevicesItemReenableRequestBuilder) {
    return NewMeManagedDevicesItemReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteLock provides operations to call the remoteLock method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) RemoteLock()(*MeManagedDevicesItemRemoteLockRequestBuilder) {
    return NewMeManagedDevicesItemRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*MeManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewMeManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*MeManagedDevicesItemReprovisionCloudPcRequestBuilder) {
    return NewMeManagedDevicesItemReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*MeManagedDevicesItemRequestRemoteAssistanceRequestBuilder) {
    return NewMeManagedDevicesItemRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteHelpSessionAccess provides operations to call the requestRemoteHelpSessionAccess method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteHelpSessionAccess()(*MeManagedDevicesItemRequestRemoteHelpSessionAccessRequestBuilder) {
    return NewMeManagedDevicesItemRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetPasscode provides operations to call the resetPasscode method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*MeManagedDevicesItemResetPasscodeRequestBuilder) {
    return NewMeManagedDevicesItemResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResizeCloudPc provides operations to call the resizeCloudPc method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) ResizeCloudPc()(*MeManagedDevicesItemResizeCloudPcRequestBuilder) {
    return NewMeManagedDevicesItemResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RestoreCloudPc provides operations to call the restoreCloudPc method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) RestoreCloudPc()(*MeManagedDevicesItemRestoreCloudPcRequestBuilder) {
    return NewMeManagedDevicesItemRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Retire provides operations to call the retire method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) Retire()(*MeManagedDevicesItemRetireRequestBuilder) {
    return NewMeManagedDevicesItemRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) RetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*MeManagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return NewMeManagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey);
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*MeManagedDevicesItemRevokeAppleVppLicensesRequestBuilder) {
    return NewMeManagedDevicesItemRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*MeManagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    return NewMeManagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*MeManagedDevicesItemRotateFileVaultKeyRequestBuilder) {
    return NewMeManagedDevicesItemRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*MeManagedDevicesItemSecurityBaselineStatesRequestBuilder) {
    return NewMeManagedDevicesItemSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStatesById provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStatesById(id string)(*MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineState%2Did"] = id
    }
    return NewMeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*MeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewMeManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*MeManagedDevicesItemSetCloudPcReviewStatusRequestBuilder) {
    return NewMeManagedDevicesItemSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetDeviceName provides operations to call the setDeviceName method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*MeManagedDevicesItemSetDeviceNameRequestBuilder) {
    return NewMeManagedDevicesItemSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShutDown provides operations to call the shutDown method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) ShutDown()(*MeManagedDevicesItemShutDownRequestBuilder) {
    return NewMeManagedDevicesItemShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncDevice provides operations to call the syncDevice method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) SyncDevice()(*MeManagedDevicesItemSyncDeviceRequestBuilder) {
    return NewMeManagedDevicesItemSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TriggerConfigurationManagerAction provides operations to call the triggerConfigurationManagerAction method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*MeManagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    return NewMeManagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*MeManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder) {
    return NewMeManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) Users()(*MeManagedDevicesItemUsersRequestBuilder) {
    return NewMeManagedDevicesItemUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*MeManagedDevicesItemWindowsDefenderScanRequestBuilder) {
    return NewMeManagedDevicesItemWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*MeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewMeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*MeManagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewMeManagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Wipe provides operations to call the wipe method.
func (m *MeManagedDevicesManagedDeviceItemRequestBuilder) Wipe()(*MeManagedDevicesItemWipeRequestBuilder) {
    return NewMeManagedDevicesItemWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
