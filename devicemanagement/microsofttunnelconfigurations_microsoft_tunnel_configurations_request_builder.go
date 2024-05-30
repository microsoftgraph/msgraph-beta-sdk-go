package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder provides operations to manage the microsoftTunnelConfigurations property of the microsoft.graph.deviceManagement entity.
type MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderGetQueryParameters collection of MicrosoftTunnelConfiguration settings associated with account.
type MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderGetQueryParameters struct {
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
// MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderGetQueryParameters
}
// MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMicrosoftTunnelConfigurationId provides operations to manage the microsoftTunnelConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationItemRequestBuilder when successful
func (m *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder) ByMicrosoftTunnelConfigurationId(microsoftTunnelConfigurationId string)(*MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if microsoftTunnelConfigurationId != "" {
        urlTplParams["microsoftTunnelConfiguration%2Did"] = microsoftTunnelConfigurationId
    }
    return NewMicrosofttunnelconfigurationsMicrosoftTunnelConfigurationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderInternal instantiates a new MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder and sets the default values.
func NewMicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder) {
    m := &MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder instantiates a new MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder and sets the default values.
func NewMicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MicrosofttunnelconfigurationsCountRequestBuilder when successful
func (m *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder) Count()(*MicrosofttunnelconfigurationsCountRequestBuilder) {
    return NewMicrosofttunnelconfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of MicrosoftTunnelConfiguration settings associated with account.
// returns a MicrosoftTunnelConfigurationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelConfigurationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelConfigurationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelConfigurationCollectionResponseable), nil
}
// Post create new navigation property to microsoftTunnelConfigurations for deviceManagement
// returns a MicrosoftTunnelConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelConfigurationable, requestConfiguration *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelConfigurationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelConfigurationable), nil
}
// ToGetRequestInformation collection of MicrosoftTunnelConfiguration settings associated with account.
// returns a *RequestInformation when successful
func (m *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to microsoftTunnelConfigurations for deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelConfigurationable, requestConfiguration *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder when successful
func (m *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder) WithUrl(rawUrl string)(*MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder) {
    return NewMicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
