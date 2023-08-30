package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder provides operations to manage the activityBasedTimeoutPolicies property of the microsoft.graph.policyRoot entity.
type ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetQueryParameters get the properties of an activityBasedTimeoutPolicy object.
type ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetQueryParameters
}
// ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppliesTo provides operations to manage the appliesTo property of the microsoft.graph.stsPolicy entity.
func (m *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) AppliesTo()(*ActivityBasedTimeoutPoliciesItemAppliesToRequestBuilder) {
    return NewActivityBasedTimeoutPoliciesItemAppliesToRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderInternal instantiates a new ActivityBasedTimeoutPolicyItemRequestBuilder and sets the default values.
func NewActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) {
    m := &ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/activityBasedTimeoutPolicies/{activityBasedTimeoutPolicy%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder instantiates a new ActivityBasedTimeoutPolicyItemRequestBuilder and sets the default values.
func NewActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an activityBasedTimeoutPolicy object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/activitybasedtimeoutpolicy-delete?view=graph-rest-1.0
func (m *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the properties of an activityBasedTimeoutPolicy object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/activitybasedtimeoutpolicy-get?view=graph-rest-1.0
func (m *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityBasedTimeoutPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateActivityBasedTimeoutPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityBasedTimeoutPolicyable), nil
}
// Patch update the properties of an activityBasedTimeoutPolicy object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/activitybasedtimeoutpolicy-update?view=graph-rest-1.0
func (m *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityBasedTimeoutPolicyable, requestConfiguration *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityBasedTimeoutPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateActivityBasedTimeoutPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityBasedTimeoutPolicyable), nil
}
// ToDeleteRequestInformation delete an activityBasedTimeoutPolicy object.
func (m *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get the properties of an activityBasedTimeoutPolicy object.
func (m *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an activityBasedTimeoutPolicy object.
func (m *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityBasedTimeoutPolicyable, requestConfiguration *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) {
    return NewActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
