package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i11b56dd67bcd5d9695e59af2c58c55472a534b18112364934482209d913c373b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages"
    i6878489e6e5c5366487958b139d296196316f6b48e91da13549ca84158b58619 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/userconfigurations"
    i8865115e0c4a7f8523c7b71edc37e7144f9ccbfcb140c502f78ed325a5b6f804 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/copy"
    i8c8bf593802b6296e9252991dc3e727323b3780ee21b0bf46d0d7681e382e439 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/singlevalueextendedproperties"
    i8d96f40e0c4d524a8a0695f4ebdfdb11706bd4042ae2bf8d271db546df18f5e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/move"
    ib3cf1b898320f5fb70f999d2aff83698cee085e8dbf2e1b033c20e55909c2147 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders"
    ib552f46594b7d9ebc22e518b339ef30511852f158eb9110fa26359c22cac8509 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messagerules"
    icb8f096fe09e8f7c4f02910d715c5b1b7763db2ea388eea3a79b9f6a1c5bce57 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/multivalueextendedproperties"
    i310b07fcb064069d6015f73b2ca186f6408a1a4dd1e21397434cfff499a2801b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messagerules/item"
    i4659db7620dc79f8074a55700b4675ba25eb98811afff3bd143e845e53916070 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/userconfigurations/item"
    i688b613b4c12c110c627749a3da9032a2c37f515bc98ee032ce8a1d7ba196a04 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/singlevalueextendedproperties/item"
    i9584fba2233b5641465d19d502543b266d2bee87dfef4b622dc1a93ebf076c8c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item"
    i99d414649a0b801dbd6f28f6fabe9640fa48f50ba20d21ff4dfc91c4e616dd1d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item"
    iddf7155795dc366c49b367caf52234a860a61a06b42dd87bc5aa2b9c61fe336a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/multivalueextendedproperties/item"
)

// MailFolderRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}
type MailFolderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MailFolderRequestBuilderDeleteOptions options for Delete
type MailFolderRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MailFolderRequestBuilderGetOptions options for Get
type MailFolderRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MailFolderRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MailFolderRequestBuilderGetQueryParameters the user's mail folders. Read-only. Nullable.
type MailFolderRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// MailFolderRequestBuilderPatchOptions options for Patch
type MailFolderRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MailFolder;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MailFolderRequestBuilder) ChildFolders()(*ib3cf1b898320f5fb70f999d2aff83698cee085e8dbf2e1b033c20e55909c2147.ChildFoldersRequestBuilder) {
    return ib3cf1b898320f5fb70f999d2aff83698cee085e8dbf2e1b033c20e55909c2147.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildFoldersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.childFolders.item collection
func (m *MailFolderRequestBuilder) ChildFoldersById(id string)(*i99d414649a0b801dbd6f28f6fabe9640fa48f50ba20d21ff4dfc91c4e616dd1d.MailFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder_id1"] = id
    }
    return i99d414649a0b801dbd6f28f6fabe9640fa48f50ba20d21ff4dfc91c4e616dd1d.NewMailFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMailFolderRequestBuilderInternal instantiates a new MailFolderRequestBuilder and sets the default values.
func NewMailFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderRequestBuilder) {
    m := &MailFolderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMailFolderRequestBuilder instantiates a new MailFolderRequestBuilder and sets the default values.
func NewMailFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMailFolderRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MailFolderRequestBuilder) Copy()(*i8865115e0c4a7f8523c7b71edc37e7144f9ccbfcb140c502f78ed325a5b6f804.CopyRequestBuilder) {
    return i8865115e0c4a7f8523c7b71edc37e7144f9ccbfcb140c502f78ed325a5b6f804.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the user's mail folders. Read-only. Nullable.
func (m *MailFolderRequestBuilder) CreateDeleteRequestInformation(options *MailFolderRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MailFolderRequestBuilder) CreateGetRequestInformation(options *MailFolderRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the user's mail folders. Read-only. Nullable.
func (m *MailFolderRequestBuilder) CreatePatchRequestInformation(options *MailFolderRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the user's mail folders. Read-only. Nullable.
func (m *MailFolderRequestBuilder) Delete(options *MailFolderRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the user's mail folders. Read-only. Nullable.
func (m *MailFolderRequestBuilder) Get(options *MailFolderRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MailFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMailFolder() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MailFolder), nil
}
func (m *MailFolderRequestBuilder) MessageRules()(*ib552f46594b7d9ebc22e518b339ef30511852f158eb9110fa26359c22cac8509.MessageRulesRequestBuilder) {
    return ib552f46594b7d9ebc22e518b339ef30511852f158eb9110fa26359c22cac8509.NewMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.messageRules.item collection
func (m *MailFolderRequestBuilder) MessageRulesById(id string)(*i310b07fcb064069d6015f73b2ca186f6408a1a4dd1e21397434cfff499a2801b.MessageRuleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRule_id"] = id
    }
    return i310b07fcb064069d6015f73b2ca186f6408a1a4dd1e21397434cfff499a2801b.NewMessageRuleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) Messages()(*i11b56dd67bcd5d9695e59af2c58c55472a534b18112364934482209d913c373b.MessagesRequestBuilder) {
    return i11b56dd67bcd5d9695e59af2c58c55472a534b18112364934482209d913c373b.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.messages.item collection
func (m *MailFolderRequestBuilder) MessagesById(id string)(*i9584fba2233b5641465d19d502543b266d2bee87dfef4b622dc1a93ebf076c8c.MessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message_id"] = id
    }
    return i9584fba2233b5641465d19d502543b266d2bee87dfef4b622dc1a93ebf076c8c.NewMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) Move()(*i8d96f40e0c4d524a8a0695f4ebdfdb11706bd4042ae2bf8d271db546df18f5e4.MoveRequestBuilder) {
    return i8d96f40e0c4d524a8a0695f4ebdfdb11706bd4042ae2bf8d271db546df18f5e4.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) MultiValueExtendedProperties()(*icb8f096fe09e8f7c4f02910d715c5b1b7763db2ea388eea3a79b9f6a1c5bce57.MultiValueExtendedPropertiesRequestBuilder) {
    return icb8f096fe09e8f7c4f02910d715c5b1b7763db2ea388eea3a79b9f6a1c5bce57.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.multiValueExtendedProperties.item collection
func (m *MailFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iddf7155795dc366c49b367caf52234a860a61a06b42dd87bc5aa2b9c61fe336a.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iddf7155795dc366c49b367caf52234a860a61a06b42dd87bc5aa2b9c61fe336a.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the user's mail folders. Read-only. Nullable.
func (m *MailFolderRequestBuilder) Patch(options *MailFolderRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *MailFolderRequestBuilder) SingleValueExtendedProperties()(*i8c8bf593802b6296e9252991dc3e727323b3780ee21b0bf46d0d7681e382e439.SingleValueExtendedPropertiesRequestBuilder) {
    return i8c8bf593802b6296e9252991dc3e727323b3780ee21b0bf46d0d7681e382e439.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.singleValueExtendedProperties.item collection
func (m *MailFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i688b613b4c12c110c627749a3da9032a2c37f515bc98ee032ce8a1d7ba196a04.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i688b613b4c12c110c627749a3da9032a2c37f515bc98ee032ce8a1d7ba196a04.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) UserConfigurations()(*i6878489e6e5c5366487958b139d296196316f6b48e91da13549ca84158b58619.UserConfigurationsRequestBuilder) {
    return i6878489e6e5c5366487958b139d296196316f6b48e91da13549ca84158b58619.NewUserConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.userConfigurations.item collection
func (m *MailFolderRequestBuilder) UserConfigurationsById(id string)(*i4659db7620dc79f8074a55700b4675ba25eb98811afff3bd143e845e53916070.UserConfigurationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userConfiguration_id"] = id
    }
    return i4659db7620dc79f8074a55700b4675ba25eb98811afff3bd143e845e53916070.NewUserConfigurationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
