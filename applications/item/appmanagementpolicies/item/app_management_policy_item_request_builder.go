package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    id8515422dad04d2c65370d4049b61c9421b519e76f3c23c14869fae96b755564 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/appmanagementpolicies/item/ref"
)

// AppManagementPolicyItemRequestBuilder builds and executes requests for operations under \applications\{application-id}\appManagementPolicies\{appManagementPolicy-id}
type AppManagementPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewAppManagementPolicyItemRequestBuilderInternal instantiates a new AppManagementPolicyItemRequestBuilder and sets the default values.
func NewAppManagementPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppManagementPolicyItemRequestBuilder) {
    m := &AppManagementPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/appManagementPolicies/{appManagementPolicy%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAppManagementPolicyItemRequestBuilder instantiates a new AppManagementPolicyItemRequestBuilder and sets the default values.
func NewAppManagementPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppManagementPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppManagementPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref the ref property
func (m *AppManagementPolicyItemRequestBuilder) Ref()(*id8515422dad04d2c65370d4049b61c9421b519e76f3c23c14869fae96b755564.RefRequestBuilder) {
    return id8515422dad04d2c65370d4049b61c9421b519e76f3c23c14869fae96b755564.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
