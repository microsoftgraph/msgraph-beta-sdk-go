package termstore

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SetsItemParentGroupSetsItemTermsCountRequestBuilder provides operations to count the resources in the collection.
type SetsItemParentGroupSetsItemTermsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SetsItemParentGroupSetsItemTermsCountRequestBuilderGetQueryParameters get the number of the resource
type SetsItemParentGroupSetsItemTermsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// SetsItemParentGroupSetsItemTermsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetsItemParentGroupSetsItemTermsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SetsItemParentGroupSetsItemTermsCountRequestBuilderGetQueryParameters
}
// NewSetsItemParentGroupSetsItemTermsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewSetsItemParentGroupSetsItemTermsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetsItemParentGroupSetsItemTermsCountRequestBuilder) {
    m := &SetsItemParentGroupSetsItemTermsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/terms/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewSetsItemParentGroupSetsItemTermsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewSetsItemParentGroupSetsItemTermsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetsItemParentGroupSetsItemTermsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetsItemParentGroupSetsItemTermsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *SetsItemParentGroupSetsItemTermsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *SetsItemParentGroupSetsItemTermsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *SetsItemParentGroupSetsItemTermsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SetsItemParentGroupSetsItemTermsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *SetsItemParentGroupSetsItemTermsCountRequestBuilder) WithUrl(rawUrl string)(*SetsItemParentGroupSetsItemTermsCountRequestBuilder) {
    return NewSetsItemParentGroupSetsItemTermsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
