package scheduleactionsforrules

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ScheduleActionsForRulesRequestBuilder provides operations to call the scheduleActionsForRules method.
type ScheduleActionsForRulesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ScheduleActionsForRulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScheduleActionsForRulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewScheduleActionsForRulesRequestBuilderInternal instantiates a new ScheduleActionsForRulesRequestBuilder and sets the default values.
func NewScheduleActionsForRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScheduleActionsForRulesRequestBuilder) {
    m := &ScheduleActionsForRulesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}/microsoft.graph.scheduleActionsForRules";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewScheduleActionsForRulesRequestBuilder instantiates a new ScheduleActionsForRulesRequestBuilder and sets the default values.
func NewScheduleActionsForRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScheduleActionsForRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduleActionsForRulesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action scheduleActionsForRules
func (m *ScheduleActionsForRulesRequestBuilder) CreatePostRequestInformation(body ScheduleActionsForRulesPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action scheduleActionsForRules
func (m *ScheduleActionsForRulesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ScheduleActionsForRulesPostRequestBodyable, requestConfiguration *ScheduleActionsForRulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action scheduleActionsForRules
func (m *ScheduleActionsForRulesRequestBuilder) Post(body ScheduleActionsForRulesPostRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action scheduleActionsForRules
func (m *ScheduleActionsForRulesRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ScheduleActionsForRulesPostRequestBodyable, requestConfiguration *ScheduleActionsForRulesRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
