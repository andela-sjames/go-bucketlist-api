# go-bucketlist-api

A simple app to demonstrate the power of Go in creating a RESTFUL API

This is an API that allows users to create bucketlists and add items to their bucketlists.Authentication via JWT is used to ensure security, allowing Users interaction only when there are authenticated

```text
Bearer Token
e.g. Bearer aaaaaaaaa.vbbbbbbbb.cvvvvvvvvv
```

EndPoint |Functionality|Public Access
---------|-------------|--------------
POST /|Home/Welcome route|TRUE
POST /api/auth/signup|User sign up|TRUE
POST api/auth/login|Logs a user in|TRUE
GET api/auth/refresh| refresh a token with 60s left to expire|FALSE
POST api/bucketlists/|Create a new bucket list|FALSE
GET api/bucketlists/|List all the created bucket lists|FASLE
GET api/bucketlists/id|Get single bucket list|FALSE
PUT api/bucketlists/id|Update this bucket list|FALSE
DELETE api/bucketlists/id|Delete this single bucket list|FALSE
POST api/bucketlists/id/items/|Create a new item in bucket list|FALSE
PUT api/bucketlists/id/items/item_id|Update a bucket list item|FALSE
DELETE api/bucketlists/id/items/item_id|Delete an item in a bucket list|FALSE

## **__RESOURCES__**

**__AUTH__**

POST api/auth/signup | Register a new user
```Parameters/Input data: {"username":"testuser", "password": "password"}```

POST api/auth/login | User login
```Parameters/Input data: {{"username":"testuser", "password": "password"}```

POST api/auth/refresh | User refresh token
```Parameters/Input data: nil```

**BUCKETLIST** __url data__: __id__ = __bucketlist id__

POST api/bucketlists/  | create a new bucketlist
```Parameters/Input data: {"name":"name of bucketlist"}```

GET api/bucketlists/ | List all the created bucket lists
```Parameters/Input data: nil```

GET api/bucketlists/id | Get single bucket list
```Parameters/Input data: nil```

PUT api/bucketlists/id | Update this bucket list
```Parameters/Input data: {"name":"update bucketlist name"}```

DELETE api/bucketlists/id | Delete this single bucket list
```Parameters/Input data: nil```

**__ITEMS__**  __url data__:__id__ = __bucketlist__ __id__, __item_id__ = __item id__

POST api/bucketlists/id/items/ | Create a new item in a bucket list
```Parameters/Input data: {"name":"my bucketlistitem", "done":false }``` 

PUT api/bucketlists/id/items/item_id | Update a bucket list item
```Parameters/Input data: {"name":"update my bucketlistitem", "done":true }```

DELETE api/bucketlists/id/items/item_id | Delete an item in a bucket list
```Parameters/Input data: nil```

## WORKFLOW

User registers via ```POST api/auth/register``` route and is given a ```jwt token```

User uses jwt token which expires after a period of 24 hour.
User obtains jwt token if he/she wants to continue using  API services.
User uses jwt token to make request to server for resources defined above.

RESTful API is STATELESS and so no user session is stored.

## OTHER FEATURES COMING SOON / TODO

User can search for bucketlist using limits and page(pagination via limit)
```GET /bucketlists/limit=2``` and ```GET /bucketlists/limit=4&page=1```
Default limit without specification is 20 and default page number is 1.
```GET /bucketlists/id/limit=5``` and ```GET /bucketlists/id/limit=4&page=2```

User can also search for bucketlist using search parameter q.
```GET /bucketlists/q=create``` and  ```GET /bucketlists/q=game&limit=4&page=1```

### USAGE

Be sure you have docker and docker-compose installed on your machine.
Run the command `docker-compose up` to spin up the server and navigate the the `localhost:8000/` to view the welcome route.

Test Api using ```POSTMAN```  or ```cURL```

## TODO

- Paginate GetAllBucketlistHandler
- Paginate GetBucketByIDlistHandler items
- Document API

## Deploy to Kubernetes using minikube

k8bucketapi folder reference: [README.md](k8bucketapi/README.md)
