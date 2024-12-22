
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

const openAIURL = "https://api.openai.com/v1/chat/completions"
const apiKey = "YOUR_OPENAI_API_KEY" // Thay bằng API key của bạn

func translateEPUB(content string) (string, error) {
    // Tạo payload
    payload := map[string]interface{}{
        "model": "gpt-4",
        "messages": []map[string]string{
            {"role": "system", "content": "Translate EPUB content to another language."},
            {"role": "user", "content": content},
        },
    }
    body, _ := json.Marshal(payload)

    // Gửi request
    req, err := http.NewRequest("POST", openAIURL, bytes.NewBuffer(body))
    if err != nil {
        return "", err
    }
    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

    // Nhận response
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // Đọc response
    responseBody, _ := ioutil.ReadAll(resp.Body)
    var result map[string]interface{}
    json.Unmarshal(responseBody, &result)

    // Trích xuất nội dung dịch
    if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
        if message, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{}); ok {
            return message["content"].(string), nil
        }
    }
    return "", fmt.Errorf("unexpected response format")
}

func main() {
    fmt.Println("EPUB Translator using OpenAI API")
    content := "This is a sample EPUB content to be translated."
    translatedContent, err := translateEPUB(content)
    if err != nil {
        fmt.Println("Error translating EPUB:", err)
    } else {
        fmt.Println("Translated Content:", translatedContent)
    }
}
