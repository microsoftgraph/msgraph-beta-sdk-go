package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder provides operations to manage the noncustodialSources property of the microsoft.graph.security.ediscoverySearch entity.
type CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetQueryParameters noncustodialDataSource sources that are included in the eDiscovery search
type CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetQueryParameters
}
// NewCasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal instantiates a new EdiscoveryNoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    m := &CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}/noncustodialSources/{ediscoveryNoncustodialDataSource%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder instantiates a new EdiscoveryNoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove an ediscoveryNoncustodialDataSource object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverysearch-delete-noncustodialsources?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get noncustodialDataSource sources that are included in the eDiscovery search
func (m *CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceable), nil
}
// ToDeleteRequestInformation remove an ediscoveryNoncustodialDataSource object.
func (m *CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation noncustodialDataSource sources that are included in the eDiscovery search
func (m *CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    return NewCasesEdiscoveryCasesItemSearchesItemNoncustodialSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
