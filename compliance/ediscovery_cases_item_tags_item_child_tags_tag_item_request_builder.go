package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder provides operations to manage the childTags property of the microsoft.graph.ediscovery.tag entity.
type EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilderGetQueryParameters returns the tags that are a child of a tag.
type EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilderGetQueryParameters
}
// NewEdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilderInternal instantiates a new EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder) {
    m := &EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/tags/{tag%2Did}/childTags/{tag%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder instantiates a new EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the tags that are a child of a tag.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a Tagable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateTagFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Tagable), nil
}
// ToGetRequestInformation returns the tags that are a child of a tag.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder when successful
func (m *EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder) {
    return NewEdiscoveryCasesItemTagsItemChildTagsTagItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
