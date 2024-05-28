package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder provides operations to manage the microsoftTunnelHealthThresholds property of the microsoft.graph.deviceManagement entity.
type MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderGetQueryParameters collection of MicrosoftTunnelHealthThreshold settings associated with account.
type MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderGetQueryParameters
}
// MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderInternal instantiates a new MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder and sets the default values.
func NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) {
    m := &MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelHealthThresholds/{microsoftTunnelHealthThreshold%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder instantiates a new MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder and sets the default values.
func NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property microsoftTunnelHealthThresholds for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of MicrosoftTunnelHealthThreshold settings associated with account.
// returns a MicrosoftTunnelHealthThresholdable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelHealthThresholdable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property microsoftTunnelHealthThresholds in deviceManagement
// returns a MicrosoftTunnelHealthThresholdable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelHealthThresholdable, requestConfiguration *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelHealthThresholdable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property microsoftTunnelHealthThresholds for deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of MicrosoftTunnelHealthThreshold settings associated with account.
// returns a *RequestInformation when successful
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property microsoftTunnelHealthThresholds in deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelHealthThresholdable, requestConfiguration *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder when successful
func (m *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) WithUrl(rawUrl string)(*MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) {
    return NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
