package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemDrivesItemListContentTypesItemPublishRequestBuilder provides operations to call the publish method.
type UsersItemDrivesItemListContentTypesItemPublishRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemDrivesItemListContentTypesItemPublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDrivesItemListContentTypesItemPublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemDrivesItemListContentTypesItemPublishRequestBuilderInternal instantiates a new PublishRequestBuilder and sets the default values.
func NewUsersItemDrivesItemListContentTypesItemPublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDrivesItemListContentTypesItemPublishRequestBuilder) {
    m := &UsersItemDrivesItemListContentTypesItemPublishRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/list/contentTypes/{contentType%2Did}/microsoft.graph.publish";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemDrivesItemListContentTypesItemPublishRequestBuilder instantiates a new PublishRequestBuilder and sets the default values.
func NewUsersItemDrivesItemListContentTypesItemPublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDrivesItemListContentTypesItemPublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemDrivesItemListContentTypesItemPublishRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation publishes a [contentType][] present in a content type hub site.
func (m *UsersItemDrivesItemListContentTypesItemPublishRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *UsersItemDrivesItemListContentTypesItemPublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post publishes a [contentType][] present in a content type hub site.
func (m *UsersItemDrivesItemListContentTypesItemPublishRequestBuilder) Post(ctx context.Context, requestConfiguration *UsersItemDrivesItemListContentTypesItemPublishRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
