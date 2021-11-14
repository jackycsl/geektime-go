1. 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

wire generate 时遇到 error, 暂时用手写， 请问老师可以帮忙看为什么会有 error

```go
    wire: /Users/JackyCSL/Developer/my_workspace/golang/geektime-go/chp4/wire.go:13:1: inject InitArticleService: no provider found for github.com/jackycsl/geekbang-go/chp4/internal/biz.ArticleRepo
        needed by *github.com/jackycsl/geekbang-go/chp4/internal/biz.ArticleUseCase in provider "NewArticleUseCase" (/Users/JackyCSL/Developer/my_workspace/golang/geektime-go/chp4/internal/biz/article.go:25:6)
        needed by *github.com/jackycsl/geekbang-go/chp4/internal/service.ArticleService in provider "NewArticleService" (/Users/JackyCSL/Developer/my_workspace/golang/geektime-go/chp4/internal/service/article.go:19:6)
        wire: github.com/jackycsl/geekbang-go/chp4: generate failed
```
