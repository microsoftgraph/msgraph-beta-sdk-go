package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder provides operations to manage the deviceStatuses property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
type WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetQueryParameters the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
type WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetQueryParameters
}
// WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderInternal instantiates a new WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder and sets the default values.
func NewWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) {
    m := &WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/wdacSupplementalPolicies/{windowsDefenderApplicationControlSupplementalPolicy%2Did}/deviceStatuses/{windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder instantiates a new WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder and sets the default values.
func NewWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceStatuses for deviceAppManagement
func (m *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable), nil
}
// Patch update the navigation property deviceStatuses in deviceAppManagement
func (m *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable, requestConfiguration *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable), nil
}
// Policy provides operations to manage the policy property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus entity.
func (m *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) Policy()(*WdacSupplementalPoliciesItemDeviceStatusesItemPolicyRequestBuilder) {
    return NewWdacSupplementalPoliciesItemDeviceStatusesItemPolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property deviceStatuses for deviceAppManagement
func (m *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPatchRequestInformation update the navigation property deviceStatuses in deviceAppManagement
func (m *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable, requestConfiguration *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) WithUrl(rawUrl string)(*WdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) {
    return NewWdacSupplementalPoliciesItemDeviceStatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
