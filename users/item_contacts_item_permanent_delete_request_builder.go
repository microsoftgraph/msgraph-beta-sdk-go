// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemContactsItemPermanentDeleteRequestBuilder provides operations to call the permanentDelete method.
type ItemContactsItemPermanentDeleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemContactsItemPermanentDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactsItemPermanentDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemContactsItemPermanentDeleteRequestBuilderInternal instantiates a new ItemContactsItemPermanentDeleteRequestBuilder and sets the default values.
func NewItemContactsItemPermanentDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactsItemPermanentDeleteRequestBuilder) {
    m := &ItemContactsItemPermanentDeleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/contacts/{contact%2Did}/permanentDelete", pathParameters),
    }
    return m
}
// NewItemContactsItemPermanentDeleteRequestBuilder instantiates a new ItemContactsItemPermanentDeleteRequestBuilder and sets the default values.
func NewItemContactsItemPermanentDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactsItemPermanentDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContactsItemPermanentDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post permanently delete a contact and place it in the Purges folder in the dumpster in the user's mailbox. Email clients such as Outlook or Outlook on the web can't access permanently deleted items. Unless there's a hold set on the mailbox, the items are permanently deleted after a set period of time. For more information about item retention, see Configure Deleted Item retention and Recoverable Items quotas.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contact-permanentdelete?view=graph-rest-beta
func (m *ItemContactsItemPermanentDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemContactsItemPermanentDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation permanently delete a contact and place it in the Purges folder in the dumpster in the user's mailbox. Email clients such as Outlook or Outlook on the web can't access permanently deleted items. Unless there's a hold set on the mailbox, the items are permanently deleted after a set period of time. For more information about item retention, see Configure Deleted Item retention and Recoverable Items quotas.
// returns a *RequestInformation when successful
func (m *ItemContactsItemPermanentDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemContactsItemPermanentDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemContactsItemPermanentDeleteRequestBuilder when successful
func (m *ItemContactsItemPermanentDeleteRequestBuilder) WithUrl(rawUrl string)(*ItemContactsItemPermanentDeleteRequestBuilder) {
    return NewItemContactsItemPermanentDeleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
