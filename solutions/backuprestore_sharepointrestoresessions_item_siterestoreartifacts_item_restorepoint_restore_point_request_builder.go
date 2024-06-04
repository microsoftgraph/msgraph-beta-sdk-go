package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder provides operations to manage the restorePoint property of the microsoft.graph.restoreArtifactBase entity.
type BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilderGetQueryParameters represents the date and time when an artifact is protected by a protectionPolicy and can be restored.
type BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilderGetQueryParameters
}
// NewBackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilderInternal instantiates a new BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder and sets the default values.
func NewBackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder) {
    m := &BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/sharePointRestoreSessions/{sharePointRestoreSession%2Did}/siteRestoreArtifacts/{siteRestoreArtifact%2Did}/restorePoint{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder instantiates a new BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder and sets the default values.
func NewBackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilderInternal(urlParams, requestAdapter)
}
// Get represents the date and time when an artifact is protected by a protectionPolicy and can be restored.
// returns a RestorePointable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRestorePointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointable), nil
}
// ToGetRequestInformation represents the date and time when an artifact is protected by a protectionPolicy and can be restored.
// returns a *RequestInformation when successful
func (m *BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder when successful
func (m *BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder) {
    return NewBackuprestoreSharepointrestoresessionsItemSiterestoreartifactsItemRestorepointRestorePointRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
