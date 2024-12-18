package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder provides operations to manage the driveRestoreArtifactsBulkAdditionRequests property of the microsoft.graph.oneDriveForBusinessRestoreSession entity.
type BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderGetQueryParameters get a list of the driveRestoreArtifactsBulkAdditionRequest objects associated with a oneDriveForBusinessRestoreSession. The drives property is deliberately omitted from the response body in order to limit the response size.
type BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderGetQueryParameters struct {
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
// BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderGetQueryParameters
}
// BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDriveRestoreArtifactsBulkAdditionRequestId provides operations to manage the driveRestoreArtifactsBulkAdditionRequests property of the microsoft.graph.oneDriveForBusinessRestoreSession entity.
// returns a *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsDriveRestoreArtifactsBulkAdditionRequestItemRequestBuilder when successful
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder) ByDriveRestoreArtifactsBulkAdditionRequestId(driveRestoreArtifactsBulkAdditionRequestId string)(*BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsDriveRestoreArtifactsBulkAdditionRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if driveRestoreArtifactsBulkAdditionRequestId != "" {
        urlTplParams["driveRestoreArtifactsBulkAdditionRequest%2Did"] = driveRestoreArtifactsBulkAdditionRequestId
    }
    return NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsDriveRestoreArtifactsBulkAdditionRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderInternal instantiates a new BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder and sets the default values.
func NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder) {
    m := &BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/oneDriveForBusinessRestoreSessions/{oneDriveForBusinessRestoreSession%2Did}/driveRestoreArtifactsBulkAdditionRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder instantiates a new BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder and sets the default values.
func NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsCountRequestBuilder when successful
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder) Count()(*BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsCountRequestBuilder) {
    return NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the driveRestoreArtifactsBulkAdditionRequest objects associated with a oneDriveForBusinessRestoreSession. The drives property is deliberately omitted from the response body in order to limit the response size.
// returns a DriveRestoreArtifactsBulkAdditionRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onedriveforbusinessrestoresession-list-driverestoreartifactsbulkadditionrequests?view=graph-rest-beta
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactsBulkAdditionRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveRestoreArtifactsBulkAdditionRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactsBulkAdditionRequestCollectionResponseable), nil
}
// Post create a driveRestoreArtifactsBulkAdditionRequest object associated with a oneDriveForBusinessRestoreSession. The initial status upon creation of the restore session is active. When all the drives are added to the corresponding OneDrive restore session and the restore session is activated, the status becomes completed. If any failures are encountered during resource resolution, the status of the restore session becomes completedWithErrors.
// returns a DriveRestoreArtifactsBulkAdditionRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onedriveforbusinessrestoresession-post-driverestoreartifactsbulkadditionrequests?view=graph-rest-beta
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactsBulkAdditionRequestable, requestConfiguration *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactsBulkAdditionRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveRestoreArtifactsBulkAdditionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactsBulkAdditionRequestable), nil
}
// ToGetRequestInformation get a list of the driveRestoreArtifactsBulkAdditionRequest objects associated with a oneDriveForBusinessRestoreSession. The drives property is deliberately omitted from the response body in order to limit the response size.
// returns a *RequestInformation when successful
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a driveRestoreArtifactsBulkAdditionRequest object associated with a oneDriveForBusinessRestoreSession. The initial status upon creation of the restore session is active. When all the drives are added to the corresponding OneDrive restore session and the restore session is activated, the status becomes completed. If any failures are encountered during resource resolution, the status of the restore session becomes completedWithErrors.
// returns a *RequestInformation when successful
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactsBulkAdditionRequestable, requestConfiguration *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder when successful
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder) {
    return NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsBulkAdditionRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
