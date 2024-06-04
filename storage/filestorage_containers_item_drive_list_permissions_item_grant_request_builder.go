package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder provides operations to call the grant method.
type FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveListPermissionsItemGrantRequestBuilderInternal instantiates a new FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListPermissionsItemGrantRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder) {
    m := &FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/permissions/{permission%2Did}/grant", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder instantiates a new FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveListPermissionsItemGrantRequestBuilderInternal(urlParams, requestAdapter)
}
// Post grant users access to a link represented by a permission.
// Deprecated: This method is obsolete. Use PostAsGrantPostResponse instead.
// returns a FilestorageContainersItemDriveListPermissionsItemGrantResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permission-grant?view=graph-rest-beta
func (m *FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder) Post(ctx context.Context, body FilestorageContainersItemDriveListPermissionsItemGrantPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilderPostRequestConfiguration)(FilestorageContainersItemDriveListPermissionsItemGrantResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveListPermissionsItemGrantResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveListPermissionsItemGrantResponseable), nil
}
// PostAsGrantPostResponse grant users access to a link represented by a permission.
// returns a FilestorageContainersItemDriveListPermissionsItemGrantPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permission-grant?view=graph-rest-beta
func (m *FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder) PostAsGrantPostResponse(ctx context.Context, body FilestorageContainersItemDriveListPermissionsItemGrantPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilderPostRequestConfiguration)(FilestorageContainersItemDriveListPermissionsItemGrantPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveListPermissionsItemGrantPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveListPermissionsItemGrantPostResponseable), nil
}
// ToPostRequestInformation grant users access to a link represented by a permission.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageContainersItemDriveListPermissionsItemGrantPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder when successful
func (m *FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder) {
    return NewFilestorageContainersItemDriveListPermissionsItemGrantRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
