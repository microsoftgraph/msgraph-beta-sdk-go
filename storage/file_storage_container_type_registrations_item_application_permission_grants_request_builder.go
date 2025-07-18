// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder provides operations to manage the applicationPermissionGrants property of the microsoft.graph.fileStorageContainerTypeRegistration entity.
type FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderGetQueryParameters get applicationPermissionGrants from storage
type FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderGetQueryParameters struct {
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
// FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderGetQueryParameters
}
// FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByFileStorageContainerTypeAppPermissionGrantAppId provides operations to manage the applicationPermissionGrants property of the microsoft.graph.fileStorageContainerTypeRegistration entity.
// returns a *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsFileStorageContainerTypeAppPermissionGrantAppItemRequestBuilder when successful
func (m *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder) ByFileStorageContainerTypeAppPermissionGrantAppId(fileStorageContainerTypeAppPermissionGrantAppId string)(*FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsFileStorageContainerTypeAppPermissionGrantAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if fileStorageContainerTypeAppPermissionGrantAppId != "" {
        urlTplParams["fileStorageContainerTypeAppPermissionGrant%2DappId"] = fileStorageContainerTypeAppPermissionGrantAppId
    }
    return NewFileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsFileStorageContainerTypeAppPermissionGrantAppItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderInternal instantiates a new FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder and sets the default values.
func NewFileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder) {
    m := &FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containerTypeRegistrations/{fileStorageContainerTypeRegistration%2Did}/applicationPermissionGrants{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder instantiates a new FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder and sets the default values.
func NewFileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsCountRequestBuilder when successful
func (m *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder) Count()(*FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsCountRequestBuilder) {
    return NewFileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get applicationPermissionGrants from storage
// returns a FileStorageContainerTypeAppPermissionGrantCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FileStorageContainerTypeAppPermissionGrantCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateFileStorageContainerTypeAppPermissionGrantCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FileStorageContainerTypeAppPermissionGrantCollectionResponseable), nil
}
// Post create new navigation property to applicationPermissionGrants for storage
// returns a FileStorageContainerTypeAppPermissionGrantable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FileStorageContainerTypeAppPermissionGrantable, requestConfiguration *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FileStorageContainerTypeAppPermissionGrantable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateFileStorageContainerTypeAppPermissionGrantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FileStorageContainerTypeAppPermissionGrantable), nil
}
// ToGetRequestInformation get applicationPermissionGrants from storage
// returns a *RequestInformation when successful
func (m *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to applicationPermissionGrants for storage
// returns a *RequestInformation when successful
func (m *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FileStorageContainerTypeAppPermissionGrantable, requestConfiguration *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder when successful
func (m *FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder) {
    return NewFileStorageContainerTypeRegistrationsItemApplicationPermissionGrantsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
