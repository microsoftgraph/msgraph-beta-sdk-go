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
// MicrosoftGraphActivateDeviceEsim provides operations to call the activateDeviceEsim method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphActivateDeviceEsim()(*ItemManagedDevicesItemMicrosoftGraphActivateDeviceEsimRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBypassActivationLock provides operations to call the bypassActivationLock method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphBypassActivationLock()(*ItemManagedDevicesItemMicrosoftGraphBypassActivationLockRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCleanWindowsDevice()(*ItemManagedDevicesItemMicrosoftGraphCleanWindowsDeviceRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCreateDeviceLogCollectionRequest()(*ItemManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateRemoteHelpSession provides operations to call the createRemoteHelpSession method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCreateRemoteHelpSession()(*ItemManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeleteUserFromSharedAppleDevice provides operations to call the deleteUserFromSharedAppleDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDeleteUserFromSharedAppleDevice()(*ItemManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeprovision provides operations to call the deprovision method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDeprovision()(*ItemManagedDevicesItemMicrosoftGraphDeprovisionRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDisable provides operations to call the disable method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDisable()(*ItemManagedDevicesItemMicrosoftGraphDisableRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDisableLostMode provides operations to call the disableLostMode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDisableLostMode()(*ItemManagedDevicesItemMicrosoftGraphDisableLostModeRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEnableLostMode provides operations to call the enableLostMode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphEnableLostMode()(*ItemManagedDevicesItemMicrosoftGraphEnableLostModeRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEndRemoteHelpSession provides operations to call the endRemoteHelpSession method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphEndRemoteHelpSession()(*ItemManagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEnrollNowAction provides operations to call the enrollNowAction method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphEnrollNowAction()(*ItemManagedDevicesItemMicrosoftGraphEnrollNowActionRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphEnrollNowActionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcRemoteActionResults provides operations to call the getCloudPcRemoteActionResults method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetCloudPcRemoteActionResults()(*ItemManagedDevicesItemMicrosoftGraphGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetCloudPcReviewStatus()(*ItemManagedDevicesItemMicrosoftGraphGetCloudPcReviewStatusRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetFileVaultKey()(*ItemManagedDevicesItemMicrosoftGraphGetFileVaultKeyRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetNonCompliantSettings()(*ItemManagedDevicesItemMicrosoftGraphGetNonCompliantSettingsRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOemWarranty provides operations to call the getOemWarranty method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetOemWarranty()(*ItemManagedDevicesItemMicrosoftGraphGetOemWarrantyRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphInitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphInitiateMobileDeviceManagementKeyRecovery()(*ItemManagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphInitiateOnDemandProactiveRemediation provides operations to call the initiateOnDemandProactiveRemediation method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphInitiateOnDemandProactiveRemediation()(*ItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLocateDevice provides operations to call the locateDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphLocateDevice()(*ItemManagedDevicesItemMicrosoftGraphLocateDeviceRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphLogoutSharedAppleDeviceActiveUser()(*ItemManagedDevicesItemMicrosoftGraphLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOverrideComplianceState provides operations to call the overrideComplianceState method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphOverrideComplianceState()(*ItemManagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPlayLostModeSound provides operations to call the playLostModeSound method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphPlayLostModeSound()(*ItemManagedDevicesItemMicrosoftGraphPlayLostModeSoundRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRebootNow provides operations to call the rebootNow method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRebootNow()(*ItemManagedDevicesItemMicrosoftGraphRebootNowRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRecoverPasscode provides operations to call the recoverPasscode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRecoverPasscode()(*ItemManagedDevicesItemMicrosoftGraphRecoverPasscodeRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReenable provides operations to call the reenable method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphReenable()(*ItemManagedDevicesItemMicrosoftGraphReenableRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoteLock provides operations to call the remoteLock method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRemoteLock()(*ItemManagedDevicesItemMicrosoftGraphRemoteLockRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagement()(*ItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphReprovisionCloudPc()(*ItemManagedDevicesItemMicrosoftGraphReprovisionCloudPcRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRequestRemoteAssistance()(*ItemManagedDevicesItemMicrosoftGraphRequestRemoteAssistanceRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRequestRemoteHelpSessionAccess provides operations to call the requestRemoteHelpSessionAccess method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRequestRemoteHelpSessionAccess()(*ItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResetPasscode provides operations to call the resetPasscode method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphResetPasscode()(*ItemManagedDevicesItemMicrosoftGraphResetPasscodeRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResizeCloudPc provides operations to call the resizeCloudPc method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphResizeCloudPc()(*ItemManagedDevicesItemMicrosoftGraphResizeCloudPcRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestoreCloudPc provides operations to call the restoreCloudPc method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRestoreCloudPc()(*ItemManagedDevicesItemMicrosoftGraphRestoreCloudPcRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetire provides operations to call the retire method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRetire()(*ItemManagedDevicesItemMicrosoftGraphRetireRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*ItemManagedDevicesItemMicrosoftGraphRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey)
}
// MicrosoftGraphRevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRevokeAppleVppLicenses()(*ItemManagedDevicesItemMicrosoftGraphRevokeAppleVppLicensesRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRotateBitLockerKeys()(*ItemManagedDevicesItemMicrosoftGraphRotateBitLockerKeysRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRotateFileVaultKey provides operations to call the rotateFileVaultKey method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRotateFileVaultKey()(*ItemManagedDevicesItemMicrosoftGraphRotateFileVaultKeyRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSendCustomNotificationToCompanyPortal()(*ItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSetCloudPcReviewStatus()(*ItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetDeviceName provides operations to call the setDeviceName method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSetDeviceName()(*ItemManagedDevicesItemMicrosoftGraphSetDeviceNameRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphShutDown provides operations to call the shutDown method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphShutDown()(*ItemManagedDevicesItemMicrosoftGraphShutDownRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSyncDevice provides operations to call the syncDevice method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSyncDevice()(*ItemManagedDevicesItemMicrosoftGraphSyncDeviceRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTriggerConfigurationManagerAction provides operations to call the triggerConfigurationManagerAction method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphTriggerConfigurationManagerAction()(*ItemManagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphUpdateWindowsDeviceAccount()(*ItemManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWindowsDefenderScan()(*ItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWindowsDefenderUpdateSignatures()(*ItemManagedDevicesItemMicrosoftGraphWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWipe provides operations to call the wipe method.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWipe()(*ItemManagedDevicesItemMicrosoftGraphWipeRequestBuilder) {
    return NewItemManagedDevicesItemMicrosoftGraphWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
// Users provides operations to manage the users property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) Users()(*ItemManagedDevicesItemUsersRequestBuilder) {
    return NewItemManagedDevicesItemUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
func (m *ItemManagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ItemManagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewItemManagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
