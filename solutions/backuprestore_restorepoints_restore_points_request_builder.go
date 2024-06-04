package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreRestorepointsRestorePointsRequestBuilder provides operations to manage the restorePoints property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreRestorepointsRestorePointsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreRestorepointsRestorePointsRequestBuilderGetQueryParameters list of restore points in the tenant.
type BackuprestoreRestorepointsRestorePointsRequestBuilderGetQueryParameters struct {
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
// BackuprestoreRestorepointsRestorePointsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreRestorepointsRestorePointsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreRestorepointsRestorePointsRequestBuilderGetQueryParameters
}
// BackuprestoreRestorepointsRestorePointsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreRestorepointsRestorePointsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRestorePointId provides operations to manage the restorePoints property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreRestorepointsRestorePointItemRequestBuilder when successful
func (m *BackuprestoreRestorepointsRestorePointsRequestBuilder) ByRestorePointId(restorePointId string)(*BackuprestoreRestorepointsRestorePointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if restorePointId != "" {
        urlTplParams["restorePoint%2Did"] = restorePointId
    }
    return NewBackuprestoreRestorepointsRestorePointItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackuprestoreRestorepointsRestorePointsRequestBuilderInternal instantiates a new BackuprestoreRestorepointsRestorePointsRequestBuilder and sets the default values.
func NewBackuprestoreRestorepointsRestorePointsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreRestorepointsRestorePointsRequestBuilder) {
    m := &BackuprestoreRestorepointsRestorePointsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/restorePoints{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackuprestoreRestorepointsRestorePointsRequestBuilder instantiates a new BackuprestoreRestorepointsRestorePointsRequestBuilder and sets the default values.
func NewBackuprestoreRestorepointsRestorePointsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreRestorepointsRestorePointsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreRestorepointsRestorePointsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackuprestoreRestorepointsCountRequestBuilder when successful
func (m *BackuprestoreRestorepointsRestorePointsRequestBuilder) Count()(*BackuprestoreRestorepointsCountRequestBuilder) {
    return NewBackuprestoreRestorepointsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of restore points in the tenant.
// returns a RestorePointCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreRestorepointsRestorePointsRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreRestorepointsRestorePointsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRestorePointCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointCollectionResponseable), nil
}
// Post create new navigation property to restorePoints for solutions
// returns a RestorePointable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreRestorepointsRestorePointsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointable, requestConfiguration *BackuprestoreRestorepointsRestorePointsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRestorePointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointable), nil
}
// Search provides operations to call the search method.
// returns a *BackuprestoreRestorepointsSearchRequestBuilder when successful
func (m *BackuprestoreRestorepointsRestorePointsRequestBuilder) Search()(*BackuprestoreRestorepointsSearchRequestBuilder) {
    return NewBackuprestoreRestorepointsSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation list of restore points in the tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreRestorepointsRestorePointsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreRestorepointsRestorePointsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to restorePoints for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreRestorepointsRestorePointsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointable, requestConfiguration *BackuprestoreRestorepointsRestorePointsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreRestorepointsRestorePointsRequestBuilder when successful
func (m *BackuprestoreRestorepointsRestorePointsRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreRestorepointsRestorePointsRequestBuilder) {
    return NewBackuprestoreRestorepointsRestorePointsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
