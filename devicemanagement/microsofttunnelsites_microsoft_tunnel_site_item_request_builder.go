package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder provides operations to manage the microsoftTunnelSites property of the microsoft.graph.deviceManagement entity.
type MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderGetQueryParameters collection of MicrosoftTunnelSite settings associated with account.
type MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderGetQueryParameters
}
// MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderInternal instantiates a new MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder and sets the default values.
func NewMicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) {
    m := &MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder instantiates a new MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder and sets the default values.
func NewMicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property microsoftTunnelSites for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of MicrosoftTunnelSite settings associated with account.
// returns a MicrosoftTunnelSiteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable), nil
}
// MicrosoftTunnelConfiguration provides operations to manage the microsoftTunnelConfiguration property of the microsoft.graph.microsoftTunnelSite entity.
// returns a *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder when successful
func (m *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) MicrosoftTunnelConfiguration()(*MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder) {
    return NewMicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftTunnelServers provides operations to manage the microsoftTunnelServers property of the microsoft.graph.microsoftTunnelSite entity.
// returns a *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServersRequestBuilder when successful
func (m *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) MicrosoftTunnelServers()(*MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServersRequestBuilder) {
    return NewMicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property microsoftTunnelSites in deviceManagement
// returns a MicrosoftTunnelSiteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, requestConfiguration *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable), nil
}
// RequestUpgrade provides operations to call the requestUpgrade method.
// returns a *MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder when successful
func (m *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) RequestUpgrade()(*MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder) {
    return NewMicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property microsoftTunnelSites for deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of MicrosoftTunnelSite settings associated with account.
// returns a *RequestInformation when successful
func (m *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property microsoftTunnelSites in deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, requestConfiguration *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder when successful
func (m *MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) WithUrl(rawUrl string)(*MicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder) {
    return NewMicrosofttunnelsitesMicrosoftTunnelSiteItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
