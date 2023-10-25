package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder provides operations to manage the microsoftTunnelSites property of the microsoft.graph.deviceManagement entity.
type MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetQueryParameters collection of MicrosoftTunnelSite settings associated with account.
type MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetQueryParameters
}
// MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderInternal instantiates a new MicrosoftTunnelSiteItemRequestBuilder and sets the default values.
func NewMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) {
    m := &MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder instantiates a new MicrosoftTunnelSiteItemRequestBuilder and sets the default values.
func NewMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property microsoftTunnelSites for deviceManagement
func (m *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of MicrosoftTunnelSite settings associated with account.
func (m *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) MicrosoftTunnelConfiguration()(*MicrosoftTunnelSitesItemMicrosoftTunnelConfigurationRequestBuilder) {
    return NewMicrosoftTunnelSitesItemMicrosoftTunnelConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftTunnelServers provides operations to manage the microsoftTunnelServers property of the microsoft.graph.microsoftTunnelSite entity.
func (m *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) MicrosoftTunnelServers()(*MicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder) {
    return NewMicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property microsoftTunnelSites in deviceManagement
func (m *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, requestConfiguration *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) RequestUpgrade()(*MicrosoftTunnelSitesItemRequestUpgradeRequestBuilder) {
    return NewMicrosoftTunnelSitesItemRequestUpgradeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property microsoftTunnelSites for deviceManagement
func (m *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of MicrosoftTunnelSite settings associated with account.
func (m *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property microsoftTunnelSites in deviceManagement
func (m *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, requestConfiguration *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) WithUrl(rawUrl string)(*MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) {
    return NewMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
