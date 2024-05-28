package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder provides operations to call the requestUpgrade method.
type MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilderInternal instantiates a new MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder and sets the default values.
func NewMicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder) {
    m := &MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}/requestUpgrade", pathParameters),
    }
    return m
}
// NewMicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder instantiates a new MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder and sets the default values.
func NewMicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action requestUpgrade
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder) Post(ctx context.Context, requestConfiguration *MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action requestUpgrade
// returns a *RequestInformation when successful
func (m *MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder when successful
func (m *MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder) WithUrl(rawUrl string)(*MicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder) {
    return NewMicrosofttunnelsitesItemRequestupgradeRequestUpgradeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
