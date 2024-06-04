package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder provides operations to manage the siteProtectionUnits property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderGetQueryParameters the list of site protection units in the tenant.
type BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderGetQueryParameters
}
// BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderInternal instantiates a new BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder and sets the default values.
func NewBackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder) {
    m := &BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/siteProtectionUnits/{siteProtectionUnit%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder instantiates a new BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder and sets the default values.
func NewBackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property siteProtectionUnits for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of site protection units in the tenant.
// returns a SiteProtectionUnitable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteProtectionUnitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitable), nil
}
// Patch update the navigation property siteProtectionUnits in solutions
// returns a SiteProtectionUnitable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitable, requestConfiguration *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteProtectionUnitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitable), nil
}
// ToDeleteRequestInformation delete navigation property siteProtectionUnits for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of site protection units in the tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property siteProtectionUnits in solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitable, requestConfiguration *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder when successful
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder) {
    return NewBackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
