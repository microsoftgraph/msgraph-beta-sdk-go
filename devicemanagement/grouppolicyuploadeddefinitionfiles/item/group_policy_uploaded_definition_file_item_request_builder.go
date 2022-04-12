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
// GroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteOptions options for Delete
type GroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// GroupPolicyUploadedDefinitionFileItemRequestBuilderGetOptions options for Get
type GroupPolicyUploadedDefinitionFileItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// GroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters the available group policy uploaded definition files for this account.
type GroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyUploadedDefinitionFileItemRequestBuilderPatchOptions options for Patch
type GroupPolicyUploadedDefinitionFileItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
// CreateDeleteRequestInformation delete navigation property groupPolicyUploadedDefinitionFiles for deviceManagement
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
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
// CreateGetRequestInformation the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) CreateGetRequestInformation(options *GroupPolicyUploadedDefinitionFileItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
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
// CreatePatchRequestInformation update the navigation property groupPolicyUploadedDefinitionFiles in deviceManagement
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyUploadedDefinitionFileItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
// Delete delete navigation property groupPolicyUploadedDefinitionFiles for deviceManagement
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) Delete(options *GroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) Get(options *GroupPolicyUploadedDefinitionFileItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyUploadedDefinitionFileFromDiscriminatorValue, nil, errorMapping)
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
// Patch update the navigation property groupPolicyUploadedDefinitionFiles in deviceManagement
func (m *GroupPolicyUploadedDefinitionFileItemRequestBuilder) Patch(options *GroupPolicyUploadedDefinitionFileItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
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
