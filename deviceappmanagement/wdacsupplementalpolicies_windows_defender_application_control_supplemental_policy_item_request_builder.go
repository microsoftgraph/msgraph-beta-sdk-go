package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder provides operations to manage the wdacSupplementalPolicies property of the microsoft.graph.deviceAppManagement entity.
type WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetQueryParameters the collection of Windows Defender Application Control Supplemental Policies.
type WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetQueryParameters
}
// WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *WdacsupplementalpoliciesItemAssignRequestBuilder when successful
func (m *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Assign()(*WdacsupplementalpoliciesItemAssignRequestBuilder) {
    return NewWdacsupplementalpoliciesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
// returns a *WdacsupplementalpoliciesItemAssignmentsRequestBuilder when successful
func (m *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Assignments()(*WdacsupplementalpoliciesItemAssignmentsRequestBuilder) {
    return NewWdacsupplementalpoliciesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal instantiates a new WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder and sets the default values.
func NewWdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) {
    m := &WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/wdacSupplementalPolicies/{windowsDefenderApplicationControlSupplementalPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder instantiates a new WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder and sets the default values.
func NewWdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property wdacSupplementalPolicies for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeploySummary provides operations to manage the deploySummary property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
// returns a *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder when successful
func (m *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) DeploySummary()(*WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder) {
    return NewWdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
// returns a *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder when successful
func (m *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) DeviceStatuses()(*WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder) {
    return NewWdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of Windows Defender Application Control Supplemental Policies.
// returns a WindowsDefenderApplicationControlSupplementalPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a WindowsDefenderApplicationControlSupplementalPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable, requestConfiguration *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of Windows Defender Application Control Supplemental Policies.
// returns a *RequestInformation when successful
func (m *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RequestInformation when successful
func (m *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyable, requestConfiguration *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder when successful
func (m *WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) WithUrl(rawUrl string)(*WdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) {
    return NewWdacsupplementalpoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
