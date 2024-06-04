package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder provides operations to manage the mailboxRestoreArtifacts property of the microsoft.graph.exchangeRestoreSession entity.
type BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderGetQueryParameters a collection of restore points and destination details that can be used to restore Exchange mailboxes.
type BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderGetQueryParameters
}
// BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderInternal instantiates a new BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder and sets the default values.
func NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder) {
    m := &BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/exchangeRestoreSessions/{exchangeRestoreSession%2Did}/mailboxRestoreArtifacts/{mailboxRestoreArtifact%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder instantiates a new BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder and sets the default values.
func NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mailboxRestoreArtifacts for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of restore points and destination details that can be used to restore Exchange mailboxes.
// returns a MailboxRestoreArtifactable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property mailboxRestoreArtifacts in solutions
// returns a MailboxRestoreArtifactable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactable, requestConfiguration *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// RestorePoint provides operations to manage the restorePoint property of the microsoft.graph.restoreArtifactBase entity.
// returns a *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsItemRestorepointRestorePointRequestBuilder when successful
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder) RestorePoint()(*BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsItemRestorepointRestorePointRequestBuilder) {
    return NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsItemRestorepointRestorePointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property mailboxRestoreArtifacts for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of restore points and destination details that can be used to restore Exchange mailboxes.
// returns a *RequestInformation when successful
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mailboxRestoreArtifacts in solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxRestoreArtifactable, requestConfiguration *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder when successful
func (m *BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder) {
    return NewBackuprestoreExchangerestoresessionsItemMailboxrestoreartifactsMailboxRestoreArtifactItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
