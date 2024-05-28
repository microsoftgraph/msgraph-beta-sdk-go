package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder provides operations to manage the zebraFotaDeployments property of the microsoft.graph.deviceManagement entity.
type ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderGetQueryParameters collection of ZebraFotaDeployments associated with account.
type ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderGetQueryParameters
}
// ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
// returns a *ZebrafotadeploymentsItemCancelRequestBuilder when successful
func (m *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder) Cancel()(*ZebrafotadeploymentsItemCancelRequestBuilder) {
    return NewZebrafotadeploymentsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderInternal instantiates a new ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder and sets the default values.
func NewZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder) {
    m := &ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/zebraFotaDeployments/{zebraFotaDeployment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder instantiates a new ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder and sets the default values.
func NewZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property zebraFotaDeployments for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of ZebraFotaDeployments associated with account.
// returns a ZebraFotaDeploymentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaDeploymentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateZebraFotaDeploymentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaDeploymentable), nil
}
// Patch update the navigation property zebraFotaDeployments in deviceManagement
// returns a ZebraFotaDeploymentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaDeploymentable, requestConfiguration *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaDeploymentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateZebraFotaDeploymentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaDeploymentable), nil
}
// ToDeleteRequestInformation delete navigation property zebraFotaDeployments for deviceManagement
// returns a *RequestInformation when successful
func (m *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of ZebraFotaDeployments associated with account.
// returns a *RequestInformation when successful
func (m *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property zebraFotaDeployments in deviceManagement
// returns a *RequestInformation when successful
func (m *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaDeploymentable, requestConfiguration *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder when successful
func (m *ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder) WithUrl(rawUrl string)(*ZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder) {
    return NewZebrafotadeploymentsZebraFotaDeploymentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
