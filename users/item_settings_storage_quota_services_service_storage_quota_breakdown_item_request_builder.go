package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder provides operations to manage the services property of the microsoft.graph.unifiedStorageQuota entity.
type ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetQueryParameters the breakdown of services contributing to the user's quota usage.
type ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetQueryParameters
}
// ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderInternal instantiates a new ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder and sets the default values.
func NewItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) {
    m := &ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/settings/storage/quota/services/{serviceStorageQuotaBreakdown%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder instantiates a new ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder and sets the default values.
func NewItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property services for users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ServiceStorageQuotaBreakdownable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceStorageQuotaBreakdownable, error) {
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
// Patch update the navigation property services in users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ServiceStorageQuotaBreakdownable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceStorageQuotaBreakdownable, requestConfiguration *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceStorageQuotaBreakdownable, error) {
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
// ToDeleteRequestInformation delete navigation property services for users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the breakdown of services contributing to the user's quota usage.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property services in users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceStorageQuotaBreakdownable, requestConfiguration *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder when successful
func (m *ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder) {
    return NewItemSettingsStorageQuotaServicesServiceStorageQuotaBreakdownItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
