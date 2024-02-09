package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder provides operations to manage the zebraFotaDeployments property of the microsoft.graph.deviceManagement entity.
type ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderGetQueryParameters collection of ZebraFotaDeployments associated with account.
type ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderGetQueryParameters
}
// ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
// returns a *ZebraFotaDeploymentsItemCancelRequestBuilder when successful
func (m *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) Cancel()(*ZebraFotaDeploymentsItemCancelRequestBuilder) {
    return NewZebraFotaDeploymentsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderInternal instantiates a new ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder and sets the default values.
func NewZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) {
    m := &ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/zebraFotaDeployments/{zebraFotaDeployment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder instantiates a new ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder and sets the default values.
func NewZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property zebraFotaDeployments for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaDeploymentable, error) {
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
func (m *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaDeploymentable, requestConfiguration *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaDeploymentable, error) {
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
func (m *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/zebraFotaDeployments/{zebraFotaDeployment%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of ZebraFotaDeployments associated with account.
// returns a *RequestInformation when successful
func (m *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ZebraFotaDeploymentable, requestConfiguration *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/zebraFotaDeployments/{zebraFotaDeployment%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder when successful
func (m *ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) WithUrl(rawUrl string)(*ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) {
    return NewZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
