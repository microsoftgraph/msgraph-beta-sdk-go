package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder provides operations to manage the domainJoinConnectors property of the microsoft.graph.deviceManagement entity.
type DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderGetQueryParameters a list of connector objects.
type DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderGetQueryParameters
}
// DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderInternal instantiates a new DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder and sets the default values.
func NewDomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) {
    m := &DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/domainJoinConnectors/{deviceManagementDomainJoinConnector%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder instantiates a new DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder and sets the default values.
func NewDomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property domainJoinConnectors for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a list of connector objects.
// returns a DeviceManagementDomainJoinConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDomainJoinConnectorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementDomainJoinConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDomainJoinConnectorable), nil
}
// Patch update the navigation property domainJoinConnectors in deviceManagement
// returns a DeviceManagementDomainJoinConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDomainJoinConnectorable, requestConfiguration *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDomainJoinConnectorable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementDomainJoinConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDomainJoinConnectorable), nil
}
// ToDeleteRequestInformation delete navigation property domainJoinConnectors for deviceManagement
// returns a *RequestInformation when successful
func (m *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a list of connector objects.
// returns a *RequestInformation when successful
func (m *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property domainJoinConnectors in deviceManagement
// returns a *RequestInformation when successful
func (m *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDomainJoinConnectorable, requestConfiguration *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder when successful
func (m *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) WithUrl(rawUrl string)(*DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) {
    return NewDomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
