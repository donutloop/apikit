package todo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type todoServiceClient struct {
	baseURL    string
	hooks      HooksClient
	ctx        context.Context
	httpClient *httpClientWrapper
	xmlMatcher *regexp.Regexp
}

func (client *todoServiceClient) DeleteTodos(request *DeleteTodosRequest) (DeleteTodosResponse, error) {
	if request == nil {
		return nil, newRequestObjectIsNilError
	}
	path := "/todos"
	method := "DELETE"
	endpoint := client.baseURL + path
	httpContext := newHttpContextWrapper(client.ctx)
	httpRequest, reqErr := http.NewRequest(method, endpoint, nil)
	if reqErr != nil {
		return nil, reqErr
	}
	// set all headers from client context
	err := setRequestHeadersFromContext(httpContext, httpRequest.Header)
	if err != nil {
		return nil, err
	}
	if len(httpRequest.Header["accept"]) == 0 && len(httpRequest.Header["Accept"]) == 0 {
		httpRequest.Header["Accept"] = []string{"application/json"}
	}
	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode == 204 {
		contentTypeOfResponse := extractContentType(httpResponse.Header.Get(contentTypeHeader))
		if contentTypeOfResponse == "" {
			httpResponse.Body.Close()
			response := new(DeleteTodos204Response)
			return response, nil
		}
		httpResponse.Body.Close()
		return nil, newNotSupportedContentType(415, contentTypeOfResponse)
	}

	if client.hooks.OnUnknownResponseCode != nil {
		message := client.hooks.OnUnknownResponseCode(httpResponse, httpRequest)
		httpResponse.Body.Close()
		return nil, errors.New(message)
	}
	httpResponse.Body.Close()
	return nil, newUnknownResponseError(httpResponse.StatusCode)
}

func (client *todoServiceClient) ListTodos(request *ListTodosRequest) (ListTodosResponse, error) {
	if request == nil {
		return nil, newRequestObjectIsNilError
	}
	path := "/todos"
	method := "GET"
	endpoint := client.baseURL + path
	httpContext := newHttpContextWrapper(client.ctx)
	httpRequest, reqErr := http.NewRequest(method, endpoint, nil)
	if reqErr != nil {
		return nil, reqErr
	}
	// set all headers from client context
	err := setRequestHeadersFromContext(httpContext, httpRequest.Header)
	if err != nil {
		return nil, err
	}
	if len(httpRequest.Header["accept"]) == 0 && len(httpRequest.Header["Accept"]) == 0 {
		httpRequest.Header["Accept"] = []string{"application/json"}
	}
	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode == 200 {
		contentTypeOfResponse := extractContentType(httpResponse.Header.Get(contentTypeHeader))
		if contentTypeOfResponse == contentTypeApplicationJson || contentTypeOfResponse == contentTypeApplicationHalJson {
			response := new(ListTodos200Response)
			decodeErr := json.NewDecoder(httpResponse.Body).Decode(&response.Body)
			httpResponse.Body.Close()
			if decodeErr != nil {
				return nil, decodeErr
			}
			return response, nil
		} else if contentTypeOfResponse == "" {
			httpResponse.Body.Close()
			response := new(ListTodos200Response)
			return response, nil
		}
		httpResponse.Body.Close()
		return nil, newNotSupportedContentType(415, contentTypeOfResponse)
	}

	if client.hooks.OnUnknownResponseCode != nil {
		message := client.hooks.OnUnknownResponseCode(httpResponse, httpRequest)
		httpResponse.Body.Close()
		return nil, errors.New(message)
	}
	httpResponse.Body.Close()
	return nil, newUnknownResponseError(httpResponse.StatusCode)
}

func (client *todoServiceClient) PostTodo(request *PostTodoRequest) (PostTodoResponse, error) {
	if request == nil {
		return nil, newRequestObjectIsNilError
	}
	path := "/todos"
	method := "POST"
	endpoint := client.baseURL + path
	httpContext := newHttpContextWrapper(client.ctx)
	jsonData := new(bytes.Buffer)
	encodeErr := json.NewEncoder(jsonData).Encode(&request.TodoPost)
	if encodeErr != nil {
		return nil, encodeErr
	}
	httpRequest, reqErr := http.NewRequest(method, endpoint, jsonData)
	if reqErr != nil {
		return nil, reqErr
	}
	httpRequest.Header[contentTypeHeader] = []string{contentTypeApplicationJson}
	// set all headers from client context
	err := setRequestHeadersFromContext(httpContext, httpRequest.Header)
	if err != nil {
		return nil, err
	}
	if len(httpRequest.Header["accept"]) == 0 && len(httpRequest.Header["Accept"]) == 0 {
		httpRequest.Header["Accept"] = []string{"application/json"}
	}
	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode == 201 {
		contentTypeOfResponse := extractContentType(httpResponse.Header.Get(contentTypeHeader))
		if contentTypeOfResponse == contentTypeApplicationJson || contentTypeOfResponse == contentTypeApplicationHalJson {
			response := new(PostTodo201Response)
			decodeErr := json.NewDecoder(httpResponse.Body).Decode(&response.Body)
			httpResponse.Body.Close()
			if decodeErr != nil {
				return nil, decodeErr
			}
			return response, nil
		} else if contentTypeOfResponse == "" {
			httpResponse.Body.Close()
			response := new(PostTodo201Response)
			return response, nil
		}
		httpResponse.Body.Close()
		return nil, newNotSupportedContentType(415, contentTypeOfResponse)
	}

	if client.hooks.OnUnknownResponseCode != nil {
		message := client.hooks.OnUnknownResponseCode(httpResponse, httpRequest)
		httpResponse.Body.Close()
		return nil, errors.New(message)
	}
	httpResponse.Body.Close()
	return nil, newUnknownResponseError(httpResponse.StatusCode)
}

func (client *todoServiceClient) DeleteTodo(request *DeleteTodoRequest) (DeleteTodoResponse, error) {
	if request == nil {
		return nil, newRequestObjectIsNilError
	}
	path := "/todos/{todoId}"
	method := "DELETE"
	endpoint := client.baseURL + path
	httpContext := newHttpContextWrapper(client.ctx)
	endpoint = strings.Replace(endpoint, "{todoId}", url.QueryEscape(toString(request.TodoId)), 1)
	httpRequest, reqErr := http.NewRequest(method, endpoint, nil)
	if reqErr != nil {
		return nil, reqErr
	}
	// set all headers from client context
	err := setRequestHeadersFromContext(httpContext, httpRequest.Header)
	if err != nil {
		return nil, err
	}
	if len(httpRequest.Header["accept"]) == 0 && len(httpRequest.Header["Accept"]) == 0 {
		httpRequest.Header["Accept"] = []string{"application/json"}
	}
	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode == 204 {
		contentTypeOfResponse := extractContentType(httpResponse.Header.Get(contentTypeHeader))
		if contentTypeOfResponse == "" {
			httpResponse.Body.Close()
			response := new(DeleteTodo204Response)
			return response, nil
		}
		httpResponse.Body.Close()
		return nil, newNotSupportedContentType(415, contentTypeOfResponse)
	}

	if httpResponse.StatusCode == 404 {
		contentTypeOfResponse := extractContentType(httpResponse.Header.Get(contentTypeHeader))
		if contentTypeOfResponse == "" {
			httpResponse.Body.Close()
			response := new(DeleteTodo404Response)
			return response, nil
		}
		httpResponse.Body.Close()
		return nil, newNotSupportedContentType(415, contentTypeOfResponse)
	}

	if client.hooks.OnUnknownResponseCode != nil {
		message := client.hooks.OnUnknownResponseCode(httpResponse, httpRequest)
		httpResponse.Body.Close()
		return nil, errors.New(message)
	}
	httpResponse.Body.Close()
	return nil, newUnknownResponseError(httpResponse.StatusCode)
}

func (client *todoServiceClient) GetTodo(request *GetTodoRequest) (GetTodoResponse, error) {
	if request == nil {
		return nil, newRequestObjectIsNilError
	}
	path := "/todos/{todoId}"
	method := "GET"
	endpoint := client.baseURL + path
	httpContext := newHttpContextWrapper(client.ctx)
	endpoint = strings.Replace(endpoint, "{todoId}", url.QueryEscape(toString(request.TodoId)), 1)
	httpRequest, reqErr := http.NewRequest(method, endpoint, nil)
	if reqErr != nil {
		return nil, reqErr
	}
	// set all headers from client context
	err := setRequestHeadersFromContext(httpContext, httpRequest.Header)
	if err != nil {
		return nil, err
	}
	if len(httpRequest.Header["accept"]) == 0 && len(httpRequest.Header["Accept"]) == 0 {
		httpRequest.Header["Accept"] = []string{"application/json"}
	}
	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode == 200 {
		contentTypeOfResponse := extractContentType(httpResponse.Header.Get(contentTypeHeader))
		if contentTypeOfResponse == contentTypeApplicationJson || contentTypeOfResponse == contentTypeApplicationHalJson {
			response := new(GetTodo200Response)
			decodeErr := json.NewDecoder(httpResponse.Body).Decode(&response.Body)
			httpResponse.Body.Close()
			if decodeErr != nil {
				return nil, decodeErr
			}
			return response, nil
		} else if contentTypeOfResponse == "" {
			httpResponse.Body.Close()
			response := new(GetTodo200Response)
			return response, nil
		}
		httpResponse.Body.Close()
		return nil, newNotSupportedContentType(415, contentTypeOfResponse)
	}

	if httpResponse.StatusCode == 404 {
		contentTypeOfResponse := extractContentType(httpResponse.Header.Get(contentTypeHeader))
		if contentTypeOfResponse == "" {
			httpResponse.Body.Close()
			response := new(GetTodo404Response)
			return response, nil
		}
		httpResponse.Body.Close()
		return nil, newNotSupportedContentType(415, contentTypeOfResponse)
	}

	if client.hooks.OnUnknownResponseCode != nil {
		message := client.hooks.OnUnknownResponseCode(httpResponse, httpRequest)
		httpResponse.Body.Close()
		return nil, errors.New(message)
	}
	httpResponse.Body.Close()
	return nil, newUnknownResponseError(httpResponse.StatusCode)
}

func (client *todoServiceClient) PatchTodo(request *PatchTodoRequest) (PatchTodoResponse, error) {
	if request == nil {
		return nil, newRequestObjectIsNilError
	}
	path := "/todos/{todoId}"
	method := "PATCH"
	endpoint := client.baseURL + path
	httpContext := newHttpContextWrapper(client.ctx)
	endpoint = strings.Replace(endpoint, "{todoId}", url.QueryEscape(toString(request.TodoId)), 1)
	jsonData := new(bytes.Buffer)
	encodeErr := json.NewEncoder(jsonData).Encode(&request.TodoPatch)
	if encodeErr != nil {
		return nil, encodeErr
	}
	httpRequest, reqErr := http.NewRequest(method, endpoint, jsonData)
	if reqErr != nil {
		return nil, reqErr
	}
	httpRequest.Header[contentTypeHeader] = []string{contentTypeApplicationJson}
	// set all headers from client context
	err := setRequestHeadersFromContext(httpContext, httpRequest.Header)
	if err != nil {
		return nil, err
	}
	if len(httpRequest.Header["accept"]) == 0 && len(httpRequest.Header["Accept"]) == 0 {
		httpRequest.Header["Accept"] = []string{"application/json"}
	}
	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode == 200 {
		contentTypeOfResponse := extractContentType(httpResponse.Header.Get(contentTypeHeader))
		if contentTypeOfResponse == contentTypeApplicationJson || contentTypeOfResponse == contentTypeApplicationHalJson {
			response := new(PatchTodo200Response)
			decodeErr := json.NewDecoder(httpResponse.Body).Decode(&response.Body)
			httpResponse.Body.Close()
			if decodeErr != nil {
				return nil, decodeErr
			}
			return response, nil
		} else if contentTypeOfResponse == "" {
			httpResponse.Body.Close()
			response := new(PatchTodo200Response)
			return response, nil
		}
		httpResponse.Body.Close()
		return nil, newNotSupportedContentType(415, contentTypeOfResponse)
	}

	if httpResponse.StatusCode == 404 {
		contentTypeOfResponse := extractContentType(httpResponse.Header.Get(contentTypeHeader))
		if contentTypeOfResponse == "" {
			httpResponse.Body.Close()
			response := new(PatchTodo404Response)
			return response, nil
		}
		httpResponse.Body.Close()
		return nil, newNotSupportedContentType(415, contentTypeOfResponse)
	}

	if client.hooks.OnUnknownResponseCode != nil {
		message := client.hooks.OnUnknownResponseCode(httpResponse, httpRequest)
		httpResponse.Body.Close()
		return nil, errors.New(message)
	}
	httpResponse.Body.Close()
	return nil, newUnknownResponseError(httpResponse.StatusCode)
}

type TodoServiceClient interface {
	DeleteTodos(request *DeleteTodosRequest) (DeleteTodosResponse, error)
	ListTodos(request *ListTodosRequest) (ListTodosResponse, error)
	PostTodo(request *PostTodoRequest) (PostTodoResponse, error)
	DeleteTodo(request *DeleteTodoRequest) (DeleteTodoResponse, error)
	GetTodo(request *GetTodoRequest) (GetTodoResponse, error)
	PatchTodo(request *PatchTodoRequest) (PatchTodoResponse, error)
}

func NewTodoServiceClient(httpClient *http.Client, baseUrl string, options Opts) TodoServiceClient {
	return &todoServiceClient{httpClient: newHttpClientWrapper(httpClient, baseUrl), baseURL: baseUrl, hooks: options.Hooks, ctx: options.Ctx, xmlMatcher: regexp.MustCompile("^application\\/(.+)xml$")}
}