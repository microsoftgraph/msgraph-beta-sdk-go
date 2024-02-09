package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemConnectorsItemMemberOfItemRefRequestBuilder provides operations to manage the collection of onPremisesPublishingProfile entities.
type ItemConnectorsItemMemberOfItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConnectorsItemMemberOfItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConnectorsItemMemberOfItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemConnectorsItemMemberOfItemRefRequestBuilderInternal instantiates a new ItemConnectorsItemMemberOfItemRefRequestBuilder and sets the default values.
func NewItemConnectorsItemMemberOfItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorsItemMemberOfItemRefRequestBuilder) {
    m := &ItemConnectorsItemMemberOfItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectors/{connector%2Did}/memberOf/{connectorGroup%2Did}/$ref", pathParameters),
    }
    return m
}
// NewItemConnectorsItemMemberOfItemRefRequestBuilder instantiates a new ItemConnectorsItemMemberOfItemRefRequestBuilder and sets the default values.
func NewItemConnectorsItemMemberOfItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConnectorsItemMemberOfItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConnectorsItemMemberOfItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete ref of navigation property memberOf for onPremisesPublishingProfiles
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemConnectorsItemMemberOfItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemConnectorsItemMemberOfItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation delete ref of navigation property memberOf for onPremisesPublishingProfiles
// returns a *RequestInformation when successful
func (m *ItemConnectorsItemMemberOfItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemConnectorsItemMemberOfItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemConnectorsItemMemberOfItemRefRequestBuilder when successful
func (m *ItemConnectorsItemMemberOfItemRefRequestBuilder) WithUrl(rawUrl string)(*ItemConnectorsItemMemberOfItemRefRequestBuilder) {
    return NewItemConnectorsItemMemberOfItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
