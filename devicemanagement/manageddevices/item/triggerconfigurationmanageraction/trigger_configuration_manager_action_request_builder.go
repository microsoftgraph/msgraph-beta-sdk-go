package triggerconfigurationmanageraction

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TriggerConfigurationManagerActionRequestBuilder provides operations to call the triggerConfigurationManagerAction method.
type TriggerConfigurationManagerActionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TriggerConfigurationManagerActionRequestBuilderPostOptions options for Post
type TriggerConfigurationManagerActionRequestBuilderPostOptions struct {
    // 
    Body TriggerConfigurationManagerActionRequestBodyable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewTriggerConfigurationManagerActionRequestBuilderInternal instantiates a new TriggerConfigurationManagerActionRequestBuilder and sets the default values.
func NewTriggerConfigurationManagerActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggerConfigurationManagerActionRequestBuilder) {
    m := &TriggerConfigurationManagerActionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/microsoft.graph.triggerConfigurationManagerAction";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTriggerConfigurationManagerActionRequestBuilder instantiates a new TriggerConfigurationManagerActionRequestBuilder and sets the default values.
func NewTriggerConfigurationManagerActionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggerConfigurationManagerActionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTriggerConfigurationManagerActionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation trigger action on ConfigurationManager client
func (m *TriggerConfigurationManagerActionRequestBuilder) CreatePostRequestInformation(options *TriggerConfigurationManagerActionRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post trigger action on ConfigurationManager client
func (m *TriggerConfigurationManagerActionRequestBuilder) Post(options *TriggerConfigurationManagerActionRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
