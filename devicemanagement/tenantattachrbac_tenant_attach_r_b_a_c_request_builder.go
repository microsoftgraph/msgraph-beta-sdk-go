package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TenantattachrbacTenantAttachRBACRequestBuilder provides operations to manage the tenantAttachRBAC property of the microsoft.graph.deviceManagement entity.
type TenantattachrbacTenantAttachRBACRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TenantattachrbacTenantAttachRBACRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantattachrbacTenantAttachRBACRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TenantattachrbacTenantAttachRBACRequestBuilderGetQueryParameters tenantAttach RBAC Enablement
type TenantattachrbacTenantAttachRBACRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantattachrbacTenantAttachRBACRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantattachrbacTenantAttachRBACRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantattachrbacTenantAttachRBACRequestBuilderGetQueryParameters
}
// TenantattachrbacTenantAttachRBACRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantattachrbacTenantAttachRBACRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTenantattachrbacTenantAttachRBACRequestBuilderInternal instantiates a new TenantattachrbacTenantAttachRBACRequestBuilder and sets the default values.
func NewTenantattachrbacTenantAttachRBACRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantattachrbacTenantAttachRBACRequestBuilder) {
    m := &TenantattachrbacTenantAttachRBACRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/tenantAttachRBAC{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTenantattachrbacTenantAttachRBACRequestBuilder instantiates a new TenantattachrbacTenantAttachRBACRequestBuilder and sets the default values.
func NewTenantattachrbacTenantAttachRBACRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantattachrbacTenantAttachRBACRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantattachrbacTenantAttachRBACRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property tenantAttachRBAC for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TenantattachrbacTenantAttachRBACRequestBuilder) Delete(ctx context.Context, requestConfiguration *TenantattachrbacTenantAttachRBACRequestBuilderDeleteRequestConfiguration)(error) {
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
// Enable provides operations to call the enable method.
// returns a *TenantattachrbacEnableRequestBuilder when successful
func (m *TenantattachrbacTenantAttachRBACRequestBuilder) Enable()(*TenantattachrbacEnableRequestBuilder) {
    return NewTenantattachrbacEnableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get tenantAttach RBAC Enablement
// returns a TenantAttachRBACable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TenantattachrbacTenantAttachRBACRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantattachrbacTenantAttachRBACRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantAttachRBACable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTenantAttachRBACFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantAttachRBACable), nil
}
// GetState provides operations to call the getState method.
// returns a *TenantattachrbacGetstateGetStateRequestBuilder when successful
func (m *TenantattachrbacTenantAttachRBACRequestBuilder) GetState()(*TenantattachrbacGetstateGetStateRequestBuilder) {
    return NewTenantattachrbacGetstateGetStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property tenantAttachRBAC in deviceManagement
// returns a TenantAttachRBACable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TenantattachrbacTenantAttachRBACRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantAttachRBACable, requestConfiguration *TenantattachrbacTenantAttachRBACRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantAttachRBACable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTenantAttachRBACFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantAttachRBACable), nil
}
// ToDeleteRequestInformation delete navigation property tenantAttachRBAC for deviceManagement
// returns a *RequestInformation when successful
func (m *TenantattachrbacTenantAttachRBACRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TenantattachrbacTenantAttachRBACRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation tenantAttach RBAC Enablement
// returns a *RequestInformation when successful
func (m *TenantattachrbacTenantAttachRBACRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TenantattachrbacTenantAttachRBACRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property tenantAttachRBAC in deviceManagement
// returns a *RequestInformation when successful
func (m *TenantattachrbacTenantAttachRBACRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantAttachRBACable, requestConfiguration *TenantattachrbacTenantAttachRBACRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TenantattachrbacTenantAttachRBACRequestBuilder when successful
func (m *TenantattachrbacTenantAttachRBACRequestBuilder) WithUrl(rawUrl string)(*TenantattachrbacTenantAttachRBACRequestBuilder) {
    return NewTenantattachrbacTenantAttachRBACRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
