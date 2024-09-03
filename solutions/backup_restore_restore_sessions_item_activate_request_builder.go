package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreRestoreSessionsItemActivateRequestBuilder provides operations to call the activate method.
type BackupRestoreRestoreSessionsItemActivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreRestoreSessionsItemActivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreRestoreSessionsItemActivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackupRestoreRestoreSessionsItemActivateRequestBuilderInternal instantiates a new BackupRestoreRestoreSessionsItemActivateRequestBuilder and sets the default values.
func NewBackupRestoreRestoreSessionsItemActivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreRestoreSessionsItemActivateRequestBuilder) {
    m := &BackupRestoreRestoreSessionsItemActivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/restoreSessions/{restoreSessionBase%2Did}/activate", pathParameters),
    }
    return m
}
// NewBackupRestoreRestoreSessionsItemActivateRequestBuilder instantiates a new BackupRestoreRestoreSessionsItemActivateRequestBuilder and sets the default values.
func NewBackupRestoreRestoreSessionsItemActivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreRestoreSessionsItemActivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreRestoreSessionsItemActivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post activate a draft restoreSessionBase object to restore a protection unit. The following points apply to restoring a protection unit:
// returns a RestoreSessionBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/restoresessionbase-activate?view=graph-rest-beta
func (m *BackupRestoreRestoreSessionsItemActivateRequestBuilder) Post(ctx context.Context, requestConfiguration *BackupRestoreRestoreSessionsItemActivateRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestoreSessionBaseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRestoreSessionBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestoreSessionBaseable), nil
}
// ToPostRequestInformation activate a draft restoreSessionBase object to restore a protection unit. The following points apply to restoring a protection unit:
// returns a *RequestInformation when successful
func (m *BackupRestoreRestoreSessionsItemActivateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreRestoreSessionsItemActivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *BackupRestoreRestoreSessionsItemActivateRequestBuilder when successful
func (m *BackupRestoreRestoreSessionsItemActivateRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreRestoreSessionsItemActivateRequestBuilder) {
    return NewBackupRestoreRestoreSessionsItemActivateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
