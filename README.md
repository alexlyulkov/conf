conf
===

[![GoDoc](https://godoc.org/github.com/alexlyulkov/conf?status.svg)](https://godoc.org/github.com/alexlyulkov/conf)

Package conf implement a http service that provides an API to an arbitrary configuration
represented by a hierarchical data structure.

Nodes can be assigned via dot-separated names. Also service supports getting or setting an entire subtree using JSON encoded maps.

API
===

The node full name is a string with dot-separated names of the node and it's parents (e.g. `tree1.subtree1.item1`).
Each name should consist only of english letters and numbers.
Empty name corresponds to the tree root.

Each node value should be a valid json, consisting only of maps and strings.

Insert
---

Insert creates the node and all the subnodes described in the value parameter.
If the node already exists, it returns an error.

**http://server_address/insert**

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

Update
---

Update updates the specified node and all the specified subnodes described in the value parameter.
If the node or subnote doesn't exist, it returns an error.

**http://server_address/update**

POST parameters:
- name: dot-separated name of the node.
- value: json value of the node.

python example :

```python
import requests

requests.post("http://server_address/update", data = {
    "name":"subtree2",
    "value":'{
        "subtree21":{
            "item3":"value3"
        },
        "subtree22":{
            "item4":"value4"
        }
    }'
})
```

Read
---

Read returns the node and all the subnodes values.
If the node doesn't exists, it returns an error.
If you don't need all the subtrees you can specify the maximum depth.

**http://server_address/read**

POST parameters:
- name: dot-separated name of the node.
- depth: maximum subtree depth.

python example :

```python
import requests

print requests.post("http://server_address/read", data = {
    "name":"subtree1.item1"
}).content
# prints: "value1"

print requests.post("http://server_address/read", data = {
    "depth":"2"
}).content
# prints: {
#   "subtree1":{
#      "item1":"value1"
#   },
#   "subtree2":{ 
#      "subtree21":{
#      },
#      "subtree22":{
#      }
#   }
# }
```

Delete
---

Delete deletes the node and all the subnodes.
If the node doesn't exists, it returns an error.
If the name is empty, all the nodes will be deleted.

**http://server_address/delete**

POST parameters:
- name: dot-separated name of the node.

python example :

```python
import requests

requests.post("http://server_address/delete", data = {
    "name":"subtree2.subtree21"
})
```

Starting the server
===

Getting the sources:
`go get github.com/alexlyulkov/conf`

Installing:
`go install github.com/alexlyulkov/conf`

Starting the service:
`conf -address "host:port" -workdir "/dir1/dir2"`
