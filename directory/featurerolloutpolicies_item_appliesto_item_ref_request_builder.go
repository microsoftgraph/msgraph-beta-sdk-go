package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder provides operations to manage the collection of directory entities.
type FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFeaturerolloutpoliciesItemAppliestoItemRefRequestBuilderInternal instantiates a new FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder and sets the default values.
func NewFeaturerolloutpoliciesItemAppliestoItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder) {
    m := &FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo/{directoryObject%2Did}/$ref", pathParameters),
    }
    return m
}
// NewFeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder instantiates a new FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder and sets the default values.
func NewFeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFeaturerolloutpoliciesItemAppliestoItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete ref of navigation property appliesTo for directory
// Deprecated: Feature Rollout Policies have been grouped with other policies under /policies. The existing /directory/featureRolloutPolicies is deprecated and will stop returning data on 06/30/2021. Please use /policies/featureRolloutPolicies. as of 2021-01/DirectoryFeatureRolloutPolicies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation delete ref of navigation property appliesTo for directory
// Deprecated: Feature Rollout Policies have been grouped with other policies under /policies. The existing /directory/featureRolloutPolicies is deprecated and will stop returning data on 06/30/2021. Please use /policies/featureRolloutPolicies. as of 2021-01/DirectoryFeatureRolloutPolicies
// returns a *RequestInformation when successful
func (m *FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: Feature Rollout Policies have been grouped with other policies under /policies. The existing /directory/featureRolloutPolicies is deprecated and will stop returning data on 06/30/2021. Please use /policies/featureRolloutPolicies. as of 2021-01/DirectoryFeatureRolloutPolicies
// returns a *FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder when successful
func (m *FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder) WithUrl(rawUrl string)(*FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder) {
    return NewFeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
