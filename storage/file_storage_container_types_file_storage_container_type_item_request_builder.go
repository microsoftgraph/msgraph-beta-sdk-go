// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder provides operations to manage the containerTypes property of the microsoft.graph.fileStorage entity.
type FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderGetQueryParameters get containerTypes from storage
type FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderGetQueryParameters
}
// FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderInternal instantiates a new FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder and sets the default values.
func NewFileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder) {
    m := &FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containerTypes/{fileStorageContainerType%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder instantiates a new FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder and sets the default values.
func NewFileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property containerTypes for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get containerTypes from storage
// returns a FileStorageContainerTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FileStorageContainerTypeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateFileStorageContainerTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FileStorageContainerTypeable), nil
}
// Patch update the navigation property containerTypes in storage
// returns a FileStorageContainerTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FileStorageContainerTypeable, requestConfiguration *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FileStorageContainerTypeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateFileStorageContainerTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FileStorageContainerTypeable), nil
}
// ToDeleteRequestInformation delete navigation property containerTypes for storage
// returns a *RequestInformation when successful
func (m *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get containerTypes from storage
// returns a *RequestInformation when successful
func (m *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property containerTypes in storage
// returns a *RequestInformation when successful
func (m *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FileStorageContainerTypeable, requestConfiguration *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder when successful
func (m *FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder) {
    return NewFileStorageContainerTypesFileStorageContainerTypeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
