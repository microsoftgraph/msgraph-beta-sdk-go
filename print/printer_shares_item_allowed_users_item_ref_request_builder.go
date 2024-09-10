package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrinterSharesItemAllowedUsersItemRefRequestBuilder provides operations to manage the collection of print entities.
type PrinterSharesItemAllowedUsersItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrinterSharesItemAllowedUsersItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrinterSharesItemAllowedUsersItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrinterSharesItemAllowedUsersItemRefRequestBuilderInternal instantiates a new PrinterSharesItemAllowedUsersItemRefRequestBuilder and sets the default values.
func NewPrinterSharesItemAllowedUsersItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemAllowedUsersItemRefRequestBuilder) {
    m := &PrinterSharesItemAllowedUsersItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printerShares/{printerShare%2Did}/allowedUsers/{user%2Did}/$ref", pathParameters),
    }
    return m
}
// NewPrinterSharesItemAllowedUsersItemRefRequestBuilder instantiates a new PrinterSharesItemAllowedUsersItemRefRequestBuilder and sets the default values.
func NewPrinterSharesItemAllowedUsersItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrinterSharesItemAllowedUsersItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterSharesItemAllowedUsersItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete revoke the specified user's access to submit print jobs to the associated printerShare.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/printershare-delete-alloweduser?view=graph-rest-beta
func (m *PrinterSharesItemAllowedUsersItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrinterSharesItemAllowedUsersItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation revoke the specified user's access to submit print jobs to the associated printerShare.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *PrinterSharesItemAllowedUsersItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrinterSharesItemAllowedUsersItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *PrinterSharesItemAllowedUsersItemRefRequestBuilder when successful
func (m *PrinterSharesItemAllowedUsersItemRefRequestBuilder) WithUrl(rawUrl string)(*PrinterSharesItemAllowedUsersItemRefRequestBuilder) {
    return NewPrinterSharesItemAllowedUsersItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
