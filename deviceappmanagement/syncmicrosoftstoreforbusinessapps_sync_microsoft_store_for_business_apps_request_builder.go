package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder provides operations to call the syncMicrosoftStoreForBusinessApps method.
type SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal instantiates a new SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder and sets the default values.
func NewSyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    m := &SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/syncMicrosoftStoreForBusinessApps", pathParameters),
    }
    return m
}
// NewSyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder instantiates a new SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder and sets the default values.
func NewSyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post syncs Intune account with Microsoft Store For Business
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder) Post(ctx context.Context, requestConfiguration *SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation syncs Intune account with Microsoft Store For Business
// returns a *RequestInformation when successful
func (m *SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder when successful
func (m *SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder) WithUrl(rawUrl string)(*SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    return NewSyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
