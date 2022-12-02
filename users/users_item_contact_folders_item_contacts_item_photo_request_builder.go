package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemContactFoldersItemContactsItemPhotoRequestBuilder provides operations to manage the photo property of the microsoft.graph.contact entity.
type UsersItemContactFoldersItemContactsItemPhotoRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemContactFoldersItemContactsItemPhotoRequestBuilderGetQueryParameters optional contact picture. You can get or set a photo for a contact.
type UsersItemContactFoldersItemContactsItemPhotoRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemContactFoldersItemContactsItemPhotoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemContactFoldersItemContactsItemPhotoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemContactFoldersItemContactsItemPhotoRequestBuilderGetQueryParameters
}
// UsersItemContactFoldersItemContactsItemPhotoRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemContactFoldersItemContactsItemPhotoRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemContactFoldersItemContactsItemPhotoRequestBuilderInternal instantiates a new PhotoRequestBuilder and sets the default values.
func NewUsersItemContactFoldersItemContactsItemPhotoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemContactFoldersItemContactsItemPhotoRequestBuilder) {
    m := &UsersItemContactFoldersItemContactsItemPhotoRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/contactFolders/{contactFolder%2Did}/contacts/{contact%2Did}/photo{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemContactFoldersItemContactsItemPhotoRequestBuilder instantiates a new PhotoRequestBuilder and sets the default values.
func NewUsersItemContactFoldersItemContactsItemPhotoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemContactFoldersItemContactsItemPhotoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemContactFoldersItemContactsItemPhotoRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *UsersItemContactFoldersItemContactsItemPhotoRequestBuilder) Content()(*UsersItemContactFoldersItemContactsItemPhotoValueContentRequestBuilder) {
    return NewUsersItemContactFoldersItemContactsItemPhotoValueContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation optional contact picture. You can get or set a photo for a contact.
func (m *UsersItemContactFoldersItemContactsItemPhotoRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemContactFoldersItemContactsItemPhotoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property photo in users
func (m *UsersItemContactFoldersItemContactsItemPhotoRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProfilePhotoable, requestConfiguration *UsersItemContactFoldersItemContactsItemPhotoRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get optional contact picture. You can get or set a photo for a contact.
func (m *UsersItemContactFoldersItemContactsItemPhotoRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemContactFoldersItemContactsItemPhotoRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProfilePhotoable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProfilePhotoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProfilePhotoable), nil
}
// Patch update the navigation property photo in users
func (m *UsersItemContactFoldersItemContactsItemPhotoRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProfilePhotoable, requestConfiguration *UsersItemContactFoldersItemContactsItemPhotoRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProfilePhotoable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProfilePhotoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProfilePhotoable), nil
}
