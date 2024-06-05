package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder provides operations to manage the mailboxRestoreArtifacts property of the microsoft.graph.exchangeRestoreSession entity.
type BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderGetQueryParameters get a list of the mailboxRestoreArtifact objects that are associated with an exchangeRestoreSession in a tenant.
type BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderGetQueryParameters struct {
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
// BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderGetQueryParameters
}
// BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMailboxRestoreArtifactId provides operations to manage the mailboxRestoreArtifacts property of the microsoft.graph.exchangeRestoreSession entity.
// returns a *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder when successful
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder) ByMailboxRestoreArtifactId(mailboxRestoreArtifactId string)(*BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mailboxRestoreArtifactId != "" {
        urlTplParams["mailboxRestoreArtifact%2Did"] = mailboxRestoreArtifactId
    }
    return NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderInternal instantiates a new BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder and sets the default values.
func NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder) {
    m := &BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/exchangeRestoreSessions/{exchangeRestoreSession%2Did}/mailboxRestoreArtifacts{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder instantiates a new BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder and sets the default values.
func NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsCountRequestBuilder when successful
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder) Count()(*BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsCountRequestBuilder) {
    return NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the mailboxRestoreArtifact objects that are associated with an exchangeRestoreSession in a tenant.
// returns a MailboxRestoreArtifactCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/exchangerestoresession-list-mailboxrestoreartifacts?view=graph-rest-beta
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxRestoreArtifactCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactCollectionResponseable), nil
}
// Post create new navigation property to mailboxRestoreArtifacts for solutions
// returns a MailboxRestoreArtifactable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactable, requestConfiguration *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxRestoreArtifactFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactable), nil
}
// ToGetRequestInformation get a list of the mailboxRestoreArtifact objects that are associated with an exchangeRestoreSession in a tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to mailboxRestoreArtifacts for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactable, requestConfiguration *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder when successful
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder) {
    return NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
