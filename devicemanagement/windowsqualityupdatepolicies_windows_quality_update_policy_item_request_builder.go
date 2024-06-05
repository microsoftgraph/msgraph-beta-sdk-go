package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder provides operations to manage the windowsQualityUpdatePolicies property of the microsoft.graph.deviceManagement entity.
type WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetQueryParameters a collection of Windows quality update policies
type WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetQueryParameters
}
// WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *WindowsqualityupdatepoliciesItemAssignRequestBuilder when successful
func (m *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder) Assign()(*WindowsqualityupdatepoliciesItemAssignRequestBuilder) {
    return NewWindowsqualityupdatepoliciesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsQualityUpdatePolicy entity.
// returns a *WindowsqualityupdatepoliciesItemAssignmentsRequestBuilder when successful
func (m *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder) Assignments()(*WindowsqualityupdatepoliciesItemAssignmentsRequestBuilder) {
    return NewWindowsqualityupdatepoliciesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderInternal instantiates a new WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder and sets the default values.
func NewWindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder) {
    m := &WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsQualityUpdatePolicies/{windowsQualityUpdatePolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder instantiates a new WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder and sets the default values.
func NewWindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsQualityUpdatePolicies for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyable, error) {
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
func (m *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyable, requestConfiguration *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyable, error) {
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
func (m *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyable, requestConfiguration *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder when successful
func (m *WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder) WithUrl(rawUrl string)(*WindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder) {
    return NewWindowsqualityupdatepoliciesWindowsQualityUpdatePolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
