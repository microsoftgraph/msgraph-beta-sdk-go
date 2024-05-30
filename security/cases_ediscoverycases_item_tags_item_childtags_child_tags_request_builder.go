package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder provides operations to manage the childTags property of the microsoft.graph.security.ediscoveryReviewTag entity.
type CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilderGetQueryParameters returns the tags that are a child of a tag.
type CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilderGetQueryParameters struct {
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
// CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilderGetQueryParameters
}
// ByEdiscoveryReviewTagId1 provides operations to manage the childTags property of the microsoft.graph.security.ediscoveryReviewTag entity.
// returns a *CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder) ByEdiscoveryReviewTagId1(ediscoveryReviewTagId1 string)(*CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ediscoveryReviewTagId1 != "" {
        urlTplParams["ediscoveryReviewTag%2Did1"] = ediscoveryReviewTagId1
    }
    return NewCasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder) {
    m := &CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/tags/{ediscoveryReviewTag%2Did}/childTags{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder instantiates a new CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CasesEdiscoverycasesItemTagsItemChildtagsCountRequestBuilder when successful
func (m *CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder) Count()(*CasesEdiscoverycasesItemTagsItemChildtagsCountRequestBuilder) {
    return NewCasesEdiscoverycasesItemTagsItemChildtagsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get returns the tags that are a child of a tag.
// returns a EdiscoveryReviewTagCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewTagCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryReviewTagCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewTagCollectionResponseable), nil
}
// ToGetRequestInformation returns the tags that are a child of a tag.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder when successful
func (m *CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder) {
    return NewCasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
