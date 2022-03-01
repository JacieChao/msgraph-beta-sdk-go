package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3c30639732fa96ae7b2270ea94ac4d733ab4ce3489526a5e5cd1c27cbb90b0df "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/identityproviders"
    ia8501a142e35d07776280d07596b55002a93c3239073e18d8e5ae9479c0a6576 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userattributeassignments"
    ib718f8165ffdc29693dd68518ac7ea324def0c6882130bda31e56b4f400e55ed "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userflowidentityproviders"
    ic8b27e7bcdb899e4a47d9ea30cb80e33646e9ca77dcb752f0bec144f3c6411be "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages"
    i4f9895372ea85cee1ee4bf9a8b414752e4270d4b2216846b499db783b5644d6c "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages/item"
    ib4590aff7f09f446c7ae6a4a70dfc243aac07fb4522106f724124d4539ac2888 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userattributeassignments/item"
)

// B2cIdentityUserFlowItemRequestBuilder builds and executes requests for operations under \identity\b2cUserFlows\{b2cIdentityUserFlow-id}
type B2cIdentityUserFlowItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// B2cIdentityUserFlowItemRequestBuilderDeleteOptions options for Delete
type B2cIdentityUserFlowItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// B2cIdentityUserFlowItemRequestBuilderGetOptions options for Get
type B2cIdentityUserFlowItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *B2cIdentityUserFlowItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// B2cIdentityUserFlowItemRequestBuilderGetQueryParameters represents entry point for B2C identity userflows.
type B2cIdentityUserFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// B2cIdentityUserFlowItemRequestBuilderPatchOptions options for Patch
type B2cIdentityUserFlowItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.B2cIdentityUserFlow;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewB2cIdentityUserFlowItemRequestBuilderInternal instantiates a new B2cIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2cIdentityUserFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*B2cIdentityUserFlowItemRequestBuilder) {
    m := &B2cIdentityUserFlowItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewB2cIdentityUserFlowItemRequestBuilder instantiates a new B2cIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2cIdentityUserFlowItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*B2cIdentityUserFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cIdentityUserFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents entry point for B2C identity userflows.
func (m *B2cIdentityUserFlowItemRequestBuilder) CreateDeleteRequestInformation(options *B2cIdentityUserFlowItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents entry point for B2C identity userflows.
func (m *B2cIdentityUserFlowItemRequestBuilder) CreateGetRequestInformation(options *B2cIdentityUserFlowItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents entry point for B2C identity userflows.
func (m *B2cIdentityUserFlowItemRequestBuilder) CreatePatchRequestInformation(options *B2cIdentityUserFlowItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents entry point for B2C identity userflows.
func (m *B2cIdentityUserFlowItemRequestBuilder) Delete(options *B2cIdentityUserFlowItemRequestBuilderDeleteOptions)(error) {
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
// Get represents entry point for B2C identity userflows.
func (m *B2cIdentityUserFlowItemRequestBuilder) Get(options *B2cIdentityUserFlowItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.B2cIdentityUserFlow, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewB2cIdentityUserFlow() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.B2cIdentityUserFlow), nil
}
func (m *B2cIdentityUserFlowItemRequestBuilder) IdentityProviders()(*i3c30639732fa96ae7b2270ea94ac4d733ab4ce3489526a5e5cd1c27cbb90b0df.IdentityProvidersRequestBuilder) {
    return i3c30639732fa96ae7b2270ea94ac4d733ab4ce3489526a5e5cd1c27cbb90b0df.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *B2cIdentityUserFlowItemRequestBuilder) Languages()(*ic8b27e7bcdb899e4a47d9ea30cb80e33646e9ca77dcb752f0bec144f3c6411be.LanguagesRequestBuilder) {
    return ic8b27e7bcdb899e4a47d9ea30cb80e33646e9ca77dcb752f0bec144f3c6411be.NewLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2cUserFlows.item.languages.item collection
func (m *B2cIdentityUserFlowItemRequestBuilder) LanguagesById(id string)(*i4f9895372ea85cee1ee4bf9a8b414752e4270d4b2216846b499db783b5644d6c.UserFlowLanguageConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguageConfiguration_id"] = id
    }
    return i4f9895372ea85cee1ee4bf9a8b414752e4270d4b2216846b499db783b5644d6c.NewUserFlowLanguageConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch represents entry point for B2C identity userflows.
func (m *B2cIdentityUserFlowItemRequestBuilder) Patch(options *B2cIdentityUserFlowItemRequestBuilderPatchOptions)(error) {
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
func (m *B2cIdentityUserFlowItemRequestBuilder) UserAttributeAssignments()(*ia8501a142e35d07776280d07596b55002a93c3239073e18d8e5ae9479c0a6576.UserAttributeAssignmentsRequestBuilder) {
    return ia8501a142e35d07776280d07596b55002a93c3239073e18d8e5ae9479c0a6576.NewUserAttributeAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserAttributeAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2cUserFlows.item.userAttributeAssignments.item collection
func (m *B2cIdentityUserFlowItemRequestBuilder) UserAttributeAssignmentsById(id string)(*ib4590aff7f09f446c7ae6a4a70dfc243aac07fb4522106f724124d4539ac2888.IdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttributeAssignment_id"] = id
    }
    return ib4590aff7f09f446c7ae6a4a70dfc243aac07fb4522106f724124d4539ac2888.NewIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *B2cIdentityUserFlowItemRequestBuilder) UserFlowIdentityProviders()(*ib718f8165ffdc29693dd68518ac7ea324def0c6882130bda31e56b4f400e55ed.UserFlowIdentityProvidersRequestBuilder) {
    return ib718f8165ffdc29693dd68518ac7ea324def0c6882130bda31e56b4f400e55ed.NewUserFlowIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}