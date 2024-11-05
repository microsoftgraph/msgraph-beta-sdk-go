package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder provides operations to manage the mailboxRestoreArtifactsBulkAdditionRequests property of the microsoft.graph.exchangeRestoreSession entity.
type BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderGetQueryParameters get mailboxRestoreArtifactsBulkAdditionRequests from solutions
type BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderGetQueryParameters
}
// BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderInternal instantiates a new BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder and sets the default values.
func NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder) {
    m := &BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/exchangeRestoreSessions/{exchangeRestoreSession%2Did}/mailboxRestoreArtifactsBulkAdditionRequests/{mailboxRestoreArtifactsBulkAdditionRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder instantiates a new BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder and sets the default values.
func NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mailboxRestoreArtifactsBulkAdditionRequests for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get mailboxRestoreArtifactsBulkAdditionRequests from solutions
// returns a MailboxRestoreArtifactsBulkAdditionRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactsBulkAdditionRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property mailboxRestoreArtifactsBulkAdditionRequests in solutions
// returns a MailboxRestoreArtifactsBulkAdditionRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactsBulkAdditionRequestable, requestConfiguration *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactsBulkAdditionRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property mailboxRestoreArtifactsBulkAdditionRequests for solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get mailboxRestoreArtifactsBulkAdditionRequests from solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mailboxRestoreArtifactsBulkAdditionRequests in solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactsBulkAdditionRequestable, requestConfiguration *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder when successful
func (m *BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder) {
    return NewBackupRestoreExchangeRestoreSessionsItemMailboxRestoreArtifactsBulkAdditionRequestsMailboxRestoreArtifactsBulkAdditionRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
