package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder provides operations to manage the targetDeviceGroups property of the microsoft.graph.remoteDesktopSecurityConfiguration entity.
type ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderGetQueryParameters the collection of target device groups that are associated with the RDS security configuration that will be enabled for SSO when a client connects to the target device over RDP using the new Microsoft Entra ID RDS authentication protocol.
type ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderGetQueryParameters struct {
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
// ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderGetQueryParameters
}
// ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTargetDeviceGroupId provides operations to manage the targetDeviceGroups property of the microsoft.graph.remoteDesktopSecurityConfiguration entity.
// returns a *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsTargetDeviceGroupItemRequestBuilder when successful
func (m *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder) ByTargetDeviceGroupId(targetDeviceGroupId string)(*ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsTargetDeviceGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if targetDeviceGroupId != "" {
        urlTplParams["targetDeviceGroup%2Did"] = targetDeviceGroupId
    }
    return NewItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsTargetDeviceGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderInternal instantiates a new ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder and sets the default values.
func NewItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder) {
    m := &ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/remoteDesktopSecurityConfiguration/targetDeviceGroups{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder instantiates a new ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder and sets the default values.
func NewItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsCountRequestBuilder when successful
func (m *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder) Count()(*ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsCountRequestBuilder) {
    return NewItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of target device groups that are associated with the RDS security configuration that will be enabled for SSO when a client connects to the target device over RDP using the new Microsoft Entra ID RDS authentication protocol.
// returns a TargetDeviceGroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetDeviceGroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTargetDeviceGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetDeviceGroupCollectionResponseable), nil
}
// Post create new navigation property to targetDeviceGroups for servicePrincipals
// returns a TargetDeviceGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetDeviceGroupable, requestConfiguration *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetDeviceGroupable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTargetDeviceGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetDeviceGroupable), nil
}
// ToGetRequestInformation the collection of target device groups that are associated with the RDS security configuration that will be enabled for SSO when a client connects to the target device over RDP using the new Microsoft Entra ID RDS authentication protocol.
// returns a *RequestInformation when successful
func (m *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to targetDeviceGroups for servicePrincipals
// returns a *RequestInformation when successful
func (m *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TargetDeviceGroupable, requestConfiguration *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder when successful
func (m *ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder) {
    return NewItemRemoteDesktopSecurityConfigurationTargetDeviceGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
