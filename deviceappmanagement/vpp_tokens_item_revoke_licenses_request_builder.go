package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VppTokensItemRevokeLicensesRequestBuilder provides operations to call the revokeLicenses method.
type VppTokensItemRevokeLicensesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VppTokensItemRevokeLicensesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VppTokensItemRevokeLicensesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVppTokensItemRevokeLicensesRequestBuilderInternal instantiates a new RevokeLicensesRequestBuilder and sets the default values.
func NewVppTokensItemRevokeLicensesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VppTokensItemRevokeLicensesRequestBuilder) {
    m := &VppTokensItemRevokeLicensesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/vppTokens/{vppToken%2Did}/revokeLicenses", pathParameters),
    }
    return m
}
// NewVppTokensItemRevokeLicensesRequestBuilder instantiates a new RevokeLicensesRequestBuilder and sets the default values.
func NewVppTokensItemRevokeLicensesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VppTokensItemRevokeLicensesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVppTokensItemRevokeLicensesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post revoke licenses associated with a specific appleVolumePurchaseProgramToken
func (m *VppTokensItemRevokeLicensesRequestBuilder) Post(ctx context.Context, body VppTokensItemRevokeLicensesPostRequestBodyable, requestConfiguration *VppTokensItemRevokeLicensesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation revoke licenses associated with a specific appleVolumePurchaseProgramToken
func (m *VppTokensItemRevokeLicensesRequestBuilder) ToPostRequestInformation(ctx context.Context, body VppTokensItemRevokeLicensesPostRequestBodyable, requestConfiguration *VppTokensItemRevokeLicensesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VppTokensItemRevokeLicensesRequestBuilder) WithUrl(rawUrl string)(*VppTokensItemRevokeLicensesRequestBuilder) {
    return NewVppTokensItemRevokeLicensesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
