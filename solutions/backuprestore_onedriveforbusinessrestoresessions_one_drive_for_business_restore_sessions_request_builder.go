package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder provides operations to manage the oneDriveForBusinessRestoreSessions property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderGetQueryParameters the list of OneDrive for Business restore sessions available in the tenant.
type BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderGetQueryParameters struct {
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
// BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderGetQueryParameters
}
// BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByOneDriveForBusinessRestoreSessionId provides operations to manage the oneDriveForBusinessRestoreSessions property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder) ByOneDriveForBusinessRestoreSessionId(oneDriveForBusinessRestoreSessionId string)(*BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if oneDriveForBusinessRestoreSessionId != "" {
        urlTplParams["oneDriveForBusinessRestoreSession%2Did"] = oneDriveForBusinessRestoreSessionId
    }
    return NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderInternal instantiates a new BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder and sets the default values.
func NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder) {
    m := &BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/oneDriveForBusinessRestoreSessions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder instantiates a new BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder and sets the default values.
func NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackuprestoreOnedriveforbusinessrestoresessionsCountRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder) Count()(*BackuprestoreOnedriveforbusinessrestoresessionsCountRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessrestoresessionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of OneDrive for Business restore sessions available in the tenant.
// returns a OneDriveForBusinessRestoreSessionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessRestoreSessionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOneDriveForBusinessRestoreSessionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessRestoreSessionCollectionResponseable), nil
}
// Post create new navigation property to oneDriveForBusinessRestoreSessions for solutions
// returns a OneDriveForBusinessRestoreSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessRestoreSessionable, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessRestoreSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOneDriveForBusinessRestoreSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessRestoreSessionable), nil
}
// ToGetRequestInformation the list of OneDrive for Business restore sessions available in the tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to oneDriveForBusinessRestoreSessions for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessRestoreSessionable, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
