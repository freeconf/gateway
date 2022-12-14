
# FreeCONF Gateway



<details><summary>API Usage Notes:</summary>

#### General API Usage Notes
* `DELETE` implementation may be disallowed or ignored depending on the context
* Lists use `../path={key}/...` instead of `.../path/key/...` to avoid API name collision

#### `GET` Query Parameters

These parameters can be combined.

> | param                            | description | example |
> |----------------------------------|-------------|---------|
> | `content=[non-config\|config]` | Show only read-only fields or only read/write fields |   `.../path?content=config`|
> | `fields=field1;field2` | Return a portion of the data limited to fields listed | `.../path?fields=user%2faddress` |
> | `depth=n` | Return a portion of the data limited to depth of the hierarchy | `.../path?depth=1`
> | `fc.xfields=field1;fields` | Return a portion of the data excluding the fields listed | `.../path?fc.xfields=user%2faddress` |
> | `fc.range=field!{startRow}-[{endRow}]` | For lists, return only limited number of rows | `.../path?fc.range=user!10-20` 

</details>





<details>
 <summary><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:fc-gateway</b></code> </summary>

#### fc-gateway


**GET Response Data**
````json
{
  "registration":[{
     "deviceId":"",
     "address":""
  }]}
````



**Data Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | registration.deviceId | string  |   | r/o |
> | registration.address | string  |   | r/o |

**Responses**
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

**HTTP response codes**
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

**Examples**
````bash
# retrieve data
curl https://server/restconf/data/acc:fc-gateway

# update existing data
curl -X PUT -d @data.json https://server/restconf/data/acc:fc-gateway

# create new data
curl -X POST -d @data.json https://server/restconf/data/acc:fc-gateway

# delete current data
curl -X DELETE https://server/restconf/data/acc:fc-gateway
````
</details>





<details>
 <summary><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:registration</b></code> </summary>

#### registration


**GET Response Data**
````json
{"registration":[
  "deviceId":"",
  "address":""}, {"..."}]}
````



**Data Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | deviceId | string  |   | r/o |
> | address | string  |   | r/o |

**Responses**
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

**HTTP response codes**
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

**Examples**
````bash
# retrieve data
curl https://server/restconf/data/acc:registration

# update existing data
curl -X PUT -d @data.json https://server/restconf/data/acc:registration

# create new data
curl -X POST -d @data.json https://server/restconf/data/acc:registration

# delete current data
curl -X DELETE https://server/restconf/data/acc:registration
````
</details>




<details>
 <summary><code>[GET|PUT|POST|DELETE]</code> <code><b>restconf/data/acc:registration={deviceId}</b></code> </summary>

#### registration={deviceId}


**GET Response Data**
````json
{
  "deviceId":"",
  "address":""}
````



**Data Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | deviceId | string  |   | r/o |
> | address | string  |   | r/o |

**Responses**
> | http method  |  request body  | response body |
> |--------------|----------------|---------------|
> | `POST`       |  *JSON data*   | - none -      |
> | `PUT`       |  *JSON data*   | - none -      |
> | `GET`       |  - none -      | *JSON data*   |
> | `DELETE`     |  - none -      | - none -      |

**HTTP response codes**
> | http code |  reason for code    |
> |-----------|---------------------|
> | 200       | success             |
> | 401       | not authorized      |
> | 400       | invalid request     |
> | 404       | data does not exist |
> | 500       | internal error      |

**Examples**
````bash
# retrieve data
curl https://server/restconf/data/acc:registration={deviceId}

# update existing data
curl -X PUT -d @data.json https://server/restconf/data/acc:registration={deviceId}

# create new data
curl -X POST -d @data.json https://server/restconf/data/acc:registration={deviceId}

# delete current data
curl -X DELETE https://server/restconf/data/acc:registration={deviceId}
````
</details>






  <details>
 <summary><code>[GET]</code> <code><b>restconf/data/acc:update</b></code> </summary>

#### update

**Response Stream** [SSE Format](https://html.spec.whatwg.org/multipage/server-sent-events.html#server-sent-events)

````
data: {first JSON message all on one line followed by 2 CRLFs}

data: {next JSON message with same format all on one line ...}
````

Each JSON message would have following data
````json
{
  "deviceId":"",
  "address":""
}
````

**Response Body Details**

> | field   |  type  |  Description |  Details |
> |---------|--------|--------------|----------|
> | deviceId | string  |   |  |
> | address | string  |   |  |

**Example**
````bash
# retrieve data stream, adjust timeout for slower streams
curl -N https://server/restconf/data/acc:update
````

</details>
  

