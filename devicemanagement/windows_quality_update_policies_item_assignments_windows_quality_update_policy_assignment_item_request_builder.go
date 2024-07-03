package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.windowsQualityUpdatePolicy entity.
type WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderGetQueryParameters list of the groups this profile is assgined to.
type WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderGetQueryParameters
}
// WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderInternal instantiates a new WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder and sets the default values.
func NewWindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder) {
    m := &WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsQualityUpdatePolicies/{windowsQualityUpdatePolicy%2Did}/assignments/{windowsQualityUpdatePolicyAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder instantiates a new WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder and sets the default values.
func NewWindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of the groups this profile is assgined to.
// returns a WindowsQualityUpdatePolicyAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsQualityUpdatePolicyAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyAssignmentable), nil
}
// Patch update the navigation property assignments in deviceManagement
// returns a WindowsQualityUpdatePolicyAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyAssignmentable, requestConfiguration *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsQualityUpdatePolicyAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyAssignmentable), nil
}
// ToDeleteRequestInformation delete navigation property assignments for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of the groups this profile is assgined to.
// returns a *RequestInformation when successful
func (m *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsQualityUpdatePolicyAssignmentable, requestConfiguration *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder when successful
func (m *WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*WindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder) {
    return NewWindowsQualityUpdatePoliciesItemAssignmentsWindowsQualityUpdatePolicyAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
