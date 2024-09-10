package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreEnableRequestBuilder provides operations to call the enable method.
type BackupRestoreEnableRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreEnableRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreEnableRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackupRestoreEnableRequestBuilderInternal instantiates a new BackupRestoreEnableRequestBuilder and sets the default values.
func NewBackupRestoreEnableRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreEnableRequestBuilder) {
    m := &BackupRestoreEnableRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/enable", pathParameters),
    }
    return m
}
// NewBackupRestoreEnableRequestBuilder instantiates a new BackupRestoreEnableRequestBuilder and sets the default values.
func NewBackupRestoreEnableRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreEnableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreEnableRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enable the Microsoft 365 Backup Storage service for a tenant. Before you call this API, call List protection policies to initialize the data store in the tenant. Data store initialization takes about 5 minutes. If you call this API before the data store is initialized, the call results in an error.
// returns a ServiceStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/backuprestoreroot-enable?view=graph-rest-beta
func (m *BackupRestoreEnableRequestBuilder) Post(ctx context.Context, body BackupRestoreEnablePostRequestBodyable, requestConfiguration *BackupRestoreEnableRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceStatusable), nil
}
// ToPostRequestInformation enable the Microsoft 365 Backup Storage service for a tenant. Before you call this API, call List protection policies to initialize the data store in the tenant. Data store initialization takes about 5 minutes. If you call this API before the data store is initialized, the call results in an error.
// returns a *RequestInformation when successful
func (m *BackupRestoreEnableRequestBuilder) ToPostRequestInformation(ctx context.Context, body BackupRestoreEnablePostRequestBodyable, requestConfiguration *BackupRestoreEnableRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreEnableRequestBuilder when successful
func (m *BackupRestoreEnableRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreEnableRequestBuilder) {
    return NewBackupRestoreEnableRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
