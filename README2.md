conf
===

[![GoDoc](https://godoc.org/github.com/alexlyulkov/conf?status.svg)](https://godoc.org/github.com/alexlyulkov/conf)

Package conf Implement a http service that provides an API to an arbitrary configuration
represented by a hierarchical data structure.

Nodes can be assigned via dot-separated names. Also it supports getting or setting an entire subtree using maps encoded in JSON.

API
===

Node full name should be a string with dot-separated names of the node parents and the node.
Each name should consist only of english letters and numbers.

Each node value should be a valid json, consisting only of maps and strings.

***Insert***

Insert creates the node and all the subnodes described in the value parameter.
If the node already exists, it returns an error.

http://server_address/insert

POST parameters:
- name: dot-separated name of the node.
- value: json value of the node.

python example :

```python
import requests

requests.post("http://server_address/insert", data = {
    "name":"subtree1.item1",
    "value":'"value1"'
})
requests.post("http://server_address/insert", data = {
    "name":"subtree2",
    "value":'{
        "subtree21":{
            "item2":"value2",
            "item3":"value3"
        },
        "subtree22":{
            "item4":"value4"
        }
    }'
})
```

