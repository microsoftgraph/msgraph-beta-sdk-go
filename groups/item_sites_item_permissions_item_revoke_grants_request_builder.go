package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemPermissionsItemRevokeGrantsRequestBuilder provides operations to call the revokeGrants method.
type ItemSitesItemPermissionsItemRevokeGrantsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPermissionsItemRevokeGrantsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPermissionsItemRevokeGrantsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemPermissionsItemRevokeGrantsRequestBuilderInternal instantiates a new RevokeGrantsRequestBuilder and sets the default values.
func NewItemSitesItemPermissionsItemRevokeGrantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPermissionsItemRevokeGrantsRequestBuilder) {
    m := &ItemSitesItemPermissionsItemRevokeGrantsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/permissions/{permission%2Did}/revokeGrants", pathParameters),
    }
    return m
}
// NewItemSitesItemPermissionsItemRevokeGrantsRequestBuilder instantiates a new RevokeGrantsRequestBuilder and sets the default values.
func NewItemSitesItemPermissionsItemRevokeGrantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPermissionsItemRevokeGrantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPermissionsItemRevokeGrantsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post revoke access to a [listItem][] or [driveItem][] granted via a sharing link by removing the specified [recipient][] from the link.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permission-revokegrants?view=graph-rest-1.0
func (m *ItemSitesItemPermissionsItemRevokeGrantsRequestBuilder) Post(ctx context.Context, body ItemSitesItemPermissionsItemRevokeGrantsPostRequestBodyable, requestConfiguration *ItemSitesItemPermissionsItemRevokeGrantsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Permissionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Permissionable), nil
}
// ToPostRequestInformation revoke access to a [listItem][] or [driveItem][] granted via a sharing link by removing the specified [recipient][] from the link.
func (m *ItemSitesItemPermissionsItemRevokeGrantsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemPermissionsItemRevokeGrantsPostRequestBodyable, requestConfiguration *ItemSitesItemPermissionsItemRevokeGrantsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSitesItemPermissionsItemRevokeGrantsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPermissionsItemRevokeGrantsRequestBuilder) {
    return NewItemSitesItemPermissionsItemRevokeGrantsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
