package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder provides operations to manage the windowsQualityUpdatePolicies property of the microsoft.graph.deviceManagement entity.
type WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetQueryParameters a collection of Windows quality update policies
type WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetQueryParameters
}
// WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *WindowsQualityUpdatePoliciesItemAssignRequestBuilder when successful
func (m *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder) Assign()(*WindowsQualityUpdatePoliciesItemAssignRequestBuilder) {
    return NewWindowsQualityUpdatePoliciesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsQualityUpdatePolicy entity.
// returns a *WindowsQualityUpdatePoliciesItemAssignmentsRequestBuilder when successful
func (m *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder) Assignments()(*WindowsQualityUpdatePoliciesItemAssignmentsRequestBuilder) {
    return NewWindowsQualityUpdatePoliciesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderInternal instantiates a new WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder and sets the default values.
func NewWindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder) {
    m := &WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsQualityUpdatePolicies/{windowsQualityUpdatePolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder instantiates a new WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder and sets the default values.
func NewWindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsQualityUpdatePolicies for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of Windows quality update policies
// returns a WindowsQualityUpdatePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsQualityUpdatePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyable), nil
}
// Patch update the navigation property windowsQualityUpdatePolicies in deviceManagement
// returns a WindowsQualityUpdatePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyable, requestConfiguration *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsQualityUpdatePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyable), nil
}
// ToDeleteRequestInformation delete navigation property windowsQualityUpdatePolicies for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of Windows quality update policies
// returns a *RequestInformation when successful
func (m *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsQualityUpdatePolicies in deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyable, requestConfiguration *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder when successful
func (m *WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder) WithUrl(rawUrl string)(*WindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder) {
    return NewWindowsQualityUpdatePoliciesWindowsQualityUpdatePolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
