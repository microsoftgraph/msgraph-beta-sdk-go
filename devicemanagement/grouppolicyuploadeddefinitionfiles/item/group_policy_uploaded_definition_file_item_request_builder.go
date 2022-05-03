package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i08aa694b5efe17e6cfb40a0b57c60c903682f60b25cedf6b77b700e69456e371 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/addlanguagefiles"
    i30872138ba0a6007708c6d50246b61e9dbb1c6693196e74b8fe1b81dbab24eed "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/remove"
    i573a2f06fa9d3a5f4c3539c0c567663d335bba6395fe3de82bf7752385eacc91 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/grouppolicyoperations"
    ib8f5596b1f4d2b0563c2e8fec9b75de399dfc5c22c6593b08984b3c43081a2f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/uploadnewversion"
    ie6c6ba293495002c0854fd282065a1b58a7aa28885d4f70b197baa6c4d5221e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/updatelanguagefiles"
    ie97617c59d880c34963563417867a7e1047adb153e1210b59f965efbd77e2b11 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/removelanguagefiles"
    ib268e3b08fe1fad55fddc2f3fcee0a20255a98eba2bf68e561d4e06a67c8ce93 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyuploadeddefinitionfiles/item/grouppolicyoperations/item"
)

// GroupPolicyUploadedDefinitionFileItemRequestBuilder provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
type GroupPolicyUploadedDefinitionFileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters the available group policy uploaded definition files for this account.
type GroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters
}
// GroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddLanguageFiles the addLanguageFiles property
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) AddLanguageFiles()(*i08aa694b5efe17e6cfb40a0b57c60c903682f60b25cedf6b77b700e69456e371.AddLanguageFilesRequestBuilder) {
    return i08aa694b5efe17e6cfb40a0b57c60c903682f60b25cedf6b77b700e69456e371.NewAddLanguageFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal instantiates a new GroupPolicyUploadedDefinitionFileItemRequestBuilder and sets the default values.
func NewGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    m := &GroupPolicyUploadedDefinitionFileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupPolicyUploadedDefinitionFileItemRequestBuilder instantiates a new GroupPolicyUploadedDefinitionFileItemRequestBuilder and sets the default values.
func NewGroupPolicyUploadedDefinitionFileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property groupPolicyUploadedDefinitionFiles for deviceManagement
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property groupPolicyUploadedDefinitionFiles for deviceManagement
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *GroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformationWithRequestConfiguration the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property groupPolicyUploadedDefinitionFiles in deviceManagement
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property groupPolicyUploadedDefinitionFiles in deviceManagement
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, requestConfiguration *GroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeleteWithResponseHandler delete navigation property groupPolicyUploadedDefinitionFiles for deviceManagement
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *GroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property groupPolicyUploadedDefinitionFiles for deviceManagement
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *GroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) GetWithResponseHandler(requestConfiguration *GroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) GetWithResponseHandler(requestConfiguration *GroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyUploadedDefinitionFileFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable), nil
}
// GroupPolicyOperations the groupPolicyOperations property
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) GroupPolicyOperations()(*i573a2f06fa9d3a5f4c3539c0c567663d335bba6395fe3de82bf7752385eacc91.GroupPolicyOperationsRequestBuilder) {
    return i573a2f06fa9d3a5f4c3539c0c567663d335bba6395fe3de82bf7752385eacc91.NewGroupPolicyOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyOperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyUploadedDefinitionFiles.item.groupPolicyOperations.item collection
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) GroupPolicyOperationsById(id string)(*ib268e3b08fe1fad55fddc2f3fcee0a20255a98eba2bf68e561d4e06a67c8ce93.GroupPolicyOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyOperation%2Did"] = id
    }
    return ib268e3b08fe1fad55fddc2f3fcee0a20255a98eba2bf68e561d4e06a67c8ce93.NewGroupPolicyOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property groupPolicyUploadedDefinitionFiles in deviceManagement
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, requestConfiguration *GroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property groupPolicyUploadedDefinitionFiles in deviceManagement
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, requestConfiguration *GroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Remove the remove property
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) Remove()(*i30872138ba0a6007708c6d50246b61e9dbb1c6693196e74b8fe1b81dbab24eed.RemoveRequestBuilder) {
    return i30872138ba0a6007708c6d50246b61e9dbb1c6693196e74b8fe1b81dbab24eed.NewRemoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveLanguageFiles the removeLanguageFiles property
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) RemoveLanguageFiles()(*ie97617c59d880c34963563417867a7e1047adb153e1210b59f965efbd77e2b11.RemoveLanguageFilesRequestBuilder) {
    return ie97617c59d880c34963563417867a7e1047adb153e1210b59f965efbd77e2b11.NewRemoveLanguageFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateLanguageFiles the updateLanguageFiles property
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) UpdateLanguageFiles()(*ie6c6ba293495002c0854fd282065a1b58a7aa28885d4f70b197baa6c4d5221e4.UpdateLanguageFilesRequestBuilder) {
    return ie6c6ba293495002c0854fd282065a1b58a7aa28885d4f70b197baa6c4d5221e4.NewUpdateLanguageFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadNewVersion the uploadNewVersion property
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) UploadNewVersion()(*ib8f5596b1f4d2b0563c2e8fec9b75de399dfc5c22c6593b08984b3c43081a2f1.UploadNewVersionRequestBuilder) {
    return ib8f5596b1f4d2b0563c2e8fec9b75de399dfc5c22c6593b08984b3c43081a2f1.NewUploadNewVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
