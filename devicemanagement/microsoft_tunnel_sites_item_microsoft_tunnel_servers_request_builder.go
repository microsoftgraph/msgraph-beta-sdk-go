package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder provides operations to manage the microsoftTunnelServers property of the microsoft.graph.microsoftTunnelSite entity.
type MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderGetQueryParameters a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
type MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderGetQueryParameters struct {
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
// MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderGetQueryParameters
}
// MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMicrosoftTunnelServerId provides operations to manage the microsoftTunnelServers property of the microsoft.graph.microsoftTunnelSite entity.
// returns a *MicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder when successful
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder) ByMicrosoftTunnelServerId(microsoftTunnelServerId string)(*MicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if microsoftTunnelServerId != "" {
        urlTplParams["microsoftTunnelServer%2Did"] = microsoftTunnelServerId
    }
    return NewMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderInternal instantiates a new MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder and sets the default values.
func NewMicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder) {
    m := &MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}/microsoftTunnelServers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder instantiates a new MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder and sets the default values.
func NewMicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MicrosoftTunnelSitesItemMicrosoftTunnelServersCountRequestBuilder when successful
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder) Count()(*MicrosoftTunnelSitesItemMicrosoftTunnelServersCountRequestBuilder) {
    return NewMicrosoftTunnelSitesItemMicrosoftTunnelServersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
// returns a MicrosoftTunnelServerCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelServerCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerCollectionResponseable), nil
}
// Post create new navigation property to microsoftTunnelServers for deviceManagement
// returns a MicrosoftTunnelServerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, requestConfiguration *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelServerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable), nil
}
// ToGetRequestInformation a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
// returns a *RequestInformation when successful
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to microsoftTunnelServers for deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, requestConfiguration *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}/microsoftTunnelServers", m.BaseRequestBuilder.PathParameters)
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
// returns a *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder when successful
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder) WithUrl(rawUrl string)(*MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder) {
    return NewMicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
