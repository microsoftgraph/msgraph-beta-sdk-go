package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder provides operations to manage the sharePointRestoreSessions property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderGetQueryParameters the list of SharePoint restore sessions available in the tenant.
type BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderGetQueryParameters struct {
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
// BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderGetQueryParameters
}
// BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySharePointRestoreSessionId provides operations to manage the sharePointRestoreSessions property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionItemRequestBuilder when successful
func (m *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder) BySharePointRestoreSessionId(sharePointRestoreSessionId string)(*BackuprestoreSharepointrestoresessionsSharePointRestoreSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if sharePointRestoreSessionId != "" {
        urlTplParams["sharePointRestoreSession%2Did"] = sharePointRestoreSessionId
    }
    return NewBackuprestoreSharepointrestoresessionsSharePointRestoreSessionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderInternal instantiates a new BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder and sets the default values.
func NewBackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder) {
    m := &BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/sharePointRestoreSessions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder instantiates a new BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder and sets the default values.
func NewBackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackuprestoreSharepointrestoresessionsCountRequestBuilder when successful
func (m *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder) Count()(*BackuprestoreSharepointrestoresessionsCountRequestBuilder) {
    return NewBackuprestoreSharepointrestoresessionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of SharePoint restore sessions available in the tenant.
// returns a SharePointRestoreSessionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharePointRestoreSessionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSharePointRestoreSessionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharePointRestoreSessionCollectionResponseable), nil
}
// Post create a new sharePointRestoreSession object.
// returns a SharePointRestoreSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/backuprestoreroot-post-sharepointrestoresessions?view=graph-rest-beta
func (m *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharePointRestoreSessionable, requestConfiguration *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharePointRestoreSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSharePointRestoreSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharePointRestoreSessionable), nil
}
// ToGetRequestInformation the list of SharePoint restore sessions available in the tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new sharePointRestoreSession object.
// returns a *RequestInformation when successful
func (m *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharePointRestoreSessionable, requestConfiguration *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder when successful
func (m *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder) {
    return NewBackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
