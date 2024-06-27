package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder provides operations to manage the siteRestoreArtifacts property of the microsoft.graph.sharePointRestoreSession entity.
type BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderGetQueryParameters a collection of restore points and destination details that can be used to restore SharePoint sites.
type BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderGetQueryParameters
}
// BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderInternal instantiates a new BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder and sets the default values.
func NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder) {
    m := &BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/sharePointRestoreSessions/{sharePointRestoreSession%2Did}/siteRestoreArtifacts/{siteRestoreArtifact%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder instantiates a new BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder and sets the default values.
func NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property siteRestoreArtifacts for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of restore points and destination details that can be used to restore SharePoint sites.
// returns a SiteRestoreArtifactable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteRestoreArtifactable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteRestoreArtifactFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteRestoreArtifactable), nil
}
// Patch update the navigation property siteRestoreArtifacts in solutions
// returns a SiteRestoreArtifactable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteRestoreArtifactable, requestConfiguration *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteRestoreArtifactable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteRestoreArtifactFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteRestoreArtifactable), nil
}
// RestorePoint provides operations to manage the restorePoint property of the microsoft.graph.restoreArtifactBase entity.
// returns a *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsItemRestorePointRequestBuilder when successful
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder) RestorePoint()(*BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsItemRestorePointRequestBuilder) {
    return NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsItemRestorePointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property siteRestoreArtifacts for solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of restore points and destination details that can be used to restore SharePoint sites.
// returns a *RequestInformation when successful
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property siteRestoreArtifacts in solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteRestoreArtifactable, requestConfiguration *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder when successful
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder) {
    return NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsSiteRestoreArtifactItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
