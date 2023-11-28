package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder provides operations to manage the wdacSupplementalPolicies property of the microsoft.graph.deviceAppManagement entity.
type WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetQueryParameters the collection of Windows Defender Application Control Supplemental Policies.
type WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetQueryParameters
}
// WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Assign()(*WdacSupplementalPoliciesItemAssignRequestBuilder) {
    return NewWdacSupplementalPoliciesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
func (m *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Assignments()(*WdacSupplementalPoliciesItemAssignmentsRequestBuilder) {
    return NewWdacSupplementalPoliciesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal instantiates a new WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder and sets the default values.
func NewWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) {
    m := &WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/wdacSupplementalPolicies/{windowsDefenderApplicationControlSupplementalPolicy%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder instantiates a new WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder and sets the default values.
func NewWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property wdacSupplementalPolicies for deviceAppManagement
func (m *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeploySummary provides operations to manage the deploySummary property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
func (m *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) DeploySummary()(*WdacSupplementalPoliciesItemDeploySummaryRequestBuilder) {
    return NewWdacSupplementalPoliciesItemDeploySummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
func (m *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) DeviceStatuses()(*WdacSupplementalPoliciesItemDeviceStatusesRequestBuilder) {
    return NewWdacSupplementalPoliciesItemDeviceStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of Windows Defender Application Control Supplemental Policies.
func (m *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDefenderApplicationControlSupplementalPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable), nil
}
// Patch update the navigation property wdacSupplementalPolicies in deviceAppManagement
func (m *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable, requestConfiguration *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDefenderApplicationControlSupplementalPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable), nil
}
// ToDeleteRequestInformation delete navigation property wdacSupplementalPolicies for deviceAppManagement
func (m *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of Windows Defender Application Control Supplemental Policies.
func (m *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property wdacSupplementalPolicies in deviceAppManagement
func (m *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable, requestConfiguration *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) WithUrl(rawUrl string)(*WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) {
    return NewWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
