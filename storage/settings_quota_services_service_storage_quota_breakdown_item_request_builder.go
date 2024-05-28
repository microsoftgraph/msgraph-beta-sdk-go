package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder provides operations to manage the services property of the microsoft.graph.unifiedStorageQuota entity.
type SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetQueryParameters the breakdown of services contributing to the user's quota usage.
type SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetQueryParameters
}
// SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderInternal instantiates a new SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder and sets the default values.
func NewSettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) {
    m := &SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/settings/quota/services/{serviceStorageQuotaBreakdown%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewSettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder instantiates a new SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder and sets the default values.
func NewSettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property services for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the breakdown of services contributing to the user's quota usage.
// returns a ServiceStorageQuotaBreakdownable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceStorageQuotaBreakdownable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceStorageQuotaBreakdownFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceStorageQuotaBreakdownable), nil
}
// Patch update the navigation property services in storage
// returns a ServiceStorageQuotaBreakdownable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceStorageQuotaBreakdownable, requestConfiguration *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceStorageQuotaBreakdownable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceStorageQuotaBreakdownFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceStorageQuotaBreakdownable), nil
}
// ToDeleteRequestInformation delete navigation property services for storage
// returns a *RequestInformation when successful
func (m *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the breakdown of services contributing to the user's quota usage.
// returns a *RequestInformation when successful
func (m *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property services in storage
// returns a *RequestInformation when successful
func (m *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceStorageQuotaBreakdownable, requestConfiguration *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder when successful
func (m *SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) WithUrl(rawUrl string)(*SettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) {
    return NewSettingsQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
