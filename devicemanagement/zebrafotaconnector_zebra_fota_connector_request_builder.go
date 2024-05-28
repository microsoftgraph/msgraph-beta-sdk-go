package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebrafotaconnectorZebraFotaConnectorRequestBuilder provides operations to manage the zebraFotaConnector property of the microsoft.graph.deviceManagement entity.
type ZebrafotaconnectorZebraFotaConnectorRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZebrafotaconnectorZebraFotaConnectorRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebrafotaconnectorZebraFotaConnectorRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ZebrafotaconnectorZebraFotaConnectorRequestBuilderGetQueryParameters the singleton ZebraFotaConnector associated with account.
type ZebrafotaconnectorZebraFotaConnectorRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ZebrafotaconnectorZebraFotaConnectorRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebrafotaconnectorZebraFotaConnectorRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ZebrafotaconnectorZebraFotaConnectorRequestBuilderGetQueryParameters
}
// ZebrafotaconnectorZebraFotaConnectorRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebrafotaconnectorZebraFotaConnectorRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApproveFotaApps provides operations to call the approveFotaApps method.
// returns a *ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder when successful
func (m *ZebrafotaconnectorZebraFotaConnectorRequestBuilder) ApproveFotaApps()(*ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder) {
    return NewZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Connect provides operations to call the connect method.
// returns a *ZebrafotaconnectorConnectRequestBuilder when successful
func (m *ZebrafotaconnectorZebraFotaConnectorRequestBuilder) Connect()(*ZebrafotaconnectorConnectRequestBuilder) {
    return NewZebrafotaconnectorConnectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewZebrafotaconnectorZebraFotaConnectorRequestBuilderInternal instantiates a new ZebrafotaconnectorZebraFotaConnectorRequestBuilder and sets the default values.
func NewZebrafotaconnectorZebraFotaConnectorRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebrafotaconnectorZebraFotaConnectorRequestBuilder) {
    m := &ZebrafotaconnectorZebraFotaConnectorRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/zebraFotaConnector{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewZebrafotaconnectorZebraFotaConnectorRequestBuilder instantiates a new ZebrafotaconnectorZebraFotaConnectorRequestBuilder and sets the default values.
func NewZebrafotaconnectorZebraFotaConnectorRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebrafotaconnectorZebraFotaConnectorRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebrafotaconnectorZebraFotaConnectorRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property zebraFotaConnector for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotaconnectorZebraFotaConnectorRequestBuilder) Delete(ctx context.Context, requestConfiguration *ZebrafotaconnectorZebraFotaConnectorRequestBuilderDeleteRequestConfiguration)(error) {
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
// Disconnect provides operations to call the disconnect method.
// returns a *ZebrafotaconnectorDisconnectRequestBuilder when successful
func (m *ZebrafotaconnectorZebraFotaConnectorRequestBuilder) Disconnect()(*ZebrafotaconnectorDisconnectRequestBuilder) {
    return NewZebrafotaconnectorDisconnectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the singleton ZebraFotaConnector associated with account.
// returns a ZebraFotaConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotaconnectorZebraFotaConnectorRequestBuilder) Get(ctx context.Context, requestConfiguration *ZebrafotaconnectorZebraFotaConnectorRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateZebraFotaConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable), nil
}
// HasActiveDeployments provides operations to call the hasActiveDeployments method.
// returns a *ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder when successful
func (m *ZebrafotaconnectorZebraFotaConnectorRequestBuilder) HasActiveDeployments()(*ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder) {
    return NewZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property zebraFotaConnector in deviceManagement
// returns a ZebraFotaConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotaconnectorZebraFotaConnectorRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable, requestConfiguration *ZebrafotaconnectorZebraFotaConnectorRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateZebraFotaConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable), nil
}
// ToDeleteRequestInformation delete navigation property zebraFotaConnector for deviceManagement
// returns a *RequestInformation when successful
func (m *ZebrafotaconnectorZebraFotaConnectorRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ZebrafotaconnectorZebraFotaConnectorRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the singleton ZebraFotaConnector associated with account.
// returns a *RequestInformation when successful
func (m *ZebrafotaconnectorZebraFotaConnectorRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ZebrafotaconnectorZebraFotaConnectorRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property zebraFotaConnector in deviceManagement
// returns a *RequestInformation when successful
func (m *ZebrafotaconnectorZebraFotaConnectorRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaConnectorable, requestConfiguration *ZebrafotaconnectorZebraFotaConnectorRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ZebrafotaconnectorZebraFotaConnectorRequestBuilder when successful
func (m *ZebrafotaconnectorZebraFotaConnectorRequestBuilder) WithUrl(rawUrl string)(*ZebrafotaconnectorZebraFotaConnectorRequestBuilder) {
    return NewZebrafotaconnectorZebraFotaConnectorRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
