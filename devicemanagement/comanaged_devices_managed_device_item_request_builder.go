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
    idPtr := &id
    return NewComanagedDevicesItemAssignmentFilterEvaluationStatusDetailsAssignmentFilterEvaluationStatusDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// NewComanagedDevicesManagedDeviceItemRequestBuilderInternal instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewComanagedDevicesManagedDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, managedDeviceId *string)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    m := &ComanagedDevicesManagedDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if managedDeviceId != nil {
        urlTplParams["managedDevice%2Did"] = *managedDeviceId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewComanagedDevicesManagedDeviceItemRequestBuilder instantiates a new ManagedDeviceItemRequestBuilder and sets the default values.
func NewComanagedDevicesManagedDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesManagedDeviceItemRequestBuilderInternal(urlParams, requestAdapter, nil)
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
    idPtr := &id
    return NewComanagedDevicesItemDetectedAppsDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewComanagedDevicesItemDeviceCompliancePolicyStatesDeviceCompliancePolicyStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewComanagedDevicesItemDeviceConfigurationStatesDeviceConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewComanagedDevicesItemLogCollectionRequestsDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewComanagedDevicesItemManagedDeviceMobileAppConfigurationStatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MicrosoftGraphActivateDeviceEsim provides operations to call the activateDeviceEsim method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphActivateDeviceEsim()(*ComanagedDevicesItemMicrosoftGraphActivateDeviceEsimActivateDeviceEsimRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphActivateDeviceEsimActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphBypassActivationLock provides operations to call the bypassActivationLock method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphBypassActivationLock()(*ComanagedDevicesItemMicrosoftGraphBypassActivationLockBypassActivationLockRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphBypassActivationLockBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCleanWindowsDevice provides operations to call the cleanWindowsDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCleanWindowsDevice()(*ComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDeviceRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphCleanWindowsDeviceCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateDeviceLogCollectionRequest provides operations to call the createDeviceLogCollectionRequest method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCreateDeviceLogCollectionRequest()(*ComanagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphCreateDeviceLogCollectionRequestCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateRemoteHelpSession provides operations to call the createRemoteHelpSession method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphCreateRemoteHelpSession()(*ComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphCreateRemoteHelpSessionCreateRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeleteUserFromSharedAppleDevice provides operations to call the deleteUserFromSharedAppleDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDeleteUserFromSharedAppleDevice()(*ComanagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDeviceRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeprovision provides operations to call the deprovision method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDeprovision()(*ComanagedDevicesItemMicrosoftGraphDeprovisionDeprovisionRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphDeprovisionDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDisable provides operations to call the disable method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDisable()(*ComanagedDevicesItemMicrosoftGraphDisableDisableRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphDisableDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDisableLostMode provides operations to call the disableLostMode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphDisableLostMode()(*ComanagedDevicesItemMicrosoftGraphDisableLostModeDisableLostModeRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphDisableLostModeDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEnableLostMode provides operations to call the enableLostMode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphEnableLostMode()(*ComanagedDevicesItemMicrosoftGraphEnableLostModeEnableLostModeRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphEnableLostModeEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEndRemoteHelpSession provides operations to call the endRemoteHelpSession method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphEndRemoteHelpSession()(*ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionEndRemoteHelpSessionRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionEndRemoteHelpSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEnrollNowAction provides operations to call the enrollNowAction method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphEnrollNowAction()(*ComanagedDevicesItemMicrosoftGraphEnrollNowActionEnrollNowActionRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphEnrollNowActionEnrollNowActionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcRemoteActionResults provides operations to call the getCloudPcRemoteActionResults method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetCloudPcRemoteActionResults()(*ComanagedDevicesItemMicrosoftGraphGetCloudPcRemoteActionResultsGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphGetCloudPcRemoteActionResultsGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcReviewStatus provides operations to call the getCloudPcReviewStatus method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetCloudPcReviewStatus()(*ComanagedDevicesItemMicrosoftGraphGetCloudPcReviewStatusGetCloudPcReviewStatusRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphGetCloudPcReviewStatusGetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetFileVaultKey provides operations to call the getFileVaultKey method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetFileVaultKey()(*ComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetNonCompliantSettings provides operations to call the getNonCompliantSettings method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetNonCompliantSettings()(*ComanagedDevicesItemMicrosoftGraphGetNonCompliantSettingsGetNonCompliantSettingsRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphGetNonCompliantSettingsGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOemWarranty provides operations to call the getOemWarranty method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphGetOemWarranty()(*ComanagedDevicesItemMicrosoftGraphGetOemWarrantyGetOemWarrantyRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphGetOemWarrantyGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphInitiateMobileDeviceManagementKeyRecovery provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphInitiateMobileDeviceManagementKeyRecovery()(*ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphInitiateOnDemandProactiveRemediation provides operations to call the initiateOnDemandProactiveRemediation method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphInitiateOnDemandProactiveRemediation()(*ComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphInitiateOnDemandProactiveRemediationInitiateOnDemandProactiveRemediationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLocateDevice provides operations to call the locateDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphLocateDevice()(*ComanagedDevicesItemMicrosoftGraphLocateDeviceLocateDeviceRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphLocateDeviceLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphLogoutSharedAppleDeviceActiveUser provides operations to call the logoutSharedAppleDeviceActiveUser method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphLogoutSharedAppleDeviceActiveUser()(*ComanagedDevicesItemMicrosoftGraphLogoutSharedAppleDeviceActiveUserLogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphLogoutSharedAppleDeviceActiveUserLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphOverrideComplianceState provides operations to call the overrideComplianceState method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphOverrideComplianceState()(*ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStateRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphPlayLostModeSound provides operations to call the playLostModeSound method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphPlayLostModeSound()(*ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRebootNow provides operations to call the rebootNow method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRebootNow()(*ComanagedDevicesItemMicrosoftGraphRebootNowRebootNowRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRebootNowRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRecoverPasscode provides operations to call the recoverPasscode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRecoverPasscode()(*ComanagedDevicesItemMicrosoftGraphRecoverPasscodeRecoverPasscodeRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRecoverPasscodeRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReenable provides operations to call the reenable method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphReenable()(*ComanagedDevicesItemMicrosoftGraphReenableReenableRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphReenableReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoteLock provides operations to call the remoteLock method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRemoteLock()(*ComanagedDevicesItemMicrosoftGraphRemoteLockRemoteLockRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRemoteLockRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagement provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagement()(*ComanagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReprovisionCloudPc provides operations to call the reprovisionCloudPc method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphReprovisionCloudPc()(*ComanagedDevicesItemMicrosoftGraphReprovisionCloudPcReprovisionCloudPcRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphReprovisionCloudPcReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRequestRemoteAssistance provides operations to call the requestRemoteAssistance method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRequestRemoteAssistance()(*ComanagedDevicesItemMicrosoftGraphRequestRemoteAssistanceRequestRemoteAssistanceRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRequestRemoteAssistanceRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRequestRemoteHelpSessionAccess provides operations to call the requestRemoteHelpSessionAccess method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRequestRemoteHelpSessionAccess()(*ComanagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRequestRemoteHelpSessionAccessRequestRemoteHelpSessionAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResetPasscode provides operations to call the resetPasscode method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphResetPasscode()(*ComanagedDevicesItemMicrosoftGraphResetPasscodeResetPasscodeRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphResetPasscodeResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphResizeCloudPc provides operations to call the resizeCloudPc method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphResizeCloudPc()(*ComanagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphResizeCloudPcResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestoreCloudPc provides operations to call the restoreCloudPc method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRestoreCloudPc()(*ComanagedDevicesItemMicrosoftGraphRestoreCloudPcRestoreCloudPcRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRestoreCloudPcRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetire provides operations to call the retire method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRetire()(*ComanagedDevicesItemMicrosoftGraphRetireRetireRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRetireRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetrieveRemoteHelpSessionWithSessionKey provides operations to call the retrieveRemoteHelpSession method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRetrieveRemoteHelpSessionWithSessionKey(sessionKey *string)(*ComanagedDevicesItemMicrosoftGraphRetrieveRemoteHelpSessionWithSessionKeyRetrieveRemoteHelpSessionWithSessionKeyRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRetrieveRemoteHelpSessionWithSessionKeyRetrieveRemoteHelpSessionWithSessionKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter, sessionKey)
}
// MicrosoftGraphRevokeAppleVppLicenses provides operations to call the revokeAppleVppLicenses method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRevokeAppleVppLicenses()(*ComanagedDevicesItemMicrosoftGraphRevokeAppleVppLicensesRevokeAppleVppLicensesRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRevokeAppleVppLicensesRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRotateBitLockerKeys provides operations to call the rotateBitLockerKeys method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRotateBitLockerKeys()(*ComanagedDevicesItemMicrosoftGraphRotateBitLockerKeysRotateBitLockerKeysRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRotateBitLockerKeysRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRotateFileVaultKey provides operations to call the rotateFileVaultKey method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphRotateFileVaultKey()(*ComanagedDevicesItemMicrosoftGraphRotateFileVaultKeyRotateFileVaultKeyRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphRotateFileVaultKeyRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSendCustomNotificationToCompanyPortal()(*ComanagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetCloudPcReviewStatus provides operations to call the setCloudPcReviewStatus method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSetCloudPcReviewStatus()(*ComanagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphSetCloudPcReviewStatusSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetDeviceName provides operations to call the setDeviceName method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSetDeviceName()(*ComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNameRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphSetDeviceNameSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphShutDown provides operations to call the shutDown method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphShutDown()(*ComanagedDevicesItemMicrosoftGraphShutDownShutDownRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphShutDownShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSyncDevice provides operations to call the syncDevice method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphSyncDevice()(*ComanagedDevicesItemMicrosoftGraphSyncDeviceSyncDeviceRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphSyncDeviceSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTriggerConfigurationManagerAction provides operations to call the triggerConfigurationManagerAction method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphTriggerConfigurationManagerAction()(*ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUpdateWindowsDeviceAccount provides operations to call the updateWindowsDeviceAccount method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphUpdateWindowsDeviceAccount()(*ComanagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphUpdateWindowsDeviceAccountUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWindowsDefenderScan provides operations to call the windowsDefenderScan method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWindowsDefenderScan()(*ComanagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphWindowsDefenderScanWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWindowsDefenderUpdateSignatures provides operations to call the windowsDefenderUpdateSignatures method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWindowsDefenderUpdateSignatures()(*ComanagedDevicesItemMicrosoftGraphWindowsDefenderUpdateSignaturesWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphWindowsDefenderUpdateSignaturesWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphWipe provides operations to call the wipe method.
func (m *ComanagedDevicesManagedDeviceItemRequestBuilder) MicrosoftGraphWipe()(*ComanagedDevicesItemMicrosoftGraphWipeWipeRequestBuilder) {
    return NewComanagedDevicesItemMicrosoftGraphWipeWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
    idPtr := &id
    return NewComanagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
