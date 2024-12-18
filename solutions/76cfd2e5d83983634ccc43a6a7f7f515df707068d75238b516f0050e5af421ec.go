package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder provides operations to manage the siteRestoreArtifactsBulkAdditionRequests property of the microsoft.graph.sharePointRestoreSession entity.
type BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderGetQueryParameters get a list of the siteRestoreArtifactsBulkAdditionRequest objects associated with a sharePointRestoreSession. The siteWebUrls property is deliberately omitted from the response body in order to limit the response size.
type BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderGetQueryParameters struct {
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
// BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderGetQueryParameters
}
// BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySiteRestoreArtifactsBulkAdditionRequestId provides operations to manage the siteRestoreArtifactsBulkAdditionRequests property of the microsoft.graph.sharePointRestoreSession entity.
// returns a *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsSiteRestoreArtifactsBulkAdditionRequestItemRequestBuilder when successful
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder) BySiteRestoreArtifactsBulkAdditionRequestId(siteRestoreArtifactsBulkAdditionRequestId string)(*BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsSiteRestoreArtifactsBulkAdditionRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if siteRestoreArtifactsBulkAdditionRequestId != "" {
        urlTplParams["siteRestoreArtifactsBulkAdditionRequest%2Did"] = siteRestoreArtifactsBulkAdditionRequestId
    }
    return NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsSiteRestoreArtifactsBulkAdditionRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderInternal instantiates a new BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder and sets the default values.
func NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder) {
    m := &BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/sharePointRestoreSessions/{sharePointRestoreSession%2Did}/siteRestoreArtifactsBulkAdditionRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder instantiates a new BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder and sets the default values.
func NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsCountRequestBuilder when successful
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder) Count()(*BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsCountRequestBuilder) {
    return NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the siteRestoreArtifactsBulkAdditionRequest objects associated with a sharePointRestoreSession. The siteWebUrls property is deliberately omitted from the response body in order to limit the response size.
// returns a SiteRestoreArtifactsBulkAdditionRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/sharepointrestoresession-list-siterestoreartifactsbulkadditionrequests?view=graph-rest-beta
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteRestoreArtifactsBulkAdditionRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteRestoreArtifactsBulkAdditionRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteRestoreArtifactsBulkAdditionRequestCollectionResponseable), nil
}
// Post create a new siteRestoreArtifactsBulkAdditionRequest object associated with a sharePointRestoreSession. The initial status upon creation of the restore session is active. When all the sites are added to the corresponding SharePoint restore session and the restore session is activated, the status becomes completed. If any failures are encountered during resource resolution, the status of the restore session becomes completedWithErrors.
// returns a SiteRestoreArtifactsBulkAdditionRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/sharepointrestoresession-post-siterestoreartifactsbulkadditionrequests?view=graph-rest-beta
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteRestoreArtifactsBulkAdditionRequestable, requestConfiguration *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteRestoreArtifactsBulkAdditionRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteRestoreArtifactsBulkAdditionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteRestoreArtifactsBulkAdditionRequestable), nil
}
// ToGetRequestInformation get a list of the siteRestoreArtifactsBulkAdditionRequest objects associated with a sharePointRestoreSession. The siteWebUrls property is deliberately omitted from the response body in order to limit the response size.
// returns a *RequestInformation when successful
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new siteRestoreArtifactsBulkAdditionRequest object associated with a sharePointRestoreSession. The initial status upon creation of the restore session is active. When all the sites are added to the corresponding SharePoint restore session and the restore session is activated, the status becomes completed. If any failures are encountered during resource resolution, the status of the restore session becomes completedWithErrors.
// returns a *RequestInformation when successful
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteRestoreArtifactsBulkAdditionRequestable, requestConfiguration *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder when successful
func (m *BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder) {
    return NewBackupRestoreSharePointRestoreSessionsItemSiteRestoreArtifactsBulkAdditionRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
