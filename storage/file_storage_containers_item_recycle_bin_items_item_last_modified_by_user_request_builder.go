package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
type FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilderGetQueryParameters get lastModifiedByUser from storage
type FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilderGetQueryParameters
}
// NewFileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilderInternal instantiates a new FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder and sets the default values.
func NewFileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder) {
    m := &FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/recycleBin/items/{recycleBinItem%2Did}/lastModifiedByUser{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder instantiates a new FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder and sets the default values.
func NewFileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get lastModifiedByUser from storage
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// MailboxSettings the mailboxSettings property
// returns a *FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserMailboxSettingsRequestBuilder when successful
func (m *FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder) MailboxSettings()(*FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserMailboxSettingsRequestBuilder) {
    return NewFileStorageContainersItemRecycleBinItemsItemLastModifiedByUserMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserServiceProvisioningErrorsRequestBuilder when successful
func (m *FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder) ServiceProvisioningErrors()(*FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserServiceProvisioningErrorsRequestBuilder) {
    return NewFileStorageContainersItemRecycleBinItemsItemLastModifiedByUserServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get lastModifiedByUser from storage
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder when successful
func (m *FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder) {
    return NewFileStorageContainersItemRecycleBinItemsItemLastModifiedByUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
