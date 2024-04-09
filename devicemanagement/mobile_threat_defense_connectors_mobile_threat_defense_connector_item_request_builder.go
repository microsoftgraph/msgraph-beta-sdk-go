package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
type MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderGetQueryParameters the list of Mobile threat Defense connectors configured by the tenant.
type MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderGetQueryParameters
}
// MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderInternal instantiates a new MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder and sets the default values.
func NewMobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder) {
    m := &MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/mobileThreatDefenseConnectors/{mobileThreatDefenseConnector%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder instantiates a new MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder and sets the default values.
func NewMobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mobileThreatDefenseConnectors for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileThreatDefenseConnectorable, error) {
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
func (m *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileThreatDefenseConnectorable, requestConfiguration *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileThreatDefenseConnectorable, error) {
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
func (m *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileThreatDefenseConnectorable, requestConfiguration *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder when successful
func (m *MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder) WithUrl(rawUrl string)(*MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder) {
    return NewMobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
