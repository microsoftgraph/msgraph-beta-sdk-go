package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder provides operations to manage the siteProtectionUnits property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderGetQueryParameters get a list of the siteProtectionUnit objects that are associated with a sharePointProtectionPolicy.
type BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderGetQueryParameters
}
// BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySiteProtectionUnitId provides operations to manage the siteProtectionUnits property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder when successful
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder) BySiteProtectionUnitId(siteProtectionUnitId string)(*BackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if siteProtectionUnitId != "" {
        urlTplParams["siteProtectionUnit%2Did"] = siteProtectionUnitId
    }
    return NewBackuprestoreSiteprotectionunitsSiteProtectionUnitItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderInternal instantiates a new BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder and sets the default values.
func NewBackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder) {
    m := &BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/siteProtectionUnits{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder instantiates a new BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder and sets the default values.
func NewBackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackuprestoreSiteprotectionunitsCountRequestBuilder when successful
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder) Count()(*BackuprestoreSiteprotectionunitsCountRequestBuilder) {
    return NewBackuprestoreSiteprotectionunitsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the siteProtectionUnit objects that are associated with a sharePointProtectionPolicy.
// returns a SiteProtectionUnitCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/backuprestoreroot-list-siteprotectionunits?view=graph-rest-beta
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteProtectionUnitCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitCollectionResponseable), nil
}
// Post create new navigation property to siteProtectionUnits for solutions
// returns a SiteProtectionUnitable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitable, requestConfiguration *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to siteProtectionUnits for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionUnitable, requestConfiguration *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder when successful
func (m *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder) {
    return NewBackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
