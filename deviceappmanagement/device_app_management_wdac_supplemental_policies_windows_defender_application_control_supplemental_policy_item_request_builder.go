package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder provides operations to manage the wdacSupplementalPolicies property of the microsoft.graph.deviceAppManagement entity.
type DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetQueryParameters the collection of Windows Defender Application Control Supplemental Policies.
type DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetQueryParameters
}
// DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Assign()(*DeviceAppManagementWdacSupplementalPoliciesItemAssignRequestBuilder) {
    return NewDeviceAppManagementWdacSupplementalPoliciesItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
func (m *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Assignments()(*DeviceAppManagementWdacSupplementalPoliciesItemAssignmentsRequestBuilder) {
    return NewDeviceAppManagementWdacSupplementalPoliciesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
func (m *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) AssignmentsById(id string)(*DeviceAppManagementWdacSupplementalPoliciesItemAssignmentsWindowsDefenderApplicationControlSupplementalPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDefenderApplicationControlSupplementalPolicyAssignment%2Did"] = id
    }
    return NewDeviceAppManagementWdacSupplementalPoliciesItemAssignmentsWindowsDefenderApplicationControlSupplementalPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal instantiates a new WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder and sets the default values.
func NewDeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) {
    m := &DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/wdacSupplementalPolicies/{windowsDefenderApplicationControlSupplementalPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder instantiates a new WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder and sets the default values.
func NewDeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property wdacSupplementalPolicies for deviceAppManagement
func (m *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of Windows Defender Application Control Supplemental Policies.
func (m *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property wdacSupplementalPolicies in deviceAppManagement
func (m *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable, requestConfiguration *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property wdacSupplementalPolicies for deviceAppManagement
func (m *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeploySummary provides operations to manage the deploySummary property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
func (m *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) DeploySummary()(*DeviceAppManagementWdacSupplementalPoliciesItemDeploySummaryRequestBuilder) {
    return NewDeviceAppManagementWdacSupplementalPoliciesItemDeploySummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
func (m *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) DeviceStatuses()(*DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesRequestBuilder) {
    return NewDeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById provides operations to manage the deviceStatuses property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
func (m *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) DeviceStatusesById(id string)(*DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus%2Did"] = id
    }
    return NewDeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the collection of Windows Defender Application Control Supplemental Policies.
func (m *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDefenderApplicationControlSupplementalPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable), nil
}
// Patch update the navigation property wdacSupplementalPolicies in deviceAppManagement
func (m *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable, requestConfiguration *DeviceAppManagementWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDefenderApplicationControlSupplementalPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable), nil
}
