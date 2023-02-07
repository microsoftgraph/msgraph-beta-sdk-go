package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesManagedDeviceItemRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ManagedDevicesManagedDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters the managed devices associated with the user.
type ManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesManagedDeviceItemRequestBuilderGetQueryParameters
}
// ManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignmentFilterEvaluationStatusDetails provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return NewManagedDevicesItemAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentFilterEvaluationStatusDetailsById provides operations to manage the assignmentFilterEvaluationStatusDetails property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) AssignmentFilterEvaluationStatusDetailsById(id string)(*ManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignmentFilterEvaluationStatusDetails%2Did"] = id
    }
    return NewManagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewManagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesManagedDeviceItemRequestBuilder) {
    m := &ManagedDevicesManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewManagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewManagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedDevices for me
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DetectedApps()(*ManagedDevicesItemDetectedAppsRequestBuilder) {
    return NewManagedDevicesItemDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DetectedAppsById provides operations to manage the detectedApps property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DetectedAppsById(id string)(*ManagedDevicesItemDetectedAppsDetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp%2Did"] = id
    }
    return NewManagedDevicesItemDetectedAppsDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceCategory provides operations to manage the deviceCategory property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceCategory()(*ManagedDevicesItemDeviceCategoryRequestBuilder) {
    return NewManagedDevicesItemDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCompliancePolicyStates provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStates()(*ManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilder) {
    return NewManagedDevicesItemDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCompliancePolicyStatesById provides operations to manage the deviceCompliancePolicyStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*ManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState%2Did"] = id
    }
    return NewManagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceConfigurationStates provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStates()(*ManagedDevicesItemDeviceConfigurationStatesRequestBuilder) {
    return NewManagedDevicesItemDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceConfigurationStatesById provides operations to manage the deviceConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceConfigurationStatesById(id string)(*ManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState%2Did"] = id
    }
    return NewManagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceHealthScriptStates provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) DeviceHealthScriptStates()(*ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) {
    return NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the managed devices associated with the user.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequests()(*ManagedDevicesItemLogCollectionRequestsRequestBuilder) {
    return NewManagedDevicesItemLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// LogCollectionRequestsById provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) LogCollectionRequestsById(id string)(*ManagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceLogCollectionResponse%2Did"] = id
    }
    return NewManagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ManagedDeviceMobileAppConfigurationStates provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*ManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return NewManagedDevicesItemManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedDeviceMobileAppConfigurationStatesById provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ManagedDeviceMobileAppConfigurationStatesById(id string)(*ManagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState%2Did"] = id
    }
    return NewManagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphActivateDeviceEsim provides operations to call the activateDeviceEsim method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphActivateDeviceEsim()(*ManagedDevicesItemMicrosoftGraphActivateDeviceEsimRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBypassActivationLock provides operations to call the bypassActivationLock method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphBypassActivationLock()(*ManagedDevicesItemMicrosoftGraphBypassActivationLockRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCleanWindowsDevice()(*ManagedDevicesItemMicrosoftGraphCleanWindowsDeviceRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCreateDeviceLogCollectionRequest()(*ManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateRemoteHelpSession provides operations to call the createRemoteHelpSession method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCreateRemoteHelpSession()(*ManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeleteUserFromSharedAppleDevice provides operations to call the deleteUserFromSharedAppleDevice method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDeleteUserFromSharedAppleDevice()(*ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeprovision provides operations to call the deprovision method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDeprovision()(*ManagedDevicesItemMicrosoftGraphDeprovisionRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDisable provides operations to call the disable method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDisable()(*ManagedDevicesItemMicrosoftGraphDisableRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDisableLostMode provides operations to call the disableLostMode method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDisableLostMode()(*ManagedDevicesItemMicrosoftGraphDisableLostModeRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEnableLostMode provides operations to call the enableLostMode method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphEnableLostMode()(*ManagedDevicesItemMicrosoftGraphEnableLostModeRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEndRemoteHelpSession provides operations to call the endRemoteHelpSession method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphEndRemoteHelpSession()(*ManagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEnrollNowAction provides operations to call the enrollNowAction method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphEnrollNowAction()(*ManagedDevicesItemMicrosoftGraphEnrollNowActionRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphEnrollNowActionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcRemoteActionResults provides operations to call the getCloudPcRemoteActionResults method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetCloudPcRemoteActionResults()(*ManagedDevicesItemMicrosoftGraphGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetCloudPcReviewStatus()(*ManagedDevicesItemMicrosoftGraphGetCloudPcReviewStatusRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetFileVaultKey()(*ManagedDevicesItemMicrosoftGraphGetFileVaultKeyRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetNonCompliantSettings()(*ManagedDevicesItemMicrosoftGraphGetNonCompliantSettingsRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOemWarranty provides operations to call the getOemWarranty method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetOemWarranty()(*ManagedDevicesItemMicrosoftGraphGetOemWarrantyRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphInitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphInitiateMobileDeviceManagementKeyRecovery()(*ManagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphInitiateOnDemandProactiveRemediation provides operations to call the initiateOnDemandProactiveRemediation method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphInitiateOnDemandProactiveRemediation()(*ManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLocateDevice provides operations to call the locateDevice method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphLocateDevice()(*ManagedDevicesItemMicrosoftGraphLocateDeviceRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphLogoutSharedAppleDeviceActiveUser()(*ManagedDevicesItemMicrosoftGraphLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOverrideComplianceState provides operations to call the overrideComplianceState method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphOverrideComplianceState()(*ManagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPlayLostModeSound provides operations to call the playLostModeSound method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphPlayLostModeSound()(*ManagedDevicesItemMicrosoftGraphPlayLostModeSoundRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRebootNow provides operations to call the rebootNow method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRebootNow()(*ManagedDevicesItemMicrosoftGraphRebootNowRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRecoverPasscode provides operations to call the recoverPasscode method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRecoverPasscode()(*ManagedDevicesItemMicrosoftGraphRecoverPasscodeRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReenable provides operations to call the reenable method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphReenable()(*ManagedDevicesItemMicrosoftGraphReenableRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoteLock provides operations to call the remoteLock method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRemoteLock()(*ManagedDevicesItemMicrosoftGraphRemoteLockRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagement()(*ManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphReprovisionCloudPc()(*ManagedDevicesItemMicrosoftGraphReprovisionCloudPcRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRequestRemoteAssistance()(*ManagedDevicesItemMicrosoftGraphRequestRemoteAssistanceRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRequestRemoteHelpSessionAccess provides operations to call the requestRemoteHelpSessionAccess method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRequestRemoteHelpSessionAccess()(*ManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResetPasscode provides operations to call the resetPasscode method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphResetPasscode()(*ManagedDevicesItemMicrosoftGraphResetPasscodeRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResizeCloudPc provides operations to call the resizeCloudPc method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphResizeCloudPc()(*ManagedDevicesItemMicrosoftGraphResizeCloudPcRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestoreCloudPc provides operations to call the restoreCloudPc method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRestoreCloudPc()(*ManagedDevicesItemMicrosoftGraphRestoreCloudPcRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetire provides operations to call the retire method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRetire()(*ManagedDevicesItemMicrosoftGraphRetireRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*ManagedDevicesItemMicrosoftGraphRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey)
}
// MicrosoftGraphRevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRevokeAppleVppLicenses()(*ManagedDevicesItemMicrosoftGraphRevokeAppleVppLicensesRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRotateBitLockerKeys()(*ManagedDevicesItemMicrosoftGraphRotateBitLockerKeysRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRotateFileVaultKey provides operations to call the rotateFileVaultKey method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRotateFileVaultKey()(*ManagedDevicesItemMicrosoftGraphRotateFileVaultKeyRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSendCustomNotificationToCompanyPortal()(*ManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSetCloudPcReviewStatus()(*ManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetDeviceName provides operations to call the setDeviceName method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSetDeviceName()(*ManagedDevicesItemMicrosoftGraphSetDeviceNameRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphShutDown provides operations to call the shutDown method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphShutDown()(*ManagedDevicesItemMicrosoftGraphShutDownRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSyncDevice provides operations to call the syncDevice method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSyncDevice()(*ManagedDevicesItemMicrosoftGraphSyncDeviceRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTriggerConfigurationManagerAction provides operations to call the triggerConfigurationManagerAction method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphTriggerConfigurationManagerAction()(*ManagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphUpdateWindowsDeviceAccount()(*ManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWindowsDefenderScan()(*ManagedDevicesItemMicrosoftGraphWindowsDefenderScanRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWindowsDefenderUpdateSignatures()(*ManagedDevicesItemMicrosoftGraphWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWipe provides operations to call the wipe method.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWipe()(*ManagedDevicesItemMicrosoftGraphWipeRequestBuilder) {
    return NewManagedDevicesItemMicrosoftGraphWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property managedDevices in me
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStates()(*ManagedDevicesItemSecurityBaselineStatesRequestBuilder) {
    return NewManagedDevicesItemSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecurityBaselineStatesById provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) SecurityBaselineStatesById(id string)(*ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineState%2Did"] = id
    }
    return NewManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property managedDevices for me
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedDevices in me
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDevicesManagedDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) Users()(*ManagedDevicesItemUsersRequestBuilder) {
    return NewManagedDevicesItemUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsProtectionState provides operations to manage the windowsProtectionState property of the microsoft.graph.managedDevice entity.
func (m *ManagedDevicesManagedDeviceItemRequestBuilder) WindowsProtectionState()(*ManagedDevicesItemWindowsProtectionStateRequestBuilder) {
    return NewManagedDevicesItemWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
