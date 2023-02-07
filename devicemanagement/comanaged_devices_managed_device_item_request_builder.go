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
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*ComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
    return NewComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
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
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewComanagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewComanagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
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
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DetectedApps()(*ComanagedDevicesItemDetectedAppsRequestBuilder) {
    return NewComanagedDevicesItemDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
    return NewComanagedDevicesItemDetectedAppsDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ComanagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewComanagedDevicesItemDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ComanagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewComanagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
    return NewComanagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ComanagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewComanagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
    return NewComanagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceHealthScriptStates provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) DeviceHealthScriptStates()(*ComanagedDevicesItemDeviceHealthScriptStatesRequestBuilder) {
    return NewComanagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ComanagedDevicesItemLogCollectionRequestsRequestBuilder) {
    return NewComanagedDevicesItemLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
    return NewComanagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*ComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
    return NewComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphActivateDeviceEsim provides operations to call the activateDeviceEsim method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphActivateDeviceEsim()(*ComanagedDevicesItemMicrosoftGraphActivateDeviceEsimRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBypassActivationLock provides operations to call the bypassActivationLock method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphBypassActivationLock()(*ComanagedDevicesItemMicrosoftGraphBypassActivationLockRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCleanWindowsDevice()(*ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCreateDeviceLogCollectionRequest()(*ComanagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateRemoteHelpSession provides operations to call the createRemoteHelpSession method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCreateRemoteHelpSession()(*ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeleteUserFromSharedAppleDevice provides operations to call the deleteUserFromSharedAppleDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDeleteUserFromSharedAppleDevice()(*ComanagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeprovision provides operations to call the deprovision method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDeprovision()(*ComanagedDevicesItemMicrosoftGraphDeprovisionRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDisable provides operations to call the disable method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDisable()(*ComanagedDevicesItemMicrosoftGraphDisableRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDisableLostMode provides operations to call the disableLostMode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDisableLostMode()(*ComanagedDevicesItemMicrosoftGraphDisableLostModeRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEnableLostMode provides operations to call the enableLostMode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphEnableLostMode()(*ComanagedDevicesItemMicrosoftGraphEnableLostModeRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEndRemoteHelpSession provides operations to call the endRemoteHelpSession method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphEndRemoteHelpSession()(*ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEnrollNowAction provides operations to call the enrollNowAction method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphEnrollNowAction()(*ComanagedDevicesItemMicrosoftGraphEnrollNowActionRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphEnrollNowActionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcRemoteActionResults provides operations to call the getCloudPcRemoteActionResults method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetCloudPcRemoteActionResults()(*ComanagedDevicesItemMicrosoftGraphGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetCloudPcReviewStatus()(*ComanagedDevicesItemMicrosoftGraphGetCloudPcReviewStatusRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetFileVaultKey()(*ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetNonCompliantSettings()(*ComanagedDevicesItemMicrosoftGraphGetNonCompliantSettingsRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOemWarranty provides operations to call the getOemWarranty method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetOemWarranty()(*ComanagedDevicesItemMicrosoftGraphGetOemWarrantyRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphInitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphInitiateMobileDeviceManagementKeyRecovery()(*ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphInitiateOnDemandProactiveRemediation provides operations to call the initiateOnDemandProactiveRemediation method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphInitiateOnDemandProactiveRemediation()(*ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLocateDevice provides operations to call the locateDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphLocateDevice()(*ComanagedDevicesItemMicrosoftGraphLocateDeviceRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphLogoutSharedAppleDeviceActiveUser()(*ComanagedDevicesItemMicrosoftGraphLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOverrideComplianceState provides operations to call the overrideComplianceState method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphOverrideComplianceState()(*ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPlayLostModeSound provides operations to call the playLostModeSound method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphPlayLostModeSound()(*ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRebootNow provides operations to call the rebootNow method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRebootNow()(*ComanagedDevicesItemMicrosoftGraphRebootNowRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRecoverPasscode provides operations to call the recoverPasscode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRecoverPasscode()(*ComanagedDevicesItemMicrosoftGraphRecoverPasscodeRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReenable provides operations to call the reenable method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphReenable()(*ComanagedDevicesItemMicrosoftGraphReenableRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoteLock provides operations to call the remoteLock method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRemoteLock()(*ComanagedDevicesItemMicrosoftGraphRemoteLockRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagement()(*ComanagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphReprovisionCloudPc()(*ComanagedDevicesItemMicrosoftGraphReprovisionCloudPcRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRequestRemoteAssistance()(*ComanagedDevicesItemMicrosoftGraphRequestRemoteAssistanceRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRequestRemoteHelpSessionAccess provides operations to call the requestRemoteHelpSessionAccess method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRequestRemoteHelpSessionAccess()(*ComanagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResetPasscode provides operations to call the resetPasscode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphResetPasscode()(*ComanagedDevicesItemMicrosoftGraphResetPasscodeRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResizeCloudPc provides operations to call the resizeCloudPc method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphResizeCloudPc()(*ComanagedDevicesItemMicrosoftGraphResizeCloudPcRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestoreCloudPc provides operations to call the restoreCloudPc method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRestoreCloudPc()(*ComanagedDevicesItemMicrosoftGraphRestoreCloudPcRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetire provides operations to call the retire method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRetire()(*ComanagedDevicesItemMicrosoftGraphRetireRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*ComanagedDevicesItemMicrosoftGraphRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey)
}
// MicrosoftGraphRevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRevokeAppleVppLicenses()(*ComanagedDevicesItemMicrosoftGraphRevokeAppleVppLicensesRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRotateBitLockerKeys()(*ComanagedDevicesItemMicrosoftGraphRotateBitLockerKeysRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRotateFileVaultKey provides operations to call the rotateFileVaultKey method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRotateFileVaultKey()(*ComanagedDevicesItemMicrosoftGraphRotateFileVaultKeyRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSendCustomNotificationToCompanyPortal()(*ComanagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSetCloudPcReviewStatus()(*ComanagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetDeviceName provides operations to call the setDeviceName method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSetDeviceName()(*ComanagedDevicesItemMicrosoftGraphSetDeviceNameRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphShutDown provides operations to call the shutDown method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphShutDown()(*ComanagedDevicesItemMicrosoftGraphShutDownRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSyncDevice provides operations to call the syncDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSyncDevice()(*ComanagedDevicesItemMicrosoftGraphSyncDeviceRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTriggerConfigurationManagerAction provides operations to call the triggerConfigurationManagerAction method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphTriggerConfigurationManagerAction()(*ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphUpdateWindowsDeviceAccount()(*ComanagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWindowsDefenderScan()(*ComanagedDevicesItemMicrosoftGraphWindowsDefenderScanRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWindowsDefenderUpdateSignatures()(*ComanagedDevicesItemMicrosoftGraphWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWipe provides operations to call the wipe method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWipe()(*ComanagedDevicesItemMicrosoftGraphWipeRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*ComanagedDevicesItemSecurityBaselineStatesRequestBuilder) {
    return NewComanagedDevicesItemSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
    return NewComanagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property comanagedDevices for deviceManagement
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the list of co-managed devices report
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property comanagedDevices in deviceManagement
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ComanagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) Users()(*ComanagedDevicesItemUsersRequestBuilder) {
    return NewComanagedDevicesItemUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ComanagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewComanagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
