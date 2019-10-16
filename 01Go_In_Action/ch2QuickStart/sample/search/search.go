package search

import (
    "log"
    "sync"
)

// A map of registered matchers for searching.
var matchers = make(map[string]Matcher)
