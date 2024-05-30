package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder provides operations to count the resources in the collection.
type B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilderGetQueryParameters get the number of the resource
type B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilderGetQueryParameters
}
// NewB2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilderInternal instantiates a new B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder and sets the default values.
func NewB2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder) {
    m := &B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/languages/{userFlowLanguageConfiguration%2Did}/defaultPages/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewB2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder instantiates a new B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder and sets the default values.
func NewB2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
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
func (m *B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder when successful
func (m *B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder) WithUrl(rawUrl string)(*B2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder) {
    return NewB2cuserflowsItemLanguagesItemDefaultpagesCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
