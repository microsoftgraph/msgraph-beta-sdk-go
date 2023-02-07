package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilder provides operations to call the wipeManagedAppRegistrationByDeviceTag method.
type ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal instantiates a new MicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilder and sets the default values.
func NewItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    m := &ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/microsoft.graph.wipeManagedAppRegistrationByDeviceTag";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilder instantiates a new MicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilder and sets the default values.
func NewItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(urlParams, requestAdapter)
}
// Post issues a wipe operation on an app registration with specified device tag.
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilder) Post(ctx context.Context, body ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagWipeManagedAppRegistrationByDeviceTagPostRequestBodyable, requestConfiguration *ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation issues a wipe operation on an app registration with specified device tag.
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagWipeManagedAppRegistrationByDeviceTagPostRequestBodyable, requestConfiguration *ItemMicrosoftGraphWipeManagedAppRegistrationByDeviceTagRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
