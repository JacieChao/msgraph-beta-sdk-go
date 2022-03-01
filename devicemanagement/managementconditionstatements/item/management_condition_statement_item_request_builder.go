package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i77e32bad600df9fdcde26f80b9bb02e4bdb8ec9d3b6c0558e7ef39af19f17027 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/managementconditionstatements/item/getmanagementconditionstatementexpressionstring"
    i8ef25e2265384c29527019c0f7a1886c2da03ccafa2470739869626ce85b62d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/managementconditionstatements/item/managementconditions"
)

// ManagementConditionStatementItemRequestBuilder builds and executes requests for operations under \deviceManagement\managementConditionStatements\{managementConditionStatement-id}
type ManagementConditionStatementItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagementConditionStatementItemRequestBuilderDeleteOptions options for Delete
type ManagementConditionStatementItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagementConditionStatementItemRequestBuilderGetOptions options for Get
type ManagementConditionStatementItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagementConditionStatementItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagementConditionStatementItemRequestBuilderGetQueryParameters the management condition statements associated with device management of the company.
type ManagementConditionStatementItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagementConditionStatementItemRequestBuilderPatchOptions options for Patch
type ManagementConditionStatementItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementConditionStatement;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewManagementConditionStatementItemRequestBuilderInternal instantiates a new ManagementConditionStatementItemRequestBuilder and sets the default values.
func NewManagementConditionStatementItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementConditionStatementItemRequestBuilder) {
    m := &ManagementConditionStatementItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managementConditionStatements/{managementConditionStatement_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagementConditionStatementItemRequestBuilder instantiates a new ManagementConditionStatementItemRequestBuilder and sets the default values.
func NewManagementConditionStatementItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementConditionStatementItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagementConditionStatementItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the management condition statements associated with device management of the company.
func (m *ManagementConditionStatementItemRequestBuilder) CreateDeleteRequestInformation(options *ManagementConditionStatementItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the management condition statements associated with device management of the company.
func (m *ManagementConditionStatementItemRequestBuilder) CreateGetRequestInformation(options *ManagementConditionStatementItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the management condition statements associated with device management of the company.
func (m *ManagementConditionStatementItemRequestBuilder) CreatePatchRequestInformation(options *ManagementConditionStatementItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the management condition statements associated with device management of the company.
func (m *ManagementConditionStatementItemRequestBuilder) Delete(options *ManagementConditionStatementItemRequestBuilderDeleteOptions)(error) {
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
// Get the management condition statements associated with device management of the company.
func (m *ManagementConditionStatementItemRequestBuilder) Get(options *ManagementConditionStatementItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementConditionStatement, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagementConditionStatement() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagementConditionStatement), nil
}
// GetManagementConditionStatementExpressionString builds and executes requests for operations under \deviceManagement\managementConditionStatements\{managementConditionStatement-id}\microsoft.graph.getManagementConditionStatementExpressionString()
func (m *ManagementConditionStatementItemRequestBuilder) GetManagementConditionStatementExpressionString()(*i77e32bad600df9fdcde26f80b9bb02e4bdb8ec9d3b6c0558e7ef39af19f17027.GetManagementConditionStatementExpressionStringRequestBuilder) {
    return i77e32bad600df9fdcde26f80b9bb02e4bdb8ec9d3b6c0558e7ef39af19f17027.NewGetManagementConditionStatementExpressionStringRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagementConditionStatementItemRequestBuilder) ManagementConditions()(*i8ef25e2265384c29527019c0f7a1886c2da03ccafa2470739869626ce85b62d1.ManagementConditionsRequestBuilder) {
    return i8ef25e2265384c29527019c0f7a1886c2da03ccafa2470739869626ce85b62d1.NewManagementConditionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the management condition statements associated with device management of the company.
func (m *ManagementConditionStatementItemRequestBuilder) Patch(options *ManagementConditionStatementItemRequestBuilderPatchOptions)(error) {
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