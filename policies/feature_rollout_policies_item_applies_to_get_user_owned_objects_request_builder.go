package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder provides operations to call the getUserOwnedObjects method.
type FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilderInternal instantiates a new GetUserOwnedObjectsRequestBuilder and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder) {
    m := &FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo/getUserOwnedObjects", pathParameters),
    }
    return m
}
// NewFeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder instantiates a new GetUserOwnedObjectsRequestBuilder and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post retrieve a list of recently deleted application and group objects owned by the specified user. This API returns up to 1,000 deleted objects owned by the user, sorted by ID, and doesn't support pagination.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directory-deleteditems-getuserownedobjects?view=graph-rest-1.0
func (m *FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder) Post(ctx context.Context, body FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsPostRequestBodyable, requestConfiguration *FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// ToPostRequestInformation retrieve a list of recently deleted application and group objects owned by the specified user. This API returns up to 1,000 deleted objects owned by the user, sorted by ID, and doesn't support pagination.
func (m *FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, body FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsPostRequestBodyable, requestConfiguration *FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
func (m *FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder) WithUrl(rawUrl string)(*FeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder) {
    return NewFeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
