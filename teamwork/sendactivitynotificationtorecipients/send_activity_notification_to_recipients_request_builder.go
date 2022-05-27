package sendactivitynotificationtorecipients

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SendActivityNotificationToRecipientsRequestBuilder provides operations to call the sendActivityNotificationToRecipients method.
type SendActivityNotificationToRecipientsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SendActivityNotificationToRecipientsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SendActivityNotificationToRecipientsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSendActivityNotificationToRecipientsRequestBuilderInternal instantiates a new SendActivityNotificationToRecipientsRequestBuilder and sets the default values.
func NewSendActivityNotificationToRecipientsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SendActivityNotificationToRecipientsRequestBuilder) {
    m := &SendActivityNotificationToRecipientsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/microsoft.graph.sendActivityNotificationToRecipients";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSendActivityNotificationToRecipientsRequestBuilder instantiates a new SendActivityNotificationToRecipientsRequestBuilder and sets the default values.
func NewSendActivityNotificationToRecipientsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SendActivityNotificationToRecipientsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSendActivityNotificationToRecipientsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action sendActivityNotificationToRecipients
func (m *SendActivityNotificationToRecipientsRequestBuilder) CreatePostRequestInformation(body SendActivityNotificationToRecipientsPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action sendActivityNotificationToRecipients
func (m *SendActivityNotificationToRecipientsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body SendActivityNotificationToRecipientsPostRequestBodyable, requestConfiguration *SendActivityNotificationToRecipientsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action sendActivityNotificationToRecipients
func (m *SendActivityNotificationToRecipientsRequestBuilder) Post(body SendActivityNotificationToRecipientsPostRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action sendActivityNotificationToRecipients
func (m *SendActivityNotificationToRecipientsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body SendActivityNotificationToRecipientsPostRequestBodyable, requestConfiguration *SendActivityNotificationToRecipientsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
