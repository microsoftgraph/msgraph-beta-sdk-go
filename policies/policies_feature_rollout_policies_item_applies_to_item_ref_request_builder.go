package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder provides operations to manage the collection of policyRoot entities.
type PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteQueryParameters delete ref of navigation property appliesTo for policies
type PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteQueryParameters
}
// NewPoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewPoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) {
    m := &PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo/{directoryObject%2Did}/$ref{?%40id*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewPoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete ref of navigation property appliesTo for policies
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete ref of navigation property appliesTo for policies
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *PoliciesFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
