package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemManagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type UsersItemManagedDevicesManagedDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters the managed devices associated with the user.
type UsersItemManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// UsersItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateDeviceEsim provides operations to call the activateDeviceEsim method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*UsersItemManagedDevicesItemActivateDeviceEsimRequestBuilder) {
    return NewUsersItemManagedDevicesItemActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*UsersItemManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewUsersItemManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetailsById provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetailsById(id string)(*UsersItemManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignmentFilterEvaluationStatusDetails%2Did"] = id
    }
    return NewUsersItemManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*UsersItemManagedDevicesItemBypassActivationLockRequestBuilder) {
    return NewUsersItemManagedDevicesItemBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*UsersItemManagedDevicesItemCleanWindowsDeviceRequestBuilder) {
    return NewUsersItemManagedDevicesItemCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemManagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewUsersItemManagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemManagedDevicesManagedDeviceItemRequestBuilder) {
    m := &UsersItemManagedDevicesManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemManagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewUsersItemManagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemManagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemManagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedDevices for users
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*UsersItemManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewUsersItemManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the managed devices associated with the user.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property managedDevices in users
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *UsersItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) CreateRemoteHelpSession()(*UsersItemManagedDevicesItemCreateRemoteHelpSessionRequestBuilder) {
    return NewUsersItemManagedDevicesItemCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property managedDevices for users
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*UsersItemManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewUsersItemManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Deprovision provides operations to call the deprovision method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) Deprovision()(*UsersItemManagedDevicesItemDeprovisionRequestBuilder) {
    return NewUsersItemManagedDevicesItemDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) DetectedApps()(*UsersItemManagedDevicesItemDetectedAppsRequestBuilder) {
    return NewUsersItemManagedDevicesItemDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedAppsById provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) DetectedAppsById(id string)(*UsersItemManagedDevicesItemDetectedAppsDetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp%2Did"] = id
    }
    return NewUsersItemManagedDevicesItemDetectedAppsDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*UsersItemManagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewUsersItemManagedDevicesItemDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*UsersItemManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewUsersItemManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStatesById provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*UsersItemManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = id
    }
    return NewUsersItemManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*UsersItemManagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewUsersItemManagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationStatesById provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*UsersItemManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState%2Did"] = id
    }
    return NewUsersItemManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Disable provides operations to call the disable method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) Disable()(*UsersItemManagedDevicesItemDisableRequestBuilder) {
    return NewUsersItemManagedDevicesItemDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DisableLostMode provides operations to call the disableLostMode method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*UsersItemManagedDevicesItemDisableLostModeRequestBuilder) {
    return NewUsersItemManagedDevicesItemDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnableLostMode provides operations to call the enableLostMode method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*UsersItemManagedDevicesItemEnableLostModeRequestBuilder) {
    return NewUsersItemManagedDevicesItemEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndRemoteHelpSession provides operations to call the endRemoteHelpSession method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) EndRemoteHelpSession()(*UsersItemManagedDevicesItemEndRemoteHelpSessionRequestBuilder) {
    return NewUsersItemManagedDevicesItemEndRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollNowAction provides operations to call the enrollNowAction method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*UsersItemManagedDevicesItemEnrollNowActionRequestBuilder) {
    return NewUsersItemManagedDevicesItemEnrollNowActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the managed devices associated with the user.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewUsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*UsersItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    return NewUsersItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*UsersItemManagedDevicesItemGetFileVaultKeyRequestBuilder) {
    return NewUsersItemManagedDevicesItemGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*UsersItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    return NewUsersItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty provides operations to call the getOemWarranty method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) GetOemWarranty()(*UsersItemManagedDevicesItemGetOemWarrantyRequestBuilder) {
    return NewUsersItemManagedDevicesItemGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*UsersItemManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewUsersItemManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocateDevice provides operations to call the locateDevice method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) LocateDevice()(*UsersItemManagedDevicesItemLocateDeviceRequestBuilder) {
    return NewUsersItemManagedDevicesItemLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*UsersItemManagedDevicesItemLogCollectionRequestsRequestBuilder) {
    return NewUsersItemManagedDevicesItemLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequestsById provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequestsById(id string)(*UsersItemManagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceLogCollectionResponse%2Did"] = id
    }
    return NewUsersItemManagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*UsersItemManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewUsersItemManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*UsersItemManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewUsersItemManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStatesById provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStatesById(id string)(*UsersItemManagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState%2Did"] = id
    }
    return NewUsersItemManagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*UsersItemManagedDevicesItemOverrideComplianceStateRequestBuilder) {
    return NewUsersItemManagedDevicesItemOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property managedDevices in users
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *UsersItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*UsersItemManagedDevicesItemPlayLostModeSoundRequestBuilder) {
    return NewUsersItemManagedDevicesItemPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RebootNow provides operations to call the rebootNow method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) RebootNow()(*UsersItemManagedDevicesItemRebootNowRequestBuilder) {
    return NewUsersItemManagedDevicesItemRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecoverPasscode provides operations to call the recoverPasscode method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*UsersItemManagedDevicesItemRecoverPasscodeRequestBuilder) {
    return NewUsersItemManagedDevicesItemRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reenable provides operations to call the reenable method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) Reenable()(*UsersItemManagedDevicesItemReenableRequestBuilder) {
    return NewUsersItemManagedDevicesItemReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteLock provides operations to call the remoteLock method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) RemoteLock()(*UsersItemManagedDevicesItemRemoteLockRequestBuilder) {
    return NewUsersItemManagedDevicesItemRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*UsersItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewUsersItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*UsersItemManagedDevicesItemReprovisionCloudPcRequestBuilder) {
    return NewUsersItemManagedDevicesItemReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*UsersItemManagedDevicesItemRequestRemoteAssistanceRequestBuilder) {
    return NewUsersItemManagedDevicesItemRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestRemoteHelpSessionAccess provides operations to call the requestRemoteHelpSessionAccess method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteHelpSessionAccess()(*UsersItemManagedDevicesItemRequestRemoteHelpSessionAccessRequestBuilder) {
    return NewUsersItemManagedDevicesItemRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetPasscode provides operations to call the resetPasscode method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*UsersItemManagedDevicesItemResetPasscodeRequestBuilder) {
    return NewUsersItemManagedDevicesItemResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResizeCloudPc provides operations to call the resizeCloudPc method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) ResizeCloudPc()(*UsersItemManagedDevicesItemResizeCloudPcRequestBuilder) {
    return NewUsersItemManagedDevicesItemResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RestoreCloudPc provides operations to call the restoreCloudPc method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) RestoreCloudPc()(*UsersItemManagedDevicesItemRestoreCloudPcRequestBuilder) {
    return NewUsersItemManagedDevicesItemRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Retire provides operations to call the retire method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) Retire()(*UsersItemManagedDevicesItemRetireRequestBuilder) {
    return NewUsersItemManagedDevicesItemRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) RetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*UsersItemManagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return NewUsersItemManagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey);
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*UsersItemManagedDevicesItemRevokeAppleVppLicensesRequestBuilder) {
    return NewUsersItemManagedDevicesItemRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*UsersItemManagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    return NewUsersItemManagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*UsersItemManagedDevicesItemRotateFileVaultKeyRequestBuilder) {
    return NewUsersItemManagedDevicesItemRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*UsersItemManagedDevicesItemSecurityBaselineStatesRequestBuilder) {
    return NewUsersItemManagedDevicesItemSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStatesById provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStatesById(id string)(*UsersItemManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineState%2Did"] = id
    }
    return NewUsersItemManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*UsersItemManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewUsersItemManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*UsersItemManagedDevicesItemSetCloudPcReviewStatusRequestBuilder) {
    return NewUsersItemManagedDevicesItemSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetDeviceName provides operations to call the setDeviceName method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*UsersItemManagedDevicesItemSetDeviceNameRequestBuilder) {
    return NewUsersItemManagedDevicesItemSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShutDown provides operations to call the shutDown method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) ShutDown()(*UsersItemManagedDevicesItemShutDownRequestBuilder) {
    return NewUsersItemManagedDevicesItemShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncDevice provides operations to call the syncDevice method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) SyncDevice()(*UsersItemManagedDevicesItemSyncDeviceRequestBuilder) {
    return NewUsersItemManagedDevicesItemSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TriggerConfigurationManagerAction provides operations to call the triggerConfigurationManagerAction method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*UsersItemManagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    return NewUsersItemManagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*UsersItemManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder) {
    return NewUsersItemManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) Users()(*UsersItemManagedDevicesItemUsersRequestBuilder) {
    return NewUsersItemManagedDevicesItemUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*UsersItemManagedDevicesItemWindowsDefenderScanRequestBuilder) {
    return NewUsersItemManagedDevicesItemWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*UsersItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewUsersItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*UsersItemManagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewUsersItemManagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Wipe provides operations to call the wipe method.
func (m *UsersItemManagedDevicesManagedDeviceItemRequestBuilder) Wipe()(*UsersItemManagedDevicesItemWipeRequestBuilder) {
    return NewUsersItemManagedDevicesItemWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
