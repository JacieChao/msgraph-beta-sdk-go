package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilder provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
type LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilderGetQueryParameters the unique identifier of the Microsoft Entra identity that last modified the workflow object.
type LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilderInternal instantiates a new UserProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilder) {
    m := &LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/executionScope/{userProcessingResult%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilder instantiates a new UserProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the unique identifier of the Microsoft Entra identity that last modified the workflow object.
func (m *LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.UserProcessingResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateUserProcessingResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.UserProcessingResultable), nil
}
// ToGetRequestInformation the unique identifier of the Microsoft Entra identity that last modified the workflow object.
func (m *LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemExecutionScopeUserProcessingResultItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
