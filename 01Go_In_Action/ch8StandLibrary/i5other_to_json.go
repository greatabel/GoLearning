package main

import (
    "encoding/json"
    "fmt"
    "log"
)

func main() {
    c := make(map[string]interface{})
    c["name"] = "Gopher"
    c["title"] = "programmer"
    c["contact"] = map[string]interface{}{
        "home": "415.333.3333",
        "cell": "415.555.5555",
    }

        // Marshal the map into a JSON string.
    data, err := json.MarshalIndent(c, "", "    ")
    if err != nil {
        log.Println("ERROR:", err)
        return
    }
    fmt.Println("##数据转化为:")
    fmt.Println(string(data))
}