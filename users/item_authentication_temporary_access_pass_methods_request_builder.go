package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
type ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationTemporaryAccessPassMethodsRequestBuilderGetQueryParameters represents a Temporary Access Pass registered to a user for authentication through time-limited passcodes.
type ItemAuthenticationTemporaryAccessPassMethodsRequestBuilderGetQueryParameters struct {
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
// ItemAuthenticationTemporaryAccessPassMethodsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationTemporaryAccessPassMethodsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilderGetQueryParameters
}
// ItemAuthenticationTemporaryAccessPassMethodsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationTemporaryAccessPassMethodsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTemporaryAccessPassAuthenticationMethodId provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder when successful
func (m *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder) ByTemporaryAccessPassAuthenticationMethodId(temporaryAccessPassAuthenticationMethodId string)(*ItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if temporaryAccessPassAuthenticationMethodId != "" {
        urlTplParams["temporaryAccessPassAuthenticationMethod%2Did"] = temporaryAccessPassAuthenticationMethodId
    }
    return NewItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAuthenticationTemporaryAccessPassMethodsRequestBuilderInternal instantiates a new ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder and sets the default values.
func NewItemAuthenticationTemporaryAccessPassMethodsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder) {
    m := &ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/temporaryAccessPassMethods{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAuthenticationTemporaryAccessPassMethodsRequestBuilder instantiates a new ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder and sets the default values.
func NewItemAuthenticationTemporaryAccessPassMethodsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationTemporaryAccessPassMethodsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAuthenticationTemporaryAccessPassMethodsCountRequestBuilder when successful
func (m *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder) Count()(*ItemAuthenticationTemporaryAccessPassMethodsCountRequestBuilder) {
    return NewItemAuthenticationTemporaryAccessPassMethodsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents a Temporary Access Pass registered to a user for authentication through time-limited passcodes.
// returns a TemporaryAccessPassAuthenticationMethodCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TemporaryAccessPassAuthenticationMethodCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTemporaryAccessPassAuthenticationMethodCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TemporaryAccessPassAuthenticationMethodCollectionResponseable), nil
}
// Post create new navigation property to temporaryAccessPassMethods for users
// returns a TemporaryAccessPassAuthenticationMethodable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TemporaryAccessPassAuthenticationMethodable, requestConfiguration *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TemporaryAccessPassAuthenticationMethodable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTemporaryAccessPassAuthenticationMethodFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TemporaryAccessPassAuthenticationMethodable), nil
}
// ToGetRequestInformation represents a Temporary Access Pass registered to a user for authentication through time-limited passcodes.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to temporaryAccessPassMethods for users
// returns a *RequestInformation when successful
func (m *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TemporaryAccessPassAuthenticationMethodable, requestConfiguration *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder when successful
func (m *ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationTemporaryAccessPassMethodsRequestBuilder) {
    return NewItemAuthenticationTemporaryAccessPassMethodsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
