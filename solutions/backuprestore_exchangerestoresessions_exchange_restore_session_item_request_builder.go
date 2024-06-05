package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder provides operations to manage the exchangeRestoreSessions property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderGetQueryParameters the list of Exchange restore sessions available in the tenant.
type BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderGetQueryParameters
}
// BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderInternal instantiates a new BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder and sets the default values.
func NewBackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder) {
    m := &BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/exchangeRestoreSessions/{exchangeRestoreSession%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder instantiates a new BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder and sets the default values.
func NewBackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property exchangeRestoreSessions for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of Exchange restore sessions available in the tenant.
// returns a ExchangeRestoreSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeRestoreSessionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExchangeRestoreSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeRestoreSessionable), nil
}
// MailboxRestoreArtifacts provides operations to manage the mailboxRestoreArtifacts property of the microsoft.graph.exchangeRestoreSession entity.
// returns a *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder when successful
func (m *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder) MailboxRestoreArtifacts()(*BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder) {
    return NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of an exchangeRestoreSession.
// returns a ExchangeRestoreSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/exchangerestoresession-update?view=graph-rest-beta
func (m *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeRestoreSessionable, requestConfiguration *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeRestoreSessionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExchangeRestoreSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeRestoreSessionable), nil
}
// ToDeleteRequestInformation delete navigation property exchangeRestoreSessions for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of Exchange restore sessions available in the tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an exchangeRestoreSession.
// returns a *RequestInformation when successful
func (m *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeRestoreSessionable, requestConfiguration *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder when successful
func (m *BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder) {
    return NewBackuprestoreExchangerestoresessionsExchangeRestoreSessionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
