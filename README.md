# GO-RESTful

This just a a simple RESTful API using GO and MySQL. I create simple operation Create, Read, Update, and Delete (CRUD) operation towards User entity. User struct looks like this:

```go
type User struct {
	ID        int64  `json:"Id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  []byte `json:"-"`
}
```

and these the routes:
* *users*: Method: GET, POST
* *users/1*: Method: GET, PUT, DELETE

I used these packages:
* *github.com/julienschmidt/httprouter* for routing, because it's simple and ease to understand
* *golang.org/x/crypto/bcrypt* for encrypting password

