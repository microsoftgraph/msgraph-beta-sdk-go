package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder provides operations to manage the driveRestoreArtifacts property of the microsoft.graph.oneDriveForBusinessRestoreSession entity.
type BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderGetQueryParameters get a list of the driveRestoreArtifact objects and their properties for a oneDriveForBusinessRestoreSession for a tenant.
type BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderGetQueryParameters struct {
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
// BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderGetQueryParameters
}
// BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDriveRestoreArtifactId provides operations to manage the driveRestoreArtifacts property of the microsoft.graph.oneDriveForBusinessRestoreSession entity.
// returns a *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactItemRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder) ByDriveRestoreArtifactId(driveRestoreArtifactId string)(*BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if driveRestoreArtifactId != "" {
        urlTplParams["driveRestoreArtifact%2Did"] = driveRestoreArtifactId
    }
    return NewBackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderInternal instantiates a new BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder and sets the default values.
func NewBackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder) {
    m := &BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/oneDriveForBusinessRestoreSessions/{oneDriveForBusinessRestoreSession%2Did}/driveRestoreArtifacts{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder instantiates a new BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder and sets the default values.
func NewBackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsCountRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder) Count()(*BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsCountRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the driveRestoreArtifact objects and their properties for a oneDriveForBusinessRestoreSession for a tenant.
// returns a DriveRestoreArtifactCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onedriveforbusinessrestoresession-list-driverestoreartifacts?view=graph-rest-beta
func (m *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveRestoreArtifactCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactCollectionResponseable), nil
}
// Post create new navigation property to driveRestoreArtifacts for solutions
// returns a DriveRestoreArtifactable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactable, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveRestoreArtifactFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactable), nil
}
// ToGetRequestInformation get a list of the driveRestoreArtifact objects and their properties for a oneDriveForBusinessRestoreSession for a tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to driveRestoreArtifacts for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactable, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
