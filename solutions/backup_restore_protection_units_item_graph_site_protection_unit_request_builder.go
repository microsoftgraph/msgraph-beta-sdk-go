package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder casts the previous resource to siteProtectionUnit.
type BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilderGetQueryParameters get a list of the siteProtectionUnit objects that are associated with a sharePointProtectionPolicy.
type BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilderGetQueryParameters
}
// NewBackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilderInternal instantiates a new BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder and sets the default values.
func NewBackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder) {
    m := &BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/protectionUnits/{protectionUnitBase%2Did}/graph.siteProtectionUnit{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder instantiates a new BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder and sets the default values.
func NewBackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a list of the siteProtectionUnit objects that are associated with a sharePointProtectionPolicy.
// returns a SiteProtectionUnitable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/backuprestoreroot-list-siteprotectionunits?view=graph-rest-beta
func (m *BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitable, error) {
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
// ToGetRequestInformation get a list of the siteProtectionUnit objects that are associated with a sharePointProtectionPolicy.
// returns a *RequestInformation when successful
func (m *BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder when successful
func (m *BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder) {
    return NewBackupRestoreProtectionUnitsItemGraphSiteProtectionUnitRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
