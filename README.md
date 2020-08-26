## go-shop

一个基本购物网站的基础API项目

## 技术栈

gin + gorm + casbin + swagger + validation

## 目录结构

```bash
├─go-shop	  	        （项目名）
    │  ├─config         （配置包）
    │  ├─docs  	        （swagger文档目录）
    │  ├─global         （全局对象）
    │  ├─initialiaze    （初始化）
    │  ├─middleware     （中间件）
    │  ├─model          （数据库模型层）
    │  ├─pkg            （公共功能）
    │  ├─router         （路由）
    |  ├─runtime         (运行时产生的文件，如日志、上传等)
    │  ├─service         (业务逻辑层)
```

## 启动项目

1. 配置数据库，打开 config/config.yaml。

``` yaml
mysql:
    Type: 'mysql'
    User: '' 		// 你的用户名
    Password: '' 	// 你的密码
    Host: '127.0.0.1:3306'
    Name: ''  			// 你的数据库名
    TablePrefix:
```

2. 配置邮箱。

``` yaml
email:
    HostName: '' // 你的邮箱名
    Password: '' // 邮箱授权码
    Host: 'smtp.qq.com'
    Port: 465
```

3. 配置alipay

``` yaml
alipay:
    AppID: ''      // 应用AppID
    PrivateKey: '' // 应用私钥
    AppPublicCertPath: 'config/alipay/appCertPublicKey_2016092500592295.crt'
    AliPayPublicCertPath: 'config/alipay/alipayCertPublicKey_RSA2.crt'
    AliPayRootCertPath: 'config/alipay/alipayRootCert.crt'
```

4. 在项目根目录下执行。

``` bash
go run main.go
```



