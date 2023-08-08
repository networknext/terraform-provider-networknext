package networknext

import (
    "context"
    "net/http"
    "encoding/json"
    "bytes"
    "time"
    "fmt"
    "io/ioutil"

    "github.com/hashicorp/terraform-plugin-log/tflog"
)

type Client struct {
    HostName string
    APIKey   string
}

func NewClient(ctx context.Context, hostname string, api_key string) (*Client, error) {
    client := Client{hostname, api_key}
    response, err := client.GetText(ctx, "ping")
    if err != nil {
        return nil, fmt.Errorf("could not ping the network next API: %v", err)
    }
    if response == "Not Authorized" {
        return nil, fmt.Errorf("could not authenticate with the network next API")
    }
    if response != "pong" {
        return nil, fmt.Errorf("invalid response from network next API ping: '%s'", response)
    }
    return &client, nil
}

func (client *Client) GetText(ctx context.Context, path string) (string, error) {

    url := client.HostName + "/" + path

    ctx = tflog.SetField(ctx, "networknext_url", url)
    ctx = tflog.SetField(ctx, "networknext_api_key", client.APIKey)

    var err error
    var response *http.Response
    for i := 0; i < 5; i++ {
        req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
        req.Header.Set("Authorization", "Bearer " + client.APIKey)
        client := &http.Client{}
        response, err = client.Do(req)
        if err == nil {
            break
        }
        time.Sleep(time.Second)
    }

    if err != nil {
        return "", fmt.Errorf("failed to read %s: %v", url, err)
    }

    if response == nil {
        return "", fmt.Errorf("no response from %s", url)
    }

    if response.StatusCode != 200 {
        return "", fmt.Errorf("got %d response for %s", response.StatusCode, url)
    }

    body, error := ioutil.ReadAll(response.Body)
    if error != nil {
        return "", fmt.Errorf("could not read response body for %s: %v", url, err)
    }

    response.Body.Close()

    return string(body), nil
}

func (client *Client) GetJSON(ctx context.Context, path string, object interface{}) error {

    url := client.HostName + "/" + path

    var err error
    var response *http.Response
    for i := 0; i < 5; i++ {
        req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
        req.Header.Set("Authorization", "Bearer " + client.APIKey)
        client := &http.Client{}
        response, err = client.Do(req)
        if err == nil {
            break
        }
        time.Sleep(time.Second)
    }

    if err != nil {
        return fmt.Errorf("failed to read %s: %v", url, err)
    }

    if response == nil {
        return fmt.Errorf("no response from %s", url)
    }

    body, error := ioutil.ReadAll(response.Body)
    if error != nil {
        return fmt.Errorf("could not read response body for %s: %v", url, err)
    }

    response.Body.Close()

    tflog.Info(ctx, string(body))

    err = json.Unmarshal([]byte(body), &object)
    if err != nil {
        return fmt.Errorf("could not parse json response for %s: %v", url, err)
    }

    return nil
}

func (client *Client) Create(ctx context.Context, path string, requestData interface{}, responseData interface{}) error {

    url := client.HostName + "/" + path

    buffer := new(bytes.Buffer)

    json.NewEncoder(buffer).Encode(requestData)

    request, err := http.NewRequest("POST", url, buffer)
    if err != nil {
        return fmt.Errorf("could not create HTTP POST request for %s: %v", url, err)
    }

    request.Header.Set("Authorization", "Bearer " + client.APIKey)

    httpClient := &http.Client{}

    var response *http.Response
    for i := 0; i < 5; i++ {
        response, err = httpClient.Do(request)
        if err == nil {
            break
        }
        time.Sleep(time.Second)
    }

    if err != nil {
        return fmt.Errorf("create failed on %s: %v", url, err)
    }

    if response == nil {
        return fmt.Errorf("no response from %s", url)
    }

    body, error := ioutil.ReadAll(response.Body)
    if error != nil {
        return fmt.Errorf("could not read response body for %s: %v", url, err)
    }

    response.Body.Close()

    tflog.Info(ctx, string(body))

    err = json.Unmarshal([]byte(body), &responseData)
    if err != nil {
        return fmt.Errorf("could not parse json response for %s: %v", url, err)
    }

    return nil
}

func (client *Client) Update(ctx context.Context, path string, requestData interface{}, responseData interface{}) error {

    url := client.HostName + "/" + path

    buffer := new(bytes.Buffer)

    json.NewEncoder(buffer).Encode(requestData)

    request, _ := http.NewRequest("PUT", url, buffer)

    request.Header.Set("Authorization", "Bearer " + client.APIKey)

    httpClient := &http.Client{}

    var err error
    var response *http.Response
    for i := 0; i < 5; i++ {
        response, err = httpClient.Do(request)
        if err == nil {
            break
        }
        time.Sleep(time.Second)
    }

    if err != nil {
        return fmt.Errorf("failed to read %s: %v", url, err)
    }

    if response == nil {
        return fmt.Errorf("no response from %s", url)
    }

    body, error := ioutil.ReadAll(response.Body)
    if error != nil {
        return fmt.Errorf("could not read response body for %s: %v", url, err)
    }

    response.Body.Close()

    tflog.Info(ctx, string(body))

    err = json.Unmarshal([]byte(body), &responseData)
    if err != nil {
        return fmt.Errorf("could not parse json response for %s: %v", url, err)
    }

    return nil
}

func (client *Client) Delete(ctx context.Context, path string, responseData interface{}) error {

    url := client.HostName + "/" + path

    request, _ := http.NewRequest("DELETE", url, nil)

    request.Header.Set("Authorization", "Bearer " + client.APIKey)

    httpClient := &http.Client{}

    var err error
    var response *http.Response
    for i := 0; i < 5; i++ {
        response, err = httpClient.Do(request)
        if err == nil {
            break
        }
        time.Sleep(time.Second)
    }

    if err != nil {
        return fmt.Errorf("failed to read %s: %v", url, err)
    }

    if response == nil {
        return fmt.Errorf("no response from %s", url)
    }

    body, error := ioutil.ReadAll(response.Body)
    if error != nil {
        return fmt.Errorf("could not read response body for %s: %v", url, err)
    }

    response.Body.Close()

    tflog.Info(ctx, string(body))

    err = json.Unmarshal([]byte(body), &responseData)
    if err != nil {
        return fmt.Errorf("could not parse json response for %s: %v", url, err)
    }

    return err
}
