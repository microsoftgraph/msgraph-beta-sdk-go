package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder provides operations to manage the deviceStatuses property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
type DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetQueryParameters the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
type DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetQueryParameters
}
// DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderInternal instantiates a new WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder and sets the default values.
func NewDeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) {
    m := &DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/wdacSupplementalPolicies/{windowsDefenderApplicationControlSupplementalPolicy%2Did}/deviceStatuses/{windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder instantiates a new WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder and sets the default values.
func NewDeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceStatuses for deviceAppManagement
func (m *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceStatuses in deviceAppManagement
func (m *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable, requestConfiguration *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceStatuses for deviceAppManagement
func (m *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable), nil
}
// Patch update the navigation property deviceStatuses in deviceAppManagement
func (m *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable, requestConfiguration *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable), nil
}
// Policy provides operations to manage the policy property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus entity.
func (m *DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) Policy()(*DeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesItemPolicyRequestBuilder) {
    return NewDeviceAppManagementWdacSupplementalPoliciesItemDeviceStatusesItemPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
