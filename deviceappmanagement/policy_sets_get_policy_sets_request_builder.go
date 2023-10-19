package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PolicySetsGetPolicySetsRequestBuilder provides operations to call the getPolicySets method.
type PolicySetsGetPolicySetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PolicySetsGetPolicySetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PolicySetsGetPolicySetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPolicySetsGetPolicySetsRequestBuilderInternal instantiates a new GetPolicySetsRequestBuilder and sets the default values.
func NewPolicySetsGetPolicySetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicySetsGetPolicySetsRequestBuilder) {
    m := &PolicySetsGetPolicySetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/policySets/getPolicySets", pathParameters),
    }
    return m
}
// NewPolicySetsGetPolicySetsRequestBuilder instantiates a new GetPolicySetsRequestBuilder and sets the default values.
func NewPolicySetsGetPolicySetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicySetsGetPolicySetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPolicySetsGetPolicySetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getPolicySets
// Deprecated: This method is obsolete. Use PostAsGetPolicySetsPostResponse instead.
func (m *PolicySetsGetPolicySetsRequestBuilder) Post(ctx context.Context, body PolicySetsGetPolicySetsPostRequestBodyable, requestConfiguration *PolicySetsGetPolicySetsRequestBuilderPostRequestConfiguration)(PolicySetsGetPolicySetsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePolicySetsGetPolicySetsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PolicySetsGetPolicySetsResponseable), nil
}
// PostAsGetPolicySetsPostResponse invoke action getPolicySets
func (m *PolicySetsGetPolicySetsRequestBuilder) PostAsGetPolicySetsPostResponse(ctx context.Context, body PolicySetsGetPolicySetsPostRequestBodyable, requestConfiguration *PolicySetsGetPolicySetsRequestBuilderPostRequestConfiguration)(PolicySetsGetPolicySetsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePolicySetsGetPolicySetsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PolicySetsGetPolicySetsPostResponseable), nil
}
// ToPostRequestInformation invoke action getPolicySets
func (m *PolicySetsGetPolicySetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body PolicySetsGetPolicySetsPostRequestBodyable, requestConfiguration *PolicySetsGetPolicySetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *PolicySetsGetPolicySetsRequestBuilder) WithUrl(rawUrl string)(*PolicySetsGetPolicySetsRequestBuilder) {
    return NewPolicySetsGetPolicySetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
