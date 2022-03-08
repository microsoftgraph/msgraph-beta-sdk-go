package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i17d1574c3c57286814540657170c4b7233c1c6f26ec60bea16e08bf0678e59a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/userconfigurations"
    i541fc3399419d1ae2cab4368370a7227e59f2c04b7e54e4929b86ebe04698638 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages"
    i58c44598f2855313799f5f0ec80b7d31cf45e85c1903cdc2ee9bd6c5776786db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/move"
    i5bed6e35b48c78b26b58da0c35269f9fbe5e11549479f8aafb35b126f8e25ab5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/singlevalueextendedproperties"
    i88c04acf0aa04d581c2773f019a56cb1fdf513359c2fe21e75dcc60ebc80d511 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messagerules"
    ibeb8c4168d977b698b05801a1b10798f68c64e767d766ce8595176f72fc5093b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders"
    id07076013853ba5080f44b31c440adb7eef26be6bf39605515485ec5ece17190 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/multivalueextendedproperties"
    if3fb3684743655c918bd082004c113091359b0d53e481b54379b0fd767244c8d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/copy"
    i5093bb4836672beb5dd5657e003961ad167c9ffe6b1cfc7663eb7832d4a6de57 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/userconfigurations/item"
    i5e5f56aacb3cd9d3f5d914c12525e9de5b7e2db88b3753d1b0e65ad922a94eb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/singlevalueextendedproperties/item"
    i779f11aeb25854b5bc2350fde107c6812b983ae14fd525417dabc92b85de7d0a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/multivalueextendedproperties/item"
    i8ffc8cfef5f5058547bf175ad9868544c797b8dde19d8a6459fe8494963d0a24 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messagerules/item"
    ic4853dfe79219076841dce140d8e62c905b0804ee816e62386841f353c4b1148 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item"
    iede75dd69fcd9e5b626950014b98ed33d606c522a9728b4f3a33a890e6632856 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item"
)

// MailFolderItemRequestBuilder provides operations to manage the mailFolders property of the microsoft.graph.user entity.
type MailFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MailFolderItemRequestBuilderDeleteOptions options for Delete
type MailFolderItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MailFolderItemRequestBuilderGetOptions options for Get
type MailFolderItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MailFolderItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MailFolderItemRequestBuilderGetQueryParameters the user's mail folders. Read-only. Nullable.
type MailFolderItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// MailFolderItemRequestBuilderPatchOptions options for Patch
type MailFolderItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MailFolderable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MailFolderItemRequestBuilder) ChildFolders()(*ibeb8c4168d977b698b05801a1b10798f68c64e767d766ce8595176f72fc5093b.ChildFoldersRequestBuilder) {
    return ibeb8c4168d977b698b05801a1b10798f68c64e767d766ce8595176f72fc5093b.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildFoldersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.childFolders.item collection
func (m *MailFolderItemRequestBuilder) ChildFoldersById(id string)(*iede75dd69fcd9e5b626950014b98ed33d606c522a9728b4f3a33a890e6632856.MailFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder_id1"] = id
    }
    return iede75dd69fcd9e5b626950014b98ed33d606c522a9728b4f3a33a890e6632856.NewMailFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMailFolderItemRequestBuilderInternal instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderItemRequestBuilder) {
    m := &MailFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMailFolderItemRequestBuilder instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMailFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MailFolderItemRequestBuilder) Copy()(*if3fb3684743655c918bd082004c113091359b0d53e481b54379b0fd767244c8d.CopyRequestBuilder) {
    return if3fb3684743655c918bd082004c113091359b0d53e481b54379b0fd767244c8d.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property mailFolders for me
func (m *MailFolderItemRequestBuilder) CreateDeleteRequestInformation(options *MailFolderItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the user's mail folders. Read-only. Nullable.
func (m *MailFolderItemRequestBuilder) CreateGetRequestInformation(options *MailFolderItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property mailFolders in me
func (m *MailFolderItemRequestBuilder) CreatePatchRequestInformation(options *MailFolderItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property mailFolders for me
func (m *MailFolderItemRequestBuilder) Delete(options *MailFolderItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the user's mail folders. Read-only. Nullable.
func (m *MailFolderItemRequestBuilder) Get(options *MailFolderItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MailFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateMailFolderFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MailFolderable), nil
}
func (m *MailFolderItemRequestBuilder) MessageRules()(*i88c04acf0aa04d581c2773f019a56cb1fdf513359c2fe21e75dcc60ebc80d511.MessageRulesRequestBuilder) {
    return i88c04acf0aa04d581c2773f019a56cb1fdf513359c2fe21e75dcc60ebc80d511.NewMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.messageRules.item collection
func (m *MailFolderItemRequestBuilder) MessageRulesById(id string)(*i8ffc8cfef5f5058547bf175ad9868544c797b8dde19d8a6459fe8494963d0a24.MessageRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRule_id"] = id
    }
    return i8ffc8cfef5f5058547bf175ad9868544c797b8dde19d8a6459fe8494963d0a24.NewMessageRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderItemRequestBuilder) Messages()(*i541fc3399419d1ae2cab4368370a7227e59f2c04b7e54e4929b86ebe04698638.MessagesRequestBuilder) {
    return i541fc3399419d1ae2cab4368370a7227e59f2c04b7e54e4929b86ebe04698638.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.messages.item collection
func (m *MailFolderItemRequestBuilder) MessagesById(id string)(*ic4853dfe79219076841dce140d8e62c905b0804ee816e62386841f353c4b1148.MessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message_id"] = id
    }
    return ic4853dfe79219076841dce140d8e62c905b0804ee816e62386841f353c4b1148.NewMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderItemRequestBuilder) Move()(*i58c44598f2855313799f5f0ec80b7d31cf45e85c1903cdc2ee9bd6c5776786db.MoveRequestBuilder) {
    return i58c44598f2855313799f5f0ec80b7d31cf45e85c1903cdc2ee9bd6c5776786db.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderItemRequestBuilder) MultiValueExtendedProperties()(*id07076013853ba5080f44b31c440adb7eef26be6bf39605515485ec5ece17190.MultiValueExtendedPropertiesRequestBuilder) {
    return id07076013853ba5080f44b31c440adb7eef26be6bf39605515485ec5ece17190.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.multiValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i779f11aeb25854b5bc2350fde107c6812b983ae14fd525417dabc92b85de7d0a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i779f11aeb25854b5bc2350fde107c6812b983ae14fd525417dabc92b85de7d0a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property mailFolders in me
func (m *MailFolderItemRequestBuilder) Patch(options *MailFolderItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *MailFolderItemRequestBuilder) SingleValueExtendedProperties()(*i5bed6e35b48c78b26b58da0c35269f9fbe5e11549479f8aafb35b126f8e25ab5.SingleValueExtendedPropertiesRequestBuilder) {
    return i5bed6e35b48c78b26b58da0c35269f9fbe5e11549479f8aafb35b126f8e25ab5.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.singleValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i5e5f56aacb3cd9d3f5d914c12525e9de5b7e2db88b3753d1b0e65ad922a94eb7.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i5e5f56aacb3cd9d3f5d914c12525e9de5b7e2db88b3753d1b0e65ad922a94eb7.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderItemRequestBuilder) UserConfigurations()(*i17d1574c3c57286814540657170c4b7233c1c6f26ec60bea16e08bf0678e59a0.UserConfigurationsRequestBuilder) {
    return i17d1574c3c57286814540657170c4b7233c1c6f26ec60bea16e08bf0678e59a0.NewUserConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.userConfigurations.item collection
func (m *MailFolderItemRequestBuilder) UserConfigurationsById(id string)(*i5093bb4836672beb5dd5657e003961ad167c9ffe6b1cfc7663eb7832d4a6de57.UserConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userConfiguration_id"] = id
    }
    return i5093bb4836672beb5dd5657e003961ad167c9ffe6b1cfc7663eb7832d4a6de57.NewUserConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
