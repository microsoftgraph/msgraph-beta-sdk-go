package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VpptokensItemRevokelicensesRevokeLicensesRequestBuilder provides operations to call the revokeLicenses method.
type VpptokensItemRevokelicensesRevokeLicensesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VpptokensItemRevokelicensesRevokeLicensesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VpptokensItemRevokelicensesRevokeLicensesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVpptokensItemRevokelicensesRevokeLicensesRequestBuilderInternal instantiates a new VpptokensItemRevokelicensesRevokeLicensesRequestBuilder and sets the default values.
func NewVpptokensItemRevokelicensesRevokeLicensesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VpptokensItemRevokelicensesRevokeLicensesRequestBuilder) {
    m := &VpptokensItemRevokelicensesRevokeLicensesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/vppTokens/{vppToken%2Did}/revokeLicenses", pathParameters),
    }
    return m
}
// NewVpptokensItemRevokelicensesRevokeLicensesRequestBuilder instantiates a new VpptokensItemRevokelicensesRevokeLicensesRequestBuilder and sets the default values.
func NewVpptokensItemRevokelicensesRevokeLicensesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VpptokensItemRevokelicensesRevokeLicensesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVpptokensItemRevokelicensesRevokeLicensesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post revoke licenses associated with a specific appleVolumePurchaseProgramToken
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VpptokensItemRevokelicensesRevokeLicensesRequestBuilder) Post(ctx context.Context, body VpptokensItemRevokelicensesRevokeLicensesPostRequestBodyable, requestConfiguration *VpptokensItemRevokelicensesRevokeLicensesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation revoke licenses associated with a specific appleVolumePurchaseProgramToken
// returns a *RequestInformation when successful
func (m *VpptokensItemRevokelicensesRevokeLicensesRequestBuilder) ToPostRequestInformation(ctx context.Context, body VpptokensItemRevokelicensesRevokeLicensesPostRequestBodyable, requestConfiguration *VpptokensItemRevokelicensesRevokeLicensesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VpptokensItemRevokelicensesRevokeLicensesRequestBuilder when successful
func (m *VpptokensItemRevokelicensesRevokeLicensesRequestBuilder) WithUrl(rawUrl string)(*VpptokensItemRevokelicensesRevokeLicensesRequestBuilder) {
    return NewVpptokensItemRevokelicensesRevokeLicensesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
