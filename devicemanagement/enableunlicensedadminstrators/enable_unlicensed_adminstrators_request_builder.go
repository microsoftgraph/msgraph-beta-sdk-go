package enableunlicensedadminstrators

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EnableUnlicensedAdminstratorsRequestBuilder provides operations to call the enableUnlicensedAdminstrators method.
type EnableUnlicensedAdminstratorsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EnableUnlicensedAdminstratorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnableUnlicensedAdminstratorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEnableUnlicensedAdminstratorsRequestBuilderInternal instantiates a new EnableUnlicensedAdminstratorsRequestBuilder and sets the default values.
func NewEnableUnlicensedAdminstratorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnableUnlicensedAdminstratorsRequestBuilder) {
    m := &EnableUnlicensedAdminstratorsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.enableUnlicensedAdminstrators";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEnableUnlicensedAdminstratorsRequestBuilder instantiates a new EnableUnlicensedAdminstratorsRequestBuilder and sets the default values.
func NewEnableUnlicensedAdminstratorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnableUnlicensedAdminstratorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnableUnlicensedAdminstratorsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration upon enabling, users assigned as administrators via Role Assignment Memberships will no longer require an assigned Intune license. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators will continue to function as-is in that transitive memberships apply and are not subject to the 350 member limit.
func (m *EnableUnlicensedAdminstratorsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration upon enabling, users assigned as administrators via Role Assignment Memberships will no longer require an assigned Intune license. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators will continue to function as-is in that transitive memberships apply and are not subject to the 350 member limit.
func (m *EnableUnlicensedAdminstratorsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *EnableUnlicensedAdminstratorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// PostWithResponseHandler upon enabling, users assigned as administrators via Role Assignment Memberships will no longer require an assigned Intune license. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators will continue to function as-is in that transitive memberships apply and are not subject to the 350 member limit.
func (m *EnableUnlicensedAdminstratorsRequestBuilder) PostWithResponseHandler(requestConfiguration *EnableUnlicensedAdminstratorsRequestBuilderPostRequestConfiguration)(error) {
    return m.PostWithResponseHandler(requestConfiguration, nil);
}
// PostWithResponseHandler upon enabling, users assigned as administrators via Role Assignment Memberships will no longer require an assigned Intune license. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators will continue to function as-is in that transitive memberships apply and are not subject to the 350 member limit.
func (m *EnableUnlicensedAdminstratorsRequestBuilder) PostWithResponseHandler(requestConfiguration *EnableUnlicensedAdminstratorsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
