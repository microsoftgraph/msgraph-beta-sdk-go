package termstore

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilder provides operations to count the resources in the collection.
type SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilderGetQueryParameters get the number of the resource
type SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilderGetQueryParameters
}
// NewSetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewSetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilder) {
    m := &SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/terms/{term%2Did}/relations/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewSetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewSetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
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
func (m *SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilder) WithUrl(rawUrl string)(*SetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilder) {
    return NewSetsItemParentGroupSetsItemTermsItemRelationsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
