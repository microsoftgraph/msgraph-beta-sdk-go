package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder provides operations to manage the photo property of the microsoft.graph.contact entity.
type ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderGetQueryParameters optional contact picture. You can get or set a photo for a contact.
type ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderGetQueryParameters
}
// ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderInternal instantiates a new ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder and sets the default values.
func NewItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder) {
    m := &ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/contactFolders/{contactFolder%2Did}/childFolders/{contactFolder%2Did1}/contacts/{contact%2Did}/photo{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder instantiates a new ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder and sets the default values.
func NewItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
// returns a *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder when successful
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder) Content()(*ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder) {
    return NewItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get optional contact picture. You can get or set a photo for a contact.
// returns a ProfilePhotoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProfilePhotoable, error) {
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
// Patch update the navigation property photo in users
// returns a ProfilePhotoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProfilePhotoable, requestConfiguration *ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProfilePhotoable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation optional contact picture. You can get or set a photo for a contact.
// returns a *RequestInformation when successful
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property photo in users
// returns a *RequestInformation when successful
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProfilePhotoable, requestConfiguration *ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder when successful
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder) WithUrl(rawUrl string)(*ItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder) {
    return NewItemContactfoldersItemChildfoldersItemContactsItemPhotoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
