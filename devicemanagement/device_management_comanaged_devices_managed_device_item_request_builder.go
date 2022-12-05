package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
type DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters the list of co-managed devices report
type DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateDeviceEsim provides operations to call the activateDeviceEsim method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*DeviceManagementComanagedDevicesItemActivateDeviceEsimRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*DeviceManagementComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetailsById provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetailsById(id string)(*DeviceManagementComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignmentFilterEvaluationStatusDetails%2Did"] = id
    }
    return NewDeviceManagementComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*DeviceManagementComanagedDevicesItemBypassActivationLockRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*DeviceManagementComanagedDevicesItemCleanWindowsDeviceRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewDeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) {
    m := &DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder{
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
// NewDeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewDeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property comanagedDevices for deviceManagement
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*DeviceManagementComanagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of co-managed devices report
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property comanagedDevices in deviceManagement
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) CreateRemoteHelpSession()(*DeviceManagementComanagedDevicesItemCreateRemoteHelpSessionRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property comanagedDevices for deviceManagement
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*DeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Deprovision provides operations to call the deprovision method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) Deprovision()(*DeviceManagementComanagedDevicesItemDeprovisionRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) DetectedApps()(*DeviceManagementComanagedDevicesItemDetectedAppsRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedAppsById provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) DetectedAppsById(id string)(*DeviceManagementComanagedDevicesItemDetectedAppsDetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp%2Did"] = id
    }
    return NewDeviceManagementComanagedDevicesItemDetectedAppsDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*DeviceManagementComanagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*DeviceManagementComanagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStatesById provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*DeviceManagementComanagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = id
    }
    return NewDeviceManagementComanagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*DeviceManagementComanagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationStatesById provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*DeviceManagementComanagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState%2Did"] = id
    }
    return NewDeviceManagementComanagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Disable provides operations to call the disable method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) Disable()(*DeviceManagementComanagedDevicesItemDisableRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DisableLostMode provides operations to call the disableLostMode method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*DeviceManagementComanagedDevicesItemDisableLostModeRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnableLostMode provides operations to call the enableLostMode method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*DeviceManagementComanagedDevicesItemEnableLostModeRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndRemoteHelpSession provides operations to call the endRemoteHelpSession method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) EndRemoteHelpSession()(*DeviceManagementComanagedDevicesItemEndRemoteHelpSessionRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemEndRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollNowAction provides operations to call the enrollNowAction method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*DeviceManagementComanagedDevicesItemEnrollNowActionRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemEnrollNowActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of co-managed devices report
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*DeviceManagementComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*DeviceManagementComanagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*DeviceManagementComanagedDevicesItemGetFileVaultKeyRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*DeviceManagementComanagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty provides operations to call the getOemWarranty method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) GetOemWarranty()(*DeviceManagementComanagedDevicesItemGetOemWarrantyRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*DeviceManagementComanagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocateDevice provides operations to call the locateDevice method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) LocateDevice()(*DeviceManagementComanagedDevicesItemLocateDeviceRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*DeviceManagementComanagedDevicesItemLogCollectionRequestsRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequestsById provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequestsById(id string)(*DeviceManagementComanagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceLogCollectionResponse%2Did"] = id
    }
    return NewDeviceManagementComanagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*DeviceManagementComanagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*DeviceManagementComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStatesById provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStatesById(id string)(*DeviceManagementComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState%2Did"] = id
    }
    return NewDeviceManagementComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*DeviceManagementComanagedDevicesItemOverrideComplianceStateRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property comanagedDevices in deviceManagement
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*DeviceManagementComanagedDevicesItemPlayLostModeSoundRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RebootNow provides operations to call the rebootNow method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) RebootNow()(*DeviceManagementComanagedDevicesItemRebootNowRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecoverPasscode provides operations to call the recoverPasscode method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*DeviceManagementComanagedDevicesItemRecoverPasscodeRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reenable provides operations to call the reenable method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) Reenable()(*DeviceManagementComanagedDevicesItemReenableRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteLock provides operations to call the remoteLock method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) RemoteLock()(*DeviceManagementComanagedDevicesItemRemoteLockRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*DeviceManagementComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*DeviceManagementComanagedDevicesItemReprovisionCloudPcRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*DeviceManagementComanagedDevicesItemRequestRemoteAssistanceRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteHelpSessionAccess provides operations to call the requestRemoteHelpSessionAccess method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteHelpSessionAccess()(*DeviceManagementComanagedDevicesItemRequestRemoteHelpSessionAccessRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetPasscode provides operations to call the resetPasscode method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*DeviceManagementComanagedDevicesItemResetPasscodeRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResizeCloudPc provides operations to call the resizeCloudPc method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) ResizeCloudPc()(*DeviceManagementComanagedDevicesItemResizeCloudPcRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RestoreCloudPc provides operations to call the restoreCloudPc method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) RestoreCloudPc()(*DeviceManagementComanagedDevicesItemRestoreCloudPcRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Retire provides operations to call the retire method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) Retire()(*DeviceManagementComanagedDevicesItemRetireRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) RetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*DeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey);
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*DeviceManagementComanagedDevicesItemRevokeAppleVppLicensesRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*DeviceManagementComanagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*DeviceManagementComanagedDevicesItemRotateFileVaultKeyRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*DeviceManagementComanagedDevicesItemSecurityBaselineStatesRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStatesById provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStatesById(id string)(*DeviceManagementComanagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineState%2Did"] = id
    }
    return NewDeviceManagementComanagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*DeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*DeviceManagementComanagedDevicesItemSetCloudPcReviewStatusRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetDeviceName provides operations to call the setDeviceName method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*DeviceManagementComanagedDevicesItemSetDeviceNameRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShutDown provides operations to call the shutDown method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) ShutDown()(*DeviceManagementComanagedDevicesItemShutDownRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncDevice provides operations to call the syncDevice method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) SyncDevice()(*DeviceManagementComanagedDevicesItemSyncDeviceRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TriggerConfigurationManagerAction provides operations to call the triggerConfigurationManagerAction method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*DeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*DeviceManagementComanagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) Users()(*DeviceManagementComanagedDevicesItemUsersRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*DeviceManagementComanagedDevicesItemWindowsDefenderScanRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*DeviceManagementComanagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*DeviceManagementComanagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Wipe provides operations to call the wipe method.
func (m *DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) Wipe()(*DeviceManagementComanagedDevicesItemWipeRequestBuilder) {
    return NewDeviceManagementComanagedDevicesItemWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
