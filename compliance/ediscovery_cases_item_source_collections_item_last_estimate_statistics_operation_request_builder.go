package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilder provides operations to manage the lastEstimateStatisticsOperation property of the microsoft.graph.ediscovery.sourceCollection entity.
type EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderGetQueryParameters the last estimate operation associated with the sourceCollection.
type EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderGetQueryParameters
}
// NewEdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderInternal instantiates a new LastEstimateStatisticsOperationRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilder) {
    m := &EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}/lastEstimateStatisticsOperation{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilder instantiates a new LastEstimateStatisticsOperationRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the last estimate operation associated with the sourceCollection.
func (m *EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderGetRequestConfiguration)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.EstimateStatisticsOperationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateEstimateStatisticsOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.EstimateStatisticsOperationable), nil
}
// ToGetRequestInformation the last estimate operation associated with the sourceCollection.
func (m *EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsItemLastEstimateStatisticsOperationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
