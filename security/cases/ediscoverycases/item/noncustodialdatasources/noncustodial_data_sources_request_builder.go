package noncustodialdatasources

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic47dce610e2866c8627520246bfe92f2621855e95a0d49a4725a2cfdd8b31696 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/noncustodialdatasources/removehold"
    ied9816c514c49600d9697fcda50b155b61cff4e841fad74a7c5f00a9877dd712 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/noncustodialdatasources/applyhold"
    iee8c9df871e99835b8cb942549b2048d831445eea844e4d9e151ee8e7131f7a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/noncustodialdatasources/count"
)

// NoncustodialDataSourcesRequestBuilder provides operations to manage the noncustodialDataSources property of the microsoft.graph.security.ediscoveryCase entity.
type NoncustodialDataSourcesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NoncustodialDataSourcesRequestBuilderGetQueryParameters returns a list of case ediscoveryNoncustodialDataSource objects for this case.
type NoncustodialDataSourcesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// NoncustodialDataSourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NoncustodialDataSourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *NoncustodialDataSourcesRequestBuilderGetQueryParameters
}
// NoncustodialDataSourcesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NoncustodialDataSourcesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplyHold provides operations to call the applyHold method.
func (m *NoncustodialDataSourcesRequestBuilder) ApplyHold()(*ied9816c514c49600d9697fcda50b155b61cff4e841fad74a7c5f00a9877dd712.ApplyHoldRequestBuilder) {
    return ied9816c514c49600d9697fcda50b155b61cff4e841fad74a7c5f00a9877dd712.NewApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewNoncustodialDataSourcesRequestBuilderInternal instantiates a new NoncustodialDataSourcesRequestBuilder and sets the default values.
func NewNoncustodialDataSourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NoncustodialDataSourcesRequestBuilder) {
    m := &NoncustodialDataSourcesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNoncustodialDataSourcesRequestBuilder instantiates a new NoncustodialDataSourcesRequestBuilder and sets the default values.
func NewNoncustodialDataSourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NoncustodialDataSourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNoncustodialDataSourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *NoncustodialDataSourcesRequestBuilder) Count()(*iee8c9df871e99835b8cb942549b2048d831445eea844e4d9e151ee8e7131f7a9.CountRequestBuilder) {
    return iee8c9df871e99835b8cb942549b2048d831445eea844e4d9e151ee8e7131f7a9.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation returns a list of case ediscoveryNoncustodialDataSource objects for this case.
func (m *NoncustodialDataSourcesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *NoncustodialDataSourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create a new ediscoveryNoncustodialDataSource object.
func (m *NoncustodialDataSourcesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceable, requestConfiguration *NoncustodialDataSourcesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get returns a list of case ediscoveryNoncustodialDataSource objects for this case.
func (m *NoncustodialDataSourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *NoncustodialDataSourcesRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryNoncustodialDataSourceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceCollectionResponseable), nil
}
// Post create a new ediscoveryNoncustodialDataSource object.
func (m *NoncustodialDataSourcesRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceable, requestConfiguration *NoncustodialDataSourcesRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryNoncustodialDataSourceable), nil
}
// RemoveHold provides operations to call the removeHold method.
func (m *NoncustodialDataSourcesRequestBuilder) RemoveHold()(*ic47dce610e2866c8627520246bfe92f2621855e95a0d49a4725a2cfdd8b31696.RemoveHoldRequestBuilder) {
    return ic47dce610e2866c8627520246bfe92f2621855e95a0d49a4725a2cfdd8b31696.NewRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
