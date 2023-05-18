# API文档

### 登录 POST /api/v1/login
发送参数  

username string 必填   用户名        
password string 必填   密码

返回参数

code int           100表示登录成功 其他表示错误
detail string

### 注册 POST /api/v1/register
发送参数

username string 必填 用户名
password string 必填 密码
email    string 选填 邮箱
phonenum string 选填 手机号

返回参数
code int
detail string