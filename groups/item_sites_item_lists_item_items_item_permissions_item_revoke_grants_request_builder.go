package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder provides operations to call the revokeGrants method.
type ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilderInternal instantiates a new ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder) {
    m := &ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/permissions/{permission%2Did}/revokeGrants", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder instantiates a new ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post revoke access to a listItem or driveItem granted via a sharing link by removing the specified recipient from the link.
// returns a Permissionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permission-revokegrants?view=graph-rest-beta
func (m *ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder) Post(ctx context.Context, body ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsPostRequestBodyable, requestConfiguration *ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Permissionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToPostRequestInformation revoke access to a listItem or driveItem granted via a sharing link by removing the specified recipient from the link.
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsPostRequestBodyable, requestConfiguration *ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder when successful
func (m *ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder) {
    return NewItemSitesItemListsItemItemsItemPermissionsItemRevokeGrantsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
