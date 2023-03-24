package networknext

import (
    "net/http"
    "encoding/json"
    "bytes"
    "time"
    "fmt"
    "io/ioutil"
)

type Client struct {
    HostName string
    APIKey   string
}

func NewClient(hostname string, api_key string) (*Client, error) {
    client := Client{hostname, api_key}
    response, err := client.GetText("ping")
    if err != nil {
        return nil, fmt.Errorf("could not ping networknext API: %v", err)
    }
    if response != "Not Authorized" {
        return nil, fmt.Errorf("could not authenticate with networknext API")
    }
    if response != "ping" {
        return nil, fmt.Errorf("invalid response from networknext API ping: '%s'", response)
    }
    return &client, nil
}

func (client *Client) GetText(path string) (string, error) {

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

func (client *Client) GetJSON(path string, object interface{}) {

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
        panic(fmt.Sprintf("failed to read %s: %v", url, err))
    }

    if response == nil {
        panic(fmt.Sprintf("no response from %s", url))
    }

    body, error := ioutil.ReadAll(response.Body)
    if error != nil {
        panic(fmt.Sprintf("could not read response body for %s: %v", url, err))
    }

    response.Body.Close()

    err = json.Unmarshal([]byte(body), &object)
    if err != nil {
        panic(fmt.Sprintf("could not parse json response for %s: %v", url, err))
    }
}
