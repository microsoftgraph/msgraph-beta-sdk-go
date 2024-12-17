package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder provides operations to manage the mailboxRestoreArtifactsBulkAdditionRequests property of the microsoft.graph.exchangeRestoreSession entity.
type BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderGetQueryParameters get a list of the maiboxRestoreArtifactsBulkAdditionRequest objects associated with an exchangeRestoreSession. The mailboxes property is deliberately omitted from the response body in order to limit the response size.
type BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderGetQueryParameters struct {
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
// BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderGetQueryParameters
}
// BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMailboxRestoreArtifactsBulkAdditionRequestId provides operations to manage the mailboxRestoreArtifactsBulkAdditionRequests property of the microsoft.graph.exchangeRestoreSession entity.
// returns a *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder when successful
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder) ByMailboxRestoreArtifactsBulkAdditionRequestId(mailboxRestoreArtifactsBulkAdditionRequestId string)(*BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mailboxRestoreArtifactsBulkAdditionRequestId != "" {
        urlTplParams["mailboxRestoreArtifactsBulkAdditionRequest%2Did"] = mailboxRestoreArtifactsBulkAdditionRequestId
    }
    return NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderInternal instantiates a new BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder and sets the default values.
func NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder) {
    m := &BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/exchangeRestoreSessions/{exchangeRestoreSession%2Did}/mailboxRestoreArtifactsBulkAdditionRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder instantiates a new BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder and sets the default values.
func NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsCountRequestBuilder when successful
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder) Count()(*BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsCountRequestBuilder) {
    return NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the maiboxRestoreArtifactsBulkAdditionRequest objects associated with an exchangeRestoreSession. The mailboxes property is deliberately omitted from the response body in order to limit the response size.
// returns a MailboxRestoreArtifactsBulkAdditionRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/exchangerestoresession-list-mailboxrestoreartifactsbulkadditionrequests?view=graph-rest-beta
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactsBulkAdditionRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxRestoreArtifactsBulkAdditionRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactsBulkAdditionRequestCollectionResponseable), nil
}
// Post create a new mailboxRestoreArtifactsBulkAdditionRequest object associated with an exchangeRestoreSession. The initial status upon creation of the restore session is active. When all the mailboxes are added to the corresponding Exchange restore session and the restore session is activated, the status becomes completed. If any failures are encountered during resource resolution, the status of the restore session becomes completedWithErrors.
// returns a MailboxRestoreArtifactsBulkAdditionRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/exchangerestoresession-post-mailboxrestoreartifactsbulkadditionrequests?view=graph-rest-beta
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactsBulkAdditionRequestable, requestConfiguration *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactsBulkAdditionRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxRestoreArtifactsBulkAdditionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactsBulkAdditionRequestable), nil
}
// ToGetRequestInformation get a list of the maiboxRestoreArtifactsBulkAdditionRequest objects associated with an exchangeRestoreSession. The mailboxes property is deliberately omitted from the response body in order to limit the response size.
// returns a *RequestInformation when successful
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new mailboxRestoreArtifactsBulkAdditionRequest object associated with an exchangeRestoreSession. The initial status upon creation of the restore session is active. When all the mailboxes are added to the corresponding Exchange restore session and the restore session is activated, the status becomes completed. If any failures are encountered during resource resolution, the status of the restore session becomes completedWithErrors.
// returns a *RequestInformation when successful
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactsBulkAdditionRequestable, requestConfiguration *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder when successful
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder) {
    return NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
