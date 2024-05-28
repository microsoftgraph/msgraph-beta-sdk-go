package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder provides operations to manage the microsoftTunnelConfiguration property of the microsoft.graph.microsoftTunnelSite entity.
type MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderGetQueryParameters the MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
type MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderGetQueryParameters
}
// MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderInternal instantiates a new MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder and sets the default values.
func NewMicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder) {
    m := &MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}/microsoftTunnelConfiguration{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder instantiates a new MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder and sets the default values.
func NewMicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property microsoftTunnelConfiguration for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder) Delete(ctx context.Context, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
// returns a MicrosoftTunnelConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property microsoftTunnelConfiguration in deviceManagement
// returns a MicrosoftTunnelConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelConfigurationable, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property microsoftTunnelConfiguration for deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
// returns a *RequestInformation when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property microsoftTunnelConfiguration in deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelConfigurationable, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder) WithUrl(rawUrl string)(*MicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder) {
    return NewMicrosofttunnelsitesItemMicrosofttunnelconfigurationMicrosoftTunnelConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
