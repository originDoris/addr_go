# 收货地址智能解析

## 1. 使用说明

```
 go version go1.20.5 darwin/arm64
```

- 索引树自动生成

  - ```shell-script
    cd /generate
    make all
    如需更新最新的行政信息，请自行修改/data下的json数据然后重新执行一遍以上命令
    ```
  
- 使用Demo
```go
package main

import (
	"fmt"

	"github.com/pupuk/addr"
)

func main() {
	parse := addr.Smart("张三 13800138000 龙华区龙华街道1980科技文化产业园3栋308 身份证120113196808214821")

	// 输出解析结果
	fmt.Println(parse.Name)     // 张三
	fmt.Println(parse.IdNumber) // 120113196808214821
	fmt.Println(parse.Mobile)   // 13800138000
	fmt.Println(parse.PostCode) // 570100
	fmt.Println(parse.Province) // 广东省
	fmt.Println(parse.City)     // 深圳市
	fmt.Println(parse.Region)   // 龙华区
	fmt.Println(parse.Street)   // 龙华街道1980科技文化产业园3栋317
	fmt.Println(parse.Address)  // 深圳市龙华区龙华街道1980科技文化产业园3栋317
}

```

- 构建HTTP服务端DEMO（使用fiber，gin类似，几行代码即可构建一个解析服务）
```go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pupuk/addr"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		parse := addr.Smart(c.Query("addr"))
		return c.JSON(parse)
	})

	app.Listen(":3000")
}

```
