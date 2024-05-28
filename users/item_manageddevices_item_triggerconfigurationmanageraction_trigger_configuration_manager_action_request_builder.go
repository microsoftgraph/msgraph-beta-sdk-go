package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder provides operations to call the triggerConfigurationManagerAction method.
type ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilderInternal instantiates a new ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder and sets the default values.
func NewItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder) {
    m := &ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/triggerConfigurationManagerAction", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder instantiates a new ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder and sets the default values.
func NewItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post trigger action on ConfigurationManager client
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder) Post(ctx context.Context, body ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionPostRequestBodyable, requestConfiguration *ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation trigger action on ConfigurationManager client
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionPostRequestBodyable, requestConfiguration *ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder when successful
func (m *ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder) {
    return NewItemManageddevicesItemTriggerconfigurationmanageractionTriggerConfigurationManagerActionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
