package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func resourceCreate(ctx context.Context, meta interface{}, url string, in, out interface{}) error {
	reqBody, err := json.Marshal(&in)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	log.Printf("POST %s: %s", url, string(reqBody))
	res, err := meta.(*Client).Post(ctx, url, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer func() {
		// Keep-Alive.
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("POST %s returned %d: %s", res.Request.URL.String(), res.StatusCode, string(body))
	}
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	log.Printf("POST %s returned %d: %s", res.Request.URL.String(), res.StatusCode, string(body))
	if err := json.Unmarshal(body, &out); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func resourceRead(ctx context.Context, meta interface{}, url string, out interface{}) (err error, ok bool) {
	log.Printf("GET %s", url)
	res, err := meta.(*Client).Get(ctx, url)
	if err != nil {
		return fmt.Errorf("%s", err), false
	}
	defer func() {
		// Keep-Alive.
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()
	if res.StatusCode == http.StatusNotFound {
		return nil, false
	}
	body, err := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s returned %d: %s", res.Request.URL.String(), res.StatusCode, string(body)), false
	}
	if err != nil {
		return fmt.Errorf("%s", err), false
	}
	log.Printf("GET %s returned %d: %s", res.Request.URL.String(), res.StatusCode, string(body))
	err = json.Unmarshal(body, &out)
	if err != nil {
		return fmt.Errorf("%s", err), false
	}
	return nil, true
}

func resourceUpdate(ctx context.Context, meta interface{}, url string, req interface{}, out interface{}) error {
	reqBody, err := json.Marshal(&req)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	log.Printf("PATCH %s: %s", url, string(reqBody))
	res, err := meta.(*Client).Patch(ctx, url, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer func() {
		// Keep-Alive.
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("PATCH %s returned %d: %s", res.Request.URL.String(), res.StatusCode, string(body))
	}
	log.Printf("PATCH %s returned %d: %s", res.Request.URL.String(), res.StatusCode, string(body))
	if err := json.Unmarshal(body, &out); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}
func resourceDelete(ctx context.Context, meta interface{}, url string) error {
	log.Printf("DELETE %s", url)
	res, err := meta.(*Client).Delete(ctx, url)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer func() {
		// Keep-Alive.
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusNotFound {
		return fmt.Errorf("DELETE %s returned %d: %s", res.Request.URL.String(), res.StatusCode, string(body))
	}
	log.Printf("DELETE %s returned %d: %s", res.Request.URL.String(), res.StatusCode, string(body))
	return nil
}
