1.我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答：应该 wrap error， 抛给上层。 dao 层管理数据库，应该如实报告情况，然后交给上层比如 biz 层， 让 biz 层去统一处理 business logic。

example code:

- dao 层

```go

  type User struct {
    Id int
    Name string
  }

  var ErrRecordNotFound = errors.New("record not found")

  func getUserById(id int64) (*User, error) {

    // user struct
    var user User

    // SQL query
    query := `SELECT id, name
    FROM users
    WHERE id= $1
    `

    err := r.DB.Query(query, id).Scan(&user.ID, &user.Name)

    if err != nil {
      switch {
      case errors.Is(err, sql.ErrNoRows):
        return nil, errors.Wrapf(ErrRecordNotFound, "data not found")
      default:
        return nil, errors.Wrapf(err, "db query system error")
      }
    }
    return &user, nil
  }
```

- biz 层

```go
  data, err = getUserById(id)
  if err != nil {
    switch {
    case errors.Is(err, ErrRecordNotFound) {
      // depends on business case, may return nil or return 404 not found.
      return 404
    default:
      // handle error
    }
    }
  }
```
