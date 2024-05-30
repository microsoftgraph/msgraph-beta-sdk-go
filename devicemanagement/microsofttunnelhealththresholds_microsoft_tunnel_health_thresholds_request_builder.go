package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder provides operations to manage the microsoftTunnelHealthThresholds property of the microsoft.graph.deviceManagement entity.
type MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderGetQueryParameters collection of MicrosoftTunnelHealthThreshold settings associated with account.
type MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderGetQueryParameters struct {
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
// MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderGetQueryParameters
}
// MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMicrosoftTunnelHealthThresholdId provides operations to manage the microsoftTunnelHealthThresholds property of the microsoft.graph.deviceManagement entity.
// returns a *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder when successful
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder) ByMicrosoftTunnelHealthThresholdId(microsoftTunnelHealthThresholdId string)(*MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if microsoftTunnelHealthThresholdId != "" {
        urlTplParams["microsoftTunnelHealthThreshold%2Did"] = microsoftTunnelHealthThresholdId
    }
    return NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderInternal instantiates a new MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder and sets the default values.
func NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder) {
    m := &MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelHealthThresholds{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder instantiates a new MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder and sets the default values.
func NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MicrosofttunnelhealththresholdsCountRequestBuilder when successful
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder) Count()(*MicrosofttunnelhealththresholdsCountRequestBuilder) {
    return NewMicrosofttunnelhealththresholdsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of MicrosoftTunnelHealthThreshold settings associated with account.
// returns a MicrosoftTunnelHealthThresholdCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelHealthThresholdCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelHealthThresholdCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelHealthThresholdCollectionResponseable), nil
}
// Post create new navigation property to microsoftTunnelHealthThresholds for deviceManagement
// returns a MicrosoftTunnelHealthThresholdable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelHealthThresholdable, requestConfiguration *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelHealthThresholdable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelHealthThresholdFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelHealthThresholdable), nil
}
// ToGetRequestInformation collection of MicrosoftTunnelHealthThreshold settings associated with account.
// returns a *RequestInformation when successful
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to microsoftTunnelHealthThresholds for deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelHealthThresholdable, requestConfiguration *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder when successful
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder) WithUrl(rawUrl string)(*MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder) {
    return NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
