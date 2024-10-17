package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder provides operations to manage the siteProtectionUnitsBulkAdditionJobs property of the microsoft.graph.backupRestoreRoot entity.
type BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetQueryParameters get siteProtectionUnitsBulkAdditionJobs from solutions
type BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetQueryParameters
}
// BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderInternal instantiates a new BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder and sets the default values.
func NewBackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) {
    m := &BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/siteProtectionUnitsBulkAdditionJobs/{siteProtectionUnitsBulkAdditionJob%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder instantiates a new BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder and sets the default values.
func NewBackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property siteProtectionUnitsBulkAdditionJobs for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get siteProtectionUnitsBulkAdditionJobs from solutions
// returns a SiteProtectionUnitsBulkAdditionJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitsBulkAdditionJobable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteProtectionUnitsBulkAdditionJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitsBulkAdditionJobable), nil
}
// Patch update the navigation property siteProtectionUnitsBulkAdditionJobs in solutions
// returns a SiteProtectionUnitsBulkAdditionJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitsBulkAdditionJobable, requestConfiguration *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitsBulkAdditionJobable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteProtectionUnitsBulkAdditionJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitsBulkAdditionJobable), nil
}
// ToDeleteRequestInformation delete navigation property siteProtectionUnitsBulkAdditionJobs for solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get siteProtectionUnitsBulkAdditionJobs from solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property siteProtectionUnitsBulkAdditionJobs in solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitsBulkAdditionJobable, requestConfiguration *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder when successful
func (m *BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) {
    return NewBackupRestoreSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
