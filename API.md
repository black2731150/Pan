# API文档

HTTP    方法使用规则  
GET     读取和检索操作
POST    新增和新建操作
PUT     更新操作，用户更新一个完整的资源
PATCH   更新操作，用于更新某一个资源的一个组成成分。
DELETE  删除的操作

## 登录 POST /api/v1/login
发送参数    
username string 必填   用户名          
password string 必填   密码  

返回参数  
code    int    0表示成功 其他表示错误  
message string 响应消息类型，如果是错误，那就是错误类型  
details string 如果发生了错误，有关错误的细节  

## 注册 POST /api/v1/register
发送参数  
username string 必填 用户名  (校验过的)  
password string 必填 密码    (6<=长度<=20>)  
email    string 必填 邮箱    (校验过的)  
phonenum string 选填 手机号     

返回参数  
code    int    0表示成功 其他表示错误  
message string 响应消息类型，如果是错误，那就是错误类型  
details string 如果发生了错误，有关错误的细节  

## 检测用户名是否已经存在 POST /api/v1/testUserName
发送参数  
username string 必填 用户名

返回参数  
code    int    0表示成功 其他表示错误  
message string 响应消息类型，如果是错误，那就是错误类型  
details string 如果发生了错误，有关错误的细节  

## 检测邮箱注册验证码是否正确 POST /api/v1/testEmailCode
发送参数  
email string 必填 邮箱 (正则匹配已经通过的)  
code  string 必填 验证码 (必须是六位)  

返回参数  
code    int    0表示成功 其他表示错误   
message string 响应消息类型，如果是错误，那就是错误类型    
details string 如果发生了错误，有关错误的细节


## 邮箱登录 POST /api/v1/loginWithEmail
发送参数
email string 必填 邮箱(正则匹配已经通过的)

返回参数
code    int    0表示成功 其他表示错误   
message string 响应消息类型，如果是错误，那就是错误类型    
details string 如果发生了错误，有关错误的细节
