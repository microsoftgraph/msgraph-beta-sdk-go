package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
type DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters the list of managed devices.
type DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateDeviceEsim provides operations to call the activateDeviceEsim method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*DeviceManagementManagedDevicesItemActivateDeviceEsimRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*DeviceManagementManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetailsById provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetailsById(id string)(*DeviceManagementManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignmentFilterEvaluationStatusDetails%2Did"] = id
    }
    return NewDeviceManagementManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*DeviceManagementManagedDevicesItemBypassActivationLockRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*DeviceManagementManagedDevicesItemCleanWindowsDeviceRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementManagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewDeviceManagementManagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) {
    m := &DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementManagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewDeviceManagementManagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementManagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedDevices for deviceManagement
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*DeviceManagementManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of managed devices.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property managedDevices in deviceManagement
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) CreateRemoteHelpSession()(*DeviceManagementManagedDevicesItemCreateRemoteHelpSessionRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property managedDevices for deviceManagement
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*DeviceManagementManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Deprovision provides operations to call the deprovision method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) Deprovision()(*DeviceManagementManagedDevicesItemDeprovisionRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) DetectedApps()(*DeviceManagementManagedDevicesItemDetectedAppsRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedAppsById provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) DetectedAppsById(id string)(*DeviceManagementManagedDevicesItemDetectedAppsDetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp%2Did"] = id
    }
    return NewDeviceManagementManagedDevicesItemDetectedAppsDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*DeviceManagementManagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStatesById provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*DeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = id
    }
    return NewDeviceManagementManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*DeviceManagementManagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationStatesById provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*DeviceManagementManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState%2Did"] = id
    }
    return NewDeviceManagementManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Disable provides operations to call the disable method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) Disable()(*DeviceManagementManagedDevicesItemDisableRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DisableLostMode provides operations to call the disableLostMode method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*DeviceManagementManagedDevicesItemDisableLostModeRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnableLostMode provides operations to call the enableLostMode method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*DeviceManagementManagedDevicesItemEnableLostModeRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndRemoteHelpSession provides operations to call the endRemoteHelpSession method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) EndRemoteHelpSession()(*DeviceManagementManagedDevicesItemEndRemoteHelpSessionRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemEndRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollNowAction provides operations to call the enrollNowAction method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*DeviceManagementManagedDevicesItemEnrollNowActionRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemEnrollNowActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of managed devices.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*DeviceManagementManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*DeviceManagementManagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*DeviceManagementManagedDevicesItemGetFileVaultKeyRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*DeviceManagementManagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty provides operations to call the getOemWarranty method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) GetOemWarranty()(*DeviceManagementManagedDevicesItemGetOemWarrantyRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*DeviceManagementManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocateDevice provides operations to call the locateDevice method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) LocateDevice()(*DeviceManagementManagedDevicesItemLocateDeviceRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*DeviceManagementManagedDevicesItemLogCollectionRequestsRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequestsById provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequestsById(id string)(*DeviceManagementManagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceLogCollectionResponse%2Did"] = id
    }
    return NewDeviceManagementManagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*DeviceManagementManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*DeviceManagementManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStatesById provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStatesById(id string)(*DeviceManagementManagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState%2Did"] = id
    }
    return NewDeviceManagementManagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*DeviceManagementManagedDevicesItemOverrideComplianceStateRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property managedDevices in deviceManagement
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*DeviceManagementManagedDevicesItemPlayLostModeSoundRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RebootNow provides operations to call the rebootNow method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) RebootNow()(*DeviceManagementManagedDevicesItemRebootNowRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecoverPasscode provides operations to call the recoverPasscode method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*DeviceManagementManagedDevicesItemRecoverPasscodeRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reenable provides operations to call the reenable method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) Reenable()(*DeviceManagementManagedDevicesItemReenableRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteLock provides operations to call the remoteLock method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) RemoteLock()(*DeviceManagementManagedDevicesItemRemoteLockRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*DeviceManagementManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*DeviceManagementManagedDevicesItemReprovisionCloudPcRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*DeviceManagementManagedDevicesItemRequestRemoteAssistanceRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteHelpSessionAccess provides operations to call the requestRemoteHelpSessionAccess method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteHelpSessionAccess()(*DeviceManagementManagedDevicesItemRequestRemoteHelpSessionAccessRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetPasscode provides operations to call the resetPasscode method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*DeviceManagementManagedDevicesItemResetPasscodeRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResizeCloudPc provides operations to call the resizeCloudPc method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) ResizeCloudPc()(*DeviceManagementManagedDevicesItemResizeCloudPcRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RestoreCloudPc provides operations to call the restoreCloudPc method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) RestoreCloudPc()(*DeviceManagementManagedDevicesItemRestoreCloudPcRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Retire provides operations to call the retire method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) Retire()(*DeviceManagementManagedDevicesItemRetireRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) RetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*DeviceManagementManagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey);
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*DeviceManagementManagedDevicesItemRevokeAppleVppLicensesRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*DeviceManagementManagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*DeviceManagementManagedDevicesItemRotateFileVaultKeyRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*DeviceManagementManagedDevicesItemSecurityBaselineStatesRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStatesById provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStatesById(id string)(*DeviceManagementManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineState%2Did"] = id
    }
    return NewDeviceManagementManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*DeviceManagementManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*DeviceManagementManagedDevicesItemSetCloudPcReviewStatusRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetDeviceName provides operations to call the setDeviceName method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*DeviceManagementManagedDevicesItemSetDeviceNameRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShutDown provides operations to call the shutDown method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) ShutDown()(*DeviceManagementManagedDevicesItemShutDownRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncDevice provides operations to call the syncDevice method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) SyncDevice()(*DeviceManagementManagedDevicesItemSyncDeviceRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TriggerConfigurationManagerAction provides operations to call the triggerConfigurationManagerAction method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*DeviceManagementManagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*DeviceManagementManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) Users()(*DeviceManagementManagedDevicesItemUsersRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*DeviceManagementManagedDevicesItemWindowsDefenderScanRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*DeviceManagementManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*DeviceManagementManagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Wipe provides operations to call the wipe method.
func (m *DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) Wipe()(*DeviceManagementManagedDevicesItemWipeRequestBuilder) {
    return NewDeviceManagementManagedDevicesItemWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
