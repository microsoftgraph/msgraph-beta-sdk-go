package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder provides operations to manage the deviceLocalCredentials property of the microsoft.graph.directory entity.
type DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicelocalcredentialsDeviceLocalCredentialsRequestBuilderGetQueryParameters get a list of the deviceLocalCredentialInfo objects and their properties excluding the credentials. 
type DevicelocalcredentialsDeviceLocalCredentialsRequestBuilderGetQueryParameters struct {
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
// DevicelocalcredentialsDeviceLocalCredentialsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicelocalcredentialsDeviceLocalCredentialsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilderGetQueryParameters
}
// DevicelocalcredentialsDeviceLocalCredentialsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicelocalcredentialsDeviceLocalCredentialsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceLocalCredentialInfoId provides operations to manage the deviceLocalCredentials property of the microsoft.graph.directory entity.
// returns a *DevicelocalcredentialsDeviceLocalCredentialInfoItemRequestBuilder when successful
func (m *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder) ByDeviceLocalCredentialInfoId(deviceLocalCredentialInfoId string)(*DevicelocalcredentialsDeviceLocalCredentialInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceLocalCredentialInfoId != "" {
        urlTplParams["deviceLocalCredentialInfo%2Did"] = deviceLocalCredentialInfoId
    }
    return NewDevicelocalcredentialsDeviceLocalCredentialInfoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicelocalcredentialsDeviceLocalCredentialsRequestBuilderInternal instantiates a new DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder and sets the default values.
func NewDevicelocalcredentialsDeviceLocalCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder) {
    m := &DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/deviceLocalCredentials{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDevicelocalcredentialsDeviceLocalCredentialsRequestBuilder instantiates a new DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder and sets the default values.
func NewDevicelocalcredentialsDeviceLocalCredentialsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicelocalcredentialsDeviceLocalCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DevicelocalcredentialsCountRequestBuilder when successful
func (m *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder) Count()(*DevicelocalcredentialsCountRequestBuilder) {
    return NewDevicelocalcredentialsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the deviceLocalCredentialInfo objects and their properties excluding the credentials. 
// returns a DeviceLocalCredentialInfoCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directory-list-devicelocalcredentials?view=graph-rest-beta
func (m *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceLocalCredentialInfoCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceLocalCredentialInfoCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceLocalCredentialInfoCollectionResponseable), nil
}
// Post create new navigation property to deviceLocalCredentials for directory
// returns a DeviceLocalCredentialInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceLocalCredentialInfoable, requestConfiguration *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceLocalCredentialInfoable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceLocalCredentialInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceLocalCredentialInfoable), nil
}
// ToGetRequestInformation get a list of the deviceLocalCredentialInfo objects and their properties excluding the credentials. 
// returns a *RequestInformation when successful
func (m *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceLocalCredentials for directory
// returns a *RequestInformation when successful
func (m *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceLocalCredentialInfoable, requestConfiguration *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder when successful
func (m *DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder) WithUrl(rawUrl string)(*DevicelocalcredentialsDeviceLocalCredentialsRequestBuilder) {
    return NewDevicelocalcredentialsDeviceLocalCredentialsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
