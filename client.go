package neura

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const DefaultEndpoint = "https://control.neura-os.com"

type Client struct {
	Endpoint string
	APIKey   string
	HTTP     *http.Client

	Memory *MemoryService
	Auth   *AuthService
}

type MemoryService struct {
	client *Client
}

type AuthService struct {
	client *Client
}

type Config struct {
	Endpoint string
	APIKey   string
}

func NewClient(cfg Config) (*Client, error) {
	apiKey := cfg.APIKey
	if apiKey == "" {
		apiKey = os.Getenv("NEURA_API_KEY")
	}
	// Allow empty API Key for registration flow

	endpoint := cfg.Endpoint
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}
	endpoint = strings.TrimRight(endpoint, "/")

	c := &Client{
		Endpoint: endpoint,
		APIKey:   apiKey,
		HTTP:     &http.Client{},
	}
	c.Memory = &MemoryService{client: c}
	c.Auth = &AuthService{client: c}
	return c, nil
}

func (c *Client) Decide(req DecisionRequest) (*DecisionResponse, error) {
	return c.post("/v1/decide", req)
}

func (c *Client) Validate(req DecisionRequest) (*ValidationResponse, error) {
	var result ValidationResponse
	err := c.doRequest("POST", "/v1/validate", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetDecision(id string) (*DecisionResponse, error) {
	var result DecisionResponse
	err := c.doRequest("GET", "/decision/"+id, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) WaitForDecision(id string, timeout time.Duration, interval time.Duration) (*DecisionResponse, error) {
	start := time.Now()
	for time.Since(start) < timeout {
		decision, err := c.GetDecision(id)
		if err != nil {
			return nil, err
		}
		if decision.Outcome == "ACT" || decision.Outcome == "DENY" {
			return decision, nil
		}
		time.Sleep(interval)
	}
	return nil, fmt.Errorf("timeout waiting for decision %s", id)
}

func (s *AuthService) Register(req AuthRequest) (*AuthResponse, error) {
	var result AuthResponse
	// Registration might not require auth headers if public/bootstrapped,
	// but doRequest adds them if APIKey is present.
	err := s.client.doRequest("POST", "/v1/auth/register", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *MemoryService) Store(req MemoryRequest) (*MemoryResponse, error) {
	var result MemoryResponse
	err := s.client.doRequest("POST", "/v1/memory", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *MemoryService) Search(req MemorySearchRequest) ([]MemoryResponse, error) {
	var result []MemoryResponse
	err := s.client.doRequest("POST", "/v1/memory/search", req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) post(path string, payload interface{}) (*DecisionResponse, error) {
	var result DecisionResponse
	err := c.doRequest("POST", path, payload, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) doRequest(method, path string, payload interface{}, target interface{}) error {
	var body io.Reader
	if payload != nil {
		jsonValue, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(jsonValue)
	}

	req, err := http.NewRequest(method, c.Endpoint+path, body)
	if err != nil {
		return err
	}

	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "neura-sdk-go/0.2.0")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("api request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	if target != nil {
		if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
			return err
		}
	}

	return nil
}
