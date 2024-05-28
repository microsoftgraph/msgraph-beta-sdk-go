package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
type MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderGetQueryParameters the list of Mobile threat Defense connectors configured by the tenant.
type MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderGetQueryParameters
}
// MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderInternal instantiates a new MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder and sets the default values.
func NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder) {
    m := &MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/mobileThreatDefenseConnectors/{mobileThreatDefenseConnector%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder instantiates a new MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder and sets the default values.
func NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mobileThreatDefenseConnectors for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of Mobile threat Defense connectors configured by the tenant.
// returns a MobileThreatDefenseConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileThreatDefenseConnectorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileThreatDefenseConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileThreatDefenseConnectorable), nil
}
// Patch update the navigation property mobileThreatDefenseConnectors in deviceManagement
// returns a MobileThreatDefenseConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileThreatDefenseConnectorable, requestConfiguration *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileThreatDefenseConnectorable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileThreatDefenseConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileThreatDefenseConnectorable), nil
}
// ToDeleteRequestInformation delete navigation property mobileThreatDefenseConnectors for deviceManagement
// returns a *RequestInformation when successful
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of Mobile threat Defense connectors configured by the tenant.
// returns a *RequestInformation when successful
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mobileThreatDefenseConnectors in deviceManagement
// returns a *RequestInformation when successful
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileThreatDefenseConnectorable, requestConfiguration *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder when successful
func (m *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder) WithUrl(rawUrl string)(*MobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder) {
    return NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
