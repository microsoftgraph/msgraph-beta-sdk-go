package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder provides operations to manage the workbook property of the microsoft.graph.driveItem entity.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderGetQueryParameters for files that are Excel spreadsheets, access to the workbook API to work with the spreadsheet's contents. Nullable.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderGetQueryParameters
}
// FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Application provides operations to manage the application property of the microsoft.graph.workbook entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookApplicationRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) Application()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookApplicationRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloseSession provides operations to call the closeSession method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookCloseSessionRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) CloseSession()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookCloseSessionRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookCloseSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Comments provides operations to manage the comments property of the microsoft.graph.workbook entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) Comments()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookCommentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateSession provides operations to call the createSession method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookCreateSessionRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) CreateSession()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookCreateSessionRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookCreateSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property workbook for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderDeleteRequestConfiguration)(error) {
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
// Functions provides operations to manage the functions property of the microsoft.graph.workbook entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) Functions()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookFunctionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get for files that are Excel spreadsheets, access to the workbook API to work with the spreadsheet's contents. Nullable.
// returns a Workbookable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Workbookable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Workbookable), nil
}
// Names provides operations to manage the names property of the microsoft.graph.workbook entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookNamesRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) Names()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookNamesRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookNamesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.workbook entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookOperationsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) Operations()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookOperationsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property workbook in storage
// returns a Workbookable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Workbookable, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Workbookable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Workbookable), nil
}
// RefreshSession provides operations to call the refreshSession method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookRefreshSessionRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) RefreshSession()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookRefreshSessionRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookRefreshSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SessionInfoResourceWithKey provides operations to call the sessionInfoResource method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookSessionInfoResourceWithKeyRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) SessionInfoResourceWithKey(key *string)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookSessionInfoResourceWithKeyRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookSessionInfoResourceWithKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, key)
}
// TableRowOperationResultWithKey provides operations to call the tableRowOperationResult method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) TableRowOperationResultWithKey(key *string)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTableRowOperationResultWithKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, key)
}
// Tables provides operations to manage the tables property of the microsoft.graph.workbook entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) Tables()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookTablesRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookTablesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property workbook for storage
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation for files that are Excel spreadsheets, access to the workbook API to work with the spreadsheet's contents. Nullable.
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property workbook in storage
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Workbookable, requestConfiguration *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheets provides operations to manage the worksheets property of the microsoft.graph.workbook entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) Worksheets()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookWorksheetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
