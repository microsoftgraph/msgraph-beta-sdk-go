package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemValidatePropertiesRequestBuilder provides operations to call the validateProperties method.
type GroupsItemValidatePropertiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemValidatePropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemValidatePropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupsItemValidatePropertiesRequestBuilderInternal instantiates a new ValidatePropertiesRequestBuilder and sets the default values.
func NewGroupsItemValidatePropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemValidatePropertiesRequestBuilder) {
    m := &GroupsItemValidatePropertiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/microsoft.graph.validateProperties";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemValidatePropertiesRequestBuilder instantiates a new ValidatePropertiesRequestBuilder and sets the default values.
func NewGroupsItemValidatePropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemValidatePropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemValidatePropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation validate if a Microsoft 365 group's display name or mail nickname complies with naming policies. Clients can use the API to determine if a display name or mail nickname is valid before trying to **update** a Microsoft 365 group. For validating properties before creating a group, use the validateProperties function for directory objects. The following validations are performed for the display name and mail nickname properties: This API returns with the first failure encountered. If one or more properties fail multiple validations, only the property with the first validation failure is returned. However, you can validate both the mail nickname and the display name and receive a collection of validation errors if you are only validating the prefix and suffix naming policy.
func (m *GroupsItemValidatePropertiesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupsItemValidatePropertiesPostRequestBodyable, requestConfiguration *GroupsItemValidatePropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post validate if a Microsoft 365 group's display name or mail nickname complies with naming policies. Clients can use the API to determine if a display name or mail nickname is valid before trying to **update** a Microsoft 365 group. For validating properties before creating a group, use the validateProperties function for directory objects. The following validations are performed for the display name and mail nickname properties: This API returns with the first failure encountered. If one or more properties fail multiple validations, only the property with the first validation failure is returned. However, you can validate both the mail nickname and the display name and receive a collection of validation errors if you are only validating the prefix and suffix naming policy.
func (m *GroupsItemValidatePropertiesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupsItemValidatePropertiesPostRequestBodyable, requestConfiguration *GroupsItemValidatePropertiesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
