package networknext

import (
    "context"
    "net/http"
    "encoding/json"
    "bytes"
    "time"
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"

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
        return nil, fmt.Errorf("could not ping networknext API: %v", err)
    }
    if response == "Not Authorized" {
        return nil, fmt.Errorf("could not authenticate with networknext API")
    }
    if response != "pong" {
        return nil, fmt.Errorf("invalid response from networknext API ping: '%s'", response)
    }
    return &client, nil
}

func (client *Client) GetText(ctx context.Context, path string) (string, error) {

    url := client.HostName + "/" + path

    ctx = tflog.SetField(ctx, "networknext_url", url)
    ctx = tflog.SetField(ctx, "networknext_api_key", client.APIKey)

    tflog.Debug(ctx, "Network Next client GetText")

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

func (client *Client) Create(path string, object interface{}) (uint64, error) {

    url := client.HostName + "/" + path

    buffer := new(bytes.Buffer)

    json.NewEncoder(buffer).Encode(object)

    request, _ := http.NewRequest("POST", url, buffer)

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
        return 0, fmt.Errorf("create failed on %s: %v\n", url, err)
    }

    if response.StatusCode != 200 {
        return 0, fmt.Errorf("bad http response for %s: %d", url, response.StatusCode)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return 0, fmt.Errorf("could not read response for %s: %v", url, err)
    }

    id, err := strconv.ParseUint(string(body), 10, 64)
    if err != nil {
        return 0, fmt.Errorf("could not id response for %s: %v\n", url, err)
    }

    if id == 0 {
        return 0, fmt.Errorf("id returned from %s should be non-zero", url)
    }

    response.Body.Close()

    return id, nil
}

func (client *Client) Update(ctx context.Context, path string, object interface{}) error {

    url := client.HostName + "/" + path

    buffer := new(bytes.Buffer)

    json.NewEncoder(buffer).Encode(object)

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
        return fmt.Errorf("create failed on %s: %v\n", url, err)
    }

    if response.StatusCode != 200 {
        return fmt.Errorf("bad http response for %s: %d", url, response.StatusCode)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return fmt.Errorf("could not read response for %s: %v", url, err)
    }

    // todo: probably want to add a response here and parse it, so we can report errors on update
    _ = body

    response.Body.Close()

    return nil
}

func (client *Client) Delete(ctx context.Context, path string, id uint64) error {

    url := client.HostName + "/" + path

    request, _ := http.NewRequest("DELETE", url, strings.NewReader(fmt.Sprintf("%d", id)))

    request.Header.Set("Authorization", "Bearer " + client.APIKey)

    httpClient := &http.Client{}

    var err error
    var response *http.Response
    for i := 0; i < 5; i++ {
        response, err = httpClient.Do(request)
        if err == nil {
            body, err := ioutil.ReadAll(response.Body)
            if err != nil {
                panic(fmt.Sprintf("could not read delete response for %s: %v", url, err))
            }
            _ = body
            response.Body.Close()
            return nil
        }
        time.Sleep(time.Second)
    }

    return err
}
