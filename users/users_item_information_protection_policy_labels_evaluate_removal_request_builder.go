package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type UsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderInternal instantiates a new EvaluateRemovalRequestBuilder and sets the default values.
func NewUsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) {
    m := &UsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/informationProtection/policy/labels/microsoft.graph.evaluateRemoval";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder instantiates a new EvaluateRemovalRequestBuilder and sets the default values.
func NewUsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
func (m *UsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBodyable, requestConfiguration *UsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
func (m *UsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBodyable, requestConfiguration *UsersItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemInformationProtectionPolicyLabelsEvaluateRemovalResponseable), nil
}
