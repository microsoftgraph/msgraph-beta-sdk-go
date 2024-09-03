package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOutlookRequestBuilder provides operations to manage the outlook property of the microsoft.graph.user entity.
type ItemOutlookRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOutlookRequestBuilderGetQueryParameters selective Outlook services available to the user. Read-only. Nullable.
type ItemOutlookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOutlookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOutlookRequestBuilderGetQueryParameters
}
// NewItemOutlookRequestBuilderInternal instantiates a new ItemOutlookRequestBuilder and sets the default values.
func NewItemOutlookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookRequestBuilder) {
    m := &ItemOutlookRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/outlook{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemOutlookRequestBuilder instantiates a new ItemOutlookRequestBuilder and sets the default values.
func NewItemOutlookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutlookRequestBuilderInternal(urlParams, requestAdapter)
}
// Get selective Outlook services available to the user. Read-only. Nullable.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a OutlookUserable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOutlookRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOutlookRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutlookUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutlookUserable), nil
}
// MasterCategories provides operations to manage the masterCategories property of the microsoft.graph.outlookUser entity.
// returns a *ItemOutlookMasterCategoriesRequestBuilder when successful
func (m *ItemOutlookRequestBuilder) MasterCategories()(*ItemOutlookMasterCategoriesRequestBuilder) {
    return NewItemOutlookMasterCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SupportedLanguages provides operations to call the supportedLanguages method.
// returns a *ItemOutlookSupportedLanguagesRequestBuilder when successful
func (m *ItemOutlookRequestBuilder) SupportedLanguages()(*ItemOutlookSupportedLanguagesRequestBuilder) {
    return NewItemOutlookSupportedLanguagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SupportedTimeZones provides operations to call the supportedTimeZones method.
// returns a *ItemOutlookSupportedTimeZonesRequestBuilder when successful
func (m *ItemOutlookRequestBuilder) SupportedTimeZones()(*ItemOutlookSupportedTimeZonesRequestBuilder) {
    return NewItemOutlookSupportedTimeZonesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SupportedTimeZonesWithTimeZoneStandard provides operations to call the supportedTimeZones method.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemOutlookSupportedTimeZonesWithTimeZoneStandardRequestBuilder when successful
func (m *ItemOutlookRequestBuilder) SupportedTimeZonesWithTimeZoneStandard(timeZoneStandard *string)(*ItemOutlookSupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    return NewItemOutlookSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, timeZoneStandard)
}
// TaskFolders provides operations to manage the taskFolders property of the microsoft.graph.outlookUser entity.
// returns a *ItemOutlookTaskFoldersRequestBuilder when successful
func (m *ItemOutlookRequestBuilder) TaskFolders()(*ItemOutlookTaskFoldersRequestBuilder) {
    return NewItemOutlookTaskFoldersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaskGroups provides operations to manage the taskGroups property of the microsoft.graph.outlookUser entity.
// returns a *ItemOutlookTaskGroupsRequestBuilder when successful
func (m *ItemOutlookRequestBuilder) TaskGroups()(*ItemOutlookTaskGroupsRequestBuilder) {
    return NewItemOutlookTaskGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.outlookUser entity.
// returns a *ItemOutlookTasksRequestBuilder when successful
func (m *ItemOutlookRequestBuilder) Tasks()(*ItemOutlookTasksRequestBuilder) {
    return NewItemOutlookTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation selective Outlook services available to the user. Read-only. Nullable.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemOutlookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOutlookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemOutlookRequestBuilder when successful
func (m *ItemOutlookRequestBuilder) WithUrl(rawUrl string)(*ItemOutlookRequestBuilder) {
    return NewItemOutlookRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
