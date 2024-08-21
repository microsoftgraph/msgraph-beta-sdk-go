package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPhotosProfilePhotoItemRequestBuilder provides operations to manage the photos property of the microsoft.graph.user entity.
type ItemPhotosProfilePhotoItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPhotosProfilePhotoItemRequestBuilderGetQueryParameters the collection of the user's profile photos in different sizes. Read-only.
type ItemPhotosProfilePhotoItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPhotosProfilePhotoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPhotosProfilePhotoItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPhotosProfilePhotoItemRequestBuilderGetQueryParameters
}
// NewItemPhotosProfilePhotoItemRequestBuilderInternal instantiates a new ItemPhotosProfilePhotoItemRequestBuilder and sets the default values.
func NewItemPhotosProfilePhotoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPhotosProfilePhotoItemRequestBuilder) {
    m := &ItemPhotosProfilePhotoItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/photos/{profilePhoto%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPhotosProfilePhotoItemRequestBuilder instantiates a new ItemPhotosProfilePhotoItemRequestBuilder and sets the default values.
func NewItemPhotosProfilePhotoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPhotosProfilePhotoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPhotosProfilePhotoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
// returns a *ItemPhotosItemValueContentRequestBuilder when successful
func (m *ItemPhotosProfilePhotoItemRequestBuilder) Content()(*ItemPhotosItemValueContentRequestBuilder) {
    return NewItemPhotosItemValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of the user's profile photos in different sizes. Read-only.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ProfilePhotoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPhotosProfilePhotoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPhotosProfilePhotoItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProfilePhotoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProfilePhotoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProfilePhotoable), nil
}
// ToGetRequestInformation the collection of the user's profile photos in different sizes. Read-only.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemPhotosProfilePhotoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPhotosProfilePhotoItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPhotosProfilePhotoItemRequestBuilder when successful
func (m *ItemPhotosProfilePhotoItemRequestBuilder) WithUrl(rawUrl string)(*ItemPhotosProfilePhotoItemRequestBuilder) {
    return NewItemPhotosProfilePhotoItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
