package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder provides operations to manage the oneDriveForBusinessRestoreSessions property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderGetQueryParameters the list of OneDrive for Business restore sessions available in the tenant.
type BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderGetQueryParameters
}
// BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderInternal instantiates a new BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder and sets the default values.
func NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder) {
    m := &BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/oneDriveForBusinessRestoreSessions/{oneDriveForBusinessRestoreSession%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder instantiates a new BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder and sets the default values.
func NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property oneDriveForBusinessRestoreSessions for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DriveRestoreArtifacts provides operations to manage the driveRestoreArtifacts property of the microsoft.graph.oneDriveForBusinessRestoreSession entity.
// returns a *BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder) DriveRestoreArtifacts()(*BackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessrestoresessionsItemDriverestoreartifactsDriveRestoreArtifactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of OneDrive for Business restore sessions available in the tenant.
// returns a OneDriveForBusinessRestoreSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessRestoreSessionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the properties of a oneDriveForBusinessRestoreSession object.
// returns a OneDriveForBusinessRestoreSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onedriveforbusinessrestoresession-update?view=graph-rest-beta
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessRestoreSessionable, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessRestoreSessionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property oneDriveForBusinessRestoreSessions for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of OneDrive for Business restore sessions available in the tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a oneDriveForBusinessRestoreSession object.
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessRestoreSessionable, requestConfiguration *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
