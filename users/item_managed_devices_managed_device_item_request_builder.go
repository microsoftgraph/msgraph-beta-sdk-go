package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ItemManagedDevicesManagedDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ActivateDeviceEsim()(*ItemManagedDevicesItemActivateDeviceEsimRequestBuilder) {
    return NewItemManagedDevicesItemActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*ItemManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewItemManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentFilterEvaluationStatusDetailsById provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetailsById(id string)(*ItemManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignmentFilterEvaluationStatusDetails%2Did"] = id
    }
    return NewItemManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// BypassActivationLock provides operations to call the bypassActivationLock method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) BypassActivationLock()(*ItemManagedDevicesItemBypassActivationLockRequestBuilder) {
    return NewItemManagedDevicesItemBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) CleanWindowsDevice()(*ItemManagedDevicesItemCleanWindowsDeviceRequestBuilder) {
    return NewItemManagedDevicesItemCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesManagedDeviceItemRequestBuilder) {
    m := &ItemManagedDevicesManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemManagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewItemManagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) CreateDeviceLogCollectionRequest()(*ItemManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewItemManagedDevicesItemCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CreateRemoteHelpSession provides operations to call the createRemoteHelpSession method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) CreateRemoteHelpSession()(*ItemManagedDevicesItemCreateRemoteHelpSessionRequestBuilder) {
    return NewItemManagedDevicesItemCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property managedDevices for users
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeleteUserFromSharedAppleDevice provides operations to call the deleteUserFromSharedAppleDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeleteUserFromSharedAppleDevice()(*ItemManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewItemManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Deprovision provides operations to call the deprovision method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Deprovision()(*ItemManagedDevicesItemDeprovisionRequestBuilder) {
    return NewItemManagedDevicesItemDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DetectedApps()(*ItemManagedDevicesItemDetectedAppsRequestBuilder) {
    return NewItemManagedDevicesItemDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DetectedAppsById provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DetectedAppsById(id string)(*ItemManagedDevicesItemDetectedAppsDetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp%2Did"] = id
    }
    return NewItemManagedDevicesItemDetectedAppsDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ItemManagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewItemManagedDevicesItemDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ItemManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewItemManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCompliancePolicyStatesById provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*ItemManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = id
    }
    return NewItemManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ItemManagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewItemManagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceConfigurationStatesById provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*ItemManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState%2Did"] = id
    }
    return NewItemManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceHealthScriptStates provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DeviceHealthScriptStates()(*ItemManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) {
    return NewItemManagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Disable provides operations to call the disable method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Disable()(*ItemManagedDevicesItemDisableRequestBuilder) {
    return NewItemManagedDevicesItemDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DisableLostMode provides operations to call the disableLostMode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) DisableLostMode()(*ItemManagedDevicesItemDisableLostModeRequestBuilder) {
    return NewItemManagedDevicesItemDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EnableLostMode provides operations to call the enableLostMode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) EnableLostMode()(*ItemManagedDevicesItemEnableLostModeRequestBuilder) {
    return NewItemManagedDevicesItemEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EndRemoteHelpSession provides operations to call the endRemoteHelpSession method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) EndRemoteHelpSession()(*ItemManagedDevicesItemEndRemoteHelpSessionRequestBuilder) {
    return NewItemManagedDevicesItemEndRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EnrollNowAction provides operations to call the enrollNowAction method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) EnrollNowAction()(*ItemManagedDevicesItemEnrollNowActionRequestBuilder) {
    return NewItemManagedDevicesItemEnrollNowActionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the managed devices associated with the user.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
// GetCloudPcRemoteActionResults provides operations to call the getCloudPcRemoteActionResults method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcRemoteActionResults()(*ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) GetCloudPcReviewStatus()(*ItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilder) {
    return NewItemManagedDevicesItemGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) GetFileVaultKey()(*ItemManagedDevicesItemGetFileVaultKeyRequestBuilder) {
    return NewItemManagedDevicesItemGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) GetNonCompliantSettings()(*ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    return NewItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetOemWarranty provides operations to call the getOemWarranty method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) GetOemWarranty()(*ItemManagedDevicesItemGetOemWarrantyRequestBuilder) {
    return NewItemManagedDevicesItemGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) InitiateMobileDeviceManagementKeyRecovery()(*ItemManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewItemManagedDevicesItemInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// InitiateOnDemandProactiveRemediation provides operations to call the initiateOnDemandProactiveRemediation method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) InitiateOnDemandProactiveRemediation()(*ItemManagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewItemManagedDevicesItemInitiateOnDemandProactiveRemediationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LocateDevice provides operations to call the locateDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) LocateDevice()(*ItemManagedDevicesItemLocateDeviceRequestBuilder) {
    return NewItemManagedDevicesItemLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LogCollectionRequests provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ItemManagedDevicesItemLogCollectionRequestsRequestBuilder) {
    return NewItemManagedDevicesItemLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LogCollectionRequestsById provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequestsById(id string)(*ItemManagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceLogCollectionResponse%2Did"] = id
    }
    return NewItemManagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// LogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ItemManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewItemManagedDevicesItemLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*ItemManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewItemManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedDeviceMobileAppConfigurationStatesById provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStatesById(id string)(*ItemManagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState%2Did"] = id
    }
    return NewItemManagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// OverrideComplianceState provides operations to call the overrideComplianceState method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) OverrideComplianceState()(*ItemManagedDevicesItemOverrideComplianceStateRequestBuilder) {
    return NewItemManagedDevicesItemOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property managedDevices in users
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
// PlayLostModeSound provides operations to call the playLostModeSound method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) PlayLostModeSound()(*ItemManagedDevicesItemPlayLostModeSoundRequestBuilder) {
    return NewItemManagedDevicesItemPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RebootNow provides operations to call the rebootNow method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RebootNow()(*ItemManagedDevicesItemRebootNowRequestBuilder) {
    return NewItemManagedDevicesItemRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RecoverPasscode provides operations to call the recoverPasscode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RecoverPasscode()(*ItemManagedDevicesItemRecoverPasscodeRequestBuilder) {
    return NewItemManagedDevicesItemRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Reenable provides operations to call the reenable method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Reenable()(*ItemManagedDevicesItemReenableRequestBuilder) {
    return NewItemManagedDevicesItemReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RemoteLock provides operations to call the remoteLock method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RemoteLock()(*ItemManagedDevicesItemRemoteLockRequestBuilder) {
    return NewItemManagedDevicesItemRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RemoveDeviceFirmwareConfigurationInterfaceManagement()(*ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ReprovisionCloudPc()(*ItemManagedDevicesItemReprovisionCloudPcRequestBuilder) {
    return NewItemManagedDevicesItemReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteAssistance()(*ItemManagedDevicesItemRequestRemoteAssistanceRequestBuilder) {
    return NewItemManagedDevicesItemRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RequestRemoteHelpSessionAccess provides operations to call the requestRemoteHelpSessionAccess method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RequestRemoteHelpSessionAccess()(*ItemManagedDevicesItemRequestRemoteHelpSessionAccessRequestBuilder) {
    return NewItemManagedDevicesItemRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResetPasscode provides operations to call the resetPasscode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ResetPasscode()(*ItemManagedDevicesItemResetPasscodeRequestBuilder) {
    return NewItemManagedDevicesItemResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResizeCloudPc provides operations to call the resizeCloudPc method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ResizeCloudPc()(*ItemManagedDevicesItemResizeCloudPcRequestBuilder) {
    return NewItemManagedDevicesItemResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RestoreCloudPc provides operations to call the restoreCloudPc method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RestoreCloudPc()(*ItemManagedDevicesItemRestoreCloudPcRequestBuilder) {
    return NewItemManagedDevicesItemRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Retire provides operations to call the retire method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Retire()(*ItemManagedDevicesItemRetireRequestBuilder) {
    return NewItemManagedDevicesItemRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*ItemManagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return NewItemManagedDevicesItemRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey)
}
// RevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RevokeAppleVppLicenses()(*ItemManagedDevicesItemRevokeAppleVppLicensesRequestBuilder) {
    return NewItemManagedDevicesItemRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RotateBitLockerKeys()(*ItemManagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    return NewItemManagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RotateFileVaultKey provides operations to call the rotateFileVaultKey method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RotateFileVaultKey()(*ItemManagedDevicesItemRotateFileVaultKeyRequestBuilder) {
    return NewItemManagedDevicesItemRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RotateLocalAdminPassword provides operations to call the rotateLocalAdminPassword method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) RotateLocalAdminPassword()(*ItemManagedDevicesItemRotateLocalAdminPasswordRequestBuilder) {
    return NewItemManagedDevicesItemRotateLocalAdminPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecurityBaselineStates provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*ItemManagedDevicesItemSecurityBaselineStatesRequestBuilder) {
    return NewItemManagedDevicesItemSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecurityBaselineStatesById provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStatesById(id string)(*ItemManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineState%2Did"] = id
    }
    return NewItemManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) SendCustomNotificationToCompanyPortal()(*ItemManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewItemManagedDevicesItemSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) SetCloudPcReviewStatus()(*ItemManagedDevicesItemSetCloudPcReviewStatusRequestBuilder) {
    return NewItemManagedDevicesItemSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SetDeviceName provides operations to call the setDeviceName method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) SetDeviceName()(*ItemManagedDevicesItemSetDeviceNameRequestBuilder) {
    return NewItemManagedDevicesItemSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ShutDown provides operations to call the shutDown method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ShutDown()(*ItemManagedDevicesItemShutDownRequestBuilder) {
    return NewItemManagedDevicesItemShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SyncDevice provides operations to call the syncDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) SyncDevice()(*ItemManagedDevicesItemSyncDeviceRequestBuilder) {
    return NewItemManagedDevicesItemSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property managedDevices for users
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the managed devices associated with the user.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedDevices in users
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ItemManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
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
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) TriggerConfigurationManagerAction()(*ItemManagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    return NewItemManagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) UpdateWindowsDeviceAccount()(*ItemManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilder) {
    return NewItemManagedDevicesItemUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Users()(*ItemManagedDevicesItemUsersRequestBuilder) {
    return NewItemManagedDevicesItemUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderScan()(*ItemManagedDevicesItemWindowsDefenderScanRequestBuilder) {
    return NewItemManagedDevicesItemWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) WindowsDefenderUpdateSignatures()(*ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ItemManagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewItemManagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Wipe provides operations to call the wipe method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Wipe()(*ItemManagedDevicesItemWipeRequestBuilder) {
    return NewItemManagedDevicesItemWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
