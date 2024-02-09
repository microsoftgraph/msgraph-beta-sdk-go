package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder provides operations to count the resources in the collection.
type CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilderGetQueryParameters get the number of the resource
type CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilderGetQueryParameters
}
// NewCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilderInternal instantiates a new CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder) {
    m := &CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/legalHolds/{ediscoveryHoldPolicy%2Did}/userSources/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder instantiates a new CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder when successful
func (m *CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder) {
    return NewCasesEdiscoveryCasesItemLegalHoldsItemUserSourcesCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
