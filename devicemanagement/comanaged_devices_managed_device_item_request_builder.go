package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
type ComanagedDevicesManagedDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
    return NewComanagedDevicesItemActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*ComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetailsById provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetailsById(id string)(*ComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignmentFilterEvaluationStatusDetails%2Did"] = id
    }
    return NewComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*ComanagedDevicesItemBypassActivationLockRequestBuilder) {
    return NewComanagedDevicesItemBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ComanagedDevicesItemCleanWindowsDeviceRequestBuilder) {
    return NewComanagedDevicesItemCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewComanagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewComanagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    m := &ComanagedDevicesManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewComanagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewComanagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property comanagedDevices for deviceManagement
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*ComanagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewComanagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of co-managed devices report
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// CreatePatchRequestInformation update the navigation property comanagedDevices in deviceManagement
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateRemoteHelpSession provides operations to call the createRemoteHelpSession method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) CreateRemoteHelpSession()(*ComanagedDevicesItemCreateRemoteHelpSessionRequestBuilder) {
    return NewComanagedDevicesItemCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property comanagedDevices for deviceManagement
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ComanagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewComanagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Deprovision provides operations to call the deprovision method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Deprovision()(*ComanagedDevicesItemDeprovisionRequestBuilder) {
    return NewComanagedDevicesItemDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DetectedApps()(*ComanagedDevicesItemDetectedAppsRequestBuilder) {
    return NewComanagedDevicesItemDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedAppsById provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DetectedAppsById(id string)(*ComanagedDevicesItemDetectedAppsDetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp%2Did"] = id
    }
    return NewComanagedDevicesItemDetectedAppsDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ComanagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewComanagedDevicesItemDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ComanagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewComanagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStatesById provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*ComanagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = id
    }
    return NewComanagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ComanagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewComanagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationStatesById provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*ComanagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState%2Did"] = id
    }
    return NewComanagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Disable provides operations to call the disable method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Disable()(*ComanagedDevicesItemDisableRequestBuilder) {
    return NewComanagedDevicesItemDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DisableLostMode provides operations to call the disableLostMode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*ComanagedDevicesItemDisableLostModeRequestBuilder) {
    return NewComanagedDevicesItemDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnableLostMode provides operations to call the enableLostMode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*ComanagedDevicesItemEnableLostModeRequestBuilder) {
    return NewComanagedDevicesItemEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndRemoteHelpSession provides operations to call the endRemoteHelpSession method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) EndRemoteHelpSession()(*ComanagedDevicesItemEndRemoteHelpSessionRequestBuilder) {
    return NewComanagedDevicesItemEndRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollNowAction provides operations to call the enrollNowAction method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*ComanagedDevicesItemEnrollNowActionRequestBuilder) {
    return NewComanagedDevicesItemEnrollNowActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of co-managed devices report
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*ComanagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    return NewComanagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*ComanagedDevicesItemGetFileVaultKeyRequestBuilder) {
    return NewComanagedDevicesItemGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*ComanagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    return NewComanagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty provides operations to call the getOemWarranty method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) GetOemWarranty()(*ComanagedDevicesItemGetOemWarrantyRequestBuilder) {
    return NewComanagedDevicesItemGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*ComanagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewComanagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocateDevice provides operations to call the locateDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) LocateDevice()(*ComanagedDevicesItemLocateDeviceRequestBuilder) {
    return NewComanagedDevicesItemLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ComanagedDevicesItemLogCollectionRequestsRequestBuilder) {
    return NewComanagedDevicesItemLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequestsById provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequestsById(id string)(*ComanagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceLogCollectionResponse%2Did"] = id
    }
    return NewComanagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ComanagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewComanagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*ComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStatesById provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStatesById(id string)(*ComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState%2Did"] = id
    }
    return NewComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*ComanagedDevicesItemOverrideComplianceStateRequestBuilder) {
    return NewComanagedDevicesItemOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property comanagedDevices in deviceManagement
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*ComanagedDevicesItemPlayLostModeSoundRequestBuilder) {
    return NewComanagedDevicesItemPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RebootNow provides operations to call the rebootNow method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RebootNow()(*ComanagedDevicesItemRebootNowRequestBuilder) {
    return NewComanagedDevicesItemRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecoverPasscode provides operations to call the recoverPasscode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*ComanagedDevicesItemRecoverPasscodeRequestBuilder) {
    return NewComanagedDevicesItemRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reenable provides operations to call the reenable method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Reenable()(*ComanagedDevicesItemReenableRequestBuilder) {
    return NewComanagedDevicesItemReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteLock provides operations to call the remoteLock method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RemoteLock()(*ComanagedDevicesItemRemoteLockRequestBuilder) {
    return NewComanagedDevicesItemRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*ComanagedDevicesItemReprovisionCloudPcRequestBuilder) {
    return NewComanagedDevicesItemReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ComanagedDevicesItemRequestRemoteAssistanceRequestBuilder) {
    return NewComanagedDevicesItemRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteHelpSessionAccess provides operations to call the requestRemoteHelpSessionAccess method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteHelpSessionAccess()(*ComanagedDevicesItemRequestRemoteHelpSessionAccessRequestBuilder) {
    return NewComanagedDevicesItemRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetPasscode provides operations to call the resetPasscode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*ComanagedDevicesItemResetPasscodeRequestBuilder) {
    return NewComanagedDevicesItemResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResizeCloudPc provides operations to call the resizeCloudPc method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ResizeCloudPc()(*ComanagedDevicesItemResizeCloudPcRequestBuilder) {
    return NewComanagedDevicesItemResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RestoreCloudPc provides operations to call the restoreCloudPc method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RestoreCloudPc()(*ComanagedDevicesItemRestoreCloudPcRequestBuilder) {
    return NewComanagedDevicesItemRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Retire provides operations to call the retire method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Retire()(*ComanagedDevicesItemRetireRequestBuilder) {
    return NewComanagedDevicesItemRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*ComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return NewComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey);
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*ComanagedDevicesItemRevokeAppleVppLicensesRequestBuilder) {
    return NewComanagedDevicesItemRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*ComanagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    return NewComanagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*ComanagedDevicesItemRotateFileVaultKeyRequestBuilder) {
    return NewComanagedDevicesItemRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*ComanagedDevicesItemSecurityBaselineStatesRequestBuilder) {
    return NewComanagedDevicesItemSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStatesById provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStatesById(id string)(*ComanagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineState%2Did"] = id
    }
    return NewComanagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*ComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*ComanagedDevicesItemSetCloudPcReviewStatusRequestBuilder) {
    return NewComanagedDevicesItemSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetDeviceName provides operations to call the setDeviceName method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*ComanagedDevicesItemSetDeviceNameRequestBuilder) {
    return NewComanagedDevicesItemSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShutDown provides operations to call the shutDown method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ShutDown()(*ComanagedDevicesItemShutDownRequestBuilder) {
    return NewComanagedDevicesItemShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncDevice provides operations to call the syncDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SyncDevice()(*ComanagedDevicesItemSyncDeviceRequestBuilder) {
    return NewComanagedDevicesItemSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TriggerConfigurationManagerAction provides operations to call the triggerConfigurationManagerAction method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    return NewComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*ComanagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder) {
    return NewComanagedDevicesItemUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Users()(*ComanagedDevicesItemUsersRequestBuilder) {
    return NewComanagedDevicesItemUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ComanagedDevicesItemWindowsDefenderScanRequestBuilder) {
    return NewComanagedDevicesItemWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*ComanagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewComanagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ComanagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewComanagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Wipe provides operations to call the wipe method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Wipe()(*ComanagedDevicesItemWipeRequestBuilder) {
    return NewComanagedDevicesItemWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
